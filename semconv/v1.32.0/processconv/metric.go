package processconv

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

type PagingFaultTypeAttr string

var (
	PagingFaultTypeMajor PagingFaultTypeAttr = "major"

	PagingFaultTypeMinor PagingFaultTypeAttr = "minor"
)

type ContextSwitches struct {
	metric.Int64Counter
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
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ContextSwitches) AttrContextSwitchType(val ContextSwitchTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUTime struct {
	metric.Float64ObservableCounter
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

func (CPUUtilization) AttrCPUMode(val CPUModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

func (DiskIO) Description() string { _ = "STUB: not implemented"; return "" }

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

type MemoryUsage struct {
	metric.Int64UpDownCounter
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

type MemoryVirtual struct {
	metric.Int64UpDownCounter
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

type NetworkIO struct {
	metric.Int64Counter
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

func (NetworkIO) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type OpenFileDescriptorCount struct {
	metric.Int64UpDownCounter
}

func NewOpenFileDescriptorCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (OpenFileDescriptorCount, error) {
	_ = "STUB: not implemented"
	return *new(OpenFileDescriptorCount), nil
}

func (m OpenFileDescriptorCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (OpenFileDescriptorCount) Name() string { _ = "STUB: not implemented"; return "" }

func (OpenFileDescriptorCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (OpenFileDescriptorCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m OpenFileDescriptorCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
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

func (PagingFaults) Description() string { _ = "STUB: not implemented"; return "" }

func (m PagingFaults) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (PagingFaults) AttrPagingFaultType(val PagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ThreadCount struct {
	metric.Int64UpDownCounter
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
