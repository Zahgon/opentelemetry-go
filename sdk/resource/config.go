package resource

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
)

type config struct {
	detectors []Detector

	schemaURL string
}

type Option interface {
	apply(config) config
}

func WithAttributes(attributes ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type detectAttributes struct {
	attributes []attribute.KeyValue
}

func (d detectAttributes) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WithDetectors(detectors ...Detector) Option { _ = "STUB: not implemented"; return *new(Option) }

type detectorsOption struct {
	detectors []Detector
}

func (o detectorsOption) apply(cfg config) config { _ = "STUB: not implemented"; return *new(config) }

func WithFromEnv() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHost() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHostID() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTelemetrySDK() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSchemaURL(schemaURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

type schemaURLOption string

func (o schemaURLOption) apply(cfg config) config { _ = "STUB: not implemented"; return *new(config) }

func WithOS() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithOSType() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithOSDescription() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProcess() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProcessPID() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProcessExecutableName() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProcessExecutablePath() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProcessCommandArgs() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProcessOwner() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProcessRuntimeName() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProcessRuntimeVersion() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProcessRuntimeDescription() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithContainer() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithContainerID() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithService() Option { _ = "STUB: not implemented"; return *new(Option) }
