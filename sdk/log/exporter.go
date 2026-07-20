package log

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type Exporter interface {
	Export(ctx context.Context, records []Record) error

	Shutdown(ctx context.Context) error

	ForceFlush(ctx context.Context) error
}

var defaultNoopExporter = &noopExporter{}

type noopExporter struct{}

func (noopExporter) Export(context.Context, []Record) error { _ = "STUB: not implemented"; return nil }

func (noopExporter) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (noopExporter) ForceFlush(context.Context) error { _ = "STUB: not implemented"; return nil }

type chunkExporter struct {
	Exporter

	size int
}

func newChunkExporter(exporter Exporter, size int) Exporter {
	_ = "STUB: not implemented"
	return *new(Exporter)
}

func (c chunkExporter) Export(ctx context.Context, records []Record) error {
	_ = "STUB: not implemented"
	return nil
}

type timeoutExporter struct {
	Exporter

	timeout time.Duration
}

func newTimeoutExporter(exp Exporter, timeout time.Duration) Exporter {
	_ = "STUB: not implemented"
	return *new(Exporter)
}

func (e *timeoutExporter) Export(ctx context.Context, records []Record) error {
	_ = "STUB: not implemented"
	return nil
}

func exportSync(input <-chan exportData, exporter Exporter) (done chan struct{}) {
	_ = "STUB: not implemented"
	return nil
}

type exportData struct {
	ctx     context.Context
	records []Record

	respCh chan<- error
}

func (e exportData) DoExport(exportFn func(context.Context, []Record) error) {
	_ = "STUB: not implemented"
	return
}

func (e exportData) respond(err error) { _ = "STUB: not implemented"; return }

type bufferExporter struct {
	Exporter

	input   chan exportData
	inputMu sync.Mutex

	done    chan struct{}
	stopped atomic.Bool
}

func newBufferExporter(exporter Exporter, size int) *bufferExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *bufferExporter) Ready() bool { _ = "STUB: not implemented"; return false }

var errStopped = errors.New("exporter stopped")

func (e *bufferExporter) enqueue(ctx context.Context, records []Record, rCh chan<- error) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *bufferExporter) EnqueueExport(records []Record) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *bufferExporter) Export(ctx context.Context, records []Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *bufferExporter) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *bufferExporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
