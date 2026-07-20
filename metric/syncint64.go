package metric

import (
	"context"

	"go.opentelemetry.io/otel/metric/embedded"
)

type Int64Counter interface {
	embedded.Int64Counter

	Add(ctx context.Context, incr int64, options ...AddOption)

	Enabled(context.Context) bool
}

type Int64CounterConfig struct {
	description string
	unit        string
}

func NewInt64CounterConfig(opts ...Int64CounterOption) Int64CounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64CounterConfig)
}

func (c Int64CounterConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Int64CounterConfig) Unit() string { _ = "STUB: not implemented"; return "" }

type Int64CounterOption interface {
	applyInt64Counter(Int64CounterConfig) Int64CounterConfig
}

type Int64UpDownCounter interface {
	embedded.Int64UpDownCounter

	Add(ctx context.Context, incr int64, options ...AddOption)

	Enabled(context.Context) bool
}

type Int64UpDownCounterConfig struct {
	description string
	unit        string
}

func NewInt64UpDownCounterConfig(opts ...Int64UpDownCounterOption) Int64UpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64UpDownCounterConfig)
}

func (c Int64UpDownCounterConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Int64UpDownCounterConfig) Unit() string { _ = "STUB: not implemented"; return "" }

type Int64UpDownCounterOption interface {
	applyInt64UpDownCounter(Int64UpDownCounterConfig) Int64UpDownCounterConfig
}

type Int64Histogram interface {
	embedded.Int64Histogram

	Record(ctx context.Context, incr int64, options ...RecordOption)

	Enabled(context.Context) bool
}

type Int64HistogramConfig struct {
	description              string
	unit                     string
	explicitBucketBoundaries []float64
}

func NewInt64HistogramConfig(opts ...Int64HistogramOption) Int64HistogramConfig {
	_ = "STUB: not implemented"
	return *new(Int64HistogramConfig)
}

func (c Int64HistogramConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Int64HistogramConfig) Unit() string { _ = "STUB: not implemented"; return "" }

func (c Int64HistogramConfig) ExplicitBucketBoundaries() []float64 {
	_ = "STUB: not implemented"
	return nil
}

type Int64HistogramOption interface {
	applyInt64Histogram(Int64HistogramConfig) Int64HistogramConfig
}

type Int64Gauge interface {
	embedded.Int64Gauge

	Record(ctx context.Context, value int64, options ...RecordOption)

	Enabled(context.Context) bool
}

type Int64GaugeConfig struct {
	description string
	unit        string
}

func NewInt64GaugeConfig(opts ...Int64GaugeOption) Int64GaugeConfig {
	_ = "STUB: not implemented"
	return *new(Int64GaugeConfig)
}

func (c Int64GaugeConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Int64GaugeConfig) Unit() string { _ = "STUB: not implemented"; return "" }

type Int64GaugeOption interface {
	applyInt64Gauge(Int64GaugeConfig) Int64GaugeConfig
}
