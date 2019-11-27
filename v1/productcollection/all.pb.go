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
	CollectionCode       string             `protobuf:"bytes,1,opt,name=CollectionCode,json=collection_code,proto3" json:"CollectionCode,omitempty"`
	Products             []*product.Product `protobuf:"bytes,2,rep,name=Products,json=products,proto3" json:"Products,omitempty"`
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
	BankID               string             `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	CollectionCode       string             `protobuf:"bytes,2,opt,name=CollectionCode,json=collection_code,proto3" json:"CollectionCode,omitempty"`
	ProductCollection    *ProductCollection `protobuf:"bytes,3,opt,name=ProductCollection,json=product_collection,proto3" json:"ProductCollection,omitempty"`
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
	BankID               string             `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	CollectionCode       string             `protobuf:"bytes,2,opt,name=CollectionCode,json=collection_code,proto3" json:"CollectionCode,omitempty"`
	ProductCollection    *ProductCollection `protobuf:"bytes,3,opt,name=ProductCollection,json=product_collection,proto3" json:"ProductCollection,omitempty"`
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
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	CollectionCode       string   `protobuf:"bytes,2,opt,name=CollectionCode,json=collection_code,proto3" json:"CollectionCode,omitempty"`
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
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	CollectionCode       string   `protobuf:"bytes,2,opt,name=CollectionCode,json=collection_code,proto3" json:"CollectionCode,omitempty"`
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
	// 1190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x57, 0xcd, 0x8b, 0x1c, 0x45,
	0x14, 0xef, 0xee, 0x49, 0x36, 0x9b, 0x0a, 0xea, 0xa6, 0x22, 0xc9, 0x30, 0x09, 0xa1, 0xd8, 0x08,
	0x3b, 0x2e, 0xa6, 0xa6, 0x67, 0xb2, 0x62, 0x32, 0x46, 0x64, 0x76, 0x37, 0xc6, 0x5d, 0x04, 0x87,
	0xd1, 0x04, 0x09, 0x84, 0xa1, 0xa7, 0xfb, 0xcd, 0x4c, 0x27, 0x3d, 0x55, 0x6d, 0x55, 0xf5, 0x4e,
	0xd6, 0x10, 0x94, 0x9c, 0x82, 0x78, 0x08, 0xeb, 0x1f, 0x10, 0xf0, 0x90, 0x93, 0xde, 0x3c, 0xf8,
	0x67, 0x08, 0x7a, 0x10, 0x3d, 0x9a, 0x43, 0xc0, 0x83, 0xe0, 0xc9, 0x83, 0x07, 0xa9, 0xee, 0xde,
	0xf9, 0xa0, 0x7b, 0x36, 0x2b, 0xae, 0x04, 0x3c, 0xed, 0x74, 0xbf, 0xf7, 0x7e, 0xef, 0xe3, 0xf7,
	0x3e, 0x7a, 0xd1, 0x9b, 0x3d, 0x5f, 0xf5, 0xa3, 0x0e, 0x75, 0xf9, 0xa0, 0xc2, 0x43, 0x60, 0x1d,
	0x87, 0xdd, 0x1e, 0xff, 0xd8, 0xaa, 0x56, 0x42, 0xc1, 0xbd, 0xc8, 0x55, 0x2e, 0x0f, 0x02, 0x70,
	0x95, 0xcf, 0x59, 0xc5, 0x09, 0x02, 0x1a, 0x0a, 0xae, 0x38, 0x3e, 0x9e, 0x11, 0x96, 0xce, 0xf4,
	0x38, 0xef, 0x05, 0x50, 0x71, 0x42, 0xbf, 0xe2, 0x30, 0xc6, 0x95, 0xa3, 0x5f, 0xcb, 0xc4, 0xa0,
	0xf4, 0x5a, 0xfc, 0xc7, 0x3d, 0xdf, 0x03, 0x76, 0x5e, 0x0e, 0x9d, 0x5e, 0x0f, 0x44, 0x85, 0x87,
	0xb1, 0x46, 0x8e, 0xf6, 0xe9, 0x14, 0x2b, 0x7e, 0xea, 0x44, 0xdd, 0x0a, 0x0c, 0x42, 0xb5, 0x9d,
	0x0a, 0xab, 0xfb, 0x0c, 0x7c, 0x1c, 0xee, 0xe2, 0xe7, 0x26, 0x3a, 0xde, 0x4c, 0xde, 0xae, 0x8d,
	0x22, 0xc6, 0x17, 0xd0, 0x8b, 0xe3, 0xa7, 0x35, 0xee, 0x41, 0xd1, 0x24, 0x66, 0xf9, 0xe8, 0x2a,
	0x9a, 0x37, 0x8a, 0x46, 0xd9, 0xb0, 0x8d, 0xa6, 0xd1, 0x7a, 0x69, 0x9c, 0x61, 0xdb, 0xe5, 0x1e,
	0xe0, 0x8b, 0x68, 0x3e, 0x45, 0x92, 0x45, 0x8b, 0x14, 0xca, 0xc7, 0x6a, 0x0b, 0x34, 0x75, 0x48,
	0x53, 0xc1, 0x14, 0xc0, 0x7c, 0x2a, 0x94, 0xf5, 0xb9, 0x79, 0x63, 0xc1, 0x28, 0x1a, 0x8b, 0xbf,
	0x9a, 0xe8, 0xec, 0x9a, 0x00, 0x47, 0x41, 0x26, 0xa4, 0x16, 0x7c, 0x1c, 0x81, 0x54, 0xf8, 0x1c,
	0x9a, 0x5b, 0x75, 0xd8, 0xed, 0x8d, 0xf5, 0x9c, 0x88, 0x8e, 0xe8, 0x34, 0xdb, 0xbe, 0x97, 0x13,
	0xbe, 0xf5, 0xec, 0xf0, 0xdb, 0x39, 0x85, 0x28, 0x16, 0x88, 0x59, 0x3e, 0x56, 0x7b, 0x85, 0x66,
	0x48, 0xa5, 0x19, 0xdd, 0x29, 0x74, 0x9c, 0x1a, 0xb4, 0xc7, 0x16, 0x53, 0x59, 0x5e, 0x0b, 0xbd,
	0xff, 0x79, 0x96, 0xf7, 0x4d, 0x74, 0x76, 0x1d, 0x02, 0x78, 0x1e, 0x59, 0x8e, 0x82, 0xf8, 0x14,
	0x9d, 0xbe, 0x0a, 0xea, 0xf9, 0x05, 0x50, 0x7b, 0xfc, 0x02, 0x2a, 0x66, 0xdc, 0x7f, 0x00, 0x62,
	0xcb, 0x77, 0x01, 0x7f, 0x53, 0x40, 0xa7, 0x66, 0xb4, 0x3b, 0xae, 0xe6, 0x90, 0xb1, 0xf7, 0x68,
	0x94, 0xf6, 0xc5, 0xdf, 0xe2, 0xb7, 0xd6, 0x4e, 0xe3, 0xa9, 0x89, 0x70, 0x2a, 0x21, 0x13, 0x3e,
	0x4f, 0x27, 0x0e, 0x88, 0x43, 0x52, 0x20, 0x32, 0xb1, 0xc4, 0x2e, 0x26, 0x42, 0x49, 0x1c, 0xc2,
	0x60, 0x98, 0xa3, 0x41, 0x1c, 0xe6, 0x11, 0x01, 0x2a, 0x12, 0x4c, 0x12, 0xd5, 0x07, 0xc2, 0x3b,
	0xb7, 0xc0, 0x55, 0x74, 0xb3, 0x8b, 0x0a, 0x35, 0xbb, 0x8a, 0xdb, 0x68, 0x29, 0x13, 0x0d, 0x71,
	0x63, 0x50, 0x8f, 0xc8, 0xc8, 0x75, 0x41, 0xca, 0x6e, 0x14, 0x04, 0xdb, 0x14, 0xaf, 0xa0, 0x5a,
	0xc9, 0x3e, 0x57, 0xf1, 0xa0, 0xeb, 0x33, 0x3f, 0xd9, 0x84, 0x99, 0xd4, 0x32, 0x58, 0x9d, 0x13,
	0xe8, 0x38, 0x9a, 0x7b, 0xbf, 0x11, 0xa9, 0x7e, 0x0d, 0x1f, 0x41, 0x87, 0x87, 0xc2, 0x57, 0x70,
	0xff, 0x87, 0x27, 0x5f, 0x5a, 0x97, 0xeb, 0xe6, 0xf2, 0xe2, 0x1b, 0x7a, 0xff, 0x69, 0x4e, 0x65,
	0xe5, 0x6e, 0x42, 0xfa, 0xbd, 0x5d, 0xd0, 0xf3, 0x63, 0x54, 0x59, 0xb9, 0x3b, 0x4d, 0xf6, 0xbd,
	0x07, 0x96, 0xf1, 0xd0, 0x8a, 0x79, 0xc6, 0xbf, 0x14, 0xd0, 0xcb, 0x79, 0xed, 0x84, 0x69, 0x4e,
	0xe5, 0xf7, 0xe8, 0xbb, 0x7d, 0x32, 0xf5, 0xd4, 0xda, 0x69, 0x3c, 0xb2, 0x72, 0x99, 0x5a, 0x6a,
	0x81, 0x12, 0x3e, 0x6c, 0x41, 0x1e, 0x0f, 0x3e, 0xeb, 0x72, 0x31, 0x88, 0xaf, 0x46, 0xa9, 0x3e,
	0x52, 0x9c, 0x78, 0x4b, 0x9c, 0x0e, 0x8f, 0x54, 0xcc, 0xd1, 0xae, 0xb9, 0x0c, 0xc1, 0xf5, 0xbb,
	0x3e, 0x78, 0xa4, 0xb3, 0x1d, 0x0b, 0x36, 0xd6, 0x37, 0x6f, 0x6a, 0xde, 0x6c, 0x7c, 0x1d, 0x9d,
	0x4d, 0xe3, 0x26, 0x70, 0x07, 0xdc, 0xe8, 0x80, 0xe8, 0xda, 0xac, 0xa2, 0xc2, 0x8a, 0xbd, 0x82,
	0x97, 0x51, 0xb9, 0x15, 0x37, 0x0d, 0x78, 0x64, 0xd8, 0x07, 0x16, 0x7b, 0x17, 0x20, 0x79, 0x24,
	0x5c, 0x20, 0xbe, 0x24, 0x8c, 0x2b, 0xd2, 0xe5, 0x11, 0xf3, 0x68, 0x07, 0xa3, 0x85, 0x11, 0xc3,
	0x73, 0xe8, 0x90, 0x00, 0xc7, 0x8b, 0x09, 0xbe, 0x84, 0x0f, 0x80, 0xdd, 0x47, 0x05, 0x74, 0x6a,
	0xc6, 0x5a, 0xce, 0x9d, 0xc6, 0xbd, 0x57, 0xf8, 0x3e, 0x39, 0x7e, 0x6c, 0xed, 0x34, 0x7e, 0x9e,
	0x31, 0x8d, 0x89, 0x83, 0xfc, 0x69, 0x5c, 0x4e, 0x84, 0x32, 0x57, 0xba, 0x24, 0x27, 0xd9, 0x4e,
	0xe6, 0x6f, 0x25, 0x7f, 0xfe, 0x26, 0x89, 0x24, 0x51, 0x8c, 0xe9, 0xfd, 0x27, 0xf3, 0x57, 0x3a,
	0x00, 0x86, 0x7e, 0xb4, 0xd0, 0xa9, 0x19, 0x27, 0x25, 0x97, 0xa1, 0xbd, 0xcf, 0x4f, 0xe9, 0x24,
	0x4d, 0xbe, 0xa5, 0xe8, 0xee, 0xb7, 0x14, 0xbd, 0xa2, 0xbf, 0xa5, 0x16, 0x7f, 0x33, 0x77, 0x1a,
	0x5f, 0xcd, 0xe0, 0x24, 0x81, 0xcc, 0xe7, 0xa4, 0xdc, 0x04, 0x31, 0x70, 0x18, 0x30, 0x15, 0x6c,
	0x13, 0x6f, 0xb6, 0x22, 0xdd, 0xb4, 0x13, 0x46, 0x5e, 0x7d, 0x26, 0x23, 0x09, 0x8a, 0x47, 0x67,
	0xd7, 0xf6, 0xd2, 0xf2, 0xbf, 0x2d, 0x6c, 0xa9, 0xf0, 0xc0, 0x32, 0x56, 0xbf, 0x9e, 0xdb, 0x69,
	0xfc, 0x75, 0x18, 0x15, 0x6a, 0xd4, 0xc6, 0x37, 0xd1, 0xc9, 0x6c, 0xee, 0xa4, 0xd1, 0xdc, 0xc0,
	0x97, 0x9b, 0x82, 0x6f, 0xf9, 0x1e, 0x48, 0xb2, 0xd6, 0xba, 0xb6, 0x4e, 0x78, 0x08, 0x22, 0xf9,
	0x3a, 0x25, 0x3c, 0x19, 0xdc, 0x1c, 0xbb, 0xdd, 0x59, 0xa6, 0xb5, 0xc3, 0x55, 0x6a, 0x53, 0x7b,
	0xd9, 0xb4, 0x6a, 0x0b, 0x4e, 0x18, 0x06, 0xbe, 0x1b, 0xdb, 0x56, 0x6e, 0x49, 0xce, 0xea, 0x99,
	0x37, 0xad, 0xb6, 0xde, 0x0f, 0x36, 0xfe, 0x08, 0x5d, 0xcf, 0xdb, 0x0f, 0xc9, 0x42, 0xea, 0x70,
	0x6f, 0x5b, 0xef, 0x88, 0x81, 0x13, 0x24, 0xfd, 0xae, 0x97, 0x13, 0x17, 0xc4, 0xe3, 0x90, 0x2c,
	0x8e, 0x81, 0xa3, 0xdc, 0x7e, 0x6c, 0x02, 0x77, 0x42, 0x70, 0xb5, 0x38, 0xb5, 0xa5, 0xad, 0xf7,
	0xb4, 0x83, 0x2a, 0xbe, 0x82, 0xd6, 0x66, 0x3b, 0x18, 0x01, 0xb9, 0x9c, 0x29, 0xc7, 0x4f, 0x2f,
	0x5b, 0x24, 0x41, 0x2c, 0x49, 0x7d, 0xbc, 0x3c, 0x60, 0xca, 0x77, 0x02, 0x49, 0x5b, 0x4d, 0x8d,
	0x76, 0x01, 0x6f, 0xa0, 0xab, 0x59, 0x34, 0xad, 0x3f, 0x86, 0xea, 0x3b, 0x7a, 0x6f, 0x83, 0x18,
	0xf8, 0x52, 0xea, 0x02, 0x29, 0x4e, 0x9c, 0x98, 0xf3, 0xa9, 0xd5, 0x47, 0x5b, 0xff, 0x7c, 0x41,
	0xb6, 0xde, 0x41, 0x85, 0xd7, 0x6d, 0x1b, 0xbf, 0x8d, 0xde, 0x9a, 0x36, 0x71, 0x18, 0x89, 0xd8,
	0xa8, 0x02, 0x20, 0x04, 0x17, 0x84, 0xbb, 0x6e, 0x24, 0x74, 0xb9, 0x12, 0x44, 0x09, 0x62, 0x0b,
	0x04, 0x91, 0xbe, 0x07, 0xf4, 0xc6, 0x1f, 0x26, 0xfa, 0xdd, 0x1c, 0x35, 0xdc, 0x13, 0x73, 0xbe,
	0x80, 0xbf, 0x30, 0x1b, 0x69, 0x90, 0x3c, 0xef, 0xe8, 0xec, 0x06, 0x24, 0x75, 0x44, 0x02, 0xa4,
	0x12, 0x7e, 0xec, 0x4b, 0x27, 0x17, 0xa9, 0xbe, 0x2e, 0x93, 0x1b, 0x5f, 0x7a, 0x5d, 0x0b, 0x49,
	0xc9, 0x87, 0x7d, 0x98, 0x14, 0x68, 0x88, 0x50, 0xf0, 0xd8, 0x41, 0x97, 0x07, 0x01, 0x1f, 0x26,
	0xd5, 0xd0, 0x01, 0x70, 0xe1, 0x7f, 0x92, 0x68, 0xe8, 0xbe, 0x25, 0xdd, 0x80, 0x0f, 0x69, 0xf9,
	0x50, 0x6d, 0x5e, 0xf7, 0xbd, 0x86, 0xa8, 0x1f, 0xd5, 0xbf, 0x14, 0xbf, 0x0d, 0x6c, 0xb5, 0x8e,
	0xce, 0xa4, 0xd3, 0x81, 0x4f, 0x5c, 0x15, 0x0e, 0x53, 0x92, 0xc4, 0x4f, 0x69, 0x85, 0x51, 0x29,
	0xb9, 0x1a, 0x18, 0xa7, 0x42, 0xfd, 0x90, 0xca, 0xde, 0x35, 0x9b, 0xc6, 0x8d, 0xec, 0x7f, 0x6f,
	0x9f, 0x99, 0xc6, 0x03, 0xd3, 0x78, 0x68, 0x1a, 0xdf, 0x99, 0xc6, 0x4f, 0xa6, 0xf1, 0xa7, 0x69,
	0x7c, 0x6f, 0x19, 0x9d, 0xb9, 0x78, 0x67, 0x5c, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x30, 0x9c,
	0xfa, 0x18, 0x2b, 0x0e, 0x00, 0x00,
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
