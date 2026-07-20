package tracetest

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type SpanStubs []SpanStub

func SpanStubsFromReadOnlySpans(ro []tracesdk.ReadOnlySpan) SpanStubs {
	_ = "STUB: not implemented"
	return *new(SpanStubs)
}

func (s SpanStubs) Snapshots() []tracesdk.ReadOnlySpan { _ = "STUB: not implemented"; return nil }

type SpanStub struct {
	Name                 string
	SpanContext          trace.SpanContext
	Parent               trace.SpanContext
	SpanKind             trace.SpanKind
	StartTime            time.Time
	EndTime              time.Time
	Attributes           []attribute.KeyValue
	Events               []tracesdk.Event
	Links                []tracesdk.Link
	Status               tracesdk.Status
	DroppedAttributes    int
	DroppedEvents        int
	DroppedLinks         int
	ChildSpanCount       int
	Resource             *resource.Resource
	InstrumentationScope instrumentation.Scope

	InstrumentationLibrary instrumentation.Library //nolint:staticcheck // This method needs to be define for backwards compatibility
}

func SpanStubFromReadOnlySpan(ro tracesdk.ReadOnlySpan) SpanStub {
	_ = "STUB: not implemented"
	return *new(SpanStub)
}

func (s SpanStub) Snapshot() tracesdk.ReadOnlySpan {
	_ = "STUB: not implemented"
	return *new(tracesdk.ReadOnlySpan)
}

type spanSnapshot struct {
	tracesdk.ReadOnlySpan

	name                 string
	spanContext          trace.SpanContext
	parent               trace.SpanContext
	spanKind             trace.SpanKind
	startTime            time.Time
	endTime              time.Time
	attributes           []attribute.KeyValue
	events               []tracesdk.Event
	links                []tracesdk.Link
	status               tracesdk.Status
	droppedAttributes    int
	droppedEvents        int
	droppedLinks         int
	childSpanCount       int
	resource             *resource.Resource
	instrumentationScope instrumentation.Scope
}

func (s spanSnapshot) Name() string { _ = "STUB: not implemented"; return "" }
func (s spanSnapshot) SpanContext() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}
func (s spanSnapshot) Parent() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}
func (s spanSnapshot) SpanKind() trace.SpanKind {
	_ = "STUB: not implemented"
	return *new(trace.SpanKind)
}
func (s spanSnapshot) StartTime() time.Time             { _ = "STUB: not implemented"; return *new(time.Time) }
func (s spanSnapshot) EndTime() time.Time               { _ = "STUB: not implemented"; return *new(time.Time) }
func (s spanSnapshot) Attributes() []attribute.KeyValue { _ = "STUB: not implemented"; return nil }
func (s spanSnapshot) Links() []tracesdk.Link           { _ = "STUB: not implemented"; return nil }
func (s spanSnapshot) Events() []tracesdk.Event         { _ = "STUB: not implemented"; return nil }
func (s spanSnapshot) Status() tracesdk.Status {
	_ = "STUB: not implemented"
	return *new(tracesdk.Status)
}
func (s spanSnapshot) DroppedAttributes() int       { _ = "STUB: not implemented"; return 0 }
func (s spanSnapshot) DroppedLinks() int            { _ = "STUB: not implemented"; return 0 }
func (s spanSnapshot) DroppedEvents() int           { _ = "STUB: not implemented"; return 0 }
func (s spanSnapshot) ChildSpanCount() int          { _ = "STUB: not implemented"; return 0 }
func (s spanSnapshot) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }
func (s spanSnapshot) InstrumentationScope() instrumentation.Scope {
	_ = "STUB: not implemented"
	return *new(instrumentation.Scope)
}

func (s spanSnapshot) InstrumentationLibrary() instrumentation.Library {
	_ = "STUB: not implemented" //nolint:staticcheck // This method needs to be define for backwards compatibility
	return *new(instrumentation.Library)
}
