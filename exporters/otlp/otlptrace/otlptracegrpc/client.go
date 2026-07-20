package otlptracegrpc

import (
	"context"
	"errors"
	"sync"
	"time"

	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc/internal/observ"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc/internal/retry"
)

type client struct {
	endpoint       string
	dialOpts       []grpc.DialOption
	metadata       metadata.MD
	exportTimeout  time.Duration
	maxRequestSize int
	requestFunc    retry.RequestFunc

	stopCtx context.Context

	stopFunc context.CancelFunc

	ourConn bool
	conn    *grpc.ClientConn
	tscMu   sync.RWMutex
	tsc     coltracepb.TraceServiceClient

	instID int64
	inst   *observ.Instrumentation
}

var _ otlptrace.Client = (*client)(nil)

func NewClient(opts ...Option) otlptrace.Client {
	_ = "STUB: not implemented"
	return *new(otlptrace.Client)
}

func newClient(opts ...Option) *client { _ = "STUB: not implemented"; return nil }

//nolint:gosec  // cancel called in client shutdown.

func (c *client) Start(context.Context) error { _ = "STUB: not implemented"; return nil }

var errAlreadyStopped = errors.New("the client is already stopped")

func (c *client) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var errShutdown = errors.New("the client is shutdown")

func (c *client) UploadTraces(ctx context.Context, protoSpans []*tracepb.ResourceSpans) (uploadErr error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) exportContext(parent context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

//nolint:gosec  // cancel called by caller when export is complete.

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

func (*client) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }
