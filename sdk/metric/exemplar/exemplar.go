package exemplar

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type Exemplar struct {
	FilteredAttributes []attribute.KeyValue

	Time time.Time

	Value Value

	SpanID []byte `json:",omitempty"`

	TraceID []byte `json:",omitempty"`
}
