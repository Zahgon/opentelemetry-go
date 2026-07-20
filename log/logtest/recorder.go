package logtest

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/embedded"
)

type enabledFn func(context.Context, log.EnabledParameters) bool

var defaultEnabledFunc = func(context.Context, log.EnabledParameters) bool {
	return true
}

type config struct {
	enabledFn enabledFn
}

func newConfig(options []Option) config { _ = "STUB: not implemented"; return *new(config) }

type Option interface {
	apply(config) config
}

type optFunc func(config) config

func (f optFunc) apply(c config) config { _ = "STUB: not implemented"; return *new(config) }

func WithEnabledFunc(fn func(context.Context, log.EnabledParameters) bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NewRecorder(options ...Option) *Recorder { _ = "STUB: not implemented"; return nil }

type Recording map[Scope][]Record

type Scope struct {
	Name string

	Version string

	SchemaURL string

	Attributes attribute.Set
}

type Record struct {
	_ [0]func()

	Context           context.Context
	EventName         string
	Timestamp         time.Time
	ObservedTimestamp time.Time
	Severity          log.Severity
	SeverityText      string
	Body              attribute.Value
	Error             error
	Attributes        []attribute.KeyValue
}

type Recorder struct {
	_ [0]func()

	embedded.LoggerProvider

	mu      sync.Mutex
	loggers map[Scope]*logger

	enabledFn enabledFn
}

var _ log.LoggerProvider = (*Recorder)(nil)

func (a Record) Clone() Record { _ = "STUB: not implemented"; return *new(Record) }

func (r *Recorder) Logger(name string, opts ...log.LoggerOption) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

func (r *Recorder) Reset() { _ = "STUB: not implemented"; return }

func (r *Recorder) Result() Recording { _ = "STUB: not implemented"; return *new(Recording) }

type logger struct {
	embedded.Logger

	mu      sync.Mutex
	records []*Record

	enabledFn enabledFn
}

func (l *logger) Enabled(ctx context.Context, param log.EnabledParameters) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *logger) Emit(ctx context.Context, record log.Record) { _ = "STUB: not implemented"; return }

func (l *logger) Reset() { _ = "STUB: not implemented"; return }
