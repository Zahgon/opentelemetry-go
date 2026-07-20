package opencensus

import (
	"context"

	"go.opencensus.io/metric/metricproducer"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type MetricProducer struct {
	manager *metricproducer.Manager
}

func NewMetricProducer(...MetricOption) *MetricProducer { _ = "STUB: not implemented"; return nil }

var _ metric.Producer = (*MetricProducer)(nil)

func (p *MetricProducer) Produce(context.Context) ([]metricdata.ScopeMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
