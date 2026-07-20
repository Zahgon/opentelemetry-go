package trace

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/internal/global"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

type ReadOnlySpan interface {
	Name() string

	SpanContext() trace.SpanContext

	Parent() trace.SpanContext

	SpanKind() trace.SpanKind

	StartTime() time.Time

	EndTime() time.Time

	Attributes() []attribute.KeyValue

	Links() []Link

	Events() []Event

	Status() Status

	InstrumentationScope() instrumentation.Scope

	InstrumentationLibrary() instrumentation.Library //nolint:staticcheck // This method needs to be define for backwards compatibility

	Resource() *resource.Resource

	DroppedAttributes() int

	DroppedLinks() int

	DroppedEvents() int

	ChildSpanCount() int

	private()
}

type ReadWriteSpan interface {
	trace.Span
	ReadOnlySpan
}

type recordingSpan struct {
	embedded.Span

	mu sync.Mutex

	parent trace.SpanContext

	spanKind trace.SpanKind

	name string

	startTime time.Time

	endTime time.Time

	status Status

	childSpanCount int

	spanContext trace.SpanContext

	attributes        []attribute.KeyValue
	droppedAttributes int
	logDropAttrsOnce  sync.Once

	events evictedQueue[Event]

	links evictedQueue[Link]

	executionTracerTaskEnd func()

	tracer *tracer

	origCtx context.Context
}

var (
	_ ReadWriteSpan = (*recordingSpan)(nil)
	_ runtimeTracer = (*recordingSpan)(nil)
)

func (s *recordingSpan) setOrigCtx(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *recordingSpan) SpanContext() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (s *recordingSpan) IsRecording() bool { _ = "STUB: not implemented"; return false }

func (s *recordingSpan) isRecording() bool { _ = "STUB: not implemented"; return false }

func (s *recordingSpan) SetStatus(code codes.Code, description string) {
	_ = "STUB: not implemented"
	return
}

func (s *recordingSpan) SetAttributes(attributes ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

var logDropAttrs = func() {
	global.Warn("limit reached: dropping trace Span attributes")
}

func (s *recordingSpan) addDroppedAttr(incr int) { _ = "STUB: not implemented"; return }

func (s *recordingSpan) addOverCapAttrs(limit int, attrs []attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func dedupAttr(attr attribute.KeyValue) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func truncateAttr(limit int, attr attribute.KeyValue) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func truncateValue(limit int, v attribute.Value) attribute.Value {
	_ = "STUB: not implemented"
	return *new(attribute.Value)
}

func stringNeedsTruncation(limit int, s string) bool { _ = "STUB: not implemented"; return false }

func needsTruncation(limit int, v attribute.Value) bool { _ = "STUB: not implemented"; return false }

func truncate(limit int, s string) string { _ = "STUB: not implemented"; return "" }

func (s *recordingSpan) End(options ...trace.SpanEndOption) { _ = "STUB: not implemented"; return }

func monotonicEndTime(start time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *recordingSpan) RecordError(err error, opts ...trace.EventOption) {
	_ = "STUB: not implemented"
	return
}

func typeStr(i any) string { _ = "STUB: not implemented"; return "" }

func recordStackTrace() string { _ = "STUB: not implemented"; return "" }

func (s *recordingSpan) AddEvent(name string, o ...trace.EventOption) {
	_ = "STUB: not implemented"
	return
}

func (s *recordingSpan) addEvent(name string, o ...trace.EventOption) {
	_ = "STUB: not implemented"
	return
}

func (s *recordingSpan) SetName(name string) { _ = "STUB: not implemented"; return }

func (s *recordingSpan) Name() string { _ = "STUB: not implemented"; return "" }

func (s *recordingSpan) Parent() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (s *recordingSpan) SpanKind() trace.SpanKind {
	_ = "STUB: not implemented"
	return *new(trace.SpanKind)
}

func (s *recordingSpan) StartTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *recordingSpan) EndTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *recordingSpan) Attributes() []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func (s *recordingSpan) dedupeAttrs() { _ = "STUB: not implemented"; return }

func (s *recordingSpan) dedupeAttrsFromRecord(record map[attribute.Key]int) {
	_ = "STUB: not implemented"
	return
}

func (s *recordingSpan) Links() []Link { _ = "STUB: not implemented"; return nil }

func (s *recordingSpan) Events() []Event { _ = "STUB: not implemented"; return nil }

func (s *recordingSpan) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

func (s *recordingSpan) InstrumentationScope() instrumentation.Scope {
	_ = "STUB: not implemented"
	return *new(instrumentation.Scope)
}

func (s *recordingSpan) InstrumentationLibrary() instrumentation.Library {
	_ = "STUB: not implemented" //nolint:staticcheck // This method needs to be define for backwards compatibility
	return *new(instrumentation.Library)
}

func (s *recordingSpan) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }

func (s *recordingSpan) AddLink(link trace.Link) { _ = "STUB: not implemented"; return }

func (s *recordingSpan) DroppedAttributes() int { _ = "STUB: not implemented"; return 0 }

func (s *recordingSpan) DroppedLinks() int { _ = "STUB: not implemented"; return 0 }

func (s *recordingSpan) DroppedEvents() int { _ = "STUB: not implemented"; return 0 }

func (s *recordingSpan) ChildSpanCount() int { _ = "STUB: not implemented"; return 0 }

func (s *recordingSpan) TracerProvider() trace.TracerProvider {
	_ = "STUB: not implemented"
	return *new(trace.TracerProvider)
}

func (s *recordingSpan) snapshot() ReadOnlySpan {
	_ = "STUB: not implemented"
	return *new(ReadOnlySpan)
}

func (s *recordingSpan) addChild() { _ = "STUB: not implemented"; return }

func (*recordingSpan) private() { _ = "STUB: not implemented"; return }

func (s *recordingSpan) runtimeTrace(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type nonRecordingSpan struct {
	embedded.Span

	tracer *tracer
	sc     trace.SpanContext
}

var _ trace.Span = nonRecordingSpan{}

func (s nonRecordingSpan) SpanContext() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (nonRecordingSpan) IsRecording() bool { _ = "STUB: not implemented"; return false }

func (nonRecordingSpan) SetStatus(codes.Code, string) { _ = "STUB: not implemented"; return }

func (nonRecordingSpan) SetError(bool) { _ = "STUB: not implemented"; return }

func (nonRecordingSpan) SetAttributes(...attribute.KeyValue) { _ = "STUB: not implemented"; return }

func (nonRecordingSpan) End(...trace.SpanEndOption) { _ = "STUB: not implemented"; return }

func (nonRecordingSpan) RecordError(error, ...trace.EventOption) { _ = "STUB: not implemented"; return }

func (nonRecordingSpan) AddEvent(string, ...trace.EventOption) { _ = "STUB: not implemented"; return }

func (nonRecordingSpan) AddLink(trace.Link) { _ = "STUB: not implemented"; return }

func (nonRecordingSpan) SetName(string) { _ = "STUB: not implemented"; return }

func (s nonRecordingSpan) TracerProvider() trace.TracerProvider {
	_ = "STUB: not implemented"
	return *new(trace.TracerProvider)
}

func isRecording(s SamplingResult) bool { _ = "STUB: not implemented"; return false }

func isSampled(s SamplingResult) bool { _ = "STUB: not implemented"; return false }

type Status struct {
	Code codes.Code

	Description string
}
