package log

import (
	"context"

	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/sdk/instrumentation"
)

type Processor interface {
	Enabled(ctx context.Context, param EnabledParameters) bool

	OnEmit(ctx context.Context, record *Record) error

	Shutdown(ctx context.Context) error

	ForceFlush(ctx context.Context) error
}

type EnabledParameters struct {
	InstrumentationScope instrumentation.Scope
	Severity             log.Severity
	EventName            string
}
