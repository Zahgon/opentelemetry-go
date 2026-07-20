package trace

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type IDGenerator interface {
	NewIDs(ctx context.Context) (trace.TraceID, trace.SpanID)

	NewSpanID(ctx context.Context, traceID trace.TraceID) trace.SpanID
}

type randomIDGenerator struct{}

var _ IDGenerator = &randomIDGenerator{}

func (*randomIDGenerator) NewSpanID(context.Context, trace.TraceID) trace.SpanID {
	_ = "STUB: not implemented"
	return *new(trace.SpanID)
}

func (*randomIDGenerator) NewIDs(context.Context) (trace.TraceID, trace.SpanID) {
	_ = "STUB: not implemented"
	return *new(trace.TraceID), *new(trace.SpanID)
}

func defaultIDGenerator() IDGenerator { _ = "STUB: not implemented"; return *new(IDGenerator) }
