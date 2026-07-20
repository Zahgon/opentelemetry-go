package log

import (
	"context"
	"sync"
	"sync/atomic"

	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/embedded"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
)

const (
	defaultAttrCntLim    = 128
	defaultAttrValLenLim = -1

	envarAttrCntLim    = "OTEL_LOGRECORD_ATTRIBUTE_COUNT_LIMIT"
	envarAttrValLenLim = "OTEL_LOGRECORD_ATTRIBUTE_VALUE_LENGTH_LIMIT"
)

type providerConfig struct {
	resource      *resource.Resource
	processors    []Processor
	attrCntLim    setting[int]
	attrValLenLim setting[int]
	allowDupKeys  setting[bool]
}

type experimentalOption interface {
	Experimental()
}

func newProviderConfig(opts []LoggerProviderOption) providerConfig {
	_ = "STUB: not implemented"
	return *new(providerConfig)
}

type LoggerProvider struct {
	embedded.LoggerProvider

	resource                  *resource.Resource
	processors                []Processor
	attributeCountLimit       int
	attributeValueLengthLimit int
	allowDupKeys              bool

	loggersMu sync.Mutex
	loggers   map[instrumentation.Scope]*logger

	stopped atomic.Bool

	noCmp [0]func() //nolint: unused  // This is indeed used.
}

var _ log.LoggerProvider = (*LoggerProvider)(nil)

func NewLoggerProvider(opts ...LoggerProviderOption) *LoggerProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p *LoggerProvider) Logger(name string, opts ...log.LoggerOption) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

func (p *LoggerProvider) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *LoggerProvider) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type LoggerProviderOption interface {
	apply(providerConfig) providerConfig
}

type loggerProviderOptionFunc func(providerConfig) providerConfig

func (fn loggerProviderOptionFunc) apply(c providerConfig) providerConfig {
	_ = "STUB: not implemented"
	return *new(providerConfig)
}

func WithResource(res *resource.Resource) LoggerProviderOption {
	_ = "STUB: not implemented"
	return *new(LoggerProviderOption)
}

func WithProcessor(processor Processor) LoggerProviderOption {
	_ = "STUB: not implemented"
	return *new(LoggerProviderOption)
}

func WithAttributeCountLimit(limit int) LoggerProviderOption {
	_ = "STUB: not implemented"
	return *new(LoggerProviderOption)
}

func WithAttributeValueLengthLimit(limit int) LoggerProviderOption {
	_ = "STUB: not implemented"
	return *new(LoggerProviderOption)
}

func WithAllowKeyDuplication() LoggerProviderOption {
	_ = "STUB: not implemented"
	return *new(LoggerProviderOption)
}
