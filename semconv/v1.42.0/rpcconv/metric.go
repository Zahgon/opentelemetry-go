package rpcconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type ErrorTypeAttr string

var (
	ErrorTypeOther ErrorTypeAttr = "_OTHER"
)

type SystemNameAttr string

var (
	SystemNameGRPC SystemNameAttr = "grpc"

	SystemNameDubbo SystemNameAttr = "dubbo"

	SystemNameConnectrpc SystemNameAttr = "connectrpc"

	SystemNameJSONRPC SystemNameAttr = "jsonrpc"
)

type ClientCallDuration struct {
	metric.Float64Histogram
}

var newClientCallDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Measures the duration of an outgoing Remote Procedure Call (RPC)."),
	metric.WithUnit("s"),
	metric.WithExplicitBucketBoundaries([]float64{0.005, 0.01, 0.025, 0.05, 0.075, 0.1, 0.25, 0.5, 0.75, 1, 2.5, 5, 7.5, 10}...),
}

func NewClientCallDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientCallDuration, error) {
	_ = "STUB: not implemented"
	return *new(ClientCallDuration), nil
}

func (m ClientCallDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientCallDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientCallDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientCallDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientCallDuration) Record(
	ctx context.Context,
	val float64,
	systemName SystemNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientCallDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientCallDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientCallDuration) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientCallDuration) AttrResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientCallDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientCallDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerCallDuration struct {
	metric.Float64Histogram
}

var newServerCallDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Measures the duration of an incoming Remote Procedure Call (RPC)."),
	metric.WithUnit("s"),
	metric.WithExplicitBucketBoundaries([]float64{0.005, 0.01, 0.025, 0.05, 0.075, 0.1, 0.25, 0.5, 0.75, 1, 2.5, 5, 7.5, 10}...),
}

func NewServerCallDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerCallDuration, error) {
	_ = "STUB: not implemented"
	return *new(ServerCallDuration), nil
}

func (m ServerCallDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ServerCallDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerCallDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerCallDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerCallDuration) Record(
	ctx context.Context,
	val float64,
	systemName SystemNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerCallDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerCallDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerCallDuration) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerCallDuration) AttrResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerCallDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerCallDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
