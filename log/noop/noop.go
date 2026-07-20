package noop

import (
	"context"

	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/embedded"
)

var (
	_ log.LoggerProvider = LoggerProvider{}
	_ log.Logger         = Logger{}
)

type LoggerProvider struct{ embedded.LoggerProvider }

func NewLoggerProvider() LoggerProvider { _ = "STUB: not implemented"; return *new(LoggerProvider) }

func (LoggerProvider) Logger(string, ...log.LoggerOption) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

type Logger struct{ embedded.Logger }

func (Logger) Emit(context.Context, log.Record) { _ = "STUB: not implemented"; return }

func (Logger) Enabled(context.Context, log.EnabledParameters) bool {
	_ = "STUB: not implemented"
	return false
}
