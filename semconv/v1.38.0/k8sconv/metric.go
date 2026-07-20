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

type PodStatusPhaseAttr string

var (
	PodStatusPhasePending PodStatusPhaseAttr = "Pending"

	PodStatusPhaseRunning PodStatusPhaseAttr = "Running"

	PodStatusPhaseSucceeded PodStatusPhaseAttr = "Succeeded"

	PodStatusPhaseFailed PodStatusPhaseAttr = "Failed"

	PodStatusPhaseUnknown PodStatusPhaseAttr = "Unknown"
)

type PodStatusReasonAttr string

var (
	PodStatusReasonEvicted PodStatusReasonAttr = "Evicted"

	PodStatusReasonNodeAffinity PodStatusReasonAttr = "NodeAffinity"

	PodStatusReasonNodeLost PodStatusReasonAttr = "NodeLost"

	PodStatusReasonShutdown PodStatusReasonAttr = "Shutdown"

	PodStatusReasonUnexpectedAdmissionError PodStatusReasonAttr = "UnexpectedAdmissionError"
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

type SystemPagingFaultTypeAttr string

var (
	SystemPagingFaultTypeMajor SystemPagingFaultTypeAttr = "major"

	SystemPagingFaultTypeMinor SystemPagingFaultTypeAttr = "minor"
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

type ContainerCPULimitUtilization struct {
	metric.Int64Gauge
}

var newContainerCPULimitUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("The ratio of container CPU usage to its CPU limit."),
	metric.WithUnit("1"),
}

func NewContainerCPULimitUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (ContainerCPULimitUtilization, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPULimitUtilization), nil
}

func (m ContainerCPULimitUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (ContainerCPULimitUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerCPULimitUtilization) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerCPULimitUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
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

type ContainerCPURequestUtilization struct {
	metric.Int64Gauge
}

var newContainerCPURequestUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("The ratio of container CPU usage to its CPU request."),
	metric.WithUnit("1"),
}

func NewContainerCPURequestUtilization(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (ContainerCPURequestUtilization, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPURequestUtilization), nil
}

func (m ContainerCPURequestUtilization) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (ContainerCPURequestUtilization) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestUtilization) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestUtilization) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerCPURequestUtilization) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerCPURequestUtilization) RecordSet(ctx context.Context, val int64, set attribute.Set) {
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

type CronJobJobActive struct {
	metric.Int64UpDownCounter
}

var newCronJobJobActiveOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of actively running jobs for a cronjob."),
	metric.WithUnit("{job}"),
}

func NewCronJobJobActive(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (CronJobJobActive, error) {
	_ = "STUB: not implemented"
	return *new(CronJobJobActive), nil
}

func (m CronJobJobActive) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (CronJobJobActive) Name() string { _ = "STUB: not implemented"; return "" }

func (CronJobJobActive) Unit() string { _ = "STUB: not implemented"; return "" }

func (CronJobJobActive) Description() string { _ = "STUB: not implemented"; return "" }

func (m CronJobJobActive) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m CronJobJobActive) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetNodeCurrentScheduled struct {
	metric.Int64UpDownCounter
}

var newDaemonSetNodeCurrentScheduledOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod."),
	metric.WithUnit("{node}"),
}

func NewDaemonSetNodeCurrentScheduled(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DaemonSetNodeCurrentScheduled, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetNodeCurrentScheduled), nil
}

func (m DaemonSetNodeCurrentScheduled) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DaemonSetNodeCurrentScheduled) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeCurrentScheduled) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeCurrentScheduled) Description() string { _ = "STUB: not implemented"; return "" }

func (m DaemonSetNodeCurrentScheduled) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m DaemonSetNodeCurrentScheduled) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetNodeDesiredScheduled struct {
	metric.Int64UpDownCounter
}

var newDaemonSetNodeDesiredScheduledOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)."),
	metric.WithUnit("{node}"),
}

func NewDaemonSetNodeDesiredScheduled(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DaemonSetNodeDesiredScheduled, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetNodeDesiredScheduled), nil
}

func (m DaemonSetNodeDesiredScheduled) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DaemonSetNodeDesiredScheduled) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeDesiredScheduled) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeDesiredScheduled) Description() string { _ = "STUB: not implemented"; return "" }

func (m DaemonSetNodeDesiredScheduled) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m DaemonSetNodeDesiredScheduled) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetNodeMisscheduled struct {
	metric.Int64UpDownCounter
}

var newDaemonSetNodeMisscheduledOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod."),
	metric.WithUnit("{node}"),
}

func NewDaemonSetNodeMisscheduled(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DaemonSetNodeMisscheduled, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetNodeMisscheduled), nil
}

func (m DaemonSetNodeMisscheduled) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DaemonSetNodeMisscheduled) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeMisscheduled) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeMisscheduled) Description() string { _ = "STUB: not implemented"; return "" }

func (m DaemonSetNodeMisscheduled) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m DaemonSetNodeMisscheduled) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DaemonSetNodeReady struct {
	metric.Int64UpDownCounter
}

var newDaemonSetNodeReadyOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready."),
	metric.WithUnit("{node}"),
}

func NewDaemonSetNodeReady(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DaemonSetNodeReady, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetNodeReady), nil
}

func (m DaemonSetNodeReady) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DaemonSetNodeReady) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeReady) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeReady) Description() string { _ = "STUB: not implemented"; return "" }

func (m DaemonSetNodeReady) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m DaemonSetNodeReady) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DeploymentPodAvailable struct {
	metric.Int64UpDownCounter
}

var newDeploymentPodAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this deployment."),
	metric.WithUnit("{pod}"),
}

func NewDeploymentPodAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DeploymentPodAvailable, error) {
	_ = "STUB: not implemented"
	return *new(DeploymentPodAvailable), nil
}

func (m DeploymentPodAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DeploymentPodAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (DeploymentPodAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DeploymentPodAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m DeploymentPodAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m DeploymentPodAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type DeploymentPodDesired struct {
	metric.Int64UpDownCounter
}

var newDeploymentPodDesiredOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this deployment."),
	metric.WithUnit("{pod}"),
}

func NewDeploymentPodDesired(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (DeploymentPodDesired, error) {
	_ = "STUB: not implemented"
	return *new(DeploymentPodDesired), nil
}

func (m DeploymentPodDesired) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (DeploymentPodDesired) Name() string { _ = "STUB: not implemented"; return "" }

func (DeploymentPodDesired) Unit() string { _ = "STUB: not implemented"; return "" }

func (DeploymentPodDesired) Description() string { _ = "STUB: not implemented"; return "" }

func (m DeploymentPodDesired) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m DeploymentPodDesired) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type HPAPodCurrent struct {
	metric.Int64UpDownCounter
}

var newHPAPodCurrentOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Current number of replica pods managed by this horizontal pod autoscaler, as last seen by the autoscaler."),
	metric.WithUnit("{pod}"),
}

func NewHPAPodCurrent(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (HPAPodCurrent, error) {
	_ = "STUB: not implemented"
	return *new(HPAPodCurrent), nil
}

func (m HPAPodCurrent) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (HPAPodCurrent) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAPodCurrent) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAPodCurrent) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPAPodCurrent) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m HPAPodCurrent) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type HPAPodDesired struct {
	metric.Int64UpDownCounter
}

var newHPAPodDesiredOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Desired number of replica pods managed by this horizontal pod autoscaler, as last calculated by the autoscaler."),
	metric.WithUnit("{pod}"),
}

func NewHPAPodDesired(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (HPAPodDesired, error) {
	_ = "STUB: not implemented"
	return *new(HPAPodDesired), nil
}

func (m HPAPodDesired) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (HPAPodDesired) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAPodDesired) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAPodDesired) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPAPodDesired) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m HPAPodDesired) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type HPAPodMax struct {
	metric.Int64UpDownCounter
}

var newHPAPodMaxOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The upper limit for the number of replica pods to which the autoscaler can scale up."),
	metric.WithUnit("{pod}"),
}

func NewHPAPodMax(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (HPAPodMax, error) {
	_ = "STUB: not implemented"
	return *new(HPAPodMax), nil
}

func (m HPAPodMax) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (HPAPodMax) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAPodMax) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAPodMax) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPAPodMax) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m HPAPodMax) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type HPAPodMin struct {
	metric.Int64UpDownCounter
}

var newHPAPodMinOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The lower limit for the number of replica pods to which the autoscaler can scale down."),
	metric.WithUnit("{pod}"),
}

func NewHPAPodMin(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (HPAPodMin, error) {
	_ = "STUB: not implemented"
	return *new(HPAPodMin), nil
}

func (m HPAPodMin) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (HPAPodMin) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAPodMin) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAPodMin) Description() string { _ = "STUB: not implemented"; return "" }

func (m HPAPodMin) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m HPAPodMin) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobPodActive struct {
	metric.Int64UpDownCounter
}

var newJobPodActiveOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of pending and actively running pods for a job."),
	metric.WithUnit("{pod}"),
}

func NewJobPodActive(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobPodActive, error) {
	_ = "STUB: not implemented"
	return *new(JobPodActive), nil
}

func (m JobPodActive) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobPodActive) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodActive) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodActive) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobPodActive) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m JobPodActive) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobPodDesiredSuccessful struct {
	metric.Int64UpDownCounter
}

var newJobPodDesiredSuccessfulOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The desired number of successfully finished pods the job should be run with."),
	metric.WithUnit("{pod}"),
}

func NewJobPodDesiredSuccessful(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobPodDesiredSuccessful, error) {
	_ = "STUB: not implemented"
	return *new(JobPodDesiredSuccessful), nil
}

func (m JobPodDesiredSuccessful) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobPodDesiredSuccessful) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodDesiredSuccessful) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodDesiredSuccessful) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobPodDesiredSuccessful) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m JobPodDesiredSuccessful) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobPodFailed struct {
	metric.Int64UpDownCounter
}

var newJobPodFailedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of pods which reached phase Failed for a job."),
	metric.WithUnit("{pod}"),
}

func NewJobPodFailed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobPodFailed, error) {
	_ = "STUB: not implemented"
	return *new(JobPodFailed), nil
}

func (m JobPodFailed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobPodFailed) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodFailed) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodFailed) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobPodFailed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m JobPodFailed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobPodMaxParallel struct {
	metric.Int64UpDownCounter
}

var newJobPodMaxParallelOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The max desired number of pods the job should run at any given time."),
	metric.WithUnit("{pod}"),
}

func NewJobPodMaxParallel(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobPodMaxParallel, error) {
	_ = "STUB: not implemented"
	return *new(JobPodMaxParallel), nil
}

func (m JobPodMaxParallel) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobPodMaxParallel) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodMaxParallel) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodMaxParallel) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobPodMaxParallel) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m JobPodMaxParallel) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type JobPodSuccessful struct {
	metric.Int64UpDownCounter
}

var newJobPodSuccessfulOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of pods which reached phase Succeeded for a job."),
	metric.WithUnit("{pod}"),
}

func NewJobPodSuccessful(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (JobPodSuccessful, error) {
	_ = "STUB: not implemented"
	return *new(JobPodSuccessful), nil
}

func (m JobPodSuccessful) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (JobPodSuccessful) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodSuccessful) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodSuccessful) Description() string { _ = "STUB: not implemented"; return "" }

func (m JobPodSuccessful) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m JobPodSuccessful) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type NodeCPUAllocatable struct {
	metric.Int64UpDownCounter
}

var newNodeCPUAllocatableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Amount of cpu allocatable on the node."),
	metric.WithUnit("{cpu}"),
}

func NewNodeCPUAllocatable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeCPUAllocatable, error) {
	_ = "STUB: not implemented"
	return *new(NodeCPUAllocatable), nil
}

func (m NodeCPUAllocatable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeCPUAllocatable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUAllocatable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUAllocatable) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeCPUAllocatable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeCPUAllocatable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type NodeEphemeralStorageAllocatable struct {
	metric.Int64UpDownCounter
}

var newNodeEphemeralStorageAllocatableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Amount of ephemeral-storage allocatable on the node."),
	metric.WithUnit("By"),
}

func NewNodeEphemeralStorageAllocatable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeEphemeralStorageAllocatable, error) {
	_ = "STUB: not implemented"
	return *new(NodeEphemeralStorageAllocatable), nil
}

func (m NodeEphemeralStorageAllocatable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeEphemeralStorageAllocatable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeEphemeralStorageAllocatable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeEphemeralStorageAllocatable) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeEphemeralStorageAllocatable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeEphemeralStorageAllocatable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type NodeMemoryAllocatable struct {
	metric.Int64UpDownCounter
}

var newNodeMemoryAllocatableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Amount of memory allocatable on the node."),
	metric.WithUnit("By"),
}

func NewNodeMemoryAllocatable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeMemoryAllocatable, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryAllocatable), nil
}

func (m NodeMemoryAllocatable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeMemoryAllocatable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryAllocatable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryAllocatable) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeMemoryAllocatable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeMemoryAllocatable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeMemoryAvailable struct {
	metric.Int64UpDownCounter
}

var newNodeMemoryAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Node memory available."),
	metric.WithUnit("By"),
}

func NewNodeMemoryAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeMemoryAvailable, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryAvailable), nil
}

func (m NodeMemoryAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeMemoryAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeMemoryAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeMemoryAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeMemoryPagingFaults struct {
	metric.Int64Counter
}

var newNodeMemoryPagingFaultsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Node memory paging faults."),
	metric.WithUnit("{fault}"),
}

func NewNodeMemoryPagingFaults(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (NodeMemoryPagingFaults, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryPagingFaults), nil
}

func (m NodeMemoryPagingFaults) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (NodeMemoryPagingFaults) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryPagingFaults) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryPagingFaults) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeMemoryPagingFaults) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m NodeMemoryPagingFaults) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (NodeMemoryPagingFaults) AttrSystemPagingFaultType(val SystemPagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type NodeMemoryRss struct {
	metric.Int64UpDownCounter
}

var newNodeMemoryRssOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Node memory RSS."),
	metric.WithUnit("By"),
}

func NewNodeMemoryRss(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeMemoryRss, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryRss), nil
}

func (m NodeMemoryRss) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeMemoryRss) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryRss) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryRss) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeMemoryRss) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeMemoryRss) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type NodeMemoryWorkingSet struct {
	metric.Int64UpDownCounter
}

var newNodeMemoryWorkingSetOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Node memory working set."),
	metric.WithUnit("By"),
}

func NewNodeMemoryWorkingSet(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeMemoryWorkingSet, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryWorkingSet), nil
}

func (m NodeMemoryWorkingSet) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeMemoryWorkingSet) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryWorkingSet) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryWorkingSet) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeMemoryWorkingSet) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeMemoryWorkingSet) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type NodePodAllocatable struct {
	metric.Int64UpDownCounter
}

var newNodePodAllocatableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Amount of pods allocatable on the node."),
	metric.WithUnit("{pod}"),
}

func NewNodePodAllocatable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodePodAllocatable, error) {
	_ = "STUB: not implemented"
	return *new(NodePodAllocatable), nil
}

func (m NodePodAllocatable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodePodAllocatable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodePodAllocatable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodePodAllocatable) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodePodAllocatable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodePodAllocatable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
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

type PodMemoryAvailable struct {
	metric.Int64UpDownCounter
}

var newPodMemoryAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Pod memory available."),
	metric.WithUnit("By"),
}

func NewPodMemoryAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodMemoryAvailable, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryAvailable), nil
}

func (m PodMemoryAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodMemoryAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodMemoryAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m PodMemoryAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodMemoryPagingFaults struct {
	metric.Int64Counter
}

var newPodMemoryPagingFaultsOpts = []metric.Int64CounterOption{
	metric.WithDescription("Pod memory paging faults."),
	metric.WithUnit("{fault}"),
}

func NewPodMemoryPagingFaults(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (PodMemoryPagingFaults, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryPagingFaults), nil
}

func (m PodMemoryPagingFaults) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (PodMemoryPagingFaults) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryPagingFaults) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryPagingFaults) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodMemoryPagingFaults) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PodMemoryPagingFaults) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (PodMemoryPagingFaults) AttrSystemPagingFaultType(val SystemPagingFaultTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PodMemoryRss struct {
	metric.Int64UpDownCounter
}

var newPodMemoryRssOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Pod memory RSS."),
	metric.WithUnit("By"),
}

func NewPodMemoryRss(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodMemoryRss, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryRss), nil
}

func (m PodMemoryRss) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodMemoryRss) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryRss) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryRss) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodMemoryRss) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m PodMemoryRss) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type PodMemoryWorkingSet struct {
	metric.Int64UpDownCounter
}

var newPodMemoryWorkingSetOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Pod memory working set."),
	metric.WithUnit("By"),
}

func NewPodMemoryWorkingSet(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodMemoryWorkingSet, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryWorkingSet), nil
}

func (m PodMemoryWorkingSet) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodMemoryWorkingSet) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryWorkingSet) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryWorkingSet) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodMemoryWorkingSet) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m PodMemoryWorkingSet) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type PodStatusPhase struct {
	metric.Int64UpDownCounter
}

var newPodStatusPhaseOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Describes number of K8s Pods that are currently in a given phase."),
	metric.WithUnit("{pod}"),
}

func NewPodStatusPhase(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodStatusPhase, error) {
	_ = "STUB: not implemented"
	return *new(PodStatusPhase), nil
}

func (m PodStatusPhase) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodStatusPhase) Name() string { _ = "STUB: not implemented"; return "" }

func (PodStatusPhase) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodStatusPhase) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodStatusPhase) Add(
	ctx context.Context,
	incr int64,
	podStatusPhase PodStatusPhaseAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PodStatusPhase) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PodStatusReason struct {
	metric.Int64UpDownCounter
}

var newPodStatusReasonOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Describes the number of K8s Pods that are currently in a state for a given reason."),
	metric.WithUnit("{pod}"),
}

func NewPodStatusReason(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PodStatusReason, error) {
	_ = "STUB: not implemented"
	return *new(PodStatusReason), nil
}

func (m PodStatusReason) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PodStatusReason) Name() string { _ = "STUB: not implemented"; return "" }

func (PodStatusReason) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodStatusReason) Description() string { _ = "STUB: not implemented"; return "" }

func (m PodStatusReason) Add(
	ctx context.Context,
	incr int64,
	podStatusReason PodStatusReasonAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PodStatusReason) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
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

type ReplicaSetPodAvailable struct {
	metric.Int64UpDownCounter
}

var newReplicaSetPodAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this replicaset."),
	metric.WithUnit("{pod}"),
}

func NewReplicaSetPodAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ReplicaSetPodAvailable, error) {
	_ = "STUB: not implemented"
	return *new(ReplicaSetPodAvailable), nil
}

func (m ReplicaSetPodAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ReplicaSetPodAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetPodAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetPodAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m ReplicaSetPodAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ReplicaSetPodAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ReplicaSetPodDesired struct {
	metric.Int64UpDownCounter
}

var newReplicaSetPodDesiredOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this replicaset."),
	metric.WithUnit("{pod}"),
}

func NewReplicaSetPodDesired(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ReplicaSetPodDesired, error) {
	_ = "STUB: not implemented"
	return *new(ReplicaSetPodDesired), nil
}

func (m ReplicaSetPodDesired) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ReplicaSetPodDesired) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetPodDesired) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetPodDesired) Description() string { _ = "STUB: not implemented"; return "" }

func (m ReplicaSetPodDesired) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ReplicaSetPodDesired) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ReplicationControllerPodAvailable struct {
	metric.Int64UpDownCounter
}

var newReplicationControllerPodAvailableOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this replication controller."),
	metric.WithUnit("{pod}"),
}

func NewReplicationControllerPodAvailable(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ReplicationControllerPodAvailable, error) {
	_ = "STUB: not implemented"
	return *new(ReplicationControllerPodAvailable), nil
}

func (m ReplicationControllerPodAvailable) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ReplicationControllerPodAvailable) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicationControllerPodAvailable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicationControllerPodAvailable) Description() string { _ = "STUB: not implemented"; return "" }

func (m ReplicationControllerPodAvailable) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ReplicationControllerPodAvailable) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ReplicationControllerPodDesired struct {
	metric.Int64UpDownCounter
}

var newReplicationControllerPodDesiredOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this replication controller."),
	metric.WithUnit("{pod}"),
}

func NewReplicationControllerPodDesired(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ReplicationControllerPodDesired, error) {
	_ = "STUB: not implemented"
	return *new(ReplicationControllerPodDesired), nil
}

func (m ReplicationControllerPodDesired) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ReplicationControllerPodDesired) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicationControllerPodDesired) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicationControllerPodDesired) Description() string { _ = "STUB: not implemented"; return "" }

func (m ReplicationControllerPodDesired) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ReplicationControllerPodDesired) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

type StatefulSetPodCurrent struct {
	metric.Int64UpDownCounter
}

var newStatefulSetPodCurrentOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of replica pods created by the statefulset controller from the statefulset version indicated by currentRevision."),
	metric.WithUnit("{pod}"),
}

func NewStatefulSetPodCurrent(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (StatefulSetPodCurrent, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetPodCurrent), nil
}

func (m StatefulSetPodCurrent) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (StatefulSetPodCurrent) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodCurrent) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodCurrent) Description() string { _ = "STUB: not implemented"; return "" }

func (m StatefulSetPodCurrent) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m StatefulSetPodCurrent) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetPodDesired struct {
	metric.Int64UpDownCounter
}

var newStatefulSetPodDesiredOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this statefulset."),
	metric.WithUnit("{pod}"),
}

func NewStatefulSetPodDesired(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (StatefulSetPodDesired, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetPodDesired), nil
}

func (m StatefulSetPodDesired) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (StatefulSetPodDesired) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodDesired) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodDesired) Description() string { _ = "STUB: not implemented"; return "" }

func (m StatefulSetPodDesired) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m StatefulSetPodDesired) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetPodReady struct {
	metric.Int64UpDownCounter
}

var newStatefulSetPodReadyOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of replica pods created for this statefulset with a Ready Condition."),
	metric.WithUnit("{pod}"),
}

func NewStatefulSetPodReady(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (StatefulSetPodReady, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetPodReady), nil
}

func (m StatefulSetPodReady) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (StatefulSetPodReady) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodReady) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodReady) Description() string { _ = "STUB: not implemented"; return "" }

func (m StatefulSetPodReady) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m StatefulSetPodReady) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type StatefulSetPodUpdated struct {
	metric.Int64UpDownCounter
}

var newStatefulSetPodUpdatedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of replica pods created by the statefulset controller from the statefulset version indicated by updateRevision."),
	metric.WithUnit("{pod}"),
}

func NewStatefulSetPodUpdated(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (StatefulSetPodUpdated, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetPodUpdated), nil
}

func (m StatefulSetPodUpdated) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (StatefulSetPodUpdated) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodUpdated) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodUpdated) Description() string { _ = "STUB: not implemented"; return "" }

func (m StatefulSetPodUpdated) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m StatefulSetPodUpdated) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}
