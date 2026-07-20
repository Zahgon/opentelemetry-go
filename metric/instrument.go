package metric

import (
	"go.opentelemetry.io/otel/attribute"
)

type Observable interface {
	observable()
}

type InstrumentOption interface {
	Int64CounterOption
	Int64UpDownCounterOption
	Int64HistogramOption
	Int64GaugeOption
	Int64ObservableCounterOption
	Int64ObservableUpDownCounterOption
	Int64ObservableGaugeOption

	Float64CounterOption
	Float64UpDownCounterOption
	Float64HistogramOption
	Float64GaugeOption
	Float64ObservableCounterOption
	Float64ObservableUpDownCounterOption
	Float64ObservableGaugeOption
}

type HistogramOption interface {
	Int64HistogramOption
	Float64HistogramOption
}

type descOpt string

func (o descOpt) applyFloat64Counter(c Float64CounterConfig) Float64CounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64CounterConfig)
}

func (o descOpt) applyFloat64UpDownCounter(c Float64UpDownCounterConfig) Float64UpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64UpDownCounterConfig)
}

func (o descOpt) applyFloat64Histogram(c Float64HistogramConfig) Float64HistogramConfig {
	_ = "STUB: not implemented"
	return *new(Float64HistogramConfig)
}

func (o descOpt) applyFloat64Gauge(c Float64GaugeConfig) Float64GaugeConfig {
	_ = "STUB: not implemented"
	return *new(Float64GaugeConfig)
}

func (o descOpt) applyFloat64ObservableCounter(c Float64ObservableCounterConfig) Float64ObservableCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableCounterConfig)
}

func (o descOpt) applyFloat64ObservableUpDownCounter(
	c Float64ObservableUpDownCounterConfig,
) Float64ObservableUpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableUpDownCounterConfig)
}

func (o descOpt) applyFloat64ObservableGauge(c Float64ObservableGaugeConfig) Float64ObservableGaugeConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableGaugeConfig)
}

func (o descOpt) applyInt64Counter(c Int64CounterConfig) Int64CounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64CounterConfig)
}

func (o descOpt) applyInt64UpDownCounter(c Int64UpDownCounterConfig) Int64UpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64UpDownCounterConfig)
}

func (o descOpt) applyInt64Histogram(c Int64HistogramConfig) Int64HistogramConfig {
	_ = "STUB: not implemented"
	return *new(Int64HistogramConfig)
}

func (o descOpt) applyInt64Gauge(c Int64GaugeConfig) Int64GaugeConfig {
	_ = "STUB: not implemented"
	return *new(Int64GaugeConfig)
}

func (o descOpt) applyInt64ObservableCounter(c Int64ObservableCounterConfig) Int64ObservableCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableCounterConfig)
}

func (o descOpt) applyInt64ObservableUpDownCounter(
	c Int64ObservableUpDownCounterConfig,
) Int64ObservableUpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableUpDownCounterConfig)
}

func (o descOpt) applyInt64ObservableGauge(c Int64ObservableGaugeConfig) Int64ObservableGaugeConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableGaugeConfig)
}

func WithDescription(desc string) InstrumentOption {
	_ = "STUB: not implemented"
	return *new(InstrumentOption)
}

type unitOpt string

func (o unitOpt) applyFloat64Counter(c Float64CounterConfig) Float64CounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64CounterConfig)
}

func (o unitOpt) applyFloat64UpDownCounter(c Float64UpDownCounterConfig) Float64UpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64UpDownCounterConfig)
}

func (o unitOpt) applyFloat64Histogram(c Float64HistogramConfig) Float64HistogramConfig {
	_ = "STUB: not implemented"
	return *new(Float64HistogramConfig)
}

func (o unitOpt) applyFloat64Gauge(c Float64GaugeConfig) Float64GaugeConfig {
	_ = "STUB: not implemented"
	return *new(Float64GaugeConfig)
}

func (o unitOpt) applyFloat64ObservableCounter(c Float64ObservableCounterConfig) Float64ObservableCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableCounterConfig)
}

func (o unitOpt) applyFloat64ObservableUpDownCounter(
	c Float64ObservableUpDownCounterConfig,
) Float64ObservableUpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableUpDownCounterConfig)
}

func (o unitOpt) applyFloat64ObservableGauge(c Float64ObservableGaugeConfig) Float64ObservableGaugeConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableGaugeConfig)
}

func (o unitOpt) applyInt64Counter(c Int64CounterConfig) Int64CounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64CounterConfig)
}

func (o unitOpt) applyInt64UpDownCounter(c Int64UpDownCounterConfig) Int64UpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64UpDownCounterConfig)
}

func (o unitOpt) applyInt64Histogram(c Int64HistogramConfig) Int64HistogramConfig {
	_ = "STUB: not implemented"
	return *new(Int64HistogramConfig)
}

func (o unitOpt) applyInt64Gauge(c Int64GaugeConfig) Int64GaugeConfig {
	_ = "STUB: not implemented"
	return *new(Int64GaugeConfig)
}

func (o unitOpt) applyInt64ObservableCounter(c Int64ObservableCounterConfig) Int64ObservableCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableCounterConfig)
}

func (o unitOpt) applyInt64ObservableUpDownCounter(
	c Int64ObservableUpDownCounterConfig,
) Int64ObservableUpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableUpDownCounterConfig)
}

func (o unitOpt) applyInt64ObservableGauge(c Int64ObservableGaugeConfig) Int64ObservableGaugeConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableGaugeConfig)
}

func WithUnit(u string) InstrumentOption { _ = "STUB: not implemented"; return *new(InstrumentOption) }

func WithExplicitBucketBoundaries(bounds ...float64) HistogramOption {
	_ = "STUB: not implemented"
	return *new(HistogramOption)
}

type bucketOpt []float64

func (o bucketOpt) applyFloat64Histogram(c Float64HistogramConfig) Float64HistogramConfig {
	_ = "STUB: not implemented"
	return *new(Float64HistogramConfig)
}

func (o bucketOpt) applyInt64Histogram(c Int64HistogramConfig) Int64HistogramConfig {
	_ = "STUB: not implemented"
	return *new(Int64HistogramConfig)
}

type AddOption interface {
	applyAdd(AddConfig) AddConfig
}

type AddConfig struct {
	attrs attribute.Set
}

func NewAddConfig(opts []AddOption) AddConfig { _ = "STUB: not implemented"; return *new(AddConfig) }

func (c AddConfig) Attributes() attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

type RecordOption interface {
	applyRecord(RecordConfig) RecordConfig
}

type RecordConfig struct {
	attrs attribute.Set
}

func NewRecordConfig(opts []RecordOption) RecordConfig {
	_ = "STUB: not implemented"
	return *new(RecordConfig)
}

func (c RecordConfig) Attributes() attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

type ObserveOption interface {
	applyObserve(ObserveConfig) ObserveConfig
}

type ObserveConfig struct {
	attrs attribute.Set
}

func NewObserveConfig(opts []ObserveOption) ObserveConfig {
	_ = "STUB: not implemented"
	return *new(ObserveConfig)
}

func (c ObserveConfig) Attributes() attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

type MeasurementOption interface {
	AddOption
	RecordOption
	ObserveOption
}

type attrOpt struct {
	set attribute.Set
}

func (o *attrOpt) Set(set attribute.Set) { _ = "STUB: not implemented"; return }

func mergeSets(a, b attribute.Set) attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

func (o *attrOpt) applyAdd(c AddConfig) AddConfig {
	_ = "STUB: not implemented"
	return *new(AddConfig)
}

func (o *attrOpt) applyRecord(c RecordConfig) RecordConfig {
	_ = "STUB: not implemented"
	return *new(RecordConfig)
}

func (o *attrOpt) applyObserve(c ObserveConfig) ObserveConfig {
	_ = "STUB: not implemented"
	return *new(ObserveConfig)
}

func WithAttributeSet(attributes attribute.Set) MeasurementOption {
	_ = "STUB: not implemented"
	return *new(MeasurementOption)
}

func WithAttributes(attributes ...attribute.KeyValue) MeasurementOption {
	_ = "STUB: not implemented"
	return *new(MeasurementOption)
}
