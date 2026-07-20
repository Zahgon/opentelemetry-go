package metric

import (
	"context"

	"go.opentelemetry.io/otel/metric/embedded"
)

type Int64Observable interface {
	Observable

	int64Observable()
}

type Int64ObservableCounter interface {
	embedded.Int64ObservableCounter

	Int64Observable
}

type Int64ObservableCounterConfig struct {
	description string
	unit        string
	callbacks   []Int64Callback
}

func NewInt64ObservableCounterConfig(opts ...Int64ObservableCounterOption) Int64ObservableCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableCounterConfig)
}

func (c Int64ObservableCounterConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Int64ObservableCounterConfig) Unit() string { _ = "STUB: not implemented"; return "" }

func (c Int64ObservableCounterConfig) Callbacks() []Int64Callback {
	_ = "STUB: not implemented"
	return nil
}

type Int64ObservableCounterOption interface {
	applyInt64ObservableCounter(Int64ObservableCounterConfig) Int64ObservableCounterConfig
}

type Int64ObservableUpDownCounter interface {
	embedded.Int64ObservableUpDownCounter

	Int64Observable
}

type Int64ObservableUpDownCounterConfig struct {
	description string
	unit        string
	callbacks   []Int64Callback
}

func NewInt64ObservableUpDownCounterConfig(
	opts ...Int64ObservableUpDownCounterOption,
) Int64ObservableUpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableUpDownCounterConfig)
}

func (c Int64ObservableUpDownCounterConfig) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (c Int64ObservableUpDownCounterConfig) Unit() string { _ = "STUB: not implemented"; return "" }

func (c Int64ObservableUpDownCounterConfig) Callbacks() []Int64Callback {
	_ = "STUB: not implemented"
	return nil
}

type Int64ObservableUpDownCounterOption interface {
	applyInt64ObservableUpDownCounter(Int64ObservableUpDownCounterConfig) Int64ObservableUpDownCounterConfig
}

type Int64ObservableGauge interface {
	embedded.Int64ObservableGauge

	Int64Observable
}

type Int64ObservableGaugeConfig struct {
	description string
	unit        string
	callbacks   []Int64Callback
}

func NewInt64ObservableGaugeConfig(opts ...Int64ObservableGaugeOption) Int64ObservableGaugeConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableGaugeConfig)
}

func (c Int64ObservableGaugeConfig) Description() string { _ = "STUB: not implemented"; return "" }

func (c Int64ObservableGaugeConfig) Unit() string { _ = "STUB: not implemented"; return "" }

func (c Int64ObservableGaugeConfig) Callbacks() []Int64Callback {
	_ = "STUB: not implemented"
	return nil
}

type Int64ObservableGaugeOption interface {
	applyInt64ObservableGauge(Int64ObservableGaugeConfig) Int64ObservableGaugeConfig
}

type Int64Observer interface {
	embedded.Int64Observer

	Observe(value int64, options ...ObserveOption)
}

type Int64Callback func(context.Context, Int64Observer) error

type Int64ObservableOption interface {
	Int64ObservableCounterOption
	Int64ObservableUpDownCounterOption
	Int64ObservableGaugeOption
}

type int64CallbackOpt struct {
	cback Int64Callback
}

func (o int64CallbackOpt) applyInt64ObservableCounter(cfg Int64ObservableCounterConfig) Int64ObservableCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableCounterConfig)
}

func (o int64CallbackOpt) applyInt64ObservableUpDownCounter(
	cfg Int64ObservableUpDownCounterConfig,
) Int64ObservableUpDownCounterConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableUpDownCounterConfig)
}

func (o int64CallbackOpt) applyInt64ObservableGauge(cfg Int64ObservableGaugeConfig) Int64ObservableGaugeConfig {
	_ = "STUB: not implemented"
	return *new(Int64ObservableGaugeConfig)
}

func WithInt64Callback(callback Int64Callback) Int64ObservableOption {
	_ = "STUB: not implemented"
	return *new(Int64ObservableOption)
}
