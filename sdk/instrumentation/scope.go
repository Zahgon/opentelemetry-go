package instrumentation

import "go.opentelemetry.io/otel/attribute"

type Scope struct {
	Name string

	Version string

	SchemaURL string

	Attributes attribute.Set
}
