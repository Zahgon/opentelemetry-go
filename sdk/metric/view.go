package metric

import (
	"errors"
)

var (
	errMultiInst = errors.New("name replacement for multiple instruments")
	errEmptyView = errors.New("no criteria provided for view")

	emptyView = func(Instrument) (Stream, bool) { return Stream{}, false }
)

type View func(Instrument) (Stream, bool)

func NewView(criteria Instrument, mask Stream) View { _ = "STUB: not implemented"; return *new(View) }

func nonZero[T comparable](v, alt T) T { _ = "STUB: not implemented"; return *new(T) }
