package hedera

import (
	protobuf "github.com/golang/protobuf/proto"
	"github.com/hashgraph/hedera-sdk-go/proto"
)

// Transaction contains the protobuf of a prepared transaction which can be signed and executed.
type Query struct {
	pb *proto.Query
	pbHeader *proto.QueryHeader

	paymentTransactionID TransactionID
	nodeID AccountID
	maxQueryPayment Hbar
	queryPayment Hbar
	nextPaymentTransactionIndex int
	nextTransactionIndex int

	paymentTransactionNodeIDs []AccountID
	paymentTransactions   []*proto.Transaction

	fff string
}

func newQuery() Query {
	return Query{
		pb: &proto.Query{},
		pbHeader: &proto.QueryHeader{},

		paymentTransactionID:                   TransactionID{},
		nextTransactionIndex: 					0,
		paymentTransactions:         			make([]*proto.Transaction, 0),
		paymentTransactionNodeIDs:              make([]AccountID, 0),
	}
}

func (query *Query) SetNodeId(accoundID AccountID) *Query {
	query.nodeID = accoundID
	return query
}

func (query *Query) getNodeId(client *Client,) AccountID {
	if query.paymentTransactionNodeIDs != nil {
		return query.paymentTransactionNodeIDs[query.nextPaymentTransactionIndex]
	}

	if query.nodeID.isZero() {
		return query.nodeID
	} else {
		return client.getNextNode()
	}
}

func (query *Query) SetQueryPayment(queryPayment Hbar) *Query {
	query.queryPayment = queryPayment
	return query
}

func (query *Query) SetMaxQueryPayment(queryMaxPayment Hbar) *Query {
	query.maxQueryPayment= queryMaxPayment
	return query
}

func (query *Query) IsPaymentRequired() bool {
	return true
}

func query_getNodeId(
	request request,
	client *Client,
) AccountID {
	return request.query.nodeID
}

func query_advanceRequest(request request) {
	request.query.nextPaymentTransactionIndex++
}

func query_makeRequest(request request) protoRequest {
	return protoRequest{
		query: request.query.pbHeader.,
		//paymentTransactions[request.query.nextPaymentTransactionIndex],
	}
}

func query_mapResponseHeader(res response) protoResponseHeader {
	return protoResponseHeader{
		responseHeader: proto.ResponseHeader{
			NodeTransactionPrecheckCode: res.query.GetResponse(),
		}
	}
}

func query_mapResponse(request request, _ response, protoRequest protoRequest) (intermediateResponse, error) {
	hash, err := protobuf.Marshal(protoRequest.query)
	if err != nil {
		return intermediateResponse{}, err
	}

	return intermediateResponse{
		query: &proto.Response{},
	}, nil
}

func query_mapResponseStatus(
	_ request,
	response response,
) Status {
	return Status(response.query.)
}

func (query *Query) getTransactionID() TransactionID {
	return query.paymentTransactionID
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//func (transaction *Transaction) UnmarshalBinary(txBytes []byte) error {
//	transaction.transactions = make([]*proto.Transaction, 0)
//	transaction.transactions = append(transaction.transactions, &proto.Transaction{})
//	if err := protobuf.Unmarshal(txBytes, transaction.transactions[0]); err != nil {
//		return err
//	}
//
//	var txBody proto.TransactionBody
//	if err := protobuf.Unmarshal(transaction.transactions[0].GetBodyBytes(), &txBody); err != nil {
//		return err
//	}
//
//	transaction.id = transactionIDFromProto(txBody.TransactionID)
//
//	return nil
//}
//
//func TransactionFromBytes(bytes []byte) Transaction {
//	tx := Transaction{}
//	(&tx).UnmarshalBinary(bytes)
//	return tx
//}
//
//func (transaction *Transaction) initFee(client *Client) {
//	if client != nil && transaction.pbBody.TransactionFee == 0 {
//		transaction.SetMaxTransactionFee(client.maxTransactionFee)
//	}
//}
//
//func (transaction *Transaction) initTransactionID(client *Client) error {
//	if transaction.pbBody.TransactionID == nil {
//		if client.operator != nil {
//			transaction.id = NewTransactionID(client.operator.accountID)
//			transaction.SetTransactionID(transaction.id)
//		} else {
//			return errNoClientOrTransactionID
//		}
//	}
//
//	return nil
//}
//
//func (transaction *Transaction) isFrozen() bool {
//	return len(transaction.transactions) > 0
//}
//
//func transaction_freezeWith(
//	transaction *Transaction,
//	client *Client,
//) error {
//	if transaction.pbBody.TransactionID != nil && transaction.pbBody.NodeAccountID != nil {
//		bodyBytes, err := protobuf.Marshal(transaction.pbBody)
//		if err != nil {
//			// This should be unreachable
//			// From the documentation this appears to only be possible if there are missing proto types
//			panic(err)
//		}
//
//		sigmap := proto.SignatureMap{
//			SigPair: make([]*proto.SignaturePair, 0),
//		}
//		transaction.signatures = append(transaction.signatures, &sigmap)
//		transaction.transactions = append(transaction.transactions, &proto.Transaction{
//			BodyBytes: bodyBytes,
//			SigMap: &sigmap,
//		})
//
//		return nil
//	}
//
//	if transaction.pbBody.TransactionID != nil && len(transaction.nodeIDs) > 0 {
//		for _, id := range transaction.nodeIDs {
//			transaction.pbBody.NodeAccountID = id.toProtobuf()
//			bodyBytes, err := protobuf.Marshal(transaction.pbBody)
//			if err != nil {
//				// This should be unreachable
//				// From the documentation this appears to only be possible if there are missing proto types
//				panic(err)
//			}
//
//			sigmap := proto.SignatureMap{
//				SigPair: make([]*proto.SignaturePair, 0),
//			}
//			transaction.signatures = append(transaction.signatures, &sigmap)
//			transaction.transactions = append(transaction.transactions, &proto.Transaction{
//				BodyBytes: bodyBytes,
//				SigMap: &sigmap,
//			})
//		}
//
//		return nil
//	}
//
//	if client != nil && transaction.pbBody.TransactionID != nil {
//		size := client.getNumberOfNodesForTransaction()
//
//		for index := 0; index < size; index++ {
//			node := client.getNextNode()
//
//			transaction.nodeIDs = append(transaction.nodeIDs, node)
//
//			transaction.pbBody.NodeAccountID = node.toProtobuf()
//			bodyBytes, err := protobuf.Marshal(transaction.pbBody)
//			if err != nil {
//				// This should be unreachable
//				// From the documentation this appears to only be possible if there are missing proto types
//				panic(err)
//			}
//
//			sigmap := proto.SignatureMap{
//				SigPair: make([]*proto.SignaturePair, 0),
//			}
//			transaction.signatures = append(transaction.signatures, &sigmap)
//			transaction.transactions = append(transaction.transactions, &proto.Transaction{
//				BodyBytes: bodyBytes,
//				SigMap:    &sigmap,
//			})
//		}
//
//		return nil
//	}
//
//	return errNoClientOrTransactionIDOrNodeId
//}
//
//func (transaction *Transaction) keyAlreadySigned(
//	pk PublicKey,
//) bool {
//	if len(transaction.signatures) > 0 {
//		for _, pair := range transaction.signatures[0].SigPair {
//			if bytes.HasPrefix(pk.keyData, pair.PubKeyPrefix) {
//				return true
//			}
//		}
//	}
//
//	return false
//}
//
//func transaction_shouldRetry(_ request, status Status) bool {
//	return status == StatusBusy
//}
//
//func transaction_makeRequest(request request) protoRequest {
//	return protoRequest{
//		transaction: request.transaction.transactions[request.transaction.nextTransactionIndex],
//	}
//}
//
//func transaction_advanceRequest(request request) {
//	request.transaction.nextTransactionIndex++
//}
//
//func transaction_getNodeId(
//	request request,
//	client *Client,
//) AccountID {
//	return request.transaction.nodeIDs[request.transaction.nextTransactionIndex]
//}
//
//func transaction_mapResponseStatus(
//	_ request,
//	response response,
//) Status {
//	return Status(response.transaction.NodeTransactionPrecheckCode)
//}
//
//func transaction_mapResponse(request request, _ response, nodeID AccountID, protoRequest protoRequest) (intermediateResponse, error) {
//	hash, err := protobuf.Marshal(protoRequest.transaction)
//	if err != nil {
//		return intermediateResponse{}, err
//	}
//
//	return intermediateResponse{
//		transaction: TransactionResponse{
//			NodeID:        nodeID,
//			TransactionID: request.transaction.id,
//			Hash:          hash,
//		},
//	}, nil
//}
//
//func (transaction *Transaction) String() string {
//	return protobuf.MarshalTextString(transaction.transactions[0]) +
//		protobuf.MarshalTextString(transaction.body())
//}
//
//// MarshalBinary implements the encoding.BinaryMarshaler interface.
//func (transaction *Transaction) MarshalBinary() ([]byte, error) {
//	return protobuf.Marshal(transaction.transactions[0])
//}
//
//func (transaction *Transaction) ToBytes() ([]byte, error) {
//	return transaction.MarshalBinary()
//}
//
//// The protobuf stores the transaction body as raw bytes so we need to first
//// decode what we have to inspect the Kind, TransactionID, and the NodeAccountID so we know how to
//// properly execute it
//func (transaction *Transaction) body() *proto.TransactionBody {
//	transactionBody := new(proto.TransactionBody)
//	err := protobuf.Unmarshal(transaction.transactions[0].GetBodyBytes(), transactionBody)
//	if err != nil {
//		// The bodyBytes inside of the transaction at this point have been verified and this should be impossible
//		panic(err)
//	}
//
//	return transactionBody
//}
//
////
//// Shared
////
//
//func (transaction *Transaction) GetMaxTransactionFee() Hbar {
//	return HbarFromTinybar(int64(transaction.pbBody.TransactionFee))
//}
//
//// SetMaxTransactionFee sets the max transaction fee for this Transaction.
//func (transaction *Transaction) SetMaxTransactionFee(fee Hbar) *Transaction {
//	transaction.pbBody.TransactionFee = uint64(fee.AsTinybar())
//	return transaction
//}
//
//func (transaction *Transaction) GetTransactionMemo() string {
//	return transaction.pbBody.Memo
//}
//
//// SetTransactionMemo sets the memo for this Transaction.
//func (transaction *Transaction) SetTransactionMemo(memo string) *Transaction {
//	transaction.pbBody.Memo = memo
//	return transaction
//}
//
//func (transaction *Transaction) GetTransactionValidDuration() time.Duration {
//	if transaction.pbBody.TransactionValidDuration != nil {
//		return durationFromProto(transaction.pbBody.TransactionValidDuration)
//	} else {
//		return 0
//	}
//}
//
//// SetTransactionValidDuration sets the valid duration for this Transaction.
//func (transaction *Transaction) SetTransactionValidDuration(duration time.Duration) *Transaction {
//	transaction.pbBody.TransactionValidDuration = durationToProto(duration)
//	return transaction
//}
//
//func (transaction *Transaction) GetTransactionID() TransactionID {
//	if transaction.pbBody.TransactionID != nil {
//		return transactionIDFromProto(transaction.pbBody.TransactionID)
//	} else {
//		return TransactionID{}
//	}
//}
//
//// SetTransactionID sets the TransactionID for this Transaction.
//func (transaction *Transaction) SetTransactionID(transactionID TransactionID) *Transaction {
//	transaction.pbBody.TransactionID = transactionID.toProtobuf()
//	return transaction
//}
//
//func (transaction *Transaction) GetNodeID() AccountID {
//	if transaction.pbBody.NodeAccountID != nil {
//		return accountIDFromProto(transaction.pbBody.NodeAccountID)
//	} else {
//		return AccountID{}
//	}
//}
//
//// SetNodeID sets the node AccountID for this Transaction.
//func (transaction *Transaction) SetNodeID(nodeID AccountID) *Transaction {
//	transaction.pbBody.NodeAccountID = nodeID.toProtobuf()
//	transaction.nodeIDs = append(transaction.nodeIDs, nodeID)
//	return transaction
//}