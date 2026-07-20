package log

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/embedded"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	semconv "go.opentelemetry.io/otel/semconv/v1.42.0"
)

var now = time.Now

const (
	exceptionTypeKey    = semconv.ExceptionTypeKey
	exceptionMessageKey = semconv.ExceptionMessageKey
)

var _ log.Logger = (*logger)(nil)

type logger struct {
	embedded.Logger

	provider             *LoggerProvider
	instrumentationScope instrumentation.Scope

	recCntIncr func(context.Context)
}

func newLogger(p *LoggerProvider, scope instrumentation.Scope) *logger {
	_ = "STUB: not implemented"
	return nil
}

func (l *logger) Emit(ctx context.Context, r log.Record) { _ = "STUB: not implemented"; return }

func (l *logger) Enabled(ctx context.Context, param log.EnabledParameters) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *logger) newRecord(ctx context.Context, r log.Record) Record {
	_ = "STUB: not implemented"
	return *new(Record)
}

func errorType(err error) string { _ = "STUB: not implemented"; return "" }

var fmtWrapErrorType = reflect.TypeOf(fmt.Errorf("wrapped: %w", errors.New("err")))

func unwrapFmtWrapped(err error) error { _ = "STUB: not implemented"; return nil }
