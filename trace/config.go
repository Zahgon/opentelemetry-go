package trace

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type TracerConfig struct {
	instrumentationVersion string

	schemaURL string
	attrs     attribute.Set
}

func (t *TracerConfig) InstrumentationVersion() string { _ = "STUB: not implemented"; return "" }

func (t *TracerConfig) InstrumentationAttributes() attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

func (t *TracerConfig) SchemaURL() string { _ = "STUB: not implemented"; return "" }

type experimentalOption interface {
	Experimental()
}

func NewTracerConfig(options ...TracerOption) TracerConfig {
	_ = "STUB: not implemented"
	return *new(TracerConfig)
}

type TracerOption interface {
	apply(TracerConfig) TracerConfig
}

type tracerOptionFunc func(TracerConfig) TracerConfig

func (fn tracerOptionFunc) apply(cfg TracerConfig) TracerConfig {
	_ = "STUB: not implemented"
	return *new(TracerConfig)
}

type SpanConfig struct {
	attributes []attribute.KeyValue
	timestamp  time.Time
	links      []Link
	newRoot    bool
	spanKind   SpanKind
	stackTrace bool
}

func (cfg *SpanConfig) Attributes() []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func (cfg *SpanConfig) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (cfg *SpanConfig) StackTrace() bool { _ = "STUB: not implemented"; return false }

func (cfg *SpanConfig) Links() []Link { _ = "STUB: not implemented"; return nil }

func (cfg *SpanConfig) NewRoot() bool { _ = "STUB: not implemented"; return false }

func (cfg *SpanConfig) SpanKind() SpanKind { _ = "STUB: not implemented"; return *new(SpanKind) }

func NewSpanStartConfig(options ...SpanStartOption) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}

func NewSpanEndConfig(options ...SpanEndOption) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}

type SpanStartOption interface {
	applySpanStart(SpanConfig) SpanConfig
}

type spanOptionFunc func(SpanConfig) SpanConfig

func (fn spanOptionFunc) applySpanStart(cfg SpanConfig) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}

type SpanEndOption interface {
	applySpanEnd(SpanConfig) SpanConfig
}

type EventConfig struct {
	attributes []attribute.KeyValue
	timestamp  time.Time
	stackTrace bool
}

func (cfg *EventConfig) Attributes() []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func (cfg *EventConfig) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (cfg *EventConfig) StackTrace() bool { _ = "STUB: not implemented"; return false }

func NewEventConfig(options ...EventOption) EventConfig {
	_ = "STUB: not implemented"
	return *new(EventConfig)
}

type EventOption interface {
	applyEvent(EventConfig) EventConfig
}

type SpanOption interface {
	SpanStartOption
	SpanEndOption
}

type SpanStartEventOption interface {
	SpanStartOption
	EventOption
}

type SpanEndEventOption interface {
	SpanEndOption
	EventOption
}

type attributeOption []attribute.KeyValue

func (o attributeOption) applySpan(c SpanConfig) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}

func (o attributeOption) applySpanStart(c SpanConfig) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}
func (o attributeOption) applyEvent(c EventConfig) EventConfig {
	_ = "STUB: not implemented"
	return *new(EventConfig)
}

var _ SpanStartEventOption = attributeOption{}

func WithAttributes(attributes ...attribute.KeyValue) SpanStartEventOption {
	_ = "STUB: not implemented"
	return *new(SpanStartEventOption)
}

type SpanEventOption interface {
	SpanOption
	EventOption
}

type timestampOption time.Time

func (o timestampOption) applySpan(c SpanConfig) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}

func (o timestampOption) applySpanStart(c SpanConfig) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}
func (o timestampOption) applySpanEnd(c SpanConfig) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}
func (o timestampOption) applyEvent(c EventConfig) EventConfig {
	_ = "STUB: not implemented"
	return *new(EventConfig)
}

var _ SpanEventOption = timestampOption{}

func WithTimestamp(t time.Time) SpanEventOption {
	_ = "STUB: not implemented"
	return *new(SpanEventOption)
}

type stackTraceOption bool

func (o stackTraceOption) applyEvent(c EventConfig) EventConfig {
	_ = "STUB: not implemented"
	return *new(EventConfig)
}

func (o stackTraceOption) applySpan(c SpanConfig) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}

func (o stackTraceOption) applySpanEnd(c SpanConfig) SpanConfig {
	_ = "STUB: not implemented"
	return *new(SpanConfig)
}

func WithStackTrace(b bool) SpanEndEventOption {
	_ = "STUB: not implemented"
	return *new(SpanEndEventOption)
}

func WithLinks(links ...Link) SpanStartOption {
	_ = "STUB: not implemented"
	return *new(SpanStartOption)
}

func WithNewRoot() SpanStartOption { _ = "STUB: not implemented"; return *new(SpanStartOption) }

func WithSpanKind(kind SpanKind) SpanStartOption {
	_ = "STUB: not implemented"
	return *new(SpanStartOption)
}

func WithInstrumentationVersion(version string) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}

func mergeSets(a, b attribute.Set) attribute.Set {
	_ = "STUB: not implemented"
	return *new(attribute.Set)
}

func WithInstrumentationAttributes(attr ...attribute.KeyValue) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}

func WithInstrumentationAttributeSet(set attribute.Set) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}

func WithSchemaURL(schemaURL string) TracerOption {
	_ = "STUB: not implemented"
	return *new(TracerOption)
}
