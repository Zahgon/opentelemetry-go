package propagation

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
)

const (
	supportedVersion  = 0
	maxVersion        = 254
	traceparentHeader = "traceparent"
	tracestateHeader  = "tracestate"
	delimiter         = "-"
)

type TraceContext struct{}

var (
	_           TextMapPropagator = TraceContext{}
	versionPart                   = fmt.Sprintf("%.2X", supportedVersion)
)

func (TraceContext) Inject(ctx context.Context, carrier TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

func (tc TraceContext) Extract(ctx context.Context, carrier TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (TraceContext) extract(carrier TextMapCarrier) trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

//nolint:gosec // slice size already checked.

func upperHex(v string) bool { _ = "STUB: not implemented"; return false }

func extractPart(dst []byte, h *string, n int) bool { _ = "STUB: not implemented"; return false }

func (TraceContext) Fields() []string { _ = "STUB: not implemented"; return nil }
