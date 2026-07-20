package stdoutmetric

import (
	"io"

	"go.opentelemetry.io/otel/sdk/metric"
)

type config struct {
	prettyPrint         bool
	encoder             *encoderHolder
	temporalitySelector metric.TemporalitySelector
	aggregationSelector metric.AggregationSelector
	redactTimestamps    bool
}

func newConfig(options ...Option) config { _ = "STUB: not implemented"; return *new(config) }

type Option interface {
	apply(config) config
}

type optionFunc func(config) config

func (o optionFunc) apply(c config) config { _ = "STUB: not implemented"; return *new(config) }

func WithEncoder(encoder Encoder) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWriter(w io.Writer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPrettyPrint() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTemporalitySelector(selector metric.TemporalitySelector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type temporalitySelectorOption struct {
	selector metric.TemporalitySelector
}

func (t temporalitySelectorOption) apply(c config) config {
	_ = "STUB: not implemented"
	return *new(config)
}

func WithAggregationSelector(selector metric.AggregationSelector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type aggregationSelectorOption struct {
	selector metric.AggregationSelector
}

func (t aggregationSelectorOption) apply(c config) config {
	_ = "STUB: not implemented"
	return *new(config)
}

func WithoutTimestamps() Option { _ = "STUB: not implemented"; return *new(Option) }
