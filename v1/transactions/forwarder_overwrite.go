// Package transactions ...
// Using customer forwarders allow us to return the appropriate 2xx status code
// for a specific operation instead of the 200 always returned by grpc-gateway.
// https://github.com/grpc-ecosystem/grpc-gateway/blob/master/docs/_docs/customizingyourgateway.md#replace-a-response-forwarder-per-method
package transactions

import (
	"net/http"

	proto "github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	context "golang.org/x/net/context"
)

// forwardCreateTransaction overwrite the default behavior of the ForwardResponseMessage.
// Set by default the status code to 201, errors are handled in runtime.ForwardResponseMessage.
func forwardCreateTransaction(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	w.WriteHeader(http.StatusCreated)
	runtime.ForwardResponseMessage(ctx, mux, marshaler, w, req, resp, opts...)
}

func init() {
	forward_TransactionService_CreateTransaction_0 = forwardCreateTransaction
}
