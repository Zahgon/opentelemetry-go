package attribute

import (
	"fmt"
)

type KeyValue struct {
	Key   Key
	Value Value
}

func (kv KeyValue) Valid() bool { _ = "STUB: not implemented"; return false }

func Bool(k string, v bool) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func BoolSlice(k string, v []bool) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func Int(k string, v int) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func IntSlice(k string, v []int) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func Int64(k string, v int64) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func Int64Slice(k string, v []int64) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func Float64(k string, v float64) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func Float64Slice(k string, v []float64) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func String(k, v string) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func StringSlice(k string, v []string) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func ByteSlice(k string, v []byte) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func Slice(k string, v ...Value) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func Map(k string, v ...KeyValue) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func Stringer(k string, v fmt.Stringer) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }
