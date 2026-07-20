package attribute

import (
	"reflect"
	"strings"
)

//go:generate stringer -type=Type

type Type int

type Value struct {
	vtype    Type
	numeric  uint64
	stringly string
	slice    any
}

const (
	EMPTY Type = iota

	BOOL

	INT64

	FLOAT64

	STRING

	BOOLSLICE

	INT64SLICE

	FLOAT64SLICE

	STRINGSLICE

	BYTESLICE

	SLICE

	MAP

	INVALID = EMPTY
)

func BoolValue(v bool) Value { _ = "STUB: not implemented"; return *new(Value) }

func BoolSliceValue(v []bool) Value { _ = "STUB: not implemented"; return *new(Value) }

func IntValue(v int) Value { _ = "STUB: not implemented"; return *new(Value) }

func IntSliceValue(v []int) Value { _ = "STUB: not implemented"; return *new(Value) }

func Int64Value(v int64) Value { _ = "STUB: not implemented"; return *new(Value) }

func Int64SliceValue(v []int64) Value { _ = "STUB: not implemented"; return *new(Value) }

func Float64Value(v float64) Value { _ = "STUB: not implemented"; return *new(Value) }

func Float64SliceValue(v []float64) Value { _ = "STUB: not implemented"; return *new(Value) }

func StringValue(v string) Value { _ = "STUB: not implemented"; return *new(Value) }

func StringSliceValue(v []string) Value { _ = "STUB: not implemented"; return *new(Value) }

func ByteSliceValue(v []byte) Value { _ = "STUB: not implemented"; return *new(Value) }

func SliceValue(v ...Value) Value { _ = "STUB: not implemented"; return *new(Value) }

func MapValue(v ...KeyValue) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v Value) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (v Value) AsBool() bool { _ = "STUB: not implemented"; return false }

func (v Value) AsBoolSlice() []bool { _ = "STUB: not implemented"; return nil }

func (v Value) asBoolSlice() []bool { _ = "STUB: not implemented"; return nil }

func (v Value) AsInt64() int64 { _ = "STUB: not implemented"; return 0 }

func (v Value) AsInt64Slice() []int64 { _ = "STUB: not implemented"; return nil }

func (v Value) asInt64Slice() []int64 { _ = "STUB: not implemented"; return nil }

func (v Value) AsFloat64() float64 { _ = "STUB: not implemented"; return 0 }

func (v Value) AsFloat64Slice() []float64 { _ = "STUB: not implemented"; return nil }

func (v Value) asFloat64Slice() []float64 { _ = "STUB: not implemented"; return nil }

func (v Value) AsString() string { _ = "STUB: not implemented"; return "" }

func (v Value) AsStringSlice() []string { _ = "STUB: not implemented"; return nil }

func (v Value) asStringSlice() []string { _ = "STUB: not implemented"; return nil }

func (v Value) AsSlice() []Value { _ = "STUB: not implemented"; return nil }

func (v Value) asSlice() []Value { _ = "STUB: not implemented"; return nil }

func asValueSliceReflect(v any) []Value { _ = "STUB: not implemented"; return nil }

func (v Value) AsMap() []KeyValue { _ = "STUB: not implemented"; return nil }

func (v Value) asMap() []KeyValue { _ = "STUB: not implemented"; return nil }

func asKeyValueSliceReflect(v any) []KeyValue { _ = "STUB: not implemented"; return nil }

func (v Value) AsByteSlice() []byte { _ = "STUB: not implemented"; return nil }

func (v Value) asByteSlice() []byte { _ = "STUB: not implemented"; return nil }

type unknownValueType struct{}

func (v Value) AsInterface() any { _ = "STUB: not implemented"; return *new(any) }

func (v Value) String() string { _ = "STUB: not implemented"; return "" }

func (v Value) Emit() string { _ = "STUB: not implemented"; return "" }

const (
	jsonArrayBracketsLen   = len("[]")
	boolArrayElemMaxLen    = len("false")
	int64ArrayElemMaxLen   = len("-9223372036854775808")
	float64ArrayElemMaxLen = len("-1.7976931348623157e+308")
	commaLen               = len(",")
)

func sliceValue(v []Value) any { _ = "STUB: not implemented"; return *new(any) }

func sliceValueReflect(v []Value) any { _ = "STUB: not implemented"; return *new(any) }

func mapValue(v []KeyValue) any { _ = "STUB: not implemented"; return *new(any) }

func mapValueReflect(v []KeyValue) any { _ = "STUB: not implemented"; return *new(any) }

func sortKeyValues(vals []KeyValue) { _ = "STUB: not implemented"; return }

func formatBoolSliceValue(v any) string { _ = "STUB: not implemented"; return "" }

func formatBoolSlice(vals []bool) string { _ = "STUB: not implemented"; return "" }

func formatBoolSliceReflect(v any) string { _ = "STUB: not implemented"; return "" }

func appendBoolSliceValue(dst *strings.Builder, v any) { _ = "STUB: not implemented"; return }

func appendBoolSlice(dst *strings.Builder, vals []bool) { _ = "STUB: not implemented"; return }

func appendBoolSliceReflect(dst *strings.Builder, rv reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func formatInt64SliceValue(v any) string { _ = "STUB: not implemented"; return "" }

func formatInt64Slice(vals []int64) string { _ = "STUB: not implemented"; return "" }

func formatInt64SliceReflect(v any) string { _ = "STUB: not implemented"; return "" }

func appendInt64SliceValue(dst *strings.Builder, v any) { _ = "STUB: not implemented"; return }

func appendInt64Slice(dst *strings.Builder, vals []int64) { _ = "STUB: not implemented"; return }

func appendInt64SliceReflect(dst *strings.Builder, rv reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func formatFloat64(v float64) string { _ = "STUB: not implemented"; return "" }

func formatFloat64SliceValue(v any) string { _ = "STUB: not implemented"; return "" }

func formatFloat64Slice(vals []float64) string { _ = "STUB: not implemented"; return "" }

func formatFloat64SliceReflect(v any) string { _ = "STUB: not implemented"; return "" }

func appendFloat64SliceValue(dst *strings.Builder, v any) { _ = "STUB: not implemented"; return }

func appendFloat64Slice(dst *strings.Builder, vals []float64) { _ = "STUB: not implemented"; return }

func appendFloat64SliceReflect(dst *strings.Builder, rv reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func formatStringSliceValue(v any) string { _ = "STUB: not implemented"; return "" }

func formatStringSlice(vals []string) string { _ = "STUB: not implemented"; return "" }

func formatStringSliceReflect(v any) string { _ = "STUB: not implemented"; return "" }

func appendStringSliceValue(dst *strings.Builder, v any) { _ = "STUB: not implemented"; return }

func appendStringSlice(dst *strings.Builder, vals []string) { _ = "STUB: not implemented"; return }

func appendStringSliceReflect(dst *strings.Builder, rv reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func formatByteSlice(v string) string { _ = "STUB: not implemented"; return "" }

func formatValueSliceValue(v any) string { _ = "STUB: not implemented"; return "" }

func formatValueSlice(vals []Value) string { _ = "STUB: not implemented"; return "" }

func formatValueSliceReflect(v any) string { _ = "STUB: not implemented"; return "" }

func formatMapValue(v any) string { _ = "STUB: not implemented"; return "" }

func formatMap(vals []KeyValue) string { _ = "STUB: not implemented"; return "" }

func formatMapReflect(v any) string { _ = "STUB: not implemented"; return "" }

func appendValueSliceValue(dst *strings.Builder, v any) { _ = "STUB: not implemented"; return }

func appendValueSlice(dst *strings.Builder, vals []Value) { _ = "STUB: not implemented"; return }

func appendValueSliceReflect(dst *strings.Builder, rv reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func appendMapValue(dst *strings.Builder, v any) { _ = "STUB: not implemented"; return }

func appendMap(dst *strings.Builder, vals []KeyValue) { _ = "STUB: not implemented"; return }

func appendMapReflect(dst *strings.Builder, rv reflect.Value) { _ = "STUB: not implemented"; return }

func appendJSONValue(dst *strings.Builder, v Value) { _ = "STUB: not implemented"; return }

func appendJSONString(dst *strings.Builder, s string) { _ = "STUB: not implemented"; return }

func appendBase64(dst *strings.Builder, s string) { _ = "STUB: not implemented"; return }

func (v Value) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
