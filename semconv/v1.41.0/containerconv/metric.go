package containerconv

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

type SystemPagingFaultTypeAttr string

var (
	SystemPagingFaultTypeMajor SystemPagingFaultTypeAttr = "major"

	SystemPagingFaultTypeMinor SystemPagingFaultTypeAttr = "minor"
)

type CPUTime struct {
	metric.Float64Counter
}

var newCPUTimeOpts = []metric.Float64CounterOption{
	metric.WithDescription("Total CPU time consumed."),
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
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CPUTime) AddSet(ctx context.Context, incr float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (CPUTime) AttrCPUMode(val CPUModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUTimeObservable struct {
	metric.Float64ObservableCounter
}

var newCPUTimeObservableOpts = []metric.Float64ObservableCounterOption{
	metric.WithDescription("Total CPU time consumed."),
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

func (CPUTimeObservable) AttrCPUMode(val CPUModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUUsage struct {
	metric.Int64Gauge
}

var newCPUUsageOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Container's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs."),
	metric.WithUnit("{cpu}"),
}

func NewCPUUsage(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (CPUUsage, error) {
	_ = "STUB: not implemented"
	return *new(CPUUsage), nil
}

func (m CPUUsage) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (CPUUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m CPUUsage) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CPUUsage) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (CPUUsage) AttrCPUMode(val CPUModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUUsageObservable struct {
	metric.Int64ObservableGauge
}

var newCPUUsageObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Container's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs."),
	metric.WithUnit("{cpu}"),
}

func NewCPUUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (CPUUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(CPUUsageObservable), nil
}

func (m CPUUsageObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (CPUUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (CPUUsageObservable) AttrCPUMode(val CPUModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskIO struct {
	metric.Int64Counter
}

var newDiskIOOpts = []metric.Int64CounterOption{
	metric.WithDescription("Disk bytes for the container."),
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

func (DiskIO) AttrSystemDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskIOObservable struct {
	metric.Int64ObservableCounter
}

var newDiskIOObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Disk bytes for the container."),
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

func (DiskIOObservable) AttrSystemDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FilesystemAvailable struct {
	metric.Int64UpDownCounter
}

var newFilesystemAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Container filesystem available bytes."),
	metric.WithUnit("By"),
}

func NewFilesystemAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (FilesystemAvailable, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemAvailable), nil
}

func (m FilesystemAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (FilesystemAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m FilesystemAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m FilesystemAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type FilesystemAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newFilesystemAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Container filesystem available bytes."),
	metric.WithUnit("By"),
}

func NewFilesystemAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (FilesystemAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemAvailableObservable), nil
}

func (m FilesystemAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (FilesystemAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

type FilesystemCapacity struct {
	metric.Int64UpDownCounter
}

var newFilesystemCapacityOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Container filesystem capacity."),
	metric.WithUnit("By"),
}

func NewFilesystemCapacity(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (FilesystemCapacity, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemCapacity), nil
}

func (m FilesystemCapacity) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (FilesystemCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemCapacity) Description() string { _ = "STUB: not implemented"; return "" }

func (m FilesystemCapacity) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m FilesystemCapacity) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type FilesystemCapacityObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newFilesystemCapacityObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Container filesystem capacity."),
	metric.WithUnit("By"),
}

func NewFilesystemCapacityObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (FilesystemCapacityObservable, error) {
	_ = "STUB: not implemented"
	return *new(FilesystemCapacityObservable), nil
}

func (m FilesystemCapacityObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (FilesystemCapacityObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (FilesystemCapacityObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (FilesystemCapacityObservable) Description() string { _ = "STUB: not implemented"; return "" }

type FilesystemUsage struct {
	metric.Int64UpDownCounter
}

var newFilesystemUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Container filesystem usage."),
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

func (m FilesystemUsage) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m FilesystemUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type FilesystemUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newFilesystemUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Container filesystem usage."),
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

type MemoryAvailable struct {
	metric.Int64UpDownCounter
}

var newMemoryAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Container memory available."),
	metric.WithUnit("By"),
}

func NewMemoryAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryAvailable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryAvailable), nil
}

func (m MemoryAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Container memory available."),
	metric.WithUnit("By"),
}

func NewMemoryAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryAvailableObservable), nil
}

func (m MemoryAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryPagingFaults struct {
	metric.Int64Counter
}

var newMemoryPagingFaultsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Container memory paging faults."),
	metric.WithUnit("{fault}"),
}

func NewMemoryPagingFaults(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (MemoryPagingFaults, error) {
	_ = "STUB: not implemented"
	return *new(MemoryPagingFaults), nil
}

func (m MemoryPagingFaults) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (MemoryPagingFaults) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryPagingFaults) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryPagingFaults) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryPagingFaults) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryPagingFaults) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (MemoryPagingFaults) AttrSystemPagingFaultType(val SystemPagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryPagingFaultsObservable struct {
	metric.Int64ObservableCounter
}

var newMemoryPagingFaultsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Container memory paging faults."),
	metric.WithUnit("{fault}"),
}

func NewMemoryPagingFaultsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (MemoryPagingFaultsObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryPagingFaultsObservable), nil
}

func (m MemoryPagingFaultsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (MemoryPagingFaultsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryPagingFaultsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryPagingFaultsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (MemoryPagingFaultsObservable) AttrSystemPagingFaultType(val SystemPagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryRss struct {
	metric.Int64UpDownCounter
}

var newMemoryRssOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Container memory RSS."),
	metric.WithUnit("By"),
}

func NewMemoryRss(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryRss, error) {
	_ = "STUB: not implemented"
	return *new(MemoryRss), nil
}

func (m MemoryRss) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryRss) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryRss) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryRss) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryRss) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryRss) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryRssObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryRssObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Container memory RSS."),
	metric.WithUnit("By"),
}

func NewMemoryRssObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryRssObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryRssObservable), nil
}

func (m MemoryRssObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryRssObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryRssObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryRssObservable) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryUsage struct {
	metric.Int64Counter
}

var newMemoryUsageOpts = []metric.Int64CounterOption{
	metric.WithDescription("Memory usage of the container."),
	metric.WithUnit("By"),
}

func NewMemoryUsage(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (MemoryUsage, error) {
	_ = "STUB: not implemented"
	return *new(MemoryUsage), nil
}

func (m MemoryUsage) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
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
	metric.Int64ObservableCounter
}

var newMemoryUsageObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Memory usage of the container."),
	metric.WithUnit("By"),
}

func NewMemoryUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (MemoryUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryUsageObservable), nil
}

func (m MemoryUsageObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (MemoryUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryWorkingSet struct {
	metric.Int64UpDownCounter
}

var newMemoryWorkingSetOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Container memory working set."),
	metric.WithUnit("By"),
}

func NewMemoryWorkingSet(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemoryWorkingSet, error) {
	_ = "STUB: not implemented"
	return *new(MemoryWorkingSet), nil
}

func (m MemoryWorkingSet) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemoryWorkingSet) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryWorkingSet) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryWorkingSet) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemoryWorkingSet) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m MemoryWorkingSet) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type MemoryWorkingSetObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemoryWorkingSetObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Container memory working set."),
	metric.WithUnit("By"),
}

func NewMemoryWorkingSetObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryWorkingSetObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemoryWorkingSetObservable), nil
}

func (m MemoryWorkingSetObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryWorkingSetObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryWorkingSetObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryWorkingSetObservable) Description() string { _ = "STUB: not implemented"; return "" }

type NetworkIO struct {
	metric.Int64Counter
}

var newNetworkIOOpts = []metric.Int64CounterOption{
	metric.WithDescription("Network bytes for the container."),
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
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkIO) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NetworkIO) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIO) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkIOObservable struct {
	metric.Int64ObservableCounter
}

var newNetworkIOObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Network bytes for the container."),
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

func (NetworkIOObservable) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIOObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Uptime struct {
	metric.Float64Gauge
}

var newUptimeOpts = []metric.Float64GaugeOption{
	metric.WithDescription("The time the container has been running."),
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
	metric.WithDescription("The time the container has been running."),
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
