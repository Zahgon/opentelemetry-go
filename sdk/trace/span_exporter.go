package trace

import "context"

type SpanExporter interface {
	ExportSpans(ctx context.Context, spans []ReadOnlySpan) error

	Shutdown(ctx context.Context) error
}
