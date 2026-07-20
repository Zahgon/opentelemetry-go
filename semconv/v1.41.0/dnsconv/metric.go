package dnsconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type ErrorTypeAttr string

var (
	ErrorTypeOther ErrorTypeAttr = "_OTHER"
)

type LookupDuration struct {
	metric.Float64Histogram
}

var newLookupDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Measures the time taken to perform a DNS lookup."),
	metric.WithUnit("s"),
}

func NewLookupDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (LookupDuration, error) {
	_ = "STUB: not implemented"
	return *new(LookupDuration), nil
}

func (m LookupDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (LookupDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (LookupDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (LookupDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m LookupDuration) Record(
	ctx context.Context,
	val float64,
	questionName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m LookupDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (LookupDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
