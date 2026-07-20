package otlptracehttp

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp/internal/observ"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp/internal/otlpconfig"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp/internal/retry"
)

const contentTypeProto = "application/x-protobuf"

var maxResponseBodySize int64 = 4 * 1024 * 1024

var gzPool = sync.Pool{
	New: func() any {
		w := gzip.NewWriter(io.Discard)
		return w
	},
}

var ourTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var errInsecureEndpointWithTLS = errors.New("insecure HTTP endpoint cannot use TLS client configuration")

type client struct {
	name        string
	cfg         otlpconfig.SignalConfig
	generalCfg  otlpconfig.Config
	requestFunc retry.RequestFunc
	client      *http.Client
	stopCh      chan struct{}
	stopOnce    sync.Once

	instID int64
	inst   *observ.Instrumentation
}

var _ otlptrace.Client = (*client)(nil)

func NewClient(opts ...Option) otlptrace.Client {
	_ = "STUB: not implemented"
	return *new(otlptrace.Client)
}

func (c *client) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *client) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *client) UploadTraces(ctx context.Context, protoSpans []*tracepb.ResourceSpans) (uploadErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) newRequest(body []byte) (request, error) {
	_ = "STUB: not implemented"
	return *new(request), nil
}

func (*client) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }

func bodyReader(buf []byte) func() io.ReadCloser { _ = "STUB: not implemented"; return nil }

func bodyReaderErr(buf []byte) func() (io.ReadCloser, error) { _ = "STUB: not implemented"; return nil }

type request struct {
	*http.Request

	bodyReader func() io.ReadCloser
}

func (r *request) reset(ctx context.Context) { _ = "STUB: not implemented"; return }

type retryableError struct {
	throttle time.Duration
	err      error
}

func newResponseError(header http.Header, wrapped error) error {
	_ = "STUB: not implemented"
	return nil
}

func retryAfterDuration(v string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (e retryableError) Error() string { _ = "STUB: not implemented"; return "" }

func (e retryableError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e retryableError) As(target any) bool { _ = "STUB: not implemented"; return false }

func evaluate(err error) (bool, time.Duration) {
	_ = "STUB: not implemented"
	return false, *new(time.Duration)
}

//nolint:errorlint

func (c *client) getScheme() string { _ = "STUB: not implemented"; return "" }

func (c *client) contextWithStop(ctx context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}
