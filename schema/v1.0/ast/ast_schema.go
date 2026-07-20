package ast

import "go.opentelemetry.io/otel/schema/v1.0/types"

type Schema struct {
	FileFormat string `yaml:"file_format"`

	SchemaURL string `yaml:"schema_url"`

	Versions map[types.TelemetryVersion]VersionDef
}

type VersionDef struct {
	All        Attributes
	Resources  Attributes
	Spans      Spans
	SpanEvents SpanEvents `yaml:"span_events"`
	Logs       Logs
	Metrics    Metrics
}

type Attributes struct {
	Changes []AttributeChange
}

type AttributeChange struct {
	RenameAttributes *RenameAttributes `yaml:"rename_attributes"`
}
