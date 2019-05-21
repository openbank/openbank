package stringid

import (
	"context"
	"net/http"
)

// contextKey is a context key.
type contextKey int

// idKey is the context key for a request.
const idKey contextKey = 0

// WithID creates a new context with the supplied ID.
func WithID(ctxt context.Context, id string) context.Context {
	return context.WithValue(ctxt, idKey, id)
}

// WithRequest returns the supplied HTTP request with a new context containing
// the supplied ID.
func WithRequest(req *http.Request, id string) *http.Request {
	return req.WithContext(WithID(req.Context(), id))
}

// FromContext returns the ID previously set on the context.
func FromContext(ctxt context.Context) string {
	id, ok := ctxt.Value(idKey).(string)
	if ok {
		return id
	}
	return ""
}

// FromRequest returns the ID previously set on the request's context.
func FromRequest(req *http.Request) string {
	return FromContext(req.Context())
}
