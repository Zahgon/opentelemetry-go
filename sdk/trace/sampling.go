package trace

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Sampler interface {
	ShouldSample(parameters SamplingParameters) SamplingResult

	Description() string
}

type SamplingParameters struct {
	ParentContext context.Context
	TraceID       trace.TraceID
	Name          string
	Kind          trace.SpanKind
	Attributes    []attribute.KeyValue
	Links         []trace.Link
}

type SamplingDecision uint8

const (
	Drop SamplingDecision = iota

	RecordOnly

	RecordAndSample
)

type SamplingResult struct {
	Decision   SamplingDecision
	Attributes []attribute.KeyValue
	Tracestate trace.TraceState
}

type traceIDRatioSampler struct {
	traceIDUpperBound uint64
	description       string
}

func (ts traceIDRatioSampler) ShouldSample(p SamplingParameters) SamplingResult {
	_ = "STUB: not implemented"
	return *new(SamplingResult)
}

func (ts traceIDRatioSampler) Description() string { _ = "STUB: not implemented"; return "" }

//nolint:revive // revive complains about stutter of `trace.TraceIDRatioBased`
func TraceIDRatioBased(fraction float64) Sampler { _ = "STUB: not implemented"; return *new(Sampler) }

type alwaysOnSampler struct{}

func (alwaysOnSampler) ShouldSample(p SamplingParameters) SamplingResult {
	_ = "STUB: not implemented"
	return *new(SamplingResult)
}

func (alwaysOnSampler) Description() string { _ = "STUB: not implemented"; return "" }

func AlwaysSample() Sampler { _ = "STUB: not implemented"; return *new(Sampler) }

type alwaysOffSampler struct{}

func (alwaysOffSampler) ShouldSample(p SamplingParameters) SamplingResult {
	_ = "STUB: not implemented"
	return *new(SamplingResult)
}

func (alwaysOffSampler) Description() string { _ = "STUB: not implemented"; return "" }

func NeverSample() Sampler { _ = "STUB: not implemented"; return *new(Sampler) }

type predeterminedSampler struct {
	description string
	decision    SamplingDecision
}

func (s predeterminedSampler) ShouldSample(p SamplingParameters) SamplingResult {
	_ = "STUB: not implemented"
	return *new(SamplingResult)
}

func (s predeterminedSampler) Description() string { _ = "STUB: not implemented"; return "" }

func ParentBased(root Sampler, samplers ...ParentBasedSamplerOption) Sampler {
	_ = "STUB: not implemented"
	return *new(Sampler)
}

type parentBased struct {
	root   Sampler
	config samplerConfig
}

func configureSamplersForParentBased(samplers []ParentBasedSamplerOption) samplerConfig {
	_ = "STUB: not implemented"
	return *new(samplerConfig)
}

type samplerConfig struct {
	remoteParentSampled, remoteParentNotSampled Sampler
	localParentSampled, localParentNotSampled   Sampler
}

type ParentBasedSamplerOption interface {
	apply(samplerConfig) samplerConfig
}

func WithRemoteParentSampled(s Sampler) ParentBasedSamplerOption {
	_ = "STUB: not implemented"
	return *new(ParentBasedSamplerOption)
}

type remoteParentSampledOption struct {
	s Sampler
}

func (o remoteParentSampledOption) apply(config samplerConfig) samplerConfig {
	_ = "STUB: not implemented"
	return *new(samplerConfig)
}

func WithRemoteParentNotSampled(s Sampler) ParentBasedSamplerOption {
	_ = "STUB: not implemented"
	return *new(ParentBasedSamplerOption)
}

type remoteParentNotSampledOption struct {
	s Sampler
}

func (o remoteParentNotSampledOption) apply(config samplerConfig) samplerConfig {
	_ = "STUB: not implemented"
	return *new(samplerConfig)
}

func WithLocalParentSampled(s Sampler) ParentBasedSamplerOption {
	_ = "STUB: not implemented"
	return *new(ParentBasedSamplerOption)
}

type localParentSampledOption struct {
	s Sampler
}

func (o localParentSampledOption) apply(config samplerConfig) samplerConfig {
	_ = "STUB: not implemented"
	return *new(samplerConfig)
}

func WithLocalParentNotSampled(s Sampler) ParentBasedSamplerOption {
	_ = "STUB: not implemented"
	return *new(ParentBasedSamplerOption)
}

type localParentNotSampledOption struct {
	s Sampler
}

func (o localParentNotSampledOption) apply(config samplerConfig) samplerConfig {
	_ = "STUB: not implemented"
	return *new(samplerConfig)
}

func (pb parentBased) ShouldSample(p SamplingParameters) SamplingResult {
	_ = "STUB: not implemented"
	return *new(SamplingResult)
}

func (pb parentBased) Description() string { _ = "STUB: not implemented"; return "" }

func AlwaysRecord(root Sampler) Sampler { _ = "STUB: not implemented"; return *new(Sampler) }

type alwaysRecord struct {
	root Sampler
}

func (ar alwaysRecord) ShouldSample(p SamplingParameters) SamplingResult {
	_ = "STUB: not implemented"
	return *new(SamplingResult)
}

func (ar alwaysRecord) Description() string { _ = "STUB: not implemented"; return "" }
