package metric

import (
	"errors"
	"fmt"
)

var errAgg = errors.New("aggregation")

type Aggregation interface {
	copy() Aggregation

	err() error
}

type AggregationDrop struct{}

var _ Aggregation = AggregationDrop{}

func (d AggregationDrop) copy() Aggregation { _ = "STUB: not implemented"; return *new(Aggregation) }

func (AggregationDrop) err() error { _ = "STUB: not implemented"; return nil }

type AggregationDefault struct{}

var _ Aggregation = AggregationDefault{}

func (d AggregationDefault) copy() Aggregation { _ = "STUB: not implemented"; return *new(Aggregation) }

func (AggregationDefault) err() error { _ = "STUB: not implemented"; return nil }

type AggregationSum struct{}

var _ Aggregation = AggregationSum{}

func (s AggregationSum) copy() Aggregation { _ = "STUB: not implemented"; return *new(Aggregation) }

func (AggregationSum) err() error { _ = "STUB: not implemented"; return nil }

type AggregationLastValue struct{}

var _ Aggregation = AggregationLastValue{}

func (l AggregationLastValue) copy() Aggregation {
	_ = "STUB: not implemented"
	return *new(Aggregation)
}

func (AggregationLastValue) err() error { _ = "STUB: not implemented"; return nil }

type AggregationExplicitBucketHistogram struct {
	Boundaries []float64

	NoMinMax bool
}

var _ Aggregation = AggregationExplicitBucketHistogram{}

var errHist = fmt.Errorf("%w: explicit bucket histogram", errAgg)

func (h AggregationExplicitBucketHistogram) err() error { _ = "STUB: not implemented"; return nil }

func (h AggregationExplicitBucketHistogram) copy() Aggregation {
	_ = "STUB: not implemented"
	return *new(Aggregation)
}

type AggregationBase2ExponentialHistogram struct {
	MaxSize int32

	MaxScale int32

	NoMinMax bool
}

var _ Aggregation = AggregationBase2ExponentialHistogram{}

func (e AggregationBase2ExponentialHistogram) copy() Aggregation {
	_ = "STUB: not implemented"
	return *new(Aggregation)
}

const (
	expoMaxScale = 20
	expoMinScale = -10
)

var errExpoHist = fmt.Errorf("%w: exponential histogram", errAgg)

func (e AggregationBase2ExponentialHistogram) err() error { _ = "STUB: not implemented"; return nil }
