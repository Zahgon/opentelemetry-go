package metric

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

var ErrExporterShutdown = errors.New("exporter is shutdown")

type Exporter interface {
	Temporality(InstrumentKind) metricdata.Temporality

	Aggregation(InstrumentKind) Aggregation

	Export(context.Context, *metricdata.ResourceMetrics) error

	ForceFlush(context.Context) error

	Shutdown(context.Context) error
}
