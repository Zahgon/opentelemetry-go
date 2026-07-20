package metric

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel/sdk/metric/internal/observ"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

const (
	defaultTimeout  = time.Millisecond * 30000
	defaultInterval = time.Millisecond * 60000
)

type periodicReaderConfig struct {
	interval                 time.Duration
	timeout                  time.Duration
	producers                []Producer
	cardinalityLimitSelector CardinalityLimitSelector
}

func newPeriodicReaderConfig(options []PeriodicReaderOption) periodicReaderConfig {
	_ = "STUB: not implemented"
	return *new(periodicReaderConfig)
}

type PeriodicReaderOption interface {
	applyPeriodic(periodicReaderConfig) periodicReaderConfig
}

type periodicReaderOptionFunc func(periodicReaderConfig) periodicReaderConfig

func (o periodicReaderOptionFunc) applyPeriodic(conf periodicReaderConfig) periodicReaderConfig {
	_ = "STUB: not implemented"
	return *new(periodicReaderConfig)
}

func WithTimeout(d time.Duration) PeriodicReaderOption {
	_ = "STUB: not implemented"
	return *new(PeriodicReaderOption)
}

func WithInterval(d time.Duration) PeriodicReaderOption {
	_ = "STUB: not implemented"
	return *new(PeriodicReaderOption)
}

func NewPeriodicReader(exporter Exporter, options ...PeriodicReaderOption) *PeriodicReader {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec  // cancel called during PeriodicReader shutdown.

var periodicReaderIDCounter atomic.Int64

func nextPeriodicReaderID() int64 { _ = "STUB: not implemented"; return 0 }

type PeriodicReader struct {
	sdkProducer atomic.Value

	mu                sync.Mutex
	isShutdown        bool
	externalProducers atomic.Value

	interval time.Duration
	timeout  time.Duration
	batcher  batcher
	exporter Exporter
	flushCh  chan chan error

	done         chan struct{}
	cancel       context.CancelFunc
	shutdownOnce sync.Once

	rmPool sync.Pool

	cardinalityLimitSelector CardinalityLimitSelector

	inst *observ.Instrumentation
}

var _ = map[Reader]struct{}{&PeriodicReader{}: {}}

var newTicker = time.NewTicker

func (r *PeriodicReader) run(ctx context.Context, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (r *PeriodicReader) register(p sdkProducer) { _ = "STUB: not implemented"; return }

func (r *PeriodicReader) temporality(kind InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func (r *PeriodicReader) aggregation(
	kind InstrumentKind,
) Aggregation {
	_ = "STUB: not implemented"
	return *new(Aggregation)
}

func (r *PeriodicReader) cardinalityLimit(kind InstrumentKind) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (r *PeriodicReader) collectAndExport(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *PeriodicReader) Collect(ctx context.Context, rm *metricdata.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *PeriodicReader) collect(ctx context.Context, p any, rm *metricdata.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *PeriodicReader) exportWithTimeout(ctx context.Context, m *metricdata.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *PeriodicReader) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *PeriodicReader) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *PeriodicReader) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }
