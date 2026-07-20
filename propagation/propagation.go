package propagation

import (
	"context"
	"net/http"
)

type TextMapCarrier interface {
	Get(key string) string

	Set(key, value string)

	Keys() []string
}

type ValuesGetter interface {
	Values(key string) []string
}

type MapCarrier map[string]string

var _ TextMapCarrier = MapCarrier{}

func (c MapCarrier) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (c MapCarrier) Set(key, value string) { _ = "STUB: not implemented"; return }

func (c MapCarrier) Keys() []string { _ = "STUB: not implemented"; return nil }

type HeaderCarrier http.Header

var _ TextMapCarrier = HeaderCarrier{}

var _ ValuesGetter = HeaderCarrier{}

func (hc HeaderCarrier) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (hc HeaderCarrier) Values(key string) []string { _ = "STUB: not implemented"; return nil }

func (hc HeaderCarrier) Set(key, value string) { _ = "STUB: not implemented"; return }

func (hc HeaderCarrier) Keys() []string { _ = "STUB: not implemented"; return nil }

type TextMapPropagator interface {
	Inject(ctx context.Context, carrier TextMapCarrier)

	Extract(ctx context.Context, carrier TextMapCarrier) context.Context

	Fields() []string
}

type compositeTextMapPropagator []TextMapPropagator

func (p compositeTextMapPropagator) Inject(ctx context.Context, carrier TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

func (p compositeTextMapPropagator) Extract(ctx context.Context, carrier TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (p compositeTextMapPropagator) Fields() []string { _ = "STUB: not implemented"; return nil }

func NewCompositeTextMapPropagator(p ...TextMapPropagator) TextMapPropagator {
	_ = "STUB: not implemented"
	return *new(TextMapPropagator)
}
