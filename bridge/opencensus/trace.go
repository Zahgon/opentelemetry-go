package opencensus

import (
	octrace "go.opencensus.io/trace"

	"go.opentelemetry.io/otel/trace"
)

func InstallTraceBridge(opts ...TraceOption) { _ = "STUB: not implemented"; return }

func newTraceBridge(opts []TraceOption) octrace.Tracer {
	_ = "STUB: not implemented"
	return *new(octrace.Tracer)
}

func OTelSpanContextToOC(sc trace.SpanContext) octrace.SpanContext {
	_ = "STUB: not implemented"
	return *new(octrace.SpanContext)
}

func OCSpanContextToOTel(sc octrace.SpanContext) trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}
