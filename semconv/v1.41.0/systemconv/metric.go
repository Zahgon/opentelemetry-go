package systemconv

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

type NetworkConnectionStateAttr string

var (
	NetworkConnectionStateClosed NetworkConnectionStateAttr = "closed"

	NetworkConnectionStateCloseWait NetworkConnectionStateAttr = "close_wait"

	NetworkConnectionStateClosing NetworkConnectionStateAttr = "closing"

	NetworkConnectionStateEstablished NetworkConnectionStateAttr = "established"

	NetworkConnectionStateFinWait1 NetworkConnectionStateAttr = "fin_wait_1"

	NetworkConnectionStateFinWait2 NetworkConnectionStateAttr = "fin_wait_2"

	NetworkConnectionStateLastAck NetworkConnectionStateAttr = "last_ack"

	NetworkConnectionStateListen NetworkConnectionStateAttr = "listen"

	NetworkConnectionStateSynReceived NetworkConnectionStateAttr = "syn_received"

	NetworkConnectionStateSynSent NetworkConnectionStateAttr = "syn_sent"

	NetworkConnectionStateTimeWait NetworkConnectionStateAttr = "time_wait"
)

type NetworkIODirectionAttr string

var (
	NetworkIODirectionTransmit NetworkIODirectionAttr = "transmit"

	NetworkIODirectionReceive NetworkIODirectionAttr = "receive"
)

type NetworkTransportAttr string

var (
	NetworkTransportTCP NetworkTransportAttr = "tcp"

	NetworkTransportUDP NetworkTransportAttr = "udp"

	NetworkTransportPipe NetworkTransportAttr = "pipe"

	NetworkTransportUnix NetworkTransportAttr = "unix"

	NetworkTransportQUIC NetworkTransportAttr = "quic"
)

type ProcessStateAttr string

var (
	ProcessStateRunning ProcessStateAttr = "running"

	ProcessStateSleeping ProcessStateAttr = "sleeping"

	ProcessStateStopped ProcessStateAttr = "stopped"

	ProcessStateDefunct ProcessStateAttr = "defunct"
)

type FilesystemStateAttr string

var (
	FilesystemStateUsed FilesystemStateAttr = "used"

	FilesystemStateFree FilesystemStateAttr = "free"

	FilesystemStateReserved FilesystemStateAttr = "reserved"
)

type FilesystemTypeAttr string

var (
	FilesystemTypeFat32 FilesystemTypeAttr = "fat32"

	FilesystemTypeExfat FilesystemTypeAttr = "exfat"

	FilesystemTypeNtfs FilesystemTypeAttr = "ntfs"

	FilesystemTypeRefs FilesystemTypeAttr = "refs"

	FilesystemTypeHfsplus FilesystemTypeAttr = "hfsplus"

	FilesystemTypeExt4 FilesystemTypeAttr = "ext4"
)

type MemoryLinuxHugepagesStateAttr string

var (
	MemoryLinuxHugepagesStateFree MemoryLinuxHugepagesStateAttr = "free"

	MemoryLinuxHugepagesStateUsed MemoryLinuxHugepagesStateAttr = "used"
)

type MemoryLinuxSlabStateAttr string

var (
	MemoryLinuxSlabStateReclaimable MemoryLinuxSlabStateAttr = "reclaimable"

	MemoryLinuxSlabStateUnreclaimable MemoryLinuxSlabStateAttr = "unreclaimable"
)

type MemoryStateAttr string

var (
	MemoryStateUsed MemoryStateAttr = "used"

	MemoryStateFree MemoryStateAttr = "free"

	MemoryStateBuffers MemoryStateAttr = "buffers"

	MemoryStateCached MemoryStateAttr = "cached"
)

type PagingDirectionAttr string

var (
	PagingDirectionIn PagingDirectionAttr = "in"

	PagingDirectionOut PagingDirectionAttr = "out"
)

type PagingFaultTypeAttr string

var (
	PagingFaultTypeMajor PagingFaultTypeAttr = "major"

	PagingFaultTypeMinor PagingFaultTypeAttr = "minor"
)

type PagingStateAttr string

var (
	PagingStateUsed PagingStateAttr = "used"

	PagingStateFree PagingStateAttr = "free"
)

type CPUFrequency struct {
	metric.Int64Gauge
}

var newCPUFrequencyOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Operating frequency of the logical CPU in Hertz."),
	metric.WithUnit("Hz"),
}

func NewCPUFrequency(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (CPUFrequency, error) {
	_ = "STUB: not implemented"
	return *new(CPUFrequency), nil
}

func (m CPUFrequency) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (CPUFrequency) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUFrequency) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUFrequency) Description() string { _ = "STUB: not implemented"; return "" }

func (m CPUFrequency) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CPUFrequency) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (CPUFrequency) AttrCPULogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUFrequencyObservable struct {
	metric.Int64ObservableGauge
}

var newCPUFrequencyObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Operating frequency of the logical CPU in Hertz."),
	metric.WithUnit("Hz"),
}

func NewCPUFrequencyObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (CPUFrequencyObservable, error) {
	_ = "STUB: not implemented"
	return *new(CPUFrequencyObservable), nil
}

func (m CPUFrequencyObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (CPUFrequencyObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUFrequencyObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUFrequencyObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (CPUFrequencyObservable) AttrCPULogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPULogicalCount struct {
	metric.Int64UpDownCounter
}

var newCPULogicalCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Reports the number of logical (virtual) processor cores created by the operating system to manage multitasking."),
	metric.WithUnit("{cpu}"),
}

func NewCPULogicalCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (CPULogicalCount, error) {
	_ = "STUB: not implemented"
	return *new(CPULogicalCount), nil
}

func (m CPULogicalCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (CPULogicalCount) Name() string { _ = "STUB: not implemented"; return "" }

func (CPULogicalCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPULogicalCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m CPULogicalCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m CPULogicalCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type CPULogicalCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newCPULogicalCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Reports the number of logical (virtual) processor cores created by the operating system to manage multitasking."),
	metric.WithUnit("{cpu}"),
}

func NewCPULogicalCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (CPULogicalCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(CPULogicalCountObservable), nil
}

func (m CPULogicalCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (CPULogicalCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (CPULogicalCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPULogicalCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

type CPUPhysicalCount struct {
	metric.Int64UpDownCounter
}

var newCPUPhysicalCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Reports the number of actual physical processor cores on the hardware."),
	metric.WithUnit("{cpu}"),
}

func NewCPUPhysicalCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (CPUPhysicalCount, error) {
	_ = "STUB: not implemented"
	return *new(CPUPhysicalCount), nil
}

func (m CPUPhysicalCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (CPUPhysicalCount) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUPhysicalCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUPhysicalCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m CPUPhysicalCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m CPUPhysicalCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type CPUPhysicalCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newCPUPhysicalCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Reports the number of actual physical processor cores on the hardware."),
	metric.WithUnit("{cpu}"),
}

func NewCPUPhysicalCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (CPUPhysicalCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(CPUPhysicalCountObservable), nil
}

func (m CPUPhysicalCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (CPUPhysicalCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUPhysicalCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUPhysicalCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

type CPUTime struct {
	metric.Float64ObservableCounter
}

var newCPUTimeOpts = []metric.Float64ObservableCounterOption{
	metric.WithDescription("Seconds each logical CPU spent on each mode."),
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

func (CPUTime) AttrCPULogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUUtilization struct {
	metric.Int64Gauge
}

var newCPUUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("For each logical CPU, the utilization is calculated as the change in cumulative CPU time (cpu.time) over a measurement interval, divided by the elapsed time."),
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
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CPUUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (CPUUtilization) AttrCPUMode(val CPUModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUUtilization) AttrCPULogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newCPUUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("For each logical CPU, the utilization is calculated as the change in cumulative CPU time (cpu.time) over a measurement interval, divided by the elapsed time."),
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

func (CPUUtilizationObservable) AttrCPULogicalNumber(val int) attribute.KeyValue {
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
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m DiskIO) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (DiskIO) AttrDiskIODirection(val DiskIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (DiskIO) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

func (DiskIOObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskIOTime struct {
	metric.Float64Counter
}

var newDiskIOTimeOpts = []metric.Float64CounterOption{
	metric.WithDescription("Time disk spent activated."),
	metric.WithUnit("s"),
}

func NewDiskIOTime(
	m metric.Meter,
	opt ...metric.Float64CounterOption,
) (DiskIOTime, error) {
	_ = "STUB: not implemented"
	return *new(DiskIOTime), nil
}

func (m DiskIOTime) Inst() metric.Float64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Float64Counter)
}

func (DiskIOTime) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskIOTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskIOTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m DiskIOTime) Add(
	ctx context.Context,
	incr float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m DiskIOTime) AddSet(ctx context.Context, incr float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (DiskIOTime) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskIOTimeObservable struct {
	metric.Float64ObservableCounter
}

var newDiskIOTimeObservableOpts = []metric.Float64ObservableCounterOption{
	metric.WithDescription("Time disk spent activated."),
	metric.WithUnit("s"),
}

func NewDiskIOTimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableCounterOption,
) (DiskIOTimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(DiskIOTimeObservable), nil
}

func (m DiskIOTimeObservable) Inst() metric.Float64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter)
}

func (DiskIOTimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskIOTimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskIOTimeObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (DiskIOTimeObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskLimit struct {
	metric.Int64UpDownCounter
}

var newDiskLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The total storage capacity of the disk."),
	metric.WithUnit("By"),
}

func NewDiskLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DiskLimit, error) {
	_ = "STUB: not implemented"
	return *new(DiskLimit), nil
}

func (m DiskLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DiskLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m DiskLimit) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m DiskLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (DiskLimit) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newDiskLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The total storage capacity of the disk."),
	metric.WithUnit("By"),
}

func NewDiskLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (DiskLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(DiskLimitObservable), nil
}

func (m DiskLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (DiskLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (DiskLimitObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskMerged struct {
	metric.Int64Counter
}

var newDiskMergedOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of disk reads/writes merged into single physical disk access operations."),
	metric.WithUnit("{operation}"),
}

func NewDiskMerged(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (DiskMerged, error) {
	_ = "STUB: not implemented"
	return *new(DiskMerged), nil
}

func (m DiskMerged) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (DiskMerged) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskMerged) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskMerged) Description() string { _ = "STUB: not implemented"; return "" }

func (m DiskMerged) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m DiskMerged) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (DiskMerged) AttrDiskIODirection(val DiskIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (DiskMerged) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskMergedObservable struct {
	metric.Int64ObservableCounter
}

var newDiskMergedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of disk reads/writes merged into single physical disk access operations."),
	metric.WithUnit("{operation}"),
}

func NewDiskMergedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (DiskMergedObservable, error) {
	_ = "STUB: not implemented"
	return *new(DiskMergedObservable), nil
}

func (m DiskMergedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (DiskMergedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskMergedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskMergedObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (DiskMergedObservable) AttrDiskIODirection(val DiskIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (DiskMergedObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskOperationTime struct {
	metric.Float64Counter
}

var newDiskOperationTimeOpts = []metric.Float64CounterOption{
	metric.WithDescription("Sum of the time each operation took to complete."),
	metric.WithUnit("s"),
}

func NewDiskOperationTime(
	m metric.Meter,
	opt ...metric.Float64CounterOption,
) (DiskOperationTime, error) {
	_ = "STUB: not implemented"
	return *new(DiskOperationTime), nil
}

func (m DiskOperationTime) Inst() metric.Float64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Float64Counter)
}

func (DiskOperationTime) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskOperationTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskOperationTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m DiskOperationTime) Add(
	ctx context.Context,
	incr float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m DiskOperationTime) AddSet(ctx context.Context, incr float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (DiskOperationTime) AttrDiskIODirection(val DiskIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (DiskOperationTime) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskOperationTimeObservable struct {
	metric.Float64ObservableCounter
}

var newDiskOperationTimeObservableOpts = []metric.Float64ObservableCounterOption{
	metric.WithDescription("Sum of the time each operation took to complete."),
	metric.WithUnit("s"),
}

func NewDiskOperationTimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableCounterOption,
) (DiskOperationTimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(DiskOperationTimeObservable), nil
}

func (m DiskOperationTimeObservable) Inst() metric.Float64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter)
}

func (DiskOperationTimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskOperationTimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskOperationTimeObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (DiskOperationTimeObservable) AttrDiskIODirection(val DiskIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (DiskOperationTimeObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskOperations struct {
	metric.Int64Counter
}

var newDiskOperationsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Disk operations count."),
	metric.WithUnit("{operation}"),
}

func NewDiskOperations(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (DiskOperations, error) {
	_ = "STUB: not implemented"
	return *new(DiskOperations), nil
}

func (m DiskOperations) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (DiskOperations) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskOperations) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskOperations) Description() string { _ = "STUB: not implemented"; return "" }

func (m DiskOperations) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m DiskOperations) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (DiskOperations) AttrDiskIODirection(val DiskIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (DiskOperations) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskOperationsObservable struct {
	metric.Int64ObservableCounter
}

var newDiskOperationsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Disk operations count."),
	metric.WithUnit("{operation}"),
}

func NewDiskOperationsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (DiskOperationsObservable, error) {
	_ = "STUB: not implemented"
	return *new(DiskOperationsObservable), nil
}

func (m DiskOperationsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (DiskOperationsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DiskOperationsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DiskOperationsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (DiskOperationsObservable) AttrDiskIODirection(val DiskIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (DiskOperationsObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FilesystemLimit struct {
	metric.Int64UpDownCounter
}

var newFilesystemLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The total storage capacity of the filesystem."),
	metric.WithUnit("By"),
}

func NewFilesystemLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (FilesystemLimit, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemLimit), nil
}

func (m FilesystemLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (FilesystemLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m FilesystemLimit) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m FilesystemLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (FilesystemLimit) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemLimit) AttrFilesystemMode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemLimit) AttrFilesystemMountpoint(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemLimit) AttrFilesystemType(val FilesystemTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FilesystemLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newFilesystemLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The total storage capacity of the filesystem."),
	metric.WithUnit("By"),
}

func NewFilesystemLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (FilesystemLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemLimitObservable), nil
}

func (m FilesystemLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (FilesystemLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (FilesystemLimitObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemLimitObservable) AttrFilesystemMode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemLimitObservable) AttrFilesystemMountpoint(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemLimitObservable) AttrFilesystemType(val FilesystemTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FilesystemUsage struct {
	metric.Int64UpDownCounter
}

var newFilesystemUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Reports a filesystem's space usage across different states."),
	metric.WithUnit("By"),
}

func NewFilesystemUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (FilesystemUsage, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemUsage), nil
}

func (m FilesystemUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (FilesystemUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m FilesystemUsage) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m FilesystemUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (FilesystemUsage) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUsage) AttrFilesystemMode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUsage) AttrFilesystemMountpoint(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUsage) AttrFilesystemState(val FilesystemStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUsage) AttrFilesystemType(val FilesystemTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FilesystemUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newFilesystemUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Reports a filesystem's space usage across different states."),
	metric.WithUnit("By"),
}

func NewFilesystemUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (FilesystemUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemUsageObservable), nil
}

func (m FilesystemUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (FilesystemUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUsageObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUsageObservable) AttrFilesystemMode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUsageObservable) AttrFilesystemMountpoint(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUsageObservable) AttrFilesystemState(val FilesystemStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUsageObservable) AttrFilesystemType(val FilesystemTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FilesystemUtilization struct {
	metric.Int64Gauge
}

var newFilesystemUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Fraction of filesystem bytes used."),
	metric.WithUnit("1"),
}

func NewFilesystemUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (FilesystemUtilization, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemUtilization), nil
}

func (m FilesystemUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (FilesystemUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m FilesystemUtilization) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m FilesystemUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (FilesystemUtilization) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUtilization) AttrFilesystemMode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUtilization) AttrFilesystemMountpoint(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUtilization) AttrFilesystemState(val FilesystemStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUtilization) AttrFilesystemType(val FilesystemTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FilesystemUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newFilesystemUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Fraction of filesystem bytes used."),
	metric.WithUnit("1"),
}

func NewFilesystemUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (FilesystemUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemUtilizationObservable), nil
}

func (m FilesystemUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (FilesystemUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUtilizationObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (FilesystemUtilizationObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUtilizationObservable) AttrFilesystemMode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUtilizationObservable) AttrFilesystemMountpoint(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUtilizationObservable) AttrFilesystemState(val FilesystemStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FilesystemUtilizationObservable) AttrFilesystemType(val FilesystemTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryLimit struct {
	metric.Int64UpDownCounter
}

var newMemoryLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Total virtual memory available in the system."),
	metric.WithUnit("By"),
}

func NewMemoryLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryLimit, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLimit), nil
}

func (m MemoryLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLimit) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Total virtual memory available in the system."),
	metric.WithUnit("By"),
}

func NewMemoryLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLimitObservable), nil
}

func (m MemoryLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryLinuxAvailable struct {
	metric.Int64UpDownCounter
}

var newMemoryLinuxAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("An estimate of how much memory is available for starting new applications, without causing swapping."),
	metric.WithUnit("By"),
}

func NewMemoryLinuxAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryLinuxAvailable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxAvailable), nil
}

func (m MemoryLinuxAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryLinuxAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLinuxAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLinuxAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryLinuxAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLinuxAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("An estimate of how much memory is available for starting new applications, without causing swapping."),
	metric.WithUnit("By"),
}

func NewMemoryLinuxAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLinuxAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxAvailableObservable), nil
}

func (m MemoryLinuxAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLinuxAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryLinuxHugepagesLimit struct {
	metric.Int64UpDownCounter
}

var newMemoryLinuxHugepagesLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Total number of hugepages available."),
	metric.WithUnit("{page}"),
}

func NewMemoryLinuxHugepagesLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryLinuxHugepagesLimit, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesLimit), nil
}

func (m MemoryLinuxHugepagesLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryLinuxHugepagesLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLinuxHugepagesLimit) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLinuxHugepagesLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryLinuxHugepagesLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLinuxHugepagesLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Total number of hugepages available."),
	metric.WithUnit("{page}"),
}

func NewMemoryLinuxHugepagesLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLinuxHugepagesLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesLimitObservable), nil
}

func (m MemoryLinuxHugepagesLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLinuxHugepagesLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesLimitObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type MemoryLinuxHugepagesPageSize struct {
	metric.Int64UpDownCounter
}

var newMemoryLinuxHugepagesPageSizeOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("System hugepage size in bytes."),
	metric.WithUnit("By"),
}

func NewMemoryLinuxHugepagesPageSize(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryLinuxHugepagesPageSize, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesPageSize), nil
}

func (m MemoryLinuxHugepagesPageSize) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryLinuxHugepagesPageSize) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesPageSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesPageSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLinuxHugepagesPageSize) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLinuxHugepagesPageSize) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryLinuxHugepagesPageSizeObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLinuxHugepagesPageSizeObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("System hugepage size in bytes."),
	metric.WithUnit("By"),
}

func NewMemoryLinuxHugepagesPageSizeObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLinuxHugepagesPageSizeObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesPageSizeObservable), nil
}

func (m MemoryLinuxHugepagesPageSizeObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLinuxHugepagesPageSizeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesPageSizeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesPageSizeObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type MemoryLinuxHugepagesReserved struct {
	metric.Int64UpDownCounter
}

var newMemoryLinuxHugepagesReservedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of reserved hugepages."),
	metric.WithUnit("{page}"),
}

func NewMemoryLinuxHugepagesReserved(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryLinuxHugepagesReserved, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesReserved), nil
}

func (m MemoryLinuxHugepagesReserved) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryLinuxHugepagesReserved) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesReserved) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesReserved) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLinuxHugepagesReserved) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLinuxHugepagesReserved) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryLinuxHugepagesReservedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLinuxHugepagesReservedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of reserved hugepages."),
	metric.WithUnit("{page}"),
}

func NewMemoryLinuxHugepagesReservedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLinuxHugepagesReservedObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesReservedObservable), nil
}

func (m MemoryLinuxHugepagesReservedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLinuxHugepagesReservedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesReservedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesReservedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type MemoryLinuxHugepagesSurplus struct {
	metric.Int64UpDownCounter
}

var newMemoryLinuxHugepagesSurplusOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of surplus hugepages."),
	metric.WithUnit("{page}"),
}

func NewMemoryLinuxHugepagesSurplus(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryLinuxHugepagesSurplus, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesSurplus), nil
}

func (m MemoryLinuxHugepagesSurplus) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryLinuxHugepagesSurplus) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesSurplus) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesSurplus) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLinuxHugepagesSurplus) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLinuxHugepagesSurplus) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryLinuxHugepagesSurplusObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLinuxHugepagesSurplusObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of surplus hugepages."),
	metric.WithUnit("{page}"),
}

func NewMemoryLinuxHugepagesSurplusObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLinuxHugepagesSurplusObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesSurplusObservable), nil
}

func (m MemoryLinuxHugepagesSurplusObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLinuxHugepagesSurplusObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesSurplusObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesSurplusObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type MemoryLinuxHugepagesUsage struct {
	metric.Int64UpDownCounter
}

var newMemoryLinuxHugepagesUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of hugepages in use by state."),
	metric.WithUnit("{page}"),
}

func NewMemoryLinuxHugepagesUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryLinuxHugepagesUsage, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesUsage), nil
}

func (m MemoryLinuxHugepagesUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryLinuxHugepagesUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLinuxHugepagesUsage) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLinuxHugepagesUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (MemoryLinuxHugepagesUsage) AttrMemoryLinuxHugepagesState(val MemoryLinuxHugepagesStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryLinuxHugepagesUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLinuxHugepagesUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of hugepages in use by state."),
	metric.WithUnit("{page}"),
}

func NewMemoryLinuxHugepagesUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLinuxHugepagesUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesUsageObservable), nil
}

func (m MemoryLinuxHugepagesUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLinuxHugepagesUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesUsageObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (MemoryLinuxHugepagesUsageObservable) AttrMemoryLinuxHugepagesState(val MemoryLinuxHugepagesStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryLinuxHugepagesUtilization struct {
	metric.Int64Gauge
}

var newMemoryLinuxHugepagesUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Percentage of hugepages in use by state."),
	metric.WithUnit("1"),
}

func NewMemoryLinuxHugepagesUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (MemoryLinuxHugepagesUtilization, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesUtilization), nil
}

func (m MemoryLinuxHugepagesUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (MemoryLinuxHugepagesUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxHugepagesUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLinuxHugepagesUtilization) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLinuxHugepagesUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (MemoryLinuxHugepagesUtilization) AttrMemoryLinuxHugepagesState(val MemoryLinuxHugepagesStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryLinuxHugepagesUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newMemoryLinuxHugepagesUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Percentage of hugepages in use by state."),
	metric.WithUnit("1"),
}

func NewMemoryLinuxHugepagesUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (MemoryLinuxHugepagesUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxHugepagesUtilizationObservable), nil
}

func (m MemoryLinuxHugepagesUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (MemoryLinuxHugepagesUtilizationObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (MemoryLinuxHugepagesUtilizationObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (MemoryLinuxHugepagesUtilizationObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (MemoryLinuxHugepagesUtilizationObservable) AttrMemoryLinuxHugepagesState(val MemoryLinuxHugepagesStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryLinuxShared struct {
	metric.Int64UpDownCounter
}

var newMemoryLinuxSharedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Shared memory used (mostly by tmpfs)."),
	metric.WithUnit("By"),
}

func NewMemoryLinuxShared(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryLinuxShared, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxShared), nil
}

func (m MemoryLinuxShared) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryLinuxShared) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxShared) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxShared) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLinuxShared) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLinuxShared) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryLinuxSharedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLinuxSharedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Shared memory used (mostly by tmpfs)."),
	metric.WithUnit("By"),
}

func NewMemoryLinuxSharedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLinuxSharedObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxSharedObservable), nil
}

func (m MemoryLinuxSharedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLinuxSharedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxSharedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxSharedObservable) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryLinuxSlabUsage struct {
	metric.Int64UpDownCounter
}

var newMemoryLinuxSlabUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Reports the memory used by the Linux kernel for managing caches of frequently used objects."),
	metric.WithUnit("By"),
}

func NewMemoryLinuxSlabUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryLinuxSlabUsage, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxSlabUsage), nil
}

func (m MemoryLinuxSlabUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryLinuxSlabUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxSlabUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxSlabUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryLinuxSlabUsage) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryLinuxSlabUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (MemoryLinuxSlabUsage) AttrMemoryLinuxSlabState(val MemoryLinuxSlabStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryLinuxSlabUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryLinuxSlabUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Reports the memory used by the Linux kernel for managing caches of frequently used objects."),
	metric.WithUnit("By"),
}

func NewMemoryLinuxSlabUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLinuxSlabUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLinuxSlabUsageObservable), nil
}

func (m MemoryLinuxSlabUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLinuxSlabUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxSlabUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxSlabUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (MemoryLinuxSlabUsageObservable) AttrMemoryLinuxSlabState(val MemoryLinuxSlabStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryUsage struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryUsageOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Reports memory in use by state."),
	metric.WithUnit("By"),
}

func NewMemoryUsage(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryUsage, error) {
	_ = "STUB: not implemented"
	return *new(MemoryUsage), nil
}

func (m MemoryUsage) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsage) AttrMemoryState(val MemoryStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryUtilization struct {
	metric.Float64ObservableGauge
}

var newMemoryUtilizationOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("Percentage of memory bytes in use."),
	metric.WithUnit("1"),
}

func NewMemoryUtilization(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (MemoryUtilization, error) {
	_ = "STUB: not implemented"
	return *new(MemoryUtilization), nil
}

func (m MemoryUtilization) Inst() metric.Float64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge)
}

func (MemoryUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (MemoryUtilization) AttrMemoryState(val MemoryStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkConnectionCount struct {
	metric.Int64UpDownCounter
}

var newNetworkConnectionCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of connections."),
	metric.WithUnit("{connection}"),
}

func NewNetworkConnectionCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NetworkConnectionCount, error) {
	_ = "STUB: not implemented"
	return *new(NetworkConnectionCount), nil
}

func (m NetworkConnectionCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NetworkConnectionCount) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkConnectionCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkConnectionCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkConnectionCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkConnectionCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NetworkConnectionCount) AttrNetworkConnectionState(val NetworkConnectionStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkConnectionCount) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkConnectionCount) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkConnectionCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNetworkConnectionCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of connections."),
	metric.WithUnit("{connection}"),
}

func NewNetworkConnectionCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NetworkConnectionCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(NetworkConnectionCountObservable), nil
}

func (m NetworkConnectionCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NetworkConnectionCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkConnectionCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkConnectionCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NetworkConnectionCountObservable) AttrNetworkConnectionState(val NetworkConnectionStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkConnectionCountObservable) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkConnectionCountObservable) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkErrors struct {
	metric.Int64Counter
}

var newNetworkErrorsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Count of network errors detected."),
	metric.WithUnit("{error}"),
}

func NewNetworkErrors(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (NetworkErrors, error) {
	_ = "STUB: not implemented"
	return *new(NetworkErrors), nil
}

func (m NetworkErrors) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (NetworkErrors) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkErrors) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkErrors) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkErrors) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkErrors) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NetworkErrors) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkErrors) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkErrorsObservable struct {
	metric.Int64ObservableCounter
}

var newNetworkErrorsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Count of network errors detected."),
	metric.WithUnit("{error}"),
}

func NewNetworkErrorsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (NetworkErrorsObservable, error) {
	_ = "STUB: not implemented"
	return *new(NetworkErrorsObservable), nil
}

func (m NetworkErrorsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (NetworkErrorsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkErrorsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkErrorsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NetworkErrorsObservable) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkErrorsObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkIO struct {
	metric.Int64ObservableCounter
}

var newNetworkIOOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of bytes transmitted and received."),
	metric.WithUnit("By"),
}

func NewNetworkIO(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (NetworkIO, error) {
	_ = "STUB: not implemented"
	return *new(NetworkIO), nil
}

func (m NetworkIO) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (NetworkIO) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkIO) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkIO) Description() string { _ = "STUB: not implemented"; return "" }

func (NetworkIO) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIO) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkPacketCount struct {
	metric.Int64Counter
}

var newNetworkPacketCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of packets transferred."),
	metric.WithUnit("{packet}"),
}

func NewNetworkPacketCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (NetworkPacketCount, error) {
	_ = "STUB: not implemented"
	return *new(NetworkPacketCount), nil
}

func (m NetworkPacketCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (NetworkPacketCount) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkPacketCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkPacketCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NetworkPacketCount) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketCount) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkPacketCountObservable struct {
	metric.Int64ObservableCounter
}

var newNetworkPacketCountObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of packets transferred."),
	metric.WithUnit("{packet}"),
}

func NewNetworkPacketCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (NetworkPacketCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(NetworkPacketCountObservable), nil
}

func (m NetworkPacketCountObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (NetworkPacketCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketCountObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketCountObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkPacketDropped struct {
	metric.Int64Counter
}

var newNetworkPacketDroppedOpts = []metric.Int64CounterOption{
	metric.WithDescription("Count of packets that are dropped or discarded even though there was no error."),
	metric.WithUnit("{packet}"),
}

func NewNetworkPacketDropped(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (NetworkPacketDropped, error) {
	_ = "STUB: not implemented"
	return *new(NetworkPacketDropped), nil
}

func (m NetworkPacketDropped) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (NetworkPacketDropped) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketDropped) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketDropped) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkPacketDropped) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkPacketDropped) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NetworkPacketDropped) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketDropped) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkPacketDroppedObservable struct {
	metric.Int64ObservableCounter
}

var newNetworkPacketDroppedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Count of packets that are dropped or discarded even though there was no error."),
	metric.WithUnit("{packet}"),
}

func NewNetworkPacketDroppedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (NetworkPacketDroppedObservable, error) {
	_ = "STUB: not implemented"
	return *new(NetworkPacketDroppedObservable), nil
}

func (m NetworkPacketDroppedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (NetworkPacketDroppedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketDroppedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketDroppedObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketDroppedObservable) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketDroppedObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingFaults struct {
	metric.Int64Counter
}

var newPagingFaultsOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of page faults."),
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

func (PagingFaults) AttrPagingFaultType(val PagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingFaultsObservable struct {
	metric.Int64ObservableCounter
}

var newPagingFaultsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of page faults."),
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

func (PagingFaultsObservable) AttrPagingFaultType(val PagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingOperations struct {
	metric.Int64Counter
}

var newPagingOperationsOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of paging operations."),
	metric.WithUnit("{operation}"),
}

func NewPagingOperations(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (PagingOperations, error) {
	_ = "STUB: not implemented"
	return *new(PagingOperations), nil
}

func (m PagingOperations) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (PagingOperations) Name() string { _ = "STUB: not implemented"; return "" }

func (PagingOperations) Unit() string { _ = "STUB: not implemented"; return "" }

func (PagingOperations) Description() string { _ = "STUB: not implemented"; return "" }

func (m PagingOperations) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PagingOperations) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PagingOperations) AttrPagingDirection(val PagingDirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PagingOperations) AttrPagingFaultType(val PagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingOperationsObservable struct {
	metric.Int64ObservableCounter
}

var newPagingOperationsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of paging operations."),
	metric.WithUnit("{operation}"),
}

func NewPagingOperationsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (PagingOperationsObservable, error) {
	_ = "STUB: not implemented"
	return *new(PagingOperationsObservable), nil
}

func (m PagingOperationsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (PagingOperationsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PagingOperationsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PagingOperationsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PagingOperationsObservable) AttrPagingDirection(val PagingDirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PagingOperationsObservable) AttrPagingFaultType(val PagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingUsage struct {
	metric.Int64UpDownCounter
}

var newPagingUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Unix swap or windows pagefile usage."),
	metric.WithUnit("By"),
}

func NewPagingUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PagingUsage, error) {
	_ = "STUB: not implemented"
	return *new(PagingUsage), nil
}

func (m PagingUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PagingUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (PagingUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (PagingUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m PagingUsage) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PagingUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PagingUsage) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PagingUsage) AttrPagingState(val PagingStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPagingUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Unix swap or windows pagefile usage."),
	metric.WithUnit("By"),
}

func NewPagingUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PagingUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(PagingUsageObservable), nil
}

func (m PagingUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PagingUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PagingUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PagingUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PagingUsageObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PagingUsageObservable) AttrPagingState(val PagingStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingUtilization struct {
	metric.Int64Gauge
}

var newPagingUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Swap (unix) or pagefile (windows) utilization."),
	metric.WithUnit("1"),
}

func NewPagingUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (PagingUtilization, error) {
	_ = "STUB: not implemented"
	return *new(PagingUtilization), nil
}

func (m PagingUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (PagingUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (PagingUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (PagingUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m PagingUtilization) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PagingUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PagingUtilization) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PagingUtilization) AttrPagingState(val PagingStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newPagingUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Swap (unix) or pagefile (windows) utilization."),
	metric.WithUnit("1"),
}

func NewPagingUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (PagingUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(PagingUtilizationObservable), nil
}

func (m PagingUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (PagingUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PagingUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PagingUtilizationObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PagingUtilizationObservable) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PagingUtilizationObservable) AttrPagingState(val PagingStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ProcessCount struct {
	metric.Int64UpDownCounter
}

var newProcessCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Total number of processes in each state."),
	metric.WithUnit("{process}"),
}

func NewProcessCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ProcessCount, error) {
	_ = "STUB: not implemented"
	return *new(ProcessCount), nil
}

func (m ProcessCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ProcessCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ProcessCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ProcessCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ProcessCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ProcessCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ProcessCount) AttrProcessState(val ProcessStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ProcessCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newProcessCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Total number of processes in each state."),
	metric.WithUnit("{process}"),
}

func NewProcessCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ProcessCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(ProcessCountObservable), nil
}

func (m ProcessCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ProcessCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ProcessCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ProcessCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ProcessCountObservable) AttrProcessState(val ProcessStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ProcessCreated struct {
	metric.Int64Counter
}

var newProcessCreatedOpts = []metric.Int64CounterOption{
	metric.WithDescription("Total number of processes created over uptime of the host."),
	metric.WithUnit("{process}"),
}

func NewProcessCreated(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ProcessCreated, error) {
	_ = "STUB: not implemented"
	return *new(ProcessCreated), nil
}

func (m ProcessCreated) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ProcessCreated) Name() string { _ = "STUB: not implemented"; return "" }

func (ProcessCreated) Unit() string { _ = "STUB: not implemented"; return "" }

func (ProcessCreated) Description() string { _ = "STUB: not implemented"; return "" }

func (m ProcessCreated) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ProcessCreated) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ProcessCreatedObservable struct {
	metric.Int64ObservableCounter
}

var newProcessCreatedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Total number of processes created over uptime of the host."),
	metric.WithUnit("{process}"),
}

func NewProcessCreatedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ProcessCreatedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ProcessCreatedObservable), nil
}

func (m ProcessCreatedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (ProcessCreatedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ProcessCreatedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ProcessCreatedObservable) Description() string { _ = "STUB: not implemented"; return "" }

type Uptime struct {
	metric.Float64Gauge
}

var newUptimeOpts = []metric.Float64GaugeOption{
	metric.WithDescription("The time the system has been running."),
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
	metric.WithDescription("The time the system has been running."),
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
