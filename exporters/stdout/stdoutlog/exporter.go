package stdoutlog

import (
	"context"
	"encoding/json"
	"sync/atomic"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog/internal/observ"

	"go.opentelemetry.io/otel/sdk/log"
)

var _ log.Exporter = &Exporter{}

type Exporter struct {
	encoder    atomic.Pointer[json.Encoder]
	timestamps bool
	inst       *observ.Instrumentation
}

func New(options ...Option) (*Exporter, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Exporter) Export(ctx context.Context, records []log.Record) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exporter) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (*Exporter) ForceFlush(context.Context) error { _ = "STUB: not implemented"; return nil }
