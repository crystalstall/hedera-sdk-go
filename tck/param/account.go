package param

import (
	"encoding/json"
)

// ---- Different methods params structs ----
type CreateAccountParams struct {
	PublicKey                     string      `json:"publicKey"`
	InitialBalance                int64       `json:"initialBalance"`
	ReceiverSignatureRequired     bool        `json:"receiverSignatureRequired"`
	MaxAutomaticTokenAssociations uint32      `json:"maxAutomaticTokenAssociations"`
	StakedAccountId               *string     `json:"stakedAccountId"`
	StakedNodeId                  json.Number `json:"stakedNodeId"`
	DeclineStakingReward          bool        `json:"declineStakingReward"`
	AccountMemo                   string      `json:"accountMemo"`
	AutoRenewPeriod               string      `json:"autoRenewPeriod"`
	PrivateKey                    string      `json:"privateKey"`
}

type AccountFromAliasParams struct {
	OperatorId     string `json:"operator_id"`
	AliasAccountId string `json:"aliasAccountId"`
	InitialBalance int64  `json:"initialBalance"`
}

type Alias string

type UpdateAccountKeyParams struct {
	AccountId     string `json:"accountId"`
	NewPublicKey  string `json:"newPublicKey"`
	OldPrivateKey string `json:"oldPrivateKey"`
	NewPrivateKey string `json:"newPrivateKey"`
}

type DeleteAccountParams struct {
	AccountId   string `json:"accountId"`
	AccountKey  string `json:"accountKey"`
	RecipientId string `json:"recipientId"`
}

type UpdateAccountParams struct {
	AccountId     string `json:"accountId"`
	NewPublicKey  string `json:"newPublicKey"`
	OldPrivateKey string `json:"oldPrivateKey"`
	NewPrivateKey string `json:"newPrivateKey"`
	Key           string `json:"key"`
	Memo          string `json:"memo"`
}