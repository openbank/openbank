package apierror

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apierrorpb "github.com/openbank/gunk/gunk/v1/apierror"
)

var baseDocURL string

// APIError embedds APIError definition from openbank/gunk.
type APIError struct {
	apierrorpb.APIError
}

// Error returns the error message.
func (a APIError) Error() string {
	return a.Message
}

// GRPCStatus returns the gRPC status according to the APIError.
func (a APIError) GRPCStatus() *status.Status {
	s := status.New(codes.Code(a.GetCode()), a.GetMessage())
	s, _ = s.WithDetails(&a.APIError)
	return s
}

// MarshalJSON is the custom json marshaling for an APIError.
func (a APIError) MarshalJSON() ([]byte, error) {
	err := map[string]interface{}{
		"message": a.GetMessage(),
		"doc_url": a.GetDocURL(),
	}

	if a.GetType() != apierrorpb.ErrorType__ {
		err["type"] = a.GetType().String()
	}

	vals := map[string]interface{}{}
	vals["error"] = err
	return json.Marshal(vals)
}

func newError(code codes.Code, errorType apierrorpb.ErrorType, docURL, format string, a ...interface{}) APIError {
	aerr := APIError{
		apierrorpb.APIError{
			Message: fmt.Sprintf(format, a...),
			Code:    int32(code),
			DocURL:  docURL,
		},
	}
	return aerr
}

// NotFound creates a not found error.
func NotFound(format string, a ...interface{}) APIError {
	return newError(codes.NotFound, apierrorpb.ErrorType__, baseDocURL, format, a...)
}

// Internal creates an internal server error.
func Internal(err error) APIError {
	return newError(codes.Internal, apierrorpb.ErrorType__, baseDocURL, err.Error())
}

// BadRequest creates a bad request error.
func BadRequest(errorType apierrorpb.ErrorType, format string, a ...interface{}) APIError {
	docURL := baseDocURL + "/" + errorType.String()
	return newError(codes.InvalidArgument, errorType, docURL, format, a...)
}

// Conflict creates a conflict error.
func Conflict(errorType apierrorpb.ErrorType, format string, a ...interface{}) APIError {
	docURL := baseDocURL + "/" + errorType.String()
	return newError(codes.AlreadyExists, errorType, docURL, format, a...)
}

func fromStatus(s *status.Status) APIError {
	aerr := APIError{
		apierrorpb.APIError{
			Message: s.Message(),
			Code:    int32(s.Code()),
		},
	}
	for _, d := range s.Details() {
		switch d.(type) {
		case *apierrorpb.APIError:
			aerr.DocURL = d.(*apierrorpb.APIError).DocURL
			aerr.Type = d.(*apierrorpb.APIError).Type
		}
	}
	return aerr
}

// CustomHTTPError overrides default http error format from gRPC runtime.
func CustomHTTPError(baseURL string) func(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	baseDocURL = baseURL

	return func(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
		const fallback = `{"error": {"message":"failed to marshal error message"}}`

		w.Header().Set("Content-type", marshaler.ContentType())
		w.WriteHeader(runtime.HTTPStatusFromCode(grpc.Code(err)))

		// the error intercepted is of type statusError
		s, _ := status.FromError(err)
		// extract details from statusError and return APIError
		aerr := fromStatus(s)

		if jerr := json.NewEncoder(w).Encode(aerr); jerr != nil {
			w.Write([]byte(fallback))
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
