package opentracing

import (
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

type TracerProvider struct {
	embedded.TracerProvider

	bridge   *BridgeTracer
	provider trace.TracerProvider

	tracers map[wrappedTracerKey]*WrapperTracer
	mtx     sync.Mutex
}

var _ trace.TracerProvider = (*TracerProvider)(nil)

func NewTracerProvider(bridge *BridgeTracer, provider trace.TracerProvider) *TracerProvider {
	_ = "STUB: not implemented"
	return nil
}

type wrappedTracerKey struct {
	name    string
	version string
	schema  string
	attrs   attribute.Set
}

func (p *TracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}
