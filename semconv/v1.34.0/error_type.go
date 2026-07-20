package semconv

import (
	"go.opentelemetry.io/otel/attribute"
)

func ErrorType(err error) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
