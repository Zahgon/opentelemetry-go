package metric

import (
	"go.opentelemetry.io/otel/attribute"
)

type MeterConfig struct {
	instrumentationVersion string
	schemaURL              string
	attrs                  attribute.Set

	noCmp [0]func() //nolint: unused  // This is indeed used.
}

func (cfg MeterConfig) InstrumentationVersion() string { _ = "STUB: not implemented"; return "" }

func (cfg MeterConfig) InstrumentationAttributes() attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

func (cfg MeterConfig) SchemaURL() string { _ = "STUB: not implemented"; return "" }

type MeterOption interface {
	applyMeter(MeterConfig) MeterConfig
}

type experimentalOption interface {
	Experimental()
}

func NewMeterConfig(opts ...MeterOption) MeterConfig {
	_ = "STUB: not implemented"
	return *new(MeterConfig)
}

type meterOptionFunc func(MeterConfig) MeterConfig

func (fn meterOptionFunc) applyMeter(cfg MeterConfig) MeterConfig {
	_ = "STUB: not implemented"
	return *new(MeterConfig)
}

func WithInstrumentationVersion(version string) MeterOption {
	_ = "STUB: not implemented"
	return *new(MeterOption)
}

func WithInstrumentationAttributes(attr ...attribute.KeyValue) MeterOption {
	_ = "STUB: not implemented"
	return *new(MeterOption)
}

func WithInstrumentationAttributeSet(set attribute.Set) MeterOption {
	_ = "STUB: not implemented"
	return *new(MeterOption)
}

func WithSchemaURL(schemaURL string) MeterOption {
	_ = "STUB: not implemented"
	return *new(MeterOption)
}
