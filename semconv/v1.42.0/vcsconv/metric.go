package vcsconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type ChangeStateAttr string

var (
	ChangeStateOpen ChangeStateAttr = "open"

	ChangeStateWip ChangeStateAttr = "wip"

	ChangeStateClosed ChangeStateAttr = "closed"

	ChangeStateMerged ChangeStateAttr = "merged"
)

type LineChangeTypeAttr string

var (
	LineChangeTypeAdded LineChangeTypeAttr = "added"

	LineChangeTypeRemoved LineChangeTypeAttr = "removed"
)

type ProviderNameAttr string

var (
	ProviderNameGithub ProviderNameAttr = "github"

	ProviderNameGitlab ProviderNameAttr = "gitlab"

	ProviderNameGitea ProviderNameAttr = "gitea"

	ProviderNameBitbucket ProviderNameAttr = "bitbucket"
)

type RefBaseTypeAttr string

var (
	RefBaseTypeBranch RefBaseTypeAttr = "branch"

	RefBaseTypeTag RefBaseTypeAttr = "tag"
)

type RefHeadTypeAttr string

var (
	RefHeadTypeBranch RefHeadTypeAttr = "branch"

	RefHeadTypeTag RefHeadTypeAttr = "tag"
)

type RefTypeAttr string

var (
	RefTypeBranch RefTypeAttr = "branch"

	RefTypeTag RefTypeAttr = "tag"
)

type RevisionDeltaDirectionAttr string

var (
	RevisionDeltaDirectionBehind RevisionDeltaDirectionAttr = "behind"

	RevisionDeltaDirectionAhead RevisionDeltaDirectionAttr = "ahead"
)

type ChangeCount struct {
	metric.Int64UpDownCounter
}

var newChangeCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of changes (pull requests/merge requests/changelists) in a repository, categorized by their state (e.g. open or merged)."),
	metric.WithUnit("{change}"),
}

func NewChangeCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ChangeCount, error) {
	_ = "STUB: not implemented"
	return *new(ChangeCount), nil
}

func (m ChangeCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ChangeCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ChangeCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ChangeCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ChangeCount) Add(
	ctx context.Context,
	incr int64,
	changeState ChangeStateAttr,
	repositoryUrlFull string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ChangeCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ChangeCount) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeCount) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeCount) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ChangeCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newChangeCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of changes (pull requests/merge requests/changelists) in a repository, categorized by their state (e.g. open or merged)."),
	metric.WithUnit("{change}"),
}

func NewChangeCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ChangeCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(ChangeCountObservable), nil
}

func (m ChangeCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ChangeCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ChangeCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ChangeCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ChangeCountObservable) AttrChangeState(val ChangeStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeCountObservable) AttrRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeCountObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeCountObservable) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeCountObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ChangeDuration struct {
	metric.Float64Gauge
}

var newChangeDurationOpts = []metric.Float64GaugeOption{
	metric.WithDescription("The time duration a change (pull request/merge request/changelist) has been in a given state."),
	metric.WithUnit("s"),
}

func NewChangeDuration(
	m metric.Meter,
	opt ...metric.Float64GaugeOption,
) (ChangeDuration, error) {
	_ = "STUB: not implemented"
	return *new(ChangeDuration), nil
}

func (m ChangeDuration) Inst() metric.Float64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge)
}

func (ChangeDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ChangeDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ChangeDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ChangeDuration) Record(
	ctx context.Context,
	val float64,
	changeState ChangeStateAttr,
	refHeadName string,
	repositoryUrlFull string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ChangeDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ChangeDuration) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeDuration) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeDuration) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ChangeDurationObservable struct {
	metric.Float64ObservableGauge
}

var newChangeDurationObservableOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("The time duration a change (pull request/merge request/changelist) has been in a given state."),
	metric.WithUnit("s"),
}

func NewChangeDurationObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (ChangeDurationObservable, error) {
	_ = "STUB: not implemented"
	return *new(ChangeDurationObservable), nil
}

func (m ChangeDurationObservable) Inst() metric.Float64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge)
}

func (ChangeDurationObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ChangeDurationObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ChangeDurationObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ChangeDurationObservable) AttrChangeState(val ChangeStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeDurationObservable) AttrRefHeadName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeDurationObservable) AttrRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeDurationObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeDurationObservable) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeDurationObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ChangeTimeToApproval struct {
	metric.Float64Gauge
}

var newChangeTimeToApprovalOpts = []metric.Float64GaugeOption{
	metric.WithDescription("The amount of time since its creation it took a change (pull request/merge request/changelist) to get the first approval."),
	metric.WithUnit("s"),
}

func NewChangeTimeToApproval(
	m metric.Meter,
	opt ...metric.Float64GaugeOption,
) (ChangeTimeToApproval, error) {
	_ = "STUB: not implemented"
	return *new(ChangeTimeToApproval), nil
}

func (m ChangeTimeToApproval) Inst() metric.Float64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge)
}

func (ChangeTimeToApproval) Name() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToApproval) Unit() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToApproval) Description() string { _ = "STUB: not implemented"; return "" }

func (m ChangeTimeToApproval) Record(
	ctx context.Context,
	val float64,
	refHeadName string,
	repositoryUrlFull string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ChangeTimeToApproval) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ChangeTimeToApproval) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApproval) AttrRefBaseName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApproval) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApproval) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApproval) AttrRefBaseRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApproval) AttrRefHeadRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ChangeTimeToApprovalObservable struct {
	metric.Float64ObservableGauge
}

var newChangeTimeToApprovalObservableOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("The amount of time since its creation it took a change (pull request/merge request/changelist) to get the first approval."),
	metric.WithUnit("s"),
}

func NewChangeTimeToApprovalObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (ChangeTimeToApprovalObservable, error) {
	_ = "STUB: not implemented"
	return *new(ChangeTimeToApprovalObservable), nil
}

func (m ChangeTimeToApprovalObservable) Inst() metric.Float64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge)
}

func (ChangeTimeToApprovalObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToApprovalObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToApprovalObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToApprovalObservable) AttrRefHeadName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApprovalObservable) AttrRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApprovalObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApprovalObservable) AttrRefBaseName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApprovalObservable) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApprovalObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApprovalObservable) AttrRefBaseRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToApprovalObservable) AttrRefHeadRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ChangeTimeToMerge struct {
	metric.Float64Gauge
}

var newChangeTimeToMergeOpts = []metric.Float64GaugeOption{
	metric.WithDescription("The amount of time since its creation it took a change (pull request/merge request/changelist) to get merged into the target(base) ref."),
	metric.WithUnit("s"),
}

func NewChangeTimeToMerge(
	m metric.Meter,
	opt ...metric.Float64GaugeOption,
) (ChangeTimeToMerge, error) {
	_ = "STUB: not implemented"
	return *new(ChangeTimeToMerge), nil
}

func (m ChangeTimeToMerge) Inst() metric.Float64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge)
}

func (ChangeTimeToMerge) Name() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToMerge) Unit() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToMerge) Description() string { _ = "STUB: not implemented"; return "" }

func (m ChangeTimeToMerge) Record(
	ctx context.Context,
	val float64,
	refHeadName string,
	repositoryUrlFull string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ChangeTimeToMerge) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ChangeTimeToMerge) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMerge) AttrRefBaseName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMerge) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMerge) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMerge) AttrRefBaseRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMerge) AttrRefHeadRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ChangeTimeToMergeObservable struct {
	metric.Float64ObservableGauge
}

var newChangeTimeToMergeObservableOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("The amount of time since its creation it took a change (pull request/merge request/changelist) to get merged into the target(base) ref."),
	metric.WithUnit("s"),
}

func NewChangeTimeToMergeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (ChangeTimeToMergeObservable, error) {
	_ = "STUB: not implemented"
	return *new(ChangeTimeToMergeObservable), nil
}

func (m ChangeTimeToMergeObservable) Inst() metric.Float64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge)
}

func (ChangeTimeToMergeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToMergeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToMergeObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ChangeTimeToMergeObservable) AttrRefHeadName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMergeObservable) AttrRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMergeObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMergeObservable) AttrRefBaseName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMergeObservable) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMergeObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMergeObservable) AttrRefBaseRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ChangeTimeToMergeObservable) AttrRefHeadRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ContributorCount struct {
	metric.Int64Gauge
}

var newContributorCountOpts = []metric.Int64GaugeOption{
	metric.WithDescription("The number of unique contributors to a repository."),
	metric.WithUnit("{contributor}"),
}

func NewContributorCount(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (ContributorCount, error) {
	_ = "STUB: not implemented"
	return *new(ContributorCount), nil
}

func (m ContributorCount) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (ContributorCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ContributorCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContributorCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ContributorCount) Record(
	ctx context.Context,
	val int64,
	repositoryUrlFull string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ContributorCount) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ContributorCount) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ContributorCount) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ContributorCount) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ContributorCountObservable struct {
	metric.Int64ObservableGauge
}

var newContributorCountObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("The number of unique contributors to a repository."),
	metric.WithUnit("{contributor}"),
}

func NewContributorCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (ContributorCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(ContributorCountObservable), nil
}

func (m ContributorCountObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (ContributorCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ContributorCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ContributorCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ContributorCountObservable) AttrRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ContributorCountObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ContributorCountObservable) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ContributorCountObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RefCount struct {
	metric.Int64UpDownCounter
}

var newRefCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of refs of type branch or tag in a repository."),
	metric.WithUnit("{ref}"),
}

func NewRefCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (RefCount, error) {
	_ = "STUB: not implemented"
	return *new(RefCount), nil
}

func (m RefCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (RefCount) Name() string { _ = "STUB: not implemented"; return "" }

func (RefCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (RefCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m RefCount) Add(
	ctx context.Context,
	incr int64,
	refType RefTypeAttr,
	repositoryUrlFull string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m RefCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (RefCount) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefCount) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefCount) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RefCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newRefCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of refs of type branch or tag in a repository."),
	metric.WithUnit("{ref}"),
}

func NewRefCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (RefCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(RefCountObservable), nil
}

func (m RefCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (RefCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (RefCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (RefCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (RefCountObservable) AttrRefType(val RefTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefCountObservable) AttrRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefCountObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefCountObservable) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefCountObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RefLinesDelta struct {
	metric.Int64Gauge
}

var newRefLinesDeltaOpts = []metric.Int64GaugeOption{
	metric.WithDescription("The number of lines added/removed in a ref (branch) relative to the ref from the `vcs.ref.base.name` attribute."),
	metric.WithUnit("{line}"),
}

func NewRefLinesDelta(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (RefLinesDelta, error) {
	_ = "STUB: not implemented"
	return *new(RefLinesDelta), nil
}

func (m RefLinesDelta) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (RefLinesDelta) Name() string { _ = "STUB: not implemented"; return "" }

func (RefLinesDelta) Unit() string { _ = "STUB: not implemented"; return "" }

func (RefLinesDelta) Description() string { _ = "STUB: not implemented"; return "" }

func (m RefLinesDelta) Record(
	ctx context.Context,
	val int64,
	lineChangeType LineChangeTypeAttr,
	refBaseName string,
	refBaseType RefBaseTypeAttr,
	refHeadName string,
	refHeadType RefHeadTypeAttr,
	repositoryUrlFull string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m RefLinesDelta) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (RefLinesDelta) AttrChangeID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDelta) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDelta) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDelta) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RefLinesDeltaObservable struct {
	metric.Int64ObservableGauge
}

var newRefLinesDeltaObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("The number of lines added/removed in a ref (branch) relative to the ref from the `vcs.ref.base.name` attribute."),
	metric.WithUnit("{line}"),
}

func NewRefLinesDeltaObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (RefLinesDeltaObservable, error) {
	_ = "STUB: not implemented"
	return *new(RefLinesDeltaObservable), nil
}

func (m RefLinesDeltaObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (RefLinesDeltaObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (RefLinesDeltaObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (RefLinesDeltaObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (RefLinesDeltaObservable) AttrLineChangeType(val LineChangeTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDeltaObservable) AttrRefBaseName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDeltaObservable) AttrRefBaseType(val RefBaseTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDeltaObservable) AttrRefHeadName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDeltaObservable) AttrRefHeadType(val RefHeadTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDeltaObservable) AttrRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDeltaObservable) AttrChangeID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDeltaObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDeltaObservable) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefLinesDeltaObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RefRevisionsDelta struct {
	metric.Int64Gauge
}

var newRefRevisionsDeltaOpts = []metric.Int64GaugeOption{
	metric.WithDescription("The number of revisions (commits) a ref (branch) is ahead/behind the branch from the `vcs.ref.base.name` attribute."),
	metric.WithUnit("{revision}"),
}

func NewRefRevisionsDelta(
	m metric.Meter,
	opt ...metric.Int64GaugeOption,
) (RefRevisionsDelta, error) {
	_ = "STUB: not implemented"
	return *new(RefRevisionsDelta), nil
}

func (m RefRevisionsDelta) Inst() metric.Int64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64Gauge)
}

func (RefRevisionsDelta) Name() string { _ = "STUB: not implemented"; return "" }

func (RefRevisionsDelta) Unit() string { _ = "STUB: not implemented"; return "" }

func (RefRevisionsDelta) Description() string { _ = "STUB: not implemented"; return "" }

func (m RefRevisionsDelta) Record(
	ctx context.Context,
	val int64,
	refBaseName string,
	refBaseType RefBaseTypeAttr,
	refHeadName string,
	refHeadType RefHeadTypeAttr,
	repositoryUrlFull string,
	revisionDeltaDirection RevisionDeltaDirectionAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m RefRevisionsDelta) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (RefRevisionsDelta) AttrChangeID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDelta) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDelta) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDelta) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RefRevisionsDeltaObservable struct {
	metric.Int64ObservableGauge
}

var newRefRevisionsDeltaObservableOpts = []metric.Int64ObservableGaugeOption{
	metric.WithDescription("The number of revisions (commits) a ref (branch) is ahead/behind the branch from the `vcs.ref.base.name` attribute."),
	metric.WithUnit("{revision}"),
}

func NewRefRevisionsDeltaObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableGaugeOption,
) (RefRevisionsDeltaObservable, error) {
	_ = "STUB: not implemented"
	return *new(RefRevisionsDeltaObservable), nil
}

func (m RefRevisionsDeltaObservable) Inst() metric.Int64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableGauge)
}

func (RefRevisionsDeltaObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (RefRevisionsDeltaObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (RefRevisionsDeltaObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (RefRevisionsDeltaObservable) AttrRefBaseName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDeltaObservable) AttrRefBaseType(val RefBaseTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDeltaObservable) AttrRefHeadName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDeltaObservable) AttrRefHeadType(val RefHeadTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDeltaObservable) AttrRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDeltaObservable) AttrRevisionDeltaDirection(val RevisionDeltaDirectionAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDeltaObservable) AttrChangeID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDeltaObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDeltaObservable) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefRevisionsDeltaObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RefTime struct {
	metric.Float64Gauge
}

var newRefTimeOpts = []metric.Float64GaugeOption{
	metric.WithDescription("Time a ref (branch) created from the default branch (trunk) has existed. The `ref.type` attribute will always be `branch`."),
	metric.WithUnit("s"),
}

func NewRefTime(
	m metric.Meter,
	opt ...metric.Float64GaugeOption,
) (RefTime, error) {
	_ = "STUB: not implemented"
	return *new(RefTime), nil
}

func (m RefTime) Inst() metric.Float64Gauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64Gauge)
}

func (RefTime) Name() string { _ = "STUB: not implemented"; return "" }

func (RefTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (RefTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m RefTime) Record(
	ctx context.Context,
	val float64,
	refHeadName string,
	refHeadType RefHeadTypeAttr,
	repositoryUrlFull string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m RefTime) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (RefTime) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefTime) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefTime) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RefTimeObservable struct {
	metric.Float64ObservableGauge
}

var newRefTimeObservableOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("Time a ref (branch) created from the default branch (trunk) has existed. The `ref.type` attribute will always be `branch`."),
	metric.WithUnit("s"),
}

func NewRefTimeObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (RefTimeObservable, error) {
	_ = "STUB: not implemented"
	return *new(RefTimeObservable), nil
}

func (m RefTimeObservable) Inst() metric.Float64ObservableGauge {
	_ = "STUB: not implemented"
	return *new(metric.Float64ObservableGauge)
}

func (RefTimeObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (RefTimeObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (RefTimeObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (RefTimeObservable) AttrRefHeadName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefTimeObservable) AttrRefHeadType(val RefHeadTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefTimeObservable) AttrRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefTimeObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefTimeObservable) AttrRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RefTimeObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RepositoryCount struct {
	metric.Int64UpDownCounter
}

var newRepositoryCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of repositories in an organization."),
	metric.WithUnit("{repository}"),
}

func NewRepositoryCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (RepositoryCount, error) {
	_ = "STUB: not implemented"
	return *new(RepositoryCount), nil
}

func (m RepositoryCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (RepositoryCount) Name() string { _ = "STUB: not implemented"; return "" }

func (RepositoryCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (RepositoryCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m RepositoryCount) Add(
	ctx context.Context,
	incr int64,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m RepositoryCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (RepositoryCount) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RepositoryCount) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type RepositoryCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newRepositoryCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of repositories in an organization."),
	metric.WithUnit("{repository}"),
}

func NewRepositoryCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (RepositoryCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(RepositoryCountObservable), nil
}

func (m RepositoryCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (RepositoryCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (RepositoryCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (RepositoryCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (RepositoryCountObservable) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RepositoryCountObservable) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
