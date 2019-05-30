package logging

import (
	"context"

	"github.com/brankas/stringid"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	correlationIDKey = "correlation-id"
)

// FromContext will check the context for a correlation id, if the correlation id
// isn't present then it will generate one, the correlation id will be a field
// in the logger
func FromContext(ctx context.Context) *logrus.Entry {
	cid := CorrelationIDFromContext(ctx)
	return NewLogger().WithFields(logrus.Fields{correlationIDKey: cid})
}

// CorrelationIdFromContext will return the correlation id from the context
// metadata if present, otherwise will return an empty string
func CorrelationIDFromContext(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		correlationIDs := md[correlationIDKey]
		if len(correlationIDs) > 0 {
			return correlationIDs[0]
		}
	}
	return ""
}

// ContextWithCorrelationID will generate a correlation id and add it to the
// contexts metadata if one doesn't already exist
func ContextWithCorrelationID(ctx context.Context) context.Context {
	cid := CorrelationIDFromContext(ctx)
	if cid == "" {
		// If no correlation id is present on the context generate one and attach it
		// to the contexts metadata
		cid := generateCorrelationID()
		// Check to see if there is any metadata on the context
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			md[correlationIDKey] = []string{cid}
		} else {
			md = metadata.Pairs(correlationIDKey, cid)
		}
		// Build a new context with the correlation id on the incoming metadata
		ctx = metadata.NewIncomingContext(ctx, md)
	}
	ctx = metadata.AppendToOutgoingContext(ctx, correlationIDKey, cid)
	return ctx
}

// UnaryServerInterceptor will make sure a correlation id exists on the context
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = ContextWithCorrelationID(ctx)
		// Add the correlation id to the grpc logrus so that we get the correlation id
		// in middleware logs
		grpc_ctxtags.Extract(ctx).Set(correlationIDKey, CorrelationIDFromContext(ctx))
		return handler(ctx, req)
	}
}

// generateCorrelationId generates a random uuid to use as the correlation id
func generateCorrelationID() string {
	return stringid.Generate()
}
