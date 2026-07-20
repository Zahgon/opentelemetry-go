package rpcconv

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

type NetworkTransportAttr string

var (
	NetworkTransportTCP NetworkTransportAttr = "tcp"

	NetworkTransportUDP NetworkTransportAttr = "udp"

	NetworkTransportPipe NetworkTransportAttr = "pipe"

	NetworkTransportUnix NetworkTransportAttr = "unix"

	NetworkTransportQUIC NetworkTransportAttr = "quic"
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
	metric.WithDescription("Measures the duration of outbound remote procedure calls (RPC)."),
	metric.WithUnit("s"),
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

func (ClientCallDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientCallDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientCallDuration) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientRequestSize struct {
	metric.Int64Histogram
}

var newClientRequestSizeOpts = []metric.Int64HistogramOption{
	metric.WithDescription("Measures the size of RPC request messages (uncompressed)."),
	metric.WithUnit("By"),
}

func NewClientRequestSize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientRequestSize, error) {
	_ = "STUB: not implemented"
	return *new(ClientRequestSize), nil
}

func (m ClientRequestSize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientRequestSize) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientRequestSize) Record(
	ctx context.Context,
	val int64,
	systemName SystemNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientRequestSize) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientRequestSize) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestSize) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestSize) AttrResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestSize) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestSize) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestSize) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestSize) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestSize) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientResponseSize struct {
	metric.Int64Histogram
}

var newClientResponseSizeOpts = []metric.Int64HistogramOption{
	metric.WithDescription("Measures the size of RPC response messages (uncompressed)."),
	metric.WithUnit("By"),
}

func NewClientResponseSize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientResponseSize, error) {
	_ = "STUB: not implemented"
	return *new(ClientResponseSize), nil
}

func (m ClientResponseSize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientResponseSize) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientResponseSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientResponseSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientResponseSize) Record(
	ctx context.Context,
	val int64,
	systemName SystemNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientResponseSize) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientResponseSize) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseSize) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseSize) AttrResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseSize) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseSize) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseSize) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseSize) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseSize) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerCallDuration struct {
	metric.Float64Histogram
}

var newServerCallDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Measures the duration of inbound remote procedure calls (RPC)."),
	metric.WithUnit("s"),
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

func (ServerCallDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerCallDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerCallDuration) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
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

type ServerRequestSize struct {
	metric.Int64Histogram
}

var newServerRequestSizeOpts = []metric.Int64HistogramOption{
	metric.WithDescription("Measures the size of RPC request messages (uncompressed)."),
	metric.WithUnit("By"),
}

func NewServerRequestSize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ServerRequestSize, error) {
	_ = "STUB: not implemented"
	return *new(ServerRequestSize), nil
}

func (m ServerRequestSize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ServerRequestSize) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerRequestSize) Record(
	ctx context.Context,
	val int64,
	systemName SystemNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerRequestSize) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerRequestSize) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestSize) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestSize) AttrResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestSize) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestSize) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestSize) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestSize) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestSize) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerResponseSize struct {
	metric.Int64Histogram
}

var newServerResponseSizeOpts = []metric.Int64HistogramOption{
	metric.WithDescription("Measures the size of RPC response messages (uncompressed)."),
	metric.WithUnit("By"),
}

func NewServerResponseSize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ServerResponseSize, error) {
	_ = "STUB: not implemented"
	return *new(ServerResponseSize), nil
}

func (m ServerResponseSize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ServerResponseSize) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerResponseSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerResponseSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerResponseSize) Record(
	ctx context.Context,
	val int64,
	systemName SystemNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerResponseSize) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerResponseSize) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseSize) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseSize) AttrResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseSize) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseSize) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseSize) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseSize) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseSize) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
