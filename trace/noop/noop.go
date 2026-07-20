package noop

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

var (
	_ trace.TracerProvider = TracerProvider{}
	_ trace.Tracer         = Tracer{}
	_ trace.Span           = Span{}
)

type TracerProvider struct{ embedded.TracerProvider }

func NewTracerProvider() TracerProvider { _ = "STUB: not implemented"; return *new(TracerProvider) }

func (TracerProvider) Tracer(string, ...trace.TracerOption) trace.Tracer {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}

type Tracer struct{ embedded.Tracer }

func (Tracer) Start(ctx context.Context, _ string, _ ...trace.SpanStartOption) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

var noopSpanInstance trace.Span = Span{}

type Span struct {
	embedded.Span

	sc trace.SpanContext
}

func (s Span) SpanContext() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (Span) IsRecording() bool { _ = "STUB: not implemented"; return false }

func (Span) SetStatus(codes.Code, string) { _ = "STUB: not implemented"; return }

func (Span) SetAttributes(...attribute.KeyValue) { _ = "STUB: not implemented"; return }

func (Span) End(...trace.SpanEndOption) { _ = "STUB: not implemented"; return }

func (Span) RecordError(error, ...trace.EventOption) { _ = "STUB: not implemented"; return }

func (Span) AddEvent(string, ...trace.EventOption) { _ = "STUB: not implemented"; return }

func (Span) AddLink(trace.Link) { _ = "STUB: not implemented"; return }

func (Span) SetName(string) { _ = "STUB: not implemented"; return }

func (Span) TracerProvider() trace.TracerProvider {
	_ = "STUB: not implemented"
	return *new(trace.TracerProvider)
}
