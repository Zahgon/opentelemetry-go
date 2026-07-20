package log

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
)

const attributesInlineCount = 5

type Record struct {
	noCmp [0]func() //nolint: unused  // This is indeed used.

	eventName         string
	timestamp         time.Time
	observedTimestamp time.Time
	severity          Severity
	severityText      string
	body              attribute.Value
	err               error

	front [attributesInlineCount]attribute.KeyValue

	nFront int

	back []attribute.KeyValue
}

func (r *Record) EventName() string { _ = "STUB: not implemented"; return "" }

func (r *Record) SetEventName(s string) { _ = "STUB: not implemented"; return }

func (r *Record) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (r *Record) SetTimestamp(t time.Time) { _ = "STUB: not implemented"; return }

func (r *Record) ObservedTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (r *Record) SetObservedTimestamp(t time.Time) { _ = "STUB: not implemented"; return }

func (r *Record) Severity() Severity { _ = "STUB: not implemented"; return *new(Severity) }

func (r *Record) SetSeverity(level Severity) { _ = "STUB: not implemented"; return }

func (r *Record) SeverityText() string { _ = "STUB: not implemented"; return "" }

func (r *Record) SetSeverityText(text string) { _ = "STUB: not implemented"; return }

func (r *Record) Body() attribute.Value { _ = "STUB: not implemented"; return *new(attribute.Value) }

func (r *Record) SetBody(v attribute.Value) { _ = "STUB: not implemented"; return }

func (r *Record) Err() error { _ = "STUB: not implemented"; return nil }

func (r *Record) SetErr(err error) { _ = "STUB: not implemented"; return }

func (r *Record) WalkAttributes(f func(attribute.KeyValue) bool) { _ = "STUB: not implemented"; return }

func (r *Record) AddAttributes(attrs ...attribute.KeyValue) { _ = "STUB: not implemented"; return }

func (r *Record) AttributesLen() int { _ = "STUB: not implemented"; return 0 }

func (r *Record) Clone() Record { _ = "STUB: not implemented"; return *new(Record) }
