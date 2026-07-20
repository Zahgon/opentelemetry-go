package opentracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func NewTracerPair(tracer trace.Tracer) (*BridgeTracer, *WrapperTracerProvider) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTracerPairWithContext(
	ctx context.Context,
	tracer trace.Tracer,
) (context.Context, *BridgeTracer, *WrapperTracerProvider) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}
