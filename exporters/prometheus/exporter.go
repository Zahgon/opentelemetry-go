package prometheus

import (
	"context"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/otlptranslator"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus/internal/observ"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
)

const (
	targetInfoDescription = "Target metadata"

	scopeLabelPrefix  = "otel_scope_"
	scopeNameLabel    = scopeLabelPrefix + "name"
	scopeVersionLabel = scopeLabelPrefix + "version"
	scopeSchemaLabel  = scopeLabelPrefix + "schema_url"

	bridgeScopeName = "go.opentelemetry.io/contrib/bridges/prometheus"
)

var metricsPool = sync.Pool{
	New: func() any {
		return &metricdata.ResourceMetrics{}
	},
}

type Exporter struct {
	metric.Reader
}

func (e *Exporter) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }

var _ metric.Reader = &Exporter{}

type keyVals struct {
	keys []string
	vals []string
}

type collector struct {
	reader metric.Reader

	withoutUnits             bool
	withoutCounterSuffixes   bool
	disableScopeInfo         bool
	namespace                string
	resourceAttributesFilter attribute.Filter

	mu                sync.Mutex
	disableTargetInfo bool
	targetInfo        prometheus.Metric
	metricFamilies    map[string]*dto.MetricFamily

	resourceKeyValsOnce sync.Once
	resourceKeyVals     keyVals
	resourceKeyValsErr  error

	metricNamer otlptranslator.MetricNamer
	labelNamer  otlptranslator.LabelNamer
	unitNamer   otlptranslator.UnitNamer

	inst *observ.Instrumentation

	bridgeErrorOnce sync.Once
}

func New(opts ...Option) (*Exporter, error) { _ = "STUB: not implemented"; return nil, nil }

func (*collector) Describe(chan<- *prometheus.Desc) { _ = "STUB: not implemented"; return }

func (c *collector) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

func downscaleExponentialBucket(bucket metricdata.ExponentialBucket, scaleDelta int32) metricdata.ExponentialBucket {
	_ = "STUB: not implemented"
	return *new(metricdata.ExponentialBucket)
}

//nolint:gosec // Length is bounded by slice allocation

//nolint:gosec // Index is bounded by loop iteration

//nolint:gosec // Length is bounded by allocation

func addExponentialHistogramMetric[N int64 | float64](
	ch chan<- prometheus.Metric,
	histogram metricdata.ExponentialHistogram[N],
	m metricdata.Metrics,
	name string,
	kv keyVals,
	labelNamer otlptranslator.LabelNamer,
	inst *observ.Instrumentation,
	ctx context.Context,
) {
	_ = "STUB: not implemented"
	return
}

func addHistogramMetric[N int64 | float64](
	ch chan<- prometheus.Metric,
	histogram metricdata.Histogram[N],
	m metricdata.Metrics,
	name string,
	kv keyVals,
	labelNamer otlptranslator.LabelNamer,
	inst *observ.Instrumentation,
	ctx context.Context,
) {
	_ = "STUB: not implemented"
	return
}

func addSumMetric[N int64 | float64](
	ch chan<- prometheus.Metric,
	sum metricdata.Sum[N],
	m metricdata.Metrics,
	name string,
	kv keyVals,
	labelNamer otlptranslator.LabelNamer,
	inst *observ.Instrumentation,
	ctx context.Context,
) {
	_ = "STUB: not implemented"
	return
}

func addGaugeMetric[N int64 | float64](
	ch chan<- prometheus.Metric,
	gauge metricdata.Gauge[N],
	m metricdata.Metrics,
	name string,
	kv keyVals,
	labelNamer otlptranslator.LabelNamer,
	inst *observ.Instrumentation,
	ctx context.Context,
) {
	_ = "STUB: not implemented"
	return
}

func getAttrs(attrs attribute.Set, labelNamer otlptranslator.LabelNamer) ([]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getScopeAttrs(attrs attribute.Set, labelNamer otlptranslator.LabelNamer) ([]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func isReservedScopeLabel(key string) bool { _ = "STUB: not implemented"; return false }

func (c *collector) createInfoMetric(name, description string, res *resource.Resource) (prometheus.Metric, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Metric), nil
}

func (c *collector) getName(m metricdata.Metrics) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (*collector) metricType(m metricdata.Metrics) *dto.MetricType {
	_ = "STUB: not implemented"
	return nil
}

func (c *collector) namingMetricType(m metricdata.Metrics) otlptranslator.MetricType {
	_ = "STUB: not implemented"
	return *new(otlptranslator.MetricType)
}

func (c *collector) createResourceAttributes(res *resource.Resource) (keyVals, error) {
	_ = "STUB: not implemented"
	return *new(keyVals), nil
}

func (c *collector) validateMetrics(name, description string, metricType *dto.MetricType) (drop bool, help string) {
	_ = "STUB: not implemented"
	return false, ""
}

func addExemplars[N int64 | float64](
	m prometheus.Metric,
	exemplars []metricdata.Exemplar[N],
	labelNamer otlptranslator.LabelNamer,
) prometheus.Metric {
	_ = "STUB: not implemented"
	return *new(prometheus.Metric)
}

func attributesToLabels(attrs []attribute.KeyValue, labelNamer otlptranslator.LabelNamer) (prometheus.Labels, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels), nil
}

func reportError(ch chan<- prometheus.Metric, desc *prometheus.Desc, err error) {
	_ = "STUB: not implemented"
	return
}
