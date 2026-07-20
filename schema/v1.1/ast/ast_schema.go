package ast

import (
	ast10 "go.opentelemetry.io/otel/schema/v1.0/ast"
	"go.opentelemetry.io/otel/schema/v1.1/types"
)

type Schema struct {
	FileFormat string `yaml:"file_format"`

	SchemaURL string `yaml:"schema_url"`

	Versions map[types.TelemetryVersion]VersionDef
}

type VersionDef struct {
	All        ast10.Attributes
	Resources  ast10.Attributes
	Spans      ast10.Spans
	SpanEvents ast10.SpanEvents `yaml:"span_events"`
	Logs       ast10.Logs
	Metrics    Metrics
}
