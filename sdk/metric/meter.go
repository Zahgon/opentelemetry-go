package metric

import (
	"errors"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric/internal/aggregate"
)

var ErrInstrumentName = errors.New("invalid instrument name")

type meter struct {
	embedded.Meter

	scope instrumentation.Scope
	pipes pipelines

	int64Insts             *cacheWithErr[instID, *int64Inst]
	float64Insts           *cacheWithErr[instID, *float64Inst]
	int64ObservableInsts   *cacheWithErr[instID, int64Observable]
	float64ObservableInsts *cacheWithErr[instID, float64Observable]

	int64Resolver   resolver[int64]
	float64Resolver resolver[float64]
}

func newMeter(s instrumentation.Scope, p pipelines) *meter { _ = "STUB: not implemented"; return nil }

var _ metric.Meter = (*meter)(nil)

func (m *meter) Int64Counter(name string, options ...metric.Int64CounterOption) (metric.Int64Counter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter), nil
}

func (m *meter) Int64UpDownCounter(
	name string,
	options ...metric.Int64UpDownCounterOption,
) (metric.Int64UpDownCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter), nil
}

func (m *meter) Int64Histogram(name string, options ...metric.Int64HistogramOption) (metric.Int64Histogram, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram), nil
}

func (m *meter) Int64Gauge(name string, options ...metric.Int64GaugeOption) (metric.Int64Gauge, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge), nil
}

func (m *meter) int64ObservableInstrument(
	id Instrument,
	allowedKeys []attribute.Key,
	callbacks []metric.Int64Callback,
) (int64Observable, error) {
	_ = "STUB: not implemented"
	return *new(int64Observable), nil
}

func (m *meter) Int64ObservableCounter(
	name string,
	options ...metric.Int64ObservableCounterOption,
) (metric.Int64ObservableCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter), nil
}

func (m *meter) Int64ObservableUpDownCounter(
	name string,
	options ...metric.Int64ObservableUpDownCounterOption,
) (metric.Int64ObservableUpDownCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter), nil
}

func (m *meter) Int64ObservableGauge(
	name string,
	options ...metric.Int64ObservableGaugeOption,
) (metric.Int64ObservableGauge, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge), nil
}

func (m *meter) Float64Counter(name string, options ...metric.Float64CounterOption) (metric.Float64Counter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64Counter), nil
}

func (m *meter) Float64UpDownCounter(
	name string,
	options ...metric.Float64UpDownCounterOption,
) (metric.Float64UpDownCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64UpDownCounter), nil
}

func (m *meter) Float64Histogram(
	name string,
	options ...metric.Float64HistogramOption,
) (metric.Float64Histogram, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram), nil
}

func (m *meter) Float64Gauge(name string, options ...metric.Float64GaugeOption) (metric.Float64Gauge, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge), nil
}

func (m *meter) float64ObservableInstrument(
	id Instrument,
	allowedKeys []attribute.Key,
	callbacks []metric.Float64Callback,
) (float64Observable, error) {
	_ = "STUB: not implemented"
	return *new(float64Observable), nil
}

func (m *meter) Float64ObservableCounter(
	name string,
	options ...metric.Float64ObservableCounterOption,
) (metric.Float64ObservableCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter), nil
}

func (m *meter) Float64ObservableUpDownCounter(
	name string,
	options ...metric.Float64ObservableUpDownCounterOption,
) (metric.Float64ObservableUpDownCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableUpDownCounter), nil
}

func (m *meter) Float64ObservableGauge(
	name string,
	options ...metric.Float64ObservableGaugeOption,
) (metric.Float64ObservableGauge, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge), nil
}

func validateInstrumentName(name string) error { _ = "STUB: not implemented"; return nil }

func isAlpha(c rune) bool { _ = "STUB: not implemented"; return false }

func isAlphanumeric(c rune) bool { _ = "STUB: not implemented"; return false }

func warnRepeatedObservableCallbacks(id Instrument) { _ = "STUB: not implemented"; return }

func (m *meter) RegisterCallback(f metric.Callback, insts ...metric.Observable) (metric.Registration, error) {
	_ = "STUB: not implemented"
	return *new(metric.Registration), nil
}

type observer struct {
	embedded.Observer

	pipe    *pipeline
	float64 map[observableID[float64]]struct{}
	int64   map[observableID[int64]]struct{}
}

func newObserver(p *pipeline) observer { _ = "STUB: not implemented"; return *new(observer) }

func (r observer) registerFloat64(id observableID[float64]) { _ = "STUB: not implemented"; return }

func (r observer) registerInt64(id observableID[int64]) { _ = "STUB: not implemented"; return }

var (
	errUnknownObserver = errors.New("unknown observable instrument")
	errUnregObserver   = errors.New("observable instrument not registered for callback")
)

func (r observer) ObserveFloat64(o metric.Float64Observable, v float64, opts ...metric.ObserveOption) {
	_ = "STUB: not implemented"
	return
}

func (r observer) ObserveInt64(o metric.Int64Observable, v int64, opts ...metric.ObserveOption) {
	_ = "STUB: not implemented"
	return
}

type noopRegister struct{ embedded.Registration }

func (noopRegister) Unregister() error { _ = "STUB: not implemented"; return nil }

type int64InstProvider struct{ *meter }

func (p int64InstProvider) aggs(
	kind InstrumentKind,
	name, desc, u string,
	allowedKeys []attribute.Key,
) ([]aggregate.Measure[int64], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p int64InstProvider) histogramAggs(
	name string,
	cfg metric.Int64HistogramConfig,
	allowedKeys []attribute.Key,
) ([]aggregate.Measure[int64], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p int64InstProvider) lookup(
	kind InstrumentKind,
	name, desc, u string,
	allowedKeys []attribute.Key,
) (*int64Inst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p int64InstProvider) lookupHistogram(
	name string,
	cfg metric.Int64HistogramConfig,
	allowedKeys []attribute.Key,
) (*int64Inst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type float64InstProvider struct{ *meter }

func (p float64InstProvider) aggs(
	kind InstrumentKind,
	name, desc, u string,
	allowedKeys []attribute.Key,
) ([]aggregate.Measure[float64], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p float64InstProvider) histogramAggs(
	name string,
	cfg metric.Float64HistogramConfig,
	allowedKeys []attribute.Key,
) ([]aggregate.Measure[float64], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p float64InstProvider) lookup(
	kind InstrumentKind,
	name, desc, u string,
	allowedKeys []attribute.Key,
) (*float64Inst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p float64InstProvider) lookupHistogram(
	name string,
	cfg metric.Float64HistogramConfig,
	allowedKeys []attribute.Key,
) (*float64Inst, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type int64Observer struct {
	embedded.Int64Observer
	measures[int64]
}

func (o int64Observer) Observe(val int64, opts ...metric.ObserveOption) {
	_ = "STUB: not implemented"
	return
}

type float64Observer struct {
	embedded.Float64Observer
	measures[float64]
}

func (o float64Observer) Observe(val float64, opts ...metric.ObserveOption) {
	_ = "STUB: not implemented"
	return
}

func defaultAttributes[T any](opts []T) []attribute.Key { _ = "STUB: not implemented"; return nil }
