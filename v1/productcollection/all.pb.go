// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/productcollection/all.proto

package productcollection

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	product "github.com/openbank/openbank/v1/product"
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

type ProductCollection struct {
	CollectionCode       string             `protobuf:"bytes,1,opt,name=CollectionCode,json=collection_code,proto3" json:"collection_code,omitempty"`
	Products             []*product.Product `protobuf:"bytes,2,rep,name=Products,json=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ProductCollection) Reset()         { *m = ProductCollection{} }
func (m *ProductCollection) String() string { return proto.CompactTextString(m) }
func (*ProductCollection) ProtoMessage()    {}
func (*ProductCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_62bb3483d57949fb, []int{0}
}

func (m *ProductCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductCollection.Unmarshal(m, b)
}
func (m *ProductCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductCollection.Marshal(b, m, deterministic)
}
func (m *ProductCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCollection.Merge(m, src)
}
func (m *ProductCollection) XXX_Size() int {
	return xxx_messageInfo_ProductCollection.Size(m)
}
func (m *ProductCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCollection.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCollection proto.InternalMessageInfo

func (m *ProductCollection) GetCollectionCode() string {
	if m != nil {
		return m.CollectionCode
	}
	return ""
}

func (m *ProductCollection) GetProducts() []*product.Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type CreateProductCollectionRequest struct {
	BankID               string             `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	CollectionCode       string             `protobuf:"bytes,2,opt,name=CollectionCode,json=collection_code,proto3" json:"collection_code,omitempty"`
	ProductCollection    *ProductCollection `protobuf:"bytes,3,opt,name=ProductCollection,json=product_collection,proto3" json:"product_collection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateProductCollectionRequest) Reset()         { *m = CreateProductCollectionRequest{} }
func (m *CreateProductCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProductCollectionRequest) ProtoMessage()    {}
func (*CreateProductCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62bb3483d57949fb, []int{1}
}

func (m *CreateProductCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductCollectionRequest.Unmarshal(m, b)
}
func (m *CreateProductCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductCollectionRequest.Marshal(b, m, deterministic)
}
func (m *CreateProductCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductCollectionRequest.Merge(m, src)
}
func (m *CreateProductCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProductCollectionRequest.Size(m)
}
func (m *CreateProductCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductCollectionRequest proto.InternalMessageInfo

func (m *CreateProductCollectionRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *CreateProductCollectionRequest) GetCollectionCode() string {
	if m != nil {
		return m.CollectionCode
	}
	return ""
}

func (m *CreateProductCollectionRequest) GetProductCollection() *ProductCollection {
	if m != nil {
		return m.ProductCollection
	}
	return nil
}

type UpdateProductCollectionRequest struct {
	BankID               string             `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	CollectionCode       string             `protobuf:"bytes,2,opt,name=CollectionCode,json=collection_code,proto3" json:"collection_code,omitempty"`
	ProductCollection    *ProductCollection `protobuf:"bytes,3,opt,name=ProductCollection,json=product_collection,proto3" json:"product_collection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateProductCollectionRequest) Reset()         { *m = UpdateProductCollectionRequest{} }
func (m *UpdateProductCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProductCollectionRequest) ProtoMessage()    {}
func (*UpdateProductCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62bb3483d57949fb, []int{2}
}

func (m *UpdateProductCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProductCollectionRequest.Unmarshal(m, b)
}
func (m *UpdateProductCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProductCollectionRequest.Marshal(b, m, deterministic)
}
func (m *UpdateProductCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProductCollectionRequest.Merge(m, src)
}
func (m *UpdateProductCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProductCollectionRequest.Size(m)
}
func (m *UpdateProductCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProductCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProductCollectionRequest proto.InternalMessageInfo

func (m *UpdateProductCollectionRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *UpdateProductCollectionRequest) GetCollectionCode() string {
	if m != nil {
		return m.CollectionCode
	}
	return ""
}

func (m *UpdateProductCollectionRequest) GetProductCollection() *ProductCollection {
	if m != nil {
		return m.ProductCollection
	}
	return nil
}

type DeleteProductCollectionRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	CollectionCode       string   `protobuf:"bytes,2,opt,name=CollectionCode,json=collection_code,proto3" json:"collection_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductCollectionRequest) Reset()         { *m = DeleteProductCollectionRequest{} }
func (m *DeleteProductCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProductCollectionRequest) ProtoMessage()    {}
func (*DeleteProductCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62bb3483d57949fb, []int{3}
}

func (m *DeleteProductCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductCollectionRequest.Unmarshal(m, b)
}
func (m *DeleteProductCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductCollectionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProductCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductCollectionRequest.Merge(m, src)
}
func (m *DeleteProductCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProductCollectionRequest.Size(m)
}
func (m *DeleteProductCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductCollectionRequest proto.InternalMessageInfo

func (m *DeleteProductCollectionRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *DeleteProductCollectionRequest) GetCollectionCode() string {
	if m != nil {
		return m.CollectionCode
	}
	return ""
}

type GetProductCollectionRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	CollectionCode       string   `protobuf:"bytes,2,opt,name=CollectionCode,json=collection_code,proto3" json:"collection_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductCollectionRequest) Reset()         { *m = GetProductCollectionRequest{} }
func (m *GetProductCollectionRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductCollectionRequest) ProtoMessage()    {}
func (*GetProductCollectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62bb3483d57949fb, []int{4}
}

func (m *GetProductCollectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductCollectionRequest.Unmarshal(m, b)
}
func (m *GetProductCollectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductCollectionRequest.Marshal(b, m, deterministic)
}
func (m *GetProductCollectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductCollectionRequest.Merge(m, src)
}
func (m *GetProductCollectionRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductCollectionRequest.Size(m)
}
func (m *GetProductCollectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductCollectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductCollectionRequest proto.InternalMessageInfo

func (m *GetProductCollectionRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *GetProductCollectionRequest) GetCollectionCode() string {
	if m != nil {
		return m.CollectionCode
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductCollection)(nil), "productcollection.ProductCollection")
	proto.RegisterType((*CreateProductCollectionRequest)(nil), "productcollection.CreateProductCollectionRequest")
	proto.RegisterType((*UpdateProductCollectionRequest)(nil), "productcollection.UpdateProductCollectionRequest")
	proto.RegisterType((*DeleteProductCollectionRequest)(nil), "productcollection.DeleteProductCollectionRequest")
	proto.RegisterType((*GetProductCollectionRequest)(nil), "productcollection.GetProductCollectionRequest")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/productcollection/all.proto", fileDescriptor_62bb3483d57949fb)
}

var fileDescriptor_62bb3483d57949fb = []byte{
	// 1227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x57, 0xd1, 0x6f, 0x14, 0x45,
	0x18, 0xdf, 0xdd, 0xc3, 0x52, 0x86, 0x04, 0x61, 0x62, 0xa0, 0xb9, 0x22, 0x99, 0x14, 0x93, 0xd6,
	0x4a, 0x67, 0xef, 0x8e, 0x2a, 0x50, 0x31, 0xe6, 0xda, 0x22, 0xb6, 0xc1, 0x78, 0x39, 0x85, 0x18,
	0x12, 0xd2, 0xcc, 0xed, 0x7e, 0x77, 0xb7, 0x74, 0x6f, 0x66, 0x9d, 0x99, 0x6d, 0xa9, 0x84, 0x88,
	0x3c, 0x11, 0x63, 0x08, 0xd6, 0x27, 0xff, 0x05, 0xff, 0x02, 0x9f, 0x7c, 0x92, 0xc4, 0x07, 0x13,
	0x8d, 0xbc, 0x18, 0xe3, 0x9b, 0xfe, 0x11, 0xc6, 0x27, 0x33, 0xbb, 0xdb, 0xbb, 0x6b, 0x76, 0xaf,
	0x94, 0x40, 0x42, 0xe2, 0x53, 0x6f, 0xe7, 0x9b, 0xef, 0xf7, 0x7d, 0x33, 0xbf, 0xdf, 0xf7, 0x7d,
	0x53, 0xf4, 0x76, 0x27, 0xd0, 0xdd, 0xb8, 0x45, 0x3d, 0xd1, 0x73, 0x45, 0x04, 0xbc, 0xc5, 0xf8,
	0xfa, 0xe0, 0xc7, 0x46, 0xd5, 0x8d, 0xa4, 0xf0, 0x63, 0x4f, 0x7b, 0x22, 0x0c, 0xc1, 0xd3, 0x81,
	0xe0, 0x2e, 0x0b, 0x43, 0x1a, 0x49, 0xa1, 0x05, 0x3e, 0x96, 0x33, 0x96, 0x4f, 0x76, 0x84, 0xe8,
	0x84, 0xe0, 0xb2, 0x28, 0x70, 0x19, 0xe7, 0x42, 0x33, 0xb3, 0xac, 0x52, 0x87, 0xf2, 0x99, 0xe4,
	0x8f, 0x37, 0xd7, 0x01, 0x3e, 0xa7, 0x36, 0x59, 0xa7, 0x03, 0xd2, 0x15, 0x51, 0xb2, 0xa3, 0x60,
	0xf7, 0x64, 0x86, 0x95, 0x7c, 0xb5, 0xe2, 0xb6, 0x0b, 0xbd, 0x48, 0x6f, 0x65, 0xc6, 0xea, 0x3e,
	0x13, 0x1f, 0xa4, 0x3b, 0xf5, 0xa5, 0x8d, 0x8e, 0x35, 0xd2, 0xd5, 0xa5, 0x7e, 0xc6, 0xf8, 0x2c,
	0x3a, 0x32, 0xf8, 0x5a, 0x12, 0x3e, 0x4c, 0xd8, 0xc4, 0x9e, 0x39, 0xb4, 0x88, 0xc6, 0xad, 0x09,
	0x6b, 0xc6, 0xaa, 0x58, 0x0d, 0xab, 0xf9, 0xf2, 0xe0, 0x84, 0x6b, 0x9e, 0xf0, 0x01, 0x9f, 0x47,
	0xe3, 0x19, 0x92, 0x9a, 0x70, 0x48, 0x69, 0xe6, 0x70, 0xed, 0x28, 0xcd, 0x02, 0xd2, 0xcc, 0xb0,
	0x0b, 0x60, 0x3c, 0x33, 0xaa, 0x85, 0xb1, 0x71, 0xeb, 0xa8, 0x35, 0x61, 0x4d, 0xfd, 0x65, 0xa3,
	0x53, 0x4b, 0x12, 0x98, 0x86, 0x5c, 0x4a, 0x4d, 0xf8, 0x34, 0x06, 0xa5, 0xf1, 0x69, 0x34, 0xb6,
	0xc8, 0xf8, 0xfa, 0xca, 0x72, 0x41, 0x46, 0x07, 0xcd, 0x31, 0xd7, 0x02, 0xbf, 0x20, 0x7d, 0xe7,
	0xc9, 0xe9, 0xaf, 0x15, 0x5c, 0xc4, 0x44, 0x89, 0xd8, 0x33, 0x87, 0x6b, 0xaf, 0xd1, 0x1c, 0xa9,
	0x34, 0xb7, 0x77, 0x17, 0x3a, 0xce, 0x1c, 0xd6, 0x06, 0x1e, 0xbb, 0x4e, 0x79, 0x35, 0xf2, 0xff,
	0xe7, 0xa7, 0xbc, 0x67, 0xa3, 0x53, 0xcb, 0x10, 0xc2, 0x8b, 0x38, 0x65, 0x3f, 0x89, 0xcf, 0xd1,
	0xe4, 0x65, 0xd0, 0x2f, 0x2e, 0x81, 0xda, 0x0f, 0x47, 0xd0, 0x44, 0x2e, 0xfc, 0x47, 0x20, 0x37,
	0x02, 0x0f, 0xf0, 0xe3, 0x12, 0x3a, 0x31, 0x42, 0xee, 0xb8, 0x5a, 0x40, 0xc6, 0xde, 0xa5, 0x51,
	0xde, 0x17, 0x7f, 0x53, 0x7f, 0x38, 0xdb, 0xf5, 0x07, 0x0e, 0xc2, 0x99, 0x85, 0x0c, 0xc5, 0x9c,
	0x4c, 0x03, 0x10, 0x46, 0x32, 0x20, 0x32, 0xd4, 0xc4, 0xce, 0xa7, 0x46, 0x45, 0x18, 0xe1, 0xb0,
	0x59, 0xb0, 0x83, 0x30, 0xee, 0x13, 0x09, 0x3a, 0x96, 0x5c, 0x11, 0xdd, 0x05, 0x22, 0x5a, 0x37,
	0xc1, 0xd3, 0x74, 0xb5, 0x8d, 0x4a, 0xb5, 0x4a, 0x15, 0xaf, 0xa1, 0xe9, 0x5c, 0x36, 0xc4, 0x4b,
	0x40, 0x7d, 0xa2, 0x62, 0xcf, 0x03, 0xa5, 0xda, 0x71, 0x18, 0x6e, 0x51, 0x3c, 0x8f, 0x6a, 0xe5,
	0xca, 0x69, 0xd7, 0x87, 0x76, 0xc0, 0x83, 0xb4, 0x13, 0xe6, 0x8e, 0x96, 0xc3, 0x6a, 0x5d, 0x40,
	0xe7, 0xd0, 0xd8, 0x87, 0xf5, 0x58, 0x77, 0x6b, 0x78, 0x0e, 0xbd, 0xd1, 0xd5, 0x3a, 0x52, 0x0b,
	0xae, 0xcb, 0x62, 0xdd, 0xa5, 0x2d, 0xbe, 0x4e, 0xb5, 0xc8, 0xa3, 0xd0, 0x4d, 0x19, 0x68, 0xb8,
	0xf7, 0xf8, 0xef, 0x6f, 0x9c, 0x8b, 0x0b, 0xf6, 0xec, 0xd4, 0x39, 0xd3, 0x25, 0x0d, 0xf3, 0xca,
	0xbd, 0x9d, 0x4a, 0xe3, 0xce, 0x8e, 0xd3, 0xdc, 0xc0, 0x4b, 0xb9, 0xb7, 0x77, 0x4b, 0xe2, 0xce,
	0x7d, 0xc7, 0x7a, 0xe8, 0x24, 0x6a, 0xc0, 0x5f, 0x1c, 0x40, 0xaf, 0x14, 0x89, 0x0e, 0xd3, 0x02,
	0x7e, 0xf6, 0x50, 0xe7, 0x3e, 0xf9, 0x7c, 0x50, 0xda, 0xae, 0x3f, 0x2a, 0xe6, 0x73, 0xba, 0x09,
	0x5a, 0x06, 0xb0, 0x01, 0x45, 0x6c, 0x05, 0xbc, 0x2d, 0x64, 0x2f, 0x99, 0x2d, 0xe5, 0x85, 0xfe,
	0xc6, 0xa1, 0x55, 0xc2, 0x5a, 0x22, 0xd6, 0x09, 0x93, 0x3b, 0xee, 0x2a, 0x02, 0x2f, 0x68, 0x07,
	0xe0, 0x93, 0xd6, 0x56, 0x62, 0x58, 0x59, 0x5e, 0xbd, 0x61, 0xd8, 0xad, 0xe0, 0x6b, 0xe8, 0x54,
	0x96, 0x37, 0x81, 0x5b, 0xe0, 0xc5, 0xcf, 0x89, 0xd4, 0xd5, 0x2a, 0x2a, 0xcd, 0x57, 0xe6, 0xf1,
	0x2c, 0x9a, 0x69, 0x26, 0xd2, 0x02, 0x9f, 0x6c, 0x76, 0x81, 0x27, 0xd1, 0x25, 0x28, 0x11, 0x4b,
	0x0f, 0x48, 0xa0, 0x08, 0x17, 0x9a, 0xb4, 0x45, 0xcc, 0x7d, 0xda, 0x3a, 0x8f, 0xde, 0xea, 0xeb,
	0xe0, 0x0c, 0x9a, 0xdd, 0x9f, 0x0e, 0x24, 0x30, 0x3f, 0x91, 0xc1, 0x05, 0xfc, 0x1c, 0x34, 0xf0,
	0xa8, 0x84, 0x4e, 0x8c, 0x68, 0xf1, 0x85, 0x95, 0xbd, 0xf7, 0x38, 0xd8, 0xa7, 0x12, 0x7e, 0x76,
	0xb6, 0xeb, 0x77, 0x47, 0x54, 0x76, 0x1a, 0xa0, 0xb8, 0xb2, 0x67, 0x53, 0xa3, 0x2a, 0xb4, 0x4e,
	0xab, 0x61, 0x4d, 0xa4, 0xb5, 0x3c, 0x5f, 0x5c, 0xcb, 0xc3, 0x74, 0x93, 0x38, 0xc1, 0xf4, 0x5f,
	0x60, 0x2d, 0x97, 0x9f, 0x03, 0x8f, 0xff, 0x3a, 0xe8, 0xc4, 0x88, 0x21, 0x56, 0xc8, 0xe3, 0xde,
	0x03, 0xaf, 0x7c, 0x9c, 0xa6, 0xaf, 0x37, 0xba, 0xf3, 0x7a, 0xa3, 0x97, 0xcc, 0xeb, 0x6d, 0xea,
	0x6b, 0x67, 0xbb, 0xfe, 0x93, 0x5d, 0xcc, 0x5c, 0x0a, 0x59, 0xcc, 0xdc, 0x4c, 0x03, 0x64, 0x8f,
	0x71, 0xe0, 0x3a, 0xdc, 0x22, 0xfe, 0xe8, 0x8d, 0x74, 0xb5, 0x92, 0xf2, 0xf6, 0xfa, 0x13, 0x79,
	0x4b, 0x51, 0x7c, 0xfa, 0xac, 0x0c, 0x5c, 0x98, 0x7d, 0xd6, 0xeb, 0x2f, 0x97, 0xee, 0x3b, 0xd6,
	0xe2, 0x77, 0x07, 0xb7, 0xeb, 0x7f, 0x8e, 0xa1, 0x52, 0x8d, 0x56, 0xf0, 0x0d, 0x74, 0x3c, 0x7f,
	0x43, 0xa4, 0xde, 0x58, 0xc1, 0x17, 0x1b, 0x52, 0x6c, 0x04, 0x3e, 0x28, 0xb2, 0xd4, 0xbc, 0xba,
	0x4c, 0x44, 0x04, 0x32, 0x7d, 0x35, 0x13, 0x91, 0xb6, 0x8a, 0x02, 0xbf, 0x9d, 0xee, 0x41, 0x6b,
	0x2f, 0x55, 0x69, 0x85, 0x56, 0x66, 0x6d, 0xa7, 0x76, 0x94, 0x45, 0x51, 0x18, 0x78, 0x89, 0xaf,
	0x7b, 0x53, 0x09, 0xbe, 0x90, 0x5b, 0x69, 0xae, 0x99, 0x8e, 0x54, 0xc1, 0x9f, 0xa0, 0x6b, 0x45,
	0x1d, 0x29, 0x6d, 0x81, 0x2d, 0xe1, 0x6f, 0x99, 0xae, 0xd4, 0x63, 0x61, 0x5a, 0x3b, 0xa6, 0x1d,
	0x0a, 0x49, 0x7c, 0x01, 0x69, 0xab, 0xea, 0x31, 0xed, 0x75, 0x13, 0x17, 0xb8, 0x15, 0x81, 0x67,
	0xcc, 0x99, 0x2f, 0x6d, 0x5e, 0x31, 0x01, 0xaa, 0xf8, 0x12, 0x5a, 0x1a, 0x1d, 0xa0, 0x0f, 0xe4,
	0x09, 0xae, 0x59, 0x90, 0x4d, 0xdc, 0x58, 0x81, 0x9c, 0x56, 0x66, 0xa8, 0xfa, 0xc0, 0x75, 0xc0,
	0x42, 0x45, 0x9b, 0x0d, 0x83, 0x76, 0x16, 0xaf, 0xa0, 0xcb, 0x79, 0x34, 0xb3, 0x7f, 0x00, 0xd5,
	0x65, 0x66, 0x52, 0x80, 0xec, 0x05, 0x4a, 0x99, 0x0b, 0xd2, 0x82, 0xb0, 0x44, 0x19, 0xbb, 0x9a,
	0x2d, 0x6d, 0x3e, 0x7d, 0x4b, 0x6e, 0xbe, 0x87, 0x4a, 0x6f, 0x56, 0x2a, 0xf8, 0x5d, 0xf4, 0xce,
	0x6e, 0x17, 0xc6, 0x49, 0xcc, 0xfb, 0x37, 0x00, 0x52, 0x0a, 0x49, 0x84, 0xe7, 0xc5, 0xd2, 0x5c,
	0x57, 0x8a, 0xa8, 0x40, 0x6e, 0x80, 0x24, 0x2a, 0xf0, 0x81, 0x5e, 0xff, 0xcd, 0x41, 0xbf, 0x38,
	0x7d, 0x59, 0xfe, 0xe8, 0x8c, 0x97, 0xf0, 0x57, 0x76, 0x3d, 0x4b, 0x52, 0x14, 0x8d, 0xb9, 0x9d,
	0x84, 0x94, 0xc9, 0x48, 0x82, 0xd2, 0x32, 0x48, 0x62, 0x99, 0xc3, 0xc5, 0xba, 0x6b, 0xae, 0xc9,
	0x4b, 0x5e, 0x20, 0xe6, 0x2e, 0x14, 0x25, 0x1f, 0x77, 0x61, 0xd8, 0x60, 0x20, 0x22, 0x29, 0x92,
	0x00, 0x6d, 0x11, 0x86, 0x62, 0x33, 0xbd, 0x0d, 0x93, 0x80, 0x90, 0xc1, 0x67, 0xe9, 0x0e, 0xa3,
	0x5b, 0xd2, 0x0e, 0xc5, 0x26, 0x9d, 0x39, 0x50, 0x1b, 0x37, 0xba, 0x37, 0x10, 0x0b, 0x87, 0xcc,
	0x2f, 0x2d, 0xd6, 0x81, 0x2f, 0x7e, 0x6b, 0xa3, 0x95, 0xa7, 0x19, 0x45, 0x78, 0xf2, 0x5a, 0x30,
	0x78, 0x65, 0x0d, 0x9d, 0xc7, 0x67, 0x9a, 0xa1, 0x2b, 0x4f, 0x55, 0x8f, 0xf8, 0xd5, 0x0f, 0x18,
	0x67, 0x1d, 0x18, 0x85, 0x76, 0x52, 0x8b, 0x04, 0x81, 0x45, 0x41, 0xbe, 0xa5, 0xbd, 0x6f, 0x37,
	0xec, 0xeb, 0xf9, 0xff, 0x69, 0xef, 0xda, 0xd6, 0x7d, 0xdb, 0x7a, 0x68, 0x5b, 0xdf, 0xdb, 0xd6,
	0xef, 0xb6, 0xf5, 0x8f, 0x6d, 0xfd, 0xea, 0x58, 0xad, 0xb1, 0xa4, 0xaf, 0x9d, 0xfd, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0xdf, 0xe8, 0x0b, 0xa9, 0x41, 0x0f, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductCollectionServiceClient is the client API for ProductCollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductCollectionServiceClient interface {
	// CreateProductCollection creates a new product collection.
	CreateProductCollection(ctx context.Context, in *CreateProductCollectionRequest, opts ...grpc.CallOption) (*ProductCollection, error)
	// GetProductCollection retrieves the details of a specific product information, selected by its ID.
	GetProductCollection(ctx context.Context, in *GetProductCollectionRequest, opts ...grpc.CallOption) (*ProductCollection, error)
	// UpdateProductCollection updates the product collection.
	UpdateProductCollection(ctx context.Context, in *UpdateProductCollectionRequest, opts ...grpc.CallOption) (*ProductCollection, error)
	// DeleteProductCollection deletes a product collection.
	DeleteProductCollection(ctx context.Context, in *DeleteProductCollectionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type productCollectionServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductCollectionServiceClient(cc *grpc.ClientConn) ProductCollectionServiceClient {
	return &productCollectionServiceClient{cc}
}

func (c *productCollectionServiceClient) CreateProductCollection(ctx context.Context, in *CreateProductCollectionRequest, opts ...grpc.CallOption) (*ProductCollection, error) {
	out := new(ProductCollection)
	err := c.cc.Invoke(ctx, "/productcollection.ProductCollectionService/CreateProductCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCollectionServiceClient) GetProductCollection(ctx context.Context, in *GetProductCollectionRequest, opts ...grpc.CallOption) (*ProductCollection, error) {
	out := new(ProductCollection)
	err := c.cc.Invoke(ctx, "/productcollection.ProductCollectionService/GetProductCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCollectionServiceClient) UpdateProductCollection(ctx context.Context, in *UpdateProductCollectionRequest, opts ...grpc.CallOption) (*ProductCollection, error) {
	out := new(ProductCollection)
	err := c.cc.Invoke(ctx, "/productcollection.ProductCollectionService/UpdateProductCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCollectionServiceClient) DeleteProductCollection(ctx context.Context, in *DeleteProductCollectionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/productcollection.ProductCollectionService/DeleteProductCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCollectionServiceServer is the server API for ProductCollectionService service.
type ProductCollectionServiceServer interface {
	// CreateProductCollection creates a new product collection.
	CreateProductCollection(context.Context, *CreateProductCollectionRequest) (*ProductCollection, error)
	// GetProductCollection retrieves the details of a specific product information, selected by its ID.
	GetProductCollection(context.Context, *GetProductCollectionRequest) (*ProductCollection, error)
	// UpdateProductCollection updates the product collection.
	UpdateProductCollection(context.Context, *UpdateProductCollectionRequest) (*ProductCollection, error)
	// DeleteProductCollection deletes a product collection.
	DeleteProductCollection(context.Context, *DeleteProductCollectionRequest) (*empty.Empty, error)
}

// UnimplementedProductCollectionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductCollectionServiceServer struct {
}

func (*UnimplementedProductCollectionServiceServer) CreateProductCollection(ctx context.Context, req *CreateProductCollectionRequest) (*ProductCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductCollection not implemented")
}
func (*UnimplementedProductCollectionServiceServer) GetProductCollection(ctx context.Context, req *GetProductCollectionRequest) (*ProductCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCollection not implemented")
}
func (*UnimplementedProductCollectionServiceServer) UpdateProductCollection(ctx context.Context, req *UpdateProductCollectionRequest) (*ProductCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductCollection not implemented")
}
func (*UnimplementedProductCollectionServiceServer) DeleteProductCollection(ctx context.Context, req *DeleteProductCollectionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductCollection not implemented")
}

func RegisterProductCollectionServiceServer(s *grpc.Server, srv ProductCollectionServiceServer) {
	s.RegisterService(&_ProductCollectionService_serviceDesc, srv)
}

func _ProductCollectionService_CreateProductCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCollectionServiceServer).CreateProductCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productcollection.ProductCollectionService/CreateProductCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCollectionServiceServer).CreateProductCollection(ctx, req.(*CreateProductCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCollectionService_GetProductCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCollectionServiceServer).GetProductCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productcollection.ProductCollectionService/GetProductCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCollectionServiceServer).GetProductCollection(ctx, req.(*GetProductCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCollectionService_UpdateProductCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCollectionServiceServer).UpdateProductCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productcollection.ProductCollectionService/UpdateProductCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCollectionServiceServer).UpdateProductCollection(ctx, req.(*UpdateProductCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCollectionService_DeleteProductCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCollectionServiceServer).DeleteProductCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productcollection.ProductCollectionService/DeleteProductCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCollectionServiceServer).DeleteProductCollection(ctx, req.(*DeleteProductCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductCollectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "productcollection.ProductCollectionService",
	HandlerType: (*ProductCollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProductCollection",
			Handler:    _ProductCollectionService_CreateProductCollection_Handler,
		},
		{
			MethodName: "GetProductCollection",
			Handler:    _ProductCollectionService_GetProductCollection_Handler,
		},
		{
			MethodName: "UpdateProductCollection",
			Handler:    _ProductCollectionService_UpdateProductCollection_Handler,
		},
		{
			MethodName: "DeleteProductCollection",
			Handler:    _ProductCollectionService_DeleteProductCollection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/productcollection/all.proto",
}
