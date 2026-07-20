package openshiftconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
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

type ClusterquotaCPULimitHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaCPULimitHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{cpu}"),
}

func NewClusterquotaCPULimitHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaCPULimitHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaCPULimitHardObservable), nil
}

func (m ClusterquotaCPULimitHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaCPULimitHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPULimitHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPULimitHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaCPULimitUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaCPULimitUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{cpu}"),
}

func NewClusterquotaCPULimitUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaCPULimitUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaCPULimitUsedObservable), nil
}

func (m ClusterquotaCPULimitUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaCPULimitUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPULimitUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPULimitUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaCPURequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaCPURequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{cpu}"),
}

func NewClusterquotaCPURequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaCPURequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaCPURequestHardObservable), nil
}

func (m ClusterquotaCPURequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaCPURequestHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPURequestHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPURequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaCPURequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaCPURequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{cpu}"),
}

func NewClusterquotaCPURequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaCPURequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaCPURequestUsedObservable), nil
}

func (m ClusterquotaCPURequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaCPURequestUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPURequestUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaCPURequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaEphemeralStorageLimitHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaEphemeralStorageLimitHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaEphemeralStorageLimitHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaEphemeralStorageLimitHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaEphemeralStorageLimitHardObservable), nil
}

func (m ClusterquotaEphemeralStorageLimitHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaEphemeralStorageLimitHardObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaEphemeralStorageLimitHardObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaEphemeralStorageLimitHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaEphemeralStorageLimitUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaEphemeralStorageLimitUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaEphemeralStorageLimitUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaEphemeralStorageLimitUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaEphemeralStorageLimitUsedObservable), nil
}

func (m ClusterquotaEphemeralStorageLimitUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaEphemeralStorageLimitUsedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaEphemeralStorageLimitUsedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaEphemeralStorageLimitUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaEphemeralStorageRequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaEphemeralStorageRequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaEphemeralStorageRequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaEphemeralStorageRequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaEphemeralStorageRequestHardObservable), nil
}

func (m ClusterquotaEphemeralStorageRequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaEphemeralStorageRequestHardObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaEphemeralStorageRequestHardObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaEphemeralStorageRequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaEphemeralStorageRequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaEphemeralStorageRequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaEphemeralStorageRequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaEphemeralStorageRequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaEphemeralStorageRequestUsedObservable), nil
}

func (m ClusterquotaEphemeralStorageRequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaEphemeralStorageRequestUsedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaEphemeralStorageRequestUsedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaEphemeralStorageRequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaHugepageCountRequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaHugepageCountRequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{hugepage}"),
}

func NewClusterquotaHugepageCountRequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaHugepageCountRequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaHugepageCountRequestHardObservable), nil
}

func (m ClusterquotaHugepageCountRequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaHugepageCountRequestHardObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaHugepageCountRequestHardObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaHugepageCountRequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaHugepageCountRequestHardObservable) AttrK8SHugepageSize(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ClusterquotaHugepageCountRequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaHugepageCountRequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{hugepage}"),
}

func NewClusterquotaHugepageCountRequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaHugepageCountRequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaHugepageCountRequestUsedObservable), nil
}

func (m ClusterquotaHugepageCountRequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaHugepageCountRequestUsedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaHugepageCountRequestUsedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaHugepageCountRequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaHugepageCountRequestUsedObservable) AttrK8SHugepageSize(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ClusterquotaMemoryLimitHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaMemoryLimitHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaMemoryLimitHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaMemoryLimitHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaMemoryLimitHardObservable), nil
}

func (m ClusterquotaMemoryLimitHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaMemoryLimitHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryLimitHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryLimitHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaMemoryLimitUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaMemoryLimitUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaMemoryLimitUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaMemoryLimitUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaMemoryLimitUsedObservable), nil
}

func (m ClusterquotaMemoryLimitUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaMemoryLimitUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryLimitUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryLimitUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaMemoryRequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaMemoryRequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaMemoryRequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaMemoryRequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaMemoryRequestHardObservable), nil
}

func (m ClusterquotaMemoryRequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaMemoryRequestHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryRequestHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryRequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaMemoryRequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaMemoryRequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaMemoryRequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaMemoryRequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaMemoryRequestUsedObservable), nil
}

func (m ClusterquotaMemoryRequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaMemoryRequestUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryRequestUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaMemoryRequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
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

type ClusterquotaObjectCountHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaObjectCountHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{object}"),
}

func NewClusterquotaObjectCountHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaObjectCountHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaObjectCountHardObservable), nil
}

func (m ClusterquotaObjectCountHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaObjectCountHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaObjectCountHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaObjectCountHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaObjectCountHardObservable) AttrK8SResourceQuotaResourceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ClusterquotaObjectCountUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaObjectCountUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{object}"),
}

func NewClusterquotaObjectCountUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaObjectCountUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaObjectCountUsedObservable), nil
}

func (m ClusterquotaObjectCountUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaObjectCountUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaObjectCountUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaObjectCountUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaObjectCountUsedObservable) AttrK8SResourceQuotaResourceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
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

type ClusterquotaPersistentvolumeclaimCountHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaPersistentvolumeclaimCountHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewClusterquotaPersistentvolumeclaimCountHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaPersistentvolumeclaimCountHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaPersistentvolumeclaimCountHardObservable), nil
}

func (m ClusterquotaPersistentvolumeclaimCountHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaPersistentvolumeclaimCountHardObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountHardObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountHardObservable) AttrK8SStorageclassName(val string) attribute.KeyValue {
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

type ClusterquotaPersistentvolumeclaimCountUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaPersistentvolumeclaimCountUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("{persistentvolumeclaim}"),
}

func NewClusterquotaPersistentvolumeclaimCountUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaPersistentvolumeclaimCountUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaPersistentvolumeclaimCountUsedObservable), nil
}

func (m ClusterquotaPersistentvolumeclaimCountUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaPersistentvolumeclaimCountUsedObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountUsedObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaPersistentvolumeclaimCountUsedObservable) AttrK8SStorageclassName(val string) attribute.KeyValue {
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

type ClusterquotaStorageRequestHardObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaStorageRequestHardObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The enforced hard limit of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaStorageRequestHardObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaStorageRequestHardObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaStorageRequestHardObservable), nil
}

func (m ClusterquotaStorageRequestHardObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaStorageRequestHardObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaStorageRequestHardObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaStorageRequestHardObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaStorageRequestHardObservable) AttrK8SStorageclassName(val string) attribute.KeyValue {
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

type ClusterquotaStorageRequestUsedObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClusterquotaStorageRequestUsedObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The current observed total usage of the resource across all projects."),
	metric.WithUnit("By"),
}

func NewClusterquotaStorageRequestUsedObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClusterquotaStorageRequestUsedObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClusterquotaStorageRequestUsedObservable), nil
}

func (m ClusterquotaStorageRequestUsedObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClusterquotaStorageRequestUsedObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaStorageRequestUsedObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClusterquotaStorageRequestUsedObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClusterquotaStorageRequestUsedObservable) AttrK8SStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
