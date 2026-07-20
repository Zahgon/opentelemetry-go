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

type PagingStateAttr string

var (
	PagingStateUsed PagingStateAttr = "used"

	PagingStateFree PagingStateAttr = "free"
)

type PagingTypeAttr string

var (
	PagingTypeMajor PagingTypeAttr = "major"

	PagingTypeMinor PagingTypeAttr = "minor"
)

type ProcessStatusAttr string

var (
	ProcessStatusRunning ProcessStatusAttr = "running"

	ProcessStatusSleeping ProcessStatusAttr = "sleeping"

	ProcessStatusStopped ProcessStatusAttr = "stopped"

	ProcessStatusDefunct ProcessStatusAttr = "defunct"
)

type CPULogicalCount struct {
	metric.Int64UpDownCounter
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

type CPUPhysicalCount struct {
	metric.Int64UpDownCounter
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

type DiskIO struct {
	metric.Int64Counter
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

func (m DiskIO) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
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

func (DiskIOTime) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskLimit struct {
	metric.Int64UpDownCounter
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

func (DiskLimit) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type DiskMerged struct {
	metric.Int64Counter
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

func (m DiskMerged) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
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

func (m DiskOperations) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
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

func (m FilesystemUtilization) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
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

type LinuxMemorySlabUsage struct {
	metric.Int64UpDownCounter
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

func (LinuxMemorySlabUsage) AttrLinuxMemorySlabState(val LinuxMemorySlabStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemoryLimit struct {
	metric.Int64UpDownCounter
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

type MemoryShared struct {
	metric.Int64UpDownCounter
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

type MemoryUsage struct {
	metric.Int64ObservableUpDownCounter
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

func (MemoryUtilization) AttrMemoryState(val MemoryStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkConnections struct {
	metric.Int64UpDownCounter
}

func NewNetworkConnections(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NetworkConnections, error) {
	_ = "STUB: not implemented"
	return *new(NetworkConnections), nil
}

func (m NetworkConnections) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NetworkConnections) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkConnections) Unit() string { _ = "STUB: not implemented"; return "" }

func (m NetworkConnections) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (NetworkConnections) AttrNetworkConnectionState(val NetworkConnectionStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkConnections) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkConnections) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkDropped struct {
	metric.Int64Counter
}

func NewNetworkDropped(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (NetworkDropped, error) {
	_ = "STUB: not implemented"
	return *new(NetworkDropped), nil
}

func (m NetworkDropped) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (NetworkDropped) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkDropped) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkDropped) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkDropped) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (NetworkDropped) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkDropped) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkErrors struct {
	metric.Int64Counter
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

func (NetworkIO) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIO) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkPackets struct {
	metric.Int64Counter
}

func NewNetworkPackets(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (NetworkPackets, error) {
	_ = "STUB: not implemented"
	return *new(NetworkPackets), nil
}

func (m NetworkPackets) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (NetworkPackets) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkPackets) Unit() string { _ = "STUB: not implemented"; return "" }

func (m NetworkPackets) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (NetworkPackets) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPackets) AttrDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingFaults struct {
	metric.Int64Counter
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

func (m PagingFaults) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (PagingFaults) AttrPagingType(val PagingTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingOperations struct {
	metric.Int64Counter
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

func (m PagingOperations) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (PagingOperations) AttrPagingDirection(val PagingDirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PagingOperations) AttrPagingType(val PagingTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PagingUsage struct {
	metric.Int64UpDownCounter
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

func (m PagingUtilization) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
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

func (ProcessCount) AttrProcessStatus(val ProcessStatusAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ProcessCreated struct {
	metric.Int64Counter
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

type Uptime struct {
	metric.Float64Gauge
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
