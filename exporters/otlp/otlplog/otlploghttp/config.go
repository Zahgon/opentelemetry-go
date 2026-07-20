package otlploghttp

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"
	"os"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp/internal/retry"
)

var (
	defaultEndpoint                              = "localhost:4318"
	defaultPath                                  = "/v1/logs"
	defaultTimeout                               = 10 * time.Second
	defaultMaxRequestSize                        = 64 * 1024 * 1024
	defaultProxy          HTTPTransportProxyFunc = http.ProxyFromEnvironment
	defaultRetryCfg                              = retry.DefaultConfig
)

var (
	envEndpoint = []string{
		"OTEL_EXPORTER_OTLP_LOGS_ENDPOINT",
		"OTEL_EXPORTER_OTLP_ENDPOINT",
	}
	envInsecure = []string{
		"OTEL_EXPORTER_OTLP_LOGS_INSECURE",
		"OTEL_EXPORTER_OTLP_INSECURE",
	}

	envPathSignal = []string{"OTEL_EXPORTER_OTLP_LOGS_ENDPOINT"}
	envPathOTLP   = []string{"OTEL_EXPORTER_OTLP_ENDPOINT"}

	envHeaders = []string{
		"OTEL_EXPORTER_OTLP_LOGS_HEADERS",
		"OTEL_EXPORTER_OTLP_HEADERS",
	}

	envCompression = []string{
		"OTEL_EXPORTER_OTLP_LOGS_COMPRESSION",
		"OTEL_EXPORTER_OTLP_COMPRESSION",
	}

	envTimeout = []string{
		"OTEL_EXPORTER_OTLP_LOGS_TIMEOUT",
		"OTEL_EXPORTER_OTLP_TIMEOUT",
	}

	envTLSCert = []string{
		"OTEL_EXPORTER_OTLP_LOGS_CERTIFICATE",
		"OTEL_EXPORTER_OTLP_CERTIFICATE",
	}
	envTLSClient = []struct {
		Certificate string
		Key         string
	}{
		{
			"OTEL_EXPORTER_OTLP_LOGS_CLIENT_CERTIFICATE",
			"OTEL_EXPORTER_OTLP_LOGS_CLIENT_KEY",
		},
		{
			"OTEL_EXPORTER_OTLP_CLIENT_CERTIFICATE",
			"OTEL_EXPORTER_OTLP_CLIENT_KEY",
		},
	}
)

type Option interface {
	applyHTTPOption(config) config
}

type fnOpt func(config) config

func (f fnOpt) applyHTTPOption(c config) config { _ = "STUB: not implemented"; return *new(config) }

type config struct {
	endpoint       setting[string]
	path           setting[string]
	insecure       setting[bool]
	tlsCfg         setting[*tls.Config]
	headers        setting[map[string]string]
	compression    setting[Compression]
	maxRequestSize setting[int]
	timeout        setting[time.Duration]
	proxy          setting[HTTPTransportProxyFunc]
	retryCfg       setting[retry.Config]
	httpClient     *http.Client
}

func newConfig(options []Option) config { _ = "STUB: not implemented"; return *new(config) }

func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEndpointURL(rawURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

type Compression int

const (
	NoCompression Compression = iota

	GzipCompression
)

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

type RetryConfig retry.Config

func WithRetry(rc RetryConfig) Option { _ = "STUB: not implemented"; return *new(Option) }

type HTTPTransportProxyFunc func(*http.Request) (*url.URL, error)

func WithProxy(pf HTTPTransportProxyFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHTTPClient(c *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

type setting[T any] struct {
	Value T
	Set   bool
}

func newSetting[T any](value T) setting[T] { _ = "STUB: not implemented"; return nil }

type resolver[T any] func(setting[T]) setting[T]

func (s setting[T]) Resolve(fn ...resolver[T]) setting[T] { _ = "STUB: not implemented"; return nil }

func loadEnvTLS[T *tls.Config]() resolver[T] { _ = "STUB: not implemented"; return nil }

var readFile = os.ReadFile

func loadCertPool(path string) (*x509.CertPool, error) { _ = "STUB: not implemented"; return nil, nil }

func loadCertificates(certPath, keyPath string) ([]tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getenv[T any](keys []string, conv func(string) (T, error)) resolver[T] {
	_ = "STUB: not implemented"
	return nil
}

func convEndpoint(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func convPathExact(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func convPath(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func convInsecure(s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func loadInsecureFromEnvEndpoint(envEndpoint []string) resolver[bool] {
	_ = "STUB: not implemented"
	return nil
}

func insecureFromScheme(prev setting[bool], scheme string) setting[bool] {
	_ = "STUB: not implemented"
	return nil
}

func convHeaders(s string) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

func convCompression(s string) (Compression, error) {
	_ = "STUB: not implemented"
	return *new(Compression), nil
}

func convDuration(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func fallback[T any](val T) resolver[T] { _ = "STUB: not implemented"; return nil }

func isValidHeaderKey(key string) bool { _ = "STUB: not implemented"; return false }

func isTokenChar(c rune) bool { _ = "STUB: not implemented"; return false }
