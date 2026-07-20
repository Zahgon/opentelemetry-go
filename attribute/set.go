package attribute

import (
	"reflect"

	"go.opentelemetry.io/otel/attribute/internal/xxhash"
)

type (
	Set struct {
		hash uint64
		data any
	}

	Distinct struct {
		hash uint64
	}

	Sortable []KeyValue
)

var (
	_ = isComparable(Set{})
	_ = isComparable(Distinct{})
)

func isComparable[T comparable](t T) T { _ = "STUB: not implemented"; return *new(T) }

var (
	keyValueType = reflect.TypeFor[KeyValue]()

	emptyHash = xxhash.New().Sum64()

	userDefinedEmptySet = &Set{
		hash: emptyHash,
		data: [0]KeyValue{},
	}

	emptySet = Set{
		hash: emptyHash,
		data: [0]KeyValue{},
	}
)

func EmptySet() *Set { _ = "STUB: not implemented"; return nil }

func (d Distinct) Valid() bool { _ = "STUB: not implemented"; return false }

func (l Set) reflectValue() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (l *Set) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *Set) Get(idx int) (KeyValue, bool) {
	_ = "STUB: not implemented"
	return *new(KeyValue), false
}

func (l *Set) Value(k Key) (Value, bool) { _ = "STUB: not implemented"; return *new(Value), false }

func (l *Set) HasValue(k Key) bool { _ = "STUB: not implemented"; return false }

func (l *Set) Iter() Iterator { _ = "STUB: not implemented"; return *new(Iterator) }

func (l *Set) ToSlice() []KeyValue { _ = "STUB: not implemented"; return nil }

func (l *Set) Equivalent() Distinct { _ = "STUB: not implemented"; return *new(Distinct) }

func (l *Set) Equals(o *Set) bool { _ = "STUB: not implemented"; return false }

func (l *Set) Encoded(encoder Encoder) string { _ = "STUB: not implemented"; return "" }

func NewSet(kvs ...KeyValue) Set { _ = "STUB: not implemented"; return *new(Set) }

func NewSetWithSortable(kvs []KeyValue, _ *Sortable) Set {
	_ = "STUB: not implemented"
	return *new(Set)
}

func NewSetWithFiltered(kvs []KeyValue, filter Filter) (Set, []KeyValue) {
	_ = "STUB: not implemented"
	return *new(Set), nil
}

func NewSetWithSortableFiltered(kvs []KeyValue, _ *Sortable, filter Filter) (Set, []KeyValue) {
	_ = "STUB: not implemented"
	return *new(Set), nil
}

func filteredToFront(slice []KeyValue, keep Filter) int { _ = "STUB: not implemented"; return 0 }

func (l *Set) Filter(re Filter) (Set, []KeyValue) { _ = "STUB: not implemented"; return *new(Set), nil }

func newSet(kvs []KeyValue) Set { _ = "STUB: not implemented"; return *new(Set) }

func computeDataFixed(kvs []KeyValue) any { _ = "STUB: not implemented"; return *new(any) }

func computeDataReflect(kvs []KeyValue) any { _ = "STUB: not implemented"; return *new(any) }

func (l *Set) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (l Set) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }

func (l *Sortable) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *Sortable) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (l *Sortable) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
