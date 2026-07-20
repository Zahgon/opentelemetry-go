package faasconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
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

var newColdstartsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of invocation cold starts."),
	metric.WithUnit("{coldstart}"),
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

func (m Coldstarts) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Coldstarts) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ColdstartsObservable struct {
	metric.Int64ObservableCounter
}

var newColdstartsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Number of invocation cold starts."),
	metric.WithUnit("{coldstart}"),
}

func NewColdstartsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ColdstartsObservable, error) {
	_ = "STUB: not implemented"
	return *new(ColdstartsObservable), nil
}

func (m ColdstartsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (ColdstartsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ColdstartsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ColdstartsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ColdstartsObservable) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUUsage struct {
	metric.Float64Histogram
}

var newCPUUsageOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Distribution of CPU usage per invocation."),
	metric.WithUnit("s"),
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

func (m CPUUsage) RecordSet(ctx context.Context, val float64, set attribute.Set) {
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

var newErrorsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of invocation errors."),
	metric.WithUnit("{error}"),
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

func (m Errors) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Errors) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ErrorsObservable struct {
	metric.Int64ObservableCounter
}

var newErrorsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Number of invocation errors."),
	metric.WithUnit("{error}"),
}

func NewErrorsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ErrorsObservable, error) {
	_ = "STUB: not implemented"
	return *new(ErrorsObservable), nil
}

func (m ErrorsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (ErrorsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ErrorsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ErrorsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ErrorsObservable) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type InitDuration struct {
	metric.Float64Histogram
}

var newInitDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Measures the duration of the function's initialization, such as a cold start."),
	metric.WithUnit("s"),
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

func (m InitDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
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

var newInvocationsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of successful invocations."),
	metric.WithUnit("{invocation}"),
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

func (m Invocations) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Invocations) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type InvocationsObservable struct {
	metric.Int64ObservableCounter
}

var newInvocationsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Number of successful invocations."),
	metric.WithUnit("{invocation}"),
}

func NewInvocationsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (InvocationsObservable, error) {
	_ = "STUB: not implemented"
	return *new(InvocationsObservable), nil
}

func (m InvocationsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (InvocationsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (InvocationsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (InvocationsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (InvocationsObservable) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type InvokeDuration struct {
	metric.Float64Histogram
}

var newInvokeDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Measures the duration of the function's logic execution."),
	metric.WithUnit("s"),
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

func (m InvokeDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
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

var newMemUsageOpts = []metric.Int64HistogramOption{
	metric.WithDescription("Distribution of max memory usage per invocation."),
	metric.WithUnit("By"),
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

func (m MemUsage) RecordSet(ctx context.Context, val int64, set attribute.Set) {
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

var newNetIOOpts = []metric.Int64HistogramOption{
	metric.WithDescription("Distribution of net I/O usage per invocation."),
	metric.WithUnit("By"),
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

func (m NetIO) RecordSet(ctx context.Context, val int64, set attribute.Set) {
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

var newTimeoutsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of invocation timeouts."),
	metric.WithUnit("{timeout}"),
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

func (m Timeouts) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Timeouts) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type TimeoutsObservable struct {
	metric.Int64ObservableCounter
}

var newTimeoutsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Number of invocation timeouts."),
	metric.WithUnit("{timeout}"),
}

func NewTimeoutsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (TimeoutsObservable, error) {
	_ = "STUB: not implemented"
	return *new(TimeoutsObservable), nil
}

func (m TimeoutsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (TimeoutsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (TimeoutsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (TimeoutsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (TimeoutsObservable) AttrTrigger(val TriggerAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
