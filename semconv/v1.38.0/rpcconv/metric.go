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

type SystemAttr string

var (
	SystemGRPC SystemAttr = "grpc"

	SystemJavaRmi SystemAttr = "java_rmi"

	SystemDotnetWcf SystemAttr = "dotnet_wcf"

	SystemApacheDubbo SystemAttr = "apache_dubbo"

	SystemConnectRPC SystemAttr = "connect_rpc"

	SystemOncRPC SystemAttr = "onc_rpc"

	SystemJSONRPC SystemAttr = "jsonrpc"
)

type ClientDuration struct {
	metric.Float64Histogram
}

var newClientDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Measures the duration of outbound RPC."),
	metric.WithUnit("ms"),
}

func NewClientDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientDuration, error) {
	_ = "STUB: not implemented"
	return *new(ClientDuration), nil
}

func (m ClientDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientDuration) Record(
	ctx context.Context,
	val float64,
	system SystemAttr,
	serverAddress string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientDuration) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientDuration) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientDuration) AttrService(val string) attribute.KeyValue {
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
	system SystemAttr,
	serverAddress string,
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

func (ClientRequestSize) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestSize) AttrService(val string) attribute.KeyValue {
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
	system SystemAttr,
	serverAddress string,
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

func (ClientResponseSize) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseSize) AttrService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerDuration struct {
	metric.Float64Histogram
}

var newServerDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Measures the duration of inbound RPC."),
	metric.WithUnit("ms"),
}

func NewServerDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerDuration, error) {
	_ = "STUB: not implemented"
	return *new(ServerDuration), nil
}

func (m ServerDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ServerDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerDuration) Record(
	ctx context.Context,
	val float64,
	system SystemAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerDuration) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerDuration) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerDuration) AttrService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerDuration) AttrServerPort(val int) attribute.KeyValue {
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
	system SystemAttr,
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

func (ServerRequestSize) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestSize) AttrService(val string) attribute.KeyValue {
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
	system SystemAttr,
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

func (ServerResponseSize) AttrMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseSize) AttrService(val string) attribute.KeyValue {
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
