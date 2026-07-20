package ast

import (
	ast10 "go.opentelemetry.io/otel/schema/v1.0/ast"
	types10 "go.opentelemetry.io/otel/schema/v1.0/types"
	types11 "go.opentelemetry.io/otel/schema/v1.1/types"
)

type Metrics struct {
	Changes []MetricsChange
}

type MetricsChange struct {
	RenameMetrics    map[types10.MetricName]types10.MetricName `yaml:"rename_metrics"`
	RenameAttributes *ast10.AttributeMapForMetrics             `yaml:"rename_attributes"`
	Split            *SplitMetric                              `yaml:"split"`
}

type SplitMetric struct {
	ApplyToMetric types10.MetricName `yaml:"apply_to_metric"`

	ByAttribute types11.AttributeName `yaml:"by_attribute"`

	MetricsFromAttributes map[types10.MetricName]types11.AttributeValue `yaml:"metrics_from_attributes"`
}
