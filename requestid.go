package requestid

import (
	"context"
)

type IDKey struct{}

const Header = "X-Request-Id"

func IntoContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, IDKey{}, id)
}

func FromContext(ctx context.Context) string {
	if id, ok := ctx.Value(IDKey{}).(string); ok {
		return id
	}
	return ""
}
