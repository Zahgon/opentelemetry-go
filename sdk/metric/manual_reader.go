package metric

import (
	"context"
	"sync"
	"sync/atomic"

	"go.opentelemetry.io/otel/sdk/metric/internal/observ"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

const (
	manualReaderType = "go.opentelemetry.io/otel/sdk/metric/metric.ManualReader"
)

type ManualReader struct {
	sdkProducer  atomic.Value
	shutdownOnce sync.Once

	mu                sync.Mutex
	isShutdown        bool
	externalProducers atomic.Value

	temporalitySelector      TemporalitySelector
	aggregationSelector      AggregationSelector
	cardinalityLimitSelector CardinalityLimitSelector

	inst *observ.Instrumentation
}

var _ = map[Reader]struct{}{&ManualReader{}: {}}

func NewManualReader(opts ...ManualReaderOption) *ManualReader {
	_ = "STUB: not implemented"
	return nil
}

var manualReaderIDCounter atomic.Int64

func nextManualReaderID() int64 { _ = "STUB: not implemented"; return 0 }

func (mr *ManualReader) register(p sdkProducer) { _ = "STUB: not implemented"; return }

func (mr *ManualReader) temporality(kind InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func (mr *ManualReader) aggregation(
	kind InstrumentKind,
) Aggregation {
	_ = "STUB: not implemented"
	return *new(Aggregation)
}

func (mr *ManualReader) cardinalityLimit(kind InstrumentKind) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (mr *ManualReader) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (mr *ManualReader) Collect(ctx context.Context, rm *metricdata.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (mr *ManualReader) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }

type manualReaderConfig struct {
	temporalitySelector      TemporalitySelector
	aggregationSelector      AggregationSelector
	cardinalityLimitSelector CardinalityLimitSelector
	producers                []Producer
}

func newManualReaderConfig(opts []ManualReaderOption) manualReaderConfig {
	_ = "STUB: not implemented"
	return *new(manualReaderConfig)
}

type ManualReaderOption interface {
	applyManual(manualReaderConfig) manualReaderConfig
}

func WithTemporalitySelector(selector TemporalitySelector) ManualReaderOption {
	_ = "STUB: not implemented"
	return *new(ManualReaderOption)
}

type temporalitySelectorOption struct {
	selector func(instrument InstrumentKind) metricdata.Temporality
}

func (t temporalitySelectorOption) applyManual(mrc manualReaderConfig) manualReaderConfig {
	_ = "STUB: not implemented"
	return *new(manualReaderConfig)
}

func WithAggregationSelector(selector AggregationSelector) ManualReaderOption {
	_ = "STUB: not implemented"
	return *new(ManualReaderOption)
}

type aggregationSelectorOption struct {
	selector AggregationSelector
}

func (t aggregationSelectorOption) applyManual(c manualReaderConfig) manualReaderConfig {
	_ = "STUB: not implemented"
	return *new(manualReaderConfig)
}
