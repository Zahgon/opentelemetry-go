package hwconv

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

type ErrorTypeAttr string

var ErrorTypeOther ErrorTypeAttr = "_OTHER"

type BatteryStateAttr string

var (
	BatteryStateCharging BatteryStateAttr = "charging"

	BatteryStateDischarging BatteryStateAttr = "discharging"
)

type GpuTaskAttr string

var (
	GpuTaskDecoder GpuTaskAttr = "decoder"

	GpuTaskEncoder GpuTaskAttr = "encoder"

	GpuTaskGeneral GpuTaskAttr = "general"
)

type LimitTypeAttr string

var (
	LimitTypeCritical LimitTypeAttr = "critical"

	LimitTypeDegraded LimitTypeAttr = "degraded"

	LimitTypeHighCritical LimitTypeAttr = "high.critical"

	LimitTypeHighDegraded LimitTypeAttr = "high.degraded"

	LimitTypeLowCritical LimitTypeAttr = "low.critical"

	LimitTypeLowDegraded LimitTypeAttr = "low.degraded"

	LimitTypeMax LimitTypeAttr = "max"

	LimitTypeThrottled LimitTypeAttr = "throttled"

	LimitTypeTurbo LimitTypeAttr = "turbo"
)

type LogicalDiskStateAttr string

var (
	LogicalDiskStateUsed LogicalDiskStateAttr = "used"

	LogicalDiskStateFree LogicalDiskStateAttr = "free"
)

type PhysicalDiskStateAttr string

var PhysicalDiskStateRemaining PhysicalDiskStateAttr = "remaining"

type StateAttr string

var (
	StateDegraded StateAttr = "degraded"

	StateFailed StateAttr = "failed"

	StateNeedsCleaning StateAttr = "needs_cleaning"

	StateOk StateAttr = "ok"

	StatePredictedFailure StateAttr = "predicted_failure"
)

type TapeDriveOperationTypeAttr string

var (
	TapeDriveOperationTypeMount TapeDriveOperationTypeAttr = "mount"

	TapeDriveOperationTypeUnmount TapeDriveOperationTypeAttr = "unmount"

	TapeDriveOperationTypeClean TapeDriveOperationTypeAttr = "clean"
)

type TypeAttr string

var (
	TypeBattery TypeAttr = "battery"

	TypeCPU TypeAttr = "cpu"

	TypeDiskController TypeAttr = "disk_controller"

	TypeEnclosure TypeAttr = "enclosure"

	TypeFan TypeAttr = "fan"

	TypeGpu TypeAttr = "gpu"

	TypeLogicalDisk TypeAttr = "logical_disk"

	TypeMemory TypeAttr = "memory"

	TypeNetwork TypeAttr = "network"

	TypePhysicalDisk TypeAttr = "physical_disk"

	TypePowerSupply TypeAttr = "power_supply"

	TypeTapeDrive TypeAttr = "tape_drive"

	TypeTemperature TypeAttr = "temperature"

	TypeVoltage TypeAttr = "voltage"
)

type NetworkIODirectionAttr string

var (
	NetworkIODirectionTransmit NetworkIODirectionAttr = "transmit"

	NetworkIODirectionReceive NetworkIODirectionAttr = "receive"
)

type BatteryCharge struct {
	metric.Int64Gauge
}

var newBatteryChargeOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Remaining fraction of battery charge."),
	metric.WithUnit("1"),
}

func NewBatteryCharge(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (BatteryCharge, error) {
	_ = "STUB: not implemented"
	return *new(BatteryCharge), nil
}

func (m BatteryCharge) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (BatteryCharge) Name() string { _ = "STUB: not implemented"; return "" }

func (BatteryCharge) Unit() string { _ = "STUB: not implemented"; return "" }

func (BatteryCharge) Description() string { _ = "STUB: not implemented"; return "" }

func (m BatteryCharge) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m BatteryCharge) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (BatteryCharge) AttrBatteryCapacity(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryCharge) AttrBatteryChemistry(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryCharge) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryCharge) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryCharge) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryCharge) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type BatteryChargeLimit struct {
	metric.Int64Gauge
}

var newBatteryChargeLimitOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Lower limit of battery charge fraction to ensure proper operation."),
	metric.WithUnit("1"),
}

func NewBatteryChargeLimit(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (BatteryChargeLimit, error) {
	_ = "STUB: not implemented"
	return *new(BatteryChargeLimit), nil
}

func (m BatteryChargeLimit) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (BatteryChargeLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (BatteryChargeLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (BatteryChargeLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m BatteryChargeLimit) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m BatteryChargeLimit) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (BatteryChargeLimit) AttrBatteryCapacity(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimit) AttrBatteryChemistry(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimit) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimit) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimit) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimit) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimit) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type BatteryTimeLeft struct {
	metric.Float64Gauge
}

var newBatteryTimeLeftOpts = []metric.Float64GaugeOption{
	metric.WithDescription("Time left before battery is completely charged or discharged."),
	metric.WithUnit("s"),
}

func NewBatteryTimeLeft(
	m metric.Meter,
	opt ...metric.Float64GaugeOption,
) (BatteryTimeLeft, error) {
	_ = "STUB: not implemented"
	return *new(BatteryTimeLeft), nil
}

func (m BatteryTimeLeft) Inst() metric.Float64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge)
}

func (BatteryTimeLeft) Name() string { _ = "STUB: not implemented"; return "" }

func (BatteryTimeLeft) Unit() string { _ = "STUB: not implemented"; return "" }

func (BatteryTimeLeft) Description() string { _ = "STUB: not implemented"; return "" }

func (m BatteryTimeLeft) Record(
	ctx context.Context,
	val float64,
	id string,
	state StateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m BatteryTimeLeft) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (BatteryTimeLeft) AttrBatteryState(val BatteryStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeft) AttrBatteryCapacity(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeft) AttrBatteryChemistry(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeft) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeft) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeft) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeft) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUSpeed struct {
	metric.Int64Gauge
}

var newCPUSpeedOpts = []metric.Int64GaugeOption{
	metric.WithDescription("CPU current frequency."),
	metric.WithUnit("Hz"),
}

func NewCPUSpeed(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (CPUSpeed, error) {
	_ = "STUB: not implemented"
	return *new(CPUSpeed), nil
}

func (m CPUSpeed) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (CPUSpeed) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeed) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeed) Description() string { _ = "STUB: not implemented"; return "" }

func (m CPUSpeed) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CPUSpeed) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (CPUSpeed) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeed) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeed) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeed) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CPUSpeedLimit struct {
	metric.Int64Gauge
}

var newCPUSpeedLimitOpts = []metric.Int64GaugeOption{
	metric.WithDescription("CPU maximum frequency."),
	metric.WithUnit("Hz"),
}

func NewCPUSpeedLimit(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (CPUSpeedLimit, error) {
	_ = "STUB: not implemented"
	return *new(CPUSpeedLimit), nil
}

func (m CPUSpeedLimit) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (CPUSpeedLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeedLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeedLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m CPUSpeedLimit) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CPUSpeedLimit) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (CPUSpeedLimit) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedLimit) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedLimit) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedLimit) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedLimit) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Energy struct {
	metric.Int64Counter
}

var newEnergyOpts = []metric.Int64CounterOption{
	metric.WithDescription("Energy consumed by the component."),
	metric.WithUnit("J"),
}

func NewEnergy(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (Energy, error) {
	_ = "STUB: not implemented"
	return *new(Energy), nil
}

func (m Energy) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (Energy) Name() string { _ = "STUB: not implemented"; return "" }

func (Energy) Unit() string { _ = "STUB: not implemented"; return "" }

func (Energy) Description() string { _ = "STUB: not implemented"; return "" }

func (m Energy) Add(
	ctx context.Context,
	incr int64,
	id string,
	hwType TypeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m Energy) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Energy) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Energy) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Errors struct {
	metric.Int64Counter
}

var newErrorsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of errors encountered by the component."),
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
	id string,
	hwType TypeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m Errors) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Errors) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Errors) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Errors) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Errors) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FanSpeed struct {
	metric.Int64Gauge
}

var newFanSpeedOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Fan speed in revolutions per minute."),
	metric.WithUnit("rpm"),
}

func NewFanSpeed(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (FanSpeed, error) {
	_ = "STUB: not implemented"
	return *new(FanSpeed), nil
}

func (m FanSpeed) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (FanSpeed) Name() string { _ = "STUB: not implemented"; return "" }

func (FanSpeed) Unit() string { _ = "STUB: not implemented"; return "" }

func (FanSpeed) Description() string { _ = "STUB: not implemented"; return "" }

func (m FanSpeed) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m FanSpeed) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (FanSpeed) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeed) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeed) AttrSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FanSpeedLimit struct {
	metric.Int64Gauge
}

var newFanSpeedLimitOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Speed limit in rpm."),
	metric.WithUnit("rpm"),
}

func NewFanSpeedLimit(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (FanSpeedLimit, error) {
	_ = "STUB: not implemented"
	return *new(FanSpeedLimit), nil
}

func (m FanSpeedLimit) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (FanSpeedLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m FanSpeedLimit) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m FanSpeedLimit) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (FanSpeedLimit) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedLimit) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedLimit) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedLimit) AttrSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type FanSpeedRatio struct {
	metric.Int64Gauge
}

var newFanSpeedRatioOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Fan speed expressed as a fraction of its maximum speed."),
	metric.WithUnit("1"),
}

func NewFanSpeedRatio(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (FanSpeedRatio, error) {
	_ = "STUB: not implemented"
	return *new(FanSpeedRatio), nil
}

func (m FanSpeedRatio) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (FanSpeedRatio) Name() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedRatio) Unit() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedRatio) Description() string { _ = "STUB: not implemented"; return "" }

func (m FanSpeedRatio) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m FanSpeedRatio) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (FanSpeedRatio) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedRatio) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedRatio) AttrSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type GpuIO struct {
	metric.Int64Counter
}

var newGpuIOOpts = []metric.Int64CounterOption{
	metric.WithDescription("Received and transmitted bytes by the GPU."),
	metric.WithUnit("By"),
}

func NewGpuIO(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (GpuIO, error) {
	_ = "STUB: not implemented"
	return *new(GpuIO), nil
}

func (m GpuIO) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (GpuIO) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuIO) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuIO) Description() string { _ = "STUB: not implemented"; return "" }

func (m GpuIO) Add(
	ctx context.Context,
	incr int64,
	id string,
	networkIoDirection NetworkIODirectionAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m GpuIO) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (GpuIO) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIO) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIO) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIO) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIO) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIO) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIO) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type GpuMemoryLimit struct {
	metric.Int64UpDownCounter
}

var newGpuMemoryLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Size of the GPU memory."),
	metric.WithUnit("By"),
}

func NewGpuMemoryLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (GpuMemoryLimit, error) {
	_ = "STUB: not implemented"
	return *new(GpuMemoryLimit), nil
}

func (m GpuMemoryLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (GpuMemoryLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m GpuMemoryLimit) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m GpuMemoryLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (GpuMemoryLimit) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimit) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimit) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimit) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimit) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimit) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimit) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type GpuMemoryUsage struct {
	metric.Int64UpDownCounter
}

var newGpuMemoryUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("GPU memory used."),
	metric.WithUnit("By"),
}

func NewGpuMemoryUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (GpuMemoryUsage, error) {
	_ = "STUB: not implemented"
	return *new(GpuMemoryUsage), nil
}

func (m GpuMemoryUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (GpuMemoryUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m GpuMemoryUsage) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m GpuMemoryUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (GpuMemoryUsage) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsage) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsage) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsage) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsage) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsage) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsage) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type GpuMemoryUtilization struct {
	metric.Int64Gauge
}

var newGpuMemoryUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Fraction of GPU memory used."),
	metric.WithUnit("1"),
}

func NewGpuMemoryUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (GpuMemoryUtilization, error) {
	_ = "STUB: not implemented"
	return *new(GpuMemoryUtilization), nil
}

func (m GpuMemoryUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (GpuMemoryUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m GpuMemoryUtilization) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m GpuMemoryUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (GpuMemoryUtilization) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilization) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilization) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilization) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilization) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilization) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilization) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type GpuUtilization struct {
	metric.Int64Gauge
}

var newGpuUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Fraction of time spent in a specific task."),
	metric.WithUnit("1"),
}

func NewGpuUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (GpuUtilization, error) {
	_ = "STUB: not implemented"
	return *new(GpuUtilization), nil
}

func (m GpuUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (GpuUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m GpuUtilization) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m GpuUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (GpuUtilization) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilization) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilization) AttrGpuTask(val GpuTaskAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilization) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilization) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilization) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilization) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilization) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type HostAmbientTemperature struct {
	metric.Int64Gauge
}

var newHostAmbientTemperatureOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Ambient (external) temperature of the physical host."),
	metric.WithUnit("Cel"),
}

func NewHostAmbientTemperature(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (HostAmbientTemperature, error) {
	_ = "STUB: not implemented"
	return *new(HostAmbientTemperature), nil
}

func (m HostAmbientTemperature) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (HostAmbientTemperature) Name() string { _ = "STUB: not implemented"; return "" }

func (HostAmbientTemperature) Unit() string { _ = "STUB: not implemented"; return "" }

func (HostAmbientTemperature) Description() string { _ = "STUB: not implemented"; return "" }

func (m HostAmbientTemperature) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m HostAmbientTemperature) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (HostAmbientTemperature) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostAmbientTemperature) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type HostEnergy struct {
	metric.Int64Counter
}

var newHostEnergyOpts = []metric.Int64CounterOption{
	metric.WithDescription("Total energy consumed by the entire physical host, in joules."),
	metric.WithUnit("J"),
}

func NewHostEnergy(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (HostEnergy, error) {
	_ = "STUB: not implemented"
	return *new(HostEnergy), nil
}

func (m HostEnergy) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (HostEnergy) Name() string { _ = "STUB: not implemented"; return "" }

func (HostEnergy) Unit() string { _ = "STUB: not implemented"; return "" }

func (HostEnergy) Description() string { _ = "STUB: not implemented"; return "" }

func (m HostEnergy) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m HostEnergy) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (HostEnergy) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostEnergy) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type HostHeatingMargin struct {
	metric.Int64Gauge
}

var newHostHeatingMarginOpts = []metric.Int64GaugeOption{
	metric.WithDescription("By how many degrees Celsius the temperature of the physical host can be increased, before reaching a warning threshold on one of the internal sensors."),
	metric.WithUnit("Cel"),
}

func NewHostHeatingMargin(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (HostHeatingMargin, error) {
	_ = "STUB: not implemented"
	return *new(HostHeatingMargin), nil
}

func (m HostHeatingMargin) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (HostHeatingMargin) Name() string { _ = "STUB: not implemented"; return "" }

func (HostHeatingMargin) Unit() string { _ = "STUB: not implemented"; return "" }

func (HostHeatingMargin) Description() string { _ = "STUB: not implemented"; return "" }

func (m HostHeatingMargin) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m HostHeatingMargin) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (HostHeatingMargin) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostHeatingMargin) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type HostPower struct {
	metric.Int64Gauge
}

var newHostPowerOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Instantaneous power consumed by the entire physical host in Watts (`hw.host.energy` is preferred)."),
	metric.WithUnit("W"),
}

func NewHostPower(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (HostPower, error) {
	_ = "STUB: not implemented"
	return *new(HostPower), nil
}

func (m HostPower) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (HostPower) Name() string { _ = "STUB: not implemented"; return "" }

func (HostPower) Unit() string { _ = "STUB: not implemented"; return "" }

func (HostPower) Description() string { _ = "STUB: not implemented"; return "" }

func (m HostPower) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m HostPower) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (HostPower) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostPower) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type LogicalDiskLimit struct {
	metric.Int64UpDownCounter
}

var newLogicalDiskLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Size of the logical disk."),
	metric.WithUnit("By"),
}

func NewLogicalDiskLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (LogicalDiskLimit, error) {
	_ = "STUB: not implemented"
	return *new(LogicalDiskLimit), nil
}

func (m LogicalDiskLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (LogicalDiskLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m LogicalDiskLimit) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m LogicalDiskLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (LogicalDiskLimit) AttrLogicalDiskRaidLevel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskLimit) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskLimit) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type LogicalDiskUsage struct {
	metric.Int64UpDownCounter
}

var newLogicalDiskUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Logical disk space usage."),
	metric.WithUnit("By"),
}

func NewLogicalDiskUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (LogicalDiskUsage, error) {
	_ = "STUB: not implemented"
	return *new(LogicalDiskUsage), nil
}

func (m LogicalDiskUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (LogicalDiskUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m LogicalDiskUsage) Add(
	ctx context.Context,
	incr int64,
	id string,
	logicalDiskState LogicalDiskStateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m LogicalDiskUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (LogicalDiskUsage) AttrLogicalDiskRaidLevel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUsage) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUsage) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type LogicalDiskUtilization struct {
	metric.Int64Gauge
}

var newLogicalDiskUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Logical disk space utilization as a fraction."),
	metric.WithUnit("1"),
}

func NewLogicalDiskUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (LogicalDiskUtilization, error) {
	_ = "STUB: not implemented"
	return *new(LogicalDiskUtilization), nil
}

func (m LogicalDiskUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (LogicalDiskUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m LogicalDiskUtilization) Record(
	ctx context.Context,
	val int64,
	id string,
	logicalDiskState LogicalDiskStateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m LogicalDiskUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (LogicalDiskUtilization) AttrLogicalDiskRaidLevel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUtilization) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUtilization) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type MemorySize struct {
	metric.Int64UpDownCounter
}

var newMemorySizeOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Size of the memory module."),
	metric.WithUnit("By"),
}

func NewMemorySize(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (MemorySize, error) {
	_ = "STUB: not implemented"
	return *new(MemorySize), nil
}

func (m MemorySize) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (MemorySize) Name() string { _ = "STUB: not implemented"; return "" }

func (MemorySize) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemorySize) Description() string { _ = "STUB: not implemented"; return "" }

func (m MemorySize) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m MemorySize) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (MemorySize) AttrMemoryType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySize) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySize) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySize) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySize) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySize) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkBandwidthLimit struct {
	metric.Int64UpDownCounter
}

var newNetworkBandwidthLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Link speed."),
	metric.WithUnit("By/s"),
}

func NewNetworkBandwidthLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NetworkBandwidthLimit, error) {
	_ = "STUB: not implemented"
	return *new(NetworkBandwidthLimit), nil
}

func (m NetworkBandwidthLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NetworkBandwidthLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkBandwidthLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkBandwidthLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkBandwidthLimit) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkBandwidthLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NetworkBandwidthLimit) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimit) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimit) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimit) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimit) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimit) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimit) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkBandwidthUtilization struct {
	metric.Int64Gauge
}

var newNetworkBandwidthUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Utilization of the network bandwidth as a fraction."),
	metric.WithUnit("1"),
}

func NewNetworkBandwidthUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (NetworkBandwidthUtilization, error) {
	_ = "STUB: not implemented"
	return *new(NetworkBandwidthUtilization), nil
}

func (m NetworkBandwidthUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (NetworkBandwidthUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkBandwidthUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkBandwidthUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkBandwidthUtilization) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkBandwidthUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NetworkBandwidthUtilization) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilization) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilization) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilization) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilization) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilization) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilization) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkIO struct {
	metric.Int64Counter
}

var newNetworkIOOpts = []metric.Int64CounterOption{
	metric.WithDescription("Received and transmitted network traffic in bytes."),
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
	id string,
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

func (NetworkIO) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIO) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIO) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIO) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIO) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIO) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIO) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkPackets struct {
	metric.Int64Counter
}

var newNetworkPacketsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Received and transmitted network traffic in packets (or frames)."),
	metric.WithUnit("{packet}"),
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

func (NetworkPackets) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkPackets) Add(
	ctx context.Context,
	incr int64,
	id string,
	networkIoDirection NetworkIODirectionAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkPackets) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NetworkPackets) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPackets) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPackets) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPackets) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPackets) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPackets) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPackets) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NetworkUp struct {
	metric.Int64UpDownCounter
}

var newNetworkUpOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Link status: `1` (up) or `0` (down)."),
	metric.WithUnit("1"),
}

func NewNetworkUp(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NetworkUp, error) {
	_ = "STUB: not implemented"
	return *new(NetworkUp), nil
}

func (m NetworkUp) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NetworkUp) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkUp) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkUp) Description() string { _ = "STUB: not implemented"; return "" }

func (m NetworkUp) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NetworkUp) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NetworkUp) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUp) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUp) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUp) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUp) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUp) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUp) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PhysicalDiskEnduranceUtilization struct {
	metric.Int64Gauge
}

var newPhysicalDiskEnduranceUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Endurance remaining for this SSD disk."),
	metric.WithUnit("1"),
}

func NewPhysicalDiskEnduranceUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (PhysicalDiskEnduranceUtilization, error) {
	_ = "STUB: not implemented"
	return *new(PhysicalDiskEnduranceUtilization), nil
}

func (m PhysicalDiskEnduranceUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (PhysicalDiskEnduranceUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskEnduranceUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskEnduranceUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m PhysicalDiskEnduranceUtilization) Record(
	ctx context.Context,
	val int64,
	id string,
	physicalDiskState PhysicalDiskStateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PhysicalDiskEnduranceUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PhysicalDiskEnduranceUtilization) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilization) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilization) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilization) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilization) AttrPhysicalDiskType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilization) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilization) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PhysicalDiskSize struct {
	metric.Int64UpDownCounter
}

var newPhysicalDiskSizeOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Size of the disk."),
	metric.WithUnit("By"),
}

func NewPhysicalDiskSize(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PhysicalDiskSize, error) {
	_ = "STUB: not implemented"
	return *new(PhysicalDiskSize), nil
}

func (m PhysicalDiskSize) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PhysicalDiskSize) Name() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m PhysicalDiskSize) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PhysicalDiskSize) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PhysicalDiskSize) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSize) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSize) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSize) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSize) AttrPhysicalDiskType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSize) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSize) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PhysicalDiskSmart struct {
	metric.Int64Gauge
}

var newPhysicalDiskSmartOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Value of the corresponding [S.M.A.R.T.](https://wikipedia.org/wiki/S.M.A.R.T.) (Self-Monitoring, Analysis, and Reporting Technology) attribute."),
	metric.WithUnit("1"),
}

func NewPhysicalDiskSmart(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (PhysicalDiskSmart, error) {
	_ = "STUB: not implemented"
	return *new(PhysicalDiskSmart), nil
}

func (m PhysicalDiskSmart) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (PhysicalDiskSmart) Name() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSmart) Unit() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSmart) Description() string { _ = "STUB: not implemented"; return "" }

func (m PhysicalDiskSmart) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PhysicalDiskSmart) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PhysicalDiskSmart) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmart) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmart) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmart) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmart) AttrPhysicalDiskSmartAttribute(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmart) AttrPhysicalDiskType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmart) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmart) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Power struct {
	metric.Int64Gauge
}

var newPowerOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Instantaneous power consumed by the component."),
	metric.WithUnit("W"),
}

func NewPower(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (Power, error) {
	_ = "STUB: not implemented"
	return *new(Power), nil
}

func (m Power) Inst() metric.Int64Gauge { _ = "STUB: not implemented"; return *new(metric.Int64Gauge) }

func (Power) Name() string { _ = "STUB: not implemented"; return "" }

func (Power) Unit() string { _ = "STUB: not implemented"; return "" }

func (Power) Description() string { _ = "STUB: not implemented"; return "" }

func (m Power) Record(
	ctx context.Context,
	val int64,
	id string,
	hwType TypeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m Power) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Power) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Power) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PowerSupplyLimit struct {
	metric.Int64UpDownCounter
}

var newPowerSupplyLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Maximum power output of the power supply."),
	metric.WithUnit("W"),
}

func NewPowerSupplyLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PowerSupplyLimit, error) {
	_ = "STUB: not implemented"
	return *new(PowerSupplyLimit), nil
}

func (m PowerSupplyLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PowerSupplyLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m PowerSupplyLimit) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PowerSupplyLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PowerSupplyLimit) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimit) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimit) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimit) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimit) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimit) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PowerSupplyUsage struct {
	metric.Int64UpDownCounter
}

var newPowerSupplyUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Current power output of the power supply."),
	metric.WithUnit("W"),
}

func NewPowerSupplyUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PowerSupplyUsage, error) {
	_ = "STUB: not implemented"
	return *new(PowerSupplyUsage), nil
}

func (m PowerSupplyUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PowerSupplyUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m PowerSupplyUsage) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PowerSupplyUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PowerSupplyUsage) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUsage) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUsage) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUsage) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUsage) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PowerSupplyUtilization struct {
	metric.Int64Gauge
}

var newPowerSupplyUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Utilization of the power supply as a fraction of its maximum output."),
	metric.WithUnit("1"),
}

func NewPowerSupplyUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (PowerSupplyUtilization, error) {
	_ = "STUB: not implemented"
	return *new(PowerSupplyUtilization), nil
}

func (m PowerSupplyUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (PowerSupplyUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m PowerSupplyUtilization) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PowerSupplyUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PowerSupplyUtilization) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUtilization) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUtilization) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUtilization) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUtilization) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Status struct {
	metric.Int64UpDownCounter
}

var newStatusOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Operational status: `1` (true) or `0` (false) for each of the possible states."),
	metric.WithUnit("1"),
}

func NewStatus(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func (m Status) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (Status) Name() string { _ = "STUB: not implemented"; return "" }

func (Status) Unit() string { _ = "STUB: not implemented"; return "" }

func (Status) Description() string { _ = "STUB: not implemented"; return "" }

func (m Status) Add(
	ctx context.Context,
	incr int64,
	id string,
	state StateAttr,
	hwType TypeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m Status) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Status) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Status) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type TapeDriveOperations struct {
	metric.Int64Counter
}

var newTapeDriveOperationsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Operations performed by the tape drive."),
	metric.WithUnit("{operation}"),
}

func NewTapeDriveOperations(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (TapeDriveOperations, error) {
	_ = "STUB: not implemented"
	return *new(TapeDriveOperations), nil
}

func (m TapeDriveOperations) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (TapeDriveOperations) Name() string { _ = "STUB: not implemented"; return "" }

func (TapeDriveOperations) Unit() string { _ = "STUB: not implemented"; return "" }

func (TapeDriveOperations) Description() string { _ = "STUB: not implemented"; return "" }

func (m TapeDriveOperations) Add(
	ctx context.Context,
	incr int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m TapeDriveOperations) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (TapeDriveOperations) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperations) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperations) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperations) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperations) AttrTapeDriveOperationType(val TapeDriveOperationTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperations) AttrVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Temperature struct {
	metric.Int64Gauge
}

var newTemperatureOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Temperature in degrees Celsius."),
	metric.WithUnit("Cel"),
}

func NewTemperature(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (Temperature, error) {
	_ = "STUB: not implemented"
	return *new(Temperature), nil
}

func (m Temperature) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (Temperature) Name() string { _ = "STUB: not implemented"; return "" }

func (Temperature) Unit() string { _ = "STUB: not implemented"; return "" }

func (Temperature) Description() string { _ = "STUB: not implemented"; return "" }

func (m Temperature) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m Temperature) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Temperature) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Temperature) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Temperature) AttrSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type TemperatureLimit struct {
	metric.Int64Gauge
}

var newTemperatureLimitOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Temperature limit in degrees Celsius."),
	metric.WithUnit("Cel"),
}

func NewTemperatureLimit(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (TemperatureLimit, error) {
	_ = "STUB: not implemented"
	return *new(TemperatureLimit), nil
}

func (m TemperatureLimit) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (TemperatureLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (TemperatureLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (TemperatureLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m TemperatureLimit) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m TemperatureLimit) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (TemperatureLimit) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureLimit) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureLimit) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureLimit) AttrSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Voltage struct {
	metric.Int64Gauge
}

var newVoltageOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Voltage measured by the sensor."),
	metric.WithUnit("V"),
}

func NewVoltage(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (Voltage, error) {
	_ = "STUB: not implemented"
	return *new(Voltage), nil
}

func (m Voltage) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (Voltage) Name() string { _ = "STUB: not implemented"; return "" }

func (Voltage) Unit() string { _ = "STUB: not implemented"; return "" }

func (Voltage) Description() string { _ = "STUB: not implemented"; return "" }

func (m Voltage) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m Voltage) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (Voltage) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Voltage) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Voltage) AttrSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type VoltageLimit struct {
	metric.Int64Gauge
}

var newVoltageLimitOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Voltage limit in Volts."),
	metric.WithUnit("V"),
}

func NewVoltageLimit(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (VoltageLimit, error) {
	_ = "STUB: not implemented"
	return *new(VoltageLimit), nil
}

func (m VoltageLimit) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (VoltageLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (VoltageLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (VoltageLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m VoltageLimit) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m VoltageLimit) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (VoltageLimit) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageLimit) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageLimit) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageLimit) AttrSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type VoltageNominal struct {
	metric.Int64Gauge
}

var newVoltageNominalOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Nominal (expected) voltage."),
	metric.WithUnit("V"),
}

func NewVoltageNominal(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (VoltageNominal, error) {
	_ = "STUB: not implemented"
	return *new(VoltageNominal), nil
}

func (m VoltageNominal) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (VoltageNominal) Name() string { _ = "STUB: not implemented"; return "" }

func (VoltageNominal) Unit() string { _ = "STUB: not implemented"; return "" }

func (VoltageNominal) Description() string { _ = "STUB: not implemented"; return "" }

func (m VoltageNominal) Record(
	ctx context.Context,
	val int64,
	id string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m VoltageNominal) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (VoltageNominal) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageNominal) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageNominal) AttrSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
