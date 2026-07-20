package metric

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/exemplar"
	"go.opentelemetry.io/otel/sdk/metric/internal/aggregate"
)

type ExemplarReservoirProviderSelector func(Aggregation) exemplar.ReservoirProvider

func reservoirFunc[N int64 | float64](
	kind InstrumentKind,
	provider exemplar.ReservoirProvider,
	filter exemplar.Filter,
) func(attribute.Set) aggregate.FilteredExemplarReservoir[N] {
	_ = "STUB: not implemented"
	return nil
}

func DefaultExemplarReservoirProviderSelector(agg Aggregation) exemplar.ReservoirProvider {
	_ = "STUB: not implemented"
	return *new(exemplar.ReservoirProvider)
}
