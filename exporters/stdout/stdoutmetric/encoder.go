package stdoutmetric

import (
	"errors"
)

type Encoder interface {
	Encode(v any) error
}

type encoderHolder struct {
	encoder Encoder
}

func (e encoderHolder) Encode(v any) error { _ = "STUB: not implemented"; return nil }

type shutdownEncoder struct{}

var errShutdown = errors.New("exporter shutdown")

func (shutdownEncoder) Encode(any) error { _ = "STUB: not implemented"; return nil }
