package log

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/sdk/log/internal/observ"
)

var _ Processor = (*SimpleProcessor)(nil)

type SimpleProcessor struct {
	mu       sync.Mutex
	exporter Exporter
	inst     *observ.SLP
	noCmp    [0]func() //nolint: unused  // This is indeed used.
}

func NewSimpleProcessor(exporter Exporter, _ ...SimpleProcessorOption) *SimpleProcessor {
	_ = "STUB: not implemented"
	return nil
}

var simpleProcRecordsPool = sync.Pool{
	New: func() any {
		records := make([]Record, 1)
		return &records
	},
}

func (*SimpleProcessor) Enabled(context.Context, EnabledParameters) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SimpleProcessor) OnEmit(ctx context.Context, r *Record) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimpleProcessor) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimpleProcessor) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type SimpleProcessorOption interface {
	apply()
}
