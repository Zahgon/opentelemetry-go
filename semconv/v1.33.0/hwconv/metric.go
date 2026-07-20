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

type StateAttr string

var (
	StateOk StateAttr = "ok"

	StateDegraded StateAttr = "degraded"

	StateFailed StateAttr = "failed"
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

type Energy struct {
	metric.Int64Counter
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

type HostAmbientTemperature struct {
	metric.Int64Gauge
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

func (HostPower) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HostPower) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Power struct {
	metric.Int64Gauge
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

func (Power) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Power) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Status struct {
	metric.Int64UpDownCounter
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

func (Status) AttrName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Status) AttrParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
