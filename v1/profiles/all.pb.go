// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/profiles/all.proto

package profiles

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	account "github.com/openbank/openbank/v1/account"
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

// GetProfileRequest is the request message for retrieving customer profiles.
type GetProfileRequest struct {
	ProfileID            string   `protobuf:"bytes,1,opt,name=ProfileID,json=profile_id,proto3" json:"ProfileID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileRequest) Reset()         { *m = GetProfileRequest{} }
func (m *GetProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetProfileRequest) ProtoMessage()    {}
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01f74d04d9bd7b36, []int{0}
}

func (m *GetProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileRequest.Unmarshal(m, b)
}
func (m *GetProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileRequest.Merge(m, src)
}
func (m *GetProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetProfileRequest.Size(m)
}
func (m *GetProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileRequest proto.InternalMessageInfo

func (m *GetProfileRequest) GetProfileID() string {
	if m != nil {
		return m.ProfileID
	}
	return ""
}

// GetProfileResponse is the response message for retrieving customer profiles.
type GetProfileResponse struct {
	Profile              *types.Profile     `protobuf:"bytes,1,opt,name=Profile,json=profile,proto3" json:"Profile,omitempty"`
	Accounts             []*account.Account `protobuf:"bytes,2,rep,name=Accounts,json=accounts,proto3" json:"Accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetProfileResponse) Reset()         { *m = GetProfileResponse{} }
func (m *GetProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetProfileResponse) ProtoMessage()    {}
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01f74d04d9bd7b36, []int{1}
}

func (m *GetProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileResponse.Unmarshal(m, b)
}
func (m *GetProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileResponse.Merge(m, src)
}
func (m *GetProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetProfileResponse.Size(m)
}
func (m *GetProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileResponse proto.InternalMessageInfo

func (m *GetProfileResponse) GetProfile() *types.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *GetProfileResponse) GetAccounts() []*account.Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// GetProfileCardsRequest is the request message for retrieving customer cards.
type GetProfileCardsRequest struct {
	ProfileID            string   `protobuf:"bytes,1,opt,name=ProfileID,json=profile_id,proto3" json:"ProfileID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileCardsRequest) Reset()         { *m = GetProfileCardsRequest{} }
func (m *GetProfileCardsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProfileCardsRequest) ProtoMessage()    {}
func (*GetProfileCardsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01f74d04d9bd7b36, []int{2}
}

func (m *GetProfileCardsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileCardsRequest.Unmarshal(m, b)
}
func (m *GetProfileCardsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileCardsRequest.Marshal(b, m, deterministic)
}
func (m *GetProfileCardsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileCardsRequest.Merge(m, src)
}
func (m *GetProfileCardsRequest) XXX_Size() int {
	return xxx_messageInfo_GetProfileCardsRequest.Size(m)
}
func (m *GetProfileCardsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileCardsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileCardsRequest proto.InternalMessageInfo

func (m *GetProfileCardsRequest) GetProfileID() string {
	if m != nil {
		return m.ProfileID
	}
	return ""
}

// GetProfileCardsResponse is the response message for retrieving customer cards.
type GetProfileCardsResponse struct {
	Cards                []*ProfileCard `protobuf:"bytes,1,rep,name=Cards,json=cards,proto3" json:"Cards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetProfileCardsResponse) Reset()         { *m = GetProfileCardsResponse{} }
func (m *GetProfileCardsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProfileCardsResponse) ProtoMessage()    {}
func (*GetProfileCardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01f74d04d9bd7b36, []int{3}
}

func (m *GetProfileCardsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileCardsResponse.Unmarshal(m, b)
}
func (m *GetProfileCardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileCardsResponse.Marshal(b, m, deterministic)
}
func (m *GetProfileCardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileCardsResponse.Merge(m, src)
}
func (m *GetProfileCardsResponse) XXX_Size() int {
	return xxx_messageInfo_GetProfileCardsResponse.Size(m)
}
func (m *GetProfileCardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileCardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileCardsResponse proto.InternalMessageInfo

func (m *GetProfileCardsResponse) GetCards() []*ProfileCard {
	if m != nil {
		return m.Cards
	}
	return nil
}

// ProfileCard holds details about a customer's credit card.
type ProfileCard struct {
	// CardToken is the card number.
	CardToken string `protobuf:"bytes,1,opt,name=CardToken,json=card_token,proto3" json:"CardToken,omitempty"`
	// Category is the card type.
	Category string `protobuf:"bytes,2,opt,name=Category,json=category,proto3" json:"Category,omitempty"`
	// LastFour is the last four digits of the card.
	LastFour string `protobuf:"bytes,3,opt,name=LastFour,json=last_four,proto3" json:"LastFour,omitempty"`
	// OwnerName is the name of the card's owner.
	OwnerName string `protobuf:"bytes,4,opt,name=OwnerName,json=owner_name,proto3" json:"OwnerName,omitempty"`
	// Type is the type of the account associated with the card.
	Type                 string   `protobuf:"bytes,5,opt,name=Type,json=type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileCard) Reset()         { *m = ProfileCard{} }
func (m *ProfileCard) String() string { return proto.CompactTextString(m) }
func (*ProfileCard) ProtoMessage()    {}
func (*ProfileCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_01f74d04d9bd7b36, []int{4}
}

func (m *ProfileCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileCard.Unmarshal(m, b)
}
func (m *ProfileCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileCard.Marshal(b, m, deterministic)
}
func (m *ProfileCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileCard.Merge(m, src)
}
func (m *ProfileCard) XXX_Size() int {
	return xxx_messageInfo_ProfileCard.Size(m)
}
func (m *ProfileCard) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileCard.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileCard proto.InternalMessageInfo

func (m *ProfileCard) GetCardToken() string {
	if m != nil {
		return m.CardToken
	}
	return ""
}

func (m *ProfileCard) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ProfileCard) GetLastFour() string {
	if m != nil {
		return m.LastFour
	}
	return ""
}

func (m *ProfileCard) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *ProfileCard) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProfileRequest)(nil), "profiles.GetProfileRequest")
	proto.RegisterType((*GetProfileResponse)(nil), "profiles.GetProfileResponse")
	proto.RegisterType((*GetProfileCardsRequest)(nil), "profiles.GetProfileCardsRequest")
	proto.RegisterType((*GetProfileCardsResponse)(nil), "profiles.GetProfileCardsResponse")
	proto.RegisterType((*ProfileCard)(nil), "profiles.ProfileCard")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/profiles/all.proto", fileDescriptor_01f74d04d9bd7b36)
}

var fileDescriptor_01f74d04d9bd7b36 = []byte{
	// 1099 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0xde, 0x5d, 0xc7, 0xae, 0x33, 0xf9, 0x29, 0x6a, 0xa7, 0xfa, 0x81, 0x65, 0xa2, 0x6a, 0x08,
	0x12, 0x35, 0xa1, 0x59, 0xff, 0x69, 0x10, 0x22, 0x12, 0x42, 0x4e, 0x4a, 0x42, 0x42, 0xd5, 0x1a,
	0x37, 0x45, 0x28, 0x37, 0xd1, 0x78, 0xf7, 0xac, 0x3d, 0x74, 0x3d, 0xb3, 0xcc, 0xcc, 0xc6, 0x35,
	0x57, 0x28, 0xdc, 0xe4, 0x02, 0xa4, 0x2a, 0x3c, 0x01, 0x3c, 0x05, 0x2f, 0xc0, 0x7d, 0x25, 0x6e,
	0x10, 0xea, 0x65, 0x91, 0x78, 0x04, 0x2e, 0xd1, 0xec, 0xae, 0xd7, 0xb6, 0x1c, 0x88, 0x10, 0x5c,
	0xad, 0x7d, 0xce, 0x99, 0xef, 0x7c, 0xe7, 0x3b, 0xe7, 0xec, 0x0e, 0x6a, 0xf5, 0x99, 0x1e, 0xc4,
	0x3d, 0xd7, 0x13, 0xc3, 0xba, 0x88, 0x80, 0xf7, 0x28, 0x7f, 0x32, 0xfd, 0x71, 0xda, 0xac, 0x47,
	0x52, 0x04, 0x2c, 0x04, 0x55, 0xa7, 0x61, 0xe8, 0x46, 0x52, 0x68, 0x81, 0xcb, 0x13, 0x5b, 0x75,
	0xad, 0x2f, 0x44, 0x3f, 0x84, 0x3a, 0x8d, 0x58, 0x9d, 0x72, 0x2e, 0x34, 0xd5, 0x4c, 0x70, 0x95,
	0xc6, 0x55, 0xef, 0x24, 0x0f, 0x6f, 0xb3, 0x0f, 0x7c, 0x53, 0x8d, 0x68, 0xbf, 0x0f, 0xb2, 0x2e,
	0xa2, 0x24, 0xe2, 0x92, 0xe8, 0xe6, 0x55, 0x4c, 0xa8, 0xe7, 0x89, 0x98, 0xeb, 0x29, 0x91, 0x6a,
	0xfd, 0xaa, 0x23, 0x7a, 0x1c, 0xcd, 0x32, 0x5f, 0xdf, 0x43, 0x37, 0xf6, 0x41, 0x77, 0x52, 0xfa,
	0x5d, 0xf8, 0x22, 0x06, 0xa5, 0xf1, 0x5b, 0x68, 0x39, 0xb3, 0x1c, 0xdc, 0xab, 0xd8, 0xc4, 0xae,
	0x2d, 0xef, 0xa0, 0xb2, 0x55, 0xb1, 0x6a, 0x56, 0xc3, 0xea, 0x58, 0x5d, 0x94, 0x55, 0x7b, 0xc2,
	0xfc, 0xed, 0x52, 0xd9, 0xba, 0x6e, 0x55, 0xac, 0xf5, 0x6f, 0x6d, 0x84, 0x67, 0x81, 0x54, 0x24,
	0xb8, 0x02, 0xbc, 0x85, 0xae, 0x65, 0xa6, 0x04, 0x67, 0xa5, 0xb5, 0xea, 0x26, 0x0c, 0xdc, 0xcc,
	0x3a, 0x87, 0x7b, 0x2d, 0xc3, 0xc5, 0xef, 0xa1, 0x72, 0x3b, 0x2d, 0x4d, 0x55, 0x1c, 0x52, 0xa8,
	0xad, 0xb4, 0x6e, 0xb8, 0x59, 0xad, 0xca, 0xcd, 0x3c, 0x73, 0x27, 0xcb, 0x13, 0x6f, 0xce, 0xe7,
	0x63, 0xf4, 0xca, 0x94, 0xce, 0x2e, 0x95, 0xbe, 0xfa, 0x17, 0xc5, 0x1d, 0xa3, 0x57, 0x17, 0xc0,
	0xb2, 0x02, 0xdf, 0x45, 0xc5, 0xc4, 0x50, 0xb1, 0x13, 0x9e, 0xff, 0x77, 0x27, 0x93, 0xe0, 0xce,
	0x84, 0xcf, 0x25, 0x28, 0x7a, 0x26, 0x3e, 0xc7, 0x7e, 0x61, 0xa3, 0x95, 0x99, 0x50, 0x43, 0xcf,
	0x3c, 0x8f, 0xc4, 0x13, 0xe0, 0x97, 0xd1, 0x33, 0xa7, 0x4f, 0xb4, 0xf1, 0xe2, 0x37, 0x51, 0x79,
	0x97, 0x6a, 0xe8, 0x0b, 0x39, 0xae, 0x38, 0x0b, 0x91, 0x65, 0x2f, 0xf3, 0xe1, 0xdb, 0xa8, 0x7c,
	0x9f, 0x2a, 0xbd, 0x27, 0x62, 0x59, 0x29, 0x2c, 0xc4, 0x2d, 0x87, 0x54, 0xe9, 0x93, 0x40, 0xc4,
	0xd2, 0xe4, 0x7e, 0x38, 0xe2, 0x20, 0x1f, 0xd0, 0x21, 0x54, 0x96, 0x16, 0x73, 0x0b, 0xe3, 0x3c,
	0xe1, 0x74, 0x08, 0xf8, 0x16, 0x5a, 0x3a, 0x1a, 0x47, 0x50, 0x29, 0x2e, 0x44, 0x2d, 0x99, 0x06,
	0x4f, 0xca, 0x6b, 0x7d, 0x53, 0x42, 0xab, 0x59, 0x79, 0x8f, 0x40, 0x9e, 0x32, 0x0f, 0xf0, 0x73,
	0x07, 0xa1, 0xa9, 0x9c, 0xf8, 0xb5, 0xa9, 0x64, 0x0b, 0x93, 0x58, 0x5d, 0xbb, 0xdc, 0x99, 0x8a,
	0xbf, 0x7e, 0xe6, 0x5c, 0xb4, 0x5f, 0xda, 0xf9, 0x88, 0xe1, 0xd5, 0x4f, 0x62, 0x90, 0x63, 0x92,
	0x2f, 0xe3, 0xdb, 0x5d, 0xd0, 0xb1, 0xe4, 0x8a, 0xe8, 0x01, 0x4c, 0xac, 0x84, 0x72, 0x9f, 0x50,
	0xa5, 0x84, 0xc7, 0xa8, 0x06, 0x9f, 0x4c, 0x26, 0xe7, 0xf0, 0x31, 0x2a, 0xb4, 0x1a, 0x0d, 0xfc,
	0x00, 0xdd, 0xca, 0x92, 0x13, 0x78, 0x0a, 0x5e, 0x6c, 0x62, 0x54, 0xec, 0x79, 0xa0, 0x54, 0x10,
	0x87, 0xe1, 0xd8, 0xc5, 0x77, 0xd0, 0x46, 0xb5, 0xf6, 0x46, 0xdd, 0x87, 0x80, 0x71, 0x96, 0xee,
	0xee, 0x24, 0xe3, 0x22, 0xc7, 0xc3, 0x26, 0x2a, 0x6c, 0x35, 0xb6, 0xf0, 0x06, 0xaa, 0xa5, 0x64,
	0xc0, 0x27, 0xa3, 0x01, 0xf0, 0x84, 0x92, 0x04, 0x25, 0x62, 0xe9, 0x01, 0x61, 0x8a, 0x70, 0xa1,
	0x49, 0x20, 0x62, 0xee, 0xbb, 0x3d, 0x8c, 0xae, 0xa3, 0xd2, 0xc3, 0x76, 0xac, 0x07, 0x2d, 0x5c,
	0x42, 0x4b, 0x12, 0xa8, 0x7f, 0xf6, 0xf3, 0xcb, 0xef, 0x9c, 0x55, 0xfc, 0xbf, 0xd9, 0xf7, 0xcf,
	0xb9, 0x63, 0x3d, 0x73, 0x12, 0xd5, 0xf1, 0xf7, 0x05, 0xb4, 0x3a, 0x3f, 0xa1, 0x98, 0x5c, 0xa6,
	0xdc, 0xec, 0x22, 0x54, 0x5f, 0xff, 0x9b, 0x88, 0x4c, 0xe0, 0x9f, 0x9c, 0x8b, 0xf6, 0x0f, 0xce,
	0x54, 0xe0, 0x95, 0x54, 0xe0, 0x64, 0x82, 0xab, 0xfd, 0x89, 0xba, 0x94, 0x13, 0x2a, 0x25, 0x1d,
	0x13, 0x11, 0x90, 0x19, 0x98, 0x59, 0x95, 0x47, 0x4c, 0x0f, 0x92, 0x92, 0x33, 0xb9, 0x49, 0x20,
	0xe4, 0x5c, 0x57, 0x7a, 0x54, 0x81, 0x4f, 0x04, 0xcf, 0x0d, 0xcc, 0x07, 0xae, 0x59, 0xc0, 0x40,
	0x1e, 0x1e, 0xa7, 0x9d, 0x79, 0x74, 0x65, 0x67, 0x9a, 0xa8, 0x5e, 0xdd, 0xbc, 0xaa, 0x33, 0x73,
	0x15, 0xfe, 0x97, 0xed, 0xb9, 0x89, 0x6f, 0xcc, 0x7d, 0x1e, 0x8c, 0x4a, 0xd3, 0x1e, 0x55, 0x0b,
	0xe7, 0x8e, 0xb5, 0xf3, 0x75, 0xe9, 0xa2, 0xfd, 0xa2, 0x88, 0x0a, 0x2d, 0xb7, 0x81, 0xf7, 0xf3,
	0xc5, 0x27, 0xed, 0xce, 0x01, 0x6e, 0x76, 0xa4, 0x38, 0x65, 0x3e, 0x28, 0xb2, 0xdb, 0x7d, 0x7c,
	0x8f, 0x88, 0x08, 0x64, 0xfa, 0x35, 0x30, 0xaa, 0x18, 0x3a, 0x93, 0xe0, 0x09, 0x2d, 0xb7, 0x55,
	0x6c, 0xba, 0x0d, 0xb7, 0xb1, 0x61, 0x3b, 0xad, 0xeb, 0x34, 0x8a, 0x42, 0xe6, 0x25, 0x07, 0xea,
	0x9f, 0x2b, 0xc1, 0xb7, 0x17, 0x2c, 0xdd, 0x3d, 0x54, 0x78, 0xa7, 0xd1, 0xc0, 0x1f, 0xa0, 0xf7,
	0xe7, 0x4b, 0xa5, 0x9c, 0xc4, 0x1c, 0x9e, 0x46, 0xe0, 0x19, 0x45, 0x41, 0x4a, 0x21, 0x89, 0xf0,
	0xbc, 0x58, 0xa6, 0x0d, 0x31, 0xa9, 0x15, 0xc8, 0x53, 0x90, 0x44, 0x31, 0x1f, 0xdc, 0xee, 0x89,
	0x91, 0xac, 0x81, 0x3f, 0x43, 0x9f, 0x5e, 0x26, 0x59, 0xda, 0xa0, 0x9e, 0xf0, 0xc7, 0x46, 0xb6,
	0x21, 0x0d, 0x03, 0x21, 0x87, 0x54, 0x1b, 0x68, 0x21, 0x89, 0x2f, 0x20, 0xd5, 0x72, 0x48, 0xb5,
	0x97, 0x4e, 0x44, 0x9e, 0x39, 0x3b, 0xeb, 0x76, 0xef, 0x9b, 0x04, 0x4d, 0xfc, 0x21, 0xda, 0xfd,
	0xeb, 0x04, 0x39, 0x90, 0x27, 0xb8, 0xa6, 0x2c, 0xdb, 0xf1, 0x58, 0x81, 0xbc, 0xad, 0x88, 0x27,
	0x21, 0x99, 0x1e, 0x1a, 0x2a, 0xb7, 0xdb, 0x31, 0x68, 0x77, 0xf1, 0x01, 0xda, 0x5f, 0x44, 0x33,
	0xf1, 0x53, 0xa8, 0x01, 0x3d, 0x05, 0x12, 0x81, 0x1c, 0x32, 0xa5, 0x98, 0xa9, 0x5c, 0x98, 0x71,
	0x05, 0xa5, 0xe6, 0xa6, 0xc1, 0xed, 0xfe, 0xf3, 0x99, 0x39, 0xfe, 0xdd, 0x46, 0xbf, 0xd9, 0xf9,
	0xd4, 0xfc, 0x6a, 0x97, 0x0b, 0xf8, 0xcc, 0x6e, 0x67, 0xe0, 0x22, 0x7f, 0x63, 0xe5, 0xc7, 0x95,
	0x39, 0x2f, 0x41, 0x69, 0xc9, 0x12, 0x6d, 0x0c, 0x95, 0x58, 0x0f, 0x4c, 0x51, 0x5e, 0xb2, 0x50,
	0x86, 0xb9, 0x72, 0xc9, 0x91, 0x59, 0xa9, 0xa9, 0x83, 0xa5, 0x0b, 0x94, 0xc0, 0x06, 0x22, 0x0c,
	0xc5, 0x28, 0xe5, 0x6e, 0xd2, 0x0a, 0xc9, 0xbe, 0x4c, 0x23, 0x76, 0x85, 0x0f, 0x24, 0x08, 0xc5,
	0xc8, 0xad, 0x2d, 0xb5, 0xca, 0xc9, 0x45, 0x22, 0xd6, 0x83, 0xed, 0xe5, 0xe4, 0x7e, 0x60, 0xbe,
	0x2c, 0x3b, 0xdb, 0xa8, 0x9a, 0x0e, 0x36, 0xc6, 0xfb, 0x92, 0x72, 0x6d, 0xb8, 0x50, 0x3f, 0x53,
	0x03, 0xad, 0xa1, 0xe2, 0x48, 0x32, 0x0d, 0xf8, 0x66, 0xe6, 0x4c, 0xfe, 0x65, 0xde, 0x8f, 0xec,
	0x8e, 0x75, 0x9c, 0xdf, 0x87, 0xbe, 0xb2, 0xad, 0x73, 0xdb, 0x7a, 0x66, 0x5b, 0x3f, 0xda, 0xd6,
	0x2f, 0xb6, 0xf5, 0x87, 0x6d, 0x3d, 0x77, 0xac, 0x5e, 0x29, 0xb9, 0x7b, 0xdc, 0xfd, 0x33, 0x00,
	0x00, 0xff, 0xff, 0x54, 0x8d, 0xae, 0xa3, 0x6b, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileServiceClient interface {
	// GetProfile retrieves user profiles.
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	// GetProfileCard retrieves cards.
	GetProfileCard(ctx context.Context, in *GetProfileCardsRequest, opts ...grpc.CallOption) (*GetProfileCardsResponse, error)
}

type profileServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfileServiceClient(cc *grpc.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfileCard(ctx context.Context, in *GetProfileCardsRequest, opts ...grpc.CallOption) (*GetProfileCardsResponse, error) {
	out := new(GetProfileCardsResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileService/GetProfileCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
type ProfileServiceServer interface {
	// GetProfile retrieves user profiles.
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	// GetProfileCard retrieves cards.
	GetProfileCard(context.Context, *GetProfileCardsRequest) (*GetProfileCardsResponse, error)
}

// UnimplementedProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (*UnimplementedProfileServiceServer) GetProfile(ctx context.Context, req *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (*UnimplementedProfileServiceServer) GetProfileCard(ctx context.Context, req *GetProfileCardsRequest) (*GetProfileCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileCard not implemented")
}

func RegisterProfileServiceServer(s *grpc.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfileCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfileCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileService/GetProfileCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfileCard(ctx, req.(*GetProfileCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profiles.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _ProfileService_GetProfile_Handler,
		},
		{
			MethodName: "GetProfileCard",
			Handler:    _ProfileService_GetProfileCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/profiles/all.proto",
}
