package requestid

import (
	"context"
)

type IDKey struct{}

const Header = "X-Request-Id"

// IntoContext injects the provided requestid into the context.
func IntoContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, IDKey{}, id)
}

// FromContext extracts the requestid from the context or returns
// the empty string.
func FromContext(ctx context.Context) string {
	if id, ok := ctx.Value(IDKey{}).(string); ok {
		return id
	}
	return ""
}
