package trace

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
)

type snapshot struct {
	name                  string
	spanContext           trace.SpanContext
	parent                trace.SpanContext
	spanKind              trace.SpanKind
	startTime             time.Time
	endTime               time.Time
	attributes            []attribute.KeyValue
	events                []Event
	links                 []Link
	status                Status
	childSpanCount        int
	droppedAttributeCount int
	droppedEventCount     int
	droppedLinkCount      int
	resource              *resource.Resource
	instrumentationScope  instrumentation.Scope
}

var _ ReadOnlySpan = snapshot{}

func (snapshot) private() { _ = "STUB: not implemented"; return }

func (s snapshot) Name() string { _ = "STUB: not implemented"; return "" }

func (s snapshot) SpanContext() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (s snapshot) Parent() trace.SpanContext {
	_ = "STUB: not implemented"
	return *new(trace.SpanContext)
}

func (s snapshot) SpanKind() trace.SpanKind { _ = "STUB: not implemented"; return *new(trace.SpanKind) }

func (s snapshot) StartTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s snapshot) EndTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s snapshot) Attributes() []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func (s snapshot) Links() []Link { _ = "STUB: not implemented"; return nil }

func (s snapshot) Events() []Event { _ = "STUB: not implemented"; return nil }

func (s snapshot) Status() Status { _ = "STUB: not implemented"; return *new(Status) }

func (s snapshot) InstrumentationScope() instrumentation.Scope {
	_ = "STUB: not implemented"
	return *new(instrumentation.Scope)
}

func (s snapshot) InstrumentationLibrary() instrumentation.Library {
	_ = "STUB: not implemented" //nolint:staticcheck // This method needs to be define for backwards compatibility
	return *new(instrumentation.Library)
}

func (s snapshot) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }

func (s snapshot) DroppedAttributes() int { _ = "STUB: not implemented"; return 0 }

func (s snapshot) DroppedLinks() int { _ = "STUB: not implemented"; return 0 }

func (s snapshot) DroppedEvents() int { _ = "STUB: not implemented"; return 0 }

func (s snapshot) ChildSpanCount() int { _ = "STUB: not implemented"; return 0 }
