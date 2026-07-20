package tracetest

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/sdk/trace"
)

var _ trace.SpanExporter = (*NoopExporter)(nil)

func NewNoopExporter() *NoopExporter { _ = "STUB: not implemented"; return nil }

type NoopExporter struct{}

func (*NoopExporter) ExportSpans(context.Context, []trace.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (*NoopExporter) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

var _ trace.SpanExporter = (*InMemoryExporter)(nil)

func NewInMemoryExporter() *InMemoryExporter { _ = "STUB: not implemented"; return nil }

type InMemoryExporter struct {
	mu sync.Mutex
	ss SpanStubs
}

func (imsb *InMemoryExporter) ExportSpans(_ context.Context, spans []trace.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (imsb *InMemoryExporter) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (imsb *InMemoryExporter) Reset() { _ = "STUB: not implemented"; return }

func (imsb *InMemoryExporter) GetSpans() SpanStubs {
	_ = "STUB: not implemented"
	return *new(SpanStubs)
}
