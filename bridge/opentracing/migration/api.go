package migration

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type DeferredContextSetupTracerExtension interface {
	DeferredContextSetupHook(ctx context.Context, span trace.Span) context.Context
}

type OverrideTracerSpanExtension interface {
	OverrideTracer(tracer trace.Tracer)
}
