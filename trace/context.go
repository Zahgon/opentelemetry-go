package trace

import "context"

type traceContextKeyType int

const currentSpanKey traceContextKeyType = iota

func ContextWithSpan(parent context.Context, span Span) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ContextWithSpanContext(parent context.Context, sc SpanContext) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ContextWithRemoteSpanContext(parent context.Context, rsc SpanContext) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SpanFromContext(ctx context.Context) Span { _ = "STUB: not implemented"; return *new(Span) }

func SpanContextFromContext(ctx context.Context) SpanContext {
	_ = "STUB: not implemented"
	return *new(SpanContext)
}
