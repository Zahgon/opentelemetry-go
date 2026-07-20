package x

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type defaultAttributesOption struct {
	metric.InstrumentOption
	keys []attribute.Key
}

func (defaultAttributesOption) Experimental() { _ = "STUB: not implemented"; return }

func (o defaultAttributesOption) AllowedKeys() []attribute.Key {
	_ = "STUB: not implemented"
	return nil
}

func WithDefaultAttributes(keys ...attribute.Key) metric.InstrumentOption {
	_ = "STUB: not implemented"
	return *new(metric.InstrumentOption)
}

type unsafeAttributesOption struct {
	metric.MeasurementOption
	kvs []attribute.KeyValue
}

func (*unsafeAttributesOption) Experimental() { _ = "STUB: not implemented"; return }

func (o *unsafeAttributesOption) RawAttributes() []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func (o *unsafeAttributesOption) Set(kvs []attribute.KeyValue) { _ = "STUB: not implemented"; return }

func WithUnsafeAttributes(kvs ...attribute.KeyValue) metric.MeasurementOption {
	_ = "STUB: not implemented"
	return *new(metric.MeasurementOption)
}

type Settable[T any] interface {
	Set(T)
}
