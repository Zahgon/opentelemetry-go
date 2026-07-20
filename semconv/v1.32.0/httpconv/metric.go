package httpconv

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

type ConnectionStateAttr string

var (
	ConnectionStateActive ConnectionStateAttr = "active"

	ConnectionStateIdle ConnectionStateAttr = "idle"
)

type RequestMethodAttr string

var (
	RequestMethodConnect RequestMethodAttr = "CONNECT"

	RequestMethodDelete RequestMethodAttr = "DELETE"

	RequestMethodGet RequestMethodAttr = "GET"

	RequestMethodHead RequestMethodAttr = "HEAD"

	RequestMethodOptions RequestMethodAttr = "OPTIONS"

	RequestMethodPatch RequestMethodAttr = "PATCH"

	RequestMethodPost RequestMethodAttr = "POST"

	RequestMethodPut RequestMethodAttr = "PUT"

	RequestMethodTrace RequestMethodAttr = "TRACE"

	RequestMethodOther RequestMethodAttr = "_OTHER"
)

type UserAgentSyntheticTypeAttr string

var (
	UserAgentSyntheticTypeBot UserAgentSyntheticTypeAttr = "bot"

	UserAgentSyntheticTypeTest UserAgentSyntheticTypeAttr = "test"
)

type ClientActiveRequests struct {
	metric.Int64UpDownCounter
}

func NewClientActiveRequests(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClientActiveRequests, error) {
	_ = "STUB: not implemented"
	return *new(ClientActiveRequests), nil
}

func (m ClientActiveRequests) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClientActiveRequests) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientActiveRequests) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientActiveRequests) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientActiveRequests) Add(
	ctx context.Context,
	incr int64,
	serverAddress string,
	serverPort int,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ClientActiveRequests) AttrURLTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientActiveRequests) AttrRequestMethod(val RequestMethodAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientActiveRequests) AttrURLScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientConnectionDuration struct {
	metric.Float64Histogram
}

func NewClientConnectionDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientConnectionDuration, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionDuration), nil
}

func (m ClientConnectionDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientConnectionDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionDuration) Record(
	ctx context.Context,
	val float64,
	serverAddress string,
	serverPort int,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ClientConnectionDuration) AttrNetworkPeerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConnectionDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConnectionDuration) AttrURLScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientOpenConnections struct {
	metric.Int64UpDownCounter
}

func NewClientOpenConnections(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClientOpenConnections, error) {
	_ = "STUB: not implemented"
	return *new(ClientOpenConnections), nil
}

func (m ClientOpenConnections) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClientOpenConnections) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientOpenConnections) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientOpenConnections) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientOpenConnections) Add(
	ctx context.Context,
	incr int64,
	connectionState ConnectionStateAttr,
	serverAddress string,
	serverPort int,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ClientOpenConnections) AttrNetworkPeerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOpenConnections) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOpenConnections) AttrURLScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientRequestBodySize struct {
	metric.Int64Histogram
}

func NewClientRequestBodySize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientRequestBodySize, error) {
	_ = "STUB: not implemented"
	return *new(ClientRequestBodySize), nil
}

func (m ClientRequestBodySize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientRequestBodySize) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestBodySize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestBodySize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientRequestBodySize) Record(
	ctx context.Context,
	val int64,
	requestMethod RequestMethodAttr,
	serverAddress string,
	serverPort int,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ClientRequestBodySize) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestBodySize) AttrResponseStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestBodySize) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestBodySize) AttrURLTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestBodySize) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestBodySize) AttrURLScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientRequestDuration struct {
	metric.Float64Histogram
}

func NewClientRequestDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientRequestDuration, error) {
	_ = "STUB: not implemented"
	return *new(ClientRequestDuration), nil
}

func (m ClientRequestDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientRequestDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientRequestDuration) Record(
	ctx context.Context,
	val float64,
	requestMethod RequestMethodAttr,
	serverAddress string,
	serverPort int,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ClientRequestDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestDuration) AttrResponseStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestDuration) AttrURLScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientRequestDuration) AttrURLTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientResponseBodySize struct {
	metric.Int64Histogram
}

func NewClientResponseBodySize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientResponseBodySize, error) {
	_ = "STUB: not implemented"
	return *new(ClientResponseBodySize), nil
}

func (m ClientResponseBodySize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientResponseBodySize) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientResponseBodySize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientResponseBodySize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientResponseBodySize) Record(
	ctx context.Context,
	val int64,
	requestMethod RequestMethodAttr,
	serverAddress string,
	serverPort int,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ClientResponseBodySize) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseBodySize) AttrResponseStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseBodySize) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseBodySize) AttrURLTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseBodySize) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseBodySize) AttrURLScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerActiveRequests struct {
	metric.Int64UpDownCounter
}

func NewServerActiveRequests(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ServerActiveRequests, error) {
	_ = "STUB: not implemented"
	return *new(ServerActiveRequests), nil
}

func (m ServerActiveRequests) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ServerActiveRequests) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerActiveRequests) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerActiveRequests) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerActiveRequests) Add(
	ctx context.Context,
	incr int64,
	requestMethod RequestMethodAttr,
	urlScheme string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ServerActiveRequests) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerActiveRequests) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerRequestBodySize struct {
	metric.Int64Histogram
}

func NewServerRequestBodySize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ServerRequestBodySize, error) {
	_ = "STUB: not implemented"
	return *new(ServerRequestBodySize), nil
}

func (m ServerRequestBodySize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ServerRequestBodySize) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestBodySize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestBodySize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerRequestBodySize) Record(
	ctx context.Context,
	val int64,
	requestMethod RequestMethodAttr,
	urlScheme string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ServerRequestBodySize) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestBodySize) AttrResponseStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestBodySize) AttrRoute(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestBodySize) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestBodySize) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestBodySize) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestBodySize) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestBodySize) AttrUserAgentSyntheticType(val UserAgentSyntheticTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerRequestDuration struct {
	metric.Float64Histogram
}

func NewServerRequestDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerRequestDuration, error) {
	_ = "STUB: not implemented"
	return *new(ServerRequestDuration), nil
}

func (m ServerRequestDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ServerRequestDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerRequestDuration) Record(
	ctx context.Context,
	val float64,
	requestMethod RequestMethodAttr,
	urlScheme string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ServerRequestDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrResponseStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrRoute(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrUserAgentSyntheticType(val UserAgentSyntheticTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerResponseBodySize struct {
	metric.Int64Histogram
}

func NewServerResponseBodySize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ServerResponseBodySize, error) {
	_ = "STUB: not implemented"
	return *new(ServerResponseBodySize), nil
}

func (m ServerResponseBodySize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ServerResponseBodySize) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerResponseBodySize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerResponseBodySize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerResponseBodySize) Record(
	ctx context.Context,
	val int64,
	requestMethod RequestMethodAttr,
	urlScheme string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (ServerResponseBodySize) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseBodySize) AttrResponseStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseBodySize) AttrRoute(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseBodySize) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseBodySize) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseBodySize) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseBodySize) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerResponseBodySize) AttrUserAgentSyntheticType(val UserAgentSyntheticTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
