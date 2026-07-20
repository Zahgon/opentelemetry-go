package otlptrace

import (
	"context"

	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
)

type Client interface {
	Start(ctx context.Context) error

	Stop(ctx context.Context) error

	UploadTraces(ctx context.Context, protoSpans []*tracepb.ResourceSpans) error
}
