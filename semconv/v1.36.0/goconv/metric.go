package goconv

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

type MemoryTypeAttr string

var (
	MemoryTypeStack MemoryTypeAttr = "stack"

	MemoryTypeOther MemoryTypeAttr = "other"
)

type ConfigGogc struct {
	metric.Int64ObservableUpDownCounter
}

func NewConfigGogc(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ConfigGogc, error) {
	_ = "STUB: not implemented"
	return *new(ConfigGogc), nil
}

func (m ConfigGogc) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ConfigGogc) Name() string { _ = "STUB: not implemented"; return "" }

func (ConfigGogc) Unit() string { _ = "STUB: not implemented"; return "" }

func (ConfigGogc) Description() string { _ = "STUB: not implemented"; return "" }

type GoroutineCount struct {
	metric.Int64ObservableUpDownCounter
}

func NewGoroutineCount(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (GoroutineCount, error) {
	_ = "STUB: not implemented"
	return *new(GoroutineCount), nil
}

func (m GoroutineCount) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (GoroutineCount) Name() string { _ = "STUB: not implemented"; return "" }

func (GoroutineCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (GoroutineCount) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryAllocated struct {
	metric.Int64ObservableCounter
}

func NewMemoryAllocated(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (MemoryAllocated, error) {
	_ = "STUB: not implemented"
	return *new(MemoryAllocated), nil
}

func (m MemoryAllocated) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (MemoryAllocated) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryAllocated) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryAllocated) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryAllocations struct {
	metric.Int64ObservableCounter
}

func NewMemoryAllocations(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (MemoryAllocations, error) {
	_ = "STUB: not implemented"
	return *new(MemoryAllocations), nil
}

func (m MemoryAllocations) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (MemoryAllocations) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryAllocations) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryAllocations) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryGCGoal struct {
	metric.Int64ObservableUpDownCounter
}

func NewMemoryGCGoal(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryGCGoal, error) {
	_ = "STUB: not implemented"
	return *new(MemoryGCGoal), nil
}

func (m MemoryGCGoal) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryGCGoal) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCGoal) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryGCGoal) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryLimit struct {
	metric.Int64ObservableUpDownCounter
}

func NewMemoryLimit(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryLimit, error) {
	_ = "STUB: not implemented"
	return *new(MemoryLimit), nil
}

func (m MemoryLimit) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryLimit) Description() string { _ = "STUB: not implemented"; return "" }

type MemoryUsed struct {
	metric.Int64ObservableUpDownCounter
}

func NewMemoryUsed(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (MemoryUsed, error) {
	_ = "STUB: not implemented"
	return *new(MemoryUsed), nil
}

func (m MemoryUsed) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (MemoryUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (MemoryUsed) AttrMemoryType(val MemoryTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ProcessorLimit struct {
	metric.Int64ObservableUpDownCounter
}

func NewProcessorLimit(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ProcessorLimit, error) {
	_ = "STUB: not implemented"
	return *new(ProcessorLimit), nil
}

func (m ProcessorLimit) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ProcessorLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (ProcessorLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (ProcessorLimit) Description() string { _ = "STUB: not implemented"; return "" }

type ScheduleDuration struct {
	metric.Float64Histogram
}

func NewScheduleDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ScheduleDuration, error) {
	_ = "STUB: not implemented"
	return *new(ScheduleDuration), nil
}

func (m ScheduleDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ScheduleDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ScheduleDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ScheduleDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ScheduleDuration) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ScheduleDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}
