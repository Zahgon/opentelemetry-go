package stdoutmetric

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric/internal/observ"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type exporter struct {
	encVal atomic.Value

	shutdownOnce sync.Once

	temporalitySelector metric.TemporalitySelector
	aggregationSelector metric.AggregationSelector

	redactTimestamps bool

	inst *observ.Instrumentation
}

func New(options ...Option) (metric.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(metric.Exporter), nil
}

func (e *exporter) Temporality(k metric.InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func (e *exporter) Aggregation(k metric.InstrumentKind) metric.Aggregation {
	_ = "STUB: not implemented"
	return *new(metric.Aggregation)
}

func (e *exporter) Export(ctx context.Context, data *metricdata.ResourceMetrics) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*exporter) ForceFlush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *exporter) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (*exporter) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }

func redactTimestamps(orig *metricdata.ResourceMetrics) { _ = "STUB: not implemented"; return }

var errUnknownAggType = errors.New("unknown aggregation type")

func redactAggregationTimestamps(orig metricdata.Aggregation) metricdata.Aggregation {
	_ = "STUB: not implemented"
	return *new(metricdata.Aggregation)
}

func redactHistogramTimestamps[T int64 | float64](
	hdp []metricdata.HistogramDataPoint[T],
) []metricdata.HistogramDataPoint[T] {
	_ = "STUB: not implemented"
	return nil
}

func redactDataPointTimestamps[T int64 | float64](sdp []metricdata.DataPoint[T]) []metricdata.DataPoint[T] {
	_ = "STUB: not implemented"
	return nil
}

func countDataPoints(rm *metricdata.ResourceMetrics) int64 { _ = "STUB: not implemented"; return 0 }
