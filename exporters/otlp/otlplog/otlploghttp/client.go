package otlploghttp

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	logpb "go.opentelemetry.io/proto/otlp/logs/v1"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp/internal/observ"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp/internal/retry"
)

type client struct {
	uploadLogs func(context.Context, []*logpb.ResourceLogs) error
}

func (c *client) UploadLogs(ctx context.Context, rl []*logpb.ResourceLogs) error {
	_ = "STUB: not implemented"
	return nil
}

func newNoopClient() *client { _ = "STUB: not implemented"; return nil }

var exporterN atomic.Int64

var errInsecureEndpointWithTLS = errors.New("insecure HTTP endpoint cannot use TLS client configuration")

var maxResponseBodySize int64 = 4 * 1024 * 1024

func nextExporterID() int64 { _ = "STUB: not implemented"; return 0 }

func newHTTPClient(ctx context.Context, cfg config) (*client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type httpClient struct {
	req            *http.Request
	compression    Compression
	maxRequestSize int
	requestFunc    retry.RequestFunc
	client         *http.Client

	inst *observ.Instrumentation
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

func (c *httpClient) uploadLogs(ctx context.Context, data []*logpb.ResourceLogs) (uploadErr error) {
	_ = "STUB: not implemented"
	return nil
}

var gzPool = sync.Pool{
	New: func() any {
		w := gzip.NewWriter(io.Discard)
		return w
	},
}

func (c *httpClient) newRequest(ctx context.Context, body []byte) (request, error) {
	_ = "STUB: not implemented"
	return *new(request), nil
}

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
