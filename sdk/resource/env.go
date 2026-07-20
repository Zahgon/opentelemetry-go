package resource

import (
	"context"
	"fmt"
)

const (
	resourceAttrKey = "OTEL_RESOURCE_ATTRIBUTES" //nolint:gosec // False positive G101: Potential hardcoded credentials

	svcNameKey = "OTEL_SERVICE_NAME"
)

var errMissingValue = fmt.Errorf("%w: missing value", ErrPartialResource)

type fromEnv struct{}

var _ Detector = fromEnv{}

func (fromEnv) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func constructOTResources(s string) (*Resource, error) { _ = "STUB: not implemented"; return nil, nil }
