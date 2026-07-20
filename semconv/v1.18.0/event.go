package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	FeatureFlagKeyKey = attribute.Key("feature_flag.key")

	FeatureFlagProviderNameKey = attribute.Key("feature_flag.provider_name")

	FeatureFlagVariantKey = attribute.Key("feature_flag.variant")
)

func FeatureFlagKey(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FeatureFlagProviderName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FeatureFlagVariant(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessageTypeKey = attribute.Key("message.type")

	MessageIDKey = attribute.Key("message.id")

	MessageCompressedSizeKey = attribute.Key("message.compressed_size")

	MessageUncompressedSizeKey = attribute.Key("message.uncompressed_size")
)

var (
	MessageTypeSent = MessageTypeKey.String("SENT")

	MessageTypeReceived = MessageTypeKey.String("RECEIVED")
)

func MessageID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessageCompressedSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessageUncompressedSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ExceptionEscapedKey = attribute.Key("exception.escaped")
)

func ExceptionEscaped(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
