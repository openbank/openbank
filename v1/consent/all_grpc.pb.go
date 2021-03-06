// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package consent

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConsentServiceClient is the client API for ConsentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsentServiceClient interface {
	// AnswerConsentChallenge sends the answer to the challenge for consnet challenge request
	AnswerConsentChallenge(ctx context.Context, in *AnswerConsentChallengeRequest, opts ...grpc.CallOption) (*Consent, error)
	// CreateConsentEmail creates a new email consent
	CreateConsentEmail(ctx context.Context, in *CreateConsentEmailRequest, opts ...grpc.CallOption) (*Consent, error)
	// CreateConsentSMS creates a new sms consent
	CreateConsentSMS(ctx context.Context, in *CreateConsentSMSRequest, opts ...grpc.CallOption) (*Consent, error)
	// GetConsents returns a list containing up to 20 consents.
	GetConsents(ctx context.Context, in *GetConsentsRequest, opts ...grpc.CallOption) (*GetConsentsResponse, error)
	// RevokeConsent revokes the consent
	RevokeConsent(ctx context.Context, in *RevokeConsentRequest, opts ...grpc.CallOption) (*Consent, error)
}

type consentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsentServiceClient(cc grpc.ClientConnInterface) ConsentServiceClient {
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

func (c *consentServiceClient) GetConsents(ctx context.Context, in *GetConsentsRequest, opts ...grpc.CallOption) (*GetConsentsResponse, error) {
	out := new(GetConsentsResponse)
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
// All implementations must embed UnimplementedConsentServiceServer
// for forward compatibility
type ConsentServiceServer interface {
	// AnswerConsentChallenge sends the answer to the challenge for consnet challenge request
	AnswerConsentChallenge(context.Context, *AnswerConsentChallengeRequest) (*Consent, error)
	// CreateConsentEmail creates a new email consent
	CreateConsentEmail(context.Context, *CreateConsentEmailRequest) (*Consent, error)
	// CreateConsentSMS creates a new sms consent
	CreateConsentSMS(context.Context, *CreateConsentSMSRequest) (*Consent, error)
	// GetConsents returns a list containing up to 20 consents.
	GetConsents(context.Context, *GetConsentsRequest) (*GetConsentsResponse, error)
	// RevokeConsent revokes the consent
	RevokeConsent(context.Context, *RevokeConsentRequest) (*Consent, error)
	mustEmbedUnimplementedConsentServiceServer()
}

// UnimplementedConsentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConsentServiceServer struct{}

func (UnimplementedConsentServiceServer) AnswerConsentChallenge(context.Context, *AnswerConsentChallengeRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerConsentChallenge not implemented")
}

func (UnimplementedConsentServiceServer) CreateConsentEmail(context.Context, *CreateConsentEmailRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsentEmail not implemented")
}

func (UnimplementedConsentServiceServer) CreateConsentSMS(context.Context, *CreateConsentSMSRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsentSMS not implemented")
}

func (UnimplementedConsentServiceServer) GetConsents(context.Context, *GetConsentsRequest) (*GetConsentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsents not implemented")
}

func (UnimplementedConsentServiceServer) RevokeConsent(context.Context, *RevokeConsentRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeConsent not implemented")
}
func (UnimplementedConsentServiceServer) mustEmbedUnimplementedConsentServiceServer() {}

// UnsafeConsentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsentServiceServer will
// result in compilation errors.
type UnsafeConsentServiceServer interface {
	mustEmbedUnimplementedConsentServiceServer()
}

func RegisterConsentServiceServer(s grpc.ServiceRegistrar, srv ConsentServiceServer) {
	s.RegisterService(&ConsentService_ServiceDesc, srv)
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

// ConsentService_ServiceDesc is the grpc.ServiceDesc for ConsentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsentService_ServiceDesc = grpc.ServiceDesc{
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
