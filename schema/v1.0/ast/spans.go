package ast

import "go.opentelemetry.io/otel/schema/v1.0/types"

type Spans struct {
	Changes []SpansChange
}

type SpanEvents struct {
	Changes []SpanEventsChange
}

type SpansChange struct {
	RenameAttributes *AttributeMapForSpans `yaml:"rename_attributes"`
}

type AttributeMapForSpans struct {
	ApplyToSpans []types.SpanName `yaml:"apply_to_spans"`
	AttributeMap AttributeMap     `yaml:"attribute_map"`
}

type SpanEventsChange struct {
	RenameEvents     *RenameSpanEvents          `yaml:"rename_events"`
	RenameAttributes *RenameSpanEventAttributes `yaml:"rename_attributes"`
}

type RenameSpanEvents struct {
	EventNameMap map[string]string `yaml:"name_map"`
}

type RenameSpanEventAttributes struct {
	ApplyToSpans  []types.SpanName  `yaml:"apply_to_spans"`
	ApplyToEvents []types.EventName `yaml:"apply_to_events"`
	AttributeMap  AttributeMap      `yaml:"attribute_map"`
}
