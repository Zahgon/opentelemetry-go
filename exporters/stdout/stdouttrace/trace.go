package stdouttrace

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace/internal/observ"
	"go.opentelemetry.io/otel/sdk/trace"
)

var zeroTime time.Time

var _ trace.SpanExporter = &Exporter{}

func New(options ...Option) (*Exporter, error) { _ = "STUB: not implemented"; return nil, nil }

type Exporter struct {
	encoder    *json.Encoder
	encoderMu  sync.Mutex
	timestamps bool

	stoppedMu sync.RWMutex
	stopped   bool

	inst *observ.Instrumentation
}

func (e *Exporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exporter) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *Exporter) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }
