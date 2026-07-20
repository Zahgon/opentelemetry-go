package logtest

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
)

type RecordFactory struct {
	EventName         string
	Timestamp         time.Time
	ObservedTimestamp time.Time
	Severity          log.Severity
	SeverityText      string
	Body              attribute.Value
	Attributes        []attribute.KeyValue
	TraceID           trace.TraceID
	SpanID            trace.SpanID
	TraceFlags        trace.TraceFlags

	Resource             *resource.Resource
	InstrumentationScope *instrumentation.Scope

	DroppedAttributes int
}

func (f RecordFactory) NewRecord() sdklog.Record {
	_ = "STUB: not implemented"
	return *new(sdklog.Record)
}

func set(r *sdklog.Record, name string, value any) { _ = "STUB: not implemented"; return }
