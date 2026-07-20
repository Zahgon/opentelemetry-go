package exemplar

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/internal/reservoir"
)

func HistogramReservoirProvider(bounds []float64) ReservoirProvider {
	_ = "STUB: not implemented"
	return *new(ReservoirProvider)
}

func NewHistogramReservoir(bounds []float64) *HistogramReservoir {
	_ = "STUB: not implemented"
	return nil
}

var _ Reservoir = &HistogramReservoir{}

type HistogramReservoir struct {
	reservoir.ConcurrentSafe
	*storage

	bounds []float64
}

func (r *HistogramReservoir) Offer(ctx context.Context, t time.Time, v Value, a []attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (r *HistogramReservoir) Collect(dest *[]Exemplar) { _ = "STUB: not implemented"; return }
