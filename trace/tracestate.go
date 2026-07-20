package trace

import (
	"encoding/json"
)

const (
	maxListMembers = 32

	listDelimiters  = ","
	memberDelimiter = "="

	errInvalidKey    errorConst = "invalid tracestate key"
	errInvalidValue  errorConst = "invalid tracestate value"
	errInvalidMember errorConst = "invalid tracestate list-member"
	errMemberNumber  errorConst = "too many list-members in tracestate"
	errDuplicate     errorConst = "duplicate list-member in tracestate"
)

type member struct {
	Key   string
	Value string
}

func checkValueChar(v byte) bool { _ = "STUB: not implemented"; return false }

func checkValueLast(v byte) bool { _ = "STUB: not implemented"; return false }

func checkValue(val string) bool { _ = "STUB: not implemented"; return false }

func checkKeyRemain(key string) bool { _ = "STUB: not implemented"; return false }

func checkKeyPart(key string, n int) bool { _ = "STUB: not implemented"; return false }

func isAlphaNumASCII[T rune | byte](c T) bool { _ = "STUB: not implemented"; return false }

func checkKeyTenant(key string, n int) bool { _ = "STUB: not implemented"; return false }

func checkKey(key string) bool { _ = "STUB: not implemented"; return false }

func newMember(key, value string) (member, error) {
	_ = "STUB: not implemented"
	return *new(member), nil
}

func parseMember(m string) (member, error) { _ = "STUB: not implemented"; return *new(member), nil }

func (m member) String() string { _ = "STUB: not implemented"; return "" }

type TraceState struct { //nolint:revive // revive complains about stutter of `trace.TraceState`

	list []member
}

var _ json.Marshaler = TraceState{}

func ParseTraceState(ts string) (TraceState, error) {
	_ = "STUB: not implemented"
	return *new(TraceState), nil
}

func (ts TraceState) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ts TraceState) String() string { _ = "STUB: not implemented"; return "" }

func (ts TraceState) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (ts TraceState) Walk(f func(key, value string) bool) { _ = "STUB: not implemented"; return }

func (ts TraceState) Insert(key, value string) (TraceState, error) {
	_ = "STUB: not implemented"
	return *new(TraceState), nil
}

func (ts TraceState) Delete(key string) TraceState {
	_ = "STUB: not implemented"
	return *new(TraceState)
}

func (ts TraceState) Len() int { _ = "STUB: not implemented"; return 0 }
