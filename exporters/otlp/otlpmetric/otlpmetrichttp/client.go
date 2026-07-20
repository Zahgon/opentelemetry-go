package otlpmetrichttp

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	metricpb "go.opentelemetry.io/proto/otlp/metrics/v1"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp/internal/observ"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp/internal/oconf"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp/internal/retry"
)

type client struct {
	req            *http.Request
	compression    Compression
	maxRequestSize int
	requestFunc    retry.RequestFunc
	httpClient     *http.Client

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

var errInsecureEndpointWithTLS = errors.New("insecure HTTP endpoint cannot use TLS client configuration")

var maxResponseBodySize int64 = 4 * 1024 * 1024

func newClient(cfg oconf.Config) (*client, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *client) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *client) UploadMetrics(ctx context.Context, protoMetrics *metricpb.ResourceMetrics) (uploadErr error) {
	_ = "STUB: not implemented"
	return nil
}

var gzPool = sync.Pool{
	New: func() any {
		w := gzip.NewWriter(io.Discard)
		return w
	},
}

func (c *client) newRequest(ctx context.Context, body []byte) (request, error) {
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
