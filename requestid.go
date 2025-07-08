package requestid

import (
	"context"
)

type RequestIDKey struct{}

func IntoContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, RequestIDKey{}, id)
}

func FromContext(ctx context.Context) string {
	if id, ok := ctx.Value(RequestIDKey{}).(string); ok {
		return id
	}
	return ""
}
