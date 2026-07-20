package k8sconv

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

type NamespacePhaseAttr string

var (
	NamespacePhaseActive NamespacePhaseAttr = "active"

	NamespacePhaseTerminating NamespacePhaseAttr = "terminating"
)

type NetworkIODirectionAttr string

var (
	NetworkIODirectionTransmit NetworkIODirectionAttr = "transmit"

	NetworkIODirectionReceive NetworkIODirectionAttr = "receive"
)

type CronJobActiveJobs struct {
	metric.Int64UpDownCounter
}

func NewCronJobActiveJobs(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (CronJobActiveJobs, error) {
	_ = "STUB: not implemented"
	return *new(CronJobActiveJobs), nil
}

func (m CronJobActiveJobs) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (CronJobActiveJobs) Name() string { _ = "STUB: not implemented"; return "" }

func (CronJobActiveJobs) Unit() string { _ = "STUB: not implemented"; return "" }

func (CronJobActiveJobs) Description() string { _ = "STUB: not implemented"; return "" }

func (m CronJobActiveJobs) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetCurrentScheduledNodes struct {
	metric.Int64UpDownCounter
}

func NewDaemonSetCurrentScheduledNodes(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DaemonSetCurrentScheduledNodes, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetCurrentScheduledNodes), nil
}

func (m DaemonSetCurrentScheduledNodes) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DaemonSetCurrentScheduledNodes) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetCurrentScheduledNodes) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetCurrentScheduledNodes) Description() string { _ = "STUB: not implemented"; return "" }

func (m DaemonSetCurrentScheduledNodes) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetDesiredScheduledNodes struct {
	metric.Int64UpDownCounter
}

func NewDaemonSetDesiredScheduledNodes(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DaemonSetDesiredScheduledNodes, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetDesiredScheduledNodes), nil
}

func (m DaemonSetDesiredScheduledNodes) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DaemonSetDesiredScheduledNodes) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetDesiredScheduledNodes) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetDesiredScheduledNodes) Description() string { _ = "STUB: not implemented"; return "" }

func (m DaemonSetDesiredScheduledNodes) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetMisscheduledNodes struct {
	metric.Int64UpDownCounter
}

func NewDaemonSetMisscheduledNodes(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DaemonSetMisscheduledNodes, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetMisscheduledNodes), nil
}

func (m DaemonSetMisscheduledNodes) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DaemonSetMisscheduledNodes) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetMisscheduledNodes) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetMisscheduledNodes) Description() string { _ = "STUB: not implemented"; return "" }

func (m DaemonSetMisscheduledNodes) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetReadyNodes struct {
	metric.Int64UpDownCounter
}

func NewDaemonSetReadyNodes(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DaemonSetReadyNodes, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetReadyNodes), nil
}

func (m DaemonSetReadyNodes) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DaemonSetReadyNodes) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetReadyNodes) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetReadyNodes) Description() string { _ = "STUB: not implemented"; return "" }

func (m DaemonSetReadyNodes) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type DeploymentAvailablePods struct {
	metric.Int64UpDownCounter
}

func NewDeploymentAvailablePods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DeploymentAvailablePods, error) {
	_ = "STUB: not implemented"
	return *new(DeploymentAvailablePods), nil
}

func (m DeploymentAvailablePods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DeploymentAvailablePods) Name() string { _ = "STUB: not implemented"; return "" }

func (DeploymentAvailablePods) Unit() string { _ = "STUB: not implemented"; return "" }

func (DeploymentAvailablePods) Description() string { _ = "STUB: not implemented"; return "" }

func (m DeploymentAvailablePods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type DeploymentDesiredPods struct {
	metric.Int64UpDownCounter
}

func NewDeploymentDesiredPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DeploymentDesiredPods, error) {
	_ = "STUB: not implemented"
	return *new(DeploymentDesiredPods), nil
}

func (m DeploymentDesiredPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DeploymentDesiredPods) Name() string { _ = "STUB: not implemented"; return "" }

func (DeploymentDesiredPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (DeploymentDesiredPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m DeploymentDesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type HPACurrentPods struct {
	metric.Int64UpDownCounter
}

func NewHPACurrentPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (HPACurrentPods, error) {
	_ = "STUB: not implemented"
	return *new(HPACurrentPods), nil
}

func (m HPACurrentPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (HPACurrentPods) Name() string { _ = "STUB: not implemented"; return "" }

func (HPACurrentPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPACurrentPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPACurrentPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type HPADesiredPods struct {
	metric.Int64UpDownCounter
}

func NewHPADesiredPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (HPADesiredPods, error) {
	_ = "STUB: not implemented"
	return *new(HPADesiredPods), nil
}

func (m HPADesiredPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (HPADesiredPods) Name() string { _ = "STUB: not implemented"; return "" }

func (HPADesiredPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPADesiredPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPADesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type HPAMaxPods struct {
	metric.Int64UpDownCounter
}

func NewHPAMaxPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (HPAMaxPods, error) {
	_ = "STUB: not implemented"
	return *new(HPAMaxPods), nil
}

func (m HPAMaxPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (HPAMaxPods) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAMaxPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAMaxPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPAMaxPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type HPAMinPods struct {
	metric.Int64UpDownCounter
}

func NewHPAMinPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (HPAMinPods, error) {
	_ = "STUB: not implemented"
	return *new(HPAMinPods), nil
}

func (m HPAMinPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (HPAMinPods) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAMinPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAMinPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPAMinPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type JobActivePods struct {
	metric.Int64UpDownCounter
}

func NewJobActivePods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobActivePods, error) {
	_ = "STUB: not implemented"
	return *new(JobActivePods), nil
}

func (m JobActivePods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobActivePods) Name() string { _ = "STUB: not implemented"; return "" }

func (JobActivePods) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobActivePods) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobActivePods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type JobDesiredSuccessfulPods struct {
	metric.Int64UpDownCounter
}

func NewJobDesiredSuccessfulPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobDesiredSuccessfulPods, error) {
	_ = "STUB: not implemented"
	return *new(JobDesiredSuccessfulPods), nil
}

func (m JobDesiredSuccessfulPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobDesiredSuccessfulPods) Name() string { _ = "STUB: not implemented"; return "" }

func (JobDesiredSuccessfulPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobDesiredSuccessfulPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobDesiredSuccessfulPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type JobFailedPods struct {
	metric.Int64UpDownCounter
}

func NewJobFailedPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobFailedPods, error) {
	_ = "STUB: not implemented"
	return *new(JobFailedPods), nil
}

func (m JobFailedPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobFailedPods) Name() string { _ = "STUB: not implemented"; return "" }

func (JobFailedPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobFailedPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobFailedPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type JobMaxParallelPods struct {
	metric.Int64UpDownCounter
}

func NewJobMaxParallelPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobMaxParallelPods, error) {
	_ = "STUB: not implemented"
	return *new(JobMaxParallelPods), nil
}

func (m JobMaxParallelPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobMaxParallelPods) Name() string { _ = "STUB: not implemented"; return "" }

func (JobMaxParallelPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobMaxParallelPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobMaxParallelPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type JobSuccessfulPods struct {
	metric.Int64UpDownCounter
}

func NewJobSuccessfulPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobSuccessfulPods, error) {
	_ = "STUB: not implemented"
	return *new(JobSuccessfulPods), nil
}

func (m JobSuccessfulPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobSuccessfulPods) Name() string { _ = "STUB: not implemented"; return "" }

func (JobSuccessfulPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobSuccessfulPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobSuccessfulPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type NamespacePhase struct {
	metric.Int64UpDownCounter
}

func NewNamespacePhase(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NamespacePhase, error) {
	_ = "STUB: not implemented"
	return *new(NamespacePhase), nil
}

func (m NamespacePhase) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NamespacePhase) Name() string { _ = "STUB: not implemented"; return "" }

func (NamespacePhase) Unit() string { _ = "STUB: not implemented"; return "" }

func (NamespacePhase) Description() string { _ = "STUB: not implemented"; return "" }

func (m NamespacePhase) Add(
	ctx context.Context,
	incr int64,
	namespacePhase NamespacePhaseAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

type NodeCPUTime struct {
	metric.Float64Counter
}

func NewNodeCPUTime(
	m metric.Meter,
	opt ...metric.Float64CounterOption,
) (NodeCPUTime, error) {
	_ = "STUB: not implemented"
	return *new(NodeCPUTime), nil
}

func (m NodeCPUTime) Inst() metric.Float64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Float64Counter)
}

func (NodeCPUTime) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeCPUTime) Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type NodeCPUUsage struct {
	metric.Int64Gauge
}

func NewNodeCPUUsage(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (NodeCPUUsage, error) {
	_ = "STUB: not implemented"
	return *new(NodeCPUUsage), nil
}

func (m NodeCPUUsage) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (NodeCPUUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeCPUUsage) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type NodeMemoryUsage struct {
	metric.Int64Gauge
}

func NewNodeMemoryUsage(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (NodeMemoryUsage, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryUsage), nil
}

func (m NodeMemoryUsage) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (NodeMemoryUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeMemoryUsage) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type NodeNetworkErrors struct {
	metric.Int64Counter
}

func NewNodeNetworkErrors(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (NodeNetworkErrors, error) {
	_ = "STUB: not implemented"
	return *new(NodeNetworkErrors), nil
}

func (m NodeNetworkErrors) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (NodeNetworkErrors) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkErrors) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkErrors) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeNetworkErrors) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (NodeNetworkErrors) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NodeNetworkErrors) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NodeNetworkIO struct {
	metric.Int64Counter
}

func NewNodeNetworkIO(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (NodeNetworkIO, error) {
	_ = "STUB: not implemented"
	return *new(NodeNetworkIO), nil
}

func (m NodeNetworkIO) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (NodeNetworkIO) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkIO) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkIO) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeNetworkIO) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (NodeNetworkIO) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NodeNetworkIO) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NodeUptime struct {
	metric.Float64Gauge
}

func NewNodeUptime(
	m metric.Meter,
	opt ...metric.Float64GaugeOption,
) (NodeUptime, error) {
	_ = "STUB: not implemented"
	return *new(NodeUptime), nil
}

func (m NodeUptime) Inst() metric.Float64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge)
}

func (NodeUptime) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeUptime) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeUptime) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeUptime) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type PodCPUTime struct {
	metric.Float64Counter
}

func NewPodCPUTime(
	m metric.Meter,
	opt ...metric.Float64CounterOption,
) (PodCPUTime, error) {
	_ = "STUB: not implemented"
	return *new(PodCPUTime), nil
}

func (m PodCPUTime) Inst() metric.Float64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Float64Counter)
}

func (PodCPUTime) Name() string { _ = "STUB: not implemented"; return "" }

func (PodCPUTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodCPUTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodCPUTime) Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type PodCPUUsage struct {
	metric.Int64Gauge
}

func NewPodCPUUsage(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (PodCPUUsage, error) {
	_ = "STUB: not implemented"
	return *new(PodCPUUsage), nil
}

func (m PodCPUUsage) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (PodCPUUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (PodCPUUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodCPUUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodCPUUsage) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type PodMemoryUsage struct {
	metric.Int64Gauge
}

func NewPodMemoryUsage(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (PodMemoryUsage, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryUsage), nil
}

func (m PodMemoryUsage) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (PodMemoryUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodMemoryUsage) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type PodNetworkErrors struct {
	metric.Int64Counter
}

func NewPodNetworkErrors(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (PodNetworkErrors, error) {
	_ = "STUB: not implemented"
	return *new(PodNetworkErrors), nil
}

func (m PodNetworkErrors) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (PodNetworkErrors) Name() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkErrors) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkErrors) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodNetworkErrors) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (PodNetworkErrors) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodNetworkErrors) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PodNetworkIO struct {
	metric.Int64Counter
}

func NewPodNetworkIO(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (PodNetworkIO, error) {
	_ = "STUB: not implemented"
	return *new(PodNetworkIO), nil
}

func (m PodNetworkIO) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (PodNetworkIO) Name() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkIO) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkIO) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodNetworkIO) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (PodNetworkIO) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodNetworkIO) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PodUptime struct {
	metric.Float64Gauge
}

func NewPodUptime(
	m metric.Meter,
	opt ...metric.Float64GaugeOption,
) (PodUptime, error) {
	_ = "STUB: not implemented"
	return *new(PodUptime), nil
}

func (m PodUptime) Inst() metric.Float64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge)
}

func (PodUptime) Name() string { _ = "STUB: not implemented"; return "" }

func (PodUptime) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodUptime) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodUptime) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ReplicaSetAvailablePods struct {
	metric.Int64UpDownCounter
}

func NewReplicaSetAvailablePods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ReplicaSetAvailablePods, error) {
	_ = "STUB: not implemented"
	return *new(ReplicaSetAvailablePods), nil
}

func (m ReplicaSetAvailablePods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ReplicaSetAvailablePods) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetAvailablePods) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetAvailablePods) Description() string { _ = "STUB: not implemented"; return "" }

func (m ReplicaSetAvailablePods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ReplicaSetDesiredPods struct {
	metric.Int64UpDownCounter
}

func NewReplicaSetDesiredPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ReplicaSetDesiredPods, error) {
	_ = "STUB: not implemented"
	return *new(ReplicaSetDesiredPods), nil
}

func (m ReplicaSetDesiredPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ReplicaSetDesiredPods) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetDesiredPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetDesiredPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m ReplicaSetDesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ReplicationControllerAvailablePods struct {
	metric.Int64UpDownCounter
}

func NewReplicationControllerAvailablePods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ReplicationControllerAvailablePods, error) {
	_ = "STUB: not implemented"
	return *new(ReplicationControllerAvailablePods), nil
}

func (m ReplicationControllerAvailablePods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ReplicationControllerAvailablePods) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicationControllerAvailablePods) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicationControllerAvailablePods) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ReplicationControllerAvailablePods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ReplicationControllerDesiredPods struct {
	metric.Int64UpDownCounter
}

func NewReplicationControllerDesiredPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ReplicationControllerDesiredPods, error) {
	_ = "STUB: not implemented"
	return *new(ReplicationControllerDesiredPods), nil
}

func (m ReplicationControllerDesiredPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ReplicationControllerDesiredPods) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicationControllerDesiredPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicationControllerDesiredPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m ReplicationControllerDesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetCurrentPods struct {
	metric.Int64UpDownCounter
}

func NewStatefulSetCurrentPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (StatefulSetCurrentPods, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetCurrentPods), nil
}

func (m StatefulSetCurrentPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (StatefulSetCurrentPods) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetCurrentPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetCurrentPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m StatefulSetCurrentPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetDesiredPods struct {
	metric.Int64UpDownCounter
}

func NewStatefulSetDesiredPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (StatefulSetDesiredPods, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetDesiredPods), nil
}

func (m StatefulSetDesiredPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (StatefulSetDesiredPods) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetDesiredPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetDesiredPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m StatefulSetDesiredPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetReadyPods struct {
	metric.Int64UpDownCounter
}

func NewStatefulSetReadyPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (StatefulSetReadyPods, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetReadyPods), nil
}

func (m StatefulSetReadyPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (StatefulSetReadyPods) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetReadyPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetReadyPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m StatefulSetReadyPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetUpdatedPods struct {
	metric.Int64UpDownCounter
}

func NewStatefulSetUpdatedPods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (StatefulSetUpdatedPods, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetUpdatedPods), nil
}

func (m StatefulSetUpdatedPods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (StatefulSetUpdatedPods) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetUpdatedPods) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetUpdatedPods) Description() string { _ = "STUB: not implemented"; return "" }

func (m StatefulSetUpdatedPods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}
