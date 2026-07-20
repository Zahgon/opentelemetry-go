package systemconv

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

type LinuxMemorySlabStateAttr string

var (
	LinuxMemorySlabStateReclaimable LinuxMemorySlabStateAttr = "reclaimable"

	LinuxMemorySlabStateUnreclaimable LinuxMemorySlabStateAttr = "unreclaimable"
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

type LinuxMemoryAvailable struct {
	metric.Int64UpDownCounter
}

var newLinuxMemoryAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("An estimate of how much memory is available for starting new applications, without causing swapping."),
	metric.WithUnit("By"),
}

func NewLinuxMemoryAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (LinuxMemoryAvailable, error) {
	_ = "STUB: not implemented"
	return *new(LinuxMemoryAvailable), nil
}

func (m LinuxMemoryAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (LinuxMemoryAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (LinuxMemoryAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (LinuxMemoryAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m LinuxMemoryAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m LinuxMemoryAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type LinuxMemorySlabUsage struct {
	metric.Int64UpDownCounter
}

var newLinuxMemorySlabUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Reports the memory used by the Linux kernel for managing caches of frequently used objects."),
	metric.WithUnit("By"),
}

func NewLinuxMemorySlabUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (LinuxMemorySlabUsage, error) {
	_ = "STUB: not implemented"
	return *new(LinuxMemorySlabUsage), nil
}

func (m LinuxMemorySlabUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (LinuxMemorySlabUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (LinuxMemorySlabUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (LinuxMemorySlabUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m LinuxMemorySlabUsage) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m LinuxMemorySlabUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (LinuxMemorySlabUsage) AttrLinuxMemorySlabState(val LinuxMemorySlabStateAttr) attribute.KeyValue {
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

type MemoryShared struct {
	metric.Int64UpDownCounter
}

var newMemorySharedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Shared memory used (mostly by tmpfs)."),
	metric.WithUnit("By"),
}

func NewMemoryShared(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryShared, error) {
	_ = "STUB: not implemented"
	return *new(MemoryShared), nil
}

func (m MemoryShared) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryShared) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryShared) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryShared) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryShared) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryShared) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
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
