// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CryptoCreate.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Create a new account. After the account is created, the AccountID for it is in the receipt, or can be retrieved with a GetByKey query, or by asking for a Record of the transaction to be created, and retrieving that. The account can then automatically generate records for large transfers into it or out of it, which each last for 25 hours. Records are generated for any transfer that exceeds the thresholds given here. This account is charged cryptocurrency for each record generated, so the thresholds are useful for limiting Record generation to happen only for large transactions. The Key field is the key used to sign transactions for this account. If the account has receiverSigRequired set to true, then all cryptocurrency transfers must be signed by this account's key, both for transfers in and out. If it is false, then only transfers out have to be signed by it. When the account is created, the payer account is charged enough hbars so that the new account will not expire for the next autoRenewPeriod seconds. When it reaches the expiration time, the new account will then be automatically charged to renew for another autoRenewPeriod seconds. If it does not have enough hbars to renew for that long, then the remaining hbars are used to extend its expiration as long as possible. If it is has a zero balance when it expires, then it is deleted. This transaction must be signed by the payer account. If receiverSigRequired is false, then the transaction does not have to be signed by the keys in the keys field. If it is true, then it must be signed by them, in addition to the keys of the payer account.
//
// An entity (account, file, or smart contract instance) must be created in a particular realm. If the realmID is left null, then a new realm will be created with the given admin key. If a new realm has a null adminKey, then anyone can create/modify/delete entities in that realm. But if an admin key is given, then any transaction to create/modify/delete an entity in that realm must be signed by that key, though anyone can still call functions on smart contract instances that exist in that realm. A realm ceases to exist when everything within it has expired and no longer exists.
//
// The current API ignores shardID, realmID, and newRealmAdminKey, and creates everything in shard 0 and realm 0, with a null key. Future versions of the API will support multiple realms and multiple shards.
type CryptoCreateTransactionBody struct {
	Key                    *Key       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	InitialBalance         uint64     `protobuf:"varint,2,opt,name=initialBalance,proto3" json:"initialBalance,omitempty"`
	ProxyAccountID         *AccountID `protobuf:"bytes,3,opt,name=proxyAccountID,proto3" json:"proxyAccountID,omitempty"`
	SendRecordThreshold    uint64     `protobuf:"varint,6,opt,name=sendRecordThreshold,proto3" json:"sendRecordThreshold,omitempty"`
	ReceiveRecordThreshold uint64     `protobuf:"varint,7,opt,name=receiveRecordThreshold,proto3" json:"receiveRecordThreshold,omitempty"`
	ReceiverSigRequired    bool       `protobuf:"varint,8,opt,name=receiverSigRequired,proto3" json:"receiverSigRequired,omitempty"`
	AutoRenewPeriod        *Duration  `protobuf:"bytes,9,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	ShardID                *ShardID   `protobuf:"bytes,10,opt,name=shardID,proto3" json:"shardID,omitempty"`
	RealmID                *RealmID   `protobuf:"bytes,11,opt,name=realmID,proto3" json:"realmID,omitempty"`
	NewRealmAdminKey       *Key       `protobuf:"bytes,12,opt,name=newRealmAdminKey,proto3" json:"newRealmAdminKey,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}   `json:"-"`
	XXX_unrecognized       []byte     `json:"-"`
	XXX_sizecache          int32      `json:"-"`
}

func (m *CryptoCreateTransactionBody) Reset()         { *m = CryptoCreateTransactionBody{} }
func (m *CryptoCreateTransactionBody) String() string { return proto.CompactTextString(m) }
func (*CryptoCreateTransactionBody) ProtoMessage()    {}
func (*CryptoCreateTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_89b414cfdad72647, []int{0}
}

func (m *CryptoCreateTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoCreateTransactionBody.Unmarshal(m, b)
}
func (m *CryptoCreateTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoCreateTransactionBody.Marshal(b, m, deterministic)
}
func (m *CryptoCreateTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoCreateTransactionBody.Merge(m, src)
}
func (m *CryptoCreateTransactionBody) XXX_Size() int {
	return xxx_messageInfo_CryptoCreateTransactionBody.Size(m)
}
func (m *CryptoCreateTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoCreateTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoCreateTransactionBody proto.InternalMessageInfo

func (m *CryptoCreateTransactionBody) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CryptoCreateTransactionBody) GetInitialBalance() uint64 {
	if m != nil {
		return m.InitialBalance
	}
	return 0
}

func (m *CryptoCreateTransactionBody) GetProxyAccountID() *AccountID {
	if m != nil {
		return m.ProxyAccountID
	}
	return nil
}

func (m *CryptoCreateTransactionBody) GetSendRecordThreshold() uint64 {
	if m != nil {
		return m.SendRecordThreshold
	}
	return 0
}

func (m *CryptoCreateTransactionBody) GetReceiveRecordThreshold() uint64 {
	if m != nil {
		return m.ReceiveRecordThreshold
	}
	return 0
}

func (m *CryptoCreateTransactionBody) GetReceiverSigRequired() bool {
	if m != nil {
		return m.ReceiverSigRequired
	}
	return false
}

func (m *CryptoCreateTransactionBody) GetAutoRenewPeriod() *Duration {
	if m != nil {
		return m.AutoRenewPeriod
	}
	return nil
}

func (m *CryptoCreateTransactionBody) GetShardID() *ShardID {
	if m != nil {
		return m.ShardID
	}
	return nil
}

func (m *CryptoCreateTransactionBody) GetRealmID() *RealmID {
	if m != nil {
		return m.RealmID
	}
	return nil
}

func (m *CryptoCreateTransactionBody) GetNewRealmAdminKey() *Key {
	if m != nil {
		return m.NewRealmAdminKey
	}
	return nil
}

func init() {
	proto.RegisterType((*CryptoCreateTransactionBody)(nil), "proto.CryptoCreateTransactionBody")
}

func init() { proto.RegisterFile("CryptoCreate.proto", fileDescriptor_89b414cfdad72647) }

var fileDescriptor_89b414cfdad72647 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x15, 0x76, 0xd9, 0x5d, 0xbc, 0x28, 0xad, 0x8c, 0x84, 0xac, 0xc2, 0xa1, 0xea, 0xa1,
	0xca, 0x29, 0x42, 0x20, 0x55, 0x70, 0x6c, 0x9a, 0x4b, 0xd5, 0x4b, 0xe5, 0xf6, 0x05, 0x06, 0x7b,
	0xd4, 0x18, 0x52, 0x3b, 0x4c, 0x92, 0x96, 0xdc, 0x78, 0x74, 0x14, 0x27, 0x45, 0x90, 0x96, 0x53,
	0xe4, 0xff, 0xfb, 0x3f, 0x8f, 0x26, 0x66, 0x7c, 0x45, 0x4d, 0x51, 0xb9, 0x15, 0x21, 0x54, 0x18,
	0x17, 0xe4, 0x2a, 0xc7, 0x5f, 0xfa, 0xcf, 0x64, 0x9c, 0x40, 0x69, 0xd4, 0xbe, 0x29, 0xb0, 0xec,
	0xc0, 0x24, 0x4c, 0x6b, 0x82, 0xca, 0x38, 0xdb, 0x9d, 0x67, 0xbf, 0xee, 0xd9, 0xbb, 0xbf, 0xfd,
	0x3d, 0x81, 0x2d, 0x41, 0xb5, 0x8d, 0xc4, 0xe9, 0x86, 0xbf, 0x67, 0x77, 0xdf, 0xb1, 0x11, 0xc1,
	0x34, 0x88, 0x9e, 0x3f, 0xb2, 0x4e, 0x8a, 0x37, 0xd8, 0xc8, 0x36, 0xe6, 0x73, 0x16, 0x1a, 0x6b,
	0x2a, 0x03, 0x79, 0x02, 0x39, 0x58, 0x85, 0xe2, 0xc5, 0x34, 0x88, 0xee, 0xe5, 0x20, 0xe5, 0x9f,
	0x59, 0x58, 0x90, 0xfb, 0xd9, 0x2c, 0x95, 0x72, 0xb5, 0xad, 0xd6, 0xa9, 0xb8, 0xf3, 0x17, 0x8e,
	0xfb, 0x0b, 0xff, 0xe4, 0x72, 0xd0, 0xe3, 0x1f, 0xd8, 0x9b, 0x12, 0xad, 0x96, 0xa8, 0x1c, 0xe9,
	0x7d, 0x46, 0x58, 0x66, 0x2e, 0xd7, 0xe2, 0xc1, 0x8f, 0xb9, 0x85, 0xf8, 0x82, 0xbd, 0x25, 0x54,
	0x68, 0x4e, 0x38, 0x94, 0x1e, 0xbd, 0xf4, 0x1f, 0xda, 0x4e, 0xea, 0x09, 0xed, 0xcc, 0x41, 0xe2,
	0x8f, 0xda, 0x10, 0x6a, 0xf1, 0x34, 0x0d, 0xa2, 0x27, 0x79, 0x0b, 0xf1, 0x2f, 0x6c, 0x04, 0x75,
	0xe5, 0x24, 0x5a, 0x3c, 0x6f, 0x91, 0x8c, 0xd3, 0xe2, 0x95, 0x5f, 0x6b, 0xd4, 0xaf, 0x75, 0xf9,
	0xd7, 0x72, 0xd8, 0xe3, 0x11, 0x7b, 0x2c, 0x33, 0x20, 0xbd, 0x4e, 0x05, 0xf3, 0x4a, 0xd8, 0x2b,
	0xbb, 0x2e, 0x95, 0x17, 0xdc, 0x36, 0x09, 0x21, 0x3f, 0xae, 0x53, 0xf1, 0xfc, 0x4f, 0x53, 0x76,
	0xa9, 0xbc, 0x60, 0xbe, 0x60, 0x63, 0x8b, 0x67, 0x1f, 0x2f, 0xf5, 0xd1, 0xd8, 0x0d, 0x36, 0xe2,
	0xf5, 0xd5, 0xbb, 0x5d, 0x75, 0x92, 0x39, 0x9b, 0x29, 0x77, 0x8c, 0x33, 0xd4, 0x48, 0x90, 0x41,
	0x99, 0x1d, 0x08, 0x8a, 0x2c, 0x86, 0xc2, 0xf4, 0xda, 0x37, 0x38, 0xc1, 0x36, 0xf8, 0xfa, 0xe0,
	0x4f, 0x9f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xb6, 0xef, 0x04, 0x70, 0x02, 0x00, 0x00,
}