package attribute

type Filter func(KeyValue) bool

func NewAllowKeysFilter(keys ...Key) Filter { _ = "STUB: not implemented"; return *new(Filter) }

func NewDenyKeysFilter(keys ...Key) Filter { _ = "STUB: not implemented"; return *new(Filter) }
