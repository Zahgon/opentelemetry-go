package otlptracehttp

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp/internal/otlpconfig"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp/internal/retry"
)

type Compression otlpconfig.Compression

type HTTPTransportProxyFunc func(*http.Request) (*url.URL, error)

const (
	NoCompression = Compression(otlpconfig.NoCompression)

	GzipCompression = Compression(otlpconfig.GzipCompression)
)

type Option interface {
	applyHTTPOption(otlpconfig.Config) otlpconfig.Config
}

func asHTTPOptions(opts []Option) []otlpconfig.HTTPOption { _ = "STUB: not implemented"; return nil }

type RetryConfig retry.Config

type wrappedOption struct {
	otlpconfig.HTTPOption
}

func (w wrappedOption) applyHTTPOption(cfg otlpconfig.Config) otlpconfig.Config {
	_ = "STUB: not implemented"
	return *new(otlpconfig.Config)
}

func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEndpointURL(u string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCompression(compression Compression) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithURLPath(urlPath string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTLSClientConfig(tlsCfg *tls.Config) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithInsecure() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeaders(headers map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTimeout(duration time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxRequestSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRetry(rc RetryConfig) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProxy(pf HTTPTransportProxyFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHTTPClient(c *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }
