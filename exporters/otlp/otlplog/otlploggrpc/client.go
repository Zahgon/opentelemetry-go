package otlploggrpc

import (
	"context"
	"sync/atomic"
	"time"

	collogpb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	logpb "go.opentelemetry.io/proto/otlp/logs/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc/internal/observ"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc/internal/retry"
)

type client struct {
	metadata       metadata.MD
	exportTimeout  time.Duration
	maxRequestSize int
	requestFunc    retry.RequestFunc

	ourConn bool
	conn    *grpc.ClientConn
	lsc     collogpb.LogsServiceClient

	instrumentation *observ.Instrumentation
}

var newGRPCClientFn = grpc.NewClient

func newClient(cfg config) (*client, error) { _ = "STUB: not implemented"; return nil, nil }

var exporterN atomic.Int64

func nextExporterID() int64 { _ = "STUB: not implemented"; return 0 }

func newGRPCDialOptions(cfg config) []grpc.DialOption { _ = "STUB: not implemented"; return nil }

func (c *client) UploadLogs(ctx context.Context, rl []*logpb.ResourceLogs) (uploadErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *client) exportContext(parent context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

//nolint:gosec  // cancel is handled by caller.

type noopClient struct{}

func newNoopClient() *noopClient { _ = "STUB: not implemented"; return nil }

func (*noopClient) UploadLogs(context.Context, []*logpb.ResourceLogs) error {
	_ = "STUB: not implemented"
	return nil
}

func (*noopClient) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

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
