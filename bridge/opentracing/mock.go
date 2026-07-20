package opentracing

import (
	"context"
	"math/rand/v2"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/bridge/opentracing/migration"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

//nolint:revive // ignoring missing comments for unexported global variables in an internal package.
var (
	statusCodeKey    = attribute.Key("status.code")
	statusMessageKey = attribute.Key("status.message")
	errorKey         = attribute.Key("error")
	nameKey          = attribute.Key("name")
)

type mockContextKeyValue struct {
	Key   any
	Value any
}

type mockTracer struct {
	embedded.Tracer

	FinishedSpans         []*mockSpan
	SpareTraceIDs         []trace.TraceID
	SpareSpanIDs          []trace.SpanID
	SpareContextKeyValues []mockContextKeyValue
	TraceFlags            trace.TraceFlags

	randLock sync.Mutex
	rand     *rand.ChaCha8
}

var (
	_ trace.Tracer                                  = &mockTracer{}
	_ migration.DeferredContextSetupTracerExtension = &mockTracer{}
)

func newMockTracer() *mockTracer { _ = "STUB: not implemented"; return nil }

func (t *mockTracer) Start(
	ctx context.Context,
	_ string,
	opts ...trace.SpanStartOption,
) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

func (t *mockTracer) addSpareContextValue(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (t *mockTracer) getTraceID(ctx context.Context, config *trace.SpanConfig) trace.TraceID {
	_ = "STUB: not implemented"
	return *new(trace.TraceID)
}

func (t *mockTracer) getParentSpanID(ctx context.Context, config *trace.SpanConfig) trace.SpanID {
	_ = "STUB: not implemented"
	return *new(trace.SpanID)
}

func (*mockTracer) getParentSpanContext(ctx context.Context, config *trace.SpanConfig) trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (t *mockTracer) getSpanID() trace.SpanID { _ = "STUB: not implemented"; return *new(trace.SpanID) }

func (t *mockTracer) getRandSpanID() trace.SpanID {
	_ = "STUB: not implemented"
	return *new(trace.SpanID)
}

func (t *mockTracer) getRandTraceID() trace.TraceID {
	_ = "STUB: not implemented"
	return *new(trace.TraceID)
}

func (t *mockTracer) DeferredContextSetupHook(ctx context.Context, _ trace.Span) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type mockEvent struct {
	Timestamp  time.Time
	Name       string
	Attributes []attribute.KeyValue
}

type mockLink struct {
	SpanContext trace.SpanContext
	Attributes  []attribute.KeyValue
}

type mockSpan struct {
	embedded.Span

	mockTracer     *mockTracer
	officialTracer trace.Tracer
	spanContext    trace.SpanContext
	SpanKind       trace.SpanKind
	recording      bool

	Attributes   []attribute.KeyValue
	StartTime    time.Time
	EndTime      time.Time
	ParentSpanID trace.SpanID
	Events       []mockEvent
	Links        []mockLink
}

var (
	_ trace.Span                            = &mockSpan{}
	_ migration.OverrideTracerSpanExtension = &mockSpan{}
)

func (s *mockSpan) SpanContext() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (s *mockSpan) IsRecording() bool { _ = "STUB: not implemented"; return false }

func (s *mockSpan) SetStatus(code codes.Code, msg string) { _ = "STUB: not implemented"; return }

func (s *mockSpan) SetName(name string) { _ = "STUB: not implemented"; return }

func (s *mockSpan) SetError(v bool) { _ = "STUB: not implemented"; return }

func (s *mockSpan) SetAttributes(attributes ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (s *mockSpan) applyUpdate(update []attribute.KeyValue) { _ = "STUB: not implemented"; return }

func (s *mockSpan) End(options ...trace.SpanEndOption) { _ = "STUB: not implemented"; return }

func (s *mockSpan) RecordError(err error, opts ...trace.EventOption) {
	_ = "STUB: not implemented"
	return
}

func (s *mockSpan) Tracer() trace.Tracer { _ = "STUB: not implemented"; return *new(trace.Tracer) }

func (s *mockSpan) AddEvent(name string, o ...trace.EventOption) { _ = "STUB: not implemented"; return }

func (s *mockSpan) AddLink(link trace.Link) { _ = "STUB: not implemented"; return }

func (s *mockSpan) OverrideTracer(tracer trace.Tracer) { _ = "STUB: not implemented"; return }

func (*mockSpan) TracerProvider() trace.TracerProvider {
	_ = "STUB: not implemented"
	return *new(trace.TracerProvider)
}
