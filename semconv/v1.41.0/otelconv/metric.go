package otelconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type ErrorTypeAttr string

var (
	ErrorTypeOther ErrorTypeAttr = "_OTHER"
)

type ComponentTypeAttr string

var (
	ComponentTypeBatchingSpanProcessor ComponentTypeAttr = "batching_span_processor"

	ComponentTypeSimpleSpanProcessor ComponentTypeAttr = "simple_span_processor"

	ComponentTypeBatchingLogProcessor ComponentTypeAttr = "batching_log_processor"

	ComponentTypeSimpleLogProcessor ComponentTypeAttr = "simple_log_processor"

	ComponentTypeOtlpGRPCSpanExporter ComponentTypeAttr = "otlp_grpc_span_exporter"

	ComponentTypeOtlpHTTPSpanExporter ComponentTypeAttr = "otlp_http_span_exporter"

	ComponentTypeOtlpHTTPJSONSpanExporter ComponentTypeAttr = "otlp_http_json_span_exporter"

	ComponentTypeZipkinHTTPSpanExporter ComponentTypeAttr = "zipkin_http_span_exporter"

	ComponentTypeOtlpGRPCLogExporter ComponentTypeAttr = "otlp_grpc_log_exporter"

	ComponentTypeOtlpHTTPLogExporter ComponentTypeAttr = "otlp_http_log_exporter"

	ComponentTypeOtlpHTTPJSONLogExporter ComponentTypeAttr = "otlp_http_json_log_exporter"

	ComponentTypePeriodicMetricReader ComponentTypeAttr = "periodic_metric_reader"

	ComponentTypeOtlpGRPCMetricExporter ComponentTypeAttr = "otlp_grpc_metric_exporter"

	ComponentTypeOtlpHTTPMetricExporter ComponentTypeAttr = "otlp_http_metric_exporter"

	ComponentTypeOtlpHTTPJSONMetricExporter ComponentTypeAttr = "otlp_http_json_metric_exporter"

	ComponentTypePrometheusHTTPTextMetricExporter ComponentTypeAttr = "prometheus_http_text_metric_exporter"
)

type SpanParentOriginAttr string

var (
	SpanParentOriginNone SpanParentOriginAttr = "none"

	SpanParentOriginLocal SpanParentOriginAttr = "local"

	SpanParentOriginRemote SpanParentOriginAttr = "remote"
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

var newSDKExporterLogExportedOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of log records for which the export has finished, either successful or failed."),
	metric.WithUnit("{log_record}"),
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

func (m SDKExporterLogExported) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type SDKExporterLogExportedObservable struct {
	metric.Int64ObservableCounter
}

var newSDKExporterLogExportedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of log records for which the export has finished, either successful or failed."),
	metric.WithUnit("{log_record}"),
}

func NewSDKExporterLogExportedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (SDKExporterLogExportedObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterLogExportedObservable), nil
}

func (m SDKExporterLogExportedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (SDKExporterLogExportedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogExportedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogExportedObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogExportedObservable) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogExportedObservable) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogExportedObservable) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogExportedObservable) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogExportedObservable) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterLogInflight struct {
	metric.Int64UpDownCounter
}

var newSDKExporterLogInflightOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of log records which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)."),
	metric.WithUnit("{log_record}"),
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

func (m SDKExporterLogInflight) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type SDKExporterLogInflightObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newSDKExporterLogInflightObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of log records which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)."),
	metric.WithUnit("{log_record}"),
}

func NewSDKExporterLogInflightObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (SDKExporterLogInflightObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterLogInflightObservable), nil
}

func (m SDKExporterLogInflightObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (SDKExporterLogInflightObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogInflightObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogInflightObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterLogInflightObservable) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogInflightObservable) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogInflightObservable) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterLogInflightObservable) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterMetricDataPointExported struct {
	metric.Int64Counter
}

var newSDKExporterMetricDataPointExportedOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of metric data points for which the export has finished, either successful or failed."),
	metric.WithUnit("{data_point}"),
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

func (m SDKExporterMetricDataPointExported) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type SDKExporterMetricDataPointExportedObservable struct {
	metric.Int64ObservableCounter
}

var newSDKExporterMetricDataPointExportedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of metric data points for which the export has finished, either successful or failed."),
	metric.WithUnit("{data_point}"),
}

func NewSDKExporterMetricDataPointExportedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (SDKExporterMetricDataPointExportedObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterMetricDataPointExportedObservable), nil
}

func (m SDKExporterMetricDataPointExportedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (SDKExporterMetricDataPointExportedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (SDKExporterMetricDataPointExportedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (SDKExporterMetricDataPointExportedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (SDKExporterMetricDataPointExportedObservable) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointExportedObservable) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointExportedObservable) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointExportedObservable) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointExportedObservable) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterMetricDataPointInflight struct {
	metric.Int64UpDownCounter
}

var newSDKExporterMetricDataPointInflightOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of metric data points which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)."),
	metric.WithUnit("{data_point}"),
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

func (m SDKExporterMetricDataPointInflight) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type SDKExporterMetricDataPointInflightObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newSDKExporterMetricDataPointInflightObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of metric data points which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)."),
	metric.WithUnit("{data_point}"),
}

func NewSDKExporterMetricDataPointInflightObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (SDKExporterMetricDataPointInflightObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterMetricDataPointInflightObservable), nil
}

func (m SDKExporterMetricDataPointInflightObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (SDKExporterMetricDataPointInflightObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (SDKExporterMetricDataPointInflightObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (SDKExporterMetricDataPointInflightObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (SDKExporterMetricDataPointInflightObservable) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointInflightObservable) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointInflightObservable) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterMetricDataPointInflightObservable) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterOperationDuration struct {
	metric.Float64Histogram
}

var newSDKExporterOperationDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The duration of exporting a batch of telemetry records."),
	metric.WithUnit("s"),
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

func (m SDKExporterOperationDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
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

func (SDKExporterOperationDuration) AttrRPCResponseStatusCode(val string) attribute.KeyValue {
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

var newSDKExporterSpanExportedOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of spans for which the export has finished, either successful or failed."),
	metric.WithUnit("{span}"),
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

func (m SDKExporterSpanExported) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type SDKExporterSpanExportedObservable struct {
	metric.Int64ObservableCounter
}

var newSDKExporterSpanExportedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of spans for which the export has finished, either successful or failed."),
	metric.WithUnit("{span}"),
}

func NewSDKExporterSpanExportedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (SDKExporterSpanExportedObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterSpanExportedObservable), nil
}

func (m SDKExporterSpanExportedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (SDKExporterSpanExportedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanExportedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanExportedObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanExportedObservable) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExportedObservable) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExportedObservable) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExportedObservable) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanExportedObservable) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKExporterSpanInflight struct {
	metric.Int64UpDownCounter
}

var newSDKExporterSpanInflightOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of spans which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)."),
	metric.WithUnit("{span}"),
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

func (m SDKExporterSpanInflight) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type SDKExporterSpanInflightObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newSDKExporterSpanInflightObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of spans which were passed to the exporter, but that have not been exported yet (neither successful, nor failed)."),
	metric.WithUnit("{span}"),
}

func NewSDKExporterSpanInflightObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (SDKExporterSpanInflightObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKExporterSpanInflightObservable), nil
}

func (m SDKExporterSpanInflightObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (SDKExporterSpanInflightObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanInflightObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanInflightObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (SDKExporterSpanInflightObservable) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanInflightObservable) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanInflightObservable) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKExporterSpanInflightObservable) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKLogCreated struct {
	metric.Int64Counter
}

var newSDKLogCreatedOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of logs submitted to enabled SDK Loggers."),
	metric.WithUnit("{log_record}"),
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

func (m SDKLogCreated) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type SDKLogCreatedObservable struct {
	metric.Int64ObservableCounter
}

var newSDKLogCreatedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of logs submitted to enabled SDK Loggers."),
	metric.WithUnit("{log_record}"),
}

func NewSDKLogCreatedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (SDKLogCreatedObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKLogCreatedObservable), nil
}

func (m SDKLogCreatedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (SDKLogCreatedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKLogCreatedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKLogCreatedObservable) Description() string { _ = "STUB: not implemented"; return "" }

type SDKMetricReaderCollectionDuration struct {
	metric.Float64Histogram
}

var newSDKMetricReaderCollectionDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The duration of the collect operation of the metric reader."),
	metric.WithUnit("s"),
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

func (m SDKMetricReaderCollectionDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
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

var newSDKProcessorLogProcessedOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of log records for which the processing has finished, either successful or failed."),
	metric.WithUnit("{log_record}"),
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

func (m SDKProcessorLogProcessed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type SDKProcessorLogProcessedObservable struct {
	metric.Int64ObservableCounter
}

var newSDKProcessorLogProcessedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of log records for which the processing has finished, either successful or failed."),
	metric.WithUnit("{log_record}"),
}

func NewSDKProcessorLogProcessedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (SDKProcessorLogProcessedObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorLogProcessedObservable), nil
}

func (m SDKProcessorLogProcessedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (SDKProcessorLogProcessedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogProcessedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogProcessedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (SDKProcessorLogProcessedObservable) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorLogProcessedObservable) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorLogProcessedObservable) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKProcessorLogQueueCapacity struct {
	metric.Int64ObservableUpDownCounter
}

var newSDKProcessorLogQueueCapacityOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The maximum number of log records the queue of a given instance of an SDK Log Record processor can hold."),
	metric.WithUnit("{log_record}"),
}

func NewSDKProcessorLogQueueCapacity(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (SDKProcessorLogQueueCapacity, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorLogQueueCapacity), nil
}

func (m SDKProcessorLogQueueCapacity) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (SDKProcessorLogQueueCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogQueueCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogQueueCapacity) Description() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogQueueCapacity) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorLogQueueCapacity) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKProcessorLogQueueSize struct {
	metric.Int64ObservableUpDownCounter
}

var newSDKProcessorLogQueueSizeOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of log records in the queue of a given instance of an SDK log processor."),
	metric.WithUnit("{log_record}"),
}

func NewSDKProcessorLogQueueSize(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (SDKProcessorLogQueueSize, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorLogQueueSize), nil
}

func (m SDKProcessorLogQueueSize) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (SDKProcessorLogQueueSize) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogQueueSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorLogQueueSize) Description() string { _ = "STUB: not implemented"; return "" }

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

var newSDKProcessorSpanProcessedOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of spans for which the processing has finished, either successful or failed."),
	metric.WithUnit("{span}"),
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

func (m SDKProcessorSpanProcessed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type SDKProcessorSpanProcessedObservable struct {
	metric.Int64ObservableCounter
}

var newSDKProcessorSpanProcessedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of spans for which the processing has finished, either successful or failed."),
	metric.WithUnit("{span}"),
}

func NewSDKProcessorSpanProcessedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (SDKProcessorSpanProcessedObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorSpanProcessedObservable), nil
}

func (m SDKProcessorSpanProcessedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (SDKProcessorSpanProcessedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanProcessedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanProcessedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (SDKProcessorSpanProcessedObservable) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanProcessedObservable) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanProcessedObservable) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKProcessorSpanQueueCapacity struct {
	metric.Int64ObservableUpDownCounter
}

var newSDKProcessorSpanQueueCapacityOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The maximum number of spans the queue of a given instance of an SDK span processor can hold."),
	metric.WithUnit("{span}"),
}

func NewSDKProcessorSpanQueueCapacity(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (SDKProcessorSpanQueueCapacity, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorSpanQueueCapacity), nil
}

func (m SDKProcessorSpanQueueCapacity) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (SDKProcessorSpanQueueCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueCapacity) Description() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueCapacity) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanQueueCapacity) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKProcessorSpanQueueSize struct {
	metric.Int64ObservableUpDownCounter
}

var newSDKProcessorSpanQueueSizeOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of spans in the queue of a given instance of an SDK span processor."),
	metric.WithUnit("{span}"),
}

func NewSDKProcessorSpanQueueSize(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (SDKProcessorSpanQueueSize, error) {
	_ = "STUB: not implemented"
	return *new(SDKProcessorSpanQueueSize), nil
}

func (m SDKProcessorSpanQueueSize) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (SDKProcessorSpanQueueSize) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueSize) Description() string { _ = "STUB: not implemented"; return "" }

func (SDKProcessorSpanQueueSize) AttrComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKProcessorSpanQueueSize) AttrComponentType(val ComponentTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKSpanLive struct {
	metric.Int64UpDownCounter
}

var newSDKSpanLiveOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of created spans with `recording=true` for which the end operation has not been called yet."),
	metric.WithUnit("{span}"),
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

func (m SDKSpanLive) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (SDKSpanLive) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKSpanLiveObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newSDKSpanLiveObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of created spans with `recording=true` for which the end operation has not been called yet."),
	metric.WithUnit("{span}"),
}

func NewSDKSpanLiveObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (SDKSpanLiveObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKSpanLiveObservable), nil
}

func (m SDKSpanLiveObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (SDKSpanLiveObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanLiveObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanLiveObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanLiveObservable) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKSpanStarted struct {
	metric.Int64Counter
}

var newSDKSpanStartedOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of created spans."),
	metric.WithUnit("{span}"),
}

func NewSDKSpanStarted(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SDKSpanStarted, error) {
	_ = "STUB: not implemented"
	return *new(SDKSpanStarted), nil
}

func (m SDKSpanStarted) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SDKSpanStarted) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanStarted) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanStarted) Description() string { _ = "STUB: not implemented"; return "" }

func (m SDKSpanStarted) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m SDKSpanStarted) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (SDKSpanStarted) AttrSpanParentOrigin(val SpanParentOriginAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKSpanStarted) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type SDKSpanStartedObservable struct {
	metric.Int64ObservableCounter
}

var newSDKSpanStartedObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of created spans."),
	metric.WithUnit("{span}"),
}

func NewSDKSpanStartedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (SDKSpanStartedObservable, error) {
	_ = "STUB: not implemented"
	return *new(SDKSpanStartedObservable), nil
}

func (m SDKSpanStartedObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (SDKSpanStartedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanStartedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanStartedObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (SDKSpanStartedObservable) AttrSpanParentOrigin(val SpanParentOriginAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (SDKSpanStartedObservable) AttrSpanSamplingResult(val SpanSamplingResultAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
