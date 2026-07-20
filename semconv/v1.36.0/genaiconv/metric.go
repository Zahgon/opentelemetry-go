package genaiconv

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

type OperationNameAttr string

var (
	OperationNameChat OperationNameAttr = "chat"

	OperationNameGenerateContent OperationNameAttr = "generate_content"

	OperationNameTextCompletion OperationNameAttr = "text_completion"

	OperationNameEmbeddings OperationNameAttr = "embeddings"

	OperationNameCreateAgent OperationNameAttr = "create_agent"

	OperationNameInvokeAgent OperationNameAttr = "invoke_agent"

	OperationNameExecuteTool OperationNameAttr = "execute_tool"
)

type SystemAttr string

var (
	SystemOpenAI SystemAttr = "openai"

	SystemGCPGenAI SystemAttr = "gcp.gen_ai"

	SystemGCPVertexAI SystemAttr = "gcp.vertex_ai"

	SystemGCPGemini SystemAttr = "gcp.gemini"

	SystemAnthropic SystemAttr = "anthropic"

	SystemCohere SystemAttr = "cohere"

	SystemAzureAIInference SystemAttr = "azure.ai.inference"

	SystemAzureAIOpenAI SystemAttr = "azure.ai.openai"

	SystemIBMWatsonxAI SystemAttr = "ibm.watsonx.ai"

	SystemAWSBedrock SystemAttr = "aws.bedrock"

	SystemPerplexity SystemAttr = "perplexity"

	SystemXai SystemAttr = "xai"

	SystemDeepseek SystemAttr = "deepseek"

	SystemGroq SystemAttr = "groq"

	SystemMistralAI SystemAttr = "mistral_ai"
)

type TokenTypeAttr string

var (
	TokenTypeInput TokenTypeAttr = "input"

	TokenTypeOutput TokenTypeAttr = "output"
)

type ClientOperationDuration struct {
	metric.Float64Histogram
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
	operationName OperationNameAttr,
	system SystemAttr,
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

func (ClientOperationDuration) AttrRequestModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrResponseModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientTokenUsage struct {
	metric.Int64Histogram
}

func NewClientTokenUsage(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientTokenUsage, error) {
	_ = "STUB: not implemented"
	return *new(ClientTokenUsage), nil
}

func (m ClientTokenUsage) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientTokenUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientTokenUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientTokenUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientTokenUsage) Record(
	ctx context.Context,
	val int64,
	operationName OperationNameAttr,
	system SystemAttr,
	tokenType TokenTypeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientTokenUsage) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientTokenUsage) AttrRequestModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientTokenUsage) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientTokenUsage) AttrResponseModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientTokenUsage) AttrServerAddress(val string) attribute.KeyValue {
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
	operationName OperationNameAttr,
	system SystemAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerRequestDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerRequestDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrRequestModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrResponseModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerRequestDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerTimePerOutputToken struct {
	metric.Float64Histogram
}

func NewServerTimePerOutputToken(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerTimePerOutputToken, error) {
	_ = "STUB: not implemented"
	return *new(ServerTimePerOutputToken), nil
}

func (m ServerTimePerOutputToken) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ServerTimePerOutputToken) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerTimePerOutputToken) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerTimePerOutputToken) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerTimePerOutputToken) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	system SystemAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerTimePerOutputToken) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerTimePerOutputToken) AttrRequestModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerTimePerOutputToken) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerTimePerOutputToken) AttrResponseModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerTimePerOutputToken) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServerTimeToFirstToken struct {
	metric.Float64Histogram
}

func NewServerTimeToFirstToken(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerTimeToFirstToken, error) {
	_ = "STUB: not implemented"
	return *new(ServerTimeToFirstToken), nil
}

func (m ServerTimeToFirstToken) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ServerTimeToFirstToken) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerTimeToFirstToken) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerTimeToFirstToken) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerTimeToFirstToken) Record(
	ctx context.Context,
	val float64,
	operationName OperationNameAttr,
	system SystemAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServerTimeToFirstToken) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServerTimeToFirstToken) AttrRequestModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerTimeToFirstToken) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerTimeToFirstToken) AttrResponseModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServerTimeToFirstToken) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
