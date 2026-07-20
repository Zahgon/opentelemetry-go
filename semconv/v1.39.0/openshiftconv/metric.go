package openshiftconv

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

type ClusterquotaCPULimitHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaCPULimitHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{cpu}"),
}

func NewClusterquotaCPULimitHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaCPULimitHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaCPULimitHard), nil
}

func (m ClusterquotaCPULimitHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaCPULimitHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPULimitHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPULimitHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaCPULimitHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaCPULimitHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaCPULimitUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaCPULimitUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{cpu}"),
}

func NewClusterquotaCPULimitUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaCPULimitUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaCPULimitUsed), nil
}

func (m ClusterquotaCPULimitUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaCPULimitUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPULimitUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPULimitUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaCPULimitUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaCPULimitUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaCPURequestHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaCPURequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{cpu}"),
}

func NewClusterquotaCPURequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaCPURequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaCPURequestHard), nil
}

func (m ClusterquotaCPURequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaCPURequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPURequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPURequestHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaCPURequestHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaCPURequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaCPURequestUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaCPURequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{cpu}"),
}

func NewClusterquotaCPURequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaCPURequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaCPURequestUsed), nil
}

func (m ClusterquotaCPURequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaCPURequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPURequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPURequestUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaCPURequestUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaCPURequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaEphemeralStorageLimitHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaEphemeralStorageLimitHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaEphemeralStorageLimitHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaEphemeralStorageLimitHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaEphemeralStorageLimitHard), nil
}

func (m ClusterquotaEphemeralStorageLimitHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaEphemeralStorageLimitHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaEphemeralStorageLimitHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaEphemeralStorageLimitHard) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ClusterquotaEphemeralStorageLimitHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaEphemeralStorageLimitHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaEphemeralStorageLimitUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaEphemeralStorageLimitUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaEphemeralStorageLimitUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaEphemeralStorageLimitUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaEphemeralStorageLimitUsed), nil
}

func (m ClusterquotaEphemeralStorageLimitUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaEphemeralStorageLimitUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaEphemeralStorageLimitUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaEphemeralStorageLimitUsed) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ClusterquotaEphemeralStorageLimitUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaEphemeralStorageLimitUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaEphemeralStorageRequestHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaEphemeralStorageRequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaEphemeralStorageRequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaEphemeralStorageRequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaEphemeralStorageRequestHard), nil
}

func (m ClusterquotaEphemeralStorageRequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaEphemeralStorageRequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaEphemeralStorageRequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaEphemeralStorageRequestHard) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ClusterquotaEphemeralStorageRequestHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaEphemeralStorageRequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaEphemeralStorageRequestUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaEphemeralStorageRequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaEphemeralStorageRequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaEphemeralStorageRequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaEphemeralStorageRequestUsed), nil
}

func (m ClusterquotaEphemeralStorageRequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaEphemeralStorageRequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaEphemeralStorageRequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaEphemeralStorageRequestUsed) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ClusterquotaEphemeralStorageRequestUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaEphemeralStorageRequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaHugepageCountRequestHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaHugepageCountRequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{hugepage}"),
}

func NewClusterquotaHugepageCountRequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaHugepageCountRequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaHugepageCountRequestHard), nil
}

func (m ClusterquotaHugepageCountRequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaHugepageCountRequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaHugepageCountRequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaHugepageCountRequestHard) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ClusterquotaHugepageCountRequestHard) Add(
	ctx context.Context,
	incr int64,
	k8sHugepageSize string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaHugepageCountRequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaHugepageCountRequestUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaHugepageCountRequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{hugepage}"),
}

func NewClusterquotaHugepageCountRequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaHugepageCountRequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaHugepageCountRequestUsed), nil
}

func (m ClusterquotaHugepageCountRequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaHugepageCountRequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaHugepageCountRequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaHugepageCountRequestUsed) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ClusterquotaHugepageCountRequestUsed) Add(
	ctx context.Context,
	incr int64,
	k8sHugepageSize string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaHugepageCountRequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaMemoryLimitHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaMemoryLimitHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaMemoryLimitHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaMemoryLimitHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaMemoryLimitHard), nil
}

func (m ClusterquotaMemoryLimitHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaMemoryLimitHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryLimitHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryLimitHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaMemoryLimitHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaMemoryLimitHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaMemoryLimitUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaMemoryLimitUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaMemoryLimitUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaMemoryLimitUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaMemoryLimitUsed), nil
}

func (m ClusterquotaMemoryLimitUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaMemoryLimitUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryLimitUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryLimitUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaMemoryLimitUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaMemoryLimitUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaMemoryRequestHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaMemoryRequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaMemoryRequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaMemoryRequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaMemoryRequestHard), nil
}

func (m ClusterquotaMemoryRequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaMemoryRequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryRequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryRequestHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaMemoryRequestHard) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaMemoryRequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaMemoryRequestUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaMemoryRequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaMemoryRequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaMemoryRequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaMemoryRequestUsed), nil
}

func (m ClusterquotaMemoryRequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaMemoryRequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryRequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryRequestUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaMemoryRequestUsed) Add(ctx context.Context, incr int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaMemoryRequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaObjectCountHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaObjectCountHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{object}"),
}

func NewClusterquotaObjectCountHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaObjectCountHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaObjectCountHard), nil
}

func (m ClusterquotaObjectCountHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaObjectCountHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaObjectCountHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaObjectCountHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaObjectCountHard) Add(
	ctx context.Context,
	incr int64,
	k8sResourcequotaResourceName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaObjectCountHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaObjectCountUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaObjectCountUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{object}"),
}

func NewClusterquotaObjectCountUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaObjectCountUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaObjectCountUsed), nil
}

func (m ClusterquotaObjectCountUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaObjectCountUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaObjectCountUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaObjectCountUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaObjectCountUsed) Add(
	ctx context.Context,
	incr int64,
	k8sResourcequotaResourceName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaObjectCountUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClusterquotaPersistentvolumeclaimCountHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaPersistentvolumeclaimCountHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewClusterquotaPersistentvolumeclaimCountHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaPersistentvolumeclaimCountHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaPersistentvolumeclaimCountHard), nil
}

func (m ClusterquotaPersistentvolumeclaimCountHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaPersistentvolumeclaimCountHard) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountHard) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountHard) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ClusterquotaPersistentvolumeclaimCountHard) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaPersistentvolumeclaimCountHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClusterquotaPersistentvolumeclaimCountHard) AttrK8SStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClusterquotaPersistentvolumeclaimCountUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaPersistentvolumeclaimCountUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewClusterquotaPersistentvolumeclaimCountUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaPersistentvolumeclaimCountUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaPersistentvolumeclaimCountUsed), nil
}

func (m ClusterquotaPersistentvolumeclaimCountUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaPersistentvolumeclaimCountUsed) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountUsed) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountUsed) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (m ClusterquotaPersistentvolumeclaimCountUsed) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaPersistentvolumeclaimCountUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClusterquotaPersistentvolumeclaimCountUsed) AttrK8SStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClusterquotaStorageRequestHard struct {
	metric.Int64UpDownCounter
}

var newClusterquotaStorageRequestHardOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaStorageRequestHard(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaStorageRequestHard, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaStorageRequestHard), nil
}

func (m ClusterquotaStorageRequestHard) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaStorageRequestHard) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaStorageRequestHard) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaStorageRequestHard) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaStorageRequestHard) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaStorageRequestHard) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClusterquotaStorageRequestHard) AttrK8SStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClusterquotaStorageRequestUsed struct {
	metric.Int64UpDownCounter
}

var newClusterquotaStorageRequestUsedOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaStorageRequestUsed(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClusterquotaStorageRequestUsed, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaStorageRequestUsed), nil
}

func (m ClusterquotaStorageRequestUsed) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClusterquotaStorageRequestUsed) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaStorageRequestUsed) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaStorageRequestUsed) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClusterquotaStorageRequestUsed) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClusterquotaStorageRequestUsed) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClusterquotaStorageRequestUsed) AttrK8SStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
