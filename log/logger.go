package log

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log/embedded"
)

type Logger interface {
	embedded.Logger

	Emit(ctx context.Context, record Record)

	Enabled(ctx context.Context, param EnabledParameters) bool
}

type LoggerOption interface {
	applyLogger(LoggerConfig) LoggerConfig
}

type LoggerConfig struct {
	noCmp [0]func() //nolint: unused  // This is indeed used.

	version   string
	schemaURL string
	attrs     attribute.Set
}

type experimentalOption interface {
	Experimental()
}

func NewLoggerConfig(options ...LoggerOption) LoggerConfig {
	_ = "STUB: not implemented"
	return *new(LoggerConfig)
}

func (cfg LoggerConfig) InstrumentationVersion() string { _ = "STUB: not implemented"; return "" }

func (cfg LoggerConfig) InstrumentationAttributes() attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

func (cfg LoggerConfig) SchemaURL() string { _ = "STUB: not implemented"; return "" }

type loggerOptionFunc func(LoggerConfig) LoggerConfig

func (fn loggerOptionFunc) applyLogger(cfg LoggerConfig) LoggerConfig {
	_ = "STUB: not implemented"
	return *new(LoggerConfig)
}

func WithInstrumentationVersion(version string) LoggerOption {
	_ = "STUB: not implemented"
	return *new(LoggerOption)
}

func mergeSets(a, b attribute.Set) attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

func WithInstrumentationAttributes(attr ...attribute.KeyValue) LoggerOption {
	_ = "STUB: not implemented"
	return *new(LoggerOption)
}

func WithInstrumentationAttributeSet(set attribute.Set) LoggerOption {
	_ = "STUB: not implemented"
	return *new(LoggerOption)
}

func WithSchemaURL(schemaURL string) LoggerOption {
	_ = "STUB: not implemented"
	return *new(LoggerOption)
}

type EnabledParameters struct {
	Severity  Severity
	EventName string
}
