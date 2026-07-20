package global

import (
	"go.opentelemetry.io/otel/log"
)

func Logger(name string, options ...log.LoggerOption) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

func GetLoggerProvider() log.LoggerProvider {
	_ = "STUB: not implemented"
	return *new(log.LoggerProvider)
}

func SetLoggerProvider(provider log.LoggerProvider) { _ = "STUB: not implemented"; return }
