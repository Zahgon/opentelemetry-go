package resource

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/otel/attribute"
)

type Resource struct {
	attrs     attribute.Set
	schemaURL string
}

var _ map[Resource]struct{} = nil

var (
	defaultResource     *Resource
	defaultResourceOnce sync.Once
)

var ErrSchemaURLConflict = errors.New("conflicting Schema URL")

func New(ctx context.Context, opts ...Option) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWithAttributes(schemaURL string, attrs ...attribute.KeyValue) *Resource {
	_ = "STUB: not implemented"
	return nil
}

func NewSchemaless(attrs ...attribute.KeyValue) *Resource { _ = "STUB: not implemented"; return nil }

func (r *Resource) String() string { _ = "STUB: not implemented"; return "" }

func (r *Resource) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }

func (r *Resource) Attributes() []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func (r *Resource) SchemaURL() string { _ = "STUB: not implemented"; return "" }

func (r *Resource) Iter() attribute.Iterator {
	_ = "STUB: not implemented"
	return *new(attribute.Iterator)
}

func (r *Resource) Equal(o *Resource) bool { _ = "STUB: not implemented"; return false }

func Merge(a, b *Resource) (*Resource, error) { _ = "STUB: not implemented"; return nil, nil }

func Empty() *Resource { _ = "STUB: not implemented"; return nil }

func Default() *Resource { _ = "STUB: not implemented"; return nil }

func DefaultWithContext(ctx context.Context) *Resource { _ = "STUB: not implemented"; return nil }

func Environment() *Resource { _ = "STUB: not implemented"; return nil }

func EnvironmentWithContext(ctx context.Context) *Resource { _ = "STUB: not implemented"; return nil }

func (r *Resource) Equivalent() attribute.Distinct {
	_ = "STUB: not implemented"
	return *new(attribute.Distinct)
}

func (r *Resource) Set() *attribute.Set { _ = "STUB: not implemented"; return nil }

func (r *Resource) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Resource) Len() int { _ = "STUB: not implemented"; return 0 }

func (r *Resource) Encoded(enc attribute.Encoder) string { _ = "STUB: not implemented"; return "" }
