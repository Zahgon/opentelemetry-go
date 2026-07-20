package opentracing

import (
	"context"

	"go.opentelemetry.io/otel/bridge/opentracing/migration"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

type WrapperTracerProvider struct {
	embedded.TracerProvider

	wTracer *WrapperTracer
}

var _ trace.TracerProvider = (*WrapperTracerProvider)(nil)

func (p *WrapperTracerProvider) Tracer(string, ...trace.TracerOption) trace.Tracer {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}

func NewWrappedTracerProvider(bridge *BridgeTracer, tracer trace.Tracer) *WrapperTracerProvider {
	_ = "STUB: not implemented"
	return nil
}

type WrapperTracer struct {
	embedded.Tracer

	bridge *BridgeTracer
	tracer trace.Tracer
}

var (
	_ trace.Tracer                                  = &WrapperTracer{}
	_ migration.DeferredContextSetupTracerExtension = &WrapperTracer{}
)

func NewWrapperTracer(bridge *BridgeTracer, tracer trace.Tracer) *WrapperTracer {
	_ = "STUB: not implemented"
	return nil
}

func (t *WrapperTracer) otelTracer() trace.Tracer {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}

func (t *WrapperTracer) Start(
	ctx context.Context,
	name string,
	opts ...trace.SpanStartOption,
) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

func (t *WrapperTracer) DeferredContextSetupHook(ctx context.Context, span trace.Span) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
