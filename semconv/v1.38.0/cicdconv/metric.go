package cicdconv

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	addOptPool = &sync.Pool{New: func() any { return &[]metric.AddOption{} }}
	recOptPool = &sync.Pool{New: func() any { return &[]metric.RecordOption{} }}
)

type PipelineResultAttr string

var (
	PipelineResultSuccess PipelineResultAttr = "success"

	PipelineResultFailure PipelineResultAttr = "failure"

	PipelineResultError PipelineResultAttr = "error"

	PipelineResultTimeout PipelineResultAttr = "timeout"

	PipelineResultCancellation PipelineResultAttr = "cancellation"

	PipelineResultSkip PipelineResultAttr = "skip"
)

type PipelineRunStateAttr string

var (
	PipelineRunStatePending PipelineRunStateAttr = "pending"

	PipelineRunStateExecuting PipelineRunStateAttr = "executing"

	PipelineRunStateFinalizing PipelineRunStateAttr = "finalizing"
)

type WorkerStateAttr string

var (
	WorkerStateAvailable WorkerStateAttr = "available"

	WorkerStateBusy WorkerStateAttr = "busy"

	WorkerStateOffline WorkerStateAttr = "offline"
)

type ErrorTypeAttr string

var ErrorTypeOther ErrorTypeAttr = "_OTHER"

type PipelineRunActive struct {
	metric.Int64UpDownCounter
}

var newPipelineRunActiveOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of pipeline runs currently active in the system by state."),
	metric.WithUnit("{run}"),
}

func NewPipelineRunActive(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PipelineRunActive, error) {
	_ = "STUB: not implemented"
	return *new(PipelineRunActive), nil
}

func (m PipelineRunActive) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PipelineRunActive) Name() string { _ = "STUB: not implemented"; return "" }

func (PipelineRunActive) Unit() string { _ = "STUB: not implemented"; return "" }

func (PipelineRunActive) Description() string { _ = "STUB: not implemented"; return "" }

func (m PipelineRunActive) Add(
	ctx context.Context,
	incr int64,
	pipelineName string,
	pipelineRunState PipelineRunStateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PipelineRunActive) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PipelineRunDuration struct {
	metric.Float64Histogram
}

var newPipelineRunDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Duration of a pipeline run grouped by pipeline, state and result."),
	metric.WithUnit("s"),
}

func NewPipelineRunDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (PipelineRunDuration, error) {
	_ = "STUB: not implemented"
	return *new(PipelineRunDuration), nil
}

func (m PipelineRunDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (PipelineRunDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (PipelineRunDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (PipelineRunDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m PipelineRunDuration) Record(
	ctx context.Context,
	val float64,
	pipelineName string,
	pipelineRunState PipelineRunStateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PipelineRunDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PipelineRunDuration) AttrPipelineResult(val PipelineResultAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PipelineRunDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PipelineRunErrors struct {
	metric.Int64Counter
}

var newPipelineRunErrorsOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of errors encountered in pipeline runs (eg. compile, test failures)."),
	metric.WithUnit("{error}"),
}

func NewPipelineRunErrors(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (PipelineRunErrors, error) {
	_ = "STUB: not implemented"
	return *new(PipelineRunErrors), nil
}

func (m PipelineRunErrors) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (PipelineRunErrors) Name() string { _ = "STUB: not implemented"; return "" }

func (PipelineRunErrors) Unit() string { _ = "STUB: not implemented"; return "" }

func (PipelineRunErrors) Description() string { _ = "STUB: not implemented"; return "" }

func (m PipelineRunErrors) Add(
	ctx context.Context,
	incr int64,
	pipelineName string,
	errorType ErrorTypeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PipelineRunErrors) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type SystemErrors struct {
	metric.Int64Counter
}

var newSystemErrorsOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of errors in a component of the CICD system (eg. controller, scheduler, agent)."),
	metric.WithUnit("{error}"),
}

func NewSystemErrors(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SystemErrors, error) {
	_ = "STUB: not implemented"
	return *new(SystemErrors), nil
}

func (m SystemErrors) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (SystemErrors) Name() string { _ = "STUB: not implemented"; return "" }

func (SystemErrors) Unit() string { _ = "STUB: not implemented"; return "" }

func (SystemErrors) Description() string { _ = "STUB: not implemented"; return "" }

func (m SystemErrors) Add(
	ctx context.Context,
	incr int64,
	systemComponent string,
	errorType ErrorTypeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m SystemErrors) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type WorkerCount struct {
	metric.Int64UpDownCounter
}

var newWorkerCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of workers on the CICD system by state."),
	metric.WithUnit("{count}"),
}

func NewWorkerCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (WorkerCount, error) {
	_ = "STUB: not implemented"
	return *new(WorkerCount), nil
}

func (m WorkerCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (WorkerCount) Name() string { _ = "STUB: not implemented"; return "" }

func (WorkerCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (WorkerCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m WorkerCount) Add(
	ctx context.Context,
	incr int64,
	workerState WorkerStateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m WorkerCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}
