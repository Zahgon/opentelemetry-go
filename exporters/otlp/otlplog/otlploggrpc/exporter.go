package otlploggrpc

import (
	"context"
	"sync"
	"sync/atomic"

	logpb "go.opentelemetry.io/proto/otlp/logs/v1"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc/internal/transform"
	"go.opentelemetry.io/otel/sdk/log"
)

type logClient interface {
	UploadLogs(ctx context.Context, rl []*logpb.ResourceLogs) error
	Shutdown(context.Context) error
}

type Exporter struct {
	clientMu sync.Mutex
	client   logClient

	stopped atomic.Bool
}

var _ log.Exporter = (*Exporter)(nil)

func New(_ context.Context, options ...Option) (*Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newExporter(c logClient) *Exporter { _ = "STUB: not implemented"; return nil }

var transformResourceLogs = transform.ResourceLogs

func (e *Exporter) Export(ctx context.Context, records []log.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (*Exporter) ForceFlush(context.Context) error { _ = "STUB: not implemented"; return nil }
