package baggage

import (
	"context"
)

func ContextWithBaggage(parent context.Context, b Baggage) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ContextWithoutBaggage(parent context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) Baggage { _ = "STUB: not implemented"; return *new(Baggage) }
