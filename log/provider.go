package log

import "go.opentelemetry.io/otel/log/embedded"

type LoggerProvider interface {
	embedded.LoggerProvider

	Logger(name string, options ...LoggerOption) Logger
}
