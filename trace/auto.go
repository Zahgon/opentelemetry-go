package trace

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace/embedded"
	"go.opentelemetry.io/otel/trace/internal/telemetry"
)

func newAutoTracerProvider() TracerProvider { _ = "STUB: not implemented"; return *new(TracerProvider) }

var tracerProviderInstance = new(autoTracerProvider)

type autoTracerProvider struct{ embedded.TracerProvider }

var _ TracerProvider = autoTracerProvider{}

func (autoTracerProvider) Tracer(name string, opts ...TracerOption) Tracer {
	_ = "STUB: not implemented"
	return *new(Tracer)
}

type autoTracer struct {
	embedded.Tracer

	name, schemaURL, version string
}

var _ Tracer = autoTracer{}

func (t autoTracer) Start(ctx context.Context, name string, opts ...SpanStartOption) (context.Context, Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(Span)
}

//go:noinline
func (*autoTracer) start(
	ctx context.Context,
	spanPtr *autoSpan,
	psc *SpanContext,
	sampled *bool,
	sc *SpanContext,
) {
	_ = "STUB: not implemented"
	return
}

var start = func(context.Context, *autoSpan, *SpanContext, *bool, *SpanContext) {}

func (t autoTracer) traces(name string, cfg SpanConfig, sc, psc SpanContext) (*telemetry.Traces, *telemetry.Span) {
	_ = "STUB: not implemented"
	return nil, nil
}

func spanKind(kind SpanKind) telemetry.SpanKind {
	_ = "STUB: not implemented"
	return *new(telemetry.SpanKind)
}

type autoSpan struct {
	embedded.Span

	spanContext SpanContext
	sampled     atomic.Bool

	mu     sync.Mutex
	traces *telemetry.Traces
	span   *telemetry.Span
}

func (s *autoSpan) SpanContext() SpanContext { _ = "STUB: not implemented"; return *new(SpanContext) }

func (s *autoSpan) IsRecording() bool { _ = "STUB: not implemented"; return false }

func (s *autoSpan) SetStatus(c codes.Code, msg string) { _ = "STUB: not implemented"; return }

func (s *autoSpan) SetAttributes(attrs ...attribute.KeyValue) { _ = "STUB: not implemented"; return }

func convCappedAttrs(limit int, attrs []attribute.KeyValue) ([]telemetry.Attr, uint32) {
	_ = "STUB: not implemented"
	return nil, 0
}

func convAttrs(attrs []attribute.KeyValue) []telemetry.Attr { _ = "STUB: not implemented"; return nil }

func convAttrValue(value attribute.Value) telemetry.Value {
	_ = "STUB: not implemented"
	return *new(telemetry.Value)
}

func truncate(limit int, s string) string { _ = "STUB: not implemented"; return "" }

func (s *autoSpan) End(opts ...SpanEndOption) { _ = "STUB: not implemented"; return }

func (s *autoSpan) end(opts []SpanEndOption) []byte { _ = "STUB: not implemented"; return nil }

//go:noinline
func (*autoSpan) ended(buf []byte) { _ = "STUB: not implemented"; return }

var ended = func([]byte) {}

func (s *autoSpan) RecordError(err error, opts ...EventOption) { _ = "STUB: not implemented"; return }

func typeStr(i any) string { _ = "STUB: not implemented"; return "" }

func (s *autoSpan) AddEvent(name string, opts ...EventOption) { _ = "STUB: not implemented"; return }

func (s *autoSpan) addEvent(name string, tStamp time.Time, attrs []attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (s *autoSpan) AddLink(link Link) { _ = "STUB: not implemented"; return }

func convLinks(links []Link) []*telemetry.SpanLink { _ = "STUB: not implemented"; return nil }

func convLink(link Link) *telemetry.SpanLink { _ = "STUB: not implemented"; return nil }

func (s *autoSpan) SetName(name string) { _ = "STUB: not implemented"; return }

func (*autoSpan) TracerProvider() TracerProvider {
	_ = "STUB: not implemented"
	return *new(TracerProvider)
}

var maxSpan = newSpanLimits()

type spanLimits struct {
	Attrs int

	AttrValueLen int

	Events int

	EventAttrs int

	Links int

	LinkAttrs int
}

func newSpanLimits() spanLimits { _ = "STUB: not implemented"; return *new(spanLimits) }

func firstEnv(defaultVal int, keys ...string) int { _ = "STUB: not implemented"; return 0 }
