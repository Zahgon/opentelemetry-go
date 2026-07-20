package vcsconv

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

	ProviderNameGittea ProviderNameAttr = "gittea"

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

type ChangeDuration struct {
	metric.Float64Gauge
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

type ChangeTimeToApproval struct {
	metric.Float64Gauge
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

type ChangeTimeToMerge struct {
	metric.Float64Gauge
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

type ContributorCount struct {
	metric.Int64Gauge
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

type RefCount struct {
	metric.Int64UpDownCounter
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

type RefLinesDelta struct {
	metric.Int64Gauge
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

type RefRevisionsDelta struct {
	metric.Int64Gauge
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

type RefTime struct {
	metric.Float64Gauge
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

type RepositoryCount struct {
	metric.Int64UpDownCounter
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

func (RepositoryCount) AttrOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (RepositoryCount) AttrProviderName(val ProviderNameAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
