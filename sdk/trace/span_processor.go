package trace

import (
	"context"
	"sync"
)

type SpanProcessor interface {
	OnStart(parent context.Context, s ReadWriteSpan)

	OnEnd(s ReadOnlySpan)

	Shutdown(ctx context.Context) error

	ForceFlush(ctx context.Context) error
}

type spanProcessorState struct {
	sp    SpanProcessor
	state sync.Once
}

func newSpanProcessorState(sp SpanProcessor) *spanProcessorState {
	_ = "STUB: not implemented"
	return nil
}

type spanProcessorStates []*spanProcessorState
