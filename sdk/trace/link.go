package trace

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Link struct {
	SpanContext trace.SpanContext

	Attributes []attribute.KeyValue

	DroppedAttributeCount int
}
