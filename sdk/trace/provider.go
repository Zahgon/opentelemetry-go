package trace

import (
	"context"
	"sync"
	"sync/atomic"

	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

const defaultTracerName = "go.opentelemetry.io/otel/sdk/tracer"

type tracerProviderConfig struct {
	processors []SpanProcessor

	sampler Sampler

	idGenerator IDGenerator

	spanLimits SpanLimits

	resource *resource.Resource

	panicRecordingDisabled bool
}

func (cfg tracerProviderConfig) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }

type TracerProvider struct {
	embedded.TracerProvider

	mu             sync.Mutex
	namedTracer    map[instrumentation.Scope]*tracer
	spanProcessors atomic.Pointer[spanProcessorStates]

	isShutdown atomic.Bool

	sampler                Sampler
	idGenerator            IDGenerator
	spanLimits             SpanLimits
	resource               *resource.Resource
	panicRecordingDisabled bool
}

var _ trace.TracerProvider = &TracerProvider{}

type experimentalOption interface {
	Experimental()
}

func NewTracerProvider(opts ...TracerProviderOption) *TracerProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p *TracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}

func (p *TracerProvider) RegisterSpanProcessor(sp SpanProcessor) { _ = "STUB: not implemented"; return }

func (p *TracerProvider) UnregisterSpanProcessor(sp SpanProcessor) {
	_ = "STUB: not implemented"
	return
}

func (p *TracerProvider) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TracerProvider) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *TracerProvider) getSpanProcessors() spanProcessorStates {
	_ = "STUB: not implemented"
	return *new(spanProcessorStates)
}

type TracerProviderOption interface {
	apply(tracerProviderConfig) tracerProviderConfig
}

type traceProviderOptionFunc func(tracerProviderConfig) tracerProviderConfig

func (fn traceProviderOptionFunc) apply(cfg tracerProviderConfig) tracerProviderConfig {
	_ = "STUB: not implemented"
	return *new(tracerProviderConfig)
}

func WithSyncer(e SpanExporter) TracerProviderOption {
	_ = "STUB: not implemented"
	return *new(TracerProviderOption)
}

func WithBatcher(e SpanExporter, opts ...BatchSpanProcessorOption) TracerProviderOption {
	_ = "STUB: not implemented"
	return *new(TracerProviderOption)
}

func WithSpanProcessor(sp SpanProcessor) TracerProviderOption {
	_ = "STUB: not implemented"
	return *new(TracerProviderOption)
}

func WithoutPanicRecording() TracerProviderOption {
	_ = "STUB: not implemented"
	return *new(TracerProviderOption)
}

func WithResource(r *resource.Resource) TracerProviderOption {
	_ = "STUB: not implemented"
	return *new(TracerProviderOption)
}

func WithIDGenerator(g IDGenerator) TracerProviderOption {
	_ = "STUB: not implemented"
	return *new(TracerProviderOption)
}

func WithSampler(s Sampler) TracerProviderOption {
	_ = "STUB: not implemented"
	return *new(TracerProviderOption)
}

func WithSpanLimits(sl SpanLimits) TracerProviderOption {
	_ = "STUB: not implemented"
	return *new(TracerProviderOption)
}

func WithRawSpanLimits(limits SpanLimits) TracerProviderOption {
	_ = "STUB: not implemented"
	return *new(TracerProviderOption)
}

func applyTracerProviderEnvConfigs(cfg tracerProviderConfig) tracerProviderConfig {
	_ = "STUB: not implemented"
	return *new(tracerProviderConfig)
}

func tracerProviderOptionsFromEnv() []TracerProviderOption { _ = "STUB: not implemented"; return nil }

func ensureValidTracerProviderConfig(cfg tracerProviderConfig) tracerProviderConfig {
	_ = "STUB: not implemented"
	return *new(tracerProviderConfig)
}
