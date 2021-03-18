// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package accountapplication

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccountApplicationServiceClient is the client API for AccountApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountApplicationServiceClient interface {
	// CreateAccountApplication creates a new account application
	CreateAccountApplication(ctx context.Context, in *CreateAccountApplicationRequest, opts ...grpc.CallOption) (*CreateAccountApplicationResponse, error)
	// GetAccountApplication retrieves the details of an account application, selected by its id.
	GetAccountApplication(ctx context.Context, in *GetAccountApplicationRequest, opts ...grpc.CallOption) (*AccountApplication, error)
	// GetAccountApplications returns a list containing up to 20 accounts.
	GetAccountApplications(ctx context.Context, in *GetAccountApplicationsRequest, opts ...grpc.CallOption) (*GetAccountApplicationsResponse, error)
	// UpdateAccountApplication status updates the status.
	UpdateAccountApplicationStatus(ctx context.Context, in *UpdateAccountApplicationStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type accountApplicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountApplicationServiceClient(cc grpc.ClientConnInterface) AccountApplicationServiceClient {
	return &accountApplicationServiceClient{cc}
}

func (c *accountApplicationServiceClient) CreateAccountApplication(ctx context.Context, in *CreateAccountApplicationRequest, opts ...grpc.CallOption) (*CreateAccountApplicationResponse, error) {
	out := new(CreateAccountApplicationResponse)
	err := c.cc.Invoke(ctx, "/accountapplication.AccountApplicationService/CreateAccountApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountApplicationServiceClient) GetAccountApplication(ctx context.Context, in *GetAccountApplicationRequest, opts ...grpc.CallOption) (*AccountApplication, error) {
	out := new(AccountApplication)
	err := c.cc.Invoke(ctx, "/accountapplication.AccountApplicationService/GetAccountApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountApplicationServiceClient) GetAccountApplications(ctx context.Context, in *GetAccountApplicationsRequest, opts ...grpc.CallOption) (*GetAccountApplicationsResponse, error) {
	out := new(GetAccountApplicationsResponse)
	err := c.cc.Invoke(ctx, "/accountapplication.AccountApplicationService/GetAccountApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountApplicationServiceClient) UpdateAccountApplicationStatus(ctx context.Context, in *UpdateAccountApplicationStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accountapplication.AccountApplicationService/UpdateAccountApplicationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountApplicationServiceServer is the server API for AccountApplicationService service.
// All implementations must embed UnimplementedAccountApplicationServiceServer
// for forward compatibility
type AccountApplicationServiceServer interface {
	// CreateAccountApplication creates a new account application
	CreateAccountApplication(context.Context, *CreateAccountApplicationRequest) (*CreateAccountApplicationResponse, error)
	// GetAccountApplication retrieves the details of an account application, selected by its id.
	GetAccountApplication(context.Context, *GetAccountApplicationRequest) (*AccountApplication, error)
	// GetAccountApplications returns a list containing up to 20 accounts.
	GetAccountApplications(context.Context, *GetAccountApplicationsRequest) (*GetAccountApplicationsResponse, error)
	// UpdateAccountApplication status updates the status.
	UpdateAccountApplicationStatus(context.Context, *UpdateAccountApplicationStatusRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAccountApplicationServiceServer()
}

// UnimplementedAccountApplicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountApplicationServiceServer struct{}

func (UnimplementedAccountApplicationServiceServer) CreateAccountApplication(context.Context, *CreateAccountApplicationRequest) (*CreateAccountApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountApplication not implemented")
}

func (UnimplementedAccountApplicationServiceServer) GetAccountApplication(context.Context, *GetAccountApplicationRequest) (*AccountApplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountApplication not implemented")
}

func (UnimplementedAccountApplicationServiceServer) GetAccountApplications(context.Context, *GetAccountApplicationsRequest) (*GetAccountApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountApplications not implemented")
}

func (UnimplementedAccountApplicationServiceServer) UpdateAccountApplicationStatus(context.Context, *UpdateAccountApplicationStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountApplicationStatus not implemented")
}

func (UnimplementedAccountApplicationServiceServer) mustEmbedUnimplementedAccountApplicationServiceServer() {
}

// UnsafeAccountApplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountApplicationServiceServer will
// result in compilation errors.
type UnsafeAccountApplicationServiceServer interface {
	mustEmbedUnimplementedAccountApplicationServiceServer()
}

func RegisterAccountApplicationServiceServer(s grpc.ServiceRegistrar, srv AccountApplicationServiceServer) {
	s.RegisterService(&AccountApplicationService_ServiceDesc, srv)
}

func _AccountApplicationService_CreateAccountApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountApplicationServiceServer).CreateAccountApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountapplication.AccountApplicationService/CreateAccountApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountApplicationServiceServer).CreateAccountApplication(ctx, req.(*CreateAccountApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountApplicationService_GetAccountApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountApplicationServiceServer).GetAccountApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountapplication.AccountApplicationService/GetAccountApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountApplicationServiceServer).GetAccountApplication(ctx, req.(*GetAccountApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountApplicationService_GetAccountApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountApplicationServiceServer).GetAccountApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountapplication.AccountApplicationService/GetAccountApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountApplicationServiceServer).GetAccountApplications(ctx, req.(*GetAccountApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountApplicationService_UpdateAccountApplicationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountApplicationStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountApplicationServiceServer).UpdateAccountApplicationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountapplication.AccountApplicationService/UpdateAccountApplicationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountApplicationServiceServer).UpdateAccountApplicationStatus(ctx, req.(*UpdateAccountApplicationStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountApplicationService_ServiceDesc is the grpc.ServiceDesc for AccountApplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountApplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accountapplication.AccountApplicationService",
	HandlerType: (*AccountApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccountApplication",
			Handler:    _AccountApplicationService_CreateAccountApplication_Handler,
		},
		{
			MethodName: "GetAccountApplication",
			Handler:    _AccountApplicationService_GetAccountApplication_Handler,
		},
		{
			MethodName: "GetAccountApplications",
			Handler:    _AccountApplicationService_GetAccountApplications_Handler,
		},
		{
			MethodName: "UpdateAccountApplicationStatus",
			Handler:    _AccountApplicationService_UpdateAccountApplicationStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/accountapplication/all.proto",
}
