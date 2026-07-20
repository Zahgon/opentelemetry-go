package otlpmetricgrpc

import (
	"context"
	"time"

	colmetricpb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	metricpb "go.opentelemetry.io/proto/otlp/metrics/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc/internal/oconf"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc/internal/retry"
)

type client struct {
	metadata       metadata.MD
	exportTimeout  time.Duration
	maxRequestSize int
	requestFunc    retry.RequestFunc

	ourConn bool
	conn    *grpc.ClientConn
	msc     colmetricpb.MetricsServiceClient
}

func newClient(_ context.Context, cfg oconf.Config) (*client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *client) UploadMetrics(ctx context.Context, protoMetrics *metricpb.ResourceMetrics) (uploadErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) exportContext(parent context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

//nolint:gosec  // cancel is handled by the caller.

func retryable(err error) (bool, time.Duration) {
	_ = "STUB: not implemented"
	return false, *new(time.Duration)
}

func retryableGRPCStatus(s *status.Status) (bool, time.Duration) {
	_ = "STUB: not implemented"
	return false, *new(time.Duration)
}

func throttleDelay(s *status.Status) (bool, time.Duration) {
	_ = "STUB: not implemented"
	return false, *new(time.Duration)
}
