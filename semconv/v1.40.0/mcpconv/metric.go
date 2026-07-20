package mcpconv

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

type GenAIOperationNameAttr string

var (
	GenAIOperationNameChat GenAIOperationNameAttr = "chat"

	GenAIOperationNameGenerateContent GenAIOperationNameAttr = "generate_content"

	GenAIOperationNameTextCompletion GenAIOperationNameAttr = "text_completion"

	GenAIOperationNameEmbeddings GenAIOperationNameAttr = "embeddings"

	GenAIOperationNameRetrieval GenAIOperationNameAttr = "retrieval"

	GenAIOperationNameCreateAgent GenAIOperationNameAttr = "create_agent"

	GenAIOperationNameInvokeAgent GenAIOperationNameAttr = "invoke_agent"

	GenAIOperationNameExecuteTool GenAIOperationNameAttr = "execute_tool"
)

type MethodNameAttr string

var (
	MethodNameNotificationsCancelled MethodNameAttr = "notifications/cancelled"

	MethodNameInitialize MethodNameAttr = "initialize"

	MethodNameNotificationsInitialized MethodNameAttr = "notifications/initialized"

	MethodNameNotificationsProgress MethodNameAttr = "notifications/progress"

	MethodNamePing MethodNameAttr = "ping"

	MethodNameResourcesList MethodNameAttr = "resources/list"

	MethodNameResourcesTemplatesList MethodNameAttr = "resources/templates/list"

	MethodNameResourcesRead MethodNameAttr = "resources/read"

	MethodNameNotificationsResourcesListChanged MethodNameAttr = "notifications/resources/list_changed"

	MethodNameResourcesSubscribe MethodNameAttr = "resources/subscribe"

	MethodNameResourcesUnsubscribe MethodNameAttr = "resources/unsubscribe"

	MethodNameNotificationsResourcesUpdated MethodNameAttr = "notifications/resources/updated"

	MethodNamePromptsList MethodNameAttr = "prompts/list"

	MethodNamePromptsGet MethodNameAttr = "prompts/get"

	MethodNameNotificationsPromptsListChanged MethodNameAttr = "notifications/prompts/list_changed"

	MethodNameToolsList MethodNameAttr = "tools/list"

	MethodNameToolsCall MethodNameAttr = "tools/call"

	MethodNameNotificationsToolsListChanged MethodNameAttr = "notifications/tools/list_changed"

	MethodNameLoggingSetLevel MethodNameAttr = "logging/setLevel"

	MethodNameNotificationsMessage MethodNameAttr = "notifications/message"

	MethodNameSamplingCreateMessage MethodNameAttr = "sampling/createMessage"

	MethodNameCompletionComplete MethodNameAttr = "completion/complete"

	MethodNameRootsList MethodNameAttr = "roots/list"

	MethodNameNotificationsRootsListChanged MethodNameAttr = "notifications/roots/list_changed"

	MethodNameElicitationCreate MethodNameAttr = "elicitation/create"
)

type NetworkTransportAttr string

var (
	NetworkTransportTCP NetworkTransportAttr = "tcp"

	NetworkTransportUDP NetworkTransportAttr = "udp"

	NetworkTransportPipe NetworkTransportAttr = "pipe"

	NetworkTransportUnix NetworkTransportAttr = "unix"

	NetworkTransportQUIC NetworkTransportAttr = "quic"
)

type ClientOperationDuration struct {
	metric.Float64Histogram
}

var newClientOperationDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The duration of the MCP request or notification as observed on the sender from the time it was sent until the response or ack is received."),
	metric.WithUnit("s"),
}

func NewClientOperationDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientOperationDuration, error) {
	_ = "STUB: not implemented"
	return *new(ClientOperationDuration), nil
}

func (m ClientOperationDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientOperationDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientOperationDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientOperationDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientOperationDuration) Record(
	ctx context.Context,
	val float64,
	methodName MethodNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientOperationDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientOperationDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrGenAIPromptName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrGenAIToolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrRPCResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrGenAIOperationName(val GenAIOperationNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrJSONRPCProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrResourceURI(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientSessionDuration struct {
	metric.Float64Histogram
}

var newClientSessionDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The duration of the MCP session as observed on the MCP client."),
	metric.WithUnit("s"),
}

func NewClientSessionDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientSessionDuration, error) {
	_ = "STUB: not implemented"
	return *new(ClientSessionDuration), nil
}

func (m ClientSessionDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientSessionDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientSessionDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientSessionDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientSessionDuration) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientSessionDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientSessionDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSessionDuration) AttrJSONRPCProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSessionDuration) AttrProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSessionDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSessionDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSessionDuration) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSessionDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSessionDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerOperationDuration struct {
	metric.Float64Histogram
}

var newServerOperationDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("MCP request or notification duration as observed on the receiver from the time it was received until the result or ack is sent."),
	metric.WithUnit("s"),
}

func NewServerOperationDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerOperationDuration, error) {
	_ = "STUB: not implemented"
	return *new(ServerOperationDuration), nil
}

func (m ServerOperationDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ServerOperationDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerOperationDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerOperationDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerOperationDuration) Record(
	ctx context.Context,
	val float64,
	methodName MethodNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerOperationDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerOperationDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrGenAIPromptName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrGenAIToolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrRPCResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrGenAIOperationName(val GenAIOperationNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrJSONRPCProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerOperationDuration) AttrResourceURI(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerSessionDuration struct {
	metric.Float64Histogram
}

var newServerSessionDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The duration of the MCP session as observed on the MCP server."),
	metric.WithUnit("s"),
}

func NewServerSessionDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerSessionDuration, error) {
	_ = "STUB: not implemented"
	return *new(ServerSessionDuration), nil
}

func (m ServerSessionDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ServerSessionDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerSessionDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerSessionDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerSessionDuration) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerSessionDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerSessionDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerSessionDuration) AttrJSONRPCProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerSessionDuration) AttrProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerSessionDuration) AttrNetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerSessionDuration) AttrNetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerSessionDuration) AttrNetworkTransport(val NetworkTransportAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
