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

type ContainerStatusReasonAttr string

var (
	ContainerStatusReasonContainerCreating ContainerStatusReasonAttr = "ContainerCreating"

	ContainerStatusReasonCrashLoopBackOff ContainerStatusReasonAttr = "CrashLoopBackOff"

	ContainerStatusReasonCreateContainerConfigError ContainerStatusReasonAttr = "CreateContainerConfigError"

	ContainerStatusReasonErrImagePull ContainerStatusReasonAttr = "ErrImagePull"

	ContainerStatusReasonImagePullBackOff ContainerStatusReasonAttr = "ImagePullBackOff"

	ContainerStatusReasonOomKilled ContainerStatusReasonAttr = "OOMKilled"

	ContainerStatusReasonCompleted ContainerStatusReasonAttr = "Completed"

	ContainerStatusReasonError ContainerStatusReasonAttr = "Error"

	ContainerStatusReasonContainerCannotRun ContainerStatusReasonAttr = "ContainerCannotRun"
)

type ContainerStatusStateAttr string

var (
	ContainerStatusStateTerminated ContainerStatusStateAttr = "terminated"

	ContainerStatusStateRunning ContainerStatusStateAttr = "running"

	ContainerStatusStateWaiting ContainerStatusStateAttr = "waiting"
)

type NamespacePhaseAttr string

var (
	NamespacePhaseActive NamespacePhaseAttr = "active"

	NamespacePhaseTerminating NamespacePhaseAttr = "terminating"
)

type NodeConditionStatusAttr string

var (
	NodeConditionStatusConditionTrue NodeConditionStatusAttr = "true"

	NodeConditionStatusConditionFalse NodeConditionStatusAttr = "false"

	NodeConditionStatusConditionUnknown NodeConditionStatusAttr = "unknown"
)

type NodeConditionTypeAttr string

var (
	NodeConditionTypeReady NodeConditionTypeAttr = "Ready"

	NodeConditionTypeDiskPressure NodeConditionTypeAttr = "DiskPressure"

	NodeConditionTypeMemoryPressure NodeConditionTypeAttr = "MemoryPressure"

	NodeConditionTypePIDPressure NodeConditionTypeAttr = "PIDPressure"

	NodeConditionTypeNetworkUnavailable NodeConditionTypeAttr = "NetworkUnavailable"
)

type VolumeTypeAttr string

var (
	VolumeTypePersistentVolumeClaim VolumeTypeAttr = "persistentVolumeClaim"

	VolumeTypeConfigMap VolumeTypeAttr = "configMap"

	VolumeTypeDownwardAPI VolumeTypeAttr = "downwardAPI"

	VolumeTypeEmptyDir VolumeTypeAttr = "emptyDir"

	VolumeTypeSecret VolumeTypeAttr = "secret"

	VolumeTypeLocal VolumeTypeAttr = "local"
)

type NetworkIODirectionAttr string

var (
	NetworkIODirectionTransmit NetworkIODirectionAttr = "transmit"

	NetworkIODirectionReceive NetworkIODirectionAttr = "receive"
)

type ContainerCPULimit struct {
	metric.Int64UpDownCounter
}

var newContainerCPULimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Maximum CPU resource limit set for the container."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPULimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerCPULimit, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPULimit), nil
}

func (m ContainerCPULimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerCPULimit) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerCPULimit) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerCPULimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerCPURequest struct {
	metric.Int64UpDownCounter
}

var newContainerCPURequestOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("CPU resource requested for the container."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPURequest(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerCPURequest, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPURequest), nil
}

func (m ContainerCPURequest) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerCPURequest) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequest) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequest) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerCPURequest) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerCPURequest) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerEphemeralStorageLimit struct {
	metric.Int64UpDownCounter
}

var newContainerEphemeralStorageLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Maximum ephemeral storage resource limit set for the container."),
	metric.WithUnit("By"),
}

func NewContainerEphemeralStorageLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerEphemeralStorageLimit, error) {
	_ = "STUB: not implemented"
	return *new(ContainerEphemeralStorageLimit), nil
}

func (m ContainerEphemeralStorageLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerEphemeralStorageLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerEphemeralStorageLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerEphemeralStorageLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerEphemeralStorageLimit) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerEphemeralStorageLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerEphemeralStorageRequest struct {
	metric.Int64UpDownCounter
}

var newContainerEphemeralStorageRequestOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Ephemeral storage resource requested for the container."),
	metric.WithUnit("By"),
}

func NewContainerEphemeralStorageRequest(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerEphemeralStorageRequest, error) {
	_ = "STUB: not implemented"
	return *new(ContainerEphemeralStorageRequest), nil
}

func (m ContainerEphemeralStorageRequest) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerEphemeralStorageRequest) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerEphemeralStorageRequest) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerEphemeralStorageRequest) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerEphemeralStorageRequest) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerEphemeralStorageRequest) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerMemoryLimit struct {
	metric.Int64UpDownCounter
}

var newContainerMemoryLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Maximum memory resource limit set for the container."),
	metric.WithUnit("By"),
}

func NewContainerMemoryLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerMemoryLimit, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryLimit), nil
}

func (m ContainerMemoryLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerMemoryLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerMemoryLimit) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerMemoryLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerMemoryRequest struct {
	metric.Int64UpDownCounter
}

var newContainerMemoryRequestOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Memory resource requested for the container."),
	metric.WithUnit("By"),
}

func NewContainerMemoryRequest(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerMemoryRequest, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryRequest), nil
}

func (m ContainerMemoryRequest) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerMemoryRequest) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequest) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequest) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerMemoryRequest) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerMemoryRequest) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerReady struct {
	metric.Int64UpDownCounter
}

var newContainerReadyOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Indicates whether the container is currently marked as ready to accept traffic, based on its readiness probe (1 = ready, 0 = not ready)."),
	metric.WithUnit("{container}"),
}

func NewContainerReady(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerReady, error) {
	_ = "STUB: not implemented"
	return *new(ContainerReady), nil
}

func (m ContainerReady) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerReady) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerReady) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerReady) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerReady) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerReady) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerRestartCount struct {
	metric.Int64UpDownCounter
}

var newContainerRestartCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Describes how many times the container has restarted (since the last counter reset)."),
	metric.WithUnit("{restart}"),
}

func NewContainerRestartCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerRestartCount, error) {
	_ = "STUB: not implemented"
	return *new(ContainerRestartCount), nil
}

func (m ContainerRestartCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerRestartCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerRestartCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerRestartCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerRestartCount) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerRestartCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerStatusReason struct {
	metric.Int64UpDownCounter
}

var newContainerStatusReasonOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Describes the number of K8s containers that are currently in a state for a given reason."),
	metric.WithUnit("{container}"),
}

func NewContainerStatusReason(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerStatusReason, error) {
	_ = "STUB: not implemented"
	return *new(ContainerStatusReason), nil
}

func (m ContainerStatusReason) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerStatusReason) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusReason) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusReason) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerStatusReason) Add(
	ctx context.Context,
	incr int64,
	containerStatusReason ContainerStatusReasonAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerStatusReason) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerStatusState struct {
	metric.Int64UpDownCounter
}

var newContainerStatusStateOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Describes the number of K8s containers that are currently in a given state."),
	metric.WithUnit("{container}"),
}

func NewContainerStatusState(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerStatusState, error) {
	_ = "STUB: not implemented"
	return *new(ContainerStatusState), nil
}

func (m ContainerStatusState) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerStatusState) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusState) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusState) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerStatusState) Add(
	ctx context.Context,
	incr int64,
	containerStatusState ContainerStatusStateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerStatusState) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerStorageLimit struct {
	metric.Int64UpDownCounter
}

var newContainerStorageLimitOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Maximum storage resource limit set for the container."),
	metric.WithUnit("By"),
}

func NewContainerStorageLimit(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerStorageLimit, error) {
	_ = "STUB: not implemented"
	return *new(ContainerStorageLimit), nil
}

func (m ContainerStorageLimit) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerStorageLimit) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerStorageLimit) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerStorageLimit) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerStorageLimit) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerStorageLimit) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerStorageRequest struct {
	metric.Int64UpDownCounter
}

var newContainerStorageRequestOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Storage resource requested for the container."),
	metric.WithUnit("By"),
}

func NewContainerStorageRequest(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerStorageRequest, error) {
	_ = "STUB: not implemented"
	return *new(ContainerStorageRequest), nil
}

func (m ContainerStorageRequest) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerStorageRequest) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerStorageRequest) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerStorageRequest) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerStorageRequest) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerStorageRequest) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type CronJobActiveJobs struct {
	metric.Int64UpDownCounter
}

var newCronJobActiveJobsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of actively running jobs for a cronjob."),
	metric.WithUnit("{job}"),
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

func (m CronJobActiveJobs) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetCurrentScheduledNodes struct {
	metric.Int64UpDownCounter
}

var newDaemonSetCurrentScheduledNodesOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod."),
	metric.WithUnit("{node}"),
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

func (m DaemonSetCurrentScheduledNodes) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetDesiredScheduledNodes struct {
	metric.Int64UpDownCounter
}

var newDaemonSetDesiredScheduledNodesOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)."),
	metric.WithUnit("{node}"),
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

func (m DaemonSetDesiredScheduledNodes) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetMisscheduledNodes struct {
	metric.Int64UpDownCounter
}

var newDaemonSetMisscheduledNodesOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod."),
	metric.WithUnit("{node}"),
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

func (m DaemonSetMisscheduledNodes) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetReadyNodes struct {
	metric.Int64UpDownCounter
}

var newDaemonSetReadyNodesOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready."),
	metric.WithUnit("{node}"),
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

func (m DaemonSetReadyNodes) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DeploymentAvailablePods struct {
	metric.Int64UpDownCounter
}

var newDeploymentAvailablePodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this deployment."),
	metric.WithUnit("{pod}"),
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

func (m DeploymentAvailablePods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DeploymentDesiredPods struct {
	metric.Int64UpDownCounter
}

var newDeploymentDesiredPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this deployment."),
	metric.WithUnit("{pod}"),
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

func (m DeploymentDesiredPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type HPACurrentPods struct {
	metric.Int64UpDownCounter
}

var newHPACurrentPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Current number of replica pods managed by this horizontal pod autoscaler, as last seen by the autoscaler."),
	metric.WithUnit("{pod}"),
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

func (m HPACurrentPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type HPADesiredPods struct {
	metric.Int64UpDownCounter
}

var newHPADesiredPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Desired number of replica pods managed by this horizontal pod autoscaler, as last calculated by the autoscaler."),
	metric.WithUnit("{pod}"),
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

func (m HPADesiredPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type HPAMaxPods struct {
	metric.Int64UpDownCounter
}

var newHPAMaxPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The upper limit for the number of replica pods to which the autoscaler can scale up."),
	metric.WithUnit("{pod}"),
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

func (m HPAMaxPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type HPAMetricTargetCPUAverageUtilization struct {
	metric.Int64Gauge
}

var newHPAMetricTargetCPUAverageUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Target average utilization, in percentage, for CPU resource in HPA config."),
	metric.WithUnit("1"),
}

func NewHPAMetricTargetCPUAverageUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (HPAMetricTargetCPUAverageUtilization, error) {
	_ = "STUB: not implemented"
	return *new(HPAMetricTargetCPUAverageUtilization), nil
}

func (m HPAMetricTargetCPUAverageUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (HPAMetricTargetCPUAverageUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUAverageUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUAverageUtilization) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m HPAMetricTargetCPUAverageUtilization) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m HPAMetricTargetCPUAverageUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (HPAMetricTargetCPUAverageUtilization) AttrContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HPAMetricTargetCPUAverageUtilization) AttrHPAMetricType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type HPAMetricTargetCPUAverageValue struct {
	metric.Int64Gauge
}

var newHPAMetricTargetCPUAverageValueOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Target average value for CPU resource in HPA config."),
	metric.WithUnit("{cpu}"),
}

func NewHPAMetricTargetCPUAverageValue(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (HPAMetricTargetCPUAverageValue, error) {
	_ = "STUB: not implemented"
	return *new(HPAMetricTargetCPUAverageValue), nil
}

func (m HPAMetricTargetCPUAverageValue) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (HPAMetricTargetCPUAverageValue) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUAverageValue) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUAverageValue) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPAMetricTargetCPUAverageValue) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m HPAMetricTargetCPUAverageValue) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (HPAMetricTargetCPUAverageValue) AttrContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HPAMetricTargetCPUAverageValue) AttrHPAMetricType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type HPAMetricTargetCPUValue struct {
	metric.Int64Gauge
}

var newHPAMetricTargetCPUValueOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Target value for CPU resource in HPA config."),
	metric.WithUnit("{cpu}"),
}

func NewHPAMetricTargetCPUValue(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (HPAMetricTargetCPUValue, error) {
	_ = "STUB: not implemented"
	return *new(HPAMetricTargetCPUValue), nil
}

func (m HPAMetricTargetCPUValue) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (HPAMetricTargetCPUValue) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUValue) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUValue) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPAMetricTargetCPUValue) Record(
	ctx context.Context,
	val int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m HPAMetricTargetCPUValue) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (HPAMetricTargetCPUValue) AttrContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HPAMetricTargetCPUValue) AttrHPAMetricType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type HPAMinPods struct {
	metric.Int64UpDownCounter
}

var newHPAMinPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The lower limit for the number of replica pods to which the autoscaler can scale down."),
	metric.WithUnit("{pod}"),
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

func (m HPAMinPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobActivePods struct {
	metric.Int64UpDownCounter
}

var newJobActivePodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of pending and actively running pods for a job."),
	metric.WithUnit("{pod}"),
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

func (m JobActivePods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobDesiredSuccessfulPods struct {
	metric.Int64UpDownCounter
}

var newJobDesiredSuccessfulPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The desired number of successfully finished pods the job should be run with."),
	metric.WithUnit("{pod}"),
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

func (m JobDesiredSuccessfulPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobFailedPods struct {
	metric.Int64UpDownCounter
}

var newJobFailedPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of pods which reached phase Failed for a job."),
	metric.WithUnit("{pod}"),
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

func (m JobFailedPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobMaxParallelPods struct {
	metric.Int64UpDownCounter
}

var newJobMaxParallelPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The max desired number of pods the job should run at any given time."),
	metric.WithUnit("{pod}"),
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

func (m JobMaxParallelPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobSuccessfulPods struct {
	metric.Int64UpDownCounter
}

var newJobSuccessfulPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of pods which reached phase Succeeded for a job."),
	metric.WithUnit("{pod}"),
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

func (m JobSuccessfulPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NamespacePhase struct {
	metric.Int64UpDownCounter
}

var newNamespacePhaseOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Describes number of K8s namespaces that are currently in a given phase."),
	metric.WithUnit("{namespace}"),
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

func (m NamespacePhase) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeAllocatableCPU struct {
	metric.Int64UpDownCounter
}

var newNodeAllocatableCPUOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Amount of cpu allocatable on the node."),
	metric.WithUnit("{cpu}"),
}

func NewNodeAllocatableCPU(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeAllocatableCPU, error) {
	_ = "STUB: not implemented"
	return *new(NodeAllocatableCPU), nil
}

func (m NodeAllocatableCPU) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeAllocatableCPU) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeAllocatableCPU) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeAllocatableCPU) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeAllocatableCPU) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeAllocatableCPU) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeAllocatableEphemeralStorage struct {
	metric.Int64UpDownCounter
}

var newNodeAllocatableEphemeralStorageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Amount of ephemeral-storage allocatable on the node."),
	metric.WithUnit("By"),
}

func NewNodeAllocatableEphemeralStorage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeAllocatableEphemeralStorage, error) {
	_ = "STUB: not implemented"
	return *new(NodeAllocatableEphemeralStorage), nil
}

func (m NodeAllocatableEphemeralStorage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeAllocatableEphemeralStorage) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeAllocatableEphemeralStorage) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeAllocatableEphemeralStorage) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeAllocatableEphemeralStorage) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeAllocatableEphemeralStorage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeAllocatableMemory struct {
	metric.Int64UpDownCounter
}

var newNodeAllocatableMemoryOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Amount of memory allocatable on the node."),
	metric.WithUnit("By"),
}

func NewNodeAllocatableMemory(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeAllocatableMemory, error) {
	_ = "STUB: not implemented"
	return *new(NodeAllocatableMemory), nil
}

func (m NodeAllocatableMemory) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeAllocatableMemory) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeAllocatableMemory) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeAllocatableMemory) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeAllocatableMemory) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeAllocatableMemory) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeAllocatablePods struct {
	metric.Int64UpDownCounter
}

var newNodeAllocatablePodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Amount of pods allocatable on the node."),
	metric.WithUnit("{pod}"),
}

func NewNodeAllocatablePods(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeAllocatablePods, error) {
	_ = "STUB: not implemented"
	return *new(NodeAllocatablePods), nil
}

func (m NodeAllocatablePods) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeAllocatablePods) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeAllocatablePods) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeAllocatablePods) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeAllocatablePods) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeAllocatablePods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeConditionStatus struct {
	metric.Int64UpDownCounter
}

var newNodeConditionStatusOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Describes the condition of a particular Node."),
	metric.WithUnit("{node}"),
}

func NewNodeConditionStatus(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeConditionStatus, error) {
	_ = "STUB: not implemented"
	return *new(NodeConditionStatus), nil
}

func (m NodeConditionStatus) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeConditionStatus) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeConditionStatus) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeConditionStatus) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeConditionStatus) Add(
	ctx context.Context,
	incr int64,
	nodeConditionStatus NodeConditionStatusAttr,
	nodeConditionType NodeConditionTypeAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NodeConditionStatus) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeCPUTime struct {
	metric.Float64Counter
}

var newNodeCPUTimeOpts = []metric.Float64CounterOption{
	metric.WithDescription("Total CPU time consumed."),
	metric.WithUnit("s"),
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

func (m NodeCPUTime) AddSet(ctx context.Context, incr float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeCPUUsage struct {
	metric.Int64Gauge
}

var newNodeCPUUsageOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Node's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs."),
	metric.WithUnit("{cpu}"),
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

func (m NodeCPUUsage) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeFilesystemAvailable struct {
	metric.Int64UpDownCounter
}

var newNodeFilesystemAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Node filesystem available bytes."),
	metric.WithUnit("By"),
}

func NewNodeFilesystemAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeFilesystemAvailable, error) {
	_ = "STUB: not implemented"
	return *new(NodeFilesystemAvailable), nil
}

func (m NodeFilesystemAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeFilesystemAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeFilesystemAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeFilesystemAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeFilesystemCapacity struct {
	metric.Int64UpDownCounter
}

var newNodeFilesystemCapacityOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Node filesystem capacity."),
	metric.WithUnit("By"),
}

func NewNodeFilesystemCapacity(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeFilesystemCapacity, error) {
	_ = "STUB: not implemented"
	return *new(NodeFilesystemCapacity), nil
}

func (m NodeFilesystemCapacity) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeFilesystemCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemCapacity) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeFilesystemCapacity) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeFilesystemCapacity) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeFilesystemUsage struct {
	metric.Int64UpDownCounter
}

var newNodeFilesystemUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Node filesystem usage."),
	metric.WithUnit("By"),
}

func NewNodeFilesystemUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeFilesystemUsage, error) {
	_ = "STUB: not implemented"
	return *new(NodeFilesystemUsage), nil
}

func (m NodeFilesystemUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeFilesystemUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeFilesystemUsage) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeFilesystemUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeMemoryUsage struct {
	metric.Int64Gauge
}

var newNodeMemoryUsageOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Memory usage of the Node."),
	metric.WithUnit("By"),
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

func (m NodeMemoryUsage) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeNetworkErrors struct {
	metric.Int64Counter
}

var newNodeNetworkErrorsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Node network errors."),
	metric.WithUnit("{error}"),
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

func (m NodeNetworkErrors) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

var newNodeNetworkIOOpts = []metric.Int64CounterOption{
	metric.WithDescription("Network bytes for the Node."),
	metric.WithUnit("By"),
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

func (m NodeNetworkIO) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

var newNodeUptimeOpts = []metric.Float64GaugeOption{
	metric.WithDescription("The time the Node has been running."),
	metric.WithUnit("s"),
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

func (m NodeUptime) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodCPUTime struct {
	metric.Float64Counter
}

var newPodCPUTimeOpts = []metric.Float64CounterOption{
	metric.WithDescription("Total CPU time consumed."),
	metric.WithUnit("s"),
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

func (m PodCPUTime) AddSet(ctx context.Context, incr float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodCPUUsage struct {
	metric.Int64Gauge
}

var newPodCPUUsageOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Pod's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs."),
	metric.WithUnit("{cpu}"),
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

func (m PodCPUUsage) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodFilesystemAvailable struct {
	metric.Int64UpDownCounter
}

var newPodFilesystemAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Pod filesystem available bytes."),
	metric.WithUnit("By"),
}

func NewPodFilesystemAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodFilesystemAvailable, error) {
	_ = "STUB: not implemented"
	return *new(PodFilesystemAvailable), nil
}

func (m PodFilesystemAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodFilesystemAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodFilesystemAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m PodFilesystemAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodFilesystemCapacity struct {
	metric.Int64UpDownCounter
}

var newPodFilesystemCapacityOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Pod filesystem capacity."),
	metric.WithUnit("By"),
}

func NewPodFilesystemCapacity(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodFilesystemCapacity, error) {
	_ = "STUB: not implemented"
	return *new(PodFilesystemCapacity), nil
}

func (m PodFilesystemCapacity) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodFilesystemCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemCapacity) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodFilesystemCapacity) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m PodFilesystemCapacity) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodFilesystemUsage struct {
	metric.Int64UpDownCounter
}

var newPodFilesystemUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Pod filesystem usage."),
	metric.WithUnit("By"),
}

func NewPodFilesystemUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodFilesystemUsage, error) {
	_ = "STUB: not implemented"
	return *new(PodFilesystemUsage), nil
}

func (m PodFilesystemUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodFilesystemUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodFilesystemUsage) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m PodFilesystemUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodMemoryUsage struct {
	metric.Int64Gauge
}

var newPodMemoryUsageOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Memory usage of the Pod."),
	metric.WithUnit("By"),
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

func (m PodMemoryUsage) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodNetworkErrors struct {
	metric.Int64Counter
}

var newPodNetworkErrorsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Pod network errors."),
	metric.WithUnit("{error}"),
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

func (m PodNetworkErrors) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

var newPodNetworkIOOpts = []metric.Int64CounterOption{
	metric.WithDescription("Network bytes for the Pod."),
	metric.WithUnit("By"),
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

func (m PodNetworkIO) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

var newPodUptimeOpts = []metric.Float64GaugeOption{
	metric.WithDescription("The time the Pod has been running."),
	metric.WithUnit("s"),
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

func (m PodUptime) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodVolumeAvailable struct {
	metric.Int64UpDownCounter
}

var newPodVolumeAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Pod volume storage space available."),
	metric.WithUnit("By"),
}

func NewPodVolumeAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodVolumeAvailable, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeAvailable), nil
}

func (m PodVolumeAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodVolumeAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodVolumeAvailable) Add(
	ctx context.Context,
	incr int64,
	volumeName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PodVolumeAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PodVolumeAvailable) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PodVolumeCapacity struct {
	metric.Int64UpDownCounter
}

var newPodVolumeCapacityOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Pod volume total capacity."),
	metric.WithUnit("By"),
}

func NewPodVolumeCapacity(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodVolumeCapacity, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeCapacity), nil
}

func (m PodVolumeCapacity) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodVolumeCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeCapacity) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodVolumeCapacity) Add(
	ctx context.Context,
	incr int64,
	volumeName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PodVolumeCapacity) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PodVolumeCapacity) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PodVolumeInodeCount struct {
	metric.Int64UpDownCounter
}

var newPodVolumeInodeCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The total inodes in the filesystem of the Pod's volume."),
	metric.WithUnit("{inode}"),
}

func NewPodVolumeInodeCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodVolumeInodeCount, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeInodeCount), nil
}

func (m PodVolumeInodeCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodVolumeInodeCount) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodVolumeInodeCount) Add(
	ctx context.Context,
	incr int64,
	volumeName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PodVolumeInodeCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PodVolumeInodeCount) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PodVolumeInodeFree struct {
	metric.Int64UpDownCounter
}

var newPodVolumeInodeFreeOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The free inodes in the filesystem of the Pod's volume."),
	metric.WithUnit("{inode}"),
}

func NewPodVolumeInodeFree(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodVolumeInodeFree, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeInodeFree), nil
}

func (m PodVolumeInodeFree) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodVolumeInodeFree) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeFree) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeFree) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodVolumeInodeFree) Add(
	ctx context.Context,
	incr int64,
	volumeName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PodVolumeInodeFree) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PodVolumeInodeFree) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PodVolumeInodeUsed struct {
	metric.Int64UpDownCounter
}

var newPodVolumeInodeUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The inodes used by the filesystem of the Pod's volume."),
	metric.WithUnit("{inode}"),
}

func NewPodVolumeInodeUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodVolumeInodeUsed, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeInodeUsed), nil
}

func (m PodVolumeInodeUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodVolumeInodeUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodVolumeInodeUsed) Add(
	ctx context.Context,
	incr int64,
	volumeName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PodVolumeInodeUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PodVolumeInodeUsed) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PodVolumeUsage struct {
	metric.Int64UpDownCounter
}

var newPodVolumeUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Pod volume usage."),
	metric.WithUnit("By"),
}

func NewPodVolumeUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodVolumeUsage, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeUsage), nil
}

func (m PodVolumeUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodVolumeUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodVolumeUsage) Add(
	ctx context.Context,
	incr int64,
	volumeName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PodVolumeUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PodVolumeUsage) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ReplicaSetAvailablePods struct {
	metric.Int64UpDownCounter
}

var newReplicaSetAvailablePodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this replicaset."),
	metric.WithUnit("{pod}"),
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

func (m ReplicaSetAvailablePods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ReplicaSetDesiredPods struct {
	metric.Int64UpDownCounter
}

var newReplicaSetDesiredPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this replicaset."),
	metric.WithUnit("{pod}"),
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

func (m ReplicaSetDesiredPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ReplicationControllerAvailablePods struct {
	metric.Int64UpDownCounter
}

var newReplicationControllerAvailablePodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this replication controller."),
	metric.WithUnit("{pod}"),
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

func (m ReplicationControllerAvailablePods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ReplicationControllerDesiredPods struct {
	metric.Int64UpDownCounter
}

var newReplicationControllerDesiredPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this replication controller."),
	metric.WithUnit("{pod}"),
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

func (m ReplicationControllerDesiredPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaCPULimitHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaCPULimitHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The CPU limits in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{cpu}"),
}

func NewResourceQuotaCPULimitHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaCPULimitHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaCPULimitHard), nil
}

func (m ResourceQuotaCPULimitHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaCPULimitHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPULimitHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPULimitHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaCPULimitHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaCPULimitHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaCPULimitUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaCPULimitUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The CPU limits in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{cpu}"),
}

func NewResourceQuotaCPULimitUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaCPULimitUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaCPULimitUsed), nil
}

func (m ResourceQuotaCPULimitUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaCPULimitUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPULimitUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPULimitUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaCPULimitUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaCPULimitUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaCPURequestHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaCPURequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The CPU requests in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{cpu}"),
}

func NewResourceQuotaCPURequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaCPURequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaCPURequestHard), nil
}

func (m ResourceQuotaCPURequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaCPURequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPURequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPURequestHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaCPURequestHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaCPURequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaCPURequestUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaCPURequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The CPU requests in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{cpu}"),
}

func NewResourceQuotaCPURequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaCPURequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaCPURequestUsed), nil
}

func (m ResourceQuotaCPURequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaCPURequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPURequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPURequestUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaCPURequestUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaCPURequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaEphemeralStorageLimitHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaEphemeralStorageLimitHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The sum of local ephemeral storage limits in the namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaEphemeralStorageLimitHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaEphemeralStorageLimitHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaEphemeralStorageLimitHard), nil
}

func (m ResourceQuotaEphemeralStorageLimitHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaEphemeralStorageLimitHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaEphemeralStorageLimitHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaEphemeralStorageLimitHard) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ResourceQuotaEphemeralStorageLimitHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaEphemeralStorageLimitHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaEphemeralStorageLimitUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaEphemeralStorageLimitUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The sum of local ephemeral storage limits in the namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaEphemeralStorageLimitUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaEphemeralStorageLimitUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaEphemeralStorageLimitUsed), nil
}

func (m ResourceQuotaEphemeralStorageLimitUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaEphemeralStorageLimitUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaEphemeralStorageLimitUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaEphemeralStorageLimitUsed) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ResourceQuotaEphemeralStorageLimitUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaEphemeralStorageLimitUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaEphemeralStorageRequestHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaEphemeralStorageRequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The sum of local ephemeral storage requests in the namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaEphemeralStorageRequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaEphemeralStorageRequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaEphemeralStorageRequestHard), nil
}

func (m ResourceQuotaEphemeralStorageRequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaEphemeralStorageRequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaEphemeralStorageRequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaEphemeralStorageRequestHard) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ResourceQuotaEphemeralStorageRequestHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaEphemeralStorageRequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaEphemeralStorageRequestUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaEphemeralStorageRequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The sum of local ephemeral storage requests in the namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaEphemeralStorageRequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaEphemeralStorageRequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaEphemeralStorageRequestUsed), nil
}

func (m ResourceQuotaEphemeralStorageRequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaEphemeralStorageRequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaEphemeralStorageRequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaEphemeralStorageRequestUsed) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ResourceQuotaEphemeralStorageRequestUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaEphemeralStorageRequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaHugepageCountRequestHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaHugepageCountRequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The huge page requests in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{hugepage}"),
}

func NewResourceQuotaHugepageCountRequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaHugepageCountRequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaHugepageCountRequestHard), nil
}

func (m ResourceQuotaHugepageCountRequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaHugepageCountRequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaHugepageCountRequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaHugepageCountRequestHard) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ResourceQuotaHugepageCountRequestHard) Add(
	ctx context.Context,
	incr int64,
	hugepageSize string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaHugepageCountRequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaHugepageCountRequestUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaHugepageCountRequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The huge page requests in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{hugepage}"),
}

func NewResourceQuotaHugepageCountRequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaHugepageCountRequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaHugepageCountRequestUsed), nil
}

func (m ResourceQuotaHugepageCountRequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaHugepageCountRequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaHugepageCountRequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaHugepageCountRequestUsed) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ResourceQuotaHugepageCountRequestUsed) Add(
	ctx context.Context,
	incr int64,
	hugepageSize string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaHugepageCountRequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaMemoryLimitHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaMemoryLimitHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The memory limits in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaMemoryLimitHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaMemoryLimitHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaMemoryLimitHard), nil
}

func (m ResourceQuotaMemoryLimitHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaMemoryLimitHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryLimitHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryLimitHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaMemoryLimitHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaMemoryLimitHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaMemoryLimitUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaMemoryLimitUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The memory limits in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaMemoryLimitUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaMemoryLimitUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaMemoryLimitUsed), nil
}

func (m ResourceQuotaMemoryLimitUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaMemoryLimitUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryLimitUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryLimitUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaMemoryLimitUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaMemoryLimitUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaMemoryRequestHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaMemoryRequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The memory requests in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaMemoryRequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaMemoryRequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaMemoryRequestHard), nil
}

func (m ResourceQuotaMemoryRequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaMemoryRequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryRequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryRequestHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaMemoryRequestHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaMemoryRequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaMemoryRequestUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaMemoryRequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The memory requests in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaMemoryRequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaMemoryRequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaMemoryRequestUsed), nil
}

func (m ResourceQuotaMemoryRequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaMemoryRequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryRequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryRequestUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaMemoryRequestUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaMemoryRequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaObjectCountHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaObjectCountHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The object count limits in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{object}"),
}

func NewResourceQuotaObjectCountHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaObjectCountHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaObjectCountHard), nil
}

func (m ResourceQuotaObjectCountHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaObjectCountHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaObjectCountHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaObjectCountHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaObjectCountHard) Add(
	ctx context.Context,
	incr int64,
	resourcequotaResourceName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaObjectCountHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaObjectCountUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaObjectCountUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The object count limits in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{object}"),
}

func NewResourceQuotaObjectCountUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaObjectCountUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaObjectCountUsed), nil
}

func (m ResourceQuotaObjectCountUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaObjectCountUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaObjectCountUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaObjectCountUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaObjectCountUsed) Add(
	ctx context.Context,
	incr int64,
	resourcequotaResourceName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaObjectCountUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ResourceQuotaPersistentvolumeclaimCountHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaPersistentvolumeclaimCountHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The total number of PersistentVolumeClaims that can exist in the namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewResourceQuotaPersistentvolumeclaimCountHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaPersistentvolumeclaimCountHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaPersistentvolumeclaimCountHard), nil
}

func (m ResourceQuotaPersistentvolumeclaimCountHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaPersistentvolumeclaimCountHard) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountHard) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountHard) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ResourceQuotaPersistentvolumeclaimCountHard) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaPersistentvolumeclaimCountHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ResourceQuotaPersistentvolumeclaimCountHard) AttrStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ResourceQuotaPersistentvolumeclaimCountUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaPersistentvolumeclaimCountUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The total number of PersistentVolumeClaims that can exist in the namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewResourceQuotaPersistentvolumeclaimCountUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaPersistentvolumeclaimCountUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaPersistentvolumeclaimCountUsed), nil
}

func (m ResourceQuotaPersistentvolumeclaimCountUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaPersistentvolumeclaimCountUsed) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountUsed) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountUsed) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ResourceQuotaPersistentvolumeclaimCountUsed) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaPersistentvolumeclaimCountUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ResourceQuotaPersistentvolumeclaimCountUsed) AttrStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ResourceQuotaStorageRequestHard struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaStorageRequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The storage requests in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaStorageRequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaStorageRequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaStorageRequestHard), nil
}

func (m ResourceQuotaStorageRequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaStorageRequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaStorageRequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaStorageRequestHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaStorageRequestHard) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaStorageRequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ResourceQuotaStorageRequestHard) AttrStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ResourceQuotaStorageRequestUsed struct {
	metric.Int64UpDownCounter
}

var newResourceQuotaStorageRequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The storage requests in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaStorageRequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ResourceQuotaStorageRequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaStorageRequestUsed), nil
}

func (m ResourceQuotaStorageRequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ResourceQuotaStorageRequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaStorageRequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaStorageRequestUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ResourceQuotaStorageRequestUsed) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ResourceQuotaStorageRequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ResourceQuotaStorageRequestUsed) AttrStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type StatefulSetCurrentPods struct {
	metric.Int64UpDownCounter
}

var newStatefulSetCurrentPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of replica pods created by the statefulset controller from the statefulset version indicated by currentRevision."),
	metric.WithUnit("{pod}"),
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

func (m StatefulSetCurrentPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetDesiredPods struct {
	metric.Int64UpDownCounter
}

var newStatefulSetDesiredPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this statefulset."),
	metric.WithUnit("{pod}"),
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

func (m StatefulSetDesiredPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetReadyPods struct {
	metric.Int64UpDownCounter
}

var newStatefulSetReadyPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of replica pods created for this statefulset with a Ready Condition."),
	metric.WithUnit("{pod}"),
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

func (m StatefulSetReadyPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetUpdatedPods struct {
	metric.Int64UpDownCounter
}

var newStatefulSetUpdatedPodsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of replica pods created by the statefulset controller from the statefulset version indicated by updateRevision."),
	metric.WithUnit("{pod}"),
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

func (m StatefulSetUpdatedPods) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}
