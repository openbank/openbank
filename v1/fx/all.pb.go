// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/fx/all.proto

package fx

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// FX holds all the details about a foreign exchange object.
type FX struct {
	// BankID is an identifier for the bank on this fx transaction.
	BankID string `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	// FromCurrencyCode is the currency to transfer from.
	FromCurrencyCode string `protobuf:"bytes,2,opt,name=FromCurrencyCode,json=from_currency_code,proto3" json:"FromCurrencyCode,omitempty"`
	// ToCurrencyCode is the currency that we are transferring to.
	ToCurrencyCode string `protobuf:"bytes,3,opt,name=ToCurrencyCode,json=to_currency_code,proto3" json:"ToCurrencyCode,omitempty"`
	// Rate is the exchange rate of the foreign exchange.
	Rate string `protobuf:"bytes,4,opt,name=Rate,json=rate,proto3" json:"Rate,omitempty"`
	// InverseRate is the inverse of the exchange rate of the foreign exchange.
	InverseRate string `protobuf:"bytes,5,opt,name=InverseRate,json=inverse_rate,proto3" json:"InverseRate,omitempty"`
	// EffectiveDate is the effective date of the foreign exchange quote.
	EffectiveDate        string   `protobuf:"bytes,6,opt,name=EffectiveDate,json=effective_date,proto3" json:"EffectiveDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FX) Reset()         { *m = FX{} }
func (m *FX) String() string { return proto.CompactTextString(m) }
func (*FX) ProtoMessage()    {}
func (*FX) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{0}
}

func (m *FX) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FX.Unmarshal(m, b)
}
func (m *FX) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FX.Marshal(b, m, deterministic)
}
func (m *FX) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FX.Merge(m, src)
}
func (m *FX) XXX_Size() int {
	return xxx_messageInfo_FX.Size(m)
}
func (m *FX) XXX_DiscardUnknown() {
	xxx_messageInfo_FX.DiscardUnknown(m)
}

var xxx_messageInfo_FX proto.InternalMessageInfo

func (m *FX) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *FX) GetFromCurrencyCode() string {
	if m != nil {
		return m.FromCurrencyCode
	}
	return ""
}

func (m *FX) GetToCurrencyCode() string {
	if m != nil {
		return m.ToCurrencyCode
	}
	return ""
}

func (m *FX) GetRate() string {
	if m != nil {
		return m.Rate
	}
	return ""
}

func (m *FX) GetInverseRate() string {
	if m != nil {
		return m.InverseRate
	}
	return ""
}

func (m *FX) GetEffectiveDate() string {
	if m != nil {
		return m.EffectiveDate
	}
	return ""
}

// CreateFXRequest is a request envelope to create a foreign exchange quote.
type CreateFXRequest struct {
	// FX is the foreign exchange information to create.
	FX                   *FX      `protobuf:"bytes,1,opt,name=FX,json=fx,proto3" json:"FX,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFXRequest) Reset()         { *m = CreateFXRequest{} }
func (m *CreateFXRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFXRequest) ProtoMessage()    {}
func (*CreateFXRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{1}
}

func (m *CreateFXRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFXRequest.Unmarshal(m, b)
}
func (m *CreateFXRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFXRequest.Marshal(b, m, deterministic)
}
func (m *CreateFXRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFXRequest.Merge(m, src)
}
func (m *CreateFXRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFXRequest.Size(m)
}
func (m *CreateFXRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFXRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFXRequest proto.InternalMessageInfo

func (m *CreateFXRequest) GetFX() *FX {
	if m != nil {
		return m.FX
	}
	return nil
}

// UpdateFXRequest is a request envelope to update a foreign exchange quote.
type UpdateFXRequest struct {
	// FX is the foreign exchange information to update.
	FX                   *FX      `protobuf:"bytes,1,opt,name=FX,json=fx,proto3" json:"FX,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFXRequest) Reset()         { *m = UpdateFXRequest{} }
func (m *UpdateFXRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFXRequest) ProtoMessage()    {}
func (*UpdateFXRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{2}
}

func (m *UpdateFXRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFXRequest.Unmarshal(m, b)
}
func (m *UpdateFXRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFXRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFXRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFXRequest.Merge(m, src)
}
func (m *UpdateFXRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFXRequest.Size(m)
}
func (m *UpdateFXRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFXRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFXRequest proto.InternalMessageInfo

func (m *UpdateFXRequest) GetFX() *FX {
	if m != nil {
		return m.FX
	}
	return nil
}

// GetCurrentFXRateRequest is a request envelope to get a foreign exchange rate.
type GetCurrentFXRateRequest struct {
	// FromCurrencyCode is the currency to transfer from.
	FromCurrencyCode string `protobuf:"bytes,1,opt,name=FromCurrencyCode,json=from_currency_code,proto3" json:"FromCurrencyCode,omitempty"`
	// ToCurrencyCode is the currency that we are transferring to.
	ToCurrencyCode       string   `protobuf:"bytes,2,opt,name=ToCurrencyCode,json=to_currency_code,proto3" json:"ToCurrencyCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentFXRateRequest) Reset()         { *m = GetCurrentFXRateRequest{} }
func (m *GetCurrentFXRateRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentFXRateRequest) ProtoMessage()    {}
func (*GetCurrentFXRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{3}
}

func (m *GetCurrentFXRateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentFXRateRequest.Unmarshal(m, b)
}
func (m *GetCurrentFXRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentFXRateRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentFXRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentFXRateRequest.Merge(m, src)
}
func (m *GetCurrentFXRateRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentFXRateRequest.Size(m)
}
func (m *GetCurrentFXRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentFXRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentFXRateRequest proto.InternalMessageInfo

func (m *GetCurrentFXRateRequest) GetFromCurrencyCode() string {
	if m != nil {
		return m.FromCurrencyCode
	}
	return ""
}

func (m *GetCurrentFXRateRequest) GetToCurrencyCode() string {
	if m != nil {
		return m.ToCurrencyCode
	}
	return ""
}

func init() {
	proto.RegisterType((*FX)(nil), "fx.FX")
	proto.RegisterType((*CreateFXRequest)(nil), "fx.CreateFXRequest")
	proto.RegisterType((*UpdateFXRequest)(nil), "fx.UpdateFXRequest")
	proto.RegisterType((*GetCurrentFXRateRequest)(nil), "fx.GetCurrentFXRateRequest")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/fx/all.proto", fileDescriptor_e3bdb6bf3dbfaace)
}

var fileDescriptor_e3bdb6bf3dbfaace = []byte{
	// 1167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6f, 0x1b, 0xc5,
	0x1b, 0xde, 0x5d, 0xa7, 0x6e, 0x3a, 0xed, 0xaf, 0xbf, 0x68, 0x40, 0x25, 0x72, 0x4b, 0x19, 0xa5,
	0x12, 0x0d, 0xa1, 0x99, 0x75, 0xdc, 0x54, 0xaa, 0x22, 0x55, 0x25, 0x49, 0xe3, 0x92, 0x08, 0x44,
	0x64, 0x4a, 0x65, 0x55, 0x42, 0xd1, 0x78, 0xf7, 0x5d, 0xef, 0x10, 0x7b, 0x66, 0x3b, 0x33, 0x6b,
	0x3b, 0x94, 0x4a, 0x7c, 0x49, 0x94, 0x5b, 0x09, 0x27, 0xce, 0x70, 0xe0, 0xc8, 0x91, 0x1b, 0x27,
	0xee, 0x95, 0xb8, 0xf0, 0x07, 0x70, 0xe3, 0x8e, 0x90, 0xb8, 0xa0, 0xd9, 0x5d, 0xbb, 0x89, 0x9d,
	0x14, 0xe8, 0xc9, 0x63, 0xcf, 0xf3, 0x3c, 0xf3, 0xbe, 0xcf, 0xfb, 0x61, 0x74, 0xa5, 0xcd, 0x4d,
	0x9c, 0xb6, 0x68, 0x20, 0xbb, 0xbe, 0x4c, 0x40, 0xb4, 0x98, 0xd8, 0x7d, 0x7a, 0xe8, 0x2d, 0xf9,
	0xd1, 0xc0, 0x67, 0x9d, 0x0e, 0x4d, 0x94, 0x34, 0x12, 0x7b, 0xd1, 0xa0, 0x72, 0xa1, 0x2d, 0x65,
	0xbb, 0x03, 0x3e, 0x4b, 0xb8, 0xcf, 0x84, 0x90, 0x86, 0x19, 0x2e, 0x85, 0xce, 0x11, 0x95, 0x2b,
	0xd9, 0x47, 0xb0, 0xd8, 0x06, 0xb1, 0xa8, 0xfb, 0xac, 0xdd, 0x06, 0xe5, 0xcb, 0x24, 0x43, 0x1c,
	0x81, 0x3e, 0x5f, 0x68, 0x65, 0xdf, 0x5a, 0x69, 0xe4, 0x43, 0x37, 0x31, 0x7b, 0xf9, 0xe5, 0xdc,
	0xb7, 0x1e, 0xf2, 0xea, 0x4d, 0x7c, 0x09, 0x95, 0xd7, 0x98, 0xd8, 0xdd, 0xbc, 0x35, 0xeb, 0x12,
	0x77, 0xfe, 0xd4, 0x1a, 0x9a, 0x76, 0x66, 0x9d, 0x79, 0xa7, 0xea, 0x6c, 0x3b, 0x8d, 0x93, 0x36,
	0xca, 0x1d, 0x1e, 0xe2, 0xeb, 0x68, 0xa6, 0xae, 0x64, 0x77, 0x3d, 0x55, 0x0a, 0x44, 0xb0, 0xb7,
	0x2e, 0x43, 0x98, 0xf5, 0x26, 0xe0, 0x38, 0x52, 0xb2, 0xbb, 0x13, 0x14, 0xa0, 0x9d, 0x40, 0x86,
	0x80, 0x97, 0xd1, 0xd9, 0x3b, 0xf2, 0x10, 0xaf, 0x34, 0xc1, 0x9b, 0x31, 0x72, 0x8c, 0x75, 0x11,
	0x4d, 0x35, 0x98, 0x81, 0xd9, 0xa9, 0x09, 0xec, 0x94, 0x62, 0x06, 0xf0, 0x22, 0x3a, 0xbd, 0x29,
	0x7a, 0xa0, 0x34, 0x64, 0xb0, 0x13, 0x13, 0xb0, 0x33, 0x3c, 0xbf, 0xde, 0xc9, 0xe0, 0x4b, 0xe8,
	0x7f, 0x1b, 0x51, 0x04, 0x81, 0xe1, 0x3d, 0xb8, 0x65, 0x09, 0xe5, 0x09, 0xc2, 0x59, 0x18, 0x02,
	0x76, 0x42, 0x66, 0x60, 0xa5, 0x3c, 0xed, 0xcc, 0x38, 0xb3, 0xce, 0xdc, 0x0d, 0xf4, 0xff, 0x75,
	0x05, 0xcc, 0x40, 0xbd, 0xd9, 0x80, 0xfb, 0x29, 0x68, 0x83, 0xe7, 0xac, 0x6f, 0x99, 0x5b, 0xa7,
	0x6b, 0x65, 0x1a, 0x0d, 0x68, 0xbd, 0x79, 0x48, 0xca, 0x8b, 0x06, 0x07, 0xe9, 0xef, 0x25, 0xe1,
	0x73, 0xd3, 0xbf, 0x72, 0xd1, 0x4b, 0xb7, 0xc1, 0xe4, 0xfe, 0x99, 0x7a, 0xd3, 0x66, 0x3b, 0xd4,
	0x39, 0xaa, 0x26, 0xee, 0x73, 0xd6, 0xc4, 0xfb, 0xe7, 0x9a, 0x0c, 0x63, 0xaa, 0xfd, 0x70, 0x12,
	0x9d, 0xaa, 0x37, 0xdf, 0x05, 0xd5, 0xe3, 0x01, 0xe0, 0xcf, 0x4a, 0x68, 0x66, 0x3c, 0x42, 0x7c,
	0xde, 0xa6, 0x75, 0x4c, 0xdc, 0x95, 0x22, 0xe7, 0xb9, 0x9f, 0xbd, 0xfd, 0xd5, 0x2f, 0xf3, 0x1e,
	0x5c, 0x6c, 0x80, 0x51, 0x1c, 0x7a, 0x40, 0x60, 0x10, 0xc4, 0x4c, 0xb4, 0x81, 0xd8, 0xba, 0x91,
	0x16, 0x98, 0x3e, 0x80, 0x20, 0xa6, 0x2f, 0x49, 0x11, 0x0c, 0x07, 0x5d, 0xb9, 0x36, 0x82, 0x9b,
	0x78, 0x9c, 0x62, 0xd3, 0x25, 0x6c, 0x88, 0xde, 0x23, 0x46, 0x12, 0x26, 0xa4, 0x89, 0x41, 0x6d,
	0xdd, 0x44, 0xa5, 0x5a, 0xb5, 0x8a, 0xaf, 0xa3, 0x8b, 0x45, 0x28, 0x04, 0x06, 0x10, 0xa4, 0x06,
	0x42, 0xa2, 0xd3, 0x20, 0x00, 0xad, 0xa3, 0xb4, 0xd3, 0xd9, 0xa3, 0xf8, 0x1c, 0x7a, 0xb1, 0x82,
	0x2f, 0xf9, 0x21, 0x44, 0x5c, 0xf0, 0x7c, 0xb6, 0xa2, 0x41, 0xbd, 0xb9, 0xb5, 0x84, 0x4a, 0xcb,
	0xd5, 0x65, 0xbc, 0x80, 0xe6, 0x1b, 0x60, 0x52, 0x25, 0x20, 0x24, 0xfd, 0xd8, 0x86, 0x17, 0x03,
	0x51, 0xa0, 0x65, 0xaa, 0x02, 0x20, 0x5c, 0x13, 0x21, 0x0d, 0x89, 0x64, 0x2a, 0x42, 0xda, 0x7a,
	0x0d, 0x5d, 0x46, 0xe5, 0x77, 0x56, 0x53, 0x13, 0xd7, 0xf0, 0xcb, 0xe8, 0x7c, 0x6c, 0x4c, 0xa2,
	0x57, 0x7c, 0x9f, 0xa5, 0x26, 0xa6, 0x2d, 0xb1, 0x4b, 0x8d, 0xf4, 0xa3, 0x01, 0x55, 0xc0, 0xc2,
	0x4f, 0x7f, 0xf9, 0xed, 0x6b, 0xef, 0x0a, 0x5e, 0x28, 0x96, 0xc2, 0x83, 0xf1, 0xda, 0x3e, 0xf4,
	0x1f, 0x1c, 0x2e, 0xda, 0xc3, 0x47, 0x9e, 0xf3, 0xd8, 0xcb, 0xea, 0x85, 0xbf, 0xf1, 0xd0, 0xf4,
	0xb0, 0x4d, 0xf1, 0x0b, 0xd6, 0xe0, 0xb1, 0xa6, 0x1d, 0xb9, 0xfe, 0xbb, 0xbb, 0xbf, 0xfa, 0xc4,
	0xcd, 0x5c, 0x7f, 0x25, 0xc7, 0x10, 0x46, 0x22, 0xa9, 0x80, 0xb7, 0xc5, 0x53, 0x2b, 0xef, 0xa7,
	0xd2, 0x40, 0x65, 0x39, 0x07, 0x68, 0xc2, 0x88, 0x80, 0xfe, 0x31, 0x28, 0xc2, 0x44, 0x48, 0x54,
	0xe6, 0x87, 0x26, 0xdc, 0xd0, 0xad, 0xdb, 0xd6, 0xe6, 0x25, 0xfc, 0x06, 0x7a, 0xb5, 0x5e, 0x10,
	0x36, 0x86, 0x84, 0x20, 0xd3, 0xfb, 0x97, 0x76, 0xb7, 0x16, 0xd0, 0xfc, 0xc8, 0xbb, 0x8b, 0xe8,
	0xc2, 0x31, 0xde, 0xf5, 0x15, 0x37, 0x90, 0x99, 0x77, 0x7a, 0xc5, 0x5d, 0x98, 0x2b, 0xe7, 0xfe,
	0x1d, 0xf0, 0xe6, 0x2f, 0x17, 0x4d, 0x0f, 0x67, 0x30, 0xf7, 0x66, 0x6c, 0x22, 0x2b, 0xe7, 0x68,
	0xbe, 0x27, 0xe9, 0x70, 0x4f, 0xd2, 0x0d, 0xbb, 0x27, 0xe7, 0x7e, 0x72, 0xf7, 0x57, 0xbf, 0x2b,
	0xbc, 0xca, 0x39, 0xc7, 0x7b, 0x45, 0x72, 0x80, 0x3e, 0x16, 0xb1, 0xe5, 0x5b, 0x5f, 0x96, 0xf1,
	0xfc, 0x11, 0xbe, 0x1c, 0xf4, 0x83, 0xa4, 0x99, 0x50, 0x48, 0x9f, 0x2b, 0xff, 0xca, 0x58, 0xfe,
	0x95, 0xd2, 0x23, 0xcf, 0x59, 0xfb, 0xbe, 0xbc, 0xbf, 0xfa, 0x45, 0x19, 0x95, 0x6a, 0xb4, 0x8a,
	0xef, 0xa2, 0x72, 0xbd, 0x49, 0x56, 0xb7, 0x37, 0xf1, 0xc6, 0xb6, 0x92, 0x3d, 0x1e, 0x82, 0x2e,
	0xea, 0x52, 0x54, 0x92, 0x85, 0x44, 0x26, 0xa0, 0xf2, 0x3f, 0x12, 0x22, 0xf3, 0xf6, 0x9e, 0xc8,
	0x69, 0xd8, 0xef, 0xb4, 0x76, 0x62, 0x89, 0x56, 0x69, 0x75, 0xc1, 0xf5, 0x6a, 0x33, 0x2c, 0x49,
	0x3a, 0x3c, 0xc8, 0x98, 0xfe, 0x07, 0x5a, 0x8a, 0x95, 0x89, 0x5f, 0x1a, 0x3b, 0x76, 0x86, 0xaa,
	0xb8, 0x89, 0xee, 0x1e, 0x35, 0x43, 0xf9, 0x58, 0xb6, 0x64, 0xb8, 0x67, 0xe7, 0xa8, 0xcb, 0x3a,
	0x91, 0x54, 0x5d, 0x66, 0x6c, 0xcf, 0x48, 0x45, 0x42, 0x09, 0xf9, 0x70, 0x75, 0x99, 0x09, 0xe2,
	0x62, 0xf8, 0x13, 0x08, 0xec, 0x75, 0xc1, 0xa5, 0x8d, 0xb7, 0xec, 0x03, 0x4b, 0x78, 0x03, 0xad,
	0x1f, 0xff, 0xc0, 0x48, 0x28, 0x90, 0xc2, 0x30, 0x2e, 0x74, 0x76, 0x9b, 0x6a, 0x50, 0x97, 0x33,
	0x33, 0x42, 0x10, 0x86, 0xb3, 0x8e, 0xa6, 0x8d, 0x6d, 0xab, 0x76, 0x15, 0x6f, 0xa2, 0xdb, 0x93,
	0x6a, 0x16, 0xff, 0x54, 0x2a, 0x66, 0x3d, 0x20, 0x09, 0xa8, 0x2e, 0xd7, 0x9a, 0x5b, 0xd7, 0x24,
	0x61, 0x59, 0x55, 0x0f, 0xad, 0x07, 0xda, 0xf8, 0xef, 0x4b, 0xa4, 0x51, 0x47, 0xa5, 0x6b, 0xd5,
	0x2a, 0xbe, 0x89, 0x6e, 0x1c, 0xa6, 0x30, 0x41, 0x52, 0x31, 0x72, 0x00, 0x94, 0x92, 0x8a, 0xc8,
	0x20, 0x48, 0x95, 0xb5, 0x2b, 0x57, 0xd4, 0xa0, 0x7a, 0xa0, 0x88, 0xe6, 0x21, 0xd0, 0x7b, 0x9f,
	0x7b, 0xe8, 0x13, 0x6f, 0xd4, 0x52, 0x7f, 0xb8, 0xd3, 0x25, 0xfc, 0xd1, 0x6a, 0x11, 0xa3, 0x24,
	0xd1, 0x60, 0xf4, 0xbe, 0xb6, 0x01, 0x28, 0xd0, 0x46, 0xf1, 0x4c, 0xda, 0xe6, 0x92, 0x9a, 0xd8,
	0xba, 0x12, 0x64, 0x03, 0x6c, 0x53, 0xd7, 0x94, 0xdc, 0x89, 0xe1, 0xe0, 0x85, 0x4d, 0x3b, 0x51,
	0x32, 0x13, 0x8c, 0x64, 0xa7, 0x23, 0xfb, 0x79, 0xf2, 0xf6, 0x3d, 0xa9, 0xf8, 0x87, 0x39, 0xc2,
	0xee, 0x2e, 0x12, 0x75, 0x64, 0x9f, 0xce, 0x4f, 0xd5, 0xa6, 0x6d, 0xbb, 0x5a, 0x89, 0x95, 0x53,
	0xf6, 0x64, 0xe4, 0x2e, 0x88, 0xb5, 0xf7, 0xd1, 0xeb, 0xcf, 0x5c, 0x95, 0xf8, 0xcc, 0x5d, 0x6e,
	0xf7, 0xd1, 0x80, 0x84, 0xcc, 0x30, 0x44, 0x9f, 0x3d, 0x1b, 0xf8, 0xec, 0xdb, 0x4c, 0xb0, 0x36,
	0x0c, 0xf1, 0x6f, 0xba, 0xdb, 0xce, 0x3d, 0x2f, 0x1a, 0x7c, 0xec, 0x3a, 0x8f, 0x5c, 0xe7, 0xb1,
	0xeb, 0xfc, 0xe8, 0x3a, 0xbf, 0xba, 0xce, 0x9f, 0xae, 0xf3, 0xc4, 0x73, 0x5a, 0xe5, 0x6c, 0x07,
	0x5c, 0xfd, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x49, 0x48, 0x43, 0xa1, 0xb9, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FXServiceClient is the client API for FXService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FXServiceClient interface {
	// GetCurrentFXRate retrieves the current foreign exchange rate of the two
	//specified currency code.
	GetCurrentFXRate(ctx context.Context, in *GetCurrentFXRateRequest, opts ...grpc.CallOption) (*FX, error)
	// CreateFX creates a new foreign exchange quote.
	CreateFX(ctx context.Context, in *CreateFXRequest, opts ...grpc.CallOption) (*FX, error)
	// UpdateFX updates a foreign exchange quote.
	UpdateFX(ctx context.Context, in *UpdateFXRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type fXServiceClient struct {
	cc *grpc.ClientConn
}

func NewFXServiceClient(cc *grpc.ClientConn) FXServiceClient {
	return &fXServiceClient{cc}
}

func (c *fXServiceClient) GetCurrentFXRate(ctx context.Context, in *GetCurrentFXRateRequest, opts ...grpc.CallOption) (*FX, error) {
	out := new(FX)
	err := c.cc.Invoke(ctx, "/fx.FXService/GetCurrentFXRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fXServiceClient) CreateFX(ctx context.Context, in *CreateFXRequest, opts ...grpc.CallOption) (*FX, error) {
	out := new(FX)
	err := c.cc.Invoke(ctx, "/fx.FXService/CreateFX", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fXServiceClient) UpdateFX(ctx context.Context, in *UpdateFXRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/fx.FXService/UpdateFX", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FXServiceServer is the server API for FXService service.
type FXServiceServer interface {
	// GetCurrentFXRate retrieves the current foreign exchange rate of the two
	//specified currency code.
	GetCurrentFXRate(context.Context, *GetCurrentFXRateRequest) (*FX, error)
	// CreateFX creates a new foreign exchange quote.
	CreateFX(context.Context, *CreateFXRequest) (*FX, error)
	// UpdateFX updates a foreign exchange quote.
	UpdateFX(context.Context, *UpdateFXRequest) (*empty.Empty, error)
}

// UnimplementedFXServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFXServiceServer struct {
}

func (*UnimplementedFXServiceServer) GetCurrentFXRate(ctx context.Context, req *GetCurrentFXRateRequest) (*FX, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentFXRate not implemented")
}
func (*UnimplementedFXServiceServer) CreateFX(ctx context.Context, req *CreateFXRequest) (*FX, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFX not implemented")
}
func (*UnimplementedFXServiceServer) UpdateFX(ctx context.Context, req *UpdateFXRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFX not implemented")
}

func RegisterFXServiceServer(s *grpc.Server, srv FXServiceServer) {
	s.RegisterService(&_FXService_serviceDesc, srv)
}

func _FXService_GetCurrentFXRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentFXRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FXServiceServer).GetCurrentFXRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.FXService/GetCurrentFXRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FXServiceServer).GetCurrentFXRate(ctx, req.(*GetCurrentFXRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FXService_CreateFX_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFXRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FXServiceServer).CreateFX(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.FXService/CreateFX",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FXServiceServer).CreateFX(ctx, req.(*CreateFXRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FXService_UpdateFX_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFXRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FXServiceServer).UpdateFX(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.FXService/UpdateFX",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FXServiceServer).UpdateFX(ctx, req.(*UpdateFXRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FXService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fx.FXService",
	HandlerType: (*FXServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentFXRate",
			Handler:    _FXService_GetCurrentFXRate_Handler,
		},
		{
			MethodName: "CreateFX",
			Handler:    _FXService_CreateFX_Handler,
		},
		{
			MethodName: "UpdateFX",
			Handler:    _FXService_UpdateFX_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/fx/all.proto",
}
