package goconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type CPUStateAttr string

var (
	CPUStateUser CPUStateAttr = "user"

	CPUStateGC CPUStateAttr = "gc"

	CPUStateScavenge CPUStateAttr = "scavenge"

	CPUStateIdle CPUStateAttr = "idle"
)

type MemoryTypeAttr string

var (
	MemoryTypeStack MemoryTypeAttr = "stack"

	MemoryTypeOther MemoryTypeAttr = "other"
)

type ConfigGogc struct {
	metric.Int64ObservableUpDownCounter
}

var newConfigGogcOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Heap size target percentage configured by the user, otherwise 100."),
	metric.WithUnit("%"),
}

func NewConfigGogc(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ConfigGogc, error) {
	_ = "STUB: not implemented"
	return *new(ConfigGogc), nil
}

func (m ConfigGogc) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ConfigGogc) Name() string { _ = "STUB: not implemented"; return "" }

func (ConfigGogc) Unit() string { _ = "STUB: not implemented"; return "" }

func (ConfigGogc) Description() string { _ = "STUB: not implemented"; return "" }

type CPUTime struct {
	metric.Float64Counter
}

var newCPUTimeOpts = []metric.Float64CounterOption{
	metric.WithDescription("Estimated CPU time spent by the Go runtime."),
	metric.WithUnit("s"),
}

func NewCPUTime(
	m metric.Meter,
	opt ...metric.Float64CounterOption,
) (CPUTime, error) {
	_ = "STUB: not implemented"
	return *new(CPUTime), nil
}

func (m CPUTime) Inst() metric.Float64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Float64Counter)
}

func (CPUTime) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m CPUTime) Add(
	ctx context.Context,
	incr float64,
	cpuState CPUStateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CPUTime) AddSet(ctx context.Context, incr float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (CPUTime) AttrCPUDetailedState(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUTimeObservable struct {
	metric.Float64ObservableCounter
}

var newCPUTimeObservableOpts = []metric.Float64ObservableCounterOption{
	metric.WithDescription("Estimated CPU time spent by the Go runtime."),
	metric.WithUnit("s"),
}

func NewCPUTimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableCounterOption,
) (CPUTimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(CPUTimeObservable), nil
}

func (m CPUTimeObservable) Inst() metric.Float64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter)
}

func (CPUTimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUTimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUTimeObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (CPUTimeObservable) AttrCPUState(val CPUStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUTimeObservable) AttrCPUDetailedState(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type GoroutineCount struct {
	metric.Int64ObservableUpDownCounter
}

var newGoroutineCountOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Count of live goroutines."),
	metric.WithUnit("{goroutine}"),
}

func NewGoroutineCount(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (GoroutineCount, error) {
	_ = "STUB: not implemented"
	return *new(GoroutineCount), nil
}

func (m GoroutineCount) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (GoroutineCount) Name() string { _ = "STUB: not implemented"; return "" }

func (GoroutineCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (GoroutineCount) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryAllocated struct {
	metric.Int64ObservableCounter
}

var newMemoryAllocatedOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Memory allocated to the heap by the application."),
	metric.WithUnit("By"),
}

func NewMemoryAllocated(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (MemoryAllocated, error) {
	_ = "STUB: not implemented"
	return *new(MemoryAllocated), nil
}

func (m MemoryAllocated) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (MemoryAllocated) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryAllocated) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryAllocated) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryAllocations struct {
	metric.Int64ObservableCounter
}

var newMemoryAllocationsOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Count of allocations to the heap by the application."),
	metric.WithUnit("{allocation}"),
}

func NewMemoryAllocations(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (MemoryAllocations, error) {
	_ = "STUB: not implemented"
	return *new(MemoryAllocations), nil
}

func (m MemoryAllocations) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (MemoryAllocations) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryAllocations) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryAllocations) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryGCCycles struct {
	metric.Int64Counter
}

var newMemoryGCCyclesOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of completed GC cycles."),
	metric.WithUnit("{gc_cycle}"),
}

func NewMemoryGCCycles(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (MemoryGCCycles, error) {
	_ = "STUB: not implemented"
	return *new(MemoryGCCycles), nil
}

func (m MemoryGCCycles) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (MemoryGCCycles) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCCycles) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCCycles) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryGCCycles) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryGCCycles) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryGCCyclesObservable struct {
	metric.Int64ObservableCounter
}

var newMemoryGCCyclesObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Number of completed GC cycles."),
	metric.WithUnit("{gc_cycle}"),
}

func NewMemoryGCCyclesObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (MemoryGCCyclesObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryGCCyclesObservable), nil
}

func (m MemoryGCCyclesObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (MemoryGCCyclesObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCCyclesObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCCyclesObservable) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryGCGoal struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryGCGoalOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Heap size target for the end of the GC cycle."),
	metric.WithUnit("By"),
}

func NewMemoryGCGoal(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryGCGoal, error) {
	_ = "STUB: not implemented"
	return *new(MemoryGCGoal), nil
}

func (m MemoryGCGoal) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryGCGoal) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCGoal) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCGoal) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryGCPauseDuration struct {
	metric.Float64Histogram
}

var newMemoryGCPauseDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Distribution of individual GC-related stop-the-world pause latencies. This is the time from deciding to stop the world until the world is started again."),
	metric.WithUnit("s"),
}

func NewMemoryGCPauseDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (MemoryGCPauseDuration, error) {
	_ = "STUB: not implemented"
	return *new(MemoryGCPauseDuration), nil
}

func (m MemoryGCPauseDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (MemoryGCPauseDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCPauseDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCPauseDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryGCPauseDuration) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryGCPauseDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryLimit struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLimitOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Go runtime memory limit configured by the user, if a limit exists."),
	metric.WithUnit("By"),
}

func NewMemoryLimit(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLimit, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLimit), nil
}

func (m MemoryLimit) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLimit) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryUsed struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryUsedOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Memory used by the Go runtime."),
	metric.WithUnit("By"),
}

func NewMemoryUsed(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryUsed, error) {
	_ = "STUB: not implemented"
	return *new(MemoryUsed), nil
}

func (m MemoryUsed) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsed) AttrMemoryType(val MemoryTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemoryUsed) AttrMemoryDetailedType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ProcessorLimit struct {
	metric.Int64ObservableUpDownCounter
}

var newProcessorLimitOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of OS threads that can execute user-level Go code simultaneously."),
	metric.WithUnit("{thread}"),
}

func NewProcessorLimit(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ProcessorLimit, error) {
	_ = "STUB: not implemented"
	return *new(ProcessorLimit), nil
}

func (m ProcessorLimit) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ProcessorLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (ProcessorLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (ProcessorLimit) Description() string { _ = "STUB: not implemented"; return "" }

type ScheduleDuration struct {
	metric.Float64Histogram
}

var newScheduleDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The time goroutines have spent in the scheduler in a runnable state before actually running."),
	metric.WithUnit("s"),
}

func NewScheduleDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ScheduleDuration, error) {
	_ = "STUB: not implemented"
	return *new(ScheduleDuration), nil
}

func (m ScheduleDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ScheduleDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ScheduleDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ScheduleDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ScheduleDuration) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ScheduleDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}
