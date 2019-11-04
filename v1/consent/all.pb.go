// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/consent/all.proto

package consent

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type Status int32

const (
	Status_UnknownStatus   Status = 0
	Status_INITIATED       Status = 1
	Status_ACCEPTED        Status = 2
	Status_REJECTED        Status = 3
	Status_REVOKED         Status = 4
	Status_RECEIVED        Status = 5
	Status_VALID           Status = 6
	Status_REVOKEDBYPSU    Status = 7
	Status_EXPIRED         Status = 8
	Status_TERMINATEDBYTPP Status = 9
)

var Status_name = map[int32]string{
	0: "UnknownStatus",
	1: "INITIATED",
	2: "ACCEPTED",
	3: "REJECTED",
	4: "REVOKED",
	5: "RECEIVED",
	6: "VALID",
	7: "REVOKEDBYPSU",
	8: "EXPIRED",
	9: "TERMINATEDBYTPP",
}

var Status_value = map[string]int32{
	"UnknownStatus":   0,
	"INITIATED":       1,
	"ACCEPTED":        2,
	"REJECTED":        3,
	"REVOKED":         4,
	"RECEIVED":        5,
	"VALID":           6,
	"REVOKEDBYPSU":    7,
	"EXPIRED":         8,
	"TERMINATEDBYTPP": 9,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{0}
}

// AnswerConsentChallengeRequest is a request mesasge to answer consent challenge request
type AnswerConsentChallengeRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	ConsentID            string   `protobuf:"bytes,2,opt,name=ConsentID,json=consent_id,proto3" json:"ConsentID,omitempty"`
	Answer               string   `protobuf:"bytes,3,opt,name=Answer,json=answer,proto3" json:"Answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerConsentChallengeRequest) Reset()         { *m = AnswerConsentChallengeRequest{} }
func (m *AnswerConsentChallengeRequest) String() string { return proto.CompactTextString(m) }
func (*AnswerConsentChallengeRequest) ProtoMessage()    {}
func (*AnswerConsentChallengeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{0}
}

func (m *AnswerConsentChallengeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerConsentChallengeRequest.Unmarshal(m, b)
}
func (m *AnswerConsentChallengeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerConsentChallengeRequest.Marshal(b, m, deterministic)
}
func (m *AnswerConsentChallengeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerConsentChallengeRequest.Merge(m, src)
}
func (m *AnswerConsentChallengeRequest) XXX_Size() int {
	return xxx_messageInfo_AnswerConsentChallengeRequest.Size(m)
}
func (m *AnswerConsentChallengeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerConsentChallengeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerConsentChallengeRequest proto.InternalMessageInfo

func (m *AnswerConsentChallengeRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *AnswerConsentChallengeRequest) GetConsentID() string {
	if m != nil {
		return m.ConsentID
	}
	return ""
}

func (m *AnswerConsentChallengeRequest) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

// Consent is a response mesasge
type Consent struct {
	ConsentID            string   `protobuf:"bytes,1,opt,name=ConsentID,json=consent_id,proto3" json:"ConsentID,omitempty"`
	Jwt                  string   `protobuf:"bytes,2,opt,name=Jwt,json=jwt,proto3" json:"Jwt,omitempty"`
	Status               Status   `protobuf:"varint,3,opt,name=Status,json=status,proto3,enum=consent.Status" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consent) Reset()         { *m = Consent{} }
func (m *Consent) String() string { return proto.CompactTextString(m) }
func (*Consent) ProtoMessage()    {}
func (*Consent) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{1}
}

func (m *Consent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consent.Unmarshal(m, b)
}
func (m *Consent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consent.Marshal(b, m, deterministic)
}
func (m *Consent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consent.Merge(m, src)
}
func (m *Consent) XXX_Size() int {
	return xxx_messageInfo_Consent.Size(m)
}
func (m *Consent) XXX_DiscardUnknown() {
	xxx_messageInfo_Consent.DiscardUnknown(m)
}

var xxx_messageInfo_Consent proto.InternalMessageInfo

func (m *Consent) GetConsentID() string {
	if m != nil {
		return m.ConsentID
	}
	return ""
}

func (m *Consent) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func (m *Consent) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UnknownStatus
}

// CreateConsentEmailRequest
type CreateConsentEmailRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	For                  string   `protobuf:"bytes,2,opt,name=For,json=for,proto3" json:"For,omitempty"`
	View                 string   `protobuf:"bytes,3,opt,name=View,json=view,proto3" json:"View,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,json=email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConsentEmailRequest) Reset()         { *m = CreateConsentEmailRequest{} }
func (m *CreateConsentEmailRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConsentEmailRequest) ProtoMessage()    {}
func (*CreateConsentEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{2}
}

func (m *CreateConsentEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConsentEmailRequest.Unmarshal(m, b)
}
func (m *CreateConsentEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConsentEmailRequest.Marshal(b, m, deterministic)
}
func (m *CreateConsentEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConsentEmailRequest.Merge(m, src)
}
func (m *CreateConsentEmailRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConsentEmailRequest.Size(m)
}
func (m *CreateConsentEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConsentEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConsentEmailRequest proto.InternalMessageInfo

func (m *CreateConsentEmailRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *CreateConsentEmailRequest) GetFor() string {
	if m != nil {
		return m.For
	}
	return ""
}

func (m *CreateConsentEmailRequest) GetView() string {
	if m != nil {
		return m.View
	}
	return ""
}

func (m *CreateConsentEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// CreateConsentSMSRequest
type CreateConsentSMSRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	For                  string   `protobuf:"bytes,2,opt,name=For,json=for,proto3" json:"For,omitempty"`
	View                 string   `protobuf:"bytes,3,opt,name=View,json=view,proto3" json:"View,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=PhoneNumber,json=phone_number,proto3" json:"PhoneNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConsentSMSRequest) Reset()         { *m = CreateConsentSMSRequest{} }
func (m *CreateConsentSMSRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConsentSMSRequest) ProtoMessage()    {}
func (*CreateConsentSMSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{3}
}

func (m *CreateConsentSMSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConsentSMSRequest.Unmarshal(m, b)
}
func (m *CreateConsentSMSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConsentSMSRequest.Marshal(b, m, deterministic)
}
func (m *CreateConsentSMSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConsentSMSRequest.Merge(m, src)
}
func (m *CreateConsentSMSRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConsentSMSRequest.Size(m)
}
func (m *CreateConsentSMSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConsentSMSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConsentSMSRequest proto.InternalMessageInfo

func (m *CreateConsentSMSRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *CreateConsentSMSRequest) GetFor() string {
	if m != nil {
		return m.For
	}
	return ""
}

func (m *CreateConsentSMSRequest) GetView() string {
	if m != nil {
		return m.View
	}
	return ""
}

func (m *CreateConsentSMSRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

// GetconsentsRequest
type GetConsentsRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConsentsRequest) Reset()         { *m = GetConsentsRequest{} }
func (m *GetConsentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetConsentsRequest) ProtoMessage()    {}
func (*GetConsentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{4}
}

func (m *GetConsentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConsentsRequest.Unmarshal(m, b)
}
func (m *GetConsentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConsentsRequest.Marshal(b, m, deterministic)
}
func (m *GetConsentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConsentsRequest.Merge(m, src)
}
func (m *GetConsentsRequest) XXX_Size() int {
	return xxx_messageInfo_GetConsentsRequest.Size(m)
}
func (m *GetConsentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConsentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConsentsRequest proto.InternalMessageInfo

func (m *GetConsentsRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

// GetconsentsResponse
type GetconsentsResponse struct {
	Consents             []*Consent `protobuf:"bytes,1,rep,name=Consents,json=consents,proto3" json:"Consents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetconsentsResponse) Reset()         { *m = GetconsentsResponse{} }
func (m *GetconsentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetconsentsResponse) ProtoMessage()    {}
func (*GetconsentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{5}
}

func (m *GetconsentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetconsentsResponse.Unmarshal(m, b)
}
func (m *GetconsentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetconsentsResponse.Marshal(b, m, deterministic)
}
func (m *GetconsentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetconsentsResponse.Merge(m, src)
}
func (m *GetconsentsResponse) XXX_Size() int {
	return xxx_messageInfo_GetconsentsResponse.Size(m)
}
func (m *GetconsentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetconsentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetconsentsResponse proto.InternalMessageInfo

func (m *GetconsentsResponse) GetConsents() []*Consent {
	if m != nil {
		return m.Consents
	}
	return nil
}

// RevokeConsentRequest
type RevokeConsentRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	ConsentID            string   `protobuf:"bytes,2,opt,name=ConsentID,json=consent_id,proto3" json:"ConsentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeConsentRequest) Reset()         { *m = RevokeConsentRequest{} }
func (m *RevokeConsentRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeConsentRequest) ProtoMessage()    {}
func (*RevokeConsentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{6}
}

func (m *RevokeConsentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeConsentRequest.Unmarshal(m, b)
}
func (m *RevokeConsentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeConsentRequest.Marshal(b, m, deterministic)
}
func (m *RevokeConsentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeConsentRequest.Merge(m, src)
}
func (m *RevokeConsentRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeConsentRequest.Size(m)
}
func (m *RevokeConsentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeConsentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeConsentRequest proto.InternalMessageInfo

func (m *RevokeConsentRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *RevokeConsentRequest) GetConsentID() string {
	if m != nil {
		return m.ConsentID
	}
	return ""
}

func init() {
	proto.RegisterEnum("consent.Status", Status_name, Status_value)
	proto.RegisterType((*AnswerConsentChallengeRequest)(nil), "consent.AnswerConsentChallengeRequest")
	proto.RegisterType((*Consent)(nil), "consent.Consent")
	proto.RegisterType((*CreateConsentEmailRequest)(nil), "consent.CreateConsentEmailRequest")
	proto.RegisterType((*CreateConsentSMSRequest)(nil), "consent.CreateConsentSMSRequest")
	proto.RegisterType((*GetConsentsRequest)(nil), "consent.GetConsentsRequest")
	proto.RegisterType((*GetconsentsResponse)(nil), "consent.GetconsentsResponse")
	proto.RegisterType((*RevokeConsentRequest)(nil), "consent.RevokeConsentRequest")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/consent/all.proto", fileDescriptor_955aa7c027ca6560)
}

var fileDescriptor_955aa7c027ca6560 = []byte{
	// 1517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xdf, 0x6f, 0x13, 0xc9,
	0x1d, 0xdf, 0x5d, 0x27, 0x8e, 0x33, 0xe1, 0x87, 0x99, 0x40, 0x31, 0x26, 0x81, 0x21, 0x20, 0x1a,
	0xd2, 0x66, 0xed, 0x38, 0x14, 0x55, 0x91, 0xaa, 0xca, 0xb1, 0x4d, 0x70, 0x48, 0x88, 0xe5, 0x84,
	0x14, 0x78, 0x68, 0x18, 0xef, 0x8e, 0xed, 0x25, 0xeb, 0x19, 0x33, 0x33, 0x1b, 0x87, 0x22, 0xa4,
	0x0a, 0x55, 0x2a, 0x6f, 0x45, 0xa9, 0xfa, 0xd4, 0x4a, 0xed, 0x4b, 0xa5, 0x3e, 0xf4, 0xad, 0x2f,
	0x7d, 0xba, 0xbf, 0xe0, 0x4e, 0x3a, 0x09, 0x9d, 0x74, 0xba, 0xd3, 0x3d, 0x9c, 0xc4, 0x49, 0xfc,
	0x09, 0xf7, 0x78, 0x9a, 0xdd, 0xf5, 0xc6, 0xc6, 0x26, 0x3a, 0x74, 0xc7, 0x3d, 0x79, 0xfd, 0xfd,
	0xf1, 0xf9, 0x7e, 0xe6, 0xb3, 0xdf, 0xfd, 0xac, 0x0d, 0x16, 0x1a, 0x8e, 0x6c, 0x7a, 0x35, 0xd3,
	0x62, 0xad, 0x0c, 0x6b, 0x13, 0x5a, 0xc3, 0x74, 0xf7, 0xf0, 0x62, 0x6f, 0x21, 0x63, 0x31, 0x2a,
	0x08, 0x95, 0x19, 0xec, 0xba, 0x66, 0x9b, 0x33, 0xc9, 0xe0, 0x58, 0x18, 0x4a, 0x4f, 0x35, 0x18,
	0x6b, 0xb8, 0x24, 0x83, 0xdb, 0x4e, 0x06, 0x53, 0xca, 0x24, 0x96, 0x0e, 0xa3, 0x22, 0x28, 0x4b,
	0xff, 0xd2, 0xff, 0xb0, 0xe6, 0x1b, 0x84, 0xce, 0x8b, 0x0e, 0x6e, 0x34, 0x08, 0xcf, 0xb0, 0xb6,
	0x5f, 0x31, 0x58, 0x3d, 0xf3, 0x77, 0x1d, 0x4c, 0xe7, 0xa9, 0xe8, 0x10, 0x5e, 0x08, 0xd0, 0x0b,
	0x4d, 0xec, 0xba, 0x84, 0x36, 0x48, 0x95, 0x3c, 0xf6, 0x88, 0x90, 0xf0, 0x32, 0x88, 0x2f, 0x63,
	0xba, 0x5b, 0x2e, 0xa6, 0x74, 0xa4, 0xcf, 0x8e, 0x2f, 0x83, 0x84, 0x96, 0xd2, 0x66, 0xb5, 0xac,
	0x56, 0xd1, 0xaa, 0x63, 0x8a, 0xed, 0x8e, 0x63, 0xc3, 0x6b, 0x60, 0x3c, 0xec, 0x2f, 0x17, 0x53,
	0xc6, 0x40, 0x1d, 0x08, 0xa9, 0xab, 0xd2, 0x19, 0x10, 0x0f, 0x06, 0xa6, 0x62, 0x03, 0x75, 0x71,
	0xec, 0x67, 0x96, 0xe2, 0x09, 0x2d, 0xa9, 0xa5, 0xb4, 0x99, 0xbf, 0xe8, 0x60, 0x2c, 0xc4, 0xed,
	0x1f, 0xa1, 0x1f, 0x39, 0x62, 0x0a, 0xc4, 0x56, 0x3b, 0x72, 0x08, 0x8f, 0xd8, 0xa3, 0x8e, 0x84,
	0x8b, 0x20, 0xbe, 0x29, 0xb1, 0xf4, 0x84, 0x4f, 0xe0, 0x44, 0xee, 0xa4, 0x19, 0xb6, 0x9a, 0x41,
	0xb8, 0x9f, 0x91, 0xf0, 0x63, 0x11, 0xa3, 0xff, 0xe8, 0xe0, 0x5c, 0x81, 0x13, 0x2c, 0x49, 0x48,
	0xa6, 0xd4, 0xc2, 0x8e, 0xfb, 0x5e, 0x5a, 0x4d, 0x81, 0xd8, 0x4d, 0xc6, 0x87, 0xb1, 0xab, 0x33,
	0x0e, 0x2f, 0x80, 0x91, 0x6d, 0x87, 0x74, 0x86, 0x88, 0x33, 0xb2, 0xe7, 0x90, 0x0e, 0x44, 0x60,
	0xd4, 0x1f, 0x99, 0x1a, 0x19, 0x28, 0x18, 0x25, 0x2a, 0x11, 0x51, 0xfd, 0x9f, 0x0e, 0xce, 0xf6,
	0x51, 0xdd, 0x5c, 0xdf, 0xfc, 0x09, 0x89, 0xce, 0x83, 0x89, 0x4a, 0x93, 0x51, 0x72, 0xc7, 0x6b,
	0xd5, 0x08, 0x1f, 0x42, 0xf7, 0x58, 0x5b, 0xa5, 0x77, 0xa8, 0x9f, 0x8f, 0x58, 0xe7, 0x01, 0x5c,
	0x21, 0x32, 0x64, 0x2c, 0xde, 0x87, 0x6f, 0x04, 0xf1, 0x3b, 0x30, 0xb9, 0x42, 0xa4, 0x15, 0x41,
	0x88, 0xb6, 0xba, 0x84, 0xbf, 0x06, 0x89, 0x2e, 0x6c, 0x4a, 0x47, 0xb1, 0xd9, 0x89, 0x5c, 0x32,
	0xba, 0xf3, 0x61, 0xa2, 0x0f, 0x37, 0xd1, 0x45, 0x88, 0x80, 0x29, 0x38, 0x5d, 0x25, 0x7b, 0x6c,
	0xb7, 0x2b, 0xe8, 0x07, 0x7a, 0x44, 0xba, 0xf3, 0xe6, 0x3e, 0xd6, 0xbb, 0xab, 0x0a, 0xcf, 0x80,
	0xe3, 0x77, 0xe9, 0x2e, 0x65, 0x1d, 0x1a, 0x04, 0x92, 0x5a, 0xda, 0x48, 0x68, 0xf0, 0x14, 0x18,
	0x2f, 0xdf, 0x29, 0x6f, 0x95, 0xf3, 0x5b, 0xa5, 0x62, 0x52, 0xf7, 0x43, 0x49, 0x90, 0xc8, 0x17,
	0x0a, 0xa5, 0x8a, 0x8a, 0x18, 0xdd, 0x48, 0xb5, 0xb4, 0x5a, 0x2a, 0xa8, 0x48, 0xcc, 0x8f, 0x9c,
	0x04, 0x63, 0xd5, 0xd2, 0xf6, 0xc6, 0xed, 0x52, 0x31, 0x39, 0x72, 0x58, 0x52, 0x28, 0x95, 0xb7,
	0x4b, 0xc5, 0xe4, 0xa8, 0x1f, 0x39, 0x0e, 0x46, 0xb7, 0xf3, 0x6b, 0xe5, 0x62, 0x32, 0xee, 0x7f,
	0x3d, 0x0d, 0x8e, 0x85, 0x1d, 0xcb, 0xf7, 0x2b, 0x9b, 0x77, 0x93, 0x63, 0x5d, 0x9c, 0xd2, 0xbd,
	0x4a, 0xb9, 0x5a, 0x2a, 0x26, 0x13, 0x7e, 0xe0, 0x2c, 0x38, 0xb9, 0x55, 0xaa, 0xae, 0x97, 0xef,
	0x28, 0x42, 0xcb, 0xf7, 0xb7, 0x2a, 0x95, 0xe4, 0xb8, 0x4a, 0xa4, 0x8d, 0x94, 0x96, 0xfb, 0xe4,
	0x18, 0x38, 0xd1, 0x5d, 0x45, 0xc2, 0xf7, 0x1c, 0x8b, 0xc0, 0xaf, 0x0c, 0xf0, 0xb3, 0xe1, 0xf6,
	0x03, 0xaf, 0x46, 0x37, 0xe7, 0x48, 0x7f, 0x4a, 0x0f, 0xdc, 0xc4, 0x99, 0x3f, 0x19, 0x07, 0xf9,
	0x8f, 0x7a, 0x9c, 0xe3, 0x4a, 0x00, 0x80, 0x64, 0x93, 0xa0, 0xb0, 0x16, 0x71, 0xf2, 0x98, 0x78,
	0x42, 0x22, 0xab, 0x0b, 0x97, 0x7e, 0x47, 0x95, 0x9a, 0x71, 0x58, 0xb5, 0x5a, 0x05, 0xb1, 0x5c,
	0x76, 0x01, 0xde, 0x06, 0x57, 0x83, 0x72, 0x62, 0x1f, 0xdd, 0x00, 0x2f, 0x81, 0x8b, 0xe9, 0xe9,
	0xcb, 0x19, 0x9b, 0xd4, 0x1d, 0xea, 0x04, 0x2e, 0xdc, 0xdd, 0xab, 0x90, 0x5f, 0x6d, 0x12, 0x9c,
	0x02, 0xf1, 0x8d, 0xbc, 0x27, 0x9b, 0x39, 0x38, 0x06, 0x46, 0x3b, 0xdc, 0x91, 0xe4, 0xf9, 0xab,
	0xd7, 0x7f, 0x35, 0x6e, 0x2c, 0xe9, 0x73, 0x33, 0x0b, 0xea, 0x6d, 0xa0, 0x96, 0x48, 0x64, 0x9e,
	0x06, 0x5b, 0xf6, 0x2c, 0x42, 0xc8, 0x3c, 0x8d, 0x56, 0xea, 0x59, 0x26, 0x9a, 0xf9, 0xc2, 0xd0,
	0x5e, 0x1a, 0xfe, 0x5e, 0xc1, 0xff, 0x1a, 0x00, 0x0e, 0xda, 0x15, 0x9c, 0x39, 0xd4, 0xec, 0x5d,
	0x5e, 0x36, 0x44, 0xd7, 0x2f, 0xf5, 0x83, 0xfc, 0xbf, 0x7b, 0x74, 0x0d, 0xdd, 0x05, 0x61, 0x8a,
	0x7c, 0xe7, 0xe9, 0xaa, 0x90, 0x3e, 0x1f, 0x24, 0x04, 0xc2, 0x88, 0x92, 0x4e, 0x7f, 0x72, 0x75,
	0x23, 0x50, 0xf0, 0x16, 0xb8, 0x5c, 0xea, 0x8d, 0x23, 0xcb, 0xef, 0xb1, 0x91, 0xf0, 0x2c, 0x8b,
	0x08, 0x51, 0xf7, 0x5c, 0xf7, 0x89, 0xf9, 0x83, 0xe4, 0xbb, 0xaa, 0xe4, 0xbb, 0x74, 0x94, 0x7c,
	0x3e, 0xb3, 0x1e, 0xb9, 0xfe, 0x65, 0x80, 0xe4, 0xdb, 0x96, 0x09, 0xd1, 0x70, 0xb1, 0x0e, 0xdd,
	0x74, 0x88, 0x54, 0xaf, 0xf4, 0x83, 0xfc, 0x3f, 0x7b, 0xa4, 0x3a, 0x73, 0x28, 0x95, 0x68, 0x89,
	0x48, 0xa8, 0x73, 0xfd, 0x42, 0xf5, 0xa4, 0x56, 0xd7, 0x03, 0x99, 0x6e, 0x82, 0x4b, 0x9b, 0xeb,
	0x9b, 0x1f, 0x50, 0xa4, 0x2b, 0x4a, 0xa4, 0x8b, 0x47, 0x89, 0x24, 0x5a, 0xa2, 0x47, 0xa2, 0xd7,
	0x06, 0x98, 0xe8, 0x31, 0x68, 0x78, 0x3e, 0x3a, 0xfb, 0xa0, 0x6d, 0xa7, 0xa7, 0x7a, 0x93, 0x6f,
	0x1b, 0xf2, 0xcc, 0xdf, 0x8c, 0x83, 0xfc, 0x9b, 0x1e, 0x91, 0x4e, 0xad, 0x39, 0x42, 0x22, 0xec,
	0xba, 0x08, 0x5b, 0x16, 0xf3, 0xa8, 0x14, 0xe9, 0xdf, 0x57, 0x89, 0xf4, 0x38, 0x55, 0x02, 0xb9,
	0x2a, 0x69, 0x31, 0x2a, 0xb1, 0x43, 0x1d, 0xda, 0x40, 0x5e, 0x1b, 0x49, 0x86, 0x72, 0xd9, 0xa8,
	0xd8, 0x44, 0x0f, 0x29, 0xd9, 0x97, 0x3b, 0x42, 0x62, 0x2e, 0x1d, 0xda, 0xd8, 0x71, 0xa8, 0x4d,
	0xf6, 0x1f, 0x22, 0x0b, 0x53, 0x54, 0x23, 0xc8, 0x13, 0xc4, 0x46, 0x75, 0xc6, 0x51, 0x1b, 0x37,
	0x1c, 0xea, 0xff, 0x08, 0x32, 0x57, 0xb7, 0x95, 0xca, 0x59, 0xb8, 0x01, 0x2e, 0x84, 0x94, 0x11,
	0xd9, 0x27, 0x96, 0x37, 0x28, 0xf1, 0x3c, 0xf8, 0x45, 0xfa, 0xda, 0x70, 0x89, 0xfb, 0xce, 0x1d,
	0x1c, 0xad, 0x06, 0x41, 0x32, 0x92, 0x3b, 0x0e, 0x46, 0x38, 0xc1, 0xb6, 0xaf, 0xf6, 0x34, 0x3c,
	0x7f, 0x84, 0xd4, 0x3d, 0x32, 0xff, 0xc3, 0x00, 0xc7, 0xfb, 0xde, 0x35, 0x70, 0x3a, 0xd2, 0x72,
	0xd8, 0x3b, 0x68, 0xc8, 0x0e, 0x7e, 0xa6, 0x1f, 0xe4, 0x0f, 0x7a, 0xe4, 0x9d, 0x0c, 0xfa, 0x7c,
	0xbf, 0x8a, 0x36, 0x10, 0xf6, 0x04, 0xbb, 0xab, 0x57, 0x0e, 0x56, 0x6f, 0x19, 0x4c, 0x17, 0x22,
	0x5b, 0x53, 0x45, 0x3f, 0xe6, 0xda, 0x2d, 0xaa, 0xb5, 0x33, 0xbf, 0xaf, 0xb5, 0x05, 0xf3, 0x0f,
	0xe5, 0x49, 0xc7, 0x5e, 0x18, 0xda, 0xf2, 0x9f, 0xe3, 0x07, 0xf9, 0xaf, 0x47, 0x41, 0x2c, 0x67,
	0x66, 0xe1, 0x1a, 0x98, 0xe8, 0x92, 0xcd, 0x57, 0xca, 0xf0, 0x46, 0x85, 0xb3, 0x3d, 0xc7, 0x26,
	0x02, 0x15, 0xaa, 0x77, 0x8b, 0x88, 0xb5, 0x09, 0x0f, 0x7e, 0xf6, 0x22, 0x46, 0xfb, 0x0c, 0xbb,
	0x8d, 0xb9, 0x3a, 0x9e, 0x60, 0x1e, 0xb7, 0x88, 0x99, 0x1b, 0x5d, 0x30, 0xb3, 0x66, 0x76, 0x4e,
	0x37, 0x72, 0x49, 0xdc, 0x6e, 0xbb, 0x8e, 0xe5, 0x77, 0x65, 0x1e, 0x09, 0x46, 0x97, 0x06, 0x22,
	0xd5, 0x0a, 0x88, 0x5d, 0xcf, 0x2e, 0xc2, 0x32, 0x58, 0x09, 0xd6, 0x94, 0xd8, 0xa8, 0xd3, 0x24,
	0xc1, 0x00, 0x4f, 0x10, 0x8e, 0x6c, 0x46, 0x04, 0xa2, 0x4c, 0xa2, 0x26, 0xde, 0x23, 0xa8, 0x4d,
	0x78, 0xcb, 0x11, 0xc2, 0x51, 0x14, 0x98, 0x5a, 0x5b, 0x22, 0x84, 0x5f, 0x1b, 0xcd, 0xaf, 0x2e,
	0x28, 0xc4, 0xeb, 0x70, 0x0e, 0xcc, 0x0e, 0x22, 0x76, 0xab, 0x90, 0x13, 0x60, 0xd6, 0x99, 0x47,
	0x6d, 0xb3, 0x7a, 0x13, 0xc4, 0x7e, 0x95, 0xcd, 0xc2, 0xdf, 0x82, 0xdf, 0xf4, 0xb7, 0x60, 0x8a,
	0x3c, 0x4a, 0xf6, 0xdb, 0xc4, 0x52, 0xdb, 0x4c, 0x38, 0x67, 0x1c, 0x31, 0xcb, 0xf2, 0xd4, 0x2b,
	0x2b, 0x14, 0x41, 0x10, 0xbe, 0x47, 0x38, 0x12, 0x8e, 0x4d, 0xcc, 0xea, 0x8e, 0x1a, 0x9d, 0x85,
	0xf7, 0xc0, 0xf6, 0xb0, 0xd1, 0xc1, 0xc3, 0x51, 0x63, 0xf6, 0x13, 0x35, 0xbe, 0x85, 0xdd, 0x3a,
	0xe3, 0x2d, 0x2c, 0x15, 0x34, 0xeb, 0x39, 0x67, 0x0b, 0x4b, 0xab, 0xe9, 0xb7, 0x44, 0x93, 0xc3,
	0x5e, 0xb3, 0xba, 0xa6, 0x06, 0x2c, 0xc0, 0x12, 0x28, 0xbc, 0x7b, 0x40, 0x04, 0x14, 0x3e, 0xe8,
	0x22, 0xd2, 0xf2, 0xe7, 0x42, 0x99, 0x9f, 0x4d, 0xa8, 0x74, 0xb0, 0x2b, 0xcc, 0x07, 0x6f, 0x74,
	0xf0, 0x8d, 0x1e, 0xed, 0xd6, 0x17, 0x7a, 0x22, 0x06, 0x9f, 0xeb, 0xf9, 0x50, 0x53, 0x16, 0x99,
	0x42, 0xa4, 0x9a, 0x50, 0xbc, 0x39, 0x11, 0x92, 0x3b, 0x3e, 0x2d, 0x55, 0xe3, 0xc9, 0xa6, 0xc2,
	0xb3, 0x7c, 0x4b, 0x55, 0x43, 0x84, 0x89, 0xb6, 0x9a, 0xa4, 0x37, 0xa1, 0x6e, 0x56, 0x9b, 0x33,
	0x1f, 0xb6, 0xce, 0x5c, 0x97, 0x75, 0x02, 0x4a, 0x6a, 0x2c, 0xe3, 0xce, 0x1f, 0x82, 0x8a, 0x02,
	0xb3, 0x09, 0xaa, 0xbb, 0xac, 0x63, 0xce, 0x8e, 0xe4, 0x12, 0x6a, 0xab, 0x15, 0xc4, 0xd2, 0xb8,
	0xba, 0x92, 0x6c, 0x97, 0xd0, 0xe5, 0x25, 0x90, 0x0e, 0x6c, 0x00, 0xc2, 0x15, 0x8e, 0x03, 0x62,
	0xd8, 0x0e, 0x97, 0x00, 0x4c, 0x85, 0x8f, 0x06, 0x9c, 0x0c, 0x93, 0xfe, 0xb7, 0x30, 0x7b, 0x4b,
	0xaf, 0x68, 0x0f, 0xba, 0x7f, 0xfe, 0xfe, 0xa8, 0x6b, 0x2f, 0x74, 0xed, 0xa5, 0xae, 0xfd, 0x5f,
	0xd7, 0x3e, 0xd7, 0xb5, 0x6f, 0x75, 0xed, 0x53, 0x43, 0xab, 0xc5, 0xfd, 0x3f, 0x73, 0x8b, 0xdf,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x06, 0xd8, 0x4d, 0x96, 0x56, 0x0e, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConsentServiceClient is the client API for ConsentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsentServiceClient interface {
	// AnswerConsentChallenge sends the answer to the challenge for consnet challenge request
	AnswerConsentChallenge(ctx context.Context, in *AnswerConsentChallengeRequest, opts ...grpc.CallOption) (*Consent, error)
	// CreateConsentEmail creates a new email consent
	CreateConsentEmail(ctx context.Context, in *CreateConsentEmailRequest, opts ...grpc.CallOption) (*Consent, error)
	// CreateConsentSMS creates a new sms consent
	CreateConsentSMS(ctx context.Context, in *CreateConsentSMSRequest, opts ...grpc.CallOption) (*Consent, error)
	// GetConsents returns a list containing up to 20 consents.
	GetConsents(ctx context.Context, in *GetConsentsRequest, opts ...grpc.CallOption) (*GetconsentsResponse, error)
	// RevokeConsent revokes the consent
	RevokeConsent(ctx context.Context, in *RevokeConsentRequest, opts ...grpc.CallOption) (*Consent, error)
}

type consentServiceClient struct {
	cc *grpc.ClientConn
}

func NewConsentServiceClient(cc *grpc.ClientConn) ConsentServiceClient {
	return &consentServiceClient{cc}
}

func (c *consentServiceClient) AnswerConsentChallenge(ctx context.Context, in *AnswerConsentChallengeRequest, opts ...grpc.CallOption) (*Consent, error) {
	out := new(Consent)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/AnswerConsentChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentServiceClient) CreateConsentEmail(ctx context.Context, in *CreateConsentEmailRequest, opts ...grpc.CallOption) (*Consent, error) {
	out := new(Consent)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/CreateConsentEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentServiceClient) CreateConsentSMS(ctx context.Context, in *CreateConsentSMSRequest, opts ...grpc.CallOption) (*Consent, error) {
	out := new(Consent)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/CreateConsentSMS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentServiceClient) GetConsents(ctx context.Context, in *GetConsentsRequest, opts ...grpc.CallOption) (*GetconsentsResponse, error) {
	out := new(GetconsentsResponse)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/GetConsents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentServiceClient) RevokeConsent(ctx context.Context, in *RevokeConsentRequest, opts ...grpc.CallOption) (*Consent, error) {
	out := new(Consent)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/RevokeConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsentServiceServer is the server API for ConsentService service.
type ConsentServiceServer interface {
	// AnswerConsentChallenge sends the answer to the challenge for consnet challenge request
	AnswerConsentChallenge(context.Context, *AnswerConsentChallengeRequest) (*Consent, error)
	// CreateConsentEmail creates a new email consent
	CreateConsentEmail(context.Context, *CreateConsentEmailRequest) (*Consent, error)
	// CreateConsentSMS creates a new sms consent
	CreateConsentSMS(context.Context, *CreateConsentSMSRequest) (*Consent, error)
	// GetConsents returns a list containing up to 20 consents.
	GetConsents(context.Context, *GetConsentsRequest) (*GetconsentsResponse, error)
	// RevokeConsent revokes the consent
	RevokeConsent(context.Context, *RevokeConsentRequest) (*Consent, error)
}

// UnimplementedConsentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConsentServiceServer struct {
}

func (*UnimplementedConsentServiceServer) AnswerConsentChallenge(ctx context.Context, req *AnswerConsentChallengeRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerConsentChallenge not implemented")
}
func (*UnimplementedConsentServiceServer) CreateConsentEmail(ctx context.Context, req *CreateConsentEmailRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsentEmail not implemented")
}
func (*UnimplementedConsentServiceServer) CreateConsentSMS(ctx context.Context, req *CreateConsentSMSRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsentSMS not implemented")
}
func (*UnimplementedConsentServiceServer) GetConsents(ctx context.Context, req *GetConsentsRequest) (*GetconsentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsents not implemented")
}
func (*UnimplementedConsentServiceServer) RevokeConsent(ctx context.Context, req *RevokeConsentRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeConsent not implemented")
}

func RegisterConsentServiceServer(s *grpc.Server, srv ConsentServiceServer) {
	s.RegisterService(&_ConsentService_serviceDesc, srv)
}

func _ConsentService_AnswerConsentChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerConsentChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).AnswerConsentChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/AnswerConsentChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).AnswerConsentChallenge(ctx, req.(*AnswerConsentChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsentService_CreateConsentEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConsentEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).CreateConsentEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/CreateConsentEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).CreateConsentEmail(ctx, req.(*CreateConsentEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsentService_CreateConsentSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConsentSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).CreateConsentSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/CreateConsentSMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).CreateConsentSMS(ctx, req.(*CreateConsentSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsentService_GetConsents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).GetConsents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/GetConsents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).GetConsents(ctx, req.(*GetConsentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsentService_RevokeConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeConsentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).RevokeConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/RevokeConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).RevokeConsent(ctx, req.(*RevokeConsentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consent.ConsentService",
	HandlerType: (*ConsentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnswerConsentChallenge",
			Handler:    _ConsentService_AnswerConsentChallenge_Handler,
		},
		{
			MethodName: "CreateConsentEmail",
			Handler:    _ConsentService_CreateConsentEmail_Handler,
		},
		{
			MethodName: "CreateConsentSMS",
			Handler:    _ConsentService_CreateConsentSMS_Handler,
		},
		{
			MethodName: "GetConsents",
			Handler:    _ConsentService_GetConsents_Handler,
		},
		{
			MethodName: "RevokeConsent",
			Handler:    _ConsentService_RevokeConsent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/consent/all.proto",
}
