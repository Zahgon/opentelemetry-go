package exemplar

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type Reservoir interface {
	Offer(ctx context.Context, t time.Time, val Value, attr []attribute.KeyValue)

	Collect(dest *[]Exemplar)
}

type ReservoirProvider func(attr attribute.Set) Reservoir
