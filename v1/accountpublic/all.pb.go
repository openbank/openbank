// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/accountpublic/all.proto

package accountpublic

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	types "github.com/openbank/openbank/v1/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Owner struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Provider             string   `protobuf:"bytes,2,opt,name=Provider,json=provider,proto3" json:"Provider,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=DisplayName,json=display_name,proto3" json:"DisplayName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Owner) Reset()         { *m = Owner{} }
func (m *Owner) String() string { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()    {}
func (*Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{0}
}

func (m *Owner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Owner.Unmarshal(m, b)
}
func (m *Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Owner.Marshal(b, m, deterministic)
}
func (m *Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Owner.Merge(m, src)
}
func (m *Owner) XXX_Size() int {
	return xxx_messageInfo_Owner.Size(m)
}
func (m *Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_Owner proto.InternalMessageInfo

func (m *Owner) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Owner) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Owner) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type AccountRouting struct {
	Scheme               string   `protobuf:"bytes,1,opt,name=Scheme,json=scheme,proto3" json:"Scheme,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRouting) Reset()         { *m = AccountRouting{} }
func (m *AccountRouting) String() string { return proto.CompactTextString(m) }
func (*AccountRouting) ProtoMessage()    {}
func (*AccountRouting) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{1}
}

func (m *AccountRouting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRouting.Unmarshal(m, b)
}
func (m *AccountRouting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRouting.Marshal(b, m, deterministic)
}
func (m *AccountRouting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRouting.Merge(m, src)
}
func (m *AccountRouting) XXX_Size() int {
	return xxx_messageInfo_AccountRouting.Size(m)
}
func (m *AccountRouting) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRouting.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRouting proto.InternalMessageInfo

func (m *AccountRouting) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *AccountRouting) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type AccountRule struct {
	Scheme               string   `protobuf:"bytes,1,opt,name=Scheme,json=scheme,proto3" json:"Scheme,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRule) Reset()         { *m = AccountRule{} }
func (m *AccountRule) String() string { return proto.CompactTextString(m) }
func (*AccountRule) ProtoMessage()    {}
func (*AccountRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{2}
}

func (m *AccountRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRule.Unmarshal(m, b)
}
func (m *AccountRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRule.Marshal(b, m, deterministic)
}
func (m *AccountRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRule.Merge(m, src)
}
func (m *AccountRule) XXX_Size() int {
	return xxx_messageInfo_AccountRule.Size(m)
}
func (m *AccountRule) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRule.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRule proto.InternalMessageInfo

func (m *AccountRule) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *AccountRule) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetPublicAccountByIDRequest struct {
	BankID               int32    `protobuf:"varint,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	AccountID            int32    `protobuf:"varint,2,opt,name=AccountID,json=account_id,proto3" json:"AccountID,omitempty"`
	ViewID               int32    `protobuf:"varint,3,opt,name=ViewID,proto3" json:"ViewID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPublicAccountByIDRequest) Reset()         { *m = GetPublicAccountByIDRequest{} }
func (m *GetPublicAccountByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetPublicAccountByIDRequest) ProtoMessage()    {}
func (*GetPublicAccountByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{3}
}

func (m *GetPublicAccountByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicAccountByIDRequest.Unmarshal(m, b)
}
func (m *GetPublicAccountByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicAccountByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetPublicAccountByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicAccountByIDRequest.Merge(m, src)
}
func (m *GetPublicAccountByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetPublicAccountByIDRequest.Size(m)
}
func (m *GetPublicAccountByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicAccountByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicAccountByIDRequest proto.InternalMessageInfo

func (m *GetPublicAccountByIDRequest) GetBankID() int32 {
	if m != nil {
		return m.BankID
	}
	return 0
}

func (m *GetPublicAccountByIDRequest) GetAccountID() int32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *GetPublicAccountByIDRequest) GetViewID() int32 {
	if m != nil {
		return m.ViewID
	}
	return 0
}

type GetPublicAccountByIDResponse struct {
	ID                   string            `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	BankID               string            `protobuf:"bytes,2,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	Label                string            `protobuf:"bytes,3,opt,name=Label,json=label,proto3" json:"Label,omitempty"`
	Number               string            `protobuf:"bytes,4,opt,name=Number,json=number,proto3" json:"Number,omitempty"`
	Owners               *Owner            `protobuf:"bytes,5,opt,name=Owners,json=owners,proto3" json:"Owners,omitempty"`
	Type                 string            `protobuf:"bytes,6,opt,name=Type,json=type,proto3" json:"Type,omitempty"`
	Balance              *types.Amount     `protobuf:"bytes,7,opt,name=Balance,json=balance,proto3" json:"Balance,omitempty"`
	AccountRoutings      []*AccountRouting `protobuf:"bytes,8,rep,name=AccountRoutings,json=account_routings,proto3" json:"AccountRoutings,omitempty"`
	AccountRules         []*AccountRule    `protobuf:"bytes,9,rep,name=AccountRules,json=account_rules,proto3" json:"AccountRules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetPublicAccountByIDResponse) Reset()         { *m = GetPublicAccountByIDResponse{} }
func (m *GetPublicAccountByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetPublicAccountByIDResponse) ProtoMessage()    {}
func (*GetPublicAccountByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{4}
}

func (m *GetPublicAccountByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicAccountByIDResponse.Unmarshal(m, b)
}
func (m *GetPublicAccountByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicAccountByIDResponse.Marshal(b, m, deterministic)
}
func (m *GetPublicAccountByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicAccountByIDResponse.Merge(m, src)
}
func (m *GetPublicAccountByIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetPublicAccountByIDResponse.Size(m)
}
func (m *GetPublicAccountByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicAccountByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicAccountByIDResponse proto.InternalMessageInfo

func (m *GetPublicAccountByIDResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GetPublicAccountByIDResponse) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *GetPublicAccountByIDResponse) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *GetPublicAccountByIDResponse) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *GetPublicAccountByIDResponse) GetOwners() *Owner {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *GetPublicAccountByIDResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetPublicAccountByIDResponse) GetBalance() *types.Amount {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *GetPublicAccountByIDResponse) GetAccountRoutings() []*AccountRouting {
	if m != nil {
		return m.AccountRoutings
	}
	return nil
}

func (m *GetPublicAccountByIDResponse) GetAccountRules() []*AccountRule {
	if m != nil {
		return m.AccountRules
	}
	return nil
}

type ViewAvailable struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	ShortName            string   `protobuf:"bytes,2,opt,name=ShortName,json=short_name,proto3" json:"ShortName,omitempty"`
	IsPublic             bool     `protobuf:"varint,3,opt,name=IsPublic,json=is_public,proto3" json:"IsPublic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewAvailable) Reset()         { *m = ViewAvailable{} }
func (m *ViewAvailable) String() string { return proto.CompactTextString(m) }
func (*ViewAvailable) ProtoMessage()    {}
func (*ViewAvailable) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{5}
}

func (m *ViewAvailable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewAvailable.Unmarshal(m, b)
}
func (m *ViewAvailable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewAvailable.Marshal(b, m, deterministic)
}
func (m *ViewAvailable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewAvailable.Merge(m, src)
}
func (m *ViewAvailable) XXX_Size() int {
	return xxx_messageInfo_ViewAvailable.Size(m)
}
func (m *ViewAvailable) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewAvailable.DiscardUnknown(m)
}

var xxx_messageInfo_ViewAvailable proto.InternalMessageInfo

func (m *ViewAvailable) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ViewAvailable) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *ViewAvailable) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

type PublicAccount struct {
	ID                   string           `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Label                string           `protobuf:"bytes,2,opt,name=Label,json=label,proto3" json:"Label,omitempty"`
	BankID               string           `protobuf:"bytes,3,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	ViewsAvailable       []*ViewAvailable `protobuf:"bytes,4,rep,name=ViewsAvailable,json=views_available,proto3" json:"ViewsAvailable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PublicAccount) Reset()         { *m = PublicAccount{} }
func (m *PublicAccount) String() string { return proto.CompactTextString(m) }
func (*PublicAccount) ProtoMessage()    {}
func (*PublicAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{6}
}

func (m *PublicAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicAccount.Unmarshal(m, b)
}
func (m *PublicAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicAccount.Marshal(b, m, deterministic)
}
func (m *PublicAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicAccount.Merge(m, src)
}
func (m *PublicAccount) XXX_Size() int {
	return xxx_messageInfo_PublicAccount.Size(m)
}
func (m *PublicAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PublicAccount proto.InternalMessageInfo

func (m *PublicAccount) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PublicAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *PublicAccount) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *PublicAccount) GetViewsAvailable() []*ViewAvailable {
	if m != nil {
		return m.ViewsAvailable
	}
	return nil
}

type GetBankPublicAccountRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBankPublicAccountRequest) Reset()         { *m = GetBankPublicAccountRequest{} }
func (m *GetBankPublicAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetBankPublicAccountRequest) ProtoMessage()    {}
func (*GetBankPublicAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{7}
}

func (m *GetBankPublicAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBankPublicAccountRequest.Unmarshal(m, b)
}
func (m *GetBankPublicAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBankPublicAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetBankPublicAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBankPublicAccountRequest.Merge(m, src)
}
func (m *GetBankPublicAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetBankPublicAccountRequest.Size(m)
}
func (m *GetBankPublicAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBankPublicAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBankPublicAccountRequest proto.InternalMessageInfo

func (m *GetBankPublicAccountRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

type GetBankPublicAccountResponse struct {
	Accounts             []*PublicAccount `protobuf:"bytes,1,rep,name=Accounts,json=accounts,proto3" json:"Accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetBankPublicAccountResponse) Reset()         { *m = GetBankPublicAccountResponse{} }
func (m *GetBankPublicAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetBankPublicAccountResponse) ProtoMessage()    {}
func (*GetBankPublicAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{8}
}

func (m *GetBankPublicAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBankPublicAccountResponse.Unmarshal(m, b)
}
func (m *GetBankPublicAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBankPublicAccountResponse.Marshal(b, m, deterministic)
}
func (m *GetBankPublicAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBankPublicAccountResponse.Merge(m, src)
}
func (m *GetBankPublicAccountResponse) XXX_Size() int {
	return xxx_messageInfo_GetBankPublicAccountResponse.Size(m)
}
func (m *GetBankPublicAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBankPublicAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBankPublicAccountResponse proto.InternalMessageInfo

func (m *GetBankPublicAccountResponse) GetAccounts() []*PublicAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type GetPublicAccountAtAllBanksResponse struct {
	Accounts             []*PublicAccount `protobuf:"bytes,1,rep,name=Accounts,json=accounts,proto3" json:"Accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetPublicAccountAtAllBanksResponse) Reset()         { *m = GetPublicAccountAtAllBanksResponse{} }
func (m *GetPublicAccountAtAllBanksResponse) String() string { return proto.CompactTextString(m) }
func (*GetPublicAccountAtAllBanksResponse) ProtoMessage()    {}
func (*GetPublicAccountAtAllBanksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5387bf3777cfc660, []int{9}
}

func (m *GetPublicAccountAtAllBanksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicAccountAtAllBanksResponse.Unmarshal(m, b)
}
func (m *GetPublicAccountAtAllBanksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicAccountAtAllBanksResponse.Marshal(b, m, deterministic)
}
func (m *GetPublicAccountAtAllBanksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicAccountAtAllBanksResponse.Merge(m, src)
}
func (m *GetPublicAccountAtAllBanksResponse) XXX_Size() int {
	return xxx_messageInfo_GetPublicAccountAtAllBanksResponse.Size(m)
}
func (m *GetPublicAccountAtAllBanksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicAccountAtAllBanksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicAccountAtAllBanksResponse proto.InternalMessageInfo

func (m *GetPublicAccountAtAllBanksResponse) GetAccounts() []*PublicAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Owner)(nil), "accountpublic.Owner")
	proto.RegisterType((*AccountRouting)(nil), "accountpublic.AccountRouting")
	proto.RegisterType((*AccountRule)(nil), "accountpublic.AccountRule")
	proto.RegisterType((*GetPublicAccountByIDRequest)(nil), "accountpublic.GetPublicAccountByIDRequest")
	proto.RegisterType((*GetPublicAccountByIDResponse)(nil), "accountpublic.GetPublicAccountByIDResponse")
	proto.RegisterType((*ViewAvailable)(nil), "accountpublic.ViewAvailable")
	proto.RegisterType((*PublicAccount)(nil), "accountpublic.PublicAccount")
	proto.RegisterType((*GetBankPublicAccountRequest)(nil), "accountpublic.GetBankPublicAccountRequest")
	proto.RegisterType((*GetBankPublicAccountResponse)(nil), "accountpublic.GetBankPublicAccountResponse")
	proto.RegisterType((*GetPublicAccountAtAllBanksResponse)(nil), "accountpublic.GetPublicAccountAtAllBanksResponse")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/accountpublic/all.proto", fileDescriptor_5387bf3777cfc660)
}

var fileDescriptor_5387bf3777cfc660 = []byte{
	// 1492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xcf, 0x6f, 0x13, 0xdb,
	0x15, 0x9e, 0xb1, 0x63, 0xc7, 0xb9, 0x21, 0x80, 0x6e, 0x69, 0x6b, 0x0c, 0x45, 0x57, 0x06, 0x41,
	0xa0, 0x65, 0xc6, 0x31, 0x54, 0xb4, 0xa9, 0x10, 0x72, 0x30, 0xa5, 0xa6, 0x14, 0x52, 0x83, 0x50,
	0x45, 0x55, 0xa5, 0xd7, 0x33, 0xc7, 0xf6, 0x90, 0xf1, 0xdc, 0xe9, 0xbd, 0x77, 0x1c, 0x5c, 0x84,
	0x54, 0xb1, 0x69, 0xda, 0x15, 0x4d, 0xf7, 0xdd, 0x74, 0x5d, 0xa9, 0xcb, 0x4a, 0x95, 0x2a, 0x75,
	0x8d, 0x2a, 0x55, 0xea, 0xa6, 0xbb, 0xb7, 0x79, 0xff, 0xc3, 0xfb, 0xb1, 0x7a, 0xba, 0x77, 0xc6,
	0x93, 0x4c, 0x3c, 0x4e, 0x78, 0x42, 0x4f, 0x6f, 0x15, 0xc7, 0xe7, 0xbb, 0xdf, 0xf9, 0xce, 0xf9,
	0xce, 0x99, 0xb9, 0x46, 0xb7, 0x06, 0x9e, 0x1c, 0x46, 0x3d, 0xcb, 0x61, 0x23, 0x9b, 0x85, 0x10,
	0xf4, 0x68, 0xb0, 0xbd, 0xff, 0x61, 0xbc, 0x66, 0x53, 0xc7, 0x61, 0x51, 0x20, 0xc3, 0xa8, 0xe7,
	0x7b, 0x8e, 0x4d, 0x7d, 0xdf, 0x0a, 0x39, 0x93, 0x0c, 0xaf, 0x64, 0x02, 0xb5, 0xf3, 0x03, 0xc6,
	0x06, 0x3e, 0xd8, 0x34, 0xf4, 0x6c, 0x1a, 0x04, 0x4c, 0x52, 0xe9, 0xb1, 0x40, 0xc4, 0xe0, 0xda,
	0xf7, 0xf4, 0x1f, 0xe7, 0xfa, 0x00, 0x82, 0xeb, 0x62, 0x87, 0x0e, 0x06, 0xc0, 0x6d, 0x16, 0x6a,
	0x44, 0x0e, 0xfa, 0x5c, 0xc2, 0xa5, 0xff, 0xeb, 0x45, 0x7d, 0x1b, 0x46, 0xa1, 0x9c, 0x24, 0x41,
	0xfb, 0x38, 0xc1, 0x72, 0x12, 0x82, 0xd8, 0x17, 0x5a, 0x7f, 0x63, 0xa2, 0xd2, 0xe3, 0x9d, 0x00,
	0x38, 0xae, 0xa1, 0x42, 0xa7, 0x5d, 0x35, 0x89, 0xb9, 0xba, 0xb4, 0x81, 0x2a, 0x46, 0xd5, 0x58,
	0x35, 0x1a, 0xc6, 0xa6, 0xd1, 0x2d, 0x78, 0x2e, 0xbe, 0x8c, 0x2a, 0x9b, 0x9c, 0x8d, 0x3d, 0x17,
	0x78, 0xb5, 0x30, 0x83, 0xa8, 0x84, 0x49, 0x0c, 0x5f, 0x47, 0xcb, 0x6d, 0x4f, 0x84, 0x3e, 0x9d,
	0x3c, 0xa2, 0x23, 0xa8, 0x16, 0x67, 0xa0, 0x27, 0xdc, 0x38, 0xbc, 0x15, 0xd0, 0x11, 0xac, 0x97,
	0x2b, 0xc6, 0x69, 0xa3, 0x6a, 0xd4, 0x7b, 0xe8, 0x64, 0x2b, 0xee, 0x57, 0x97, 0x45, 0xd2, 0x0b,
	0x06, 0xb8, 0x8e, 0xca, 0x4f, 0x9c, 0x21, 0x8c, 0x20, 0x47, 0x50, 0x59, 0xe8, 0x08, 0xbe, 0x84,
	0x16, 0x5b, 0xae, 0xcb, 0x41, 0x88, 0x1c, 0x4d, 0x8b, 0x34, 0x0e, 0xa5, 0x39, 0x7e, 0x89, 0x96,
	0xa7, 0x39, 0x22, 0x1f, 0xde, 0x2b, 0x01, 0x41, 0xa5, 0x67, 0xd4, 0x8f, 0x20, 0x87, 0xbe, 0x34,
	0x56, 0x81, 0x94, 0xfc, 0x4f, 0x26, 0x3a, 0x77, 0x1f, 0xe4, 0xa6, 0x76, 0x3b, 0x49, 0xb3, 0x31,
	0xe9, 0xb4, 0xbb, 0xf0, 0x9b, 0x08, 0x84, 0xc4, 0x17, 0x51, 0x79, 0x83, 0x06, 0xdb, 0x49, 0x7f,
	0x4b, 0x59, 0xa5, 0xca, 0x9a, 0x2d, 0xcf, 0xc5, 0x57, 0xd1, 0x52, 0x72, 0xb4, 0xd3, 0xd6, 0x29,
	0xb3, 0x38, 0x94, 0x8c, 0x94, 0x82, 0xd6, 0x50, 0xf9, 0x99, 0x07, 0x3b, 0x9d, 0xb6, 0x6e, 0x71,
	0x06, 0x97, 0x6a, 0xfa, 0xa8, 0x88, 0xce, 0xe7, 0x6b, 0x12, 0x21, 0x0b, 0x04, 0x1c, 0x69, 0xf8,
	0xbe, 0xe0, 0x9c, 0xd6, 0x4e, 0x05, 0x13, 0x54, 0x7a, 0x48, 0x7b, 0xe0, 0xe7, 0xf8, 0x5c, 0xf2,
	0x55, 0x40, 0x75, 0xf9, 0x51, 0x34, 0xea, 0x01, 0xaf, 0x2e, 0xcc, 0x76, 0x39, 0xd0, 0x11, 0xfc,
	0x03, 0x54, 0xd6, 0x03, 0x28, 0xaa, 0x25, 0x62, 0xae, 0x2e, 0x37, 0xcf, 0x58, 0x99, 0xdd, 0xb1,
	0x74, 0x30, 0x7b, 0x92, 0x69, 0x3c, 0xbe, 0x80, 0x16, 0x9e, 0x4e, 0x42, 0xa8, 0x96, 0x67, 0xb8,
	0x17, 0xd4, 0x98, 0xe3, 0x1b, 0x68, 0x71, 0x83, 0xfa, 0x34, 0x70, 0xa0, 0xba, 0xa8, 0xa9, 0x57,
	0x2c, 0x3d, 0xfe, 0x56, 0x6b, 0xa4, 0x9b, 0x71, 0xa8, 0x28, 0x8d, 0xc4, 0x4f, 0xd1, 0xa9, 0xec,
	0x2c, 0x8a, 0x6a, 0x85, 0x14, 0x57, 0x97, 0x9b, 0xdf, 0x39, 0xa4, 0x2b, 0x8b, 0xca, 0x90, 0x9d,
	0x9e, 0x5a, 0xc5, 0x13, 0x0a, 0xfc, 0x53, 0x74, 0xe2, 0xc0, 0xf4, 0x89, 0xea, 0x92, 0xa6, 0xac,
	0xcd, 0xa1, 0x8c, 0x7c, 0xc8, 0xf0, 0xad, 0xa4, 0x7c, 0xea, 0x70, 0xea, 0xf0, 0xef, 0x4d, 0xb4,
	0xa2, 0xc6, 0xa0, 0x35, 0xa6, 0x9e, 0x4f, 0x7b, 0xfe, 0xd1, 0x96, 0x5e, 0x45, 0x4b, 0x4f, 0x86,
	0x8c, 0x4b, 0xbd, 0x99, 0xb3, 0xae, 0x22, 0xa1, 0x82, 0x7a, 0x2f, 0xf1, 0x15, 0x54, 0xe9, 0x88,
	0x78, 0x70, 0xb4, 0xb7, 0x95, 0x0c, 0x72, 0xc9, 0x13, 0x5b, 0xb1, 0xe0, 0x54, 0xc9, 0x7f, 0x4c,
	0xb4, 0x92, 0x19, 0xb4, 0x23, 0x95, 0xa4, 0x73, 0x53, 0x98, 0x37, 0x37, 0xfb, 0xe3, 0x57, 0x9c,
	0x3f, 0x7e, 0x3f, 0x47, 0x27, 0x55, 0xf5, 0x22, 0x2d, 0xbf, 0xba, 0xa0, 0xbb, 0x7a, 0xfe, 0x50,
	0x57, 0x33, 0x2d, 0xca, 0x50, 0x9d, 0x1a, 0xab, 0xf3, 0x5b, 0x74, 0x1a, 0x4c, 0xeb, 0x79, 0xa0,
	0xd7, 0x59, 0x49, 0xc8, 0x54, 0x95, 0xbf, 0xce, 0xf9, 0xf2, 0x52, 0x2e, 0x5f, 0xaf, 0x61, 0x0e,
	0x57, 0xb2, 0x86, 0x6d, 0x54, 0x49, 0xbe, 0x12, 0x55, 0x33, 0xb7, 0x80, 0xec, 0x0a, 0x67, 0x9e,
	0xbc, 0x09, 0x70, 0x7f, 0x26, 0x38, 0xaa, 0x1f, 0x5e, 0xfa, 0x96, 0x6c, 0xf9, 0xbe, 0xca, 0x2f,
	0xbe, 0x9a, 0x9c, 0xcd, 0x3d, 0x84, 0xce, 0x24, 0xc8, 0xf8, 0xd8, 0x13, 0xe0, 0x63, 0xcf, 0x01,
	0xfc, 0xef, 0x22, 0x3a, 0x93, 0xf7, 0x08, 0xc2, 0xd7, 0x0e, 0x65, 0x3b, 0xe2, 0xd9, 0x59, 0xfb,
	0xee, 0x7b, 0x61, 0xe3, 0xc2, 0xea, 0xef, 0x0a, 0x7b, 0xad, 0xcf, 0xcd, 0xf4, 0x75, 0x42, 0x62,
	0x24, 0x3e, 0x7b, 0x1f, 0xa6, 0x9f, 0xc9, 0x34, 0xb4, 0x31, 0x21, 0x9d, 0xb6, 0x55, 0x5b, 0xeb,
	0x82, 0x8c, 0x78, 0x40, 0x9c, 0x48, 0x48, 0x36, 0x02, 0x4e, 0xe2, 0x34, 0x24, 0x49, 0x4a, 0xfa,
	0x8c, 0x13, 0x11, 0x82, 0xe3, 0xf5, 0x3d, 0x87, 0x28, 0x6f, 0xad, 0x07, 0x2e, 0x2a, 0x36, 0x1b,
	0x0d, 0xfc, 0x2b, 0x74, 0x21, 0xd1, 0x48, 0xe0, 0x25, 0x38, 0x91, 0x04, 0x97, 0x88, 0xc8, 0x71,
	0x40, 0x88, 0x7e, 0xe4, 0xfb, 0x13, 0x0b, 0xff, 0x08, 0xfd, 0xb0, 0x76, 0xeb, 0xa2, 0xed, 0x42,
	0xdf, 0x0b, 0xbc, 0xe4, 0x7d, 0x7e, 0xb0, 0x96, 0xa3, 0x4a, 0xe9, 0xdd, 0x44, 0x4d, 0x54, 0x7e,
	0xdc, 0x8a, 0xe4, 0xb0, 0x89, 0x57, 0xd1, 0xe5, 0xa1, 0x94, 0xa1, 0x58, 0xb7, 0x6d, 0x1a, 0xc9,
	0xa1, 0xd5, 0x0b, 0xb6, 0x2d, 0xc9, 0xb2, 0x64, 0x16, 0x07, 0xea, 0xbe, 0xf9, 0xdf, 0xc7, 0x7f,
	0x2e, 0xb4, 0xf0, 0x1d, 0xf5, 0xaa, 0x57, 0x62, 0x85, 0xfd, 0x2a, 0x9e, 0xd4, 0xd7, 0xf6, 0xf4,
	0x8e, 0x92, 0xb8, 0x67, 0xbf, 0x4a, 0x5f, 0x36, 0xaf, 0xed, 0x57, 0xf1, 0xdb, 0xe4, 0xf5, 0x34,
	0xba, 0x5b, 0x30, 0xde, 0x16, 0xb4, 0xe1, 0xf8, 0xaf, 0xb1, 0x87, 0x33, 0xf3, 0x9b, 0xe7, 0xe1,
	0xbc, 0x85, 0xc9, 0xf3, 0x70, 0xee, 0x42, 0xd4, 0xff, 0x56, 0xd8, 0x6b, 0x7d, 0x36, 0xeb, 0x61,
	0x55, 0x79, 0xa8, 0x6a, 0x3a, 0xe4, 0xd0, 0xd7, 0x6e, 0xe1, 0xdc, 0x4a, 0x3e, 0xc0, 0xc2, 0x4b,
	0xb8, 0x9e, 0x63, 0x61, 0xea, 0x5d, 0x0c, 0x3f, 0xe0, 0xd2, 0x1f, 0x8b, 0xa8, 0x36, 0x7f, 0xef,
	0xf1, 0xb7, 0xac, 0xf8, 0xd2, 0x68, 0x4d, 0x2f, 0x8d, 0xd6, 0x3d, 0x75, 0x69, 0xac, 0xad, 0x1d,
	0xb3, 0x5b, 0xb3, 0x8f, 0x8e, 0xfa, 0x5f, 0xf2, 0xdd, 0x21, 0xca, 0x9d, 0x43, 0x7d, 0xa7, 0x92,
	0x50, 0xdf, 0xd7, 0x1d, 0x17, 0x56, 0xed, 0x62, 0xe2, 0xd2, 0x51, 0xa0, 0x07, 0x2f, 0x62, 0x5f,
	0x9c, 0x63, 0x7d, 0x69, 0xa1, 0x3b, 0xb5, 0xdb, 0x5f, 0x62, 0xb5, 0x66, 0x2b, 0xf9, 0x00, 0x77,
	0xbe, 0x89, 0xbf, 0x71, 0xe0, 0xf2, 0x9f, 0xb5, 0xa3, 0x61, 0xd4, 0x8a, 0xbb, 0x05, 0x63, 0xe3,
	0x93, 0xf2, 0x5e, 0xeb, 0x5d, 0x19, 0x15, 0x9b, 0x56, 0x03, 0xff, 0x1a, 0xe1, 0x6c, 0xbb, 0x48,
	0x6b, 0xb3, 0x83, 0xef, 0x26, 0x57, 0x6a, 0x41, 0x1c, 0x0e, 0x54, 0x02, 0xa1, 0x81, 0x4b, 0x54,
	0x0e, 0xc2, 0x42, 0xe0, 0xf1, 0x9d, 0x9f, 0xb0, 0x80, 0xc8, 0x21, 0xa4, 0xdd, 0x4a, 0x9a, 0xc7,
	0x41, 0xb0, 0x88, 0x3b, 0x60, 0x35, 0x4b, 0x6b, 0x56, 0xc3, 0x6a, 0x5c, 0x33, 0x0b, 0xcd, 0xd3,
	0x34, 0x0c, 0x7d, 0xcf, 0xd1, 0xe7, 0xec, 0x17, 0x82, 0x05, 0xeb, 0x33, 0xdf, 0x74, 0xd7, 0x50,
	0xf1, 0x66, 0xe3, 0x26, 0xbe, 0x86, 0x56, 0x63, 0x33, 0xc0, 0x25, 0x3b, 0x43, 0x88, 0x53, 0x4c,
	0x39, 0x89, 0x27, 0x48, 0xc0, 0xd4, 0xd6, 0x44, 0x81, 0x6b, 0x75, 0x7f, 0x8c, 0x8a, 0xdf, 0x6f,
	0x34, 0xf0, 0x1d, 0x74, 0x3b, 0x7b, 0x84, 0x06, 0x24, 0x0a, 0xe0, 0x65, 0x08, 0x8e, 0x32, 0x07,
	0x38, 0x67, 0x9c, 0x30, 0xc7, 0x89, 0x38, 0xb8, 0x53, 0xd1, 0x02, 0xf8, 0x18, 0x38, 0x11, 0x9e,
	0x0b, 0x56, 0x77, 0x4b, 0xa5, 0x6e, 0xe0, 0x5f, 0xa0, 0x67, 0x79, 0xa9, 0x63, 0xaf, 0x7b, 0xcc,
	0x9d, 0xa8, 0xf4, 0x23, 0xea, 0xf7, 0x19, 0x1f, 0x51, 0xa9, 0xa8, 0x19, 0x27, 0x2e, 0x83, 0x58,
	0xd3, 0x88, 0x4a, 0x67, 0xa8, 0x8f, 0xa4, 0x99, 0x93, 0xb3, 0x56, 0xf7, 0xa1, 0x4a, 0xb0, 0x86,
	0xef, 0xa1, 0xbb, 0xf3, 0x13, 0xa4, 0x44, 0x0e, 0x0b, 0x24, 0xf5, 0x02, 0xa1, 0xa3, 0x91, 0x00,
	0x7e, 0x45, 0xbb, 0xe0, 0x42, 0x20, 0x3d, 0xea, 0x0b, 0xab, 0xbb, 0xa9, 0xd8, 0x6e, 0xe0, 0x0e,
	0xba, 0x3f, 0xcb, 0xa6, 0xf0, 0xfb, 0x54, 0x43, 0x3a, 0x06, 0x12, 0x02, 0x1f, 0x79, 0x42, 0x78,
	0xaa, 0x72, 0xa6, 0xdc, 0x02, 0x21, 0x32, 0x5d, 0xb5, 0x9e, 0xff, 0xab, 0x80, 0xfe, 0x59, 0x48,
	0x07, 0xee, 0xef, 0x85, 0x4a, 0x11, 0xef, 0x9a, 0xad, 0x04, 0xc9, 0xe6, 0x39, 0x2c, 0x54, 0x3f,
	0x38, 0x08, 0xc9, 0x3d, 0x5d, 0xae, 0x42, 0x46, 0x72, 0xa8, 0x74, 0x3a, 0x54, 0x7d, 0xa1, 0xc4,
	0x08, 0x8b, 0x3c, 0x55, 0x43, 0xb2, 0x1f, 0x50, 0x42, 0x42, 0xce, 0x34, 0x79, 0x9f, 0xf9, 0x3e,
	0xdb, 0x89, 0xe5, 0xa8, 0xe4, 0x8c, 0x7b, 0xbf, 0x8d, 0x11, 0x77, 0x99, 0x0b, 0xa4, 0xef, 0xb3,
	0x1d, 0x6b, 0x75, 0xa1, 0x59, 0xd1, 0x23, 0x1d, 0xc9, 0xe1, 0xfa, 0x92, 0xfe, 0xa1, 0xc8, 0xb6,
	0x21, 0xd8, 0xf8, 0x83, 0x89, 0x5a, 0xef, 0xbb, 0x1a, 0xf8, 0xdb, 0xea, 0x2d, 0x42, 0x32, 0x01,
	0xe2, 0x52, 0x49, 0x51, 0x1b, 0x5d, 0x39, 0x9e, 0x62, 0x87, 0x7b, 0x12, 0xf0, 0xd9, 0x9f, 0xd1,
	0x80, 0x0e, 0x20, 0x87, 0xe5, 0x27, 0xe6, 0xa6, 0xf1, 0x3c, 0xfb, 0xbb, 0xfa, 0x77, 0xa6, 0xb1,
	0x6b, 0x1a, 0x6f, 0x4d, 0xe3, 0x1f, 0xa6, 0xf1, 0x7f, 0xd3, 0xf8, 0xd4, 0x34, 0xfe, 0x5b, 0x30,
	0x7a, 0x65, 0xfd, 0x90, 0xbb, 0xf1, 0x45, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xfa, 0x99, 0x90,
	0xbd, 0x0f, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountPublicServiceClient is the client API for AccountPublicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountPublicServiceClient interface {
	// GetPublicAccountByID
	GetPublicAccountByID(ctx context.Context, in *GetPublicAccountByIDRequest, opts ...grpc.CallOption) (*GetPublicAccountByIDResponse, error)
	// GetBankPublicAccount
	GetBankPublicAccount(ctx context.Context, in *GetBankPublicAccountRequest, opts ...grpc.CallOption) (*GetBankPublicAccountResponse, error)
	// GetPublicAccountAtAllBanks
	GetPublicAccountAtAllBanks(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetPublicAccountAtAllBanksResponse, error)
}

type accountPublicServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountPublicServiceClient(cc *grpc.ClientConn) AccountPublicServiceClient {
	return &accountPublicServiceClient{cc}
}

func (c *accountPublicServiceClient) GetPublicAccountByID(ctx context.Context, in *GetPublicAccountByIDRequest, opts ...grpc.CallOption) (*GetPublicAccountByIDResponse, error) {
	out := new(GetPublicAccountByIDResponse)
	err := c.cc.Invoke(ctx, "/accountpublic.AccountPublicService/GetPublicAccountByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountPublicServiceClient) GetBankPublicAccount(ctx context.Context, in *GetBankPublicAccountRequest, opts ...grpc.CallOption) (*GetBankPublicAccountResponse, error) {
	out := new(GetBankPublicAccountResponse)
	err := c.cc.Invoke(ctx, "/accountpublic.AccountPublicService/GetBankPublicAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountPublicServiceClient) GetPublicAccountAtAllBanks(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetPublicAccountAtAllBanksResponse, error) {
	out := new(GetPublicAccountAtAllBanksResponse)
	err := c.cc.Invoke(ctx, "/accountpublic.AccountPublicService/GetPublicAccountAtAllBanks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountPublicServiceServer is the server API for AccountPublicService service.
type AccountPublicServiceServer interface {
	// GetPublicAccountByID
	GetPublicAccountByID(context.Context, *GetPublicAccountByIDRequest) (*GetPublicAccountByIDResponse, error)
	// GetBankPublicAccount
	GetBankPublicAccount(context.Context, *GetBankPublicAccountRequest) (*GetBankPublicAccountResponse, error)
	// GetPublicAccountAtAllBanks
	GetPublicAccountAtAllBanks(context.Context, *empty.Empty) (*GetPublicAccountAtAllBanksResponse, error)
}

// UnimplementedAccountPublicServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountPublicServiceServer struct {
}

func (*UnimplementedAccountPublicServiceServer) GetPublicAccountByID(ctx context.Context, req *GetPublicAccountByIDRequest) (*GetPublicAccountByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicAccountByID not implemented")
}
func (*UnimplementedAccountPublicServiceServer) GetBankPublicAccount(ctx context.Context, req *GetBankPublicAccountRequest) (*GetBankPublicAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankPublicAccount not implemented")
}
func (*UnimplementedAccountPublicServiceServer) GetPublicAccountAtAllBanks(ctx context.Context, req *empty.Empty) (*GetPublicAccountAtAllBanksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicAccountAtAllBanks not implemented")
}

func RegisterAccountPublicServiceServer(s *grpc.Server, srv AccountPublicServiceServer) {
	s.RegisterService(&_AccountPublicService_serviceDesc, srv)
}

func _AccountPublicService_GetPublicAccountByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicAccountByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountPublicServiceServer).GetPublicAccountByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountpublic.AccountPublicService/GetPublicAccountByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountPublicServiceServer).GetPublicAccountByID(ctx, req.(*GetPublicAccountByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountPublicService_GetBankPublicAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankPublicAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountPublicServiceServer).GetBankPublicAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountpublic.AccountPublicService/GetBankPublicAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountPublicServiceServer).GetBankPublicAccount(ctx, req.(*GetBankPublicAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountPublicService_GetPublicAccountAtAllBanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountPublicServiceServer).GetPublicAccountAtAllBanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountpublic.AccountPublicService/GetPublicAccountAtAllBanks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountPublicServiceServer).GetPublicAccountAtAllBanks(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountPublicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "accountpublic.AccountPublicService",
	HandlerType: (*AccountPublicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicAccountByID",
			Handler:    _AccountPublicService_GetPublicAccountByID_Handler,
		},
		{
			MethodName: "GetBankPublicAccount",
			Handler:    _AccountPublicService_GetBankPublicAccount_Handler,
		},
		{
			MethodName: "GetPublicAccountAtAllBanks",
			Handler:    _AccountPublicService_GetPublicAccountAtAllBanks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/accountpublic/all.proto",
}
