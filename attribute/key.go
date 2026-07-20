package attribute

type Key string

func (k Key) Bool(v bool) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) BoolSlice(v []bool) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) Int(v int) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) IntSlice(v []int) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) Int64(v int64) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) Int64Slice(v []int64) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) Float64(v float64) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) Float64Slice(v []float64) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) String(v string) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) StringSlice(v []string) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) ByteSlice(v []byte) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) Slice(v ...Value) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) Map(v ...KeyValue) KeyValue { _ = "STUB: not implemented"; return *new(KeyValue) }

func (k Key) Defined() bool { _ = "STUB: not implemented"; return false }
