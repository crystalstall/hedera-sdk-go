package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/handler"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/hashgraph/hedera-sdk-go/tck/methods"
	"github.com/hashgraph/hedera-sdk-go/tck/response"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv
	_ = godotenv.Load()

	// Initialize the SDK service
	sdkService := new(methods.SDKService)
	accountService := new(methods.AccountService)
	accountService.SetSdkService(sdkService)

	// Create a new RPC server
	assigner := handler.Map{
		"setup":                  errorHandlerWrapper(HandleError, handler.New(sdkService.Setup)),
		"reset":                  errorHandlerWrapper(HandleError, handler.New(sdkService.Reset)),
		"createAccount":          errorHandlerWrapper(HandleError, handler.New(accountService.CreateAccount)),
		"createAccountFromAlias": errorHandlerWrapper(HandleError, handler.New(accountService.CreateAccountFromAlias)),
		"getAccountInfo":         errorHandlerWrapper(HandleError, handler.New(accountService.GetAccountInfo)),
		"updateAccountMemo":      errorHandlerWrapper(HandleError, handler.New(accountService.UpdateAccountMemo)),
		"updateAccountKey":       errorHandlerWrapper(HandleError, handler.New(accountService.UpdateAccountKey)),
		"deleteAccount":          errorHandlerWrapper(HandleError, handler.New(accountService.DeleteAccount)),
		"generatePublicKey":      errorHandlerWrapper(HandleError, handler.New(methods.GeneratePublicKey)),
		"generatePrivateKey":     errorHandlerWrapper(HandleError, handler.New(methods.GeneratePrivateKey)),
	}

	bridge := jhttp.NewBridge(assigner, nil)

	// Listen and redirect to bridge
	http.HandleFunc("/", bridge.ServeHTTP)
	port := os.Getenv("TCK_PORT")
	if port == "" {
		port = "80"
	}
	fmt.Println("Server is listening on port: " + port)

	server := &http.Server{Addr: ":" + port}

	// Start the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Wait for the termination signal
	sig := <-signalCh
	fmt.Printf("Received signal: %v\n", sig)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %s\n", err)
	}

	fmt.Println("Server shutdown complete.")
}

// ErrorHandler is a function that handles errors reported by a method handler.
type ErrorHandler func(context.Context, *jrpc2.Request, error) error

func HandleError(_ context.Context, request *jrpc2.Request, err error) error {
	if err != nil {
		// jrpc generic error
		if jrpcError, ok := err.(*jrpc2.Error); ok {
			return jrpcError
		}
		// hedera specific errors
		if hederaErr, ok := err.(hedera.ErrHederaReceiptStatus); ok {
			return response.NewHederaReceiptError(hederaErr)
		}
		if hederaErr, ok := err.(hedera.ErrHederaPreCheckStatus); ok {
			return response.NewHederaPrecheckError(hederaErr)
		}
		// other errors
		return response.InternalError
	}
	return nil
}

// this wraps the jrpc2.Handler as it invokes the ErrorHandler func if error is returned
// essentially acts as a "postHandler"
func errorHandlerWrapper(eh ErrorHandler, h jrpc2.Handler) jrpc2.Handler {
	return func(ctx context.Context, req *jrpc2.Request) (any, error) {
		res, err := h(ctx, req)
		if err != nil {
			errorMessage := fmt.Sprintf("Error occurred processing JSON-RPC request: %s, Response error: %s", req, err)
			fmt.Println(errorMessage)
			return nil, eh(ctx, req, err)
		}
		return res, nil
	}
}