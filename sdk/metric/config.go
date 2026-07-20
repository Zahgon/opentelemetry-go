package metric

import (
	"context"

	"go.opentelemetry.io/otel/sdk/metric/exemplar"
	"go.opentelemetry.io/otel/sdk/resource"
)

type config struct {
	res              *resource.Resource
	readers          []Reader
	views            []View
	exemplarFilter   exemplar.Filter
	cardinalityLimit int
}

const defaultCardinalityLimit = 2000

func (c config) readerSignals() (forceFlush, shutdown func(context.Context) error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unify(funcs []func(context.Context) error) func(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func unifyShutdown(funcs []func(context.Context) error) func(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type experimentalOption interface {
	Experimental()
}

func newConfig(options []Option) config { _ = "STUB: not implemented"; return *new(config) }

type Option interface {
	apply(config) config
}

type optionFunc func(config) config

func (o optionFunc) apply(conf config) config { _ = "STUB: not implemented"; return *new(config) }

func WithResource(res *resource.Resource) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReader(r Reader) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithView(views ...View) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithExemplarFilter(filter exemplar.Filter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCardinalityLimit(limit int) Option { _ = "STUB: not implemented"; return *new(Option) }

func meterProviderOptionsFromEnv() []Option { _ = "STUB: not implemented"; return nil }

func cardinalityLimitFromEnv() int { _ = "STUB: not implemented"; return 0 }
