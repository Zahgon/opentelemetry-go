package otlploghttp

import (
	"context"
	"sync/atomic"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp/internal/transform"
	"go.opentelemetry.io/otel/sdk/log"
)

type Exporter struct {
	client  atomic.Pointer[client]
	stopped atomic.Bool
}

var _ log.Exporter = (*Exporter)(nil)

func New(ctx context.Context, options ...Option) (*Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newExporter(c *client, _ config) (*Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var transformResourceLogs = transform.ResourceLogs

func (e *Exporter) Export(ctx context.Context, records []log.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exporter) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (*Exporter) ForceFlush(context.Context) error { _ = "STUB: not implemented"; return nil }
