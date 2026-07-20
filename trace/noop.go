package trace

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace/embedded"
)

func NewNoopTracerProvider() TracerProvider { _ = "STUB: not implemented"; return *new(TracerProvider) }

type noopTracerProvider struct{ embedded.TracerProvider }

var _ TracerProvider = noopTracerProvider{}

func (noopTracerProvider) Tracer(string, ...TracerOption) Tracer {
	_ = "STUB: not implemented"
	return *new(Tracer)
}

type noopTracer struct{ embedded.Tracer }

var _ Tracer = noopTracer{}

func (noopTracer) Start(ctx context.Context, _ string, _ ...SpanStartOption) (context.Context, Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(Span)
}

type noopSpan struct{ embedded.Span }

var noopSpanInstance Span = noopSpan{}

func (noopSpan) SpanContext() SpanContext { _ = "STUB: not implemented"; return *new(SpanContext) }

func (noopSpan) IsRecording() bool { _ = "STUB: not implemented"; return false }

func (noopSpan) SetStatus(codes.Code, string) { _ = "STUB: not implemented"; return }

func (noopSpan) SetError(bool) { _ = "STUB: not implemented"; return }

func (noopSpan) SetAttributes(...attribute.KeyValue) { _ = "STUB: not implemented"; return }

func (noopSpan) End(...SpanEndOption) { _ = "STUB: not implemented"; return }

func (noopSpan) RecordError(error, ...EventOption) { _ = "STUB: not implemented"; return }

func (noopSpan) AddEvent(string, ...EventOption) { _ = "STUB: not implemented"; return }

func (noopSpan) AddLink(Link) { _ = "STUB: not implemented"; return }

func (noopSpan) SetName(string) { _ = "STUB: not implemented"; return }

func (s noopSpan) TracerProvider() TracerProvider {
	_ = "STUB: not implemented"
	return *new(TracerProvider)
}

var autoInstEnabled = new(bool)

//go:noinline
func (noopSpan) tracerProvider(autoEnabled *bool) TracerProvider {
	_ = "STUB: not implemented"
	return *new(TracerProvider)
}
