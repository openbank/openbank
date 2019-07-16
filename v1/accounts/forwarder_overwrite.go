// Package accounts ...
// Using customer forwarders allow us to return the appropriate 2xx status code
// for a specific operation instead of the 200 always returned by grpc-gateway.
// https://github.com/grpc-ecosystem/grpc-gateway/blob/master/docs/_docs/customizingyourgateway.md#replace-a-response-forwarder-per-method
package accounts

import (
	"net/http"

	proto "github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	context "golang.org/x/net/context"
)

// forwardCreateAccount overwrite the default behavior of the ForwardResponseMessage.
// Set by default the status code to 201, errors are handled in runtime.ForwardResponseMessage.
func forwardCreateAccount(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	w.WriteHeader(http.StatusCreated)
	runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
}

// forwardUpdateAccount overwrite the default behavior of the ForwardResponseMessage.
// Set by default the status code to 204, errors are handled in runtime.ForwardResponseMessage.
func forwardUpdateAccount(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	w.WriteHeader(http.StatusNoContent)
	runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
}

// forwardDeleteAccount overwrite the default behavior of the ForwardResponseMessage.
// Set by default the status code to 204, errors are handled in runtime.ForwardResponseMessage.
func forwardDeleteAccount(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	w.WriteHeader(http.StatusNoContent)
	runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
}

func init() {
	forward_AccountService_CreateAccount_0 = forwardCreateAccount

	forward_AccountService_UpdateAccount_0 = forwardUpdateAccount

	forward_AccountService_DeleteAccount_0 = forwardDeleteAccount
}
