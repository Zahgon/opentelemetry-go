package azureconv

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

type CosmosDBConsistencyLevelAttr string

var (
	CosmosDBConsistencyLevelStrong CosmosDBConsistencyLevelAttr = "Strong"

	CosmosDBConsistencyLevelBoundedStaleness CosmosDBConsistencyLevelAttr = "BoundedStaleness"

	CosmosDBConsistencyLevelSession CosmosDBConsistencyLevelAttr = "Session"

	CosmosDBConsistencyLevelEventual CosmosDBConsistencyLevelAttr = "Eventual"

	CosmosDBConsistencyLevelConsistentPrefix CosmosDBConsistencyLevelAttr = "ConsistentPrefix"
)

type ErrorTypeAttr string

var ErrorTypeOther ErrorTypeAttr = "_OTHER"

type CosmosDBClientActiveInstanceCount struct {
	metric.Int64UpDownCounter
}

func NewCosmosDBClientActiveInstanceCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (CosmosDBClientActiveInstanceCount, error) {
	_ = "STUB: not implemented"
	return *new(CosmosDBClientActiveInstanceCount), nil
}

func (m CosmosDBClientActiveInstanceCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (CosmosDBClientActiveInstanceCount) Name() string { _ = "STUB: not implemented"; return "" }

func (CosmosDBClientActiveInstanceCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (CosmosDBClientActiveInstanceCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m CosmosDBClientActiveInstanceCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CosmosDBClientActiveInstanceCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (CosmosDBClientActiveInstanceCount) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CosmosDBClientActiveInstanceCount) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type CosmosDBClientOperationRequestCharge struct {
	metric.Int64Histogram
}

func NewCosmosDBClientOperationRequestCharge(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (CosmosDBClientOperationRequestCharge, error) {
	_ = "STUB: not implemented"
	return *new(CosmosDBClientOperationRequestCharge), nil
}

func (m CosmosDBClientOperationRequestCharge) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (CosmosDBClientOperationRequestCharge) Name() string { _ = "STUB: not implemented"; return "" }

func (CosmosDBClientOperationRequestCharge) Unit() string { _ = "STUB: not implemented"; return "" }

func (CosmosDBClientOperationRequestCharge) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m CosmosDBClientOperationRequestCharge) Record(
	ctx context.Context,
	val int64,
	dbOperationName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m CosmosDBClientOperationRequestCharge) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (CosmosDBClientOperationRequestCharge) AttrCosmosDBConsistencyLevel(val CosmosDBConsistencyLevelAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CosmosDBClientOperationRequestCharge) AttrCosmosDBResponseSubStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CosmosDBClientOperationRequestCharge) AttrDBCollectionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CosmosDBClientOperationRequestCharge) AttrDBNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CosmosDBClientOperationRequestCharge) AttrDBResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CosmosDBClientOperationRequestCharge) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CosmosDBClientOperationRequestCharge) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CosmosDBClientOperationRequestCharge) AttrCosmosDBOperationContactedRegions(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (CosmosDBClientOperationRequestCharge) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
