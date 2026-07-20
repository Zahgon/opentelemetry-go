package otlpmetrichttp

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp/internal/oconf"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp/internal/retry"
	"go.opentelemetry.io/otel/sdk/metric"
)

type Compression oconf.Compression

type HTTPTransportProxyFunc func(*http.Request) (*url.URL, error)

const (
	NoCompression = Compression(oconf.NoCompression)

	GzipCompression = Compression(oconf.GzipCompression)
)

type Option interface {
	applyHTTPOption(oconf.Config) oconf.Config
}

func asHTTPOptions(opts []Option) []oconf.HTTPOption { _ = "STUB: not implemented"; return nil }

type RetryConfig retry.Config

type wrappedOption struct {
	oconf.HTTPOption
}

func (w wrappedOption) applyHTTPOption(cfg oconf.Config) oconf.Config {
	_ = "STUB: not implemented"
	return *new(oconf.Config)
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

func WithTemporalitySelector(selector metric.TemporalitySelector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithAggregationSelector(selector metric.AggregationSelector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithProxy(pf HTTPTransportProxyFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHTTPClient(c *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }
