package metric

import (
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type batcher struct {
	size int
}

func (b batcher) splitResourceMetrics(src *metricdata.ResourceMetrics) []*metricdata.ResourceMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (b batcher) splitScopeMetrics(sm metricdata.ScopeMetrics, firstSize int) []metricdata.ScopeMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (b batcher) splitMetric(m metricdata.Metrics, firstSize int) []metricdata.Metrics {
	_ = "STUB: not implemented"
	return nil
}

func copyMetricData(m metricdata.Metrics, offset, take int) metricdata.Metrics {
	_ = "STUB: not implemented"
	return *new(metricdata.Metrics)
}

func scopeMetricsDPC(sm metricdata.ScopeMetrics) int { _ = "STUB: not implemented"; return 0 }

func metricDPC(m metricdata.Metrics) int { _ = "STUB: not implemented"; return 0 }
