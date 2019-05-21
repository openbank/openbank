package services

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// RegisterServices creates and registers grpc services.
func RegisterServices(srv *grpc.Server, services ...func(srv *grpc.Server) error) error {
	for _, svc := range services {
		if err := svc(srv); err != nil {
			return err
		}
	}
	return nil
}

// RegisterGateway registers services for grpc-gateway server.
func RegisterGateway(mux *runtime.ServeMux, gwAddr string, opts []grpc.DialOption,
	services ...func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error) error {
	ctx := context.Background()
	for _, svc := range services {
		if err := svc(ctx, mux, gwAddr, opts); err != nil {
			return err
		}
	}
	return nil
}
