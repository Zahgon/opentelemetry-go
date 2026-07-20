package signalrconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type ConnectionStatusAttr string

var (
	ConnectionStatusNormalClosure ConnectionStatusAttr = "normal_closure"

	ConnectionStatusTimeout ConnectionStatusAttr = "timeout"

	ConnectionStatusAppShutdown ConnectionStatusAttr = "app_shutdown"
)

type TransportAttr string

var (
	TransportServerSentEvents TransportAttr = "server_sent_events"

	TransportLongPolling TransportAttr = "long_polling"

	TransportWebSockets TransportAttr = "web_sockets"
)

type ServerActiveConnections struct {
	metric.Int64UpDownCounter
}

var newServerActiveConnectionsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of connections that are currently active on the server."),
	metric.WithUnit("{connection}"),
}

func NewServerActiveConnections(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ServerActiveConnections, error) {
	_ = "STUB: not implemented"
	return *new(ServerActiveConnections), nil
}

func (m ServerActiveConnections) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ServerActiveConnections) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerActiveConnections) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerActiveConnections) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerActiveConnections) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerActiveConnections) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerActiveConnections) AttrConnectionStatus(val ConnectionStatusAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerActiveConnections) AttrTransport(val TransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerActiveConnectionsObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newServerActiveConnectionsObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of connections that are currently active on the server."),
	metric.WithUnit("{connection}"),
}

func NewServerActiveConnectionsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ServerActiveConnectionsObservable, error) {
	_ = "STUB: not implemented"
	return *new(ServerActiveConnectionsObservable), nil
}

func (m ServerActiveConnectionsObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ServerActiveConnectionsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerActiveConnectionsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerActiveConnectionsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ServerActiveConnectionsObservable) AttrConnectionStatus(val ConnectionStatusAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerActiveConnectionsObservable) AttrTransport(val TransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerConnectionDuration struct {
	metric.Float64Histogram
}

var newServerConnectionDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The duration of connections on the server."),
	metric.WithUnit("s"),
}

func NewServerConnectionDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerConnectionDuration, error) {
	_ = "STUB: not implemented"
	return *new(ServerConnectionDuration), nil
}

func (m ServerConnectionDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ServerConnectionDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerConnectionDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerConnectionDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerConnectionDuration) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerConnectionDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerConnectionDuration) AttrConnectionStatus(val ConnectionStatusAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerConnectionDuration) AttrTransport(val TransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
