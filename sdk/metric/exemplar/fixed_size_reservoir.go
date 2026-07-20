package exemplar

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/internal/reservoir"
)

func FixedSizeReservoirProvider(k int) ReservoirProvider {
	_ = "STUB: not implemented"
	return *new(ReservoirProvider)
}

func NewFixedSizeReservoir(k int) *FixedSizeReservoir { _ = "STUB: not implemented"; return nil }

var _ Reservoir = &FixedSizeReservoir{}

type FixedSizeReservoir struct {
	reservoir.ConcurrentSafe
	*storage
	mu sync.Mutex

	count int64

	next int64

	w float64
}

func newFixedSizeReservoir(s *storage) *FixedSizeReservoir { _ = "STUB: not implemented"; return nil }

func (*FixedSizeReservoir) randomFloat64() float64 { _ = "STUB: not implemented"; return 0 }

func (r *FixedSizeReservoir) Offer(ctx context.Context, t time.Time, n Value, a []attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (r *FixedSizeReservoir) reset() { _ = "STUB: not implemented"; return }

func (r *FixedSizeReservoir) advance() { _ = "STUB: not implemented"; return }

func (r *FixedSizeReservoir) Collect(dest *[]Exemplar) { _ = "STUB: not implemented"; return }
