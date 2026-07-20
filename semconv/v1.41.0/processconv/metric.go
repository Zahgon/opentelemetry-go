package processconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type CPUModeAttr string

var (
	CPUModeUser CPUModeAttr = "user"

	CPUModeSystem CPUModeAttr = "system"

	CPUModeNice CPUModeAttr = "nice"

	CPUModeIdle CPUModeAttr = "idle"

	CPUModeIOWait CPUModeAttr = "iowait"

	CPUModeInterrupt CPUModeAttr = "interrupt"

	CPUModeSteal CPUModeAttr = "steal"

	CPUModeKernel CPUModeAttr = "kernel"
)

type DiskIODirectionAttr string

var (
	DiskIODirectionRead DiskIODirectionAttr = "read"

	DiskIODirectionWrite DiskIODirectionAttr = "write"
)

type NetworkIODirectionAttr string

var (
	NetworkIODirectionTransmit NetworkIODirectionAttr = "transmit"

	NetworkIODirectionReceive NetworkIODirectionAttr = "receive"
)

type ContextSwitchTypeAttr string

var (
	ContextSwitchTypeVoluntary ContextSwitchTypeAttr = "voluntary"

	ContextSwitchTypeInvoluntary ContextSwitchTypeAttr = "involuntary"
)

type SystemPagingFaultTypeAttr string

var (
	SystemPagingFaultTypeMajor SystemPagingFaultTypeAttr = "major"

	SystemPagingFaultTypeMinor SystemPagingFaultTypeAttr = "minor"
)

type ContextSwitches struct {
	metric.Int64Counter
}

var newContextSwitchesOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of times the process has been context switched."),
	metric.WithUnit("{context_switch}"),
}

func NewContextSwitches(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ContextSwitches, error) {
	_ = "STUB: not implemented"
	return *new(ContextSwitches), nil
}

func (m ContextSwitches) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ContextSwitches) Name() string { _ = "STUB: not implemented"; return "" }

func (ContextSwitches) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContextSwitches) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContextSwitches) Add(
	ctx context.Context,
	incr int64,
	contextSwitchType ContextSwitchTypeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ContextSwitches) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContextSwitchesObservable struct {
	metric.Int64ObservableCounter
}

var newContextSwitchesObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Number of times the process has been context switched."),
	metric.WithUnit("{context_switch}"),
}

func NewContextSwitchesObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ContextSwitchesObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContextSwitchesObservable), nil
}

func (m ContextSwitchesObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (ContextSwitchesObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContextSwitchesObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContextSwitchesObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ContextSwitchesObservable) AttrContextSwitchType(val ContextSwitchTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUTime struct {
	metric.Float64ObservableCounter
}

var newCPUTimeOpts = []metric.Float64ObservableCounterOption{
	metric.WithDescription("Total CPU seconds broken down by different CPU modes."),
	metric.WithUnit("s"),
}

func NewCPUTime(
	m metric.Meter,
	opt ...metric.Float64ObservableCounterOption,
) (CPUTime, error) {
	_ = "STUB: not implemented"
	return *new(CPUTime), nil
}

func (m CPUTime) Inst() metric.Float64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter)
}

func (CPUTime) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUTime) Description() string { _ = "STUB: not implemented"; return "" }

func (CPUTime) AttrCPUMode(val CPUModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUUtilization struct {
	metric.Int64Gauge
}

var newCPUUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Difference in process.cpu.time since the last measurement, divided by the elapsed time and number of CPUs available to the process."),
	metric.WithUnit("1"),
}

func NewCPUUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (CPUUtilization, error) {
	_ = "STUB: not implemented"
	return *new(CPUUtilization), nil
}

func (m CPUUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (CPUUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m CPUUtilization) Record(
	ctx context.Context,
	val int64,
	cpuMode CPUModeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CPUUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type CPUUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newCPUUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Difference in process.cpu.time since the last measurement, divided by the elapsed time and number of CPUs available to the process."),
	metric.WithUnit("1"),
}

func NewCPUUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (CPUUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(CPUUtilizationObservable), nil
}

func (m CPUUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (CPUUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUUtilizationObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (CPUUtilizationObservable) AttrCPUMode(val CPUModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskIO struct {
	metric.Int64Counter
}

var newDiskIOOpts = []metric.Int64CounterOption{
	metric.WithDescription("Disk bytes transferred."),
	metric.WithUnit("By"),
}

func NewDiskIO(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (DiskIO, error) {
	_ = "STUB: not implemented"
	return *new(DiskIO), nil
}

func (m DiskIO) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (DiskIO) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskIO) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskIO) Description() string { _ = "STUB: not implemented"; return "" }

func (m DiskIO) Add(
	ctx context.Context,
	incr int64,
	diskIoDirection DiskIODirectionAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m DiskIO) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DiskIOObservable struct {
	metric.Int64ObservableCounter
}

var newDiskIOObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Disk bytes transferred."),
	metric.WithUnit("By"),
}

func NewDiskIOObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (DiskIOObservable, error) {
	_ = "STUB: not implemented"
	return *new(DiskIOObservable), nil
}

func (m DiskIOObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (DiskIOObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskIOObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskIOObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (DiskIOObservable) AttrDiskIODirection(val DiskIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryUsage struct {
	metric.Int64UpDownCounter
}

var newMemoryUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The amount of physical memory in use."),
	metric.WithUnit("By"),
}

func NewMemoryUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryUsage, error) {
	_ = "STUB: not implemented"
	return *new(MemoryUsage), nil
}

func (m MemoryUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryUsage) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The amount of physical memory in use."),
	metric.WithUnit("By"),
}

func NewMemoryUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryUsageObservable), nil
}

func (m MemoryUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryVirtual struct {
	metric.Int64UpDownCounter
}

var newMemoryVirtualOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The amount of committed virtual memory."),
	metric.WithUnit("By"),
}

func NewMemoryVirtual(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryVirtual, error) {
	_ = "STUB: not implemented"
	return *new(MemoryVirtual), nil
}

func (m MemoryVirtual) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryVirtual) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryVirtual) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryVirtual) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryVirtual) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryVirtual) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryVirtualObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryVirtualObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The amount of committed virtual memory."),
	metric.WithUnit("By"),
}

func NewMemoryVirtualObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryVirtualObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryVirtualObservable), nil
}

func (m MemoryVirtualObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryVirtualObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryVirtualObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryVirtualObservable) Description() string { _ = "STUB: not implemented"; return "" }

type NetworkIO struct {
	metric.Int64Counter
}

var newNetworkIOOpts = []metric.Int64CounterOption{
	metric.WithDescription("Network bytes transferred."),
	metric.WithUnit("By"),
}

func NewNetworkIO(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (NetworkIO, error) {
	_ = "STUB: not implemented"
	return *new(NetworkIO), nil
}

func (m NetworkIO) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (NetworkIO) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkIO) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkIO) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkIO) Add(
	ctx context.Context,
	incr int64,
	networkIoDirection NetworkIODirectionAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkIO) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NetworkIOObservable struct {
	metric.Int64ObservableCounter
}

var newNetworkIOObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Network bytes transferred."),
	metric.WithUnit("By"),
}

func NewNetworkIOObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (NetworkIOObservable, error) {
	_ = "STUB: not implemented"
	return *new(NetworkIOObservable), nil
}

func (m NetworkIOObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (NetworkIOObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkIOObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkIOObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NetworkIOObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingFaults struct {
	metric.Int64Counter
}

var newPagingFaultsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of page faults the process has made."),
	metric.WithUnit("{fault}"),
}

func NewPagingFaults(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (PagingFaults, error) {
	_ = "STUB: not implemented"
	return *new(PagingFaults), nil
}

func (m PagingFaults) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (PagingFaults) Name() string { _ = "STUB: not implemented"; return "" }

func (PagingFaults) Unit() string { _ = "STUB: not implemented"; return "" }

func (PagingFaults) Description() string { _ = "STUB: not implemented"; return "" }

func (m PagingFaults) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PagingFaults) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PagingFaults) AttrSystemPagingFaultType(val SystemPagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingFaultsObservable struct {
	metric.Int64ObservableCounter
}

var newPagingFaultsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Number of page faults the process has made."),
	metric.WithUnit("{fault}"),
}

func NewPagingFaultsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (PagingFaultsObservable, error) {
	_ = "STUB: not implemented"
	return *new(PagingFaultsObservable), nil
}

func (m PagingFaultsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (PagingFaultsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PagingFaultsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PagingFaultsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PagingFaultsObservable) AttrSystemPagingFaultType(val SystemPagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ThreadCount struct {
	metric.Int64UpDownCounter
}

var newThreadCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Process threads count."),
	metric.WithUnit("{thread}"),
}

func NewThreadCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ThreadCount, error) {
	_ = "STUB: not implemented"
	return *new(ThreadCount), nil
}

func (m ThreadCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ThreadCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ThreadCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ThreadCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ThreadCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ThreadCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ThreadCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newThreadCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Process threads count."),
	metric.WithUnit("{thread}"),
}

func NewThreadCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ThreadCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(ThreadCountObservable), nil
}

func (m ThreadCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ThreadCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ThreadCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ThreadCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

type UnixFileDescriptorCount struct {
	metric.Int64UpDownCounter
}

var newUnixFileDescriptorCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of unix file descriptors in use by the process."),
	metric.WithUnit("{file_descriptor}"),
}

func NewUnixFileDescriptorCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (UnixFileDescriptorCount, error) {
	_ = "STUB: not implemented"
	return *new(UnixFileDescriptorCount), nil
}

func (m UnixFileDescriptorCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (UnixFileDescriptorCount) Name() string { _ = "STUB: not implemented"; return "" }

func (UnixFileDescriptorCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (UnixFileDescriptorCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m UnixFileDescriptorCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m UnixFileDescriptorCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type UnixFileDescriptorCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newUnixFileDescriptorCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of unix file descriptors in use by the process."),
	metric.WithUnit("{file_descriptor}"),
}

func NewUnixFileDescriptorCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (UnixFileDescriptorCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(UnixFileDescriptorCountObservable), nil
}

func (m UnixFileDescriptorCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (UnixFileDescriptorCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (UnixFileDescriptorCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (UnixFileDescriptorCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

type Uptime struct {
	metric.Float64Gauge
}

var newUptimeOpts = []metric.Float64GaugeOption{
	metric.WithDescription("The time the process has been running."),
	metric.WithUnit("s"),
}

func NewUptime(
	m metric.Meter,
	opt ...metric.Float64GaugeOption,
) (Uptime, error) {
	_ = "STUB: not implemented"
	return *new(Uptime), nil
}

func (m Uptime) Inst() metric.Float64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge)
}

func (Uptime) Name() string { _ = "STUB: not implemented"; return "" }

func (Uptime) Unit() string { _ = "STUB: not implemented"; return "" }

func (Uptime) Description() string { _ = "STUB: not implemented"; return "" }

func (m Uptime) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m Uptime) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type UptimeObservable struct {
	metric.Float64ObservableGauge
}

var newUptimeObservableOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("The time the process has been running."),
	metric.WithUnit("s"),
}

func NewUptimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (UptimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(UptimeObservable), nil
}

func (m UptimeObservable) Inst() metric.Float64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge)
}

func (UptimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (UptimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (UptimeObservable) Description() string { _ = "STUB: not implemented"; return "" }

type WindowsHandleCount struct {
	metric.Int64UpDownCounter
}

var newWindowsHandleCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of handles held by the process."),
	metric.WithUnit("{handle}"),
}

func NewWindowsHandleCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (WindowsHandleCount, error) {
	_ = "STUB: not implemented"
	return *new(WindowsHandleCount), nil
}

func (m WindowsHandleCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (WindowsHandleCount) Name() string { _ = "STUB: not implemented"; return "" }

func (WindowsHandleCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (WindowsHandleCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m WindowsHandleCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m WindowsHandleCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type WindowsHandleCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newWindowsHandleCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of handles held by the process."),
	metric.WithUnit("{handle}"),
}

func NewWindowsHandleCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (WindowsHandleCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(WindowsHandleCountObservable), nil
}

func (m WindowsHandleCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (WindowsHandleCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (WindowsHandleCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (WindowsHandleCountObservable) Description() string { _ = "STUB: not implemented"; return "" }
