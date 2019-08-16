// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/statements/all.proto

package statements

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	types "github.com/openbank/openbank/v1/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Status define the status of a statment.
type Status int32

const (
	Status_UnknownStatus Status = 0
	// Status_Received is the value for a received statment.
	Status_Received Status = 1
	// Status_Pending is the value for a pending statement.
	Status_Pending Status = 2
)

var Status_name = map[int32]string{
	0: "UnknownStatus",
	1: "Received",
	2: "Pending",
}

var Status_value = map[string]int32{
	"UnknownStatus": 0,
	"Received":      1,
	"Pending":       2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1cebb344c4e7431, []int{0}
}

// Statement holds details regarding a customer's statement.
type Statement struct {
	// ID is the unique identifier of a statement.
	ID string `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	// Status is the status of the statement.
	Status Status `protobuf:"varint,2,opt,name=Status,json=status,proto3,enum=statements.Status" json:"Status,omitempty"`
	// Date is the date of the statement.
	Date string `protobuf:"bytes,3,opt,name=Date,json=date,proto3" json:"Date,omitempty"`
	// Description is the description of the statement.
	Description string `protobuf:"bytes,4,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// Amount is the amount if the transcation that writes on the statement.
	Amount *types.Amount `protobuf:"bytes,5,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	// Balance is remaining balance after related transaction.
	Balance              *types.Amount `protobuf:"bytes,6,opt,name=Balance,json=balance,proto3" json:"Balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cebb344c4e7431, []int{0}
}

func (m *Statement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statement.Unmarshal(m, b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
}
func (m *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(m, src)
}
func (m *Statement) XXX_Size() int {
	return xxx_messageInfo_Statement.Size(m)
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

func (m *Statement) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Statement) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UnknownStatus
}

func (m *Statement) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Statement) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Statement) GetAmount() *types.Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Statement) GetBalance() *types.Amount {
	if m != nil {
		return m.Balance
	}
	return nil
}

// GetStatementRequest is the request envelope to retrieve a statement.
type GetStatementRequest struct {
	// ID is the unique identifier of the statement
	ID                   string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatementRequest) Reset()         { *m = GetStatementRequest{} }
func (m *GetStatementRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatementRequest) ProtoMessage()    {}
func (*GetStatementRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cebb344c4e7431, []int{1}
}

func (m *GetStatementRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatementRequest.Unmarshal(m, b)
}
func (m *GetStatementRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatementRequest.Marshal(b, m, deterministic)
}
func (m *GetStatementRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatementRequest.Merge(m, src)
}
func (m *GetStatementRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatementRequest.Size(m)
}
func (m *GetStatementRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatementRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatementRequest proto.InternalMessageInfo

func (m *GetStatementRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

// GetStatementsRequest is the request envelope to retrieve several statements.
type GetStatementsRequest struct {
	// BankCode is a bank code to get related statement.
	BankCode types.BankCode `protobuf:"varint,1,opt,name=BankCode,json=bank_code,proto3,enum=types.BankCode" json:"BankCode,omitempty"`
	// PeriodDays is statement period to be query.
	PeriodDays int32 `protobuf:"varint,2,opt,name=PeriodDays,json=period_days,proto3" json:"PeriodDays,omitempty"`
	// Status is status of the statement to be query.
	Status               Status   `protobuf:"varint,3,opt,name=Status,json=status,proto3,enum=statements.Status" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatementsRequest) Reset()         { *m = GetStatementsRequest{} }
func (m *GetStatementsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatementsRequest) ProtoMessage()    {}
func (*GetStatementsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cebb344c4e7431, []int{2}
}

func (m *GetStatementsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatementsRequest.Unmarshal(m, b)
}
func (m *GetStatementsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatementsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatementsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatementsRequest.Merge(m, src)
}
func (m *GetStatementsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatementsRequest.Size(m)
}
func (m *GetStatementsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatementsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatementsRequest proto.InternalMessageInfo

func (m *GetStatementsRequest) GetBankCode() types.BankCode {
	if m != nil {
		return m.BankCode
	}
	return types.BankCode_UnknownBankCode
}

func (m *GetStatementsRequest) GetPeriodDays() int32 {
	if m != nil {
		return m.PeriodDays
	}
	return 0
}

func (m *GetStatementsRequest) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UnknownStatus
}

// GetStatementResponse is the response envelope for retrieving
type GetStatementsResponse struct {
	// Result is a list containing the statements.
	Result []*Statement `protobuf:"bytes,1,rep,name=Result,json=result,proto3" json:"Result,omitempty"`
	// HasMore indicates if there are more statements available.
	HasMore              bool     `protobuf:"varint,2,opt,name=HasMore,json=has_more,proto3" json:"HasMore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatementsResponse) Reset()         { *m = GetStatementsResponse{} }
func (m *GetStatementsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatementsResponse) ProtoMessage()    {}
func (*GetStatementsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cebb344c4e7431, []int{3}
}

func (m *GetStatementsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatementsResponse.Unmarshal(m, b)
}
func (m *GetStatementsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatementsResponse.Marshal(b, m, deterministic)
}
func (m *GetStatementsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatementsResponse.Merge(m, src)
}
func (m *GetStatementsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatementsResponse.Size(m)
}
func (m *GetStatementsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatementsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatementsResponse proto.InternalMessageInfo

func (m *GetStatementsResponse) GetResult() []*Statement {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetStatementsResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func init() {
	proto.RegisterEnum("statements.Status", Status_name, Status_value)
	proto.RegisterType((*Statement)(nil), "statements.Statement")
	proto.RegisterType((*GetStatementRequest)(nil), "statements.GetStatementRequest")
	proto.RegisterType((*GetStatementsRequest)(nil), "statements.GetStatementsRequest")
	proto.RegisterType((*GetStatementsResponse)(nil), "statements.GetStatementsResponse")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/statements/all.proto", fileDescriptor_e1cebb344c4e7431)
}

var fileDescriptor_e1cebb344c4e7431 = []byte{
	// 1192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xde, 0x5d, 0xc7, 0x8e, 0x33, 0xa1, 0xad, 0x35, 0x6d, 0x24, 0xcb, 0xaa, 0xca, 0x34, 0x55,
	0x85, 0x55, 0xda, 0xb5, 0xb3, 0x0d, 0x88, 0x46, 0x42, 0xc8, 0xa9, 0x69, 0x9a, 0x8a, 0x82, 0xe5,
	0x52, 0x84, 0x7a, 0x49, 0xc7, 0xbb, 0xcf, 0xf6, 0x90, 0xf5, 0xcc, 0x32, 0x33, 0xeb, 0x24, 0x20,
	0xa4, 0x82, 0x38, 0x54, 0x9c, 0xaa, 0x72, 0x46, 0xe2, 0x2f, 0x70, 0xe3, 0xc2, 0x7f, 0x40, 0xe2,
	0x82, 0x10, 0x47, 0x10, 0x12, 0x12, 0xe2, 0x8a, 0xc4, 0x05, 0xcd, 0xae, 0xbd, 0x5e, 0x37, 0x6e,
	0x23, 0x4e, 0x71, 0xde, 0xfb, 0xde, 0x37, 0xef, 0xfb, 0xde, 0xdb, 0xd9, 0x45, 0x9b, 0x03, 0xa6,
	0x87, 0x71, 0xcf, 0xf5, 0xc5, 0xa8, 0x21, 0x22, 0xe0, 0x3d, 0xca, 0xf7, 0x67, 0x3f, 0xc6, 0x1b,
	0x0d, 0xa5, 0xa9, 0x86, 0x11, 0x70, 0xad, 0x1a, 0x34, 0x0c, 0xdd, 0x48, 0x0a, 0x2d, 0x30, 0x9a,
	0x45, 0x6b, 0xe7, 0x07, 0x42, 0x0c, 0x42, 0x68, 0xd0, 0x88, 0x35, 0x28, 0xe7, 0x42, 0x53, 0xcd,
	0x04, 0x57, 0x29, 0xb2, 0x76, 0x35, 0xf9, 0xe3, 0x5f, 0x1b, 0x00, 0xbf, 0xa6, 0x0e, 0xe8, 0x60,
	0x00, 0xb2, 0x21, 0xa2, 0x04, 0xb1, 0x00, 0xdd, 0x38, 0xa9, 0x1b, 0x7d, 0x14, 0x41, 0xae, 0x91,
	0xf5, 0x6f, 0x1c, 0xb4, 0x72, 0x6f, 0xda, 0x0b, 0xae, 0x21, 0x67, 0xb7, 0x5d, 0xb5, 0x89, 0x5d,
	0x5f, 0xd9, 0x46, 0x65, 0xab, 0x6a, 0xd5, 0xad, 0xa6, 0xd5, 0xb1, 0xba, 0x0e, 0x0b, 0xf0, 0xeb,
	0xa8, 0x64, 0x80, 0xb1, 0xaa, 0x3a, 0xc4, 0xae, 0x9f, 0xf6, 0xb0, 0x3b, 0xd3, 0xe0, 0xa6, 0x99,
	0xb9, 0x9a, 0x92, 0x4a, 0x62, 0xf8, 0x02, 0x5a, 0x6a, 0x53, 0x0d, 0xd5, 0xc2, 0x31, 0xd6, 0xa5,
	0x80, 0x6a, 0xc0, 0x57, 0xd1, 0x6a, 0x1b, 0x94, 0x2f, 0x59, 0x22, 0xaa, 0xba, 0x74, 0x0c, 0xb6,
	0x1a, 0xcc, 0xd2, 0x78, 0x03, 0x95, 0x5a, 0x23, 0x11, 0x73, 0x5d, 0x2d, 0x12, 0xbb, 0xbe, 0xea,
	0x9d, 0x72, 0x13, 0x45, 0x6e, 0x1a, 0x9c, 0x6f, 0x80, 0x26, 0x31, 0x7c, 0x1d, 0x2d, 0x6f, 0xd3,
	0x90, 0x72, 0x1f, 0xaa, 0xa5, 0x93, 0x6a, 0x96, 0x7b, 0x29, 0x72, 0xab, 0x54, 0xb6, 0x2a, 0x56,
	0xd5, 0x5a, 0xbf, 0x81, 0xce, 0xee, 0x80, 0xce, 0x1c, 0xea, 0xc2, 0xc7, 0x31, 0xa8, 0x17, 0x1a,
	0x95, 0x95, 0xfe, 0x60, 0xa3, 0x73, 0xf9, 0x5a, 0x35, 0x2d, 0x7e, 0x03, 0x95, 0xb7, 0x29, 0xdf,
	0xbf, 0x29, 0x02, 0x48, 0x28, 0x4e, 0x7b, 0x67, 0x26, 0x1d, 0x4d, 0xc3, 0x73, 0x9c, 0x2b, 0x66,
	0x76, 0x7b, 0xbe, 0x08, 0x00, 0xbf, 0x8a, 0x50, 0x07, 0x24, 0x13, 0x41, 0x9b, 0x1e, 0xa5, 0x73,
	0x28, 0xce, 0x5b, 0x15, 0x25, 0xd9, 0xbd, 0x80, 0x1e, 0xa9, 0xdc, 0xc0, 0x0a, 0xff, 0x67, 0x60,
	0x59, 0xff, 0x9f, 0xdb, 0x68, 0xed, 0x99, 0xfe, 0x55, 0x24, 0xb8, 0x02, 0x7c, 0x03, 0x95, 0xba,
	0xa0, 0xe2, 0x50, 0x57, 0x6d, 0x52, 0xa8, 0xaf, 0x7a, 0x6b, 0xcf, 0x32, 0x27, 0x3f, 0xe7, 0xc9,
	0x65, 0x52, 0x80, 0x2f, 0xa3, 0xe5, 0xdb, 0x54, 0xdd, 0x15, 0x12, 0x92, 0xf6, 0xcb, 0x73, 0xa0,
	0xf2, 0x90, 0xaa, 0xbd, 0x91, 0x90, 0x99, 0xfd, 0x57, 0x6e, 0x4d, 0x35, 0xe0, 0x35, 0x74, 0xea,
	0x3e, 0xdf, 0xe7, 0xe2, 0x80, 0xa7, 0x81, 0x8a, 0x55, 0x73, 0xca, 0x16, 0xae, 0xa0, 0x72, 0x17,
	0x7c, 0x60, 0x63, 0x08, 0x2a, 0x76, 0x12, 0x39, 0x83, 0x96, 0x3b, 0xc0, 0x03, 0xc6, 0x07, 0x15,
	0xc7, 0x04, 0x6a, 0x4e, 0xd5, 0xf2, 0xfe, 0x2e, 0xa2, 0x4a, 0xd6, 0xd8, 0x3d, 0x90, 0x63, 0xe6,
	0x03, 0xfe, 0xcb, 0x41, 0x2f, 0xe5, 0x05, 0xe2, 0x97, 0xf3, 0x3a, 0x16, 0x8c, 0xbd, 0xb6, 0x58,
	0xe8, 0xfa, 0xb7, 0xce, 0xd3, 0xd6, 0xbf, 0x76, 0xfe, 0x41, 0x3a, 0xd7, 0x05, 0x2d, 0x19, 0x8c,
	0x81, 0x50, 0x92, 0x95, 0xd4, 0x5a, 0x79, 0x4e, 0x22, 0x27, 0x10, 0x65, 0x30, 0x11, 0xf8, 0xac,
	0xcf, 0xfc, 0x19, 0xf8, 0xea, 0x34, 0x06, 0x01, 0xe9, 0x1d, 0x11, 0xa6, 0x15, 0xd9, 0x6d, 0xbb,
	0x77, 0x36, 0x50, 0x61, 0xb3, 0xb9, 0x89, 0xaf, 0xa0, 0x7a, 0x17, 0x74, 0x2c, 0x39, 0x04, 0xe4,
	0x60, 0x08, 0x9c, 0xe8, 0x21, 0x10, 0x09, 0x4a, 0xc4, 0xd2, 0x07, 0xc2, 0x14, 0xe1, 0x42, 0x93,
	0xbe, 0x88, 0x79, 0xe0, 0xde, 0x79, 0x17, 0x15, 0xbc, 0x66, 0x13, 0xef, 0xa0, 0x0b, 0x13, 0x19,
	0x04, 0x0e, 0xc1, 0x8f, 0x35, 0x04, 0x44, 0xc5, 0xbe, 0x0f, 0x4a, 0xf5, 0xe3, 0x30, 0x3c, 0x72,
	0xf1, 0x65, 0x74, 0xa9, 0x76, 0xf1, 0x52, 0x23, 0x80, 0x3e, 0xe3, 0x2c, 0xbd, 0x61, 0x66, 0x72,
	0xb3, 0xbe, 0x7b, 0x18, 0x55, 0x50, 0xe9, 0xbd, 0x56, 0xac, 0x87, 0x1e, 0x2e, 0xa1, 0x25, 0x09,
	0x34, 0xf8, 0xe2, 0xa7, 0xdf, 0xbe, 0x76, 0xd6, 0xf0, 0xd9, 0x67, 0xee, 0xbb, 0x4f, 0x77, 0xdb,
	0x9f, 0x3d, 0x76, 0xac, 0x27, 0x4e, 0x32, 0x60, 0xfc, 0xbb, 0x83, 0x4e, 0xcd, 0x2d, 0x14, 0x26,
	0xcf, 0x33, 0x7c, 0xfa, 0xac, 0xd4, 0x2e, 0xbe, 0x00, 0x91, 0x6e, 0xe3, 0xfa, 0x57, 0xce, 0xd3,
	0xd6, 0xa3, 0xb9, 0x6b, 0xec, 0x7c, 0xe6, 0xfe, 0x28, 0x0e, 0x35, 0x8b, 0x42, 0x20, 0xb9, 0xfb,
	0x16, 0x52, 0xe7, 0x8c, 0xed, 0x21, 0x53, 0x9a, 0xf8, 0x82, 0x6b, 0xca, 0x38, 0xe3, 0x03, 0x12,
	0x47, 0x44, 0x0b, 0xe2, 0x35, 0x73, 0x70, 0x97, 0x3c, 0xe4, 0x70, 0xa8, 0xf7, 0x94, 0xa6, 0x52,
	0x33, 0x3e, 0xd8, 0x63, 0x3c, 0x80, 0xc3, 0x87, 0xc4, 0xa7, 0x9c, 0xf4, 0x20, 0x37, 0xa4, 0xbe,
	0x90, 0x24, 0xa2, 0x03, 0xc6, 0x93, 0x0b, 0xd9, 0xbd, 0xf3, 0x20, 0xb5, 0xfd, 0xde, 0x89, 0xb6,
	0x6f, 0xa0, 0x46, 0xed, 0xda, 0xf3, 0x6c, 0x5f, 0x28, 0x79, 0xd1, 0x08, 0x66, 0x3e, 0xd7, 0x0a,
	0x8f, 0x1d, 0x6b, 0xfb, 0x97, 0xe2, 0xd3, 0xd6, 0x77, 0x45, 0x54, 0xf0, 0xdc, 0x26, 0xbe, 0x8b,
	0x4e, 0xcf, 0x48, 0x48, 0xab, 0xb3, 0x8b, 0x37, 0x3b, 0x52, 0x8c, 0x59, 0x00, 0x8a, 0xdc, 0xec,
	0xde, 0x6f, 0x13, 0x11, 0x81, 0x4c, 0xdf, 0x23, 0x44, 0xa4, 0xbb, 0x34, 0xeb, 0x20, 0x5b, 0x2b,
	0xd7, 0x2b, 0x6e, 0xb8, 0x4d, 0xb7, 0x79, 0xc5, 0x76, 0xbc, 0x0a, 0x8d, 0xa2, 0x90, 0xf9, 0x49,
	0x4d, 0xe3, 0x23, 0x25, 0xf8, 0xd6, 0xb1, 0x48, 0x77, 0xcf, 0xac, 0x6a, 0x13, 0x7f, 0x88, 0x3e,
	0x58, 0xb4, 0xaa, 0xa9, 0x25, 0x3d, 0x11, 0x1c, 0x99, 0x75, 0x1d, 0xd1, 0xb0, 0x2f, 0xe4, 0x88,
	0x6a, 0x63, 0x8f, 0x90, 0x24, 0x10, 0x90, 0xee, 0xf0, 0x88, 0x6a, 0x7f, 0x98, 0x94, 0xc0, 0x61,
	0x04, 0xbe, 0x49, 0x4f, 0x6a, 0xdd, 0xee, 0x3b, 0xe6, 0x80, 0x0d, 0xfc, 0x36, 0xba, 0xf9, 0xfc,
	0x03, 0x32, 0xa2, 0xc9, 0x94, 0x55, 0x92, 0x8d, 0x15, 0xc8, 0x57, 0x14, 0xf1, 0x25, 0x04, 0xc0,
	0x35, 0xa3, 0xa1, 0x72, 0xbb, 0x1d, 0xc3, 0x76, 0x1d, 0xef, 0xa2, 0x9d, 0xe3, 0x6c, 0x06, 0x3f,
	0xa3, 0x1a, 0xd2, 0x31, 0x90, 0x08, 0xe4, 0x88, 0x29, 0xc5, 0x8c, 0x5f, 0x82, 0xd0, 0x64, 0xa0,
	0x73, 0x4f, 0xa1, 0xdb, 0xbd, 0x85, 0x0a, 0xaf, 0x35, 0x9b, 0xf8, 0x2d, 0xf4, 0xe6, 0x3c, 0x23,
	0xe5, 0x24, 0xe6, 0x99, 0x1c, 0x90, 0x52, 0x48, 0x22, 0x7c, 0x3f, 0x96, 0x46, 0xfb, 0xc4, 0x7e,
	0x90, 0x63, 0x90, 0x44, 0xb1, 0x00, 0xdc, 0x07, 0x7f, 0xda, 0xe8, 0x0f, 0x3b, 0x9b, 0xf7, 0xaf,
	0x76, 0xb9, 0x80, 0xbf, 0xb4, 0x5b, 0x93, 0x13, 0xc5, 0xa2, 0x51, 0x29, 0xe3, 0xa9, 0x04, 0xa5,
	0x25, 0x4b, 0xce, 0x30, 0x1d, 0xc6, 0x7a, 0x68, 0xb4, 0xfa, 0xd4, 0x04, 0x8c, 0x20, 0xe5, 0x92,
	0xf7, 0x87, 0x90, 0x4f, 0x18, 0x31, 0x91, 0x14, 0x09, 0x71, 0x5f, 0x84, 0xa1, 0x38, 0x48, 0x25,
	0x99, 0x83, 0x85, 0x64, 0x9f, 0xa4, 0x08, 0xf3, 0x82, 0x22, 0xfd, 0x50, 0x1c, 0xb8, 0xf5, 0x25,
	0xaf, 0x6c, 0x1e, 0x79, 0x43, 0xb1, 0xb5, 0x92, 0x7c, 0x5e, 0x88, 0x7d, 0xe0, 0xdb, 0x5b, 0xa8,
	0x96, 0x2e, 0x25, 0xc6, 0x3b, 0x92, 0xa6, 0x8d, 0xd1, 0x60, 0x62, 0x12, 0x3a, 0x8f, 0x8a, 0x07,
	0x92, 0x69, 0xc0, 0x67, 0x27, 0xc9, 0xe4, 0xbf, 0x49, 0xf6, 0xb6, 0xdd, 0xb1, 0x1e, 0xe4, 0xbe,
	0x8e, 0x1e, 0xd9, 0xd6, 0x63, 0xdb, 0x7a, 0x62, 0x5b, 0xdf, 0xdb, 0xd6, 0xcf, 0xb6, 0xf5, 0x8f,
	0x6d, 0xfd, 0xe8, 0x58, 0xbd, 0x52, 0xf2, 0xf1, 0x72, 0xfd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x0e, 0xf0, 0xa8, 0xf7, 0x7d, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatementServiceClient is the client API for StatementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatementServiceClient interface {
	// GetStatement retrieves a specific statement, specified by its ID.
	GetStatement(ctx context.Context, in *GetStatementRequest, opts ...grpc.CallOption) (*Statement, error)
	GetStatements(ctx context.Context, in *GetStatementsRequest, opts ...grpc.CallOption) (*GetStatementsResponse, error)
}

type statementServiceClient struct {
	cc *grpc.ClientConn
}

func NewStatementServiceClient(cc *grpc.ClientConn) StatementServiceClient {
	return &statementServiceClient{cc}
}

func (c *statementServiceClient) GetStatement(ctx context.Context, in *GetStatementRequest, opts ...grpc.CallOption) (*Statement, error) {
	out := new(Statement)
	err := c.cc.Invoke(ctx, "/statements.StatementService/GetStatement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statementServiceClient) GetStatements(ctx context.Context, in *GetStatementsRequest, opts ...grpc.CallOption) (*GetStatementsResponse, error) {
	out := new(GetStatementsResponse)
	err := c.cc.Invoke(ctx, "/statements.StatementService/GetStatements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatementServiceServer is the server API for StatementService service.
type StatementServiceServer interface {
	// GetStatement retrieves a specific statement, specified by its ID.
	GetStatement(context.Context, *GetStatementRequest) (*Statement, error)
	GetStatements(context.Context, *GetStatementsRequest) (*GetStatementsResponse, error)
}

func RegisterStatementServiceServer(s *grpc.Server, srv StatementServiceServer) {
	s.RegisterService(&_StatementService_serviceDesc, srv)
}

func _StatementService_GetStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).GetStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statements.StatementService/GetStatement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).GetStatement(ctx, req.(*GetStatementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatementService_GetStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatementServiceServer).GetStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statements.StatementService/GetStatements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatementServiceServer).GetStatements(ctx, req.(*GetStatementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "statements.StatementService",
	HandlerType: (*StatementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatement",
			Handler:    _StatementService_GetStatement_Handler,
		},
		{
			MethodName: "GetStatements",
			Handler:    _StatementService_GetStatements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/statements/all.proto",
}
