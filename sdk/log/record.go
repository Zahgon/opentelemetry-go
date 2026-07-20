package log

import (
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/internal/global"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
)

const (
	attributesInlineCount = 5

	maxUniqueSize = 1028
)

var logAttrDropped = sync.OnceFunc(func() {
	global.Warn("limit reached: dropping log Record attributes")
})

var logKeyValuePairDropped = sync.OnceFunc(func() {
	global.Warn("key duplication: dropping key-value pair")
})

var uniquePool = sync.Pool{
	New: func() any { return new([]attribute.KeyValue) },
}

func getUnique() *[]attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func putUnique(v *[]attribute.KeyValue) { _ = "STUB: not implemented"; return }

var indexPool = sync.Pool{
	New: func() any { return make(map[attribute.Key]int) },
}

func getIndex() map[attribute.Key]int { _ = "STUB: not implemented"; return nil }

func putIndex(index map[attribute.Key]int) { _ = "STUB: not implemented"; return }

type Record struct {
	eventName         string
	timestamp         time.Time
	observedTimestamp time.Time
	severity          log.Severity
	severityText      string
	body              attribute.Value

	front [attributesInlineCount]attribute.KeyValue

	nFront int

	back []attribute.KeyValue

	dropped int

	traceID    trace.TraceID
	spanID     trace.SpanID
	traceFlags trace.TraceFlags

	resource *resource.Resource

	scope *instrumentation.Scope

	attributeValueLengthLimit int
	attributeCountLimit       int

	allowDupKeys bool

	noCmp [0]func() //nolint: unused  // This is indeed used.
}

func (r *Record) addDropped(n int) { _ = "STUB: not implemented"; return }

func (r *Record) EventName() string { _ = "STUB: not implemented"; return "" }

func (r *Record) SetEventName(s string) { _ = "STUB: not implemented"; return }

func (r *Record) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (r *Record) SetTimestamp(t time.Time) { _ = "STUB: not implemented"; return }

func (r *Record) ObservedTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (r *Record) SetObservedTimestamp(t time.Time) { _ = "STUB: not implemented"; return }

func (r *Record) Severity() log.Severity { _ = "STUB: not implemented"; return *new(log.Severity) }

func (r *Record) SetSeverity(level log.Severity) { _ = "STUB: not implemented"; return }

func (r *Record) SeverityText() string { _ = "STUB: not implemented"; return "" }

func (r *Record) SetSeverityText(text string) { _ = "STUB: not implemented"; return }

func (r *Record) Body() attribute.Value { _ = "STUB: not implemented"; return *new(attribute.Value) }

func (r *Record) SetBody(v attribute.Value) { _ = "STUB: not implemented"; return }

func (r *Record) WalkAttributes(f func(attribute.KeyValue) bool) { _ = "STUB: not implemented"; return }

func (r *Record) AddAttributes(attrs ...attribute.KeyValue) { _ = "STUB: not implemented"; return }

func (r *Record) attrIndex() map[attribute.Key]int { _ = "STUB: not implemented"; return nil }

func (r *Record) addAttrs(attrs []attribute.KeyValue) { _ = "STUB: not implemented"; return }

func (r *Record) SetAttributes(attrs ...attribute.KeyValue) { _ = "STUB: not implemented"; return }

func (r *Record) head(kvs []attribute.KeyValue) (out []attribute.KeyValue, dropped int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (r *Record) hasAttributeCountLimit() bool { _ = "STUB: not implemented"; return false }

func dedup(kvs []attribute.KeyValue) (unique []attribute.KeyValue, dropped int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (r *Record) AttributesLen() int { _ = "STUB: not implemented"; return 0 }

func (r *Record) DroppedAttributes() int { _ = "STUB: not implemented"; return 0 }

func (r *Record) TraceID() trace.TraceID { _ = "STUB: not implemented"; return *new(trace.TraceID) }

func (r *Record) SetTraceID(id trace.TraceID) { _ = "STUB: not implemented"; return }

func (r *Record) SpanID() trace.SpanID { _ = "STUB: not implemented"; return *new(trace.SpanID) }

func (r *Record) SetSpanID(id trace.SpanID) { _ = "STUB: not implemented"; return }

func (r *Record) TraceFlags() trace.TraceFlags {
	_ = "STUB: not implemented"
	return *new(trace.TraceFlags)
}

func (r *Record) SetTraceFlags(flags trace.TraceFlags) { _ = "STUB: not implemented"; return }

func (r *Record) Resource() *resource.Resource { _ = "STUB: not implemented"; return nil }

func (r *Record) InstrumentationScope() instrumentation.Scope {
	_ = "STUB: not implemented"
	return *new(instrumentation.Scope)
}

func (r *Record) Clone() Record { _ = "STUB: not implemented"; return *new(Record) }

func (r *Record) applyAttrLimitsAndDedup(attr attribute.KeyValue) attribute.KeyValue {
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
