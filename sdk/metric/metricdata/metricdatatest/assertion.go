package metricdatatest

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type Datatypes interface {
	metricdata.DataPoint[float64] |
		metricdata.DataPoint[int64] |
		metricdata.Gauge[float64] |
		metricdata.Gauge[int64] |
		metricdata.Histogram[float64] |
		metricdata.Histogram[int64] |
		metricdata.HistogramDataPoint[float64] |
		metricdata.HistogramDataPoint[int64] |
		metricdata.Extrema[int64] |
		metricdata.Extrema[float64] |
		metricdata.Metrics |
		metricdata.ResourceMetrics |
		metricdata.ScopeMetrics |
		metricdata.Sum[float64] |
		metricdata.Sum[int64] |
		metricdata.Exemplar[float64] |
		metricdata.Exemplar[int64] |
		metricdata.ExponentialHistogram[float64] |
		metricdata.ExponentialHistogram[int64] |
		metricdata.ExponentialHistogramDataPoint[float64] |
		metricdata.ExponentialHistogramDataPoint[int64] |
		metricdata.ExponentialBucket |
		metricdata.Summary |
		metricdata.SummaryDataPoint |
		metricdata.QuantileValue
}

type TestingT interface {
	Helper()

	Error(...any)
}

type config struct {
	ignoreTimestamp bool
	ignoreExemplars bool
	ignoreValue     bool
}

func newConfig(opts []Option) config { _ = "STUB: not implemented"; return *new(config) }

type Option interface {
	apply(cfg config) config
}

type fnOption func(cfg config) config

func (fn fnOption) apply(cfg config) config { _ = "STUB: not implemented"; return *new(config) }

func IgnoreTimestamp() Option { _ = "STUB: not implemented"; return *new(Option) }

func IgnoreExemplars() Option { _ = "STUB: not implemented"; return *new(Option) }

func IgnoreValue() Option { _ = "STUB: not implemented"; return *new(Option) }

func AssertEqual[T Datatypes](t TestingT, expected, actual T, opts ...Option) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertAggregationsEqual(t TestingT, expected, actual metricdata.Aggregation, opts ...Option) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertHasAttributes[T Datatypes](t TestingT, actual T, attrs ...attribute.KeyValue) bool {
	_ = "STUB: not implemented"
	return false
}
