package noop

import (
	"context"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
)

var (
	_ metric.MeterProvider                  = MeterProvider{}
	_ metric.Meter                          = Meter{}
	_ metric.Observer                       = Observer{}
	_ metric.Registration                   = Registration{}
	_ metric.Int64Counter                   = Int64Counter{}
	_ metric.Float64Counter                 = Float64Counter{}
	_ metric.Int64UpDownCounter             = Int64UpDownCounter{}
	_ metric.Float64UpDownCounter           = Float64UpDownCounter{}
	_ metric.Int64Histogram                 = Int64Histogram{}
	_ metric.Float64Histogram               = Float64Histogram{}
	_ metric.Int64Gauge                     = Int64Gauge{}
	_ metric.Float64Gauge                   = Float64Gauge{}
	_ metric.Int64ObservableCounter         = Int64ObservableCounter{}
	_ metric.Float64ObservableCounter       = Float64ObservableCounter{}
	_ metric.Int64ObservableGauge           = Int64ObservableGauge{}
	_ metric.Float64ObservableGauge         = Float64ObservableGauge{}
	_ metric.Int64ObservableUpDownCounter   = Int64ObservableUpDownCounter{}
	_ metric.Float64ObservableUpDownCounter = Float64ObservableUpDownCounter{}
	_ metric.Int64Observer                  = Int64Observer{}
	_ metric.Float64Observer                = Float64Observer{}
)

type MeterProvider struct{ embedded.MeterProvider }

func NewMeterProvider() MeterProvider { _ = "STUB: not implemented"; return *new(MeterProvider) }

func (MeterProvider) Meter(string, ...metric.MeterOption) metric.Meter {
	_ = "STUB: not implemented"
	return *new(metric.Meter)
}

type Meter struct{ embedded.Meter }

func (Meter) Int64Counter(string, ...metric.Int64CounterOption) (metric.Int64Counter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter), nil
}

func (Meter) Int64UpDownCounter(string, ...metric.Int64UpDownCounterOption) (metric.Int64UpDownCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter), nil
}

func (Meter) Int64Histogram(string, ...metric.Int64HistogramOption) (metric.Int64Histogram, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram), nil
}

func (Meter) Int64Gauge(string, ...metric.Int64GaugeOption) (metric.Int64Gauge, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge), nil
}

func (Meter) Int64ObservableCounter(
	string,
	...metric.Int64ObservableCounterOption,
) (metric.Int64ObservableCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter), nil
}

func (Meter) Int64ObservableUpDownCounter(
	string,
	...metric.Int64ObservableUpDownCounterOption,
) (metric.Int64ObservableUpDownCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter), nil
}

func (Meter) Int64ObservableGauge(string, ...metric.Int64ObservableGaugeOption) (metric.Int64ObservableGauge, error) {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge), nil
}

func (Meter) Float64Counter(string, ...metric.Float64CounterOption) (metric.Float64Counter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64Counter), nil
}

func (Meter) Float64UpDownCounter(string, ...metric.Float64UpDownCounterOption) (metric.Float64UpDownCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64UpDownCounter), nil
}

func (Meter) Float64Histogram(string, ...metric.Float64HistogramOption) (metric.Float64Histogram, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram), nil
}

func (Meter) Float64Gauge(string, ...metric.Float64GaugeOption) (metric.Float64Gauge, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge), nil
}

func (Meter) Float64ObservableCounter(
	string,
	...metric.Float64ObservableCounterOption,
) (metric.Float64ObservableCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter), nil
}

func (Meter) Float64ObservableUpDownCounter(
	string,
	...metric.Float64ObservableUpDownCounterOption,
) (metric.Float64ObservableUpDownCounter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableUpDownCounter), nil
}

func (Meter) Float64ObservableGauge(
	string,
	...metric.Float64ObservableGaugeOption,
) (metric.Float64ObservableGauge, error) {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge), nil
}

func (Meter) RegisterCallback(metric.Callback, ...metric.Observable) (metric.Registration, error) {
	_ = "STUB: not implemented"
	return *new(metric.Registration), nil
}

type Observer struct{ embedded.Observer }

func (Observer) ObserveFloat64(metric.Float64Observable, float64, ...metric.ObserveOption) {
	_ = "STUB: not implemented"
	return
}

func (Observer) ObserveInt64(metric.Int64Observable, int64, ...metric.ObserveOption) {
	_ = "STUB: not implemented"
	return
}

type Registration struct{ embedded.Registration }

func (Registration) Unregister() error { _ = "STUB: not implemented"; return nil }

type Int64Counter struct{ embedded.Int64Counter }

func (Int64Counter) Add(context.Context, int64, ...metric.AddOption) {
	_ = "STUB: not implemented"
	return
}

func (Int64Counter) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

type Float64Counter struct{ embedded.Float64Counter }

func (Float64Counter) Add(context.Context, float64, ...metric.AddOption) {
	_ = "STUB: not implemented"
	return
}

func (Float64Counter) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

type Int64UpDownCounter struct{ embedded.Int64UpDownCounter }

func (Int64UpDownCounter) Add(context.Context, int64, ...metric.AddOption) {
	_ = "STUB: not implemented"
	return
}

func (Int64UpDownCounter) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

type Float64UpDownCounter struct{ embedded.Float64UpDownCounter }

func (Float64UpDownCounter) Add(context.Context, float64, ...metric.AddOption) {
	_ = "STUB: not implemented"
	return
}

func (Float64UpDownCounter) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

type Int64Histogram struct{ embedded.Int64Histogram }

func (Int64Histogram) Record(context.Context, int64, ...metric.RecordOption) {
	_ = "STUB: not implemented"
	return
}

func (Int64Histogram) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

type Float64Histogram struct{ embedded.Float64Histogram }

func (Float64Histogram) Record(context.Context, float64, ...metric.RecordOption) {
	_ = "STUB: not implemented"
	return
}

func (Float64Histogram) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

type Int64Gauge struct{ embedded.Int64Gauge }

func (Int64Gauge) Record(context.Context, int64, ...metric.RecordOption) {
	_ = "STUB: not implemented"
	return
}

func (Int64Gauge) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

type Float64Gauge struct{ embedded.Float64Gauge }

func (Float64Gauge) Record(context.Context, float64, ...metric.RecordOption) {
	_ = "STUB: not implemented"
	return
}

func (Float64Gauge) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

type Int64ObservableCounter struct {
	metric.Int64Observable
	embedded.Int64ObservableCounter
}

type Float64ObservableCounter struct {
	metric.Float64Observable
	embedded.Float64ObservableCounter
}

type Int64ObservableGauge struct {
	metric.Int64Observable
	embedded.Int64ObservableGauge
}

type Float64ObservableGauge struct {
	metric.Float64Observable
	embedded.Float64ObservableGauge
}

type Int64ObservableUpDownCounter struct {
	metric.Int64Observable
	embedded.Int64ObservableUpDownCounter
}

type Float64ObservableUpDownCounter struct {
	metric.Float64Observable
	embedded.Float64ObservableUpDownCounter
}

type Int64Observer struct{ embedded.Int64Observer }

func (Int64Observer) Observe(int64, ...metric.ObserveOption) { _ = "STUB: not implemented"; return }

type Float64Observer struct{ embedded.Float64Observer }

func (Float64Observer) Observe(float64, ...metric.ObserveOption) { _ = "STUB: not implemented"; return }
