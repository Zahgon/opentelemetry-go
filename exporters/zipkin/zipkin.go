package zipkin

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go-logr/logr"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	defaultCollectorURL = "http://localhost:9411/api/v2/spans"
)

type Exporter struct {
	url     string
	client  *http.Client
	logger  logr.Logger
	headers map[string]string

	stoppedMu sync.RWMutex
	stopped   bool
}

var _ sdktrace.SpanExporter = &Exporter{}

var emptyLogger = logr.Logger{}

type config struct {
	client  *http.Client
	logger  logr.Logger
	headers map[string]string
}

type Option interface {
	apply(config) config
}

type optionFunc func(config) config

func (fn optionFunc) apply(cfg config) config { _ = "STUB: not implemented"; return *new(config) }

func WithLogger(logger *log.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLogr(logger logr.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeaders(headers map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(collectorURL string, opts ...Option) (*Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Exporter) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Exporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *Exporter) logf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (e *Exporter) errf(format string, args ...any) error { _ = "STUB: not implemented"; return nil }

func (*Exporter) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }
