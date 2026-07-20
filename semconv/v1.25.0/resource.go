package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	AWSECSTaskIDKey = attribute.Key("aws.ecs.task.id")

	AWSECSClusterARNKey = attribute.Key("aws.ecs.cluster.arn")

	AWSECSContainerARNKey = attribute.Key("aws.ecs.container.arn")

	AWSECSLaunchtypeKey = attribute.Key("aws.ecs.launchtype")

	AWSECSTaskARNKey = attribute.Key("aws.ecs.task.arn")

	AWSECSTaskFamilyKey = attribute.Key("aws.ecs.task.family")

	AWSECSTaskRevisionKey = attribute.Key("aws.ecs.task.revision")
)

var (
	AWSECSLaunchtypeEC2 = AWSECSLaunchtypeKey.String("ec2")

	AWSECSLaunchtypeFargate = AWSECSLaunchtypeKey.String("fargate")
)

func AWSECSTaskID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSClusterARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSContainerARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSTaskARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSTaskFamily(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSTaskRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSEKSClusterARNKey = attribute.Key("aws.eks.cluster.arn")
)

func AWSEKSClusterARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSLogGroupARNsKey = attribute.Key("aws.log.group.arns")

	AWSLogGroupNamesKey = attribute.Key("aws.log.group.names")

	AWSLogStreamARNsKey = attribute.Key("aws.log.stream.arns")

	AWSLogStreamNamesKey = attribute.Key("aws.log.stream.names")
)

func AWSLogGroupARNs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLogGroupNames(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLogStreamARNs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLogStreamNames(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HerokuAppIDKey = attribute.Key("heroku.app.id")

	HerokuReleaseCommitKey = attribute.Key("heroku.release.commit")

	HerokuReleaseCreationTimestampKey = attribute.Key("heroku.release.creation_timestamp")
)

func HerokuAppID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HerokuReleaseCommit(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HerokuReleaseCreationTimestamp(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	WebEngineNameKey = attribute.Key("webengine.name")

	WebEngineDescriptionKey = attribute.Key("webengine.description")

	WebEngineVersionKey = attribute.Key("webengine.version")
)

func WebEngineName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func WebEngineDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func WebEngineVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	OTelScopeNameKey = attribute.Key("otel.scope.name")

	OTelScopeVersionKey = attribute.Key("otel.scope.version")
)

func OTelScopeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OTelScopeVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	OTelLibraryNameKey = attribute.Key("otel.library.name")

	OTelLibraryVersionKey = attribute.Key("otel.library.version")
)

func OTelLibraryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OTelLibraryVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
