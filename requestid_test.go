package requestid

import (
	"context"
	"testing"
)

func TestFromContext(t *testing.T) {
	t.Run("WithRealString", func(t *testing.T) {
		id := "abcd1234"
		actual := FromContext(IntoContext(context.Background(), id))
		if actual != id {
			t.Errorf("FromContext(%v) = %v; want %v", id, actual, id)
		}
	})
	t.Run("WhenEmpty", func(t *testing.T) {
		actual := FromContext(context.Background())
		if actual != "" {
			t.Errorf("FromContext(%v) = %v; want ''", actual, actual)
		}
	})
}
