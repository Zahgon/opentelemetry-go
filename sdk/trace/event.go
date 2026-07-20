package trace

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type Event struct {
	Name string

	Attributes []attribute.KeyValue

	DroppedAttributeCount int

	Time time.Time
}
