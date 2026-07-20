package nfsconv

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

type NetworkIODirectionAttr string

var (
	NetworkIODirectionTransmit NetworkIODirectionAttr = "transmit"

	NetworkIODirectionReceive NetworkIODirectionAttr = "receive"
)

type NetworkTransportAttr string

var (
	NetworkTransportTCP NetworkTransportAttr = "tcp"

	NetworkTransportUDP NetworkTransportAttr = "udp"

	NetworkTransportPipe NetworkTransportAttr = "pipe"

	NetworkTransportUnix NetworkTransportAttr = "unix"

	NetworkTransportQUIC NetworkTransportAttr = "quic"
)

type ClientNetCount struct {
	metric.Int64Counter
}

var newClientNetCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS client TCP segments and UDP datagrams handled."),
	metric.WithUnit("{record}"),
}

func NewClientNetCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientNetCount, error) {
	_ = "STUB: not implemented"
	return *new(ClientNetCount), nil
}

func (m ClientNetCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientNetCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientNetCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientNetCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientNetCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientNetCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientNetCount) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientNetTCPConnectionAccepted struct {
	metric.Int64Counter
}

var newClientNetTCPConnectionAcceptedOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS client TCP connections accepted."),
	metric.WithUnit("{connection}"),
}

func NewClientNetTCPConnectionAccepted(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientNetTCPConnectionAccepted, error) {
	_ = "STUB: not implemented"
	return *new(ClientNetTCPConnectionAccepted), nil
}

func (m ClientNetTCPConnectionAccepted) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientNetTCPConnectionAccepted) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientNetTCPConnectionAccepted) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientNetTCPConnectionAccepted) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientNetTCPConnectionAccepted) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClientNetTCPConnectionAccepted) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientOperationCount struct {
	metric.Int64Counter
}

var newClientOperationCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFSv4+ client operations."),
	metric.WithUnit("{operation}"),
}

func NewClientOperationCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientOperationCount, error) {
	_ = "STUB: not implemented"
	return *new(ClientOperationCount), nil
}

func (m ClientOperationCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientOperationCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientOperationCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientOperationCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientOperationCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientOperationCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientOperationCount) AttrOperationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationCount) AttrOncRPCVersion(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientProcedureCount struct {
	metric.Int64Counter
}

var newClientProcedureCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS client procedures."),
	metric.WithUnit("{procedure}"),
}

func NewClientProcedureCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientProcedureCount, error) {
	_ = "STUB: not implemented"
	return *new(ClientProcedureCount), nil
}

func (m ClientProcedureCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientProcedureCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientProcedureCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientProcedureCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientProcedureCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientProcedureCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientProcedureCount) AttrOncRPCProcedureName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientProcedureCount) AttrOncRPCVersion(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientRPCAuthrefreshCount struct {
	metric.Int64Counter
}

var newClientRPCAuthrefreshCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS client RPC authentication refreshes."),
	metric.WithUnit("{authrefresh}"),
}

func NewClientRPCAuthrefreshCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientRPCAuthrefreshCount, error) {
	_ = "STUB: not implemented"
	return *new(ClientRPCAuthrefreshCount), nil
}

func (m ClientRPCAuthrefreshCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientRPCAuthrefreshCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientRPCAuthrefreshCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientRPCAuthrefreshCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientRPCAuthrefreshCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClientRPCAuthrefreshCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientRPCCount struct {
	metric.Int64Counter
}

var newClientRPCCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS client RPCs sent, regardless of whether they're accepted/rejected by the server."),
	metric.WithUnit("{request}"),
}

func NewClientRPCCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientRPCCount, error) {
	_ = "STUB: not implemented"
	return *new(ClientRPCCount), nil
}

func (m ClientRPCCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientRPCCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientRPCCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientRPCCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientRPCCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClientRPCCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientRPCRetransmitCount struct {
	metric.Int64Counter
}

var newClientRPCRetransmitCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS client RPC retransmits."),
	metric.WithUnit("{retransmit}"),
}

func NewClientRPCRetransmitCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientRPCRetransmitCount, error) {
	_ = "STUB: not implemented"
	return *new(ClientRPCRetransmitCount), nil
}

func (m ClientRPCRetransmitCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientRPCRetransmitCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientRPCRetransmitCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientRPCRetransmitCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientRPCRetransmitCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClientRPCRetransmitCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ServerFhStaleCount struct {
	metric.Int64Counter
}

var newServerFhStaleCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS server stale file handles."),
	metric.WithUnit("{fh}"),
}

func NewServerFhStaleCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ServerFhStaleCount, error) {
	_ = "STUB: not implemented"
	return *new(ServerFhStaleCount), nil
}

func (m ServerFhStaleCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ServerFhStaleCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerFhStaleCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerFhStaleCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerFhStaleCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ServerFhStaleCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ServerIO struct {
	metric.Int64Counter
}

var newServerIOOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS server bytes returned to receive and transmit (read and write) requests."),
	metric.WithUnit("By"),
}

func NewServerIO(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ServerIO, error) {
	_ = "STUB: not implemented"
	return *new(ServerIO), nil
}

func (m ServerIO) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ServerIO) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerIO) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerIO) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerIO) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerIO) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerIO) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerNetCount struct {
	metric.Int64Counter
}

var newServerNetCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS server TCP segments and UDP datagrams handled."),
	metric.WithUnit("{record}"),
}

func NewServerNetCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ServerNetCount, error) {
	_ = "STUB: not implemented"
	return *new(ServerNetCount), nil
}

func (m ServerNetCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ServerNetCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerNetCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerNetCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerNetCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerNetCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerNetCount) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerNetTCPConnectionAccepted struct {
	metric.Int64Counter
}

var newServerNetTCPConnectionAcceptedOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS server TCP connections accepted."),
	metric.WithUnit("{connection}"),
}

func NewServerNetTCPConnectionAccepted(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ServerNetTCPConnectionAccepted, error) {
	_ = "STUB: not implemented"
	return *new(ServerNetTCPConnectionAccepted), nil
}

func (m ServerNetTCPConnectionAccepted) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ServerNetTCPConnectionAccepted) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerNetTCPConnectionAccepted) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerNetTCPConnectionAccepted) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerNetTCPConnectionAccepted) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ServerNetTCPConnectionAccepted) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ServerOperationCount struct {
	metric.Int64Counter
}

var newServerOperationCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFSv4+ server operations."),
	metric.WithUnit("{operation}"),
}

func NewServerOperationCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ServerOperationCount, error) {
	_ = "STUB: not implemented"
	return *new(ServerOperationCount), nil
}

func (m ServerOperationCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ServerOperationCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerOperationCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerOperationCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerOperationCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerOperationCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerOperationCount) AttrOperationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationCount) AttrOncRPCVersion(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerProcedureCount struct {
	metric.Int64Counter
}

var newServerProcedureCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS server procedures."),
	metric.WithUnit("{procedure}"),
}

func NewServerProcedureCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ServerProcedureCount, error) {
	_ = "STUB: not implemented"
	return *new(ServerProcedureCount), nil
}

func (m ServerProcedureCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ServerProcedureCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerProcedureCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerProcedureCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerProcedureCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerProcedureCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerProcedureCount) AttrOncRPCProcedureName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerProcedureCount) AttrOncRPCVersion(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerRepcacheRequests struct {
	metric.Int64Counter
}

var newServerRepcacheRequestsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the kernel NFS server reply cache request count by cache hit status."),
	metric.WithUnit("{request}"),
}

func NewServerRepcacheRequests(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ServerRepcacheRequests, error) {
	_ = "STUB: not implemented"
	return *new(ServerRepcacheRequests), nil
}

func (m ServerRepcacheRequests) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ServerRepcacheRequests) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerRepcacheRequests) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerRepcacheRequests) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerRepcacheRequests) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerRepcacheRequests) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerRepcacheRequests) AttrServerRepcacheStatus(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerRPCCount struct {
	metric.Int64Counter
}

var newServerRPCCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("Reports the count of kernel NFS server RPCs handled."),
	metric.WithUnit("{request}"),
}

func NewServerRPCCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ServerRPCCount, error) {
	_ = "STUB: not implemented"
	return *new(ServerRPCCount), nil
}

func (m ServerRPCCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ServerRPCCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerRPCCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerRPCCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerRPCCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerRPCCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerRPCCount) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerThreadCount struct {
	metric.Int64UpDownCounter
}

var newServerThreadCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Reports the count of kernel NFS server available threads."),
	metric.WithUnit("{thread}"),
}

func NewServerThreadCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ServerThreadCount, error) {
	_ = "STUB: not implemented"
	return *new(ServerThreadCount), nil
}

func (m ServerThreadCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ServerThreadCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerThreadCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerThreadCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerThreadCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ServerThreadCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}
