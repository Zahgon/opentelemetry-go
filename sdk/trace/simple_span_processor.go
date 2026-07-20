package trace

import (
	"context"
	"sync"
	"sync/atomic"

	"go.opentelemetry.io/otel/sdk/trace/internal/observ"
)

type simpleSpanProcessor struct {
	exporterMu sync.Mutex
	exporter   SpanExporter
	stopOnce   sync.Once

	inst *observ.SSP
}

var _ SpanProcessor = (*simpleSpanProcessor)(nil)

func NewSimpleSpanProcessor(exporter SpanExporter) SpanProcessor {
	_ = "STUB: not implemented"
	return *new(SpanProcessor)
}

var simpleProcessorIDCounter atomic.Int64

func nextSimpleProcessorID() int64 { _ = "STUB: not implemented"; return 0 }

func (*simpleSpanProcessor) OnStart(context.Context, ReadWriteSpan) {
	_ = "STUB: not implemented"
	return
}

func (ssp *simpleSpanProcessor) OnEnd(s ReadOnlySpan) { _ = "STUB: not implemented"; return }

func (ssp *simpleSpanProcessor) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (*simpleSpanProcessor) ForceFlush(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (ssp *simpleSpanProcessor) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }
