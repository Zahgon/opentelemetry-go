package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/otlptranslator"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric"
)

type config struct {
	registerer               prometheus.Registerer
	disableTargetInfo        bool
	translationStrategy      otlptranslator.TranslationStrategyOption
	withoutUnits             bool
	withoutCounterSuffixes   bool
	readerOpts               []metric.ManualReaderOption
	disableScopeInfo         bool
	namespace                string
	resourceAttributesFilter attribute.Filter
}

func newConfig(opts ...Option) config { _ = "STUB: not implemented"; return *new(config) }

type Option interface {
	apply(config) config
}

type optionFunc func(config) config

func (fn optionFunc) apply(cfg config) config { _ = "STUB: not implemented"; return *new(config) }

func WithRegisterer(reg prometheus.Registerer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithAggregationSelector(agg metric.AggregationSelector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithProducer(producer metric.Producer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithoutTargetInfo() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTranslationStrategy(strategy otlptranslator.TranslationStrategyOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithoutUnits() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithoutCounterSuffixes() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithoutScopeInfo() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithNamespace(ns string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithResourceAsConstantLabels(resourceFilter attribute.Filter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
