package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	PeerServiceKey = attribute.Key("peer.service")
)

func PeerService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSLambdaInvokedARNKey = attribute.Key("aws.lambda.invoked_arn")
)

func AWSLambdaInvokedARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	OpentracingRefTypeKey = attribute.Key("opentracing.ref_type")
)

var (
	OpentracingRefTypeChildOf = OpentracingRefTypeKey.String("child_of")

	OpentracingRefTypeFollowsFrom = OpentracingRefTypeKey.String("follows_from")
)

const (
	OTelStatusCodeKey = attribute.Key("otel.status_code")

	OTelStatusDescriptionKey = attribute.Key("otel.status_description")
)

var (
	OTelStatusCodeOk = OTelStatusCodeKey.String("OK")

	OTelStatusCodeError = OTelStatusCodeKey.String("ERROR")
)

func OTelStatusDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSRequestIDKey = attribute.Key("aws.request_id")
)

func AWSRequestID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSS3BucketKey = attribute.Key("aws.s3.bucket")

	AWSS3CopySourceKey = attribute.Key("aws.s3.copy_source")

	AWSS3DeleteKey = attribute.Key("aws.s3.delete")

	AWSS3KeyKey = attribute.Key("aws.s3.key")

	AWSS3PartNumberKey = attribute.Key("aws.s3.part_number")

	AWSS3UploadIDKey = attribute.Key("aws.s3.upload_id")
)

func AWSS3Bucket(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSS3CopySource(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSS3Delete(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSS3Key(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSS3PartNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSS3UploadID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	GraphqlDocumentKey = attribute.Key("graphql.document")

	GraphqlOperationNameKey = attribute.Key("graphql.operation.name")

	GraphqlOperationTypeKey = attribute.Key("graphql.operation.type")
)

var (
	GraphqlOperationTypeQuery = GraphqlOperationTypeKey.String("query")

	GraphqlOperationTypeMutation = GraphqlOperationTypeKey.String("mutation")

	GraphqlOperationTypeSubscription = GraphqlOperationTypeKey.String("subscription")
)

func GraphqlDocument(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GraphqlOperationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
