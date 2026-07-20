package trace

import (
	"encoding/json"
)

const (
	FlagsSampled = TraceFlags(0x01)

	FlagsRandom = TraceFlags(0x02)

	errInvalidHexID errorConst = "trace-id and span-id can only contain [0-9a-f] characters, all lowercase"

	errInvalidTraceIDLength errorConst = "hex encoded trace-id must have length equals to 32"
	errNilTraceID           errorConst = "trace-id can't be all zero"

	errInvalidSpanIDLength errorConst = "hex encoded span-id must have length equals to 16"
	errNilSpanID           errorConst = "span-id can't be all zero"
)

type errorConst string

func (e errorConst) Error() string { _ = "STUB: not implemented"; return "" }

type TraceID [16]byte

var (
	nilTraceID TraceID
	_          json.Marshaler = nilTraceID
)

func (t TraceID) IsValid() bool { _ = "STUB: not implemented"; return false }

func (t TraceID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t TraceID) String() string { _ = "STUB: not implemented"; return "" }

func (t TraceID) hexBytes() [32]byte { _ = "STUB: not implemented"; return [32]byte{} }

type SpanID [8]byte

var (
	nilSpanID SpanID
	_         json.Marshaler = nilSpanID
)

func (s SpanID) IsValid() bool { _ = "STUB: not implemented"; return false }

func (s SpanID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s SpanID) String() string { _ = "STUB: not implemented"; return "" }

func (s SpanID) hexBytes() [16]byte { _ = "STUB: not implemented"; return [16]byte{} }

func TraceIDFromHex(h string) (TraceID, error) {
	_ = "STUB: not implemented"
	return *new(TraceID), nil
}

func SpanIDFromHex(h string) (SpanID, error) { _ = "STUB: not implemented"; return *new(SpanID), nil }

type TraceFlags byte //nolint:revive // revive complains about stutter of `trace.TraceFlags`.

func (tf TraceFlags) IsSampled() bool { _ = "STUB: not implemented"; return false }

func (tf TraceFlags) WithSampled(sampled bool) TraceFlags {
	_ = "STUB: not implemented"
	return *new(TraceFlags)
}

func (tf TraceFlags) IsRandom() bool { _ = "STUB: not implemented"; return false }

func (tf TraceFlags) WithRandom(random bool) TraceFlags {
	_ = "STUB: not implemented"
	return *new(TraceFlags)
}

func (tf TraceFlags) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (tf TraceFlags) String() string { _ = "STUB: not implemented"; return "" }

func (tf TraceFlags) hexBytes() [2]byte { _ = "STUB: not implemented"; return [2]byte{} }

type SpanContextConfig struct {
	TraceID    TraceID
	SpanID     SpanID
	TraceFlags TraceFlags
	TraceState TraceState
	Remote     bool
}

func NewSpanContext(config SpanContextConfig) SpanContext {
	_ = "STUB: not implemented"
	return *new(SpanContext)
}

type SpanContext struct {
	traceID    TraceID
	spanID     SpanID
	traceFlags TraceFlags
	traceState TraceState
	remote     bool
}

var _ json.Marshaler = SpanContext{}

func (sc SpanContext) IsValid() bool { _ = "STUB: not implemented"; return false }

func (sc SpanContext) IsRemote() bool { _ = "STUB: not implemented"; return false }

func (sc SpanContext) WithRemote(remote bool) SpanContext {
	_ = "STUB: not implemented"
	return *new(SpanContext)
}

func (sc SpanContext) TraceID() TraceID { _ = "STUB: not implemented"; return *new(TraceID) }

func (sc SpanContext) HasTraceID() bool { _ = "STUB: not implemented"; return false }

func (sc SpanContext) WithTraceID(traceID TraceID) SpanContext {
	_ = "STUB: not implemented"
	return *new(SpanContext)
}

func (sc SpanContext) SpanID() SpanID { _ = "STUB: not implemented"; return *new(SpanID) }

func (sc SpanContext) HasSpanID() bool { _ = "STUB: not implemented"; return false }

func (sc SpanContext) WithSpanID(spanID SpanID) SpanContext {
	_ = "STUB: not implemented"
	return *new(SpanContext)
}

func (sc SpanContext) TraceFlags() TraceFlags { _ = "STUB: not implemented"; return *new(TraceFlags) }

func (sc SpanContext) IsSampled() bool { _ = "STUB: not implemented"; return false }

func (sc SpanContext) IsRandom() bool { _ = "STUB: not implemented"; return false }

func (sc SpanContext) WithTraceFlags(flags TraceFlags) SpanContext {
	_ = "STUB: not implemented"
	return *new(SpanContext)
}

func (sc SpanContext) TraceState() TraceState { _ = "STUB: not implemented"; return *new(TraceState) }

func (sc SpanContext) WithTraceState(state TraceState) SpanContext {
	_ = "STUB: not implemented"
	return *new(SpanContext)
}

func (sc SpanContext) Equal(other SpanContext) bool { _ = "STUB: not implemented"; return false }

func (sc SpanContext) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
