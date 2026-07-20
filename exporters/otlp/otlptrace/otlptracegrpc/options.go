package otlptracegrpc

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc/internal/otlpconfig"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc/internal/retry"
)

type Option interface {
	applyGRPCOption(otlpconfig.Config) otlpconfig.Config
}

func asGRPCOptions(opts []Option) []otlpconfig.GRPCOption { _ = "STUB: not implemented"; return nil }

type RetryConfig retry.Config

type wrappedOption struct {
	otlpconfig.GRPCOption
}

func (w wrappedOption) applyGRPCOption(cfg otlpconfig.Config) otlpconfig.Config {
	_ = "STUB: not implemented"
	return *new(otlpconfig.Config)
}

func WithInsecure() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEndpointURL(u string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReconnectionPeriod(rp time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func compressorToCompression(compressor string) otlpconfig.Compression {
	_ = "STUB: not implemented"
	return *new(otlpconfig.Compression)
}

func WithCompressor(compressor string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHeaders(headers map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTLSCredentials(creds credentials.TransportCredentials) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithServiceConfig(serviceConfig string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDialOption(opts ...grpc.DialOption) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCConn(conn *grpc.ClientConn) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTimeout(duration time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxRequestSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRetry(settings RetryConfig) Option { _ = "STUB: not implemented"; return *new(Option) }
