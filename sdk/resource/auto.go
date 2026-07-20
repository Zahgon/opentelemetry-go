package resource

import (
	"context"
	"errors"
)

var ErrPartialResource = errors.New("partial resource")

type Detector interface {
	Detect(ctx context.Context) (*Resource, error)
}

func Detect(ctx context.Context, detectors ...Detector) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func detect(ctx context.Context, res *Resource, detectors []Detector) error {
	_ = "STUB: not implemented"
	return nil
}
