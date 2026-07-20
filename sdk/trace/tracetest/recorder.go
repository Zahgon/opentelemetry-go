package tracetest

import (
	"context"
	"sync"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type SpanRecorder struct {
	startedMu sync.RWMutex
	started   []sdktrace.ReadWriteSpan

	endedMu sync.RWMutex
	ended   []sdktrace.ReadOnlySpan
}

var _ sdktrace.SpanProcessor = (*SpanRecorder)(nil)

func NewSpanRecorder() *SpanRecorder { _ = "STUB: not implemented"; return nil }

func (sr *SpanRecorder) OnStart(_ context.Context, s sdktrace.ReadWriteSpan) {
	_ = "STUB: not implemented"
	return
}

func (sr *SpanRecorder) OnEnd(s sdktrace.ReadOnlySpan) { _ = "STUB: not implemented"; return }

func (*SpanRecorder) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (*SpanRecorder) ForceFlush(context.Context) error { _ = "STUB: not implemented"; return nil }

func (sr *SpanRecorder) Started() []sdktrace.ReadWriteSpan { _ = "STUB: not implemented"; return nil }

func (sr *SpanRecorder) Reset() { _ = "STUB: not implemented"; return }

func (sr *SpanRecorder) Ended() []sdktrace.ReadOnlySpan { _ = "STUB: not implemented"; return nil }
