package metricdata

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
)

type ResourceMetrics struct {
	Resource *resource.Resource

	ScopeMetrics []ScopeMetrics
}

type ScopeMetrics struct {
	Scope instrumentation.Scope

	Metrics []Metrics
}

type Metrics struct {
	Name string

	Description string

	Unit string

	Data Aggregation
}

type Aggregation interface {
	privateAggregation()
}

type Gauge[N int64 | float64] struct {
	DataPoints []DataPoint[N]
}

func (Gauge[N]) privateAggregation() { _ = "STUB: not implemented"; return }

type Sum[N int64 | float64] struct {
	DataPoints []DataPoint[N]

	Temporality Temporality

	IsMonotonic bool
}

func (Sum[N]) privateAggregation() { _ = "STUB: not implemented"; return }

type DataPoint[N int64 | float64] struct {
	Attributes attribute.Set

	StartTime time.Time `json:",omitempty"`

	Time time.Time `json:",omitempty"`

	Value N

	Exemplars []Exemplar[N] `json:",omitempty"`
}

type Histogram[N int64 | float64] struct {
	DataPoints []HistogramDataPoint[N]

	Temporality Temporality
}

func (Histogram[N]) privateAggregation() { _ = "STUB: not implemented"; return }

type HistogramDataPoint[N int64 | float64] struct {
	Attributes attribute.Set

	StartTime time.Time

	Time time.Time

	Count uint64

	Bounds []float64

	BucketCounts []uint64

	Min Extrema[N]

	Max Extrema[N]

	Sum N

	Exemplars []Exemplar[N] `json:",omitempty"`
}

type ExponentialHistogram[N int64 | float64] struct {
	DataPoints []ExponentialHistogramDataPoint[N]

	Temporality Temporality
}

func (ExponentialHistogram[N]) privateAggregation() { _ = "STUB: not implemented"; return }

type ExponentialHistogramDataPoint[N int64 | float64] struct {
	Attributes attribute.Set

	StartTime time.Time

	Time time.Time

	Count uint64

	Min Extrema[N]

	Max Extrema[N]

	Sum N

	Scale int32

	ZeroCount uint64

	PositiveBucket ExponentialBucket

	NegativeBucket ExponentialBucket

	ZeroThreshold float64

	Exemplars []Exemplar[N] `json:",omitempty"`
}

type ExponentialBucket struct {
	Offset int32

	Counts []uint64
}

type Extrema[N int64 | float64] struct {
	value N
	valid bool
}

func (e Extrema[N]) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Extrema[N]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewExtrema[N int64 | float64](v N) Extrema[N] { _ = "STUB: not implemented"; return nil }

func (e Extrema[N]) Value() (v N, defined bool) { _ = "STUB: not implemented"; return *new(N), false }

type Exemplar[N int64 | float64] struct {
	FilteredAttributes []attribute.KeyValue

	Time time.Time

	Value N

	SpanID []byte `json:",omitempty"`

	TraceID []byte `json:",omitempty"`
}

type Summary struct {
	DataPoints []SummaryDataPoint
}

func (Summary) privateAggregation() { _ = "STUB: not implemented"; return }

type SummaryDataPoint struct {
	Attributes attribute.Set

	StartTime time.Time

	Time time.Time

	Count uint64

	Sum float64

	QuantileValues []QuantileValue
}

type QuantileValue struct {
	Quantile float64

	Value float64
}
