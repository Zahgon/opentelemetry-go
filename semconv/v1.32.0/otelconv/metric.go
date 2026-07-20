package otelconv

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

type ComponentTypeAttr string

var (
	ComponentTypeBatchingSpanProcessor ComponentTypeAttr = "batching_span_processor"

	ComponentTypeSimpleSpanProcessor ComponentTypeAttr = "simple_span_processor"

	ComponentTypeBatchingLogProcessor ComponentTypeAttr = "batching_log_processor"

	ComponentTypeSimpleLogProcessor ComponentTypeAttr = "simple_log_processor"

	ComponentTypeOtlpGRPCSpanExporter ComponentTypeAttr = "otlp_grpc_span_exporter"

	ComponentTypeOtlpHTTPSpanExporter ComponentTypeAttr = "otlp_http_span_exporter"

	ComponentTypeOtlpHTTPJSONSpanExporter ComponentTypeAttr = "otlp_http_json_span_exporter"

	ComponentTypeOtlpGRPCLogExporter ComponentTypeAttr = "otlp_grpc_log_exporter"

	ComponentTypeOtlpHTTPLogExporter ComponentTypeAttr = "otlp_http_log_exporter"

	ComponentTypeOtlpHTTPJSONLogExporter ComponentTypeAttr = "otlp_http_json_log_exporter"
)

type SpanSamplingResultAttr string

var (
	SpanSamplingResultDrop SpanSamplingResultAttr = "DROP"

	SpanSamplingResultRecordOnly SpanSamplingResultAttr = "RECORD_ONLY"

	SpanSamplingResultRecordAndSample SpanSamplingResultAttr = "RECORD_AND_SAMPLE"
)

type SDKExporterLogExported struct {
	metric.Int64Counter
}

func NewSDKExporterLogExported(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKExporterLogExported, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterLogExported), nil
}

func (m SDKExporterLogExported) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKExporterLogExported) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogExported) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogExported) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKExporterLogExported) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKExporterLogExported) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogExported) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogExported) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogExported) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogExported) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterLogInflight struct {
	metric.Int64UpDownCounter
}

func NewSDKExporterLogInflight(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKExporterLogInflight, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterLogInflight), nil
}

func (m SDKExporterLogInflight) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKExporterLogInflight) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogInflight) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogInflight) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKExporterLogInflight) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKExporterLogInflight) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogInflight) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogInflight) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogInflight) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterSpanExportedCount struct {
	metric.Int64Counter
}

func NewSDKExporterSpanExportedCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKExporterSpanExportedCount, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterSpanExportedCount), nil
}

func (m SDKExporterSpanExportedCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKExporterSpanExportedCount) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanExportedCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanExportedCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKExporterSpanExportedCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKExporterSpanExportedCount) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExportedCount) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExportedCount) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExportedCount) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExportedCount) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterSpanInflightCount struct {
	metric.Int64UpDownCounter
}

func NewSDKExporterSpanInflightCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKExporterSpanInflightCount, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterSpanInflightCount), nil
}

func (m SDKExporterSpanInflightCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKExporterSpanInflightCount) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanInflightCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanInflightCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKExporterSpanInflightCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKExporterSpanInflightCount) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanInflightCount) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanInflightCount) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanInflightCount) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKLogCreated struct {
	metric.Int64Counter
}

func NewSDKLogCreated(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKLogCreated, error) {
	_ = "STUB: not implemented"
	return *new(SDKLogCreated), nil
}

func (m SDKLogCreated) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKLogCreated) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKLogCreated) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKLogCreated) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKLogCreated) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type SDKProcessorLogProcessed struct {
	metric.Int64Counter
}

func NewSDKProcessorLogProcessed(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKProcessorLogProcessed, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorLogProcessed), nil
}

func (m SDKProcessorLogProcessed) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKProcessorLogProcessed) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogProcessed) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogProcessed) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKProcessorLogProcessed) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKProcessorLogProcessed) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorLogProcessed) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorLogProcessed) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKProcessorLogQueueCapacity struct {
	metric.Int64UpDownCounter
}

func NewSDKProcessorLogQueueCapacity(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKProcessorLogQueueCapacity, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorLogQueueCapacity), nil
}

func (m SDKProcessorLogQueueCapacity) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKProcessorLogQueueCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogQueueCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogQueueCapacity) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKProcessorLogQueueCapacity) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKProcessorLogQueueCapacity) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorLogQueueCapacity) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKProcessorLogQueueSize struct {
	metric.Int64UpDownCounter
}

func NewSDKProcessorLogQueueSize(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKProcessorLogQueueSize, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorLogQueueSize), nil
}

func (m SDKProcessorLogQueueSize) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKProcessorLogQueueSize) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogQueueSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogQueueSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKProcessorLogQueueSize) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKProcessorLogQueueSize) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorLogQueueSize) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKProcessorSpanProcessedCount struct {
	metric.Int64Counter
}

func NewSDKProcessorSpanProcessedCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKProcessorSpanProcessedCount, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorSpanProcessedCount), nil
}

func (m SDKProcessorSpanProcessedCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKProcessorSpanProcessedCount) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanProcessedCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanProcessedCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKProcessorSpanProcessedCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKProcessorSpanProcessedCount) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanProcessedCount) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanProcessedCount) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKProcessorSpanQueueCapacity struct {
	metric.Int64UpDownCounter
}

func NewSDKProcessorSpanQueueCapacity(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKProcessorSpanQueueCapacity, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorSpanQueueCapacity), nil
}

func (m SDKProcessorSpanQueueCapacity) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKProcessorSpanQueueCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueCapacity) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKProcessorSpanQueueCapacity) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKProcessorSpanQueueCapacity) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanQueueCapacity) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKProcessorSpanQueueSize struct {
	metric.Int64UpDownCounter
}

func NewSDKProcessorSpanQueueSize(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKProcessorSpanQueueSize, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorSpanQueueSize), nil
}

func (m SDKProcessorSpanQueueSize) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKProcessorSpanQueueSize) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKProcessorSpanQueueSize) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKProcessorSpanQueueSize) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanQueueSize) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKSpanEndedCount struct {
	metric.Int64Counter
}

func NewSDKSpanEndedCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKSpanEndedCount, error) {
	_ = "STUB: not implemented"
	return *new(SDKSpanEndedCount), nil
}

func (m SDKSpanEndedCount) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKSpanEndedCount) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanEndedCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanEndedCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKSpanEndedCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKSpanEndedCount) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKSpanLiveCount struct {
	metric.Int64UpDownCounter
}

func NewSDKSpanLiveCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKSpanLiveCount, error) {
	_ = "STUB: not implemented"
	return *new(SDKSpanLiveCount), nil
}

func (m SDKSpanLiveCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKSpanLiveCount) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanLiveCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanLiveCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKSpanLiveCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKSpanLiveCount) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
