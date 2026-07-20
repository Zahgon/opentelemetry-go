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

	ComponentTypePeriodicMetricReader ComponentTypeAttr = "periodic_metric_reader"

	ComponentTypeOtlpGRPCMetricExporter ComponentTypeAttr = "otlp_grpc_metric_exporter"

	ComponentTypeOtlpHTTPMetricExporter ComponentTypeAttr = "otlp_http_metric_exporter"

	ComponentTypeOtlpHTTPJSONMetricExporter ComponentTypeAttr = "otlp_http_json_metric_exporter"
)

type SpanSamplingResultAttr string

var (
	SpanSamplingResultDrop SpanSamplingResultAttr = "DROP"

	SpanSamplingResultRecordOnly SpanSamplingResultAttr = "RECORD_ONLY"

	SpanSamplingResultRecordAndSample SpanSamplingResultAttr = "RECORD_AND_SAMPLE"
)

type RPCGRPCStatusCodeAttr int64

var (
	RPCGRPCStatusCodeOk RPCGRPCStatusCodeAttr = 0

	RPCGRPCStatusCodeCancelled RPCGRPCStatusCodeAttr = 1

	RPCGRPCStatusCodeUnknown RPCGRPCStatusCodeAttr = 2

	RPCGRPCStatusCodeInvalidArgument RPCGRPCStatusCodeAttr = 3

	RPCGRPCStatusCodeDeadlineExceeded RPCGRPCStatusCodeAttr = 4

	RPCGRPCStatusCodeNotFound RPCGRPCStatusCodeAttr = 5

	RPCGRPCStatusCodeAlreadyExists RPCGRPCStatusCodeAttr = 6

	RPCGRPCStatusCodePermissionDenied RPCGRPCStatusCodeAttr = 7

	RPCGRPCStatusCodeResourceExhausted RPCGRPCStatusCodeAttr = 8

	RPCGRPCStatusCodeFailedPrecondition RPCGRPCStatusCodeAttr = 9

	RPCGRPCStatusCodeAborted RPCGRPCStatusCodeAttr = 10

	RPCGRPCStatusCodeOutOfRange RPCGRPCStatusCodeAttr = 11

	RPCGRPCStatusCodeUnimplemented RPCGRPCStatusCodeAttr = 12

	RPCGRPCStatusCodeInternal RPCGRPCStatusCodeAttr = 13

	RPCGRPCStatusCodeUnavailable RPCGRPCStatusCodeAttr = 14

	RPCGRPCStatusCodeDataLoss RPCGRPCStatusCodeAttr = 15

	RPCGRPCStatusCodeUnauthenticated RPCGRPCStatusCodeAttr = 16
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

type SDKExporterMetricDataPointExported struct {
	metric.Int64Counter
}

func NewSDKExporterMetricDataPointExported(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKExporterMetricDataPointExported, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterMetricDataPointExported), nil
}

func (m SDKExporterMetricDataPointExported) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKExporterMetricDataPointExported) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterMetricDataPointExported) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterMetricDataPointExported) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m SDKExporterMetricDataPointExported) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKExporterMetricDataPointExported) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointExported) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointExported) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointExported) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointExported) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterMetricDataPointInflight struct {
	metric.Int64UpDownCounter
}

func NewSDKExporterMetricDataPointInflight(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKExporterMetricDataPointInflight, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterMetricDataPointInflight), nil
}

func (m SDKExporterMetricDataPointInflight) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKExporterMetricDataPointInflight) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterMetricDataPointInflight) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterMetricDataPointInflight) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m SDKExporterMetricDataPointInflight) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKExporterMetricDataPointInflight) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointInflight) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointInflight) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointInflight) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterOperationDuration struct {
	metric.Float64Histogram
}

func NewSDKExporterOperationDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (SDKExporterOperationDuration, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterOperationDuration), nil
}

func (m SDKExporterOperationDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (SDKExporterOperationDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterOperationDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterOperationDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKExporterOperationDuration) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKExporterOperationDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterOperationDuration) AttrHTTPResponseStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterOperationDuration) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterOperationDuration) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterOperationDuration) AttrRPCGRPCStatusCode(val RPCGRPCStatusCodeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterOperationDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterOperationDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterSpanExported struct {
	metric.Int64Counter
}

func NewSDKExporterSpanExported(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKExporterSpanExported, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterSpanExported), nil
}

func (m SDKExporterSpanExported) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKExporterSpanExported) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanExported) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanExported) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKExporterSpanExported) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKExporterSpanExported) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExported) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExported) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExported) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExported) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterSpanInflight struct {
	metric.Int64UpDownCounter
}

func NewSDKExporterSpanInflight(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKExporterSpanInflight, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterSpanInflight), nil
}

func (m SDKExporterSpanInflight) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKExporterSpanInflight) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanInflight) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanInflight) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKExporterSpanInflight) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKExporterSpanInflight) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanInflight) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanInflight) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanInflight) AttrServerPort(val int) attribute.KeyValue {
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

type SDKMetricReaderCollectionDuration struct {
	metric.Float64Histogram
}

func NewSDKMetricReaderCollectionDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (SDKMetricReaderCollectionDuration, error) {
	_ = "STUB: not implemented"
	return *new(SDKMetricReaderCollectionDuration), nil
}

func (m SDKMetricReaderCollectionDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (SDKMetricReaderCollectionDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKMetricReaderCollectionDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKMetricReaderCollectionDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKMetricReaderCollectionDuration) Record(
	ctx context.Context,
	val float64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKMetricReaderCollectionDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKMetricReaderCollectionDuration) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKMetricReaderCollectionDuration) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type SDKProcessorSpanProcessed struct {
	metric.Int64Counter
}

func NewSDKProcessorSpanProcessed(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKProcessorSpanProcessed, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorSpanProcessed), nil
}

func (m SDKProcessorSpanProcessed) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKProcessorSpanProcessed) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanProcessed) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanProcessed) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKProcessorSpanProcessed) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKProcessorSpanProcessed) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanProcessed) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanProcessed) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
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

type SDKSpanEnded struct {
	metric.Int64Counter
}

func NewSDKSpanEnded(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKSpanEnded, error) {
	_ = "STUB: not implemented"
	return *new(SDKSpanEnded), nil
}

func (m SDKSpanEnded) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKSpanEnded) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanEnded) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanEnded) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKSpanEnded) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKSpanEnded) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKSpanLive struct {
	metric.Int64UpDownCounter
}

func NewSDKSpanLive(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (SDKSpanLive, error) {
	_ = "STUB: not implemented"
	return *new(SDKSpanLive), nil
}

func (m SDKSpanLive) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (SDKSpanLive) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanLive) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanLive) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKSpanLive) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (SDKSpanLive) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
