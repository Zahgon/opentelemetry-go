package opentracing

import (
	"context"
	"sync"

	ot "github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	iBaggage "go.opentelemetry.io/otel/internal/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var (
	noopTracer = noop.NewTracerProvider().Tracer("")
	noopSpan   = func() trace.Span {
		_, s := noopTracer.Start(context.Background(), "")
		return s
	}()
)

type bridgeSpanContext struct {
	bag baggage.Baggage
	trace.SpanContext
}

var _ ot.SpanContext = &bridgeSpanContext{}

func newBridgeSpanContext(otelSpanContext trace.SpanContext, parentOtSpanContext ot.SpanContext) *bridgeSpanContext {
	_ = "STUB: not implemented"
	return nil
}

func (c *bridgeSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {
	_ = "STUB: not implemented"
	return
}

func (c *bridgeSpanContext) setBaggageItem(restrictedKey, value string) {
	_ = "STUB: not implemented"
	return
}

func (c *bridgeSpanContext) baggageItem(restrictedKey string) baggage.Member {
	_ = "STUB: not implemented"
	return *new(baggage.Member)
}

type bridgeSpan struct {
	otelSpan          trace.Span
	ctx               *bridgeSpanContext
	tracer            *BridgeTracer
	skipDeferHook     bool
	extraBaggageItems map[string]string
}

var _ ot.Span = &bridgeSpan{}

func newBridgeSpan(otelSpan trace.Span, bridgeSC *bridgeSpanContext, tracer *BridgeTracer) *bridgeSpan {
	_ = "STUB: not implemented"
	return nil
}

func (s *bridgeSpan) Finish() { _ = "STUB: not implemented"; return }

func (s *bridgeSpan) FinishWithOptions(opts ot.FinishOptions) { _ = "STUB: not implemented"; return }

func (s *bridgeSpan) logRecord(record ot.LogRecord) { _ = "STUB: not implemented"; return }

func (s *bridgeSpan) Context() ot.SpanContext {
	_ = "STUB: not implemented"
	return *new(ot.SpanContext)
}

func (s *bridgeSpan) SetOperationName(operationName string) ot.Span {
	_ = "STUB: not implemented"
	return *new(ot.Span)
}

func (s *bridgeSpan) SetTag(key string, value any) ot.Span {
	_ = "STUB: not implemented"
	return *new(ot.Span)
}

func (s *bridgeSpan) LogFields(fields ...otlog.Field) { _ = "STUB: not implemented"; return }

type bridgeFieldEncoder struct {
	pairs []attribute.KeyValue
}

var _ otlog.Encoder = &bridgeFieldEncoder{}

func (e *bridgeFieldEncoder) EmitString(key, value string) { _ = "STUB: not implemented"; return }

func (e *bridgeFieldEncoder) EmitBool(key string, value bool) { _ = "STUB: not implemented"; return }

func (e *bridgeFieldEncoder) EmitInt(key string, value int) { _ = "STUB: not implemented"; return }

func (e *bridgeFieldEncoder) EmitInt32(key string, value int32) { _ = "STUB: not implemented"; return }

func (e *bridgeFieldEncoder) EmitInt64(key string, value int64) { _ = "STUB: not implemented"; return }

func (e *bridgeFieldEncoder) EmitUint32(key string, value uint32) {
	_ = "STUB: not implemented"
	return
}

func (e *bridgeFieldEncoder) EmitUint64(key string, value uint64) {
	_ = "STUB: not implemented"
	return
}

func (e *bridgeFieldEncoder) EmitFloat32(key string, value float32) {
	_ = "STUB: not implemented"
	return
}

func (e *bridgeFieldEncoder) EmitFloat64(key string, value float64) {
	_ = "STUB: not implemented"
	return
}

func (e *bridgeFieldEncoder) EmitObject(key string, value any) { _ = "STUB: not implemented"; return }

func (e *bridgeFieldEncoder) EmitLazyLogger(value otlog.LazyLogger) {
	_ = "STUB: not implemented"
	return
}

func (e *bridgeFieldEncoder) emitCommon(key string, value any) { _ = "STUB: not implemented"; return }

func otLogFieldsToOTelAttrs(fields []otlog.Field) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *bridgeSpan) LogKV(alternatingKeyValues ...any) { _ = "STUB: not implemented"; return }

func (s *bridgeSpan) SetBaggageItem(restrictedKey, value string) ot.Span {
	_ = "STUB: not implemented"
	return *new(ot.Span)
}

func (s *bridgeSpan) setBaggageItemOnly(restrictedKey, value string) {
	_ = "STUB: not implemented"
	return
}

func (s *bridgeSpan) updateOTelContext(restrictedKey, value string) {
	_ = "STUB: not implemented"
	return
}

func (s *bridgeSpan) BaggageItem(restrictedKey string) string { _ = "STUB: not implemented"; return "" }

func (s *bridgeSpan) Tracer() ot.Tracer { _ = "STUB: not implemented"; return *new(ot.Tracer) }

func (s *bridgeSpan) LogEvent(event string) { _ = "STUB: not implemented"; return }

func (s *bridgeSpan) LogEventWithPayload(event string, payload any) {
	_ = "STUB: not implemented"
	return
}

func (s *bridgeSpan) Log(data ot.LogData) { _ = "STUB: not implemented"; return }

type bridgeSetTracer struct {
	isSet      bool
	otelTracer trace.Tracer

	warningHandler BridgeWarningHandler
	warnOnce       sync.Once
}

func (s *bridgeSetTracer) tracer() trace.Tracer {
	_ = "STUB: not implemented"
	return *new(trace.Tracer)
}

type BridgeWarningHandler func(msg string)

type BridgeTracer struct {
	setTracer bridgeSetTracer

	warningHandler BridgeWarningHandler
	warnOnce       sync.Once

	propagator propagation.TextMapPropagator
}

var (
	_ ot.Tracer                         = &BridgeTracer{}
	_ ot.TracerContextWithSpanExtension = &BridgeTracer{}
)

func NewBridgeTracer() *BridgeTracer { _ = "STUB: not implemented"; return nil }

func (t *BridgeTracer) SetWarningHandler(handler BridgeWarningHandler) {
	_ = "STUB: not implemented"
	return
}

func (t *BridgeTracer) SetOpenTelemetryTracer(tracer trace.Tracer) {
	_ = "STUB: not implemented"
	return
}

func (t *BridgeTracer) SetTextMapPropagator(propagator propagation.TextMapPropagator) {
	_ = "STUB: not implemented"
	return
}

func (t *BridgeTracer) NewHookedContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (t *BridgeTracer) baggageSetHook(ctx context.Context, list iBaggage.List) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (t *BridgeTracer) baggageGetHook(ctx context.Context, list iBaggage.List) iBaggage.List {
	_ = "STUB: not implemented"
	return *new(iBaggage.List)
}

func (t *BridgeTracer) StartSpan(operationName string, opts ...ot.StartSpanOption) ot.Span {
	_ = "STUB: not implemented"
	return *new(ot.Span)
}

func (t *BridgeTracer) ContextWithBridgeSpan(ctx context.Context, span trace.Span) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (t *BridgeTracer) ContextWithSpanHook(ctx context.Context, span ot.Span) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func otTagsToOTelAttributesKindAndError(tags map[string]any) ([]attribute.KeyValue, trace.SpanKind, bool) {
	_ = "STUB: not implemented"
	return nil, *new(trace.SpanKind), false
}

func otTagToOTelAttr(k string, v any) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func otTagToOTelAttrKey(k string) attribute.Key {
	_ = "STUB: not implemented"
	return *new(attribute.Key)
}

func otSpanReferencesToParentAndLinks(references []ot.SpanReference) (*bridgeSpanContext, []trace.Link) {
	_ = "STUB: not implemented"
	return nil, nil
}

func otSpanReferenceToOTelLink(bridgeSC *bridgeSpanContext, refType ot.SpanReferenceType) trace.Link {
	_ = "STUB: not implemented"
	return *new(trace.Link)
}

func otSpanReferenceTypeToOTelLinkAttributes(refType ot.SpanReferenceType) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func otSpanReferenceTypeToString(refType ot.SpanReferenceType) string {
	_ = "STUB: not implemented"
	return ""
}

type fakeSpan struct {
	trace.Span
	sc trace.SpanContext
}

func (s fakeSpan) SpanContext() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (t *BridgeTracer) Inject(sm ot.SpanContext, format, carrier any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *BridgeTracer) Extract(format, carrier any) (ot.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(ot.SpanContext), nil
}

func (t *BridgeTracer) getPropagator() propagation.TextMapPropagator {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator)
}

type textMapWrapper struct {
	ot.TextMapWriter
	ot.TextMapReader
	readerMap map[string]string
}

func (t *textMapWrapper) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (t *textMapWrapper) Set(key, value string) { _ = "STUB: not implemented"; return }

func (t *textMapWrapper) Keys() []string { _ = "STUB: not implemented"; return nil }

func (t *textMapWrapper) loadMap() { _ = "STUB: not implemented"; return }

func newTextMapWrapperForExtract(carrier any) (*textMapWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTextMapWrapperForInject(carrier any) (*textMapWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type textMapWriter struct{}

func (*textMapWriter) Set(string, string) { _ = "STUB: not implemented"; return }

type textMapReader struct{}

func (*textMapReader) ForeachKey(func(string, string) error) error {
	_ = "STUB: not implemented"
	return nil
}
