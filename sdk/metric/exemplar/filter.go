package exemplar

import (
	"context"
)

type Filter func(context.Context) bool

func TraceBasedFilter(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func AlwaysOnFilter(context.Context) bool { _ = "STUB: not implemented"; return false }

func AlwaysOffFilter(context.Context) bool { _ = "STUB: not implemented"; return false }
