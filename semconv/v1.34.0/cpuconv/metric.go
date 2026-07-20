package cpuconv

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

type ModeAttr string

var (
	ModeUser ModeAttr = "user"

	ModeSystem ModeAttr = "system"

	ModeNice ModeAttr = "nice"

	ModeIdle ModeAttr = "idle"

	ModeIOWait ModeAttr = "iowait"

	ModeInterrupt ModeAttr = "interrupt"

	ModeSteal ModeAttr = "steal"

	ModeKernel ModeAttr = "kernel"
)

type Frequency struct {
	metric.Int64Gauge
}

func NewFrequency(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (Frequency, error) {
	_ = "STUB: not implemented"
	return *new(Frequency), nil
}

func (m Frequency) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (Frequency) Name() string { _ = "STUB: not implemented"; return "" }

func (Frequency) Unit() string { _ = "STUB: not implemented"; return "" }

func (Frequency) Description() string { _ = "STUB: not implemented"; return "" }

func (m Frequency) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (Frequency) AttrLogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Time struct {
	metric.Float64ObservableCounter
}

func NewTime(
	m metric.Meter,
	opt ...metric.Float64ObservableCounterOption,
) (Time, error) {
	_ = "STUB: not implemented"
	return *new(Time), nil
}

func (m Time) Inst() metric.Float64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter)
}

func (Time) Name() string { _ = "STUB: not implemented"; return "" }

func (Time) Unit() string { _ = "STUB: not implemented"; return "" }

func (Time) Description() string { _ = "STUB: not implemented"; return "" }

func (Time) AttrLogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Time) AttrMode(val ModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type Utilization struct {
	metric.Int64Gauge
}

func NewUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (Utilization, error) {
	_ = "STUB: not implemented"
	return *new(Utilization), nil
}

func (m Utilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (Utilization) Name() string { _ = "STUB: not implemented"; return "" }

func (Utilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (Utilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m Utilization) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (Utilization) AttrLogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (Utilization) AttrMode(val ModeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
