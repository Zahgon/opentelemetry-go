package trace

import (
	"context"

	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/trace/internal/observ"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

type tracer struct {
	embedded.Tracer

	provider             *TracerProvider
	instrumentationScope instrumentation.Scope

	inst observ.Tracer
}

var _ trace.Tracer = &tracer{}

func (tr *tracer) Start(
	ctx context.Context,
	name string,
	options ...trace.SpanStartOption,
) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

type runtimeTracer interface {
	runtimeTrace(ctx context.Context) context.Context
}

func (tr *tracer) newSpan(ctx context.Context, name string, config *trace.SpanConfig) trace.Span {
	_ = "STUB: not implemented"
	return *new(trace.Span)
}

func (tr *tracer) newRecordingSpan(
	ctx context.Context,
	psc, sc trace.SpanContext,
	name string,
	sr SamplingResult,
	config *trace.SpanConfig,
) *recordingSpan {
	_ = "STUB: not implemented"
	return nil
}

func (tr *tracer) newNonRecordingSpan(sc trace.SpanContext) nonRecordingSpan {
	_ = "STUB: not implemented"
	return *new(nonRecordingSpan)
}
