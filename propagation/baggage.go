package propagation

import (
	"context"
	"sync"
)

const (
	baggageHeader = "baggage"

	maxParseErrors = 5

	maxMembers               = 64
	maxBytesPerBaggageString = 8192
)

var handleExtractErrOnce sync.Once

type Baggage struct{}

var _ TextMapPropagator = Baggage{}

func (Baggage) Inject(ctx context.Context, carrier TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

func (Baggage) Extract(parent context.Context, carrier TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (Baggage) Fields() []string { _ = "STUB: not implemented"; return nil }

func extractSingleBaggage(parent context.Context, carrier TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func extractMultiBaggage(parent context.Context, carrier ValuesGetter) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
