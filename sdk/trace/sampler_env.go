package trace

import (
	"errors"
)

const (
	tracesSamplerKey    = "OTEL_TRACES_SAMPLER"
	tracesSamplerArgKey = "OTEL_TRACES_SAMPLER_ARG"

	samplerAlwaysOn                = "always_on"
	samplerAlwaysOff               = "always_off"
	samplerTraceIDRatio            = "traceidratio"
	samplerParentBasedAlwaysOn     = "parentbased_always_on"
	samplerParsedBasedAlwaysOff    = "parentbased_always_off"
	samplerParentBasedTraceIDRatio = "parentbased_traceidratio"
)

type errUnsupportedSampler string

func (e errUnsupportedSampler) Error() string { _ = "STUB: not implemented"; return "" }

var (
	errNegativeTraceIDRatio       = errors.New("invalid trace ID ratio: less than 0.0")
	errGreaterThanOneTraceIDRatio = errors.New("invalid trace ID ratio: greater than 1.0")
)

type samplerArgParseError struct {
	parseErr error
}

func (e samplerArgParseError) Error() string { _ = "STUB: not implemented"; return "" }

func (e samplerArgParseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func samplerFromEnv() (Sampler, error) { _ = "STUB: not implemented"; return *new(Sampler), nil }

func parseTraceIDRatio(arg string) (Sampler, error) {
	_ = "STUB: not implemented"
	return *new(Sampler), nil
}
