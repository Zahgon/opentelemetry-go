package trace

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel/sdk/trace/internal/observ"
	"go.opentelemetry.io/otel/trace"
)

const (
	DefaultMaxQueueSize = 2048

	DefaultScheduleDelay = 5000

	DefaultExportTimeout      = 30000
	DefaultMaxExportBatchSize = 512
)

type BatchSpanProcessorOption func(o *BatchSpanProcessorOptions)

type BatchSpanProcessorOptions struct {
	MaxQueueSize int

	BatchTimeout time.Duration

	ExportTimeout time.Duration

	MaxExportBatchSize int

	BlockOnQueueFull bool
}

type batchSpanProcessor struct {
	e SpanExporter
	o BatchSpanProcessorOptions

	queue   chan ReadOnlySpan
	dropped atomic.Uint32

	inst *observ.BSP

	batch      []ReadOnlySpan
	batchMutex sync.Mutex
	timer      *time.Timer
	stopWait   sync.WaitGroup
	stopOnce   sync.Once
	stopCh     chan struct{}
	stopped    atomic.Bool
}

var _ SpanProcessor = (*batchSpanProcessor)(nil)

func NewBatchSpanProcessor(exporter SpanExporter, options ...BatchSpanProcessorOption) SpanProcessor {
	_ = "STUB: not implemented"
	return *new(SpanProcessor)
}

var processorIDCounter atomic.Int64

func nextProcessorID() int64 { _ = "STUB: not implemented"; return 0 }

func (*batchSpanProcessor) OnStart(context.Context, ReadWriteSpan) {
	_ = "STUB: not implemented"
	return
}

func (bsp *batchSpanProcessor) OnEnd(s ReadOnlySpan) { _ = "STUB: not implemented"; return }

func (bsp *batchSpanProcessor) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type forceFlushSpan struct {
	ReadOnlySpan
	flushed chan struct{}
}

func (forceFlushSpan) SpanContext() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (bsp *batchSpanProcessor) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func WithMaxQueueSize(size int) BatchSpanProcessorOption {
	_ = "STUB: not implemented"
	return *new(BatchSpanProcessorOption)
}

func WithMaxExportBatchSize(size int) BatchSpanProcessorOption {
	_ = "STUB: not implemented"
	return *new(BatchSpanProcessorOption)
}

func WithBatchTimeout(delay time.Duration) BatchSpanProcessorOption {
	_ = "STUB: not implemented"
	return *new(BatchSpanProcessorOption)
}

func WithExportTimeout(timeout time.Duration) BatchSpanProcessorOption {
	_ = "STUB: not implemented"
	return *new(BatchSpanProcessorOption)
}

func WithBlocking() BatchSpanProcessorOption {
	_ = "STUB: not implemented"
	return *new(BatchSpanProcessorOption)
}

func (bsp *batchSpanProcessor) exportSpans(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (bsp *batchSpanProcessor) processQueue() { _ = "STUB: not implemented"; return }

func (bsp *batchSpanProcessor) drainQueue() { _ = "STUB: not implemented"; return }

func (bsp *batchSpanProcessor) enqueue(sd ReadOnlySpan) { _ = "STUB: not implemented"; return }

func (bsp *batchSpanProcessor) enqueueBlockOnQueueFull(ctx context.Context, sd ReadOnlySpan) bool {
	_ = "STUB: not implemented"
	return false
}

func (bsp *batchSpanProcessor) enqueueDrop(ctx context.Context, sd ReadOnlySpan) bool {
	_ = "STUB: not implemented"
	return false
}

func (bsp *batchSpanProcessor) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }
