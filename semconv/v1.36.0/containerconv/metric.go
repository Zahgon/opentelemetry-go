package containerconv

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

type CPUTime struct {
	metric.Float64Counter
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

type CPUUsage struct {
	metric.Int64Gauge
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

type MemoryUsage struct {
	metric.Int64Counter
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

func (m Uptime) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}
