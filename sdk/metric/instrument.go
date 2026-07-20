//go:generate stringer -type=InstrumentKind -trimprefix=InstrumentKind

package metric

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric/internal/aggregate"
)

var zeroScope instrumentation.Scope

type InstrumentKind uint8

const (
	instrumentKindUndefined InstrumentKind = 0

	InstrumentKindCounter InstrumentKind = 1

	InstrumentKindUpDownCounter InstrumentKind = 2

	InstrumentKindHistogram InstrumentKind = 3

	InstrumentKindObservableCounter InstrumentKind = 4

	InstrumentKindObservableUpDownCounter InstrumentKind = 5

	InstrumentKindObservableGauge InstrumentKind = 6

	InstrumentKindGauge InstrumentKind = 7
)

type nonComparable [0]func()

type Instrument struct {
	Name string

	Description string

	Kind InstrumentKind

	Unit string

	Scope instrumentation.Scope

	nonComparable
}

func (i Instrument) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (i Instrument) matches(other Instrument) bool { _ = "STUB: not implemented"; return false }

func (i Instrument) matchesName(other Instrument) bool { _ = "STUB: not implemented"; return false }

func (i Instrument) matchesDescription(other Instrument) bool {
	_ = "STUB: not implemented"
	return false
}

func (i Instrument) matchesKind(other Instrument) bool { _ = "STUB: not implemented"; return false }

func (i Instrument) matchesUnit(other Instrument) bool { _ = "STUB: not implemented"; return false }

func (i Instrument) matchesScope(other Instrument) bool { _ = "STUB: not implemented"; return false }

type Stream struct {
	Name string

	Description string

	Unit string

	Aggregation Aggregation

	AttributeFilter attribute.Filter

	ExemplarReservoirProviderSelector ExemplarReservoirProviderSelector
}

type instID struct {
	Name string

	Description string

	Kind InstrumentKind

	Unit string

	Number string
}

func (i instID) normalize() instID { _ = "STUB: not implemented"; return *new(instID) }

type rawAttributesOption interface {
	RawAttributes() []attribute.KeyValue
	Experimental()
}

func extractRawKVs[T any](opts []T) []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func resolveAttributes(configAttrs attribute.Set, rawKVs []attribute.KeyValue) attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

type int64Inst struct {
	measures []aggregate.Measure[int64]

	embedded.Int64Counter
	embedded.Int64UpDownCounter
	embedded.Int64Histogram
	embedded.Int64Gauge
}

var (
	_ metric.Int64Counter       = (*int64Inst)(nil)
	_ metric.Int64UpDownCounter = (*int64Inst)(nil)
	_ metric.Int64Histogram     = (*int64Inst)(nil)
	_ metric.Int64Gauge         = (*int64Inst)(nil)
)

func (i *int64Inst) Add(ctx context.Context, val int64, opts ...metric.AddOption) {
	_ = "STUB: not implemented"
	return
}

func (i *int64Inst) Record(ctx context.Context, val int64, opts ...metric.RecordOption) {
	_ = "STUB: not implemented"
	return
}

func (i *int64Inst) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

func (i *int64Inst) aggregate(
	ctx context.Context,
	val int64,
	s attribute.Set,
) {
	_ = "STUB: not implemented"
	return
}

type float64Inst struct {
	measures []aggregate.Measure[float64]

	embedded.Float64Counter
	embedded.Float64UpDownCounter
	embedded.Float64Histogram
	embedded.Float64Gauge
}

var (
	_ metric.Float64Counter       = (*float64Inst)(nil)
	_ metric.Float64UpDownCounter = (*float64Inst)(nil)
	_ metric.Float64Histogram     = (*float64Inst)(nil)
	_ metric.Float64Gauge         = (*float64Inst)(nil)
)

func (i *float64Inst) Add(ctx context.Context, val float64, opts ...metric.AddOption) {
	_ = "STUB: not implemented"
	return
}

func (i *float64Inst) Record(ctx context.Context, val float64, opts ...metric.RecordOption) {
	_ = "STUB: not implemented"
	return
}

func (i *float64Inst) Enabled(context.Context) bool { _ = "STUB: not implemented"; return false }

func (i *float64Inst) aggregate(ctx context.Context, val float64, s attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type observableID[N int64 | float64] struct {
	name        string
	description string
	kind        InstrumentKind
	unit        string
	scope       instrumentation.Scope
}

type float64Observable struct {
	metric.Float64Observable
	*observable[float64]

	embedded.Float64ObservableCounter
	embedded.Float64ObservableUpDownCounter
	embedded.Float64ObservableGauge
}

var (
	_ metric.Float64ObservableCounter       = float64Observable{}
	_ metric.Float64ObservableUpDownCounter = float64Observable{}
	_ metric.Float64ObservableGauge         = float64Observable{}
)

func newFloat64Observable(m *meter, kind InstrumentKind, name, desc, u string) float64Observable {
	_ = "STUB: not implemented"
	return *new(float64Observable)
}

type int64Observable struct {
	metric.Int64Observable
	*observable[int64]

	embedded.Int64ObservableCounter
	embedded.Int64ObservableUpDownCounter
	embedded.Int64ObservableGauge
}

var (
	_ metric.Int64ObservableCounter       = int64Observable{}
	_ metric.Int64ObservableUpDownCounter = int64Observable{}
	_ metric.Int64ObservableGauge         = int64Observable{}
)

func newInt64Observable(m *meter, kind InstrumentKind, name, desc, u string) int64Observable {
	_ = "STUB: not implemented"
	return *new(int64Observable)
}

type observable[N int64 | float64] struct {
	metric.Observable
	observableID[N]

	meter           *meter
	measures        measures[N]
	dropAggregation bool
}

func newObservable[N int64 | float64](m *meter, kind InstrumentKind, name, desc, u string) *observable[N] {
	_ = "STUB: not implemented"
	return nil
}

func (o *observable[N]) observe(val N, s attribute.Set) { _ = "STUB: not implemented"; return }

func (o *observable[N]) appendMeasures(meas []aggregate.Measure[N]) {
	_ = "STUB: not implemented"
	return
}

type measures[N int64 | float64] []aggregate.Measure[N]

func (m measures[N]) observe(val N, s attribute.Set) { _ = "STUB: not implemented"; return }

var errEmptyAgg = errors.New("no aggregators for observable instrument")

func (o *observable[N]) registerable(m *meter) error { _ = "STUB: not implemented"; return nil }
