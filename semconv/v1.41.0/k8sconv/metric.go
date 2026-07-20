package k8sconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
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

type PersistentvolumeStatusPhaseAttr string

var (
	PersistentvolumeStatusPhaseAvailable PersistentvolumeStatusPhaseAttr = "Available"

	PersistentvolumeStatusPhaseBound PersistentvolumeStatusPhaseAttr = "Bound"

	PersistentvolumeStatusPhaseFailed PersistentvolumeStatusPhaseAttr = "Failed"

	PersistentvolumeStatusPhasePending PersistentvolumeStatusPhaseAttr = "Pending"

	PersistentvolumeStatusPhaseReleased PersistentvolumeStatusPhaseAttr = "Released"
)

type PersistentvolumeclaimStatusPhaseAttr string

var (
	PersistentvolumeclaimStatusPhaseBound PersistentvolumeclaimStatusPhaseAttr = "Bound"

	PersistentvolumeclaimStatusPhaseLost PersistentvolumeclaimStatusPhaseAttr = "Lost"

	PersistentvolumeclaimStatusPhasePending PersistentvolumeclaimStatusPhaseAttr = "Pending"
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

type ServiceEndpointAddressTypeAttr string

var (
	ServiceEndpointAddressTypeIPv4 ServiceEndpointAddressTypeAttr = "IPv4"

	ServiceEndpointAddressTypeIPv6 ServiceEndpointAddressTypeAttr = "IPv6"

	ServiceEndpointAddressTypeFqdn ServiceEndpointAddressTypeAttr = "FQDN"
)

type ServiceEndpointConditionAttr string

var (
	ServiceEndpointConditionReady ServiceEndpointConditionAttr = "ready"

	ServiceEndpointConditionServing ServiceEndpointConditionAttr = "serving"

	ServiceEndpointConditionTerminating ServiceEndpointConditionAttr = "terminating"
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

type ContainerCPULimitCurrent struct {
	metric.Int64UpDownCounter
}

var newContainerCPULimitCurrentOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Maximum CPU resource limit currently configured for a running container."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPULimitCurrent(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerCPULimitCurrent, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPULimitCurrent), nil
}

func (m ContainerCPULimitCurrent) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerCPULimitCurrent) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitCurrent) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitCurrent) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerCPULimitCurrent) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerCPULimitCurrent) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerCPULimitCurrentObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerCPULimitCurrentObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Maximum CPU resource limit currently configured for a running container."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPULimitCurrentObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerCPULimitCurrentObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPULimitCurrentObservable), nil
}

func (m ContainerCPULimitCurrentObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerCPULimitCurrentObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitCurrentObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitCurrentObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type ContainerCPULimitDesired struct {
	metric.Int64UpDownCounter
}

var newContainerCPULimitDesiredOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Maximum CPU resource limit as defined by the container spec."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPULimitDesired(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerCPULimitDesired, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPULimitDesired), nil
}

func (m ContainerCPULimitDesired) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerCPULimitDesired) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitDesired) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitDesired) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerCPULimitDesired) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerCPULimitDesired) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerCPULimitDesiredObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerCPULimitDesiredObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Maximum CPU resource limit as defined by the container spec."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPULimitDesiredObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerCPULimitDesiredObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPULimitDesiredObservable), nil
}

func (m ContainerCPULimitDesiredObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerCPULimitDesiredObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitDesiredObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitDesiredObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type ContainerCPULimitUtilization struct {
	metric.Int64Gauge
}

var newContainerCPULimitUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("The ratio of container CPU usage to its current CPU limit."),
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

type ContainerCPULimitUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newContainerCPULimitUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("The ratio of container CPU usage to its current CPU limit."),
	metric.WithUnit("1"),
}

func NewContainerCPULimitUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (ContainerCPULimitUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPULimitUtilizationObservable), nil
}

func (m ContainerCPULimitUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (ContainerCPULimitUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPULimitUtilizationObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type ContainerCPURequestCurrent struct {
	metric.Int64UpDownCounter
}

var newContainerCPURequestCurrentOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("CPU resource requested currently configured for a running container."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPURequestCurrent(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerCPURequestCurrent, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPURequestCurrent), nil
}

func (m ContainerCPURequestCurrent) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerCPURequestCurrent) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestCurrent) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestCurrent) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerCPURequestCurrent) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerCPURequestCurrent) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerCPURequestCurrentObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerCPURequestCurrentObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("CPU resource requested currently configured for a running container."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPURequestCurrentObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerCPURequestCurrentObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPURequestCurrentObservable), nil
}

func (m ContainerCPURequestCurrentObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerCPURequestCurrentObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestCurrentObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestCurrentObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type ContainerCPURequestDesired struct {
	metric.Int64UpDownCounter
}

var newContainerCPURequestDesiredOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("CPU resource requested as defined by the container spec."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPURequestDesired(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerCPURequestDesired, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPURequestDesired), nil
}

func (m ContainerCPURequestDesired) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerCPURequestDesired) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestDesired) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestDesired) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerCPURequestDesired) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerCPURequestDesired) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerCPURequestDesiredObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerCPURequestDesiredObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("CPU resource requested as defined by the container spec."),
	metric.WithUnit("{cpu}"),
}

func NewContainerCPURequestDesiredObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerCPURequestDesiredObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPURequestDesiredObservable), nil
}

func (m ContainerCPURequestDesiredObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerCPURequestDesiredObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestDesiredObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestDesiredObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type ContainerCPURequestUtilization struct {
	metric.Int64Gauge
}

var newContainerCPURequestUtilizationOpts = []metric.Int64GaugeOption{
	metric.WithDescription("The ratio of container CPU usage to its current CPU request."),
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

type ContainerCPURequestUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newContainerCPURequestUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("The ratio of container CPU usage to its current CPU request."),
	metric.WithUnit("1"),
}

func NewContainerCPURequestUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (ContainerCPURequestUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerCPURequestUtilizationObservable), nil
}

func (m ContainerCPURequestUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (ContainerCPURequestUtilizationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestUtilizationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerCPURequestUtilizationObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ContainerEphemeralStorageLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerEphemeralStorageLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Maximum ephemeral storage resource limit set for the container."),
	metric.WithUnit("By"),
}

func NewContainerEphemeralStorageLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerEphemeralStorageLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerEphemeralStorageLimitObservable), nil
}

func (m ContainerEphemeralStorageLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerEphemeralStorageLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerEphemeralStorageLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerEphemeralStorageLimitObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ContainerEphemeralStorageRequestObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerEphemeralStorageRequestObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Ephemeral storage resource requested for the container."),
	metric.WithUnit("By"),
}

func NewContainerEphemeralStorageRequestObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerEphemeralStorageRequestObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerEphemeralStorageRequestObservable), nil
}

func (m ContainerEphemeralStorageRequestObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerEphemeralStorageRequestObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ContainerEphemeralStorageRequestObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ContainerEphemeralStorageRequestObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type ContainerMemoryLimitCurrent struct {
	metric.Int64UpDownCounter
}

var newContainerMemoryLimitCurrentOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Maximum memory resource limit currently configured for a running container."),
	metric.WithUnit("By"),
}

func NewContainerMemoryLimitCurrent(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerMemoryLimitCurrent, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryLimitCurrent), nil
}

func (m ContainerMemoryLimitCurrent) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerMemoryLimitCurrent) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimitCurrent) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimitCurrent) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerMemoryLimitCurrent) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerMemoryLimitCurrent) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerMemoryLimitCurrentObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerMemoryLimitCurrentObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Maximum memory resource limit currently configured for a running container."),
	metric.WithUnit("By"),
}

func NewContainerMemoryLimitCurrentObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerMemoryLimitCurrentObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryLimitCurrentObservable), nil
}

func (m ContainerMemoryLimitCurrentObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerMemoryLimitCurrentObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimitCurrentObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimitCurrentObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type ContainerMemoryLimitDesired struct {
	metric.Int64UpDownCounter
}

var newContainerMemoryLimitDesiredOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Maximum memory resource limit as defined by the container spec."),
	metric.WithUnit("By"),
}

func NewContainerMemoryLimitDesired(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerMemoryLimitDesired, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryLimitDesired), nil
}

func (m ContainerMemoryLimitDesired) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerMemoryLimitDesired) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimitDesired) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimitDesired) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerMemoryLimitDesired) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerMemoryLimitDesired) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerMemoryLimitDesiredObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerMemoryLimitDesiredObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Maximum memory resource limit as defined by the container spec."),
	metric.WithUnit("By"),
}

func NewContainerMemoryLimitDesiredObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerMemoryLimitDesiredObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryLimitDesiredObservable), nil
}

func (m ContainerMemoryLimitDesiredObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerMemoryLimitDesiredObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimitDesiredObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryLimitDesiredObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type ContainerMemoryRequestCurrent struct {
	metric.Int64UpDownCounter
}

var newContainerMemoryRequestCurrentOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Memory resource request currently configured for a running container."),
	metric.WithUnit("By"),
}

func NewContainerMemoryRequestCurrent(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerMemoryRequestCurrent, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryRequestCurrent), nil
}

func (m ContainerMemoryRequestCurrent) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerMemoryRequestCurrent) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequestCurrent) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequestCurrent) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerMemoryRequestCurrent) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerMemoryRequestCurrent) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerMemoryRequestCurrentObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerMemoryRequestCurrentObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Memory resource request currently configured for a running container."),
	metric.WithUnit("By"),
}

func NewContainerMemoryRequestCurrentObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerMemoryRequestCurrentObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryRequestCurrentObservable), nil
}

func (m ContainerMemoryRequestCurrentObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerMemoryRequestCurrentObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequestCurrentObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequestCurrentObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type ContainerMemoryRequestDesired struct {
	metric.Int64UpDownCounter
}

var newContainerMemoryRequestDesiredOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Memory resource requested as defined by the container spec."),
	metric.WithUnit("By"),
}

func NewContainerMemoryRequestDesired(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ContainerMemoryRequestDesired, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryRequestDesired), nil
}

func (m ContainerMemoryRequestDesired) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ContainerMemoryRequestDesired) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequestDesired) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequestDesired) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContainerMemoryRequestDesired) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ContainerMemoryRequestDesired) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ContainerMemoryRequestDesiredObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerMemoryRequestDesiredObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Memory resource requested as defined by the container spec."),
	metric.WithUnit("By"),
}

func NewContainerMemoryRequestDesiredObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerMemoryRequestDesiredObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerMemoryRequestDesiredObservable), nil
}

func (m ContainerMemoryRequestDesiredObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerMemoryRequestDesiredObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequestDesiredObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerMemoryRequestDesiredObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ContainerReadyObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerReadyObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Indicates whether the container is currently marked as ready to accept traffic, based on its readiness probe (1 = ready, 0 = not ready)."),
	metric.WithUnit("{container}"),
}

func NewContainerReadyObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerReadyObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerReadyObservable), nil
}

func (m ContainerReadyObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerReadyObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerReadyObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerReadyObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type ContainerRestartCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerRestartCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Describes how many times the container has restarted (since the last counter reset)."),
	metric.WithUnit("{restart}"),
}

func NewContainerRestartCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerRestartCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerRestartCountObservable), nil
}

func (m ContainerRestartCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerRestartCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerRestartCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerRestartCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type ContainerStatusReasonObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerStatusReasonObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Describes the number of K8s containers that are currently in a state for a given reason."),
	metric.WithUnit("{container}"),
}

func NewContainerStatusReasonObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerStatusReasonObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerStatusReasonObservable), nil
}

func (m ContainerStatusReasonObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerStatusReasonObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusReasonObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusReasonObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusReasonObservable) AttrContainerStatusReason(val ContainerStatusReasonAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ContainerStatusStateObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerStatusStateObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Describes the number of K8s containers that are currently in a given state."),
	metric.WithUnit("{container}"),
}

func NewContainerStatusStateObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerStatusStateObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerStatusStateObservable), nil
}

func (m ContainerStatusStateObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerStatusStateObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusStateObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusStateObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ContainerStatusStateObservable) AttrContainerStatusState(val ContainerStatusStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ContainerStorageLimitObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerStorageLimitObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Maximum storage resource limit set for the container."),
	metric.WithUnit("By"),
}

func NewContainerStorageLimitObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerStorageLimitObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerStorageLimitObservable), nil
}

func (m ContainerStorageLimitObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerStorageLimitObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerStorageLimitObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerStorageLimitObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type ContainerStorageRequestObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newContainerStorageRequestObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Storage resource requested for the container."),
	metric.WithUnit("By"),
}

func NewContainerStorageRequestObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ContainerStorageRequestObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContainerStorageRequestObservable), nil
}

func (m ContainerStorageRequestObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ContainerStorageRequestObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContainerStorageRequestObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContainerStorageRequestObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type CronJobJobActiveObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newCronJobJobActiveObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of actively running jobs for a cronjob."),
	metric.WithUnit("{job}"),
}

func NewCronJobJobActiveObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (CronJobJobActiveObservable, error) {
	_ = "STUB: not implemented"
	return *new(CronJobJobActiveObservable), nil
}

func (m CronJobJobActiveObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (CronJobJobActiveObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (CronJobJobActiveObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (CronJobJobActiveObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type DaemonSetNodeCurrentScheduledObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newDaemonSetNodeCurrentScheduledObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod."),
	metric.WithUnit("{node}"),
}

func NewDaemonSetNodeCurrentScheduledObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (DaemonSetNodeCurrentScheduledObservable, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetNodeCurrentScheduledObservable), nil
}

func (m DaemonSetNodeCurrentScheduledObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (DaemonSetNodeCurrentScheduledObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeCurrentScheduledObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeCurrentScheduledObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type DaemonSetNodeDesiredScheduledObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newDaemonSetNodeDesiredScheduledObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of nodes that should be running the daemon pod (including nodes currently running the daemon pod)."),
	metric.WithUnit("{node}"),
}

func NewDaemonSetNodeDesiredScheduledObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (DaemonSetNodeDesiredScheduledObservable, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetNodeDesiredScheduledObservable), nil
}

func (m DaemonSetNodeDesiredScheduledObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (DaemonSetNodeDesiredScheduledObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeDesiredScheduledObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeDesiredScheduledObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type DaemonSetNodeMisscheduledObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newDaemonSetNodeMisscheduledObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of nodes that are running the daemon pod, but are not supposed to run the daemon pod."),
	metric.WithUnit("{node}"),
}

func NewDaemonSetNodeMisscheduledObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (DaemonSetNodeMisscheduledObservable, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetNodeMisscheduledObservable), nil
}

func (m DaemonSetNodeMisscheduledObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (DaemonSetNodeMisscheduledObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeMisscheduledObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeMisscheduledObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type DaemonSetNodeReadyObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newDaemonSetNodeReadyObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready."),
	metric.WithUnit("{node}"),
}

func NewDaemonSetNodeReadyObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (DaemonSetNodeReadyObservable, error) {
	_ = "STUB: not implemented"
	return *new(DaemonSetNodeReadyObservable), nil
}

func (m DaemonSetNodeReadyObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (DaemonSetNodeReadyObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeReadyObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DaemonSetNodeReadyObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type DeploymentPodAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newDeploymentPodAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this deployment."),
	metric.WithUnit("{pod}"),
}

func NewDeploymentPodAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (DeploymentPodAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(DeploymentPodAvailableObservable), nil
}

func (m DeploymentPodAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (DeploymentPodAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DeploymentPodAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DeploymentPodAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type DeploymentPodDesiredObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newDeploymentPodDesiredObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this deployment."),
	metric.WithUnit("{pod}"),
}

func NewDeploymentPodDesiredObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (DeploymentPodDesiredObservable, error) {
	_ = "STUB: not implemented"
	return *new(DeploymentPodDesiredObservable), nil
}

func (m DeploymentPodDesiredObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (DeploymentPodDesiredObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (DeploymentPodDesiredObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (DeploymentPodDesiredObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type HPAMetricTargetCPUAverageUtilizationObservable struct {
	metric.Int64ObservableGauge
}

var newHPAMetricTargetCPUAverageUtilizationObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Target average utilization, in percentage, for CPU resource in HPA config."),
	metric.WithUnit("1"),
}

func NewHPAMetricTargetCPUAverageUtilizationObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (HPAMetricTargetCPUAverageUtilizationObservable, error) {
	_ = "STUB: not implemented"
	return *new(HPAMetricTargetCPUAverageUtilizationObservable), nil
}

func (m HPAMetricTargetCPUAverageUtilizationObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (HPAMetricTargetCPUAverageUtilizationObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (HPAMetricTargetCPUAverageUtilizationObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (HPAMetricTargetCPUAverageUtilizationObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (HPAMetricTargetCPUAverageUtilizationObservable) AttrContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HPAMetricTargetCPUAverageUtilizationObservable) AttrHPAMetricType(val string) attribute.KeyValue {
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

type HPAMetricTargetCPUAverageValueObservable struct {
	metric.Int64ObservableGauge
}

var newHPAMetricTargetCPUAverageValueObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Target average value for CPU resource in HPA config."),
	metric.WithUnit("{cpu}"),
}

func NewHPAMetricTargetCPUAverageValueObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (HPAMetricTargetCPUAverageValueObservable, error) {
	_ = "STUB: not implemented"
	return *new(HPAMetricTargetCPUAverageValueObservable), nil
}

func (m HPAMetricTargetCPUAverageValueObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (HPAMetricTargetCPUAverageValueObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUAverageValueObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUAverageValueObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (HPAMetricTargetCPUAverageValueObservable) AttrContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HPAMetricTargetCPUAverageValueObservable) AttrHPAMetricType(val string) attribute.KeyValue {
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

type HPAMetricTargetCPUValueObservable struct {
	metric.Int64ObservableGauge
}

var newHPAMetricTargetCPUValueObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Target value for CPU resource in HPA config."),
	metric.WithUnit("{cpu}"),
}

func NewHPAMetricTargetCPUValueObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (HPAMetricTargetCPUValueObservable, error) {
	_ = "STUB: not implemented"
	return *new(HPAMetricTargetCPUValueObservable), nil
}

func (m HPAMetricTargetCPUValueObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (HPAMetricTargetCPUValueObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUValueObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUValueObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (HPAMetricTargetCPUValueObservable) AttrContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (HPAMetricTargetCPUValueObservable) AttrHPAMetricType(val string) attribute.KeyValue {
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

type HPAPodCurrentObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newHPAPodCurrentObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Current number of replica pods managed by this horizontal pod autoscaler, as last seen by the autoscaler."),
	metric.WithUnit("{pod}"),
}

func NewHPAPodCurrentObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (HPAPodCurrentObservable, error) {
	_ = "STUB: not implemented"
	return *new(HPAPodCurrentObservable), nil
}

func (m HPAPodCurrentObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (HPAPodCurrentObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAPodCurrentObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAPodCurrentObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type HPAPodDesiredObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newHPAPodDesiredObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Desired number of replica pods managed by this horizontal pod autoscaler, as last calculated by the autoscaler."),
	metric.WithUnit("{pod}"),
}

func NewHPAPodDesiredObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (HPAPodDesiredObservable, error) {
	_ = "STUB: not implemented"
	return *new(HPAPodDesiredObservable), nil
}

func (m HPAPodDesiredObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (HPAPodDesiredObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAPodDesiredObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAPodDesiredObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type HPAPodMaxObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newHPAPodMaxObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The upper limit for the number of replica pods to which the autoscaler can scale up."),
	metric.WithUnit("{pod}"),
}

func NewHPAPodMaxObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (HPAPodMaxObservable, error) {
	_ = "STUB: not implemented"
	return *new(HPAPodMaxObservable), nil
}

func (m HPAPodMaxObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (HPAPodMaxObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAPodMaxObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAPodMaxObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type HPAPodMinObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newHPAPodMinObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The lower limit for the number of replica pods to which the autoscaler can scale down."),
	metric.WithUnit("{pod}"),
}

func NewHPAPodMinObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (HPAPodMinObservable, error) {
	_ = "STUB: not implemented"
	return *new(HPAPodMinObservable), nil
}

func (m HPAPodMinObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (HPAPodMinObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (HPAPodMinObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (HPAPodMinObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type JobPodActiveObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newJobPodActiveObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of pending and actively running pods for a job."),
	metric.WithUnit("{pod}"),
}

func NewJobPodActiveObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (JobPodActiveObservable, error) {
	_ = "STUB: not implemented"
	return *new(JobPodActiveObservable), nil
}

func (m JobPodActiveObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (JobPodActiveObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodActiveObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodActiveObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type JobPodDesiredSuccessfulObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newJobPodDesiredSuccessfulObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The desired number of successfully finished pods the job should be run with."),
	metric.WithUnit("{pod}"),
}

func NewJobPodDesiredSuccessfulObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (JobPodDesiredSuccessfulObservable, error) {
	_ = "STUB: not implemented"
	return *new(JobPodDesiredSuccessfulObservable), nil
}

func (m JobPodDesiredSuccessfulObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (JobPodDesiredSuccessfulObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodDesiredSuccessfulObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodDesiredSuccessfulObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type JobPodFailedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newJobPodFailedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of pods which reached phase Failed for a job."),
	metric.WithUnit("{pod}"),
}

func NewJobPodFailedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (JobPodFailedObservable, error) {
	_ = "STUB: not implemented"
	return *new(JobPodFailedObservable), nil
}

func (m JobPodFailedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (JobPodFailedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodFailedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodFailedObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type JobPodMaxParallelObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newJobPodMaxParallelObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The max desired number of pods the job should run at any given time."),
	metric.WithUnit("{pod}"),
}

func NewJobPodMaxParallelObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (JobPodMaxParallelObservable, error) {
	_ = "STUB: not implemented"
	return *new(JobPodMaxParallelObservable), nil
}

func (m JobPodMaxParallelObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (JobPodMaxParallelObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodMaxParallelObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodMaxParallelObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type JobPodSuccessfulObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newJobPodSuccessfulObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of pods which reached phase Succeeded for a job."),
	metric.WithUnit("{pod}"),
}

func NewJobPodSuccessfulObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (JobPodSuccessfulObservable, error) {
	_ = "STUB: not implemented"
	return *new(JobPodSuccessfulObservable), nil
}

func (m JobPodSuccessfulObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (JobPodSuccessfulObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (JobPodSuccessfulObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (JobPodSuccessfulObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NamespacePhaseObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNamespacePhaseObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Describes number of K8s namespaces that are currently in a given phase."),
	metric.WithUnit("{namespace}"),
}

func NewNamespacePhaseObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NamespacePhaseObservable, error) {
	_ = "STUB: not implemented"
	return *new(NamespacePhaseObservable), nil
}

func (m NamespacePhaseObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NamespacePhaseObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NamespacePhaseObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NamespacePhaseObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NamespacePhaseObservable) AttrNamespacePhase(val NamespacePhaseAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type NodeConditionStatusObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeConditionStatusObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Describes the condition of a particular Node."),
	metric.WithUnit("{node}"),
}

func NewNodeConditionStatusObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeConditionStatusObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeConditionStatusObservable), nil
}

func (m NodeConditionStatusObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeConditionStatusObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeConditionStatusObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeConditionStatusObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NodeConditionStatusObservable) AttrNodeConditionStatus(val NodeConditionStatusAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NodeConditionStatusObservable) AttrNodeConditionType(val NodeConditionTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type NodeCPUAllocatableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeCPUAllocatableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Amount of cpu allocatable on the node."),
	metric.WithUnit("{cpu}"),
}

func NewNodeCPUAllocatableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeCPUAllocatableObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeCPUAllocatableObservable), nil
}

func (m NodeCPUAllocatableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeCPUAllocatableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUAllocatableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUAllocatableObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeCPUTimeObservable struct {
	metric.Float64ObservableCounter
}

var newNodeCPUTimeObservableOpts = []metric.Float64ObservableCounterOption{
	metric.WithDescription("Total CPU time consumed."),
	metric.WithUnit("s"),
}

func NewNodeCPUTimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableCounterOption,
) (NodeCPUTimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeCPUTimeObservable), nil
}

func (m NodeCPUTimeObservable) Inst() metric.Float64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter)
}

func (NodeCPUTimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUTimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUTimeObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeCPUUsageObservable struct {
	metric.Int64ObservableGauge
}

var newNodeCPUUsageObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Node's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs."),
	metric.WithUnit("{cpu}"),
}

func NewNodeCPUUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (NodeCPUUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeCPUUsageObservable), nil
}

func (m NodeCPUUsageObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (NodeCPUUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeCPUUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeEphemeralStorageAllocatableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeEphemeralStorageAllocatableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Amount of ephemeral-storage allocatable on the node."),
	metric.WithUnit("By"),
}

func NewNodeEphemeralStorageAllocatableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeEphemeralStorageAllocatableObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeEphemeralStorageAllocatableObservable), nil
}

func (m NodeEphemeralStorageAllocatableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeEphemeralStorageAllocatableObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (NodeEphemeralStorageAllocatableObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (NodeEphemeralStorageAllocatableObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type NodeFilesystemAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeFilesystemAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Node filesystem available bytes."),
	metric.WithUnit("By"),
}

func NewNodeFilesystemAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeFilesystemAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeFilesystemAvailableObservable), nil
}

func (m NodeFilesystemAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeFilesystemAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeFilesystemCapacityObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeFilesystemCapacityObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Node filesystem capacity."),
	metric.WithUnit("By"),
}

func NewNodeFilesystemCapacityObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeFilesystemCapacityObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeFilesystemCapacityObservable), nil
}

func (m NodeFilesystemCapacityObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeFilesystemCapacityObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemCapacityObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemCapacityObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeFilesystemUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeFilesystemUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Node filesystem usage."),
	metric.WithUnit("By"),
}

func NewNodeFilesystemUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeFilesystemUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeFilesystemUsageObservable), nil
}

func (m NodeFilesystemUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeFilesystemUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeFilesystemUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeMemoryAllocatableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeMemoryAllocatableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Amount of memory allocatable on the node."),
	metric.WithUnit("By"),
}

func NewNodeMemoryAllocatableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeMemoryAllocatableObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryAllocatableObservable), nil
}

func (m NodeMemoryAllocatableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeMemoryAllocatableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryAllocatableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryAllocatableObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeMemoryAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeMemoryAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Node memory available."),
	metric.WithUnit("By"),
}

func NewNodeMemoryAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeMemoryAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryAvailableObservable), nil
}

func (m NodeMemoryAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeMemoryAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeMemoryPagingFaultsObservable struct {
	metric.Int64ObservableCounter
}

var newNodeMemoryPagingFaultsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Node memory paging faults."),
	metric.WithUnit("{fault}"),
}

func NewNodeMemoryPagingFaultsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (NodeMemoryPagingFaultsObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryPagingFaultsObservable), nil
}

func (m NodeMemoryPagingFaultsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (NodeMemoryPagingFaultsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryPagingFaultsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryPagingFaultsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryPagingFaultsObservable) AttrSystemPagingFaultType(val SystemPagingFaultTypeAttr) attribute.KeyValue {
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

type NodeMemoryRssObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeMemoryRssObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Node memory RSS."),
	metric.WithUnit("By"),
}

func NewNodeMemoryRssObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeMemoryRssObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryRssObservable), nil
}

func (m NodeMemoryRssObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeMemoryRssObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryRssObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryRssObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeMemoryUsageObservable struct {
	metric.Int64ObservableGauge
}

var newNodeMemoryUsageObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Memory usage of the Node."),
	metric.WithUnit("By"),
}

func NewNodeMemoryUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (NodeMemoryUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryUsageObservable), nil
}

func (m NodeMemoryUsageObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (NodeMemoryUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeMemoryWorkingSetObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeMemoryWorkingSetObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Node memory working set."),
	metric.WithUnit("By"),
}

func NewNodeMemoryWorkingSetObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeMemoryWorkingSetObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeMemoryWorkingSetObservable), nil
}

func (m NodeMemoryWorkingSetObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeMemoryWorkingSetObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryWorkingSetObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeMemoryWorkingSetObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type NodeNetworkErrorsObservable struct {
	metric.Int64ObservableCounter
}

var newNodeNetworkErrorsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Node network errors."),
	metric.WithUnit("{error}"),
}

func NewNodeNetworkErrorsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (NodeNetworkErrorsObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeNetworkErrorsObservable), nil
}

func (m NodeNetworkErrorsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (NodeNetworkErrorsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkErrorsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkErrorsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkErrorsObservable) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NodeNetworkErrorsObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
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

type NodeNetworkIOObservable struct {
	metric.Int64ObservableCounter
}

var newNodeNetworkIOObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Network bytes for the Node."),
	metric.WithUnit("By"),
}

func NewNodeNetworkIOObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (NodeNetworkIOObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeNetworkIOObservable), nil
}

func (m NodeNetworkIOObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (NodeNetworkIOObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkIOObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkIOObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (NodeNetworkIOObservable) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (NodeNetworkIOObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
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

type NodePodAllocatableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodePodAllocatableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Amount of pods allocatable on the node."),
	metric.WithUnit("{pod}"),
}

func NewNodePodAllocatableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodePodAllocatableObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodePodAllocatableObservable), nil
}

func (m NodePodAllocatableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodePodAllocatableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodePodAllocatableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodePodAllocatableObservable) Description() string { _ = "STUB: not implemented"; return "" }

type NodeSystemContainerCPUTime struct {
	metric.Float64Counter
}

var newNodeSystemContainerCPUTimeOpts = []metric.Float64CounterOption{
	metric.WithDescription("Node's system container CPU time."),
	metric.WithUnit("s"),
}

func NewNodeSystemContainerCPUTime(
	m metric.Meter,
	opt ...metric.Float64CounterOption,
) (NodeSystemContainerCPUTime, error) {
	_ = "STUB: not implemented"
	return *new(NodeSystemContainerCPUTime), nil
}

func (m NodeSystemContainerCPUTime) Inst() metric.Float64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Float64Counter)
}

func (NodeSystemContainerCPUTime) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerCPUTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerCPUTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeSystemContainerCPUTime) Add(ctx context.Context, incr float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeSystemContainerCPUTime) AddSet(ctx context.Context, incr float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeSystemContainerCPUTimeObservable struct {
	metric.Float64ObservableCounter
}

var newNodeSystemContainerCPUTimeObservableOpts = []metric.Float64ObservableCounterOption{
	metric.WithDescription("Node's system container CPU time."),
	metric.WithUnit("s"),
}

func NewNodeSystemContainerCPUTimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableCounterOption,
) (NodeSystemContainerCPUTimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeSystemContainerCPUTimeObservable), nil
}

func (m NodeSystemContainerCPUTimeObservable) Inst() metric.Float64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter)
}

func (NodeSystemContainerCPUTimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerCPUTimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerCPUTimeObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type NodeSystemContainerCPUUsage struct {
	metric.Int64Gauge
}

var newNodeSystemContainerCPUUsageOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Node's system container CPU usage, measured in cpus."),
	metric.WithUnit("{cpu}"),
}

func NewNodeSystemContainerCPUUsage(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (NodeSystemContainerCPUUsage, error) {
	_ = "STUB: not implemented"
	return *new(NodeSystemContainerCPUUsage), nil
}

func (m NodeSystemContainerCPUUsage) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (NodeSystemContainerCPUUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerCPUUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerCPUUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeSystemContainerCPUUsage) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeSystemContainerCPUUsage) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeSystemContainerCPUUsageObservable struct {
	metric.Int64ObservableGauge
}

var newNodeSystemContainerCPUUsageObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Node's system container CPU usage, measured in cpus."),
	metric.WithUnit("{cpu}"),
}

func NewNodeSystemContainerCPUUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (NodeSystemContainerCPUUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeSystemContainerCPUUsageObservable), nil
}

func (m NodeSystemContainerCPUUsageObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (NodeSystemContainerCPUUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerCPUUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerCPUUsageObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type NodeSystemContainerMemoryUsage struct {
	metric.Int64UpDownCounter
}

var newNodeSystemContainerMemoryUsageOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Node's system container memory usage."),
	metric.WithUnit("By"),
}

func NewNodeSystemContainerMemoryUsage(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeSystemContainerMemoryUsage, error) {
	_ = "STUB: not implemented"
	return *new(NodeSystemContainerMemoryUsage), nil
}

func (m NodeSystemContainerMemoryUsage) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeSystemContainerMemoryUsage) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerMemoryUsage) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerMemoryUsage) Description() string { _ = "STUB: not implemented"; return "" }

func (m NodeSystemContainerMemoryUsage) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeSystemContainerMemoryUsage) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeSystemContainerMemoryUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeSystemContainerMemoryUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Node's system container memory usage."),
	metric.WithUnit("By"),
}

func NewNodeSystemContainerMemoryUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeSystemContainerMemoryUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeSystemContainerMemoryUsageObservable), nil
}

func (m NodeSystemContainerMemoryUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeSystemContainerMemoryUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerMemoryUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerMemoryUsageObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type NodeSystemContainerMemoryWorkingSet struct {
	metric.Int64UpDownCounter
}

var newNodeSystemContainerMemoryWorkingSetOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The amount of working set memory."),
	metric.WithUnit("By"),
}

func NewNodeSystemContainerMemoryWorkingSet(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (NodeSystemContainerMemoryWorkingSet, error) {
	_ = "STUB: not implemented"
	return *new(NodeSystemContainerMemoryWorkingSet), nil
}

func (m NodeSystemContainerMemoryWorkingSet) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (NodeSystemContainerMemoryWorkingSet) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerMemoryWorkingSet) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeSystemContainerMemoryWorkingSet) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m NodeSystemContainerMemoryWorkingSet) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m NodeSystemContainerMemoryWorkingSet) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type NodeSystemContainerMemoryWorkingSetObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newNodeSystemContainerMemoryWorkingSetObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The amount of working set memory."),
	metric.WithUnit("By"),
}

func NewNodeSystemContainerMemoryWorkingSetObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (NodeSystemContainerMemoryWorkingSetObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeSystemContainerMemoryWorkingSetObservable), nil
}

func (m NodeSystemContainerMemoryWorkingSetObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (NodeSystemContainerMemoryWorkingSetObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (NodeSystemContainerMemoryWorkingSetObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (NodeSystemContainerMemoryWorkingSetObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type NodeUptimeObservable struct {
	metric.Float64ObservableGauge
}

var newNodeUptimeObservableOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("The time the Node has been running."),
	metric.WithUnit("s"),
}

func NewNodeUptimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (NodeUptimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(NodeUptimeObservable), nil
}

func (m NodeUptimeObservable) Inst() metric.Float64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge)
}

func (NodeUptimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (NodeUptimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (NodeUptimeObservable) Description() string { _ = "STUB: not implemented"; return "" }

type PersistentvolumeStatusPhase struct {
	metric.Int64UpDownCounter
}

var newPersistentvolumeStatusPhaseOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of PersistentVolumes in a given phase."),
	metric.WithUnit("{persistentvolume}"),
}

func NewPersistentvolumeStatusPhase(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PersistentvolumeStatusPhase, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeStatusPhase), nil
}

func (m PersistentvolumeStatusPhase) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PersistentvolumeStatusPhase) Name() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeStatusPhase) Unit() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeStatusPhase) Description() string { _ = "STUB: not implemented"; return "" }

func (m PersistentvolumeStatusPhase) Add(
	ctx context.Context,
	incr int64,
	persistentvolumeStatusPhase PersistentvolumeStatusPhaseAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PersistentvolumeStatusPhase) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PersistentvolumeStatusPhaseObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPersistentvolumeStatusPhaseObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of PersistentVolumes in a given phase."),
	metric.WithUnit("{persistentvolume}"),
}

func NewPersistentvolumeStatusPhaseObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PersistentvolumeStatusPhaseObservable, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeStatusPhaseObservable), nil
}

func (m PersistentvolumeStatusPhaseObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PersistentvolumeStatusPhaseObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeStatusPhaseObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeStatusPhaseObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeStatusPhaseObservable) AttrPersistentvolumeStatusPhase(val PersistentvolumeStatusPhaseAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PersistentvolumeStorageCapacity struct {
	metric.Int64UpDownCounter
}

var newPersistentvolumeStorageCapacityOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The storage capacity of the PersistentVolume."),
	metric.WithUnit("By"),
}

func NewPersistentvolumeStorageCapacity(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PersistentvolumeStorageCapacity, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeStorageCapacity), nil
}

func (m PersistentvolumeStorageCapacity) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PersistentvolumeStorageCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeStorageCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeStorageCapacity) Description() string { _ = "STUB: not implemented"; return "" }

func (m PersistentvolumeStorageCapacity) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m PersistentvolumeStorageCapacity) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PersistentvolumeStorageCapacityObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPersistentvolumeStorageCapacityObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The storage capacity of the PersistentVolume."),
	metric.WithUnit("By"),
}

func NewPersistentvolumeStorageCapacityObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PersistentvolumeStorageCapacityObservable, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeStorageCapacityObservable), nil
}

func (m PersistentvolumeStorageCapacityObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PersistentvolumeStorageCapacityObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeStorageCapacityObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeStorageCapacityObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type PersistentvolumeclaimStatusPhase struct {
	metric.Int64UpDownCounter
}

var newPersistentvolumeclaimStatusPhaseOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("Number of PersistentVolumeClaims in a given phase."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewPersistentvolumeclaimStatusPhase(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PersistentvolumeclaimStatusPhase, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeclaimStatusPhase), nil
}

func (m PersistentvolumeclaimStatusPhase) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PersistentvolumeclaimStatusPhase) Name() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeclaimStatusPhase) Unit() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeclaimStatusPhase) Description() string { _ = "STUB: not implemented"; return "" }

func (m PersistentvolumeclaimStatusPhase) Add(
	ctx context.Context,
	incr int64,
	persistentvolumeclaimStatusPhase PersistentvolumeclaimStatusPhaseAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m PersistentvolumeclaimStatusPhase) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PersistentvolumeclaimStatusPhaseObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPersistentvolumeclaimStatusPhaseObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of PersistentVolumeClaims in a given phase."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewPersistentvolumeclaimStatusPhaseObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PersistentvolumeclaimStatusPhaseObservable, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeclaimStatusPhaseObservable), nil
}

func (m PersistentvolumeclaimStatusPhaseObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PersistentvolumeclaimStatusPhaseObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeclaimStatusPhaseObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeclaimStatusPhaseObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeclaimStatusPhaseObservable) AttrPersistentvolumeclaimStatusPhase(val PersistentvolumeclaimStatusPhaseAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type PersistentvolumeclaimStorageCapacity struct {
	metric.Int64UpDownCounter
}

var newPersistentvolumeclaimStorageCapacityOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The actual storage capacity provisioned for the PersistentVolumeClaim."),
	metric.WithUnit("By"),
}

func NewPersistentvolumeclaimStorageCapacity(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PersistentvolumeclaimStorageCapacity, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeclaimStorageCapacity), nil
}

func (m PersistentvolumeclaimStorageCapacity) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PersistentvolumeclaimStorageCapacity) Name() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeclaimStorageCapacity) Unit() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeclaimStorageCapacity) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m PersistentvolumeclaimStorageCapacity) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m PersistentvolumeclaimStorageCapacity) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PersistentvolumeclaimStorageCapacityObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPersistentvolumeclaimStorageCapacityObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The actual storage capacity provisioned for the PersistentVolumeClaim."),
	metric.WithUnit("By"),
}

func NewPersistentvolumeclaimStorageCapacityObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PersistentvolumeclaimStorageCapacityObservable, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeclaimStorageCapacityObservable), nil
}

func (m PersistentvolumeclaimStorageCapacityObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PersistentvolumeclaimStorageCapacityObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeclaimStorageCapacityObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeclaimStorageCapacityObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

type PersistentvolumeclaimStorageRequest struct {
	metric.Int64UpDownCounter
}

var newPersistentvolumeclaimStorageRequestOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The storage requested by the PersistentVolumeClaim."),
	metric.WithUnit("By"),
}

func NewPersistentvolumeclaimStorageRequest(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (PersistentvolumeclaimStorageRequest, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeclaimStorageRequest), nil
}

func (m PersistentvolumeclaimStorageRequest) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (PersistentvolumeclaimStorageRequest) Name() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeclaimStorageRequest) Unit() string { _ = "STUB: not implemented"; return "" }

func (PersistentvolumeclaimStorageRequest) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m PersistentvolumeclaimStorageRequest) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m PersistentvolumeclaimStorageRequest) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type PersistentvolumeclaimStorageRequestObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPersistentvolumeclaimStorageRequestObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The storage requested by the PersistentVolumeClaim."),
	metric.WithUnit("By"),
}

func NewPersistentvolumeclaimStorageRequestObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PersistentvolumeclaimStorageRequestObservable, error) {
	_ = "STUB: not implemented"
	return *new(PersistentvolumeclaimStorageRequestObservable), nil
}

func (m PersistentvolumeclaimStorageRequestObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PersistentvolumeclaimStorageRequestObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeclaimStorageRequestObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (PersistentvolumeclaimStorageRequestObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type PodCPUTimeObservable struct {
	metric.Float64ObservableCounter
}

var newPodCPUTimeObservableOpts = []metric.Float64ObservableCounterOption{
	metric.WithDescription("Total CPU time consumed."),
	metric.WithUnit("s"),
}

func NewPodCPUTimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableCounterOption,
) (PodCPUTimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodCPUTimeObservable), nil
}

func (m PodCPUTimeObservable) Inst() metric.Float64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableCounter)
}

func (PodCPUTimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodCPUTimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodCPUTimeObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodCPUUsageObservable struct {
	metric.Int64ObservableGauge
}

var newPodCPUUsageObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Pod's CPU usage, measured in cpus. Range from 0 to the number of allocatable CPUs."),
	metric.WithUnit("{cpu}"),
}

func NewPodCPUUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (PodCPUUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodCPUUsageObservable), nil
}

func (m PodCPUUsageObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (PodCPUUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodCPUUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodCPUUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodFilesystemAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodFilesystemAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Pod filesystem available bytes."),
	metric.WithUnit("By"),
}

func NewPodFilesystemAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodFilesystemAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodFilesystemAvailableObservable), nil
}

func (m PodFilesystemAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodFilesystemAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodFilesystemCapacityObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodFilesystemCapacityObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Pod filesystem capacity."),
	metric.WithUnit("By"),
}

func NewPodFilesystemCapacityObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodFilesystemCapacityObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodFilesystemCapacityObservable), nil
}

func (m PodFilesystemCapacityObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodFilesystemCapacityObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemCapacityObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemCapacityObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodFilesystemUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodFilesystemUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Pod filesystem usage."),
	metric.WithUnit("By"),
}

func NewPodFilesystemUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodFilesystemUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodFilesystemUsageObservable), nil
}

func (m PodFilesystemUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodFilesystemUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodFilesystemUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodMemoryAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodMemoryAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Pod memory available."),
	metric.WithUnit("By"),
}

func NewPodMemoryAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodMemoryAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryAvailableObservable), nil
}

func (m PodMemoryAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodMemoryAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodMemoryPagingFaultsObservable struct {
	metric.Int64ObservableCounter
}

var newPodMemoryPagingFaultsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Pod memory paging faults."),
	metric.WithUnit("{fault}"),
}

func NewPodMemoryPagingFaultsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (PodMemoryPagingFaultsObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryPagingFaultsObservable), nil
}

func (m PodMemoryPagingFaultsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (PodMemoryPagingFaultsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryPagingFaultsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryPagingFaultsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryPagingFaultsObservable) AttrSystemPagingFaultType(val SystemPagingFaultTypeAttr) attribute.KeyValue {
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

type PodMemoryRssObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodMemoryRssObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Pod memory RSS."),
	metric.WithUnit("By"),
}

func NewPodMemoryRssObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodMemoryRssObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryRssObservable), nil
}

func (m PodMemoryRssObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodMemoryRssObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryRssObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryRssObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodMemoryUsageObservable struct {
	metric.Int64ObservableGauge
}

var newPodMemoryUsageObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Memory usage of the Pod."),
	metric.WithUnit("By"),
}

func NewPodMemoryUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (PodMemoryUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryUsageObservable), nil
}

func (m PodMemoryUsageObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (PodMemoryUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodMemoryWorkingSetObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodMemoryWorkingSetObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Pod memory working set."),
	metric.WithUnit("By"),
}

func NewPodMemoryWorkingSetObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodMemoryWorkingSetObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodMemoryWorkingSetObservable), nil
}

func (m PodMemoryWorkingSetObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodMemoryWorkingSetObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryWorkingSetObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodMemoryWorkingSetObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodNetworkErrorsObservable struct {
	metric.Int64ObservableCounter
}

var newPodNetworkErrorsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Pod network errors."),
	metric.WithUnit("{error}"),
}

func NewPodNetworkErrorsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (PodNetworkErrorsObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodNetworkErrorsObservable), nil
}

func (m PodNetworkErrorsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (PodNetworkErrorsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkErrorsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkErrorsObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkErrorsObservable) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodNetworkErrorsObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
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

type PodNetworkIOObservable struct {
	metric.Int64ObservableCounter
}

var newPodNetworkIOObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("Network bytes for the Pod."),
	metric.WithUnit("By"),
}

func NewPodNetworkIOObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (PodNetworkIOObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodNetworkIOObservable), nil
}

func (m PodNetworkIOObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (PodNetworkIOObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkIOObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkIOObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodNetworkIOObservable) AttrNetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodNetworkIOObservable) AttrNetworkIODirection(val NetworkIODirectionAttr) attribute.KeyValue {
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

type PodStatusPhaseObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodStatusPhaseObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Describes number of K8s Pods that are currently in a given phase."),
	metric.WithUnit("{pod}"),
}

func NewPodStatusPhaseObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodStatusPhaseObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodStatusPhaseObservable), nil
}

func (m PodStatusPhaseObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodStatusPhaseObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodStatusPhaseObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodStatusPhaseObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodStatusPhaseObservable) AttrPodStatusPhase(val PodStatusPhaseAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type PodStatusReasonObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodStatusReasonObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Describes the number of K8s Pods that are currently in a state for a given reason."),
	metric.WithUnit("{pod}"),
}

func NewPodStatusReasonObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodStatusReasonObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodStatusReasonObservable), nil
}

func (m PodStatusReasonObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodStatusReasonObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodStatusReasonObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodStatusReasonObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodStatusReasonObservable) AttrPodStatusReason(val PodStatusReasonAttr) attribute.KeyValue {
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

type PodUptimeObservable struct {
	metric.Float64ObservableGauge
}

var newPodUptimeObservableOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("The time the Pod has been running."),
	metric.WithUnit("s"),
}

func NewPodUptimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (PodUptimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodUptimeObservable), nil
}

func (m PodUptimeObservable) Inst() metric.Float64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge)
}

func (PodUptimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodUptimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodUptimeObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type PodVolumeAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodVolumeAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Pod volume storage space available."),
	metric.WithUnit("By"),
}

func NewPodVolumeAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodVolumeAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeAvailableObservable), nil
}

func (m PodVolumeAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodVolumeAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeAvailableObservable) AttrVolumeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodVolumeAvailableObservable) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
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

type PodVolumeCapacityObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodVolumeCapacityObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Pod volume total capacity."),
	metric.WithUnit("By"),
}

func NewPodVolumeCapacityObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodVolumeCapacityObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeCapacityObservable), nil
}

func (m PodVolumeCapacityObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodVolumeCapacityObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeCapacityObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeCapacityObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeCapacityObservable) AttrVolumeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodVolumeCapacityObservable) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
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

type PodVolumeInodeCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodVolumeInodeCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The total inodes in the filesystem of the Pod's volume."),
	metric.WithUnit("{inode}"),
}

func NewPodVolumeInodeCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodVolumeInodeCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeInodeCountObservable), nil
}

func (m PodVolumeInodeCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodVolumeInodeCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeCountObservable) AttrVolumeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodVolumeInodeCountObservable) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
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

type PodVolumeInodeFreeObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodVolumeInodeFreeObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The free inodes in the filesystem of the Pod's volume."),
	metric.WithUnit("{inode}"),
}

func NewPodVolumeInodeFreeObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodVolumeInodeFreeObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeInodeFreeObservable), nil
}

func (m PodVolumeInodeFreeObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodVolumeInodeFreeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeFreeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeFreeObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeFreeObservable) AttrVolumeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodVolumeInodeFreeObservable) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
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

type PodVolumeInodeUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodVolumeInodeUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The inodes used by the filesystem of the Pod's volume."),
	metric.WithUnit("{inode}"),
}

func NewPodVolumeInodeUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodVolumeInodeUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeInodeUsedObservable), nil
}

func (m PodVolumeInodeUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodVolumeInodeUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeUsedObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeInodeUsedObservable) AttrVolumeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodVolumeInodeUsedObservable) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
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

type PodVolumeUsageObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newPodVolumeUsageObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Pod volume usage."),
	metric.WithUnit("By"),
}

func NewPodVolumeUsageObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (PodVolumeUsageObservable, error) {
	_ = "STUB: not implemented"
	return *new(PodVolumeUsageObservable), nil
}

func (m PodVolumeUsageObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (PodVolumeUsageObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeUsageObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeUsageObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (PodVolumeUsageObservable) AttrVolumeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (PodVolumeUsageObservable) AttrVolumeType(val VolumeTypeAttr) attribute.KeyValue {
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

type ReplicaSetPodAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newReplicaSetPodAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this replicaset."),
	metric.WithUnit("{pod}"),
}

func NewReplicaSetPodAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ReplicaSetPodAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(ReplicaSetPodAvailableObservable), nil
}

func (m ReplicaSetPodAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ReplicaSetPodAvailableObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetPodAvailableObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetPodAvailableObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type ReplicaSetPodDesiredObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newReplicaSetPodDesiredObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this replicaset."),
	metric.WithUnit("{pod}"),
}

func NewReplicaSetPodDesiredObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ReplicaSetPodDesiredObservable, error) {
	_ = "STUB: not implemented"
	return *new(ReplicaSetPodDesiredObservable), nil
}

func (m ReplicaSetPodDesiredObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ReplicaSetPodDesiredObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetPodDesiredObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ReplicaSetPodDesiredObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type ReplicationControllerPodAvailableObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newReplicationControllerPodAvailableObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Total number of available replica pods (ready for at least minReadySeconds) targeted by this replication controller."),
	metric.WithUnit("{pod}"),
}

func NewReplicationControllerPodAvailableObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ReplicationControllerPodAvailableObservable, error) {
	_ = "STUB: not implemented"
	return *new(ReplicationControllerPodAvailableObservable), nil
}

func (m ReplicationControllerPodAvailableObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ReplicationControllerPodAvailableObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ReplicationControllerPodAvailableObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ReplicationControllerPodAvailableObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ReplicationControllerPodDesiredObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newReplicationControllerPodDesiredObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this replication controller."),
	metric.WithUnit("{pod}"),
}

func NewReplicationControllerPodDesiredObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ReplicationControllerPodDesiredObservable, error) {
	_ = "STUB: not implemented"
	return *new(ReplicationControllerPodDesiredObservable), nil
}

func (m ReplicationControllerPodDesiredObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ReplicationControllerPodDesiredObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ReplicationControllerPodDesiredObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ReplicationControllerPodDesiredObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaCPULimitHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaCPULimitHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The CPU limits in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{cpu}"),
}

func NewResourceQuotaCPULimitHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaCPULimitHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaCPULimitHardObservable), nil
}

func (m ResourceQuotaCPULimitHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaCPULimitHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPULimitHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPULimitHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaCPULimitUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaCPULimitUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The CPU limits in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{cpu}"),
}

func NewResourceQuotaCPULimitUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaCPULimitUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaCPULimitUsedObservable), nil
}

func (m ResourceQuotaCPULimitUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaCPULimitUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPULimitUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPULimitUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaCPURequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaCPURequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The CPU requests in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{cpu}"),
}

func NewResourceQuotaCPURequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaCPURequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaCPURequestHardObservable), nil
}

func (m ResourceQuotaCPURequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaCPURequestHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPURequestHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPURequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaCPURequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaCPURequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The CPU requests in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{cpu}"),
}

func NewResourceQuotaCPURequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaCPURequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaCPURequestUsedObservable), nil
}

func (m ResourceQuotaCPURequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaCPURequestUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPURequestUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaCPURequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaEphemeralStorageLimitHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaEphemeralStorageLimitHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The sum of local ephemeral storage limits in the namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaEphemeralStorageLimitHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaEphemeralStorageLimitHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaEphemeralStorageLimitHardObservable), nil
}

func (m ResourceQuotaEphemeralStorageLimitHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaEphemeralStorageLimitHardObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaEphemeralStorageLimitHardObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaEphemeralStorageLimitHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaEphemeralStorageLimitUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaEphemeralStorageLimitUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The sum of local ephemeral storage limits in the namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaEphemeralStorageLimitUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaEphemeralStorageLimitUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaEphemeralStorageLimitUsedObservable), nil
}

func (m ResourceQuotaEphemeralStorageLimitUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaEphemeralStorageLimitUsedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaEphemeralStorageLimitUsedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaEphemeralStorageLimitUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaEphemeralStorageRequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaEphemeralStorageRequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The sum of local ephemeral storage requests in the namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaEphemeralStorageRequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaEphemeralStorageRequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaEphemeralStorageRequestHardObservable), nil
}

func (m ResourceQuotaEphemeralStorageRequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaEphemeralStorageRequestHardObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaEphemeralStorageRequestHardObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaEphemeralStorageRequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaEphemeralStorageRequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaEphemeralStorageRequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The sum of local ephemeral storage requests in the namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaEphemeralStorageRequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaEphemeralStorageRequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaEphemeralStorageRequestUsedObservable), nil
}

func (m ResourceQuotaEphemeralStorageRequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaEphemeralStorageRequestUsedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaEphemeralStorageRequestUsedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaEphemeralStorageRequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaHugepageCountRequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaHugepageCountRequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The huge page requests in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{hugepage}"),
}

func NewResourceQuotaHugepageCountRequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaHugepageCountRequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaHugepageCountRequestHardObservable), nil
}

func (m ResourceQuotaHugepageCountRequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaHugepageCountRequestHardObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaHugepageCountRequestHardObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaHugepageCountRequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaHugepageCountRequestHardObservable) AttrHugepageSize(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ResourceQuotaHugepageCountRequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaHugepageCountRequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The huge page requests in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{hugepage}"),
}

func NewResourceQuotaHugepageCountRequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaHugepageCountRequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaHugepageCountRequestUsedObservable), nil
}

func (m ResourceQuotaHugepageCountRequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaHugepageCountRequestUsedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaHugepageCountRequestUsedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaHugepageCountRequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaHugepageCountRequestUsedObservable) AttrHugepageSize(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ResourceQuotaMemoryLimitHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaMemoryLimitHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The memory limits in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaMemoryLimitHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaMemoryLimitHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaMemoryLimitHardObservable), nil
}

func (m ResourceQuotaMemoryLimitHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaMemoryLimitHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryLimitHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryLimitHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaMemoryLimitUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaMemoryLimitUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The memory limits in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaMemoryLimitUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaMemoryLimitUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaMemoryLimitUsedObservable), nil
}

func (m ResourceQuotaMemoryLimitUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaMemoryLimitUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryLimitUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryLimitUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaMemoryRequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaMemoryRequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The memory requests in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaMemoryRequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaMemoryRequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaMemoryRequestHardObservable), nil
}

func (m ResourceQuotaMemoryRequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaMemoryRequestHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryRequestHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryRequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaMemoryRequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaMemoryRequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The memory requests in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaMemoryRequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaMemoryRequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaMemoryRequestUsedObservable), nil
}

func (m ResourceQuotaMemoryRequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaMemoryRequestUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryRequestUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaMemoryRequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ResourceQuotaObjectCountHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaObjectCountHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The object count limits in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{object}"),
}

func NewResourceQuotaObjectCountHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaObjectCountHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaObjectCountHardObservable), nil
}

func (m ResourceQuotaObjectCountHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaObjectCountHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaObjectCountHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaObjectCountHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaObjectCountHardObservable) AttrResourceQuotaResourceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ResourceQuotaObjectCountUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaObjectCountUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The object count limits in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{object}"),
}

func NewResourceQuotaObjectCountUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaObjectCountUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaObjectCountUsedObservable), nil
}

func (m ResourceQuotaObjectCountUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaObjectCountUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaObjectCountUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ResourceQuotaObjectCountUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaObjectCountUsedObservable) AttrResourceQuotaResourceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ResourceQuotaPersistentvolumeclaimCountHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaPersistentvolumeclaimCountHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The total number of PersistentVolumeClaims that can exist in the namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewResourceQuotaPersistentvolumeclaimCountHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaPersistentvolumeclaimCountHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaPersistentvolumeclaimCountHardObservable), nil
}

func (m ResourceQuotaPersistentvolumeclaimCountHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaPersistentvolumeclaimCountHardObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountHardObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountHardObservable) AttrStorageclassName(val string) attribute.KeyValue {
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

type ResourceQuotaPersistentvolumeclaimCountUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaPersistentvolumeclaimCountUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The total number of PersistentVolumeClaims that can exist in the namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewResourceQuotaPersistentvolumeclaimCountUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaPersistentvolumeclaimCountUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaPersistentvolumeclaimCountUsedObservable), nil
}

func (m ResourceQuotaPersistentvolumeclaimCountUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaPersistentvolumeclaimCountUsedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountUsedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaPersistentvolumeclaimCountUsedObservable) AttrStorageclassName(val string) attribute.KeyValue {
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

type ResourceQuotaStorageRequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaStorageRequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The storage requests in a specific namespace. The value represents the configured quota limit of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaStorageRequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaStorageRequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaStorageRequestHardObservable), nil
}

func (m ResourceQuotaStorageRequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaStorageRequestHardObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaStorageRequestHardObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaStorageRequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaStorageRequestHardObservable) AttrStorageclassName(val string) attribute.KeyValue {
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

type ResourceQuotaStorageRequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newResourceQuotaStorageRequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The storage requests in a specific namespace. The value represents the current observed total usage of the resource in the namespace."),
	metric.WithUnit("By"),
}

func NewResourceQuotaStorageRequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ResourceQuotaStorageRequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ResourceQuotaStorageRequestUsedObservable), nil
}

func (m ResourceQuotaStorageRequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ResourceQuotaStorageRequestUsedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaStorageRequestUsedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaStorageRequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ResourceQuotaStorageRequestUsedObservable) AttrStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServiceEndpointCount struct {
	metric.Int64Gauge
}

var newServiceEndpointCountOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Number of endpoints for a service by condition and address type."),
	metric.WithUnit("{endpoint}"),
}

func NewServiceEndpointCount(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (ServiceEndpointCount, error) {
	_ = "STUB: not implemented"
	return *new(ServiceEndpointCount), nil
}

func (m ServiceEndpointCount) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (ServiceEndpointCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ServiceEndpointCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServiceEndpointCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServiceEndpointCount) Record(
	ctx context.Context,
	val int64,
	serviceEndpointAddressType ServiceEndpointAddressTypeAttr,
	serviceEndpointCondition ServiceEndpointConditionAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ServiceEndpointCount) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ServiceEndpointCount) AttrServiceEndpointZone(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServiceEndpointCountObservable struct {
	metric.Int64ObservableGauge
}

var newServiceEndpointCountObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Number of endpoints for a service by condition and address type."),
	metric.WithUnit("{endpoint}"),
}

func NewServiceEndpointCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (ServiceEndpointCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(ServiceEndpointCountObservable), nil
}

func (m ServiceEndpointCountObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (ServiceEndpointCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ServiceEndpointCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServiceEndpointCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ServiceEndpointCountObservable) AttrServiceEndpointAddressType(val ServiceEndpointAddressTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServiceEndpointCountObservable) AttrServiceEndpointCondition(val ServiceEndpointConditionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ServiceEndpointCountObservable) AttrServiceEndpointZone(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ServiceLoadBalancerIngressCount struct {
	metric.Int64Gauge
}

var newServiceLoadBalancerIngressCountOpts = []metric.Int64GaugeOption{
	metric.WithDescription("Number of load balancer ingress points (external IPs/hostnames) assigned to the service."),
	metric.WithUnit("{ingress}"),
}

func NewServiceLoadBalancerIngressCount(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (ServiceLoadBalancerIngressCount, error) {
	_ = "STUB: not implemented"
	return *new(ServiceLoadBalancerIngressCount), nil
}

func (m ServiceLoadBalancerIngressCount) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (ServiceLoadBalancerIngressCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ServiceLoadBalancerIngressCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServiceLoadBalancerIngressCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServiceLoadBalancerIngressCount) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ServiceLoadBalancerIngressCount) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ServiceLoadBalancerIngressCountObservable struct {
	metric.Int64ObservableGauge
}

var newServiceLoadBalancerIngressCountObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("Number of load balancer ingress points (external IPs/hostnames) assigned to the service."),
	metric.WithUnit("{ingress}"),
}

func NewServiceLoadBalancerIngressCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (ServiceLoadBalancerIngressCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(ServiceLoadBalancerIngressCountObservable), nil
}

func (m ServiceLoadBalancerIngressCountObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (ServiceLoadBalancerIngressCountObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ServiceLoadBalancerIngressCountObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ServiceLoadBalancerIngressCountObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type StatefulSetPodCurrentObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newStatefulSetPodCurrentObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of replica pods created by the statefulset controller from the statefulset version indicated by currentRevision."),
	metric.WithUnit("{pod}"),
}

func NewStatefulSetPodCurrentObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (StatefulSetPodCurrentObservable, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetPodCurrentObservable), nil
}

func (m StatefulSetPodCurrentObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (StatefulSetPodCurrentObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodCurrentObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodCurrentObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type StatefulSetPodDesiredObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newStatefulSetPodDesiredObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of desired replica pods in this statefulset."),
	metric.WithUnit("{pod}"),
}

func NewStatefulSetPodDesiredObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (StatefulSetPodDesiredObservable, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetPodDesiredObservable), nil
}

func (m StatefulSetPodDesiredObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (StatefulSetPodDesiredObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodDesiredObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodDesiredObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type StatefulSetPodReadyObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newStatefulSetPodReadyObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of replica pods created for this statefulset with a Ready Condition."),
	metric.WithUnit("{pod}"),
}

func NewStatefulSetPodReadyObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (StatefulSetPodReadyObservable, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetPodReadyObservable), nil
}

func (m StatefulSetPodReadyObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (StatefulSetPodReadyObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodReadyObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodReadyObservable) Description() string { _ = "STUB: not implemented"; return "" }

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

type StatefulSetPodUpdatedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newStatefulSetPodUpdatedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("Number of replica pods created by the statefulset controller from the statefulset version indicated by updateRevision."),
	metric.WithUnit("{pod}"),
}

func NewStatefulSetPodUpdatedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (StatefulSetPodUpdatedObservable, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetPodUpdatedObservable), nil
}

func (m StatefulSetPodUpdatedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (StatefulSetPodUpdatedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodUpdatedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (StatefulSetPodUpdatedObservable) Description() string { _ = "STUB: not implemented"; return "" }
