package migration

import (
	"context"
)

type doDeferredContextSetupType struct{}

var (
	doDeferredContextSetupTypeKey   = doDeferredContextSetupType{}
	doDeferredContextSetupTypeValue = doDeferredContextSetupType{}
)

func WithDeferredSetup(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SkipContextSetup(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
