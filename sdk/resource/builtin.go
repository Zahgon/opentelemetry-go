package resource

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
)

type (
	telemetrySDK struct{}

	host struct{}

	stringDetector struct {
		schemaURL string
		K         attribute.Key
		F         func() (string, error)
	}

	defaultServiceNameDetector struct{}

	defaultServiceInstanceIDDetector struct{}
)

var (
	_ Detector = telemetrySDK{}
	_ Detector = host{}
	_ Detector = stringDetector{}
	_ Detector = defaultServiceNameDetector{}
	_ Detector = defaultServiceInstanceIDDetector{}
)

func (telemetrySDK) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (host) Detect(ctx context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func StringDetector(schemaURL string, k attribute.Key, f func() (string, error)) Detector {
	_ = "STUB: not implemented"
	return *new(Detector)
}

func (sd stringDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (defaultServiceNameDetector) Detect(ctx context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (defaultServiceInstanceIDDetector) Detect(ctx context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
