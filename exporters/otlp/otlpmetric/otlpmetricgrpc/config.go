package otlpmetricgrpc

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc/internal/oconf"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc/internal/retry"
	"go.opentelemetry.io/otel/sdk/metric"
)

type Option interface {
	applyGRPCOption(oconf.Config) oconf.Config
}

func asGRPCOptions(opts []Option) []oconf.GRPCOption { _ = "STUB: not implemented"; return nil }

type RetryConfig retry.Config

type wrappedOption struct {
	oconf.GRPCOption
}

func (w wrappedOption) applyGRPCOption(cfg oconf.Config) oconf.Config {
	_ = "STUB: not implemented"
	return *new(oconf.Config)
}

func WithInsecure() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEndpointURL(u string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReconnectionPeriod(rp time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func compressorToCompression(compressor string) oconf.Compression {
	_ = "STUB: not implemented"
	return *new(oconf.Compression)
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

func WithTemporalitySelector(selector metric.TemporalitySelector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithAggregationSelector(selector metric.AggregationSelector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
