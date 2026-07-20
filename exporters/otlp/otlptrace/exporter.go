package otlptrace

import (
	"context"
	"errors"
	"sync"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

var errAlreadyStarted = errors.New("already started")

type Exporter struct {
	client Client

	mu      sync.RWMutex
	started bool

	startOnce sync.Once
	stopOnce  sync.Once
}

func (e *Exporter) ExportSpans(ctx context.Context, ss []tracesdk.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exporter) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *Exporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var _ tracesdk.SpanExporter = (*Exporter)(nil)

func New(ctx context.Context, client Client) (*Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewUnstarted(client Client) *Exporter { _ = "STUB: not implemented"; return nil }

func (e *Exporter) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }
