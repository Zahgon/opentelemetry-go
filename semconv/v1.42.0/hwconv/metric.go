package hwconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type ErrorTypeAttr string

var (
	ErrorTypeOther ErrorTypeAttr = "_OTHER"
)

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

var (
	PhysicalDiskStateRemaining PhysicalDiskStateAttr = "remaining"
)

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

type BatteryChargeObservable struct {
	metric.Int64ObservableGauge
}

var newBatteryChargeObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Remaining fraction of battery charge."),
	metric.WithUnit("1"),
}

func NewBatteryChargeObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (BatteryChargeObservable, error) {
	_ = "STUB: not implemented"
	return *new(BatteryChargeObservable), nil
}

func (m BatteryChargeObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (BatteryChargeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (BatteryChargeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (BatteryChargeObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (BatteryChargeObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeObservable) AttrBatteryCapacity(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeObservable) AttrBatteryChemistry(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeObservable) AttrVendor(val string) attribute.KeyValue {
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

type BatteryChargeLimitObservable struct {
	metric.Int64ObservableGauge
}

var newBatteryChargeLimitObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Lower limit of battery charge fraction to ensure proper operation."),
	metric.WithUnit("1"),
}

func NewBatteryChargeLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (BatteryChargeLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(BatteryChargeLimitObservable), nil
}

func (m BatteryChargeLimitObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (BatteryChargeLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (BatteryChargeLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (BatteryChargeLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (BatteryChargeLimitObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimitObservable) AttrBatteryCapacity(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimitObservable) AttrBatteryChemistry(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimitObservable) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimitObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimitObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimitObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryChargeLimitObservable) AttrVendor(val string) attribute.KeyValue {
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

type BatteryTimeLeftObservable struct {
	metric.Float64ObservableGauge
}

var newBatteryTimeLeftObservableOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("Time left before battery is completely charged or discharged."),
	metric.WithUnit("s"),
}

func NewBatteryTimeLeftObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (BatteryTimeLeftObservable, error) {
	_ = "STUB: not implemented"
	return *new(BatteryTimeLeftObservable), nil
}

func (m BatteryTimeLeftObservable) Inst() metric.Float64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge)
}

func (BatteryTimeLeftObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (BatteryTimeLeftObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (BatteryTimeLeftObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (BatteryTimeLeftObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeftObservable) AttrState(val StateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeftObservable) AttrBatteryState(val BatteryStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeftObservable) AttrBatteryCapacity(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeftObservable) AttrBatteryChemistry(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeftObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeftObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeftObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (BatteryTimeLeftObservable) AttrVendor(val string) attribute.KeyValue {
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

type CPUSpeedObservable struct {
	metric.Int64ObservableGauge
}

var newCPUSpeedObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("CPU current frequency."),
	metric.WithUnit("Hz"),
}

func NewCPUSpeedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (CPUSpeedObservable, error) {
	_ = "STUB: not implemented"
	return *new(CPUSpeedObservable), nil
}

func (m CPUSpeedObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (CPUSpeedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeedObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeedObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedObservable) AttrVendor(val string) attribute.KeyValue {
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

type CPUSpeedLimitObservable struct {
	metric.Int64ObservableGauge
}

var newCPUSpeedLimitObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("CPU maximum frequency."),
	metric.WithUnit("Hz"),
}

func NewCPUSpeedLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (CPUSpeedLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(CPUSpeedLimitObservable), nil
}

func (m CPUSpeedLimitObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (CPUSpeedLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeedLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeedLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (CPUSpeedLimitObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedLimitObservable) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedLimitObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedLimitObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedLimitObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CPUSpeedLimitObservable) AttrVendor(val string) attribute.KeyValue {
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

type EnergyObservable struct {
	metric.Int64ObservableCounter
}

var newEnergyObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Energy consumed by the component."),
	metric.WithUnit("J"),
}

func NewEnergyObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (EnergyObservable, error) {
	_ = "STUB: not implemented"
	return *new(EnergyObservable), nil
}

func (m EnergyObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (EnergyObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (EnergyObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (EnergyObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (EnergyObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (EnergyObservable) AttrType(val TypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (EnergyObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (EnergyObservable) AttrParent(val string) attribute.KeyValue {
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

type ErrorsObservable struct {
	metric.Int64ObservableCounter
}

var newErrorsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Number of errors encountered by the component."),
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

func (ErrorsObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ErrorsObservable) AttrType(val TypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ErrorsObservable) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ErrorsObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ErrorsObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ErrorsObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
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

type FanSpeedObservable struct {
	metric.Int64ObservableGauge
}

var newFanSpeedObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Fan speed in revolutions per minute."),
	metric.WithUnit("rpm"),
}

func NewFanSpeedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (FanSpeedObservable, error) {
	_ = "STUB: not implemented"
	return *new(FanSpeedObservable), nil
}

func (m FanSpeedObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (FanSpeedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedObservable) AttrSensorLocation(val string) attribute.KeyValue {
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

type FanSpeedLimitObservable struct {
	metric.Int64ObservableGauge
}

var newFanSpeedLimitObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Speed limit in rpm."),
	metric.WithUnit("rpm"),
}

func NewFanSpeedLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (FanSpeedLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(FanSpeedLimitObservable), nil
}

func (m FanSpeedLimitObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (FanSpeedLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedLimitObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedLimitObservable) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedLimitObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedLimitObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedLimitObservable) AttrSensorLocation(val string) attribute.KeyValue {
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

type FanSpeedRatioObservable struct {
	metric.Int64ObservableGauge
}

var newFanSpeedRatioObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Fan speed expressed as a fraction of its maximum speed."),
	metric.WithUnit("1"),
}

func NewFanSpeedRatioObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (FanSpeedRatioObservable, error) {
	_ = "STUB: not implemented"
	return *new(FanSpeedRatioObservable), nil
}

func (m FanSpeedRatioObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (FanSpeedRatioObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedRatioObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedRatioObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (FanSpeedRatioObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedRatioObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedRatioObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (FanSpeedRatioObservable) AttrSensorLocation(val string) attribute.KeyValue {
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

type GpuIOObservable struct {
	metric.Int64ObservableCounter
}

var newGpuIOObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Received and transmitted bytes by the GPU."),
	metric.WithUnit("By"),
}

func NewGpuIOObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (GpuIOObservable, error) {
	_ = "STUB: not implemented"
	return *new(GpuIOObservable), nil
}

func (m GpuIOObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (GpuIOObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuIOObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuIOObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (GpuIOObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIOObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIOObservable) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIOObservable) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIOObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIOObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIOObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIOObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuIOObservable) AttrVendor(val string) attribute.KeyValue {
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

type GpuMemoryLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newGpuMemoryLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Size of the GPU memory."),
	metric.WithUnit("By"),
}

func NewGpuMemoryLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (GpuMemoryLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(GpuMemoryLimitObservable), nil
}

func (m GpuMemoryLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (GpuMemoryLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryLimitObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimitObservable) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimitObservable) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimitObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimitObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimitObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimitObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryLimitObservable) AttrVendor(val string) attribute.KeyValue {
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

type GpuMemoryUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newGpuMemoryUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("GPU memory used."),
	metric.WithUnit("By"),
}

func NewGpuMemoryUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (GpuMemoryUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(GpuMemoryUsageObservable), nil
}

func (m GpuMemoryUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (GpuMemoryUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUsageObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsageObservable) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsageObservable) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsageObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsageObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsageObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsageObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUsageObservable) AttrVendor(val string) attribute.KeyValue {
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

type GpuMemoryUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newGpuMemoryUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Fraction of GPU memory used."),
	metric.WithUnit("1"),
}

func NewGpuMemoryUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (GpuMemoryUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(GpuMemoryUtilizationObservable), nil
}

func (m GpuMemoryUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (GpuMemoryUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUtilizationObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (GpuMemoryUtilizationObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilizationObservable) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilizationObservable) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilizationObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilizationObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilizationObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilizationObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuMemoryUtilizationObservable) AttrVendor(val string) attribute.KeyValue {
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

type GpuUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newGpuUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Fraction of time spent in a specific task."),
	metric.WithUnit("1"),
}

func NewGpuUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (GpuUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(GpuUtilizationObservable), nil
}

func (m GpuUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (GpuUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (GpuUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (GpuUtilizationObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (GpuUtilizationObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilizationObservable) AttrDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilizationObservable) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilizationObservable) AttrGpuTask(val GpuTaskAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilizationObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilizationObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilizationObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilizationObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (GpuUtilizationObservable) AttrVendor(val string) attribute.KeyValue {
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

type HostAmbientTemperatureObservable struct {
	metric.Int64ObservableGauge
}

var newHostAmbientTemperatureObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Ambient (external) temperature of the physical host."),
	metric.WithUnit("Cel"),
}

func NewHostAmbientTemperatureObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (HostAmbientTemperatureObservable, error) {
	_ = "STUB: not implemented"
	return *new(HostAmbientTemperatureObservable), nil
}

func (m HostAmbientTemperatureObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (HostAmbientTemperatureObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HostAmbientTemperatureObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HostAmbientTemperatureObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (HostAmbientTemperatureObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostAmbientTemperatureObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostAmbientTemperatureObservable) AttrParent(val string) attribute.KeyValue {
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

type HostEnergyObservable struct {
	metric.Int64ObservableCounter
}

var newHostEnergyObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Total energy consumed by the entire physical host, in joules."),
	metric.WithUnit("J"),
}

func NewHostEnergyObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (HostEnergyObservable, error) {
	_ = "STUB: not implemented"
	return *new(HostEnergyObservable), nil
}

func (m HostEnergyObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (HostEnergyObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HostEnergyObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HostEnergyObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (HostEnergyObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostEnergyObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostEnergyObservable) AttrParent(val string) attribute.KeyValue {
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

type HostHeatingMarginObservable struct {
	metric.Int64ObservableGauge
}

var newHostHeatingMarginObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("By how many degrees Celsius the temperature of the physical host can be increased, before reaching a warning threshold on one of the internal sensors."),
	metric.WithUnit("Cel"),
}

func NewHostHeatingMarginObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (HostHeatingMarginObservable, error) {
	_ = "STUB: not implemented"
	return *new(HostHeatingMarginObservable), nil
}

func (m HostHeatingMarginObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (HostHeatingMarginObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HostHeatingMarginObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HostHeatingMarginObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (HostHeatingMarginObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostHeatingMarginObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostHeatingMarginObservable) AttrParent(val string) attribute.KeyValue {
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

type HostPowerObservable struct {
	metric.Int64ObservableGauge
}

var newHostPowerObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Instantaneous power consumed by the entire physical host in Watts (`hw.host.energy` is preferred)."),
	metric.WithUnit("W"),
}

func NewHostPowerObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (HostPowerObservable, error) {
	_ = "STUB: not implemented"
	return *new(HostPowerObservable), nil
}

func (m HostPowerObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (HostPowerObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HostPowerObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HostPowerObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (HostPowerObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostPowerObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostPowerObservable) AttrParent(val string) attribute.KeyValue {
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

type LogicalDiskLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newLogicalDiskLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Size of the logical disk."),
	metric.WithUnit("By"),
}

func NewLogicalDiskLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (LogicalDiskLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(LogicalDiskLimitObservable), nil
}

func (m LogicalDiskLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (LogicalDiskLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskLimitObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskLimitObservable) AttrLogicalDiskRaidLevel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskLimitObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskLimitObservable) AttrParent(val string) attribute.KeyValue {
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

type LogicalDiskUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newLogicalDiskUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Logical disk space usage."),
	metric.WithUnit("By"),
}

func NewLogicalDiskUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (LogicalDiskUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(LogicalDiskUsageObservable), nil
}

func (m LogicalDiskUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (LogicalDiskUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUsageObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUsageObservable) AttrLogicalDiskState(val LogicalDiskStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUsageObservable) AttrLogicalDiskRaidLevel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUsageObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUsageObservable) AttrParent(val string) attribute.KeyValue {
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

type LogicalDiskUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newLogicalDiskUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Logical disk space utilization as a fraction."),
	metric.WithUnit("1"),
}

func NewLogicalDiskUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (LogicalDiskUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(LogicalDiskUtilizationObservable), nil
}

func (m LogicalDiskUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (LogicalDiskUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUtilizationObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (LogicalDiskUtilizationObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUtilizationObservable) AttrLogicalDiskState(val LogicalDiskStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUtilizationObservable) AttrLogicalDiskRaidLevel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUtilizationObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (LogicalDiskUtilizationObservable) AttrParent(val string) attribute.KeyValue {
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

type MemorySizeObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newMemorySizeObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Size of the memory module."),
	metric.WithUnit("By"),
}

func NewMemorySizeObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemorySizeObservable, error) {
	_ = "STUB: not implemented"
	return *new(MemorySizeObservable), nil
}

func (m MemorySizeObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemorySizeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (MemorySizeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemorySizeObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (MemorySizeObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySizeObservable) AttrMemoryType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySizeObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySizeObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySizeObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySizeObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (MemorySizeObservable) AttrVendor(val string) attribute.KeyValue {
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

type NetworkBandwidthLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNetworkBandwidthLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Link speed."),
	metric.WithUnit("By/s"),
}

func NewNetworkBandwidthLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NetworkBandwidthLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(NetworkBandwidthLimitObservable), nil
}

func (m NetworkBandwidthLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NetworkBandwidthLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkBandwidthLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkBandwidthLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NetworkBandwidthLimitObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimitObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimitObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimitObservable) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimitObservable) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimitObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimitObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthLimitObservable) AttrVendor(val string) attribute.KeyValue {
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

type NetworkBandwidthUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newNetworkBandwidthUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Utilization of the network bandwidth as a fraction."),
	metric.WithUnit("1"),
}

func NewNetworkBandwidthUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (NetworkBandwidthUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(NetworkBandwidthUtilizationObservable), nil
}

func (m NetworkBandwidthUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (NetworkBandwidthUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkBandwidthUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkBandwidthUtilizationObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (NetworkBandwidthUtilizationObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilizationObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilizationObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilizationObservable) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilizationObservable) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilizationObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilizationObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkBandwidthUtilizationObservable) AttrVendor(val string) attribute.KeyValue {
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

type NetworkIOObservable struct {
	metric.Int64ObservableCounter
}

var newNetworkIOObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Received and transmitted network traffic in bytes."),
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

func (NetworkIOObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIOObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIOObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIOObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIOObservable) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIOObservable) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIOObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIOObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkIOObservable) AttrVendor(val string) attribute.KeyValue {
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

type NetworkPacketsObservable struct {
	metric.Int64ObservableCounter
}

var newNetworkPacketsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Received and transmitted network traffic in packets (or frames)."),
	metric.WithUnit("{packet}"),
}

func NewNetworkPacketsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (NetworkPacketsObservable, error) {
	_ = "STUB: not implemented"
	return *new(NetworkPacketsObservable), nil
}

func (m NetworkPacketsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (NetworkPacketsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NetworkPacketsObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketsObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketsObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketsObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketsObservable) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketsObservable) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketsObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketsObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkPacketsObservable) AttrVendor(val string) attribute.KeyValue {
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

type NetworkUpObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNetworkUpObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Link status: `1` (up) or `0` (down)."),
	metric.WithUnit("1"),
}

func NewNetworkUpObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NetworkUpObservable, error) {
	_ = "STUB: not implemented"
	return *new(NetworkUpObservable), nil
}

func (m NetworkUpObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NetworkUpObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NetworkUpObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NetworkUpObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NetworkUpObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUpObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUpObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUpObservable) AttrNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUpObservable) AttrNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUpObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUpObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NetworkUpObservable) AttrVendor(val string) attribute.KeyValue {
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

type PhysicalDiskEnduranceUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newPhysicalDiskEnduranceUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Endurance remaining for this SSD disk."),
	metric.WithUnit("1"),
}

func NewPhysicalDiskEnduranceUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (PhysicalDiskEnduranceUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(PhysicalDiskEnduranceUtilizationObservable), nil
}

func (m PhysicalDiskEnduranceUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (PhysicalDiskEnduranceUtilizationObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (PhysicalDiskEnduranceUtilizationObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (PhysicalDiskEnduranceUtilizationObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (PhysicalDiskEnduranceUtilizationObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilizationObservable) AttrPhysicalDiskState(val PhysicalDiskStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilizationObservable) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilizationObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilizationObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilizationObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilizationObservable) AttrPhysicalDiskType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilizationObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskEnduranceUtilizationObservable) AttrVendor(val string) attribute.KeyValue {
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

type PhysicalDiskSizeObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPhysicalDiskSizeObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Size of the disk."),
	metric.WithUnit("By"),
}

func NewPhysicalDiskSizeObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PhysicalDiskSizeObservable, error) {
	_ = "STUB: not implemented"
	return *new(PhysicalDiskSizeObservable), nil
}

func (m PhysicalDiskSizeObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PhysicalDiskSizeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSizeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSizeObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSizeObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSizeObservable) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSizeObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSizeObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSizeObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSizeObservable) AttrPhysicalDiskType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSizeObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSizeObservable) AttrVendor(val string) attribute.KeyValue {
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

type PhysicalDiskSmartObservable struct {
	metric.Int64ObservableGauge
}

var newPhysicalDiskSmartObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Value of the corresponding [S.M.A.R.T.](https://wikipedia.org/wiki/S.M.A.R.T.) (Self-Monitoring, Analysis, and Reporting Technology) attribute."),
	metric.WithUnit("1"),
}

func NewPhysicalDiskSmartObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (PhysicalDiskSmartObservable, error) {
	_ = "STUB: not implemented"
	return *new(PhysicalDiskSmartObservable), nil
}

func (m PhysicalDiskSmartObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (PhysicalDiskSmartObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSmartObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSmartObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PhysicalDiskSmartObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmartObservable) AttrFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmartObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmartObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmartObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmartObservable) AttrPhysicalDiskSmartAttribute(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmartObservable) AttrPhysicalDiskType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmartObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PhysicalDiskSmartObservable) AttrVendor(val string) attribute.KeyValue {
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

type PowerObservable struct {
	metric.Int64ObservableGauge
}

var newPowerObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Instantaneous power consumed by the component."),
	metric.WithUnit("W"),
}

func NewPowerObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (PowerObservable, error) {
	_ = "STUB: not implemented"
	return *new(PowerObservable), nil
}

func (m PowerObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (PowerObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PowerObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PowerObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PowerObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerObservable) AttrType(val TypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerObservable) AttrParent(val string) attribute.KeyValue {
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

type PowerSupplyLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPowerSupplyLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Maximum power output of the power supply."),
	metric.WithUnit("W"),
}

func NewPowerSupplyLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PowerSupplyLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(PowerSupplyLimitObservable), nil
}

func (m PowerSupplyLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PowerSupplyLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyLimitObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimitObservable) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimitObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimitObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimitObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimitObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyLimitObservable) AttrVendor(val string) attribute.KeyValue {
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

type PowerSupplyUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPowerSupplyUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Current power output of the power supply."),
	metric.WithUnit("W"),
}

func NewPowerSupplyUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PowerSupplyUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(PowerSupplyUsageObservable), nil
}

func (m PowerSupplyUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PowerSupplyUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUsageObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUsageObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUsageObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUsageObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUsageObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUsageObservable) AttrVendor(val string) attribute.KeyValue {
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

type PowerSupplyUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newPowerSupplyUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Utilization of the power supply as a fraction of its maximum output."),
	metric.WithUnit("1"),
}

func NewPowerSupplyUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (PowerSupplyUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(PowerSupplyUtilizationObservable), nil
}

func (m PowerSupplyUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (PowerSupplyUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUtilizationObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PowerSupplyUtilizationObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUtilizationObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUtilizationObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUtilizationObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUtilizationObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PowerSupplyUtilizationObservable) AttrVendor(val string) attribute.KeyValue {
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

type StatusObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newStatusObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Operational status: `1` (true) or `0` (false) for each of the possible states."),
	metric.WithUnit("1"),
}

func NewStatusObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (StatusObservable, error) {
	_ = "STUB: not implemented"
	return *new(StatusObservable), nil
}

func (m StatusObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (StatusObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (StatusObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatusObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (StatusObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (StatusObservable) AttrState(val StateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (StatusObservable) AttrType(val TypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (StatusObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (StatusObservable) AttrParent(val string) attribute.KeyValue {
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

type TapeDriveOperationsObservable struct {
	metric.Int64ObservableCounter
}

var newTapeDriveOperationsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Operations performed by the tape drive."),
	metric.WithUnit("{operation}"),
}

func NewTapeDriveOperationsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (TapeDriveOperationsObservable, error) {
	_ = "STUB: not implemented"
	return *new(TapeDriveOperationsObservable), nil
}

func (m TapeDriveOperationsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (TapeDriveOperationsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (TapeDriveOperationsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (TapeDriveOperationsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (TapeDriveOperationsObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperationsObservable) AttrModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperationsObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperationsObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperationsObservable) AttrSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperationsObservable) AttrTapeDriveOperationType(val TapeDriveOperationTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TapeDriveOperationsObservable) AttrVendor(val string) attribute.KeyValue {
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

type TemperatureObservable struct {
	metric.Int64ObservableGauge
}

var newTemperatureObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Temperature in degrees Celsius."),
	metric.WithUnit("Cel"),
}

func NewTemperatureObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (TemperatureObservable, error) {
	_ = "STUB: not implemented"
	return *new(TemperatureObservable), nil
}

func (m TemperatureObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (TemperatureObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (TemperatureObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (TemperatureObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (TemperatureObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureObservable) AttrSensorLocation(val string) attribute.KeyValue {
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

type TemperatureLimitObservable struct {
	metric.Int64ObservableGauge
}

var newTemperatureLimitObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Temperature limit in degrees Celsius."),
	metric.WithUnit("Cel"),
}

func NewTemperatureLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (TemperatureLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(TemperatureLimitObservable), nil
}

func (m TemperatureLimitObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (TemperatureLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (TemperatureLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (TemperatureLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (TemperatureLimitObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureLimitObservable) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureLimitObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureLimitObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (TemperatureLimitObservable) AttrSensorLocation(val string) attribute.KeyValue {
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

type VoltageObservable struct {
	metric.Int64ObservableGauge
}

var newVoltageObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Voltage measured by the sensor."),
	metric.WithUnit("V"),
}

func NewVoltageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (VoltageObservable, error) {
	_ = "STUB: not implemented"
	return *new(VoltageObservable), nil
}

func (m VoltageObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (VoltageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (VoltageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (VoltageObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (VoltageObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageObservable) AttrSensorLocation(val string) attribute.KeyValue {
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

type VoltageLimitObservable struct {
	metric.Int64ObservableGauge
}

var newVoltageLimitObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Voltage limit in Volts."),
	metric.WithUnit("V"),
}

func NewVoltageLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (VoltageLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(VoltageLimitObservable), nil
}

func (m VoltageLimitObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (VoltageLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (VoltageLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (VoltageLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (VoltageLimitObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageLimitObservable) AttrLimitType(val LimitTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageLimitObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageLimitObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageLimitObservable) AttrSensorLocation(val string) attribute.KeyValue {
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

type VoltageNominalObservable struct {
	metric.Int64ObservableGauge
}

var newVoltageNominalObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Nominal (expected) voltage."),
	metric.WithUnit("V"),
}

func NewVoltageNominalObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (VoltageNominalObservable, error) {
	_ = "STUB: not implemented"
	return *new(VoltageNominalObservable), nil
}

func (m VoltageNominalObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (VoltageNominalObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (VoltageNominalObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (VoltageNominalObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (VoltageNominalObservable) AttrID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageNominalObservable) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageNominalObservable) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (VoltageNominalObservable) AttrSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
