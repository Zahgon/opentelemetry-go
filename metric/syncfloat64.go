package metric

import (
	"context"

	"go.opentelemetry.io/otel/metric/embedded"
)

type Float64Counter interface {
	embedded.Float64Counter

	Add(ctx context.Context, incr float64, options ...AddOption)

	Enabled(context.Context) bool
}

type Float64CounterConfig struct {
	description string
	unit        string
}

func NewFloat64CounterConfig(opts ...Float64CounterOption) Float64CounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64CounterConfig)
}

func (c Float64CounterConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Float64CounterConfig) Unit() string { _ = "STUB: not implemented"; return "" }

type Float64CounterOption interface {
	applyFloat64Counter(Float64CounterConfig) Float64CounterConfig
}

type Float64UpDownCounter interface {
	embedded.Float64UpDownCounter

	Add(ctx context.Context, incr float64, options ...AddOption)

	Enabled(context.Context) bool
}

type Float64UpDownCounterConfig struct {
	description string
	unit        string
}

func NewFloat64UpDownCounterConfig(opts ...Float64UpDownCounterOption) Float64UpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64UpDownCounterConfig)
}

func (c Float64UpDownCounterConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Float64UpDownCounterConfig) Unit() string { _ = "STUB: not implemented"; return "" }

type Float64UpDownCounterOption interface {
	applyFloat64UpDownCounter(Float64UpDownCounterConfig) Float64UpDownCounterConfig
}

type Float64Histogram interface {
	embedded.Float64Histogram

	Record(ctx context.Context, incr float64, options ...RecordOption)

	Enabled(context.Context) bool
}

type Float64HistogramConfig struct {
	description              string
	unit                     string
	explicitBucketBoundaries []float64
}

func NewFloat64HistogramConfig(opts ...Float64HistogramOption) Float64HistogramConfig {
	_ = "STUB: not implemented"
	return *new(Float64HistogramConfig)
}

func (c Float64HistogramConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Float64HistogramConfig) Unit() string { _ = "STUB: not implemented"; return "" }

func (c Float64HistogramConfig) ExplicitBucketBoundaries() []float64 {
	_ = "STUB: not implemented"
	return nil
}

type Float64HistogramOption interface {
	applyFloat64Histogram(Float64HistogramConfig) Float64HistogramConfig
}

type Float64Gauge interface {
	embedded.Float64Gauge

	Record(ctx context.Context, value float64, options ...RecordOption)

	Enabled(context.Context) bool
}

type Float64GaugeConfig struct {
	description string
	unit        string
}

func NewFloat64GaugeConfig(opts ...Float64GaugeOption) Float64GaugeConfig {
	_ = "STUB: not implemented"
	return *new(Float64GaugeConfig)
}

func (c Float64GaugeConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Float64GaugeConfig) Unit() string { _ = "STUB: not implemented"; return "" }

type Float64GaugeOption interface {
	applyFloat64Gauge(Float64GaugeConfig) Float64GaugeConfig
}
