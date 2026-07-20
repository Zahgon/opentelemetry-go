package metric

import (
	"context"

	"go.opentelemetry.io/otel/metric/embedded"
)

type Float64Observable interface {
	Observable

	float64Observable()
}

type Float64ObservableCounter interface {
	embedded.Float64ObservableCounter

	Float64Observable
}

type Float64ObservableCounterConfig struct {
	description string
	unit        string
	callbacks   []Float64Callback
}

func NewFloat64ObservableCounterConfig(opts ...Float64ObservableCounterOption) Float64ObservableCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableCounterConfig)
}

func (c Float64ObservableCounterConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Float64ObservableCounterConfig) Unit() string { _ = "STUB: not implemented"; return "" }

func (c Float64ObservableCounterConfig) Callbacks() []Float64Callback {
	_ = "STUB: not implemented"
	return nil
}

type Float64ObservableCounterOption interface {
	applyFloat64ObservableCounter(Float64ObservableCounterConfig) Float64ObservableCounterConfig
}

type Float64ObservableUpDownCounter interface {
	embedded.Float64ObservableUpDownCounter

	Float64Observable
}

type Float64ObservableUpDownCounterConfig struct {
	description string
	unit        string
	callbacks   []Float64Callback
}

func NewFloat64ObservableUpDownCounterConfig(
	opts ...Float64ObservableUpDownCounterOption,
) Float64ObservableUpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableUpDownCounterConfig)
}

func (c Float64ObservableUpDownCounterConfig) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (c Float64ObservableUpDownCounterConfig) Unit() string { _ = "STUB: not implemented"; return "" }

func (c Float64ObservableUpDownCounterConfig) Callbacks() []Float64Callback {
	_ = "STUB: not implemented"
	return nil
}

type Float64ObservableUpDownCounterOption interface {
	applyFloat64ObservableUpDownCounter(Float64ObservableUpDownCounterConfig) Float64ObservableUpDownCounterConfig
}

type Float64ObservableGauge interface {
	embedded.Float64ObservableGauge

	Float64Observable
}

type Float64ObservableGaugeConfig struct {
	description string
	unit        string
	callbacks   []Float64Callback
}

func NewFloat64ObservableGaugeConfig(opts ...Float64ObservableGaugeOption) Float64ObservableGaugeConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableGaugeConfig)
}

func (c Float64ObservableGaugeConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Float64ObservableGaugeConfig) Unit() string { _ = "STUB: not implemented"; return "" }

func (c Float64ObservableGaugeConfig) Callbacks() []Float64Callback {
	_ = "STUB: not implemented"
	return nil
}

type Float64ObservableGaugeOption interface {
	applyFloat64ObservableGauge(Float64ObservableGaugeConfig) Float64ObservableGaugeConfig
}

type Float64Observer interface {
	embedded.Float64Observer

	Observe(value float64, options ...ObserveOption)
}

type Float64Callback func(context.Context, Float64Observer) error

type Float64ObservableOption interface {
	Float64ObservableCounterOption
	Float64ObservableUpDownCounterOption
	Float64ObservableGaugeOption
}

type float64CallbackOpt struct {
	cback Float64Callback
}

func (o float64CallbackOpt) applyFloat64ObservableCounter(
	cfg Float64ObservableCounterConfig,
) Float64ObservableCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableCounterConfig)
}

func (o float64CallbackOpt) applyFloat64ObservableUpDownCounter(
	cfg Float64ObservableUpDownCounterConfig,
) Float64ObservableUpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableUpDownCounterConfig)
}

func (o float64CallbackOpt) applyFloat64ObservableGauge(cfg Float64ObservableGaugeConfig) Float64ObservableGaugeConfig {
	_ = "STUB: not implemented"
	return *new(Float64ObservableGaugeConfig)
}

func WithFloat64Callback(callback Float64Callback) Float64ObservableOption {
	_ = "STUB: not implemented"
	return *new(Float64ObservableOption)
}
