package resource

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
)

type osDescriptionProvider func() (string, error)

var defaultOSDescriptionProvider osDescriptionProvider = platformOSDescription

var osDescription = defaultOSDescriptionProvider

func setDefaultOSDescriptionProvider() { _ = "STUB: not implemented"; return }

func setOSDescriptionProvider(osDescriptionProvider osDescriptionProvider) {
	_ = "STUB: not implemented"
	return
}

type (
	osTypeDetector        struct{}
	osDescriptionDetector struct{}
)

func (osTypeDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (osDescriptionDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mapRuntimeOSToSemconvOSType(osType string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
