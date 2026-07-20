package faasconv

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	addOptPool = &sync.Pool{New: func() any { return &[]metric.AddOption{} }}
	recOptPool = &sync.Pool{New: func() any { return &[]metric.RecordOption{} }}
)

type TriggerAttr string

var (
	TriggerDatasource TriggerAttr = "datasource"

	TriggerHTTP TriggerAttr = "http"

	TriggerPubSub TriggerAttr = "pubsub"

	TriggerTimer TriggerAttr = "timer"

	TriggerOther TriggerAttr = "other"
)

type Coldstarts struct {
	metric.Int64Counter
}

func NewColdstarts(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (Coldstarts, error) {
	_ = "STUB: not implemented"
	return *new(Coldstarts), nil
}

func (m Coldstarts) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (Coldstarts) Name() string { _ = "STUB: not implemented"; return "" }

func (Coldstarts) Unit() string { _ = "STUB: not implemented"; return "" }

func (Coldstarts) Description() string { _ = "STUB: not implemented"; return "" }

func (m Coldstarts) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (Coldstarts) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUUsage struct {
	metric.Float64Histogram
}

func NewCPUUsage(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (CPUUsage, error) {
	_ = "STUB: not implemented"
	return *new(CPUUsage), nil
}

func (m CPUUsage) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (CPUUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m CPUUsage) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (CPUUsage) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Errors struct {
	metric.Int64Counter
}

func NewErrors(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (Errors, error) {
	_ = "STUB: not implemented"
	return *new(Errors), nil
}

func (m Errors) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (Errors) Name() string { _ = "STUB: not implemented"; return "" }

func (Errors) Unit() string { _ = "STUB: not implemented"; return "" }

func (Errors) Description() string { _ = "STUB: not implemented"; return "" }

func (m Errors) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (Errors) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type InitDuration struct {
	metric.Float64Histogram
}

func NewInitDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (InitDuration, error) {
	_ = "STUB: not implemented"
	return *new(InitDuration), nil
}

func (m InitDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (InitDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (InitDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (InitDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m InitDuration) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (InitDuration) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Invocations struct {
	metric.Int64Counter
}

func NewInvocations(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (Invocations, error) {
	_ = "STUB: not implemented"
	return *new(Invocations), nil
}

func (m Invocations) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (Invocations) Name() string { _ = "STUB: not implemented"; return "" }

func (Invocations) Unit() string { _ = "STUB: not implemented"; return "" }

func (Invocations) Description() string { _ = "STUB: not implemented"; return "" }

func (m Invocations) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (Invocations) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type InvokeDuration struct {
	metric.Float64Histogram
}

func NewInvokeDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (InvokeDuration, error) {
	_ = "STUB: not implemented"
	return *new(InvokeDuration), nil
}

func (m InvokeDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (InvokeDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (InvokeDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (InvokeDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m InvokeDuration) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (InvokeDuration) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemUsage struct {
	metric.Int64Histogram
}

func NewMemUsage(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (MemUsage, error) {
	_ = "STUB: not implemented"
	return *new(MemUsage), nil
}

func (m MemUsage) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (MemUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (MemUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemUsage) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (MemUsage) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetIO struct {
	metric.Int64Histogram
}

func NewNetIO(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (NetIO, error) {
	_ = "STUB: not implemented"
	return *new(NetIO), nil
}

func (m NetIO) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (NetIO) Name() string { _ = "STUB: not implemented"; return "" }

func (NetIO) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetIO) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetIO) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (NetIO) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Timeouts struct {
	metric.Int64Counter
}

func NewTimeouts(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (Timeouts, error) {
	_ = "STUB: not implemented"
	return *new(Timeouts), nil
}

func (m Timeouts) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (Timeouts) Name() string { _ = "STUB: not implemented"; return "" }

func (Timeouts) Unit() string { _ = "STUB: not implemented"; return "" }

func (Timeouts) Description() string { _ = "STUB: not implemented"; return "" }

func (m Timeouts) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (Timeouts) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
