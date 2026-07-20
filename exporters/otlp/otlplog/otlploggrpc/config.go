package otlploggrpc

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc/internal/retry"
)

var (
	defaultEndpoint       = "localhost:4317"
	defaultTimeout        = 10 * time.Second
	defaultMaxRequestSize = 64 * 1024 * 1024
	defaultRetryCfg       = retry.DefaultConfig
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

type fnOpt func(config) config

func (f fnOpt) applyOption(c config) config { _ = "STUB: not implemented"; return *new(config) }

type Option interface {
	applyOption(config) config
}

type config struct {
	endpoint       setting[string]
	insecure       setting[bool]
	tlsCfg         setting[*tls.Config]
	headers        setting[map[string]string]
	compression    setting[Compression]
	maxRequestSize setting[int]
	timeout        setting[time.Duration]
	retryCfg       setting[retry.Config]

	gRPCCredentials    setting[credentials.TransportCredentials]
	serviceConfig      setting[string]
	reconnectionPeriod setting[time.Duration]
	dialOptions        setting[[]grpc.DialOption]
	gRPCConn           setting[*grpc.ClientConn]
}

func newConfig(options []Option) config { _ = "STUB: not implemented"; return *new(config) }

type RetryConfig retry.Config

func WithInsecure() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEndpointURL(rawURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReconnectionPeriod(rp time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type Compression int

const (
	NoCompression Compression = iota

	GzipCompression
)

func WithCompressor(compressor string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeaders(headers map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTLSCredentials(credential credentials.TransportCredentials) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithServiceConfig(serviceConfig string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDialOption(opts ...grpc.DialOption) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCConn(conn *grpc.ClientConn) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTimeout(duration time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxRequestSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRetry(rc RetryConfig) Option { _ = "STUB: not implemented"; return *new(Option) }

func convCompression(s string) (Compression, error) {
	_ = "STUB: not implemented"
	return *new(Compression), nil
}

func convEndpoint(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func convInsecure(s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func loadInsecureFromEnvEndpoint(envEndpoint []string) resolver[bool] {
	_ = "STUB: not implemented"
	return nil
}

func convHeaders(s string) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

func convDuration(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func loadEnvTLS[T *tls.Config]() resolver[T] { _ = "STUB: not implemented"; return nil }

var readFile = os.ReadFile

func loadCertPool(path string) (*x509.CertPool, error) { _ = "STUB: not implemented"; return nil, nil }

func loadCertificates(certPath, keyPath string) ([]tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func insecureFromScheme(prev setting[bool], scheme string) setting[bool] {
	_ = "STUB: not implemented"
	return nil
}

func compressorToCompression(compressor string) Compression {
	_ = "STUB: not implemented"
	return *new(Compression)
}

type setting[T any] struct {
	Value T
	Set   bool
}

func newSetting[T any](value T) setting[T] { _ = "STUB: not implemented"; return nil }

type resolver[T any] func(setting[T]) setting[T]

func (s setting[T]) Resolve(fn ...resolver[T]) setting[T] { _ = "STUB: not implemented"; return nil }

func getEnv[T any](keys []string, conv func(string) (T, error)) resolver[T] {
	_ = "STUB: not implemented"
	return nil
}

func fallback[T any](val T) resolver[T] { _ = "STUB: not implemented"; return nil }

func isValidHeaderKey(key string) bool { _ = "STUB: not implemented"; return false }

func isTokenChar(c rune) bool { _ = "STUB: not implemented"; return false }
