package otel

import (
	"go.opentelemetry.io/otel/trace"
)

func Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}

func GetTracerProvider() trace.TracerProvider {
	_ = "STUB: not implemented"
	return *new(trace.TracerProvider)
}

func SetTracerProvider(tp trace.TracerProvider) { _ = "STUB: not implemented"; return }
