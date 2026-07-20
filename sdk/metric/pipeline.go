package metric

import (
	"container/list"
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric/exemplar"
	"go.opentelemetry.io/otel/sdk/metric/internal/aggregate"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
)

var (
	errCreatingAggregators     = errors.New("could not create all aggregators")
	errIncompatibleAggregation = errors.New("incompatible aggregation")
	errUnknownAggregation      = errors.New("unrecognized aggregation")
)

type instrumentSync struct {
	name        string
	description string
	unit        string
	compAgg     aggregate.ComputeAggregation
}

func newPipeline(
	res *resource.Resource,
	reader Reader,
	views []View,
	exemplarFilter exemplar.Filter,
	cardinalityLimit int,
) *pipeline {
	_ = "STUB: not implemented"
	return nil
}

type pipeline struct {
	resource *resource.Resource

	reader Reader
	views  []View

	sync.Mutex
	int64Measures    map[observableID[int64]][]aggregate.Measure[int64]
	float64Measures  map[observableID[float64]][]aggregate.Measure[float64]
	aggregations     map[instrumentation.Scope][]instrumentSync
	callbacks        []func(context.Context) error
	multiCallbacks   list.List
	exemplarFilter   exemplar.Filter
	cardinalityLimit int
}

func (p *pipeline) addInt64Measure(id observableID[int64], m []aggregate.Measure[int64]) {
	_ = "STUB: not implemented"
	return
}

func (p *pipeline) addFloat64Measure(id observableID[float64], m []aggregate.Measure[float64]) {
	_ = "STUB: not implemented"
	return
}

func (p *pipeline) addSync(scope instrumentation.Scope, iSync instrumentSync) {
	_ = "STUB: not implemented"
	return
}

type multiCallback func(context.Context) error

func (p *pipeline) addMultiCallback(c multiCallback) (unregister func()) {
	_ = "STUB: not implemented"
	return nil
}

func (p *pipeline) produce(ctx context.Context, rm *metricdata.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

type inserter[N int64 | float64] struct {
	aggregators *cache[instID, aggVal[N]]

	views *cache[string, instID]

	pipeline *pipeline
}

func newInserter[N int64 | float64](p *pipeline, vc *cache[string, instID]) *inserter[N] {
	_ = "STUB: not implemented"
	return nil
}

func (i *inserter[N]) Instrument(
	inst Instrument,
	allowedKeys []attribute.Key,
	readerAggregation Aggregation,
) ([]aggregate.Measure[N], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *inserter[N]) addCallback(cback func(context.Context) error) {
	_ = "STUB: not implemented"
	return
}

var aggIDCount atomic.Uint64

type aggVal[N int64 | float64] struct {
	ID      uint64
	Measure aggregate.Measure[N]
	Err     error
}

func (i *inserter[N]) readerDefaultAggregation(kind InstrumentKind) Aggregation {
	_ = "STUB: not implemented"
	return *new(Aggregation)
}

func (i *inserter[N]) cachedAggregator(
	scope instrumentation.Scope,
	kind InstrumentKind,
	stream Stream,
	readerAggregation Aggregation,
) (meas aggregate.Measure[N], aggID uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (i *inserter[N]) getCardinalityLimit(kind InstrumentKind) int {
	_ = "STUB: not implemented"
	return 0
}

func (i *inserter[N]) logConflict(id instID) { _ = "STUB: not implemented"; return }

func (*inserter[N]) instID(kind InstrumentKind, stream Stream) instID {
	_ = "STUB: not implemented"
	return *new(instID)
}

func (i *inserter[N]) aggregateFunc(
	b aggregate.Builder[N],
	agg Aggregation,
	kind InstrumentKind,
) (meas aggregate.Measure[N], comp aggregate.ComputeAggregation, err error) {
	_ = "STUB: not implemented"
	return nil, *new(aggregate.ComputeAggregation), nil
}

func isAggregatorCompatible(kind InstrumentKind, agg Aggregation) error {
	_ = "STUB: not implemented"
	return nil
}

type pipelines []*pipeline

func newPipelines(
	res *resource.Resource,
	readers []Reader,
	views []View,
	exemplarFilter exemplar.Filter,
	cardinalityLimit int,
) pipelines {
	_ = "STUB: not implemented"
	return *new(pipelines)
}

type unregisterFuncs struct {
	embedded.Registration
	f []func()
}

func (u unregisterFuncs) Unregister() error { _ = "STUB: not implemented"; return nil }

type resolver[N int64 | float64] struct {
	inserters []*inserter[N]
}

func newResolver[N int64 | float64](p pipelines, vc *cache[string, instID]) resolver[N] {
	_ = "STUB: not implemented"
	return nil
}

func (r resolver[N]) Aggregators(id Instrument, allowedKeys []attribute.Key) ([]aggregate.Measure[N], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r resolver[N]) HistogramAggregators(
	id Instrument,
	allowedKeys []attribute.Key,
	boundaries []float64,
) ([]aggregate.Measure[N], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
