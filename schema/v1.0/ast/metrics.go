package ast

import "go.opentelemetry.io/otel/schema/v1.0/types"

type Metrics struct {
	Changes []MetricsChange
}

type MetricsChange struct {
	RenameMetrics    map[types.MetricName]types.MetricName `yaml:"rename_metrics"`
	RenameAttributes *AttributeMapForMetrics               `yaml:"rename_attributes"`
}

type AttributeMapForMetrics struct {
	ApplyToMetrics []types.MetricName `yaml:"apply_to_metrics"`
	AttributeMap   AttributeMap       `yaml:"attribute_map"`
}
