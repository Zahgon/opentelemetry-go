package otel

import (
	"go.opentelemetry.io/otel/propagation"
)

func GetTextMapPropagator() propagation.TextMapPropagator {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator)
}

func SetTextMapPropagator(propagator propagation.TextMapPropagator) {
	_ = "STUB: not implemented"
	return
}
