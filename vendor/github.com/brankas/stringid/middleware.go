package stringid

import (
	"net/http"
	"os"
)

// GeneratorMiddleware provides standard HTTP middleware that sets a HTTP
// request ID based on the Prefix and a Generator's generated ID.
//
// Useful for logging / tracing purposes.
type GeneratorMiddleware struct {
	Prefix    string
	Generator Generator
}

// Middleware creates a standard HTTP middleware handler that sets a request ID
// on the HTTP context.
//
// If no options are supplied, then a new push-style generator will be used.
//
// By default, the Prefix will be set as "<hostname>/".
func Middleware(opts ...MiddlewareOption) func(http.Handler) http.Handler {
	var err error

	// retrieve hostname
	var hostname string
	hostname, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	if hostname == "" {
		hostname = "localhost"
	}

	// create middleware
	gm := &GeneratorMiddleware{
		Prefix: hostname + "/",
	}

	// apply options
	for _, o := range opts {
		o(gm)
	}

	// ensure generator is set
	if gm.Generator == nil {
		gm.Generator = NewPushGenerator(nil)
	}

	return gm.Handler
}

// Handler returns a standard HTTP middleware handler for the generator.
func (gm *GeneratorMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(res, WithRequest(req, gm.Prefix+gm.Generator.Generate()))
	})
}

// MiddlewareOption is a middleware option.
type MiddlewareOption func(*GeneratorMiddleware)

// WithGenerator is a middleware option to set the generator used.
func WithGenerator(generator Generator) MiddlewareOption {
	return func(gm *GeneratorMiddleware) {
		gm.Generator = generator
	}
}

// WithPrefix is a middleware option to set a prefix on generated IDs.
func WithPrefix(prefix string) MiddlewareOption {
	return func(gm *GeneratorMiddleware) {
		gm.Prefix = prefix
	}
}

// HeaderMiddleware creates a standard HTTP middleware handler that sets the
// request ID as read from one of the supplied headers.
//
// Note: this should be used in a middleware chain only *AFTER* verifying the
// request header can be trusted.
func HeaderMiddleware(headers ...string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			var id string
			for _, header := range headers {
				if id = req.Header.Get(header); id != "" {
					break
				}
			}
			h.ServeHTTP(res, WithRequest(req, id))
		})
	}
}
