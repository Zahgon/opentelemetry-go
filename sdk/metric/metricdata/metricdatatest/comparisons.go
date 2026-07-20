package metricdatatest

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

func equalResourceMetrics(a, b metricdata.ResourceMetrics, cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalScopeMetrics(a, b metricdata.ScopeMetrics, cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalMetrics(a, b metricdata.Metrics, cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalAggregations(a, b metricdata.Aggregation, cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalGauges[N int64 | float64](a, b metricdata.Gauge[N], cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalSums[N int64 | float64](a, b metricdata.Sum[N], cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalHistograms[N int64 | float64](a, b metricdata.Histogram[N], cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalDataPoints[N int64 | float64](
	a, b metricdata.DataPoint[N],
	cfg config,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalHistogramDataPoints[N int64 | float64](
	a, b metricdata.HistogramDataPoint[N],
	cfg config,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalExponentialHistograms[N int64 | float64](
	a, b metricdata.ExponentialHistogram[N],
	cfg config,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalExponentialHistogramDataPoints[N int64 | float64](
	a, b metricdata.ExponentialHistogramDataPoint[N],
	cfg config,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalExponentialBuckets(a, b metricdata.ExponentialBucket, _ config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalSummary(a, b metricdata.Summary, cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalSummaryDataPoint(a, b metricdata.SummaryDataPoint, cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func equalQuantileValue(a, b metricdata.QuantileValue, _ config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func notEqualStr(prefix string, expected, actual any) string { _ = "STUB: not implemented"; return "" }

func equalExtrema[N int64 | float64](a, b metricdata.Extrema[N], _ config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func eqExtrema[N int64 | float64](a, b metricdata.Extrema[N]) bool {
	_ = "STUB: not implemented"
	return false
}

func equalKeyValue(a, b attribute.KeyValue) bool { _ = "STUB: not implemented"; return false }

func equalExemplars[N int64 | float64](a, b metricdata.Exemplar[N], cfg config) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func diffSlices[T any](a, b []T, formatContext func(T) string, compare func(T, T) []string) string {
	_ = "STUB: not implemented"
	return ""
}

func missingAttrStr(name string) string { _ = "STUB: not implemented"; return "" }

func hasAttributesExemplars[T int64 | float64](
	exemplar metricdata.Exemplar[T],
	attrs ...attribute.KeyValue,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesDataPoints[T int64 | float64](
	dp metricdata.DataPoint[T],
	attrs ...attribute.KeyValue,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesGauge[T int64 | float64](gauge metricdata.Gauge[T], attrs ...attribute.KeyValue) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesSum[T int64 | float64](sum metricdata.Sum[T], attrs ...attribute.KeyValue) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesHistogramDataPoints[T int64 | float64](
	dp metricdata.HistogramDataPoint[T],
	attrs ...attribute.KeyValue,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesHistogram[T int64 | float64](
	histogram metricdata.Histogram[T],
	attrs ...attribute.KeyValue,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesExponentialHistogramDataPoints[T int64 | float64](
	dp metricdata.ExponentialHistogramDataPoint[T],
	attrs ...attribute.KeyValue,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesExponentialHistogram[T int64 | float64](
	histogram metricdata.ExponentialHistogram[T],
	attrs ...attribute.KeyValue,
) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesAggregation(agg metricdata.Aggregation, attrs ...attribute.KeyValue) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesMetrics(metrics metricdata.Metrics, attrs ...attribute.KeyValue) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesScopeMetrics(sm metricdata.ScopeMetrics, attrs ...attribute.KeyValue) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesResourceMetrics(rm metricdata.ResourceMetrics, attrs ...attribute.KeyValue) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesSummary(summary metricdata.Summary, attrs ...attribute.KeyValue) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}

func hasAttributesSummaryDataPoint(dp metricdata.SummaryDataPoint, attrs ...attribute.KeyValue) (reasons []string) {
	_ = "STUB: not implemented"
	return nil
}
