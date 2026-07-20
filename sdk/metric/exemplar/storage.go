package exemplar

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type storage struct {
	measurements []measurement
}

func newStorage(n int) *storage { _ = "STUB: not implemented"; return nil }

func (r *storage) store(ctx context.Context, idx int, ts time.Time, v Value, droppedAttr []attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (r *storage) Collect(dest *[]Exemplar) { _ = "STUB: not implemented"; return }

type measurement struct {
	mux sync.Mutex

	FilteredAttributes []attribute.KeyValue

	Time time.Time

	Value Value

	SpanContext trace.SpanContext

	valid bool
}

func (m *measurement) exemplar(dest *Exemplar) bool { _ = "STUB: not implemented"; return false }

func reset[T any](s []T, length, capacity int) []T { _ = "STUB: not implemented"; return nil }
