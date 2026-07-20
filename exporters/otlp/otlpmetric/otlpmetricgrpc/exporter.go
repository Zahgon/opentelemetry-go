package otlpmetricgrpc

import (
	"context"
	"errors"
	"sync"

	metricpb "go.opentelemetry.io/proto/otlp/metrics/v1"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc/internal/observ"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc/internal/oconf"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type Exporter struct {
	clientMu sync.Mutex
	client   interface {
		UploadMetrics(context.Context, *metricpb.ResourceMetrics) error
		Shutdown(context.Context) error
	}

	temporalitySelector metric.TemporalitySelector
	aggregationSelector metric.AggregationSelector

	shutdownOnce sync.Once

	inst *observ.Instrumentation
}

func newExporter(c *client, cfg oconf.Config) (*Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Exporter) Temporality(k metric.InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func (e *Exporter) Aggregation(k metric.InstrumentKind) metric.Aggregation {
	_ = "STUB: not implemented"
	return *new(metric.Aggregation)
}

func (e *Exporter) Export(ctx context.Context, rm *metricdata.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Exporter) ForceFlush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *Exporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var errShutdown = errors.New("gRPC exporter is shutdown")

type shutdownClient struct{}

func (shutdownClient) err(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c shutdownClient) UploadMetrics(ctx context.Context, _ *metricpb.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (c shutdownClient) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (*Exporter) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }

func New(ctx context.Context, options ...Option) (*Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
