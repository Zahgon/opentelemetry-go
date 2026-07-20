package exemplar

type ValueType uint8

const (
	UnknownValueType ValueType = 0

	Int64ValueType ValueType = 1

	Float64ValueType ValueType = 2
)

type Value struct {
	t   ValueType
	val uint64
}

func NewValue[N int64 | float64](value N) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v Value) Type() ValueType { _ = "STUB: not implemented"; return *new(ValueType) }

func (v Value) Int64() int64 { _ = "STUB: not implemented"; return 0 }

func (v Value) Float64() float64 { _ = "STUB: not implemented"; return 0 }
