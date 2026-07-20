package metric

import (
	"context"
	"sync/atomic"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/sdk/instrumentation"
)

type MeterProvider struct {
	embedded.MeterProvider

	pipes  pipelines
	meters cache[instrumentation.Scope, *meter]

	forceFlush, shutdown func(context.Context) error
	stopped              atomic.Bool
}

var _ metric.MeterProvider = (*MeterProvider)(nil)

func NewMeterProvider(options ...Option) *MeterProvider { _ = "STUB: not implemented"; return nil }

func (mp *MeterProvider) Meter(name string, options ...metric.MeterOption) metric.Meter {
	_ = "STUB: not implemented"
	return *new(metric.Meter)
}

func (mp *MeterProvider) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *MeterProvider) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
