package logtest

import (
	"github.com/google/go-cmp/cmp"
)

type TestingT interface {
	Errorf(format string, args ...any)
}

func AssertEqual[T Recording | Record](t TestingT, want, got T, opts ...AssertOption) bool {
	_ = "STUB: not implemented"
	return false
}

type assertConfig struct {
	cmpOpts []cmp.Option
	msg     string
	args    []any
}

type AssertOption interface {
	apply(cfg assertConfig) assertConfig
}

type fnOption func(cfg assertConfig) assertConfig

func (fn fnOption) apply(cfg assertConfig) assertConfig {
	_ = "STUB: not implemented"
	return *new(assertConfig)
}

func Transform[A, B any](f func(A) B) AssertOption {
	_ = "STUB: not implemented"
	return *new(AssertOption)
}

func Desc(text string, args ...any) AssertOption {
	_ = "STUB: not implemented"
	return *new(AssertOption)
}
