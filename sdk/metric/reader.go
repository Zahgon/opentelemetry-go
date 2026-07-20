package metric

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

var errDuplicateRegister = errors.New("duplicate reader registration")

var ErrReaderNotRegistered = errors.New("reader is not registered")

var ErrReaderShutdown = errors.New("reader is shutdown")

var errNonPositiveDuration = errors.New("non-positive duration")

type Reader interface {
	register(sdkProducer)

	temporality(InstrumentKind) metricdata.Temporality

	aggregation(InstrumentKind) Aggregation

	cardinalityLimit(InstrumentKind) (limit int, fallback bool)

	Collect(ctx context.Context, rm *metricdata.ResourceMetrics) error

	Shutdown(context.Context) error
}

type sdkProducer interface {
	produce(context.Context, *metricdata.ResourceMetrics) error
}

type Producer interface {
	Produce(context.Context) ([]metricdata.ScopeMetrics, error)
}

type produceHolder struct {
	produce func(context.Context, *metricdata.ResourceMetrics) error
}

type shutdownProducer struct{}

func (shutdownProducer) produce(context.Context, *metricdata.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

type TemporalitySelector func(InstrumentKind) metricdata.Temporality

func DefaultTemporalitySelector(k InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func CumulativeTemporalitySelector(InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func DeltaTemporalitySelector(k InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func LowMemoryTemporalitySelector(k InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

type AggregationSelector func(InstrumentKind) Aggregation

func DefaultAggregationSelector(ik InstrumentKind) Aggregation {
	_ = "STUB: not implemented"
	return *new(Aggregation)
}

type CardinalityLimitSelector func(InstrumentKind) (limit int, fallback bool)

func defaultCardinalityLimitSelector(_ InstrumentKind) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type ReaderOption interface {
	PeriodicReaderOption
	ManualReaderOption
}

func WithProducer(p Producer) ReaderOption { _ = "STUB: not implemented"; return *new(ReaderOption) }

type producerOption struct {
	p Producer
}

func (o producerOption) applyManual(c manualReaderConfig) manualReaderConfig {
	_ = "STUB: not implemented"
	return *new(manualReaderConfig)
}

func (o producerOption) applyPeriodic(c periodicReaderConfig) periodicReaderConfig {
	_ = "STUB: not implemented"
	return *new(periodicReaderConfig)
}

func WithCardinalityLimitSelector(selector CardinalityLimitSelector) ReaderOption {
	_ = "STUB: not implemented"
	return *new(ReaderOption)
}

type cardinalityLimitSelectorOption struct {
	selector CardinalityLimitSelector
}

func (o cardinalityLimitSelectorOption) applyManual(c manualReaderConfig) manualReaderConfig {
	_ = "STUB: not implemented"
	return *new(manualReaderConfig)
}

func (o cardinalityLimitSelectorOption) applyPeriodic(c periodicReaderConfig) periodicReaderConfig {
	_ = "STUB: not implemented"
	return *new(periodicReaderConfig)
}
