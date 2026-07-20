package otel

import (
	"go.opentelemetry.io/otel/metric"
)

func Meter(name string, opts ...metric.MeterOption) metric.Meter {
	_ = "STUB: not implemented"
	return *new(metric.Meter)
}

func GetMeterProvider() metric.MeterProvider {
	_ = "STUB: not implemented"
	return *new(metric.MeterProvider)
}

func SetMeterProvider(mp metric.MeterProvider) { _ = "STUB: not implemented"; return }
