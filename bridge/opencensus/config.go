package opencensus

import (
	"go.opentelemetry.io/otel/trace"
)

const scopeName = "go.opentelemetry.io/otel/bridge/opencensus"

func newTraceConfig(options []TraceOption) traceConfig {
	_ = "STUB: not implemented"
	return *new(traceConfig)
}

type traceConfig struct {
	tp trace.TracerProvider
}

type TraceOption interface {
	apply(traceConfig) traceConfig
}

type traceOptionFunc func(traceConfig) traceConfig

func (o traceOptionFunc) apply(conf traceConfig) traceConfig {
	_ = "STUB: not implemented"
	return *new(traceConfig)
}

func WithTracerProvider(tp trace.TracerProvider) TraceOption {
	_ = "STUB: not implemented"
	return *new(TraceOption)
}

type metricConfig struct{}

type MetricOption interface {
	apply(metricConfig) metricConfig
}
