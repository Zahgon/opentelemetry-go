package stdoutlog

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
)

type recordJSON struct {
	Timestamp         *time.Time `json:",omitempty"`
	ObservedTimestamp *time.Time `json:",omitempty"`
	EventName         string     `json:",omitempty"`
	Severity          log.Severity
	SeverityText      string
	Body              attribute.Value
	Attributes        []attribute.KeyValue
	TraceID           trace.TraceID
	SpanID            trace.SpanID
	TraceFlags        trace.TraceFlags
	Resource          *resource.Resource
	Scope             instrumentation.Scope
	DroppedAttributes int
}

func (e *Exporter) newRecordJSON(r sdklog.Record) recordJSON {
	_ = "STUB: not implemented"
	return *new(recordJSON)
}
