package otlptracegrpc

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
)

func New(ctx context.Context, opts ...Option) (*otlptrace.Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewUnstarted(opts ...Option) *otlptrace.Exporter { _ = "STUB: not implemented"; return nil }
