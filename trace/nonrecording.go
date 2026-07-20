package trace

type nonRecordingSpan struct {
	noopSpan

	sc SpanContext
}

func (s nonRecordingSpan) SpanContext() SpanContext {
	_ = "STUB: not implemented"
	return *new(SpanContext)
}
