package trace

const (
	DefaultAttributeValueLengthLimit = -1

	DefaultAttributeCountLimit = 128

	DefaultEventCountLimit = 128

	DefaultLinkCountLimit = 128

	DefaultAttributePerEventCountLimit = 128

	DefaultAttributePerLinkCountLimit = 128
)

type SpanLimits struct {
	AttributeValueLengthLimit int

	AttributeCountLimit int

	EventCountLimit int

	LinkCountLimit int

	AttributePerEventCountLimit int

	AttributePerLinkCountLimit int
}

func NewSpanLimits() SpanLimits { _ = "STUB: not implemented"; return *new(SpanLimits) }
