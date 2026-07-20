package log

type setting[T any] struct {
	Value T
	Set   bool
}

func newSetting[T any](value T) setting[T] { _ = "STUB: not implemented"; return nil }

type resolver[T any] func(setting[T]) setting[T]

func (s setting[T]) Resolve(fn ...resolver[T]) setting[T] { _ = "STUB: not implemented"; return nil }

func clampMax[T ~int | ~int64](n T) resolver[T] { _ = "STUB: not implemented"; return nil }

func clearLessThanOne[T ~int | ~int64]() resolver[T] { _ = "STUB: not implemented"; return nil }

func getenv[T ~int | ~int64](key string) resolver[T] { _ = "STUB: not implemented"; return nil }

func fallback[T any](val T) resolver[T] { _ = "STUB: not implemented"; return nil }
