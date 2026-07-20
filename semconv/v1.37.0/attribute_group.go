package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	AndroidAppStateKey = attribute.Key("android.app.state")

	AndroidOSAPILevelKey = attribute.Key("android.os.api_level")
)

func AndroidOSAPILevel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	AndroidAppStateCreated = AndroidAppStateKey.String("created")

	AndroidAppStateBackground = AndroidAppStateKey.String("background")

	AndroidAppStateForeground = AndroidAppStateKey.String("foreground")
)

const (
	AppBuildIDKey = attribute.Key("app.build_id")

	AppInstallationIDKey = attribute.Key("app.installation.id")

	AppJankFrameCountKey = attribute.Key("app.jank.frame_count")

	AppJankPeriodKey = attribute.Key("app.jank.period")

	AppJankThresholdKey = attribute.Key("app.jank.threshold")

	AppScreenCoordinateXKey = attribute.Key("app.screen.coordinate.x")

	AppScreenCoordinateYKey = attribute.Key("app.screen.coordinate.y")

	AppWidgetIDKey = attribute.Key("app.widget.id")

	AppWidgetNameKey = attribute.Key("app.widget.name")
)

func AppBuildID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AppInstallationID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AppJankFrameCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AppJankPeriod(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AppJankThreshold(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AppScreenCoordinateX(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AppScreenCoordinateY(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AppWidgetID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AppWidgetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ArtifactAttestationFilenameKey = attribute.Key("artifact.attestation.filename")

	ArtifactAttestationHashKey = attribute.Key("artifact.attestation.hash")

	ArtifactAttestationIDKey = attribute.Key("artifact.attestation.id")

	ArtifactFilenameKey = attribute.Key("artifact.filename")

	ArtifactHashKey = attribute.Key("artifact.hash")

	ArtifactPurlKey = attribute.Key("artifact.purl")

	ArtifactVersionKey = attribute.Key("artifact.version")
)

func ArtifactAttestationFilename(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ArtifactAttestationHash(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ArtifactAttestationID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ArtifactFilename(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ArtifactHash(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ArtifactPurl(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ArtifactVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSBedrockGuardrailIDKey = attribute.Key("aws.bedrock.guardrail.id")

	AWSBedrockKnowledgeBaseIDKey = attribute.Key("aws.bedrock.knowledge_base.id")

	AWSDynamoDBAttributeDefinitionsKey = attribute.Key("aws.dynamodb.attribute_definitions")

	AWSDynamoDBAttributesToGetKey = attribute.Key("aws.dynamodb.attributes_to_get")

	AWSDynamoDBConsistentReadKey = attribute.Key("aws.dynamodb.consistent_read")

	AWSDynamoDBConsumedCapacityKey = attribute.Key("aws.dynamodb.consumed_capacity")

	AWSDynamoDBCountKey = attribute.Key("aws.dynamodb.count")

	AWSDynamoDBExclusiveStartTableKey = attribute.Key("aws.dynamodb.exclusive_start_table")

	AWSDynamoDBGlobalSecondaryIndexUpdatesKey = attribute.Key("aws.dynamodb.global_secondary_index_updates")

	AWSDynamoDBGlobalSecondaryIndexesKey = attribute.Key("aws.dynamodb.global_secondary_indexes")

	AWSDynamoDBIndexNameKey = attribute.Key("aws.dynamodb.index_name")

	AWSDynamoDBItemCollectionMetricsKey = attribute.Key("aws.dynamodb.item_collection_metrics")

	AWSDynamoDBLimitKey = attribute.Key("aws.dynamodb.limit")

	AWSDynamoDBLocalSecondaryIndexesKey = attribute.Key("aws.dynamodb.local_secondary_indexes")

	AWSDynamoDBProjectionKey = attribute.Key("aws.dynamodb.projection")

	AWSDynamoDBProvisionedReadCapacityKey = attribute.Key("aws.dynamodb.provisioned_read_capacity")

	AWSDynamoDBProvisionedWriteCapacityKey = attribute.Key("aws.dynamodb.provisioned_write_capacity")

	AWSDynamoDBScanForwardKey = attribute.Key("aws.dynamodb.scan_forward")

	AWSDynamoDBScannedCountKey = attribute.Key("aws.dynamodb.scanned_count")

	AWSDynamoDBSegmentKey = attribute.Key("aws.dynamodb.segment")

	AWSDynamoDBSelectKey = attribute.Key("aws.dynamodb.select")

	AWSDynamoDBTableCountKey = attribute.Key("aws.dynamodb.table_count")

	AWSDynamoDBTableNamesKey = attribute.Key("aws.dynamodb.table_names")

	AWSDynamoDBTotalSegmentsKey = attribute.Key("aws.dynamodb.total_segments")

	AWSECSClusterARNKey = attribute.Key("aws.ecs.cluster.arn")

	AWSECSContainerARNKey = attribute.Key("aws.ecs.container.arn")

	AWSECSLaunchtypeKey = attribute.Key("aws.ecs.launchtype")

	AWSECSTaskARNKey = attribute.Key("aws.ecs.task.arn")

	AWSECSTaskFamilyKey = attribute.Key("aws.ecs.task.family")

	AWSECSTaskIDKey = attribute.Key("aws.ecs.task.id")

	AWSECSTaskRevisionKey = attribute.Key("aws.ecs.task.revision")

	AWSEKSClusterARNKey = attribute.Key("aws.eks.cluster.arn")

	AWSExtendedRequestIDKey = attribute.Key("aws.extended_request_id")

	AWSKinesisStreamNameKey = attribute.Key("aws.kinesis.stream_name")

	AWSLambdaInvokedARNKey = attribute.Key("aws.lambda.invoked_arn")

	AWSLambdaResourceMappingIDKey = attribute.Key("aws.lambda.resource_mapping.id")

	AWSLogGroupARNsKey = attribute.Key("aws.log.group.arns")

	AWSLogGroupNamesKey = attribute.Key("aws.log.group.names")

	AWSLogStreamARNsKey = attribute.Key("aws.log.stream.arns")

	AWSLogStreamNamesKey = attribute.Key("aws.log.stream.names")

	AWSRequestIDKey = attribute.Key("aws.request_id")

	AWSS3BucketKey = attribute.Key("aws.s3.bucket")

	AWSS3CopySourceKey = attribute.Key("aws.s3.copy_source")

	AWSS3DeleteKey = attribute.Key("aws.s3.delete")

	AWSS3KeyKey = attribute.Key("aws.s3.key")

	AWSS3PartNumberKey = attribute.Key("aws.s3.part_number")

	AWSS3UploadIDKey = attribute.Key("aws.s3.upload_id")

	AWSSecretsmanagerSecretARNKey = attribute.Key("aws.secretsmanager.secret.arn")

	AWSSNSTopicARNKey = attribute.Key("aws.sns.topic.arn")

	AWSSQSQueueURLKey = attribute.Key("aws.sqs.queue.url")

	AWSStepFunctionsActivityARNKey = attribute.Key("aws.step_functions.activity.arn")

	AWSStepFunctionsStateMachineARNKey = attribute.Key("aws.step_functions.state_machine.arn")
)

func AWSBedrockGuardrailID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSBedrockKnowledgeBaseID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBAttributeDefinitions(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBAttributesToGet(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBConsistentRead(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBConsumedCapacity(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBExclusiveStartTable(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBGlobalSecondaryIndexUpdates(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBGlobalSecondaryIndexes(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBIndexName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBItemCollectionMetrics(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBLimit(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBLocalSecondaryIndexes(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBProjection(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBProvisionedReadCapacity(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBProvisionedWriteCapacity(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBScanForward(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBScannedCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBSegment(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBSelect(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBTableCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBTableNames(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBTotalSegments(val int) attribute.KeyValue {
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

func AWSECSTaskID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSECSTaskRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSEKSClusterARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSExtendedRequestID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSKinesisStreamName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLambdaInvokedARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSLambdaResourceMappingID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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

func AWSRequestID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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

func AWSSecretsmanagerSecretARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSSNSTopicARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSSQSQueueURL(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSStepFunctionsActivityARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSStepFunctionsStateMachineARN(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	AWSECSLaunchtypeEC2 = AWSECSLaunchtypeKey.String("ec2")

	AWSECSLaunchtypeFargate = AWSECSLaunchtypeKey.String("fargate")
)

const (
	AzureClientIDKey = attribute.Key("azure.client.id")

	AzureCosmosDBConnectionModeKey = attribute.Key("azure.cosmosdb.connection.mode")

	AzureCosmosDBConsistencyLevelKey = attribute.Key("azure.cosmosdb.consistency.level")

	AzureCosmosDBOperationContactedRegionsKey = attribute.Key("azure.cosmosdb.operation.contacted_regions")

	AzureCosmosDBOperationRequestChargeKey = attribute.Key("azure.cosmosdb.operation.request_charge")

	AzureCosmosDBRequestBodySizeKey = attribute.Key("azure.cosmosdb.request.body.size")

	AzureCosmosDBResponseSubStatusCodeKey = attribute.Key("azure.cosmosdb.response.sub_status_code")

	AzureResourceProviderNamespaceKey = attribute.Key("azure.resource_provider.namespace")

	AzureServiceRequestIDKey = attribute.Key("azure.service.request.id")
)

func AzureClientID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AzureCosmosDBOperationContactedRegions(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AzureCosmosDBOperationRequestCharge(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AzureCosmosDBRequestBodySize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AzureCosmosDBResponseSubStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AzureResourceProviderNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AzureServiceRequestID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	AzureCosmosDBConnectionModeGateway = AzureCosmosDBConnectionModeKey.String("gateway")

	AzureCosmosDBConnectionModeDirect = AzureCosmosDBConnectionModeKey.String("direct")
)

var (
	AzureCosmosDBConsistencyLevelStrong = AzureCosmosDBConsistencyLevelKey.String("Strong")

	AzureCosmosDBConsistencyLevelBoundedStaleness = AzureCosmosDBConsistencyLevelKey.String("BoundedStaleness")

	AzureCosmosDBConsistencyLevelSession = AzureCosmosDBConsistencyLevelKey.String("Session")

	AzureCosmosDBConsistencyLevelEventual = AzureCosmosDBConsistencyLevelKey.String("Eventual")

	AzureCosmosDBConsistencyLevelConsistentPrefix = AzureCosmosDBConsistencyLevelKey.String("ConsistentPrefix")
)

const (
	BrowserBrandsKey = attribute.Key("browser.brands")

	BrowserLanguageKey = attribute.Key("browser.language")

	BrowserMobileKey = attribute.Key("browser.mobile")

	BrowserPlatformKey = attribute.Key("browser.platform")
)

func BrowserBrands(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func BrowserLanguage(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func BrowserMobile(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func BrowserPlatform(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	CassandraConsistencyLevelKey = attribute.Key("cassandra.consistency.level")

	CassandraCoordinatorDCKey = attribute.Key("cassandra.coordinator.dc")

	CassandraCoordinatorIDKey = attribute.Key("cassandra.coordinator.id")

	CassandraPageSizeKey = attribute.Key("cassandra.page.size")

	CassandraQueryIdempotentKey = attribute.Key("cassandra.query.idempotent")

	CassandraSpeculativeExecutionCountKey = attribute.Key("cassandra.speculative_execution.count")
)

func CassandraCoordinatorDC(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CassandraCoordinatorID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CassandraPageSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CassandraQueryIdempotent(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CassandraSpeculativeExecutionCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	CassandraConsistencyLevelAll = CassandraConsistencyLevelKey.String("all")

	CassandraConsistencyLevelEachQuorum = CassandraConsistencyLevelKey.String("each_quorum")

	CassandraConsistencyLevelQuorum = CassandraConsistencyLevelKey.String("quorum")

	CassandraConsistencyLevelLocalQuorum = CassandraConsistencyLevelKey.String("local_quorum")

	CassandraConsistencyLevelOne = CassandraConsistencyLevelKey.String("one")

	CassandraConsistencyLevelTwo = CassandraConsistencyLevelKey.String("two")

	CassandraConsistencyLevelThree = CassandraConsistencyLevelKey.String("three")

	CassandraConsistencyLevelLocalOne = CassandraConsistencyLevelKey.String("local_one")

	CassandraConsistencyLevelAny = CassandraConsistencyLevelKey.String("any")

	CassandraConsistencyLevelSerial = CassandraConsistencyLevelKey.String("serial")

	CassandraConsistencyLevelLocalSerial = CassandraConsistencyLevelKey.String("local_serial")
)

const (
	CICDPipelineActionNameKey = attribute.Key("cicd.pipeline.action.name")

	CICDPipelineNameKey = attribute.Key("cicd.pipeline.name")

	CICDPipelineResultKey = attribute.Key("cicd.pipeline.result")

	CICDPipelineRunIDKey = attribute.Key("cicd.pipeline.run.id")

	CICDPipelineRunStateKey = attribute.Key("cicd.pipeline.run.state")

	CICDPipelineRunURLFullKey = attribute.Key("cicd.pipeline.run.url.full")

	CICDPipelineTaskNameKey = attribute.Key("cicd.pipeline.task.name")

	CICDPipelineTaskRunIDKey = attribute.Key("cicd.pipeline.task.run.id")

	CICDPipelineTaskRunResultKey = attribute.Key("cicd.pipeline.task.run.result")

	CICDPipelineTaskRunURLFullKey = attribute.Key("cicd.pipeline.task.run.url.full")

	CICDPipelineTaskTypeKey = attribute.Key("cicd.pipeline.task.type")

	CICDSystemComponentKey = attribute.Key("cicd.system.component")

	CICDWorkerIDKey = attribute.Key("cicd.worker.id")

	CICDWorkerNameKey = attribute.Key("cicd.worker.name")

	CICDWorkerStateKey = attribute.Key("cicd.worker.state")

	CICDWorkerURLFullKey = attribute.Key("cicd.worker.url.full")
)

func CICDPipelineName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CICDPipelineRunID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CICDPipelineRunURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CICDPipelineTaskName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CICDPipelineTaskRunID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CICDPipelineTaskRunURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CICDSystemComponent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CICDWorkerID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CICDWorkerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CICDWorkerURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	CICDPipelineActionNameBuild = CICDPipelineActionNameKey.String("BUILD")

	CICDPipelineActionNameRun = CICDPipelineActionNameKey.String("RUN")

	CICDPipelineActionNameSync = CICDPipelineActionNameKey.String("SYNC")
)

var (
	CICDPipelineResultSuccess = CICDPipelineResultKey.String("success")

	CICDPipelineResultFailure = CICDPipelineResultKey.String("failure")

	CICDPipelineResultError = CICDPipelineResultKey.String("error")

	CICDPipelineResultTimeout = CICDPipelineResultKey.String("timeout")

	CICDPipelineResultCancellation = CICDPipelineResultKey.String("cancellation")

	CICDPipelineResultSkip = CICDPipelineResultKey.String("skip")
)

var (
	CICDPipelineRunStatePending = CICDPipelineRunStateKey.String("pending")

	CICDPipelineRunStateExecuting = CICDPipelineRunStateKey.String("executing")

	CICDPipelineRunStateFinalizing = CICDPipelineRunStateKey.String("finalizing")
)

var (
	CICDPipelineTaskRunResultSuccess = CICDPipelineTaskRunResultKey.String("success")

	CICDPipelineTaskRunResultFailure = CICDPipelineTaskRunResultKey.String("failure")

	CICDPipelineTaskRunResultError = CICDPipelineTaskRunResultKey.String("error")

	CICDPipelineTaskRunResultTimeout = CICDPipelineTaskRunResultKey.String("timeout")

	CICDPipelineTaskRunResultCancellation = CICDPipelineTaskRunResultKey.String("cancellation")

	CICDPipelineTaskRunResultSkip = CICDPipelineTaskRunResultKey.String("skip")
)

var (
	CICDPipelineTaskTypeBuild = CICDPipelineTaskTypeKey.String("build")

	CICDPipelineTaskTypeTest = CICDPipelineTaskTypeKey.String("test")

	CICDPipelineTaskTypeDeploy = CICDPipelineTaskTypeKey.String("deploy")
)

var (
	CICDWorkerStateAvailable = CICDWorkerStateKey.String("available")

	CICDWorkerStateBusy = CICDWorkerStateKey.String("busy")

	CICDWorkerStateOffline = CICDWorkerStateKey.String("offline")
)

const (
	ClientAddressKey = attribute.Key("client.address")

	ClientPortKey = attribute.Key("client.port")
)

func ClientAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ClientPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	CloudAccountIDKey = attribute.Key("cloud.account.id")

	CloudAvailabilityZoneKey = attribute.Key("cloud.availability_zone")

	CloudPlatformKey = attribute.Key("cloud.platform")

	CloudProviderKey = attribute.Key("cloud.provider")

	CloudRegionKey = attribute.Key("cloud.region")

	CloudResourceIDKey = attribute.Key("cloud.resource_id")
)

func CloudAccountID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudAvailabilityZone(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudRegion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudResourceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	CloudPlatformAlibabaCloudECS = CloudPlatformKey.String("alibaba_cloud_ecs")

	CloudPlatformAlibabaCloudFC = CloudPlatformKey.String("alibaba_cloud_fc")

	CloudPlatformAlibabaCloudOpenShift = CloudPlatformKey.String("alibaba_cloud_openshift")

	CloudPlatformAWSEC2 = CloudPlatformKey.String("aws_ec2")

	CloudPlatformAWSECS = CloudPlatformKey.String("aws_ecs")

	CloudPlatformAWSEKS = CloudPlatformKey.String("aws_eks")

	CloudPlatformAWSLambda = CloudPlatformKey.String("aws_lambda")

	CloudPlatformAWSElasticBeanstalk = CloudPlatformKey.String("aws_elastic_beanstalk")

	CloudPlatformAWSAppRunner = CloudPlatformKey.String("aws_app_runner")

	CloudPlatformAWSOpenShift = CloudPlatformKey.String("aws_openshift")

	CloudPlatformAzureVM = CloudPlatformKey.String("azure.vm")

	CloudPlatformAzureContainerApps = CloudPlatformKey.String("azure.container_apps")

	CloudPlatformAzureContainerInstances = CloudPlatformKey.String("azure.container_instances")

	CloudPlatformAzureAKS = CloudPlatformKey.String("azure.aks")

	CloudPlatformAzureFunctions = CloudPlatformKey.String("azure.functions")

	CloudPlatformAzureAppService = CloudPlatformKey.String("azure.app_service")

	CloudPlatformAzureOpenShift = CloudPlatformKey.String("azure.openshift")

	CloudPlatformGCPBareMetalSolution = CloudPlatformKey.String("gcp_bare_metal_solution")

	CloudPlatformGCPComputeEngine = CloudPlatformKey.String("gcp_compute_engine")

	CloudPlatformGCPCloudRun = CloudPlatformKey.String("gcp_cloud_run")

	CloudPlatformGCPKubernetesEngine = CloudPlatformKey.String("gcp_kubernetes_engine")

	CloudPlatformGCPCloudFunctions = CloudPlatformKey.String("gcp_cloud_functions")

	CloudPlatformGCPAppEngine = CloudPlatformKey.String("gcp_app_engine")

	CloudPlatformGCPOpenShift = CloudPlatformKey.String("gcp_openshift")

	CloudPlatformIBMCloudOpenShift = CloudPlatformKey.String("ibm_cloud_openshift")

	CloudPlatformOracleCloudCompute = CloudPlatformKey.String("oracle_cloud_compute")

	CloudPlatformOracleCloudOKE = CloudPlatformKey.String("oracle_cloud_oke")

	CloudPlatformTencentCloudCVM = CloudPlatformKey.String("tencent_cloud_cvm")

	CloudPlatformTencentCloudEKS = CloudPlatformKey.String("tencent_cloud_eks")

	CloudPlatformTencentCloudSCF = CloudPlatformKey.String("tencent_cloud_scf")
)

var (
	CloudProviderAlibabaCloud = CloudProviderKey.String("alibaba_cloud")

	CloudProviderAWS = CloudProviderKey.String("aws")

	CloudProviderAzure = CloudProviderKey.String("azure")

	CloudProviderGCP = CloudProviderKey.String("gcp")

	CloudProviderHeroku = CloudProviderKey.String("heroku")

	CloudProviderIBMCloud = CloudProviderKey.String("ibm_cloud")

	CloudProviderOracleCloud = CloudProviderKey.String("oracle_cloud")

	CloudProviderTencentCloud = CloudProviderKey.String("tencent_cloud")
)

const (
	CloudEventsEventIDKey = attribute.Key("cloudevents.event_id")

	CloudEventsEventSourceKey = attribute.Key("cloudevents.event_source")

	CloudEventsEventSpecVersionKey = attribute.Key("cloudevents.event_spec_version")

	CloudEventsEventSubjectKey = attribute.Key("cloudevents.event_subject")

	CloudEventsEventTypeKey = attribute.Key("cloudevents.event_type")
)

func CloudEventsEventID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudEventsEventSource(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudEventsEventSpecVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudEventsEventSubject(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudEventsEventType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	CloudFoundryAppIDKey = attribute.Key("cloudfoundry.app.id")

	CloudFoundryAppInstanceIDKey = attribute.Key("cloudfoundry.app.instance.id")

	CloudFoundryAppNameKey = attribute.Key("cloudfoundry.app.name")

	CloudFoundryOrgIDKey = attribute.Key("cloudfoundry.org.id")

	CloudFoundryOrgNameKey = attribute.Key("cloudfoundry.org.name")

	CloudFoundryProcessIDKey = attribute.Key("cloudfoundry.process.id")

	CloudFoundryProcessTypeKey = attribute.Key("cloudfoundry.process.type")

	CloudFoundrySpaceIDKey = attribute.Key("cloudfoundry.space.id")

	CloudFoundrySpaceNameKey = attribute.Key("cloudfoundry.space.name")

	CloudFoundrySystemIDKey = attribute.Key("cloudfoundry.system.id")

	CloudFoundrySystemInstanceIDKey = attribute.Key("cloudfoundry.system.instance.id")
)

func CloudFoundryAppID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundryAppInstanceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundryAppName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundryOrgID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundryOrgName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundryProcessID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundryProcessType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundrySpaceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundrySpaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundrySystemID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudFoundrySystemInstanceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	CodeColumnNumberKey = attribute.Key("code.column.number")

	CodeFilePathKey = attribute.Key("code.file.path")

	CodeFunctionNameKey = attribute.Key("code.function.name")

	CodeLineNumberKey = attribute.Key("code.line.number")

	CodeStacktraceKey = attribute.Key("code.stacktrace")
)

func CodeColumnNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeFilePath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeFunctionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeLineNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeStacktrace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ContainerCommandKey = attribute.Key("container.command")

	ContainerCommandArgsKey = attribute.Key("container.command_args")

	ContainerCommandLineKey = attribute.Key("container.command_line")

	ContainerCSIPluginNameKey = attribute.Key("container.csi.plugin.name")

	ContainerCSIVolumeIDKey = attribute.Key("container.csi.volume.id")

	ContainerIDKey = attribute.Key("container.id")

	ContainerImageIDKey = attribute.Key("container.image.id")

	ContainerImageNameKey = attribute.Key("container.image.name")

	ContainerImageRepoDigestsKey = attribute.Key("container.image.repo_digests")

	ContainerImageTagsKey = attribute.Key("container.image.tags")

	ContainerNameKey = attribute.Key("container.name")

	ContainerRuntimeDescriptionKey = attribute.Key("container.runtime.description")

	ContainerRuntimeNameKey = attribute.Key("container.runtime.name")

	ContainerRuntimeVersionKey = attribute.Key("container.runtime.version")
)

func ContainerCommand(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerCommandArgs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerCommandLine(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerCSIPluginName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerCSIVolumeID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageRepoDigests(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerImageTags(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerRuntimeDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerRuntimeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ContainerRuntimeVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	CPULogicalNumberKey = attribute.Key("cpu.logical_number")

	CPUModeKey = attribute.Key("cpu.mode")
)

func CPULogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	CPUModeUser = CPUModeKey.String("user")

	CPUModeSystem = CPUModeKey.String("system")

	CPUModeNice = CPUModeKey.String("nice")

	CPUModeIdle = CPUModeKey.String("idle")

	CPUModeIOWait = CPUModeKey.String("iowait")

	CPUModeInterrupt = CPUModeKey.String("interrupt")

	CPUModeSteal = CPUModeKey.String("steal")

	CPUModeKernel = CPUModeKey.String("kernel")
)

const (
	DBClientConnectionPoolNameKey = attribute.Key("db.client.connection.pool.name")

	DBClientConnectionStateKey = attribute.Key("db.client.connection.state")

	DBCollectionNameKey = attribute.Key("db.collection.name")

	DBNamespaceKey = attribute.Key("db.namespace")

	DBOperationBatchSizeKey = attribute.Key("db.operation.batch.size")

	DBOperationNameKey = attribute.Key("db.operation.name")

	DBQuerySummaryKey = attribute.Key("db.query.summary")

	DBQueryTextKey = attribute.Key("db.query.text")

	DBResponseReturnedRowsKey = attribute.Key("db.response.returned_rows")

	DBResponseStatusCodeKey = attribute.Key("db.response.status_code")

	DBStoredProcedureNameKey = attribute.Key("db.stored_procedure.name")

	DBSystemNameKey = attribute.Key("db.system.name")
)

func DBClientConnectionPoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCollectionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBOperationBatchSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBOperationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBOperationParameter(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBQueryParameter(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBQuerySummary(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBQueryText(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBResponseReturnedRows(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBStoredProcedureName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	DBClientConnectionStateIdle = DBClientConnectionStateKey.String("idle")

	DBClientConnectionStateUsed = DBClientConnectionStateKey.String("used")
)

var (
	DBSystemNameOtherSQL = DBSystemNameKey.String("other_sql")

	DBSystemNameSoftwareagAdabas = DBSystemNameKey.String("softwareag.adabas")

	DBSystemNameActianIngres = DBSystemNameKey.String("actian.ingres")

	DBSystemNameAWSDynamoDB = DBSystemNameKey.String("aws.dynamodb")

	DBSystemNameAWSRedshift = DBSystemNameKey.String("aws.redshift")

	DBSystemNameAzureCosmosDB = DBSystemNameKey.String("azure.cosmosdb")

	DBSystemNameIntersystemsCache = DBSystemNameKey.String("intersystems.cache")

	DBSystemNameCassandra = DBSystemNameKey.String("cassandra")

	DBSystemNameClickHouse = DBSystemNameKey.String("clickhouse")

	DBSystemNameCockroachDB = DBSystemNameKey.String("cockroachdb")

	DBSystemNameCouchbase = DBSystemNameKey.String("couchbase")

	DBSystemNameCouchDB = DBSystemNameKey.String("couchdb")

	DBSystemNameDerby = DBSystemNameKey.String("derby")

	DBSystemNameElasticsearch = DBSystemNameKey.String("elasticsearch")

	DBSystemNameFirebirdSQL = DBSystemNameKey.String("firebirdsql")

	DBSystemNameGCPSpanner = DBSystemNameKey.String("gcp.spanner")

	DBSystemNameGeode = DBSystemNameKey.String("geode")

	DBSystemNameH2database = DBSystemNameKey.String("h2database")

	DBSystemNameHBase = DBSystemNameKey.String("hbase")

	DBSystemNameHive = DBSystemNameKey.String("hive")

	DBSystemNameHSQLDB = DBSystemNameKey.String("hsqldb")

	DBSystemNameIBMDB2 = DBSystemNameKey.String("ibm.db2")

	DBSystemNameIBMInformix = DBSystemNameKey.String("ibm.informix")

	DBSystemNameIBMNetezza = DBSystemNameKey.String("ibm.netezza")

	DBSystemNameInfluxDB = DBSystemNameKey.String("influxdb")

	DBSystemNameInstantDB = DBSystemNameKey.String("instantdb")

	DBSystemNameMariaDB = DBSystemNameKey.String("mariadb")

	DBSystemNameMemcached = DBSystemNameKey.String("memcached")

	DBSystemNameMongoDB = DBSystemNameKey.String("mongodb")

	DBSystemNameMicrosoftSQLServer = DBSystemNameKey.String("microsoft.sql_server")

	DBSystemNameMySQL = DBSystemNameKey.String("mysql")

	DBSystemNameNeo4j = DBSystemNameKey.String("neo4j")

	DBSystemNameOpenSearch = DBSystemNameKey.String("opensearch")

	DBSystemNameOracleDB = DBSystemNameKey.String("oracle.db")

	DBSystemNamePostgreSQL = DBSystemNameKey.String("postgresql")

	DBSystemNameRedis = DBSystemNameKey.String("redis")

	DBSystemNameSAPHANA = DBSystemNameKey.String("sap.hana")

	DBSystemNameSAPMaxDB = DBSystemNameKey.String("sap.maxdb")

	DBSystemNameSQLite = DBSystemNameKey.String("sqlite")

	DBSystemNameTeradata = DBSystemNameKey.String("teradata")

	DBSystemNameTrino = DBSystemNameKey.String("trino")
)

const (
	DeploymentEnvironmentNameKey = attribute.Key("deployment.environment.name")

	DeploymentIDKey = attribute.Key("deployment.id")

	DeploymentNameKey = attribute.Key("deployment.name")

	DeploymentStatusKey = attribute.Key("deployment.status")
)

func DeploymentEnvironmentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DeploymentID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DeploymentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	DeploymentStatusFailed = DeploymentStatusKey.String("failed")

	DeploymentStatusSucceeded = DeploymentStatusKey.String("succeeded")
)

const (
	DestinationAddressKey = attribute.Key("destination.address")

	DestinationPortKey = attribute.Key("destination.port")
)

func DestinationAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DestinationPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DeviceIDKey = attribute.Key("device.id")

	DeviceManufacturerKey = attribute.Key("device.manufacturer")

	DeviceModelIdentifierKey = attribute.Key("device.model.identifier")

	DeviceModelNameKey = attribute.Key("device.model.name")
)

func DeviceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DeviceManufacturer(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DeviceModelIdentifier(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DeviceModelName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DiskIODirectionKey = attribute.Key("disk.io.direction")
)

var (
	DiskIODirectionRead = DiskIODirectionKey.String("read")

	DiskIODirectionWrite = DiskIODirectionKey.String("write")
)

const (
	DNSAnswersKey = attribute.Key("dns.answers")

	DNSQuestionNameKey = attribute.Key("dns.question.name")
)

func DNSAnswers(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DNSQuestionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ElasticsearchNodeNameKey = attribute.Key("elasticsearch.node.name")
)

func ElasticsearchNodeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	EnduserIDKey = attribute.Key("enduser.id")

	EnduserPseudoIDKey = attribute.Key("enduser.pseudo.id")
)

func EnduserID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func EnduserPseudoID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ErrorMessageKey = attribute.Key("error.message")

	ErrorTypeKey = attribute.Key("error.type")
)

func ErrorMessage(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	ErrorTypeOther = ErrorTypeKey.String("_OTHER")
)

const (
	ExceptionMessageKey = attribute.Key("exception.message")

	ExceptionStacktraceKey = attribute.Key("exception.stacktrace")

	ExceptionTypeKey = attribute.Key("exception.type")
)

func ExceptionMessage(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ExceptionStacktrace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ExceptionType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	FaaSColdstartKey = attribute.Key("faas.coldstart")

	FaaSCronKey = attribute.Key("faas.cron")

	FaaSDocumentCollectionKey = attribute.Key("faas.document.collection")

	FaaSDocumentNameKey = attribute.Key("faas.document.name")

	FaaSDocumentOperationKey = attribute.Key("faas.document.operation")

	FaaSDocumentTimeKey = attribute.Key("faas.document.time")

	FaaSInstanceKey = attribute.Key("faas.instance")

	FaaSInvocationIDKey = attribute.Key("faas.invocation_id")

	FaaSInvokedNameKey = attribute.Key("faas.invoked_name")

	FaaSInvokedProviderKey = attribute.Key("faas.invoked_provider")

	FaaSInvokedRegionKey = attribute.Key("faas.invoked_region")

	FaaSMaxMemoryKey = attribute.Key("faas.max_memory")

	FaaSNameKey = attribute.Key("faas.name")

	FaaSTimeKey = attribute.Key("faas.time")

	FaaSTriggerKey = attribute.Key("faas.trigger")

	FaaSVersionKey = attribute.Key("faas.version")
)

func FaaSColdstart(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSCron(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSDocumentCollection(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSDocumentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSDocumentTime(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSInstance(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSInvocationID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSInvokedName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSInvokedRegion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSMaxMemory(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSTime(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	FaaSDocumentOperationInsert = FaaSDocumentOperationKey.String("insert")

	FaaSDocumentOperationEdit = FaaSDocumentOperationKey.String("edit")

	FaaSDocumentOperationDelete = FaaSDocumentOperationKey.String("delete")
)

var (
	FaaSInvokedProviderAlibabaCloud = FaaSInvokedProviderKey.String("alibaba_cloud")

	FaaSInvokedProviderAWS = FaaSInvokedProviderKey.String("aws")

	FaaSInvokedProviderAzure = FaaSInvokedProviderKey.String("azure")

	FaaSInvokedProviderGCP = FaaSInvokedProviderKey.String("gcp")

	FaaSInvokedProviderTencentCloud = FaaSInvokedProviderKey.String("tencent_cloud")
)

var (
	FaaSTriggerDatasource = FaaSTriggerKey.String("datasource")

	FaaSTriggerHTTP = FaaSTriggerKey.String("http")

	FaaSTriggerPubSub = FaaSTriggerKey.String("pubsub")

	FaaSTriggerTimer = FaaSTriggerKey.String("timer")

	FaaSTriggerOther = FaaSTriggerKey.String("other")
)

const (
	FeatureFlagContextIDKey = attribute.Key("feature_flag.context.id")

	FeatureFlagKeyKey = attribute.Key("feature_flag.key")

	FeatureFlagProviderNameKey = attribute.Key("feature_flag.provider.name")

	FeatureFlagResultReasonKey = attribute.Key("feature_flag.result.reason")

	FeatureFlagResultValueKey = attribute.Key("feature_flag.result.value")

	FeatureFlagResultVariantKey = attribute.Key("feature_flag.result.variant")

	FeatureFlagSetIDKey = attribute.Key("feature_flag.set.id")

	FeatureFlagVersionKey = attribute.Key("feature_flag.version")
)

func FeatureFlagContextID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FeatureFlagKey(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FeatureFlagProviderName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FeatureFlagResultVariant(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FeatureFlagSetID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FeatureFlagVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	FeatureFlagResultReasonStatic = FeatureFlagResultReasonKey.String("static")

	FeatureFlagResultReasonDefault = FeatureFlagResultReasonKey.String("default")

	FeatureFlagResultReasonTargetingMatch = FeatureFlagResultReasonKey.String("targeting_match")

	FeatureFlagResultReasonSplit = FeatureFlagResultReasonKey.String("split")

	FeatureFlagResultReasonCached = FeatureFlagResultReasonKey.String("cached")

	FeatureFlagResultReasonDisabled = FeatureFlagResultReasonKey.String("disabled")

	FeatureFlagResultReasonUnknown = FeatureFlagResultReasonKey.String("unknown")

	FeatureFlagResultReasonStale = FeatureFlagResultReasonKey.String("stale")

	FeatureFlagResultReasonError = FeatureFlagResultReasonKey.String("error")
)

const (
	FileAccessedKey = attribute.Key("file.accessed")

	FileAttributesKey = attribute.Key("file.attributes")

	FileChangedKey = attribute.Key("file.changed")

	FileCreatedKey = attribute.Key("file.created")

	FileDirectoryKey = attribute.Key("file.directory")

	FileExtensionKey = attribute.Key("file.extension")

	FileForkNameKey = attribute.Key("file.fork_name")

	FileGroupIDKey = attribute.Key("file.group.id")

	FileGroupNameKey = attribute.Key("file.group.name")

	FileInodeKey = attribute.Key("file.inode")

	FileModeKey = attribute.Key("file.mode")

	FileModifiedKey = attribute.Key("file.modified")

	FileNameKey = attribute.Key("file.name")

	FileOwnerIDKey = attribute.Key("file.owner.id")

	FileOwnerNameKey = attribute.Key("file.owner.name")

	FilePathKey = attribute.Key("file.path")

	FileSizeKey = attribute.Key("file.size")

	FileSymbolicLinkTargetPathKey = attribute.Key("file.symbolic_link.target_path")
)

func FileAccessed(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileAttributes(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileChanged(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileCreated(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileDirectory(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileExtension(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileForkName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileGroupID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileGroupName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileInode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileMode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileModified(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileOwnerID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FilePath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FileSymbolicLinkTargetPath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	GCPAppHubApplicationContainerKey = attribute.Key("gcp.apphub.application.container")

	GCPAppHubApplicationIDKey = attribute.Key("gcp.apphub.application.id")

	GCPAppHubApplicationLocationKey = attribute.Key("gcp.apphub.application.location")

	GCPAppHubServiceCriticalityTypeKey = attribute.Key("gcp.apphub.service.criticality_type")

	GCPAppHubServiceEnvironmentTypeKey = attribute.Key("gcp.apphub.service.environment_type")

	GCPAppHubServiceIDKey = attribute.Key("gcp.apphub.service.id")

	GCPAppHubWorkloadCriticalityTypeKey = attribute.Key("gcp.apphub.workload.criticality_type")

	GCPAppHubWorkloadEnvironmentTypeKey = attribute.Key("gcp.apphub.workload.environment_type")

	GCPAppHubWorkloadIDKey = attribute.Key("gcp.apphub.workload.id")

	GCPClientServiceKey = attribute.Key("gcp.client.service")

	GCPCloudRunJobExecutionKey = attribute.Key("gcp.cloud_run.job.execution")

	GCPCloudRunJobTaskIndexKey = attribute.Key("gcp.cloud_run.job.task_index")

	GCPGCEInstanceHostnameKey = attribute.Key("gcp.gce.instance.hostname")

	GCPGCEInstanceNameKey = attribute.Key("gcp.gce.instance.name")
)

func GCPAppHubApplicationContainer(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPAppHubApplicationID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPAppHubApplicationLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPAppHubServiceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPAppHubWorkloadID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPClientService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPCloudRunJobExecution(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPCloudRunJobTaskIndex(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPGCEInstanceHostname(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GCPGCEInstanceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	GCPAppHubServiceCriticalityTypeMissionCritical = GCPAppHubServiceCriticalityTypeKey.String("MISSION_CRITICAL")

	GCPAppHubServiceCriticalityTypeHigh = GCPAppHubServiceCriticalityTypeKey.String("HIGH")

	GCPAppHubServiceCriticalityTypeMedium = GCPAppHubServiceCriticalityTypeKey.String("MEDIUM")

	GCPAppHubServiceCriticalityTypeLow = GCPAppHubServiceCriticalityTypeKey.String("LOW")
)

var (
	GCPAppHubServiceEnvironmentTypeProduction = GCPAppHubServiceEnvironmentTypeKey.String("PRODUCTION")

	GCPAppHubServiceEnvironmentTypeStaging = GCPAppHubServiceEnvironmentTypeKey.String("STAGING")

	GCPAppHubServiceEnvironmentTypeTest = GCPAppHubServiceEnvironmentTypeKey.String("TEST")

	GCPAppHubServiceEnvironmentTypeDevelopment = GCPAppHubServiceEnvironmentTypeKey.String("DEVELOPMENT")
)

var (
	GCPAppHubWorkloadCriticalityTypeMissionCritical = GCPAppHubWorkloadCriticalityTypeKey.String("MISSION_CRITICAL")

	GCPAppHubWorkloadCriticalityTypeHigh = GCPAppHubWorkloadCriticalityTypeKey.String("HIGH")

	GCPAppHubWorkloadCriticalityTypeMedium = GCPAppHubWorkloadCriticalityTypeKey.String("MEDIUM")

	GCPAppHubWorkloadCriticalityTypeLow = GCPAppHubWorkloadCriticalityTypeKey.String("LOW")
)

var (
	GCPAppHubWorkloadEnvironmentTypeProduction = GCPAppHubWorkloadEnvironmentTypeKey.String("PRODUCTION")

	GCPAppHubWorkloadEnvironmentTypeStaging = GCPAppHubWorkloadEnvironmentTypeKey.String("STAGING")

	GCPAppHubWorkloadEnvironmentTypeTest = GCPAppHubWorkloadEnvironmentTypeKey.String("TEST")

	GCPAppHubWorkloadEnvironmentTypeDevelopment = GCPAppHubWorkloadEnvironmentTypeKey.String("DEVELOPMENT")
)

const (
	GenAIAgentDescriptionKey = attribute.Key("gen_ai.agent.description")

	GenAIAgentIDKey = attribute.Key("gen_ai.agent.id")

	GenAIAgentNameKey = attribute.Key("gen_ai.agent.name")

	GenAIConversationIDKey = attribute.Key("gen_ai.conversation.id")

	GenAIDataSourceIDKey = attribute.Key("gen_ai.data_source.id")

	GenAIInputMessagesKey = attribute.Key("gen_ai.input.messages")

	GenAIOperationNameKey = attribute.Key("gen_ai.operation.name")

	GenAIOutputMessagesKey = attribute.Key("gen_ai.output.messages")

	GenAIOutputTypeKey = attribute.Key("gen_ai.output.type")

	GenAIProviderNameKey = attribute.Key("gen_ai.provider.name")

	GenAIRequestChoiceCountKey = attribute.Key("gen_ai.request.choice.count")

	GenAIRequestEncodingFormatsKey = attribute.Key("gen_ai.request.encoding_formats")

	GenAIRequestFrequencyPenaltyKey = attribute.Key("gen_ai.request.frequency_penalty")

	GenAIRequestMaxTokensKey = attribute.Key("gen_ai.request.max_tokens")

	GenAIRequestModelKey = attribute.Key("gen_ai.request.model")

	GenAIRequestPresencePenaltyKey = attribute.Key("gen_ai.request.presence_penalty")

	GenAIRequestSeedKey = attribute.Key("gen_ai.request.seed")

	GenAIRequestStopSequencesKey = attribute.Key("gen_ai.request.stop_sequences")

	GenAIRequestTemperatureKey = attribute.Key("gen_ai.request.temperature")

	GenAIRequestTopKKey = attribute.Key("gen_ai.request.top_k")

	GenAIRequestTopPKey = attribute.Key("gen_ai.request.top_p")

	GenAIResponseFinishReasonsKey = attribute.Key("gen_ai.response.finish_reasons")

	GenAIResponseIDKey = attribute.Key("gen_ai.response.id")

	GenAIResponseModelKey = attribute.Key("gen_ai.response.model")

	GenAISystemInstructionsKey = attribute.Key("gen_ai.system_instructions")

	GenAITokenTypeKey = attribute.Key("gen_ai.token.type")

	GenAIToolCallIDKey = attribute.Key("gen_ai.tool.call.id")

	GenAIToolDescriptionKey = attribute.Key("gen_ai.tool.description")

	GenAIToolNameKey = attribute.Key("gen_ai.tool.name")

	GenAIToolTypeKey = attribute.Key("gen_ai.tool.type")

	GenAIUsageInputTokensKey = attribute.Key("gen_ai.usage.input_tokens")

	GenAIUsageOutputTokensKey = attribute.Key("gen_ai.usage.output_tokens")
)

func GenAIAgentDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIAgentID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIAgentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIConversationID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIDataSourceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestChoiceCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestEncodingFormats(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestFrequencyPenalty(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestMaxTokens(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestPresencePenalty(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestSeed(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestStopSequences(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestTemperature(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestTopK(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIRequestTopP(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIResponseFinishReasons(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIResponseID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIResponseModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIToolCallID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIToolDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIToolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIToolType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIUsageInputTokens(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GenAIUsageOutputTokens(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	GenAIOperationNameChat = GenAIOperationNameKey.String("chat")

	GenAIOperationNameGenerateContent = GenAIOperationNameKey.String("generate_content")

	GenAIOperationNameTextCompletion = GenAIOperationNameKey.String("text_completion")

	GenAIOperationNameEmbeddings = GenAIOperationNameKey.String("embeddings")

	GenAIOperationNameCreateAgent = GenAIOperationNameKey.String("create_agent")

	GenAIOperationNameInvokeAgent = GenAIOperationNameKey.String("invoke_agent")

	GenAIOperationNameExecuteTool = GenAIOperationNameKey.String("execute_tool")
)

var (
	GenAIOutputTypeText = GenAIOutputTypeKey.String("text")

	GenAIOutputTypeJSON = GenAIOutputTypeKey.String("json")

	GenAIOutputTypeImage = GenAIOutputTypeKey.String("image")

	GenAIOutputTypeSpeech = GenAIOutputTypeKey.String("speech")
)

var (
	GenAIProviderNameOpenAI = GenAIProviderNameKey.String("openai")

	GenAIProviderNameGCPGenAI = GenAIProviderNameKey.String("gcp.gen_ai")

	GenAIProviderNameGCPVertexAI = GenAIProviderNameKey.String("gcp.vertex_ai")

	GenAIProviderNameGCPGemini = GenAIProviderNameKey.String("gcp.gemini")

	GenAIProviderNameAnthropic = GenAIProviderNameKey.String("anthropic")

	GenAIProviderNameCohere = GenAIProviderNameKey.String("cohere")

	GenAIProviderNameAzureAIInference = GenAIProviderNameKey.String("azure.ai.inference")

	GenAIProviderNameAzureAIOpenAI = GenAIProviderNameKey.String("azure.ai.openai")

	GenAIProviderNameIBMWatsonxAI = GenAIProviderNameKey.String("ibm.watsonx.ai")

	GenAIProviderNameAWSBedrock = GenAIProviderNameKey.String("aws.bedrock")

	GenAIProviderNamePerplexity = GenAIProviderNameKey.String("perplexity")

	GenAIProviderNameXAI = GenAIProviderNameKey.String("x_ai")

	GenAIProviderNameDeepseek = GenAIProviderNameKey.String("deepseek")

	GenAIProviderNameGroq = GenAIProviderNameKey.String("groq")

	GenAIProviderNameMistralAI = GenAIProviderNameKey.String("mistral_ai")
)

var (
	GenAITokenTypeInput = GenAITokenTypeKey.String("input")

	GenAITokenTypeOutput = GenAITokenTypeKey.String("output")
)

const (
	GeoContinentCodeKey = attribute.Key("geo.continent.code")

	GeoCountryISOCodeKey = attribute.Key("geo.country.iso_code")

	GeoLocalityNameKey = attribute.Key("geo.locality.name")

	GeoLocationLatKey = attribute.Key("geo.location.lat")

	GeoLocationLonKey = attribute.Key("geo.location.lon")

	GeoPostalCodeKey = attribute.Key("geo.postal_code")

	GeoRegionISOCodeKey = attribute.Key("geo.region.iso_code")
)

func GeoCountryISOCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GeoLocalityName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GeoLocationLat(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GeoLocationLon(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GeoPostalCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GeoRegionISOCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	GeoContinentCodeAf = GeoContinentCodeKey.String("AF")

	GeoContinentCodeAn = GeoContinentCodeKey.String("AN")

	GeoContinentCodeAs = GeoContinentCodeKey.String("AS")

	GeoContinentCodeEu = GeoContinentCodeKey.String("EU")

	GeoContinentCodeNa = GeoContinentCodeKey.String("NA")

	GeoContinentCodeOc = GeoContinentCodeKey.String("OC")

	GeoContinentCodeSa = GeoContinentCodeKey.String("SA")
)

const (
	GoMemoryTypeKey = attribute.Key("go.memory.type")
)

var (
	GoMemoryTypeStack = GoMemoryTypeKey.String("stack")

	GoMemoryTypeOther = GoMemoryTypeKey.String("other")
)

const (
	GraphQLDocumentKey = attribute.Key("graphql.document")

	GraphQLOperationNameKey = attribute.Key("graphql.operation.name")

	GraphQLOperationTypeKey = attribute.Key("graphql.operation.type")
)

func GraphQLDocument(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GraphQLOperationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	GraphQLOperationTypeQuery = GraphQLOperationTypeKey.String("query")

	GraphQLOperationTypeMutation = GraphQLOperationTypeKey.String("mutation")

	GraphQLOperationTypeSubscription = GraphQLOperationTypeKey.String("subscription")
)

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
	HostArchKey = attribute.Key("host.arch")

	HostCPUCacheL2SizeKey = attribute.Key("host.cpu.cache.l2.size")

	HostCPUFamilyKey = attribute.Key("host.cpu.family")

	HostCPUModelIDKey = attribute.Key("host.cpu.model.id")

	HostCPUModelNameKey = attribute.Key("host.cpu.model.name")

	HostCPUSteppingKey = attribute.Key("host.cpu.stepping")

	HostCPUVendorIDKey = attribute.Key("host.cpu.vendor.id")

	HostIDKey = attribute.Key("host.id")

	HostImageIDKey = attribute.Key("host.image.id")

	HostImageNameKey = attribute.Key("host.image.name")

	HostImageVersionKey = attribute.Key("host.image.version")

	HostIPKey = attribute.Key("host.ip")

	HostMacKey = attribute.Key("host.mac")

	HostNameKey = attribute.Key("host.name")

	HostTypeKey = attribute.Key("host.type")
)

func HostCPUCacheL2Size(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUFamily(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUModelID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUModelName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUStepping(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostCPUVendorID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostImageID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostImageName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostImageVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostIP(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostMac(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HostType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	HostArchAMD64 = HostArchKey.String("amd64")

	HostArchARM32 = HostArchKey.String("arm32")

	HostArchARM64 = HostArchKey.String("arm64")

	HostArchIA64 = HostArchKey.String("ia64")

	HostArchPPC32 = HostArchKey.String("ppc32")

	HostArchPPC64 = HostArchKey.String("ppc64")

	HostArchS390x = HostArchKey.String("s390x")

	HostArchX86 = HostArchKey.String("x86")
)

const (
	HTTPConnectionStateKey = attribute.Key("http.connection.state")

	HTTPRequestBodySizeKey = attribute.Key("http.request.body.size")

	HTTPRequestMethodKey = attribute.Key("http.request.method")

	HTTPRequestMethodOriginalKey = attribute.Key("http.request.method_original")

	HTTPRequestResendCountKey = attribute.Key("http.request.resend_count")

	HTTPRequestSizeKey = attribute.Key("http.request.size")

	HTTPResponseBodySizeKey = attribute.Key("http.response.body.size")

	HTTPResponseSizeKey = attribute.Key("http.response.size")

	HTTPResponseStatusCodeKey = attribute.Key("http.response.status_code")

	HTTPRouteKey = attribute.Key("http.route")
)

func HTTPRequestBodySize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPRequestHeader(key string, val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPRequestMethodOriginal(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPRequestResendCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPRequestSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPResponseBodySize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPResponseHeader(key string, val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPResponseSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPResponseStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPRoute(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	HTTPConnectionStateActive = HTTPConnectionStateKey.String("active")

	HTTPConnectionStateIdle = HTTPConnectionStateKey.String("idle")
)

var (
	HTTPRequestMethodConnect = HTTPRequestMethodKey.String("CONNECT")

	HTTPRequestMethodDelete = HTTPRequestMethodKey.String("DELETE")

	HTTPRequestMethodGet = HTTPRequestMethodKey.String("GET")

	HTTPRequestMethodHead = HTTPRequestMethodKey.String("HEAD")

	HTTPRequestMethodOptions = HTTPRequestMethodKey.String("OPTIONS")

	HTTPRequestMethodPatch = HTTPRequestMethodKey.String("PATCH")

	HTTPRequestMethodPost = HTTPRequestMethodKey.String("POST")

	HTTPRequestMethodPut = HTTPRequestMethodKey.String("PUT")

	HTTPRequestMethodTrace = HTTPRequestMethodKey.String("TRACE")

	HTTPRequestMethodOther = HTTPRequestMethodKey.String("_OTHER")
)

const (
	HwBatteryCapacityKey = attribute.Key("hw.battery.capacity")

	HwBatteryChemistryKey = attribute.Key("hw.battery.chemistry")

	HwBatteryStateKey = attribute.Key("hw.battery.state")

	HwBiosVersionKey = attribute.Key("hw.bios_version")

	HwDriverVersionKey = attribute.Key("hw.driver_version")

	HwEnclosureTypeKey = attribute.Key("hw.enclosure.type")

	HwFirmwareVersionKey = attribute.Key("hw.firmware_version")

	HwGpuTaskKey = attribute.Key("hw.gpu.task")

	HwIDKey = attribute.Key("hw.id")

	HwLimitTypeKey = attribute.Key("hw.limit_type")

	HwLogicalDiskRaidLevelKey = attribute.Key("hw.logical_disk.raid_level")

	HwLogicalDiskStateKey = attribute.Key("hw.logical_disk.state")

	HwMemoryTypeKey = attribute.Key("hw.memory.type")

	HwModelKey = attribute.Key("hw.model")

	HwNameKey = attribute.Key("hw.name")

	HwNetworkLogicalAddressesKey = attribute.Key("hw.network.logical_addresses")

	HwNetworkPhysicalAddressKey = attribute.Key("hw.network.physical_address")

	HwParentKey = attribute.Key("hw.parent")

	HwPhysicalDiskSmartAttributeKey = attribute.Key("hw.physical_disk.smart_attribute")

	HwPhysicalDiskStateKey = attribute.Key("hw.physical_disk.state")

	HwPhysicalDiskTypeKey = attribute.Key("hw.physical_disk.type")

	HwSensorLocationKey = attribute.Key("hw.sensor_location")

	HwSerialNumberKey = attribute.Key("hw.serial_number")

	HwStateKey = attribute.Key("hw.state")

	HwTapeDriveOperationTypeKey = attribute.Key("hw.tape_drive.operation_type")

	HwTypeKey = attribute.Key("hw.type")

	HwVendorKey = attribute.Key("hw.vendor")
)

func HwBatteryCapacity(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwBatteryChemistry(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwBiosVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwDriverVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwEnclosureType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwFirmwareVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwLogicalDiskRaidLevel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwMemoryType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwModel(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwNetworkLogicalAddresses(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwNetworkPhysicalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwParent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwPhysicalDiskSmartAttribute(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwPhysicalDiskType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwSensorLocation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwSerialNumber(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HwVendor(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	HwBatteryStateCharging = HwBatteryStateKey.String("charging")

	HwBatteryStateDischarging = HwBatteryStateKey.String("discharging")
)

var (
	HwGpuTaskDecoder = HwGpuTaskKey.String("decoder")

	HwGpuTaskEncoder = HwGpuTaskKey.String("encoder")

	HwGpuTaskGeneral = HwGpuTaskKey.String("general")
)

var (
	HwLimitTypeCritical = HwLimitTypeKey.String("critical")

	HwLimitTypeDegraded = HwLimitTypeKey.String("degraded")

	HwLimitTypeHighCritical = HwLimitTypeKey.String("high.critical")

	HwLimitTypeHighDegraded = HwLimitTypeKey.String("high.degraded")

	HwLimitTypeLowCritical = HwLimitTypeKey.String("low.critical")

	HwLimitTypeLowDegraded = HwLimitTypeKey.String("low.degraded")

	HwLimitTypeMax = HwLimitTypeKey.String("max")

	HwLimitTypeThrottled = HwLimitTypeKey.String("throttled")

	HwLimitTypeTurbo = HwLimitTypeKey.String("turbo")
)

var (
	HwLogicalDiskStateUsed = HwLogicalDiskStateKey.String("used")

	HwLogicalDiskStateFree = HwLogicalDiskStateKey.String("free")
)

var (
	HwPhysicalDiskStateRemaining = HwPhysicalDiskStateKey.String("remaining")
)

var (
	HwStateDegraded = HwStateKey.String("degraded")

	HwStateFailed = HwStateKey.String("failed")

	HwStateNeedsCleaning = HwStateKey.String("needs_cleaning")

	HwStateOk = HwStateKey.String("ok")

	HwStatePredictedFailure = HwStateKey.String("predicted_failure")
)

var (
	HwTapeDriveOperationTypeMount = HwTapeDriveOperationTypeKey.String("mount")

	HwTapeDriveOperationTypeUnmount = HwTapeDriveOperationTypeKey.String("unmount")

	HwTapeDriveOperationTypeClean = HwTapeDriveOperationTypeKey.String("clean")
)

var (
	HwTypeBattery = HwTypeKey.String("battery")

	HwTypeCPU = HwTypeKey.String("cpu")

	HwTypeDiskController = HwTypeKey.String("disk_controller")

	HwTypeEnclosure = HwTypeKey.String("enclosure")

	HwTypeFan = HwTypeKey.String("fan")

	HwTypeGpu = HwTypeKey.String("gpu")

	HwTypeLogicalDisk = HwTypeKey.String("logical_disk")

	HwTypeMemory = HwTypeKey.String("memory")

	HwTypeNetwork = HwTypeKey.String("network")

	HwTypePhysicalDisk = HwTypeKey.String("physical_disk")

	HwTypePowerSupply = HwTypeKey.String("power_supply")

	HwTypeTapeDrive = HwTypeKey.String("tape_drive")

	HwTypeTemperature = HwTypeKey.String("temperature")

	HwTypeVoltage = HwTypeKey.String("voltage")
)

const (
	IOSAppStateKey = attribute.Key("ios.app.state")
)

var (
	IOSAppStateActive = IOSAppStateKey.String("active")

	IOSAppStateInactive = IOSAppStateKey.String("inactive")

	IOSAppStateBackground = IOSAppStateKey.String("background")

	IOSAppStateForeground = IOSAppStateKey.String("foreground")

	IOSAppStateTerminate = IOSAppStateKey.String("terminate")
)

const (
	K8SClusterNameKey = attribute.Key("k8s.cluster.name")

	K8SClusterUIDKey = attribute.Key("k8s.cluster.uid")

	K8SContainerNameKey = attribute.Key("k8s.container.name")

	K8SContainerRestartCountKey = attribute.Key("k8s.container.restart_count")

	K8SContainerStatusLastTerminatedReasonKey = attribute.Key("k8s.container.status.last_terminated_reason")

	K8SContainerStatusReasonKey = attribute.Key("k8s.container.status.reason")

	K8SContainerStatusStateKey = attribute.Key("k8s.container.status.state")

	K8SCronJobNameKey = attribute.Key("k8s.cronjob.name")

	K8SCronJobUIDKey = attribute.Key("k8s.cronjob.uid")

	K8SDaemonSetNameKey = attribute.Key("k8s.daemonset.name")

	K8SDaemonSetUIDKey = attribute.Key("k8s.daemonset.uid")

	K8SDeploymentNameKey = attribute.Key("k8s.deployment.name")

	K8SDeploymentUIDKey = attribute.Key("k8s.deployment.uid")

	K8SHPAMetricTypeKey = attribute.Key("k8s.hpa.metric.type")

	K8SHPANameKey = attribute.Key("k8s.hpa.name")

	K8SHPAScaletargetrefAPIVersionKey = attribute.Key("k8s.hpa.scaletargetref.api_version")

	K8SHPAScaletargetrefKindKey = attribute.Key("k8s.hpa.scaletargetref.kind")

	K8SHPAScaletargetrefNameKey = attribute.Key("k8s.hpa.scaletargetref.name")

	K8SHPAUIDKey = attribute.Key("k8s.hpa.uid")

	K8SHugepageSizeKey = attribute.Key("k8s.hugepage.size")

	K8SJobNameKey = attribute.Key("k8s.job.name")

	K8SJobUIDKey = attribute.Key("k8s.job.uid")

	K8SNamespaceNameKey = attribute.Key("k8s.namespace.name")

	K8SNamespacePhaseKey = attribute.Key("k8s.namespace.phase")

	K8SNodeConditionStatusKey = attribute.Key("k8s.node.condition.status")

	K8SNodeConditionTypeKey = attribute.Key("k8s.node.condition.type")

	K8SNodeNameKey = attribute.Key("k8s.node.name")

	K8SNodeUIDKey = attribute.Key("k8s.node.uid")

	K8SPodNameKey = attribute.Key("k8s.pod.name")

	K8SPodUIDKey = attribute.Key("k8s.pod.uid")

	K8SReplicaSetNameKey = attribute.Key("k8s.replicaset.name")

	K8SReplicaSetUIDKey = attribute.Key("k8s.replicaset.uid")

	K8SReplicationControllerNameKey = attribute.Key("k8s.replicationcontroller.name")

	K8SReplicationControllerUIDKey = attribute.Key("k8s.replicationcontroller.uid")

	K8SResourceQuotaNameKey = attribute.Key("k8s.resourcequota.name")

	K8SResourceQuotaResourceNameKey = attribute.Key("k8s.resourcequota.resource_name")

	K8SResourceQuotaUIDKey = attribute.Key("k8s.resourcequota.uid")

	K8SStatefulSetNameKey = attribute.Key("k8s.statefulset.name")

	K8SStatefulSetUIDKey = attribute.Key("k8s.statefulset.uid")

	K8SStorageclassNameKey = attribute.Key("k8s.storageclass.name")

	K8SVolumeNameKey = attribute.Key("k8s.volume.name")

	K8SVolumeTypeKey = attribute.Key("k8s.volume.type")
)

func K8SClusterName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SClusterUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SContainerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SContainerRestartCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SContainerStatusLastTerminatedReason(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SCronJobAnnotation(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SCronJobLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SCronJobName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SCronJobUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDaemonSetAnnotation(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDaemonSetLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDaemonSetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDaemonSetUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDeploymentAnnotation(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDeploymentLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDeploymentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SDeploymentUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SHPAMetricType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SHPAName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SHPAScaletargetrefAPIVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SHPAScaletargetrefKind(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SHPAScaletargetrefName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SHPAUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SHugepageSize(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SJobAnnotation(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SJobLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SJobName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SJobUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SNamespaceAnnotation(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SNamespaceLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SNamespaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SNodeAnnotation(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SNodeLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SNodeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SNodeUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SPodAnnotation(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SPodLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SPodName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SPodUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SReplicaSetAnnotation(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SReplicaSetLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SReplicaSetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SReplicaSetUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SReplicationControllerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SReplicationControllerUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SResourceQuotaName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SResourceQuotaResourceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SResourceQuotaUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SStatefulSetAnnotation(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SStatefulSetLabel(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SStatefulSetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SStatefulSetUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SStorageclassName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func K8SVolumeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	K8SContainerStatusReasonContainerCreating = K8SContainerStatusReasonKey.String("ContainerCreating")

	K8SContainerStatusReasonCrashLoopBackOff = K8SContainerStatusReasonKey.String("CrashLoopBackOff")

	K8SContainerStatusReasonCreateContainerConfigError = K8SContainerStatusReasonKey.String("CreateContainerConfigError")

	K8SContainerStatusReasonErrImagePull = K8SContainerStatusReasonKey.String("ErrImagePull")

	K8SContainerStatusReasonImagePullBackOff = K8SContainerStatusReasonKey.String("ImagePullBackOff")

	K8SContainerStatusReasonOomKilled = K8SContainerStatusReasonKey.String("OOMKilled")

	K8SContainerStatusReasonCompleted = K8SContainerStatusReasonKey.String("Completed")

	K8SContainerStatusReasonError = K8SContainerStatusReasonKey.String("Error")

	K8SContainerStatusReasonContainerCannotRun = K8SContainerStatusReasonKey.String("ContainerCannotRun")
)

var (
	K8SContainerStatusStateTerminated = K8SContainerStatusStateKey.String("terminated")

	K8SContainerStatusStateRunning = K8SContainerStatusStateKey.String("running")

	K8SContainerStatusStateWaiting = K8SContainerStatusStateKey.String("waiting")
)

var (
	K8SNamespacePhaseActive = K8SNamespacePhaseKey.String("active")

	K8SNamespacePhaseTerminating = K8SNamespacePhaseKey.String("terminating")
)

var (
	K8SNodeConditionStatusConditionTrue = K8SNodeConditionStatusKey.String("true")

	K8SNodeConditionStatusConditionFalse = K8SNodeConditionStatusKey.String("false")

	K8SNodeConditionStatusConditionUnknown = K8SNodeConditionStatusKey.String("unknown")
)

var (
	K8SNodeConditionTypeReady = K8SNodeConditionTypeKey.String("Ready")

	K8SNodeConditionTypeDiskPressure = K8SNodeConditionTypeKey.String("DiskPressure")

	K8SNodeConditionTypeMemoryPressure = K8SNodeConditionTypeKey.String("MemoryPressure")

	K8SNodeConditionTypePIDPressure = K8SNodeConditionTypeKey.String("PIDPressure")

	K8SNodeConditionTypeNetworkUnavailable = K8SNodeConditionTypeKey.String("NetworkUnavailable")
)

var (
	K8SVolumeTypePersistentVolumeClaim = K8SVolumeTypeKey.String("persistentVolumeClaim")

	K8SVolumeTypeConfigMap = K8SVolumeTypeKey.String("configMap")

	K8SVolumeTypeDownwardAPI = K8SVolumeTypeKey.String("downwardAPI")

	K8SVolumeTypeEmptyDir = K8SVolumeTypeKey.String("emptyDir")

	K8SVolumeTypeSecret = K8SVolumeTypeKey.String("secret")

	K8SVolumeTypeLocal = K8SVolumeTypeKey.String("local")
)

const (
	LinuxMemorySlabStateKey = attribute.Key("linux.memory.slab.state")
)

var (
	LinuxMemorySlabStateReclaimable = LinuxMemorySlabStateKey.String("reclaimable")

	LinuxMemorySlabStateUnreclaimable = LinuxMemorySlabStateKey.String("unreclaimable")
)

const (
	LogFileNameKey = attribute.Key("log.file.name")

	LogFileNameResolvedKey = attribute.Key("log.file.name_resolved")

	LogFilePathKey = attribute.Key("log.file.path")

	LogFilePathResolvedKey = attribute.Key("log.file.path_resolved")

	LogIostreamKey = attribute.Key("log.iostream")

	LogRecordOriginalKey = attribute.Key("log.record.original")

	LogRecordUIDKey = attribute.Key("log.record.uid")
)

func LogFileName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func LogFileNameResolved(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func LogFilePath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func LogFilePathResolved(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func LogRecordOriginal(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func LogRecordUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	LogIostreamStdout = LogIostreamKey.String("stdout")

	LogIostreamStderr = LogIostreamKey.String("stderr")
)

const (
	MainframeLparNameKey = attribute.Key("mainframe.lpar.name")
)

func MainframeLparName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingBatchMessageCountKey = attribute.Key("messaging.batch.message_count")

	MessagingClientIDKey = attribute.Key("messaging.client.id")

	MessagingConsumerGroupNameKey = attribute.Key("messaging.consumer.group.name")

	MessagingDestinationAnonymousKey = attribute.Key("messaging.destination.anonymous")

	MessagingDestinationNameKey = attribute.Key("messaging.destination.name")

	MessagingDestinationPartitionIDKey = attribute.Key("messaging.destination.partition.id")

	MessagingDestinationSubscriptionNameKey = attribute.Key("messaging.destination.subscription.name")

	MessagingDestinationTemplateKey = attribute.Key("messaging.destination.template")

	MessagingDestinationTemporaryKey = attribute.Key("messaging.destination.temporary")

	MessagingEventHubsMessageEnqueuedTimeKey = attribute.Key("messaging.eventhubs.message.enqueued_time")

	MessagingGCPPubSubMessageAckDeadlineKey = attribute.Key("messaging.gcp_pubsub.message.ack_deadline")

	MessagingGCPPubSubMessageAckIDKey = attribute.Key("messaging.gcp_pubsub.message.ack_id")

	MessagingGCPPubSubMessageDeliveryAttemptKey = attribute.Key("messaging.gcp_pubsub.message.delivery_attempt")

	MessagingGCPPubSubMessageOrderingKeyKey = attribute.Key("messaging.gcp_pubsub.message.ordering_key")

	MessagingKafkaMessageKeyKey = attribute.Key("messaging.kafka.message.key")

	MessagingKafkaMessageTombstoneKey = attribute.Key("messaging.kafka.message.tombstone")

	MessagingKafkaOffsetKey = attribute.Key("messaging.kafka.offset")

	MessagingMessageBodySizeKey = attribute.Key("messaging.message.body.size")

	MessagingMessageConversationIDKey = attribute.Key("messaging.message.conversation_id")

	MessagingMessageEnvelopeSizeKey = attribute.Key("messaging.message.envelope.size")

	MessagingMessageIDKey = attribute.Key("messaging.message.id")

	MessagingOperationNameKey = attribute.Key("messaging.operation.name")

	MessagingOperationTypeKey = attribute.Key("messaging.operation.type")

	MessagingRabbitMQDestinationRoutingKeyKey = attribute.Key("messaging.rabbitmq.destination.routing_key")

	MessagingRabbitMQMessageDeliveryTagKey = attribute.Key("messaging.rabbitmq.message.delivery_tag")

	MessagingRocketMQConsumptionModelKey = attribute.Key("messaging.rocketmq.consumption_model")

	MessagingRocketMQMessageDelayTimeLevelKey = attribute.Key("messaging.rocketmq.message.delay_time_level")

	MessagingRocketMQMessageDeliveryTimestampKey = attribute.Key("messaging.rocketmq.message.delivery_timestamp")

	MessagingRocketMQMessageGroupKey = attribute.Key("messaging.rocketmq.message.group")

	MessagingRocketMQMessageKeysKey = attribute.Key("messaging.rocketmq.message.keys")

	MessagingRocketMQMessageTagKey = attribute.Key("messaging.rocketmq.message.tag")

	MessagingRocketMQMessageTypeKey = attribute.Key("messaging.rocketmq.message.type")

	MessagingRocketMQNamespaceKey = attribute.Key("messaging.rocketmq.namespace")

	MessagingServiceBusDispositionStatusKey = attribute.Key("messaging.servicebus.disposition_status")

	MessagingServiceBusMessageDeliveryCountKey = attribute.Key("messaging.servicebus.message.delivery_count")

	MessagingServiceBusMessageEnqueuedTimeKey = attribute.Key("messaging.servicebus.message.enqueued_time")

	MessagingSystemKey = attribute.Key("messaging.system")
)

func MessagingBatchMessageCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingClientID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingConsumerGroupName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingDestinationAnonymous(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingDestinationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingDestinationPartitionID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingDestinationSubscriptionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingDestinationTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingDestinationTemporary(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingEventHubsMessageEnqueuedTime(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingGCPPubSubMessageAckDeadline(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingGCPPubSubMessageAckID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingGCPPubSubMessageDeliveryAttempt(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingGCPPubSubMessageOrderingKey(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaMessageKey(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaMessageTombstone(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaOffset(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingMessageBodySize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingMessageConversationID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingMessageEnvelopeSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingMessageID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingOperationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRabbitMQDestinationRoutingKey(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRabbitMQMessageDeliveryTag(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketMQMessageDelayTimeLevel(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketMQMessageDeliveryTimestamp(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketMQMessageGroup(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketMQMessageKeys(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketMQMessageTag(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketMQNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingServiceBusMessageDeliveryCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingServiceBusMessageEnqueuedTime(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	MessagingOperationTypeCreate = MessagingOperationTypeKey.String("create")

	MessagingOperationTypeSend = MessagingOperationTypeKey.String("send")

	MessagingOperationTypeReceive = MessagingOperationTypeKey.String("receive")

	MessagingOperationTypeProcess = MessagingOperationTypeKey.String("process")

	MessagingOperationTypeSettle = MessagingOperationTypeKey.String("settle")
)

var (
	MessagingRocketMQConsumptionModelClustering = MessagingRocketMQConsumptionModelKey.String("clustering")

	MessagingRocketMQConsumptionModelBroadcasting = MessagingRocketMQConsumptionModelKey.String("broadcasting")
)

var (
	MessagingRocketMQMessageTypeNormal = MessagingRocketMQMessageTypeKey.String("normal")

	MessagingRocketMQMessageTypeFifo = MessagingRocketMQMessageTypeKey.String("fifo")

	MessagingRocketMQMessageTypeDelay = MessagingRocketMQMessageTypeKey.String("delay")

	MessagingRocketMQMessageTypeTransaction = MessagingRocketMQMessageTypeKey.String("transaction")
)

var (
	MessagingServiceBusDispositionStatusComplete = MessagingServiceBusDispositionStatusKey.String("complete")

	MessagingServiceBusDispositionStatusAbandon = MessagingServiceBusDispositionStatusKey.String("abandon")

	MessagingServiceBusDispositionStatusDeadLetter = MessagingServiceBusDispositionStatusKey.String("dead_letter")

	MessagingServiceBusDispositionStatusDefer = MessagingServiceBusDispositionStatusKey.String("defer")
)

var (
	MessagingSystemActiveMQ = MessagingSystemKey.String("activemq")

	MessagingSystemAWSSNS = MessagingSystemKey.String("aws.sns")

	MessagingSystemAWSSQS = MessagingSystemKey.String("aws_sqs")

	MessagingSystemEventGrid = MessagingSystemKey.String("eventgrid")

	MessagingSystemEventHubs = MessagingSystemKey.String("eventhubs")

	MessagingSystemServiceBus = MessagingSystemKey.String("servicebus")

	MessagingSystemGCPPubSub = MessagingSystemKey.String("gcp_pubsub")

	MessagingSystemJMS = MessagingSystemKey.String("jms")

	MessagingSystemKafka = MessagingSystemKey.String("kafka")

	MessagingSystemRabbitMQ = MessagingSystemKey.String("rabbitmq")

	MessagingSystemRocketMQ = MessagingSystemKey.String("rocketmq")

	MessagingSystemPulsar = MessagingSystemKey.String("pulsar")
)

const (
	NetworkCarrierICCKey = attribute.Key("network.carrier.icc")

	NetworkCarrierMCCKey = attribute.Key("network.carrier.mcc")

	NetworkCarrierMNCKey = attribute.Key("network.carrier.mnc")

	NetworkCarrierNameKey = attribute.Key("network.carrier.name")

	NetworkConnectionStateKey = attribute.Key("network.connection.state")

	NetworkConnectionSubtypeKey = attribute.Key("network.connection.subtype")

	NetworkConnectionTypeKey = attribute.Key("network.connection.type")

	NetworkInterfaceNameKey = attribute.Key("network.interface.name")

	NetworkIODirectionKey = attribute.Key("network.io.direction")

	NetworkLocalAddressKey = attribute.Key("network.local.address")

	NetworkLocalPortKey = attribute.Key("network.local.port")

	NetworkPeerAddressKey = attribute.Key("network.peer.address")

	NetworkPeerPortKey = attribute.Key("network.peer.port")

	NetworkProtocolNameKey = attribute.Key("network.protocol.name")

	NetworkProtocolVersionKey = attribute.Key("network.protocol.version")

	NetworkTransportKey = attribute.Key("network.transport")

	NetworkTypeKey = attribute.Key("network.type")
)

func NetworkCarrierICC(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkCarrierMCC(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkCarrierMNC(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkCarrierName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkInterfaceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkLocalAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkLocalPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkPeerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkPeerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	NetworkConnectionStateClosed = NetworkConnectionStateKey.String("closed")

	NetworkConnectionStateCloseWait = NetworkConnectionStateKey.String("close_wait")

	NetworkConnectionStateClosing = NetworkConnectionStateKey.String("closing")

	NetworkConnectionStateEstablished = NetworkConnectionStateKey.String("established")

	NetworkConnectionStateFinWait1 = NetworkConnectionStateKey.String("fin_wait_1")

	NetworkConnectionStateFinWait2 = NetworkConnectionStateKey.String("fin_wait_2")

	NetworkConnectionStateLastAck = NetworkConnectionStateKey.String("last_ack")

	NetworkConnectionStateListen = NetworkConnectionStateKey.String("listen")

	NetworkConnectionStateSynReceived = NetworkConnectionStateKey.String("syn_received")

	NetworkConnectionStateSynSent = NetworkConnectionStateKey.String("syn_sent")

	NetworkConnectionStateTimeWait = NetworkConnectionStateKey.String("time_wait")
)

var (
	NetworkConnectionSubtypeGprs = NetworkConnectionSubtypeKey.String("gprs")

	NetworkConnectionSubtypeEdge = NetworkConnectionSubtypeKey.String("edge")

	NetworkConnectionSubtypeUmts = NetworkConnectionSubtypeKey.String("umts")

	NetworkConnectionSubtypeCdma = NetworkConnectionSubtypeKey.String("cdma")

	NetworkConnectionSubtypeEvdo0 = NetworkConnectionSubtypeKey.String("evdo_0")

	NetworkConnectionSubtypeEvdoA = NetworkConnectionSubtypeKey.String("evdo_a")

	NetworkConnectionSubtypeCdma20001xrtt = NetworkConnectionSubtypeKey.String("cdma2000_1xrtt")

	NetworkConnectionSubtypeHsdpa = NetworkConnectionSubtypeKey.String("hsdpa")

	NetworkConnectionSubtypeHsupa = NetworkConnectionSubtypeKey.String("hsupa")

	NetworkConnectionSubtypeHspa = NetworkConnectionSubtypeKey.String("hspa")

	NetworkConnectionSubtypeIden = NetworkConnectionSubtypeKey.String("iden")

	NetworkConnectionSubtypeEvdoB = NetworkConnectionSubtypeKey.String("evdo_b")

	NetworkConnectionSubtypeLte = NetworkConnectionSubtypeKey.String("lte")

	NetworkConnectionSubtypeEhrpd = NetworkConnectionSubtypeKey.String("ehrpd")

	NetworkConnectionSubtypeHspap = NetworkConnectionSubtypeKey.String("hspap")

	NetworkConnectionSubtypeGsm = NetworkConnectionSubtypeKey.String("gsm")

	NetworkConnectionSubtypeTdScdma = NetworkConnectionSubtypeKey.String("td_scdma")

	NetworkConnectionSubtypeIwlan = NetworkConnectionSubtypeKey.String("iwlan")

	NetworkConnectionSubtypeNr = NetworkConnectionSubtypeKey.String("nr")

	NetworkConnectionSubtypeNrnsa = NetworkConnectionSubtypeKey.String("nrnsa")

	NetworkConnectionSubtypeLteCa = NetworkConnectionSubtypeKey.String("lte_ca")
)

var (
	NetworkConnectionTypeWifi = NetworkConnectionTypeKey.String("wifi")

	NetworkConnectionTypeWired = NetworkConnectionTypeKey.String("wired")

	NetworkConnectionTypeCell = NetworkConnectionTypeKey.String("cell")

	NetworkConnectionTypeUnavailable = NetworkConnectionTypeKey.String("unavailable")

	NetworkConnectionTypeUnknown = NetworkConnectionTypeKey.String("unknown")
)

var (
	NetworkIODirectionTransmit = NetworkIODirectionKey.String("transmit")

	NetworkIODirectionReceive = NetworkIODirectionKey.String("receive")
)

var (
	NetworkTransportTCP = NetworkTransportKey.String("tcp")

	NetworkTransportUDP = NetworkTransportKey.String("udp")

	NetworkTransportPipe = NetworkTransportKey.String("pipe")

	NetworkTransportUnix = NetworkTransportKey.String("unix")

	NetworkTransportQUIC = NetworkTransportKey.String("quic")
)

var (
	NetworkTypeIPv4 = NetworkTypeKey.String("ipv4")

	NetworkTypeIPv6 = NetworkTypeKey.String("ipv6")
)

const (
	OCIManifestDigestKey = attribute.Key("oci.manifest.digest")
)

func OCIManifestDigest(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	OpenAIRequestServiceTierKey = attribute.Key("openai.request.service_tier")

	OpenAIResponseServiceTierKey = attribute.Key("openai.response.service_tier")

	OpenAIResponseSystemFingerprintKey = attribute.Key("openai.response.system_fingerprint")
)

func OpenAIResponseServiceTier(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OpenAIResponseSystemFingerprint(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	OpenAIRequestServiceTierAuto = OpenAIRequestServiceTierKey.String("auto")

	OpenAIRequestServiceTierDefault = OpenAIRequestServiceTierKey.String("default")
)

const (
	OpenTracingRefTypeKey = attribute.Key("opentracing.ref_type")
)

var (
	OpenTracingRefTypeChildOf = OpenTracingRefTypeKey.String("child_of")

	OpenTracingRefTypeFollowsFrom = OpenTracingRefTypeKey.String("follows_from")
)

const (
	OSBuildIDKey = attribute.Key("os.build_id")

	OSDescriptionKey = attribute.Key("os.description")

	OSNameKey = attribute.Key("os.name")

	OSTypeKey = attribute.Key("os.type")

	OSVersionKey = attribute.Key("os.version")
)

func OSBuildID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OSDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OSName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OSVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	OSTypeWindows = OSTypeKey.String("windows")

	OSTypeLinux = OSTypeKey.String("linux")

	OSTypeDarwin = OSTypeKey.String("darwin")

	OSTypeFreeBSD = OSTypeKey.String("freebsd")

	OSTypeNetBSD = OSTypeKey.String("netbsd")

	OSTypeOpenBSD = OSTypeKey.String("openbsd")

	OSTypeDragonflyBSD = OSTypeKey.String("dragonflybsd")

	OSTypeHPUX = OSTypeKey.String("hpux")

	OSTypeAIX = OSTypeKey.String("aix")

	OSTypeSolaris = OSTypeKey.String("solaris")

	OSTypeZOS = OSTypeKey.String("zos")
)

const (
	OTelComponentNameKey = attribute.Key("otel.component.name")

	OTelComponentTypeKey = attribute.Key("otel.component.type")

	OTelScopeNameKey = attribute.Key("otel.scope.name")

	OTelScopeSchemaURLKey = attribute.Key("otel.scope.schema_url")

	OTelScopeVersionKey = attribute.Key("otel.scope.version")

	OTelSpanParentOriginKey = attribute.Key("otel.span.parent.origin")

	OTelSpanSamplingResultKey = attribute.Key("otel.span.sampling_result")

	OTelStatusCodeKey = attribute.Key("otel.status_code")

	OTelStatusDescriptionKey = attribute.Key("otel.status_description")
)

func OTelComponentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OTelScopeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OTelScopeSchemaURL(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OTelScopeVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func OTelStatusDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	OTelComponentTypeBatchingSpanProcessor = OTelComponentTypeKey.String("batching_span_processor")

	OTelComponentTypeSimpleSpanProcessor = OTelComponentTypeKey.String("simple_span_processor")

	OTelComponentTypeBatchingLogProcessor = OTelComponentTypeKey.String("batching_log_processor")

	OTelComponentTypeSimpleLogProcessor = OTelComponentTypeKey.String("simple_log_processor")

	OTelComponentTypeOtlpGRPCSpanExporter = OTelComponentTypeKey.String("otlp_grpc_span_exporter")

	OTelComponentTypeOtlpHTTPSpanExporter = OTelComponentTypeKey.String("otlp_http_span_exporter")

	OTelComponentTypeOtlpHTTPJSONSpanExporter = OTelComponentTypeKey.String("otlp_http_json_span_exporter")

	OTelComponentTypeZipkinHTTPSpanExporter = OTelComponentTypeKey.String("zipkin_http_span_exporter")

	OTelComponentTypeOtlpGRPCLogExporter = OTelComponentTypeKey.String("otlp_grpc_log_exporter")

	OTelComponentTypeOtlpHTTPLogExporter = OTelComponentTypeKey.String("otlp_http_log_exporter")

	OTelComponentTypeOtlpHTTPJSONLogExporter = OTelComponentTypeKey.String("otlp_http_json_log_exporter")

	OTelComponentTypePeriodicMetricReader = OTelComponentTypeKey.String("periodic_metric_reader")

	OTelComponentTypeOtlpGRPCMetricExporter = OTelComponentTypeKey.String("otlp_grpc_metric_exporter")

	OTelComponentTypeOtlpHTTPMetricExporter = OTelComponentTypeKey.String("otlp_http_metric_exporter")

	OTelComponentTypeOtlpHTTPJSONMetricExporter = OTelComponentTypeKey.String("otlp_http_json_metric_exporter")

	OTelComponentTypePrometheusHTTPTextMetricExporter = OTelComponentTypeKey.String("prometheus_http_text_metric_exporter")
)

var (
	OTelSpanParentOriginNone = OTelSpanParentOriginKey.String("none")

	OTelSpanParentOriginLocal = OTelSpanParentOriginKey.String("local")

	OTelSpanParentOriginRemote = OTelSpanParentOriginKey.String("remote")
)

var (
	OTelSpanSamplingResultDrop = OTelSpanSamplingResultKey.String("DROP")

	OTelSpanSamplingResultRecordOnly = OTelSpanSamplingResultKey.String("RECORD_ONLY")

	OTelSpanSamplingResultRecordAndSample = OTelSpanSamplingResultKey.String("RECORD_AND_SAMPLE")
)

var (
	OTelStatusCodeOk = OTelStatusCodeKey.String("OK")

	OTelStatusCodeError = OTelStatusCodeKey.String("ERROR")
)

const (
	PeerServiceKey = attribute.Key("peer.service")
)

func PeerService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ProcessArgsCountKey = attribute.Key("process.args_count")

	ProcessCommandKey = attribute.Key("process.command")

	ProcessCommandArgsKey = attribute.Key("process.command_args")

	ProcessCommandLineKey = attribute.Key("process.command_line")

	ProcessContextSwitchTypeKey = attribute.Key("process.context_switch_type")

	ProcessCreationTimeKey = attribute.Key("process.creation.time")

	ProcessExecutableBuildIDGNUKey = attribute.Key("process.executable.build_id.gnu")

	ProcessExecutableBuildIDGoKey = attribute.Key("process.executable.build_id.go")

	ProcessExecutableBuildIDHtlhashKey = attribute.Key("process.executable.build_id.htlhash")

	ProcessExecutableNameKey = attribute.Key("process.executable.name")

	ProcessExecutablePathKey = attribute.Key("process.executable.path")

	ProcessExitCodeKey = attribute.Key("process.exit.code")

	ProcessExitTimeKey = attribute.Key("process.exit.time")

	ProcessGroupLeaderPIDKey = attribute.Key("process.group_leader.pid")

	ProcessInteractiveKey = attribute.Key("process.interactive")

	ProcessLinuxCgroupKey = attribute.Key("process.linux.cgroup")

	ProcessOwnerKey = attribute.Key("process.owner")

	ProcessPagingFaultTypeKey = attribute.Key("process.paging.fault_type")

	ProcessParentPIDKey = attribute.Key("process.parent_pid")

	ProcessPIDKey = attribute.Key("process.pid")

	ProcessRealUserIDKey = attribute.Key("process.real_user.id")

	ProcessRealUserNameKey = attribute.Key("process.real_user.name")

	ProcessRuntimeDescriptionKey = attribute.Key("process.runtime.description")

	ProcessRuntimeNameKey = attribute.Key("process.runtime.name")

	ProcessRuntimeVersionKey = attribute.Key("process.runtime.version")

	ProcessSavedUserIDKey = attribute.Key("process.saved_user.id")

	ProcessSavedUserNameKey = attribute.Key("process.saved_user.name")

	ProcessSessionLeaderPIDKey = attribute.Key("process.session_leader.pid")

	ProcessTitleKey = attribute.Key("process.title")

	ProcessUserIDKey = attribute.Key("process.user.id")

	ProcessUserNameKey = attribute.Key("process.user.name")

	ProcessVpidKey = attribute.Key("process.vpid")

	ProcessWorkingDirectoryKey = attribute.Key("process.working_directory")
)

func ProcessArgsCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessCommand(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessCommandArgs(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessCommandLine(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessCreationTime(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessEnvironmentVariable(key string, val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessExecutableBuildIDGNU(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessExecutableBuildIDGo(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessExecutableBuildIDHtlhash(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessExecutableName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessExecutablePath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessExitCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessExitTime(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessGroupLeaderPID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessInteractive(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessLinuxCgroup(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessOwner(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessParentPID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessPID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessRealUserID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessRealUserName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessRuntimeDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessRuntimeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessRuntimeVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessSavedUserID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessSavedUserName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessSessionLeaderPID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessTitle(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessUserID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessUserName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessVpid(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ProcessWorkingDirectory(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	ProcessContextSwitchTypeVoluntary = ProcessContextSwitchTypeKey.String("voluntary")

	ProcessContextSwitchTypeInvoluntary = ProcessContextSwitchTypeKey.String("involuntary")
)

var (
	ProcessPagingFaultTypeMajor = ProcessPagingFaultTypeKey.String("major")

	ProcessPagingFaultTypeMinor = ProcessPagingFaultTypeKey.String("minor")
)

const (
	ProfileFrameTypeKey = attribute.Key("profile.frame.type")
)

var (
	ProfileFrameTypeDotnet = ProfileFrameTypeKey.String("dotnet")

	ProfileFrameTypeJVM = ProfileFrameTypeKey.String("jvm")

	ProfileFrameTypeKernel = ProfileFrameTypeKey.String("kernel")

	ProfileFrameTypeNative = ProfileFrameTypeKey.String("native")

	ProfileFrameTypePerl = ProfileFrameTypeKey.String("perl")

	ProfileFrameTypePHP = ProfileFrameTypeKey.String("php")

	ProfileFrameTypeCpython = ProfileFrameTypeKey.String("cpython")

	ProfileFrameTypeRuby = ProfileFrameTypeKey.String("ruby")

	ProfileFrameTypeV8JS = ProfileFrameTypeKey.String("v8js")

	ProfileFrameTypeBeam = ProfileFrameTypeKey.String("beam")

	ProfileFrameTypeGo = ProfileFrameTypeKey.String("go")

	ProfileFrameTypeRust = ProfileFrameTypeKey.String("rust")
)

const (
	RPCConnectRPCErrorCodeKey = attribute.Key("rpc.connect_rpc.error_code")

	RPCGRPCStatusCodeKey = attribute.Key("rpc.grpc.status_code")

	RPCJSONRPCErrorCodeKey = attribute.Key("rpc.jsonrpc.error_code")

	RPCJSONRPCErrorMessageKey = attribute.Key("rpc.jsonrpc.error_message")

	RPCJSONRPCRequestIDKey = attribute.Key("rpc.jsonrpc.request_id")

	RPCJSONRPCVersionKey = attribute.Key("rpc.jsonrpc.version")

	RPCMessageCompressedSizeKey = attribute.Key("rpc.message.compressed_size")

	RPCMessageIDKey = attribute.Key("rpc.message.id")

	RPCMessageTypeKey = attribute.Key("rpc.message.type")

	RPCMessageUncompressedSizeKey = attribute.Key("rpc.message.uncompressed_size")

	RPCMethodKey = attribute.Key("rpc.method")

	RPCServiceKey = attribute.Key("rpc.service")

	RPCSystemKey = attribute.Key("rpc.system")
)

func RPCConnectRPCRequestMetadata(key string, val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCConnectRPCResponseMetadata(key string, val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCGRPCRequestMetadata(key string, val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCGRPCResponseMetadata(key string, val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJSONRPCErrorCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJSONRPCErrorMessage(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJSONRPCRequestID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJSONRPCVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCMessageCompressedSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCMessageID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCMessageUncompressedSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	RPCConnectRPCErrorCodeCancelled = RPCConnectRPCErrorCodeKey.String("cancelled")

	RPCConnectRPCErrorCodeUnknown = RPCConnectRPCErrorCodeKey.String("unknown")

	RPCConnectRPCErrorCodeInvalidArgument = RPCConnectRPCErrorCodeKey.String("invalid_argument")

	RPCConnectRPCErrorCodeDeadlineExceeded = RPCConnectRPCErrorCodeKey.String("deadline_exceeded")

	RPCConnectRPCErrorCodeNotFound = RPCConnectRPCErrorCodeKey.String("not_found")

	RPCConnectRPCErrorCodeAlreadyExists = RPCConnectRPCErrorCodeKey.String("already_exists")

	RPCConnectRPCErrorCodePermissionDenied = RPCConnectRPCErrorCodeKey.String("permission_denied")

	RPCConnectRPCErrorCodeResourceExhausted = RPCConnectRPCErrorCodeKey.String("resource_exhausted")

	RPCConnectRPCErrorCodeFailedPrecondition = RPCConnectRPCErrorCodeKey.String("failed_precondition")

	RPCConnectRPCErrorCodeAborted = RPCConnectRPCErrorCodeKey.String("aborted")

	RPCConnectRPCErrorCodeOutOfRange = RPCConnectRPCErrorCodeKey.String("out_of_range")

	RPCConnectRPCErrorCodeUnimplemented = RPCConnectRPCErrorCodeKey.String("unimplemented")

	RPCConnectRPCErrorCodeInternal = RPCConnectRPCErrorCodeKey.String("internal")

	RPCConnectRPCErrorCodeUnavailable = RPCConnectRPCErrorCodeKey.String("unavailable")

	RPCConnectRPCErrorCodeDataLoss = RPCConnectRPCErrorCodeKey.String("data_loss")

	RPCConnectRPCErrorCodeUnauthenticated = RPCConnectRPCErrorCodeKey.String("unauthenticated")
)

var (
	RPCGRPCStatusCodeOk = RPCGRPCStatusCodeKey.Int(0)

	RPCGRPCStatusCodeCancelled = RPCGRPCStatusCodeKey.Int(1)

	RPCGRPCStatusCodeUnknown = RPCGRPCStatusCodeKey.Int(2)

	RPCGRPCStatusCodeInvalidArgument = RPCGRPCStatusCodeKey.Int(3)

	RPCGRPCStatusCodeDeadlineExceeded = RPCGRPCStatusCodeKey.Int(4)

	RPCGRPCStatusCodeNotFound = RPCGRPCStatusCodeKey.Int(5)

	RPCGRPCStatusCodeAlreadyExists = RPCGRPCStatusCodeKey.Int(6)

	RPCGRPCStatusCodePermissionDenied = RPCGRPCStatusCodeKey.Int(7)

	RPCGRPCStatusCodeResourceExhausted = RPCGRPCStatusCodeKey.Int(8)

	RPCGRPCStatusCodeFailedPrecondition = RPCGRPCStatusCodeKey.Int(9)

	RPCGRPCStatusCodeAborted = RPCGRPCStatusCodeKey.Int(10)

	RPCGRPCStatusCodeOutOfRange = RPCGRPCStatusCodeKey.Int(11)

	RPCGRPCStatusCodeUnimplemented = RPCGRPCStatusCodeKey.Int(12)

	RPCGRPCStatusCodeInternal = RPCGRPCStatusCodeKey.Int(13)

	RPCGRPCStatusCodeUnavailable = RPCGRPCStatusCodeKey.Int(14)

	RPCGRPCStatusCodeDataLoss = RPCGRPCStatusCodeKey.Int(15)

	RPCGRPCStatusCodeUnauthenticated = RPCGRPCStatusCodeKey.Int(16)
)

var (
	RPCMessageTypeSent = RPCMessageTypeKey.String("SENT")

	RPCMessageTypeReceived = RPCMessageTypeKey.String("RECEIVED")
)

var (
	RPCSystemGRPC = RPCSystemKey.String("grpc")

	RPCSystemJavaRmi = RPCSystemKey.String("java_rmi")

	RPCSystemDotnetWcf = RPCSystemKey.String("dotnet_wcf")

	RPCSystemApacheDubbo = RPCSystemKey.String("apache_dubbo")

	RPCSystemConnectRPC = RPCSystemKey.String("connect_rpc")
)

const (
	SecurityRuleCategoryKey = attribute.Key("security_rule.category")

	SecurityRuleDescriptionKey = attribute.Key("security_rule.description")

	SecurityRuleLicenseKey = attribute.Key("security_rule.license")

	SecurityRuleNameKey = attribute.Key("security_rule.name")

	SecurityRuleReferenceKey = attribute.Key("security_rule.reference")

	SecurityRuleRulesetNameKey = attribute.Key("security_rule.ruleset.name")

	SecurityRuleUUIDKey = attribute.Key("security_rule.uuid")

	SecurityRuleVersionKey = attribute.Key("security_rule.version")
)

func SecurityRuleCategory(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SecurityRuleDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SecurityRuleLicense(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SecurityRuleName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SecurityRuleReference(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SecurityRuleRulesetName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SecurityRuleUUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SecurityRuleVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ServerAddressKey = attribute.Key("server.address")

	ServerPortKey = attribute.Key("server.port")
)

func ServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ServiceInstanceIDKey = attribute.Key("service.instance.id")

	ServiceNameKey = attribute.Key("service.name")

	ServiceNamespaceKey = attribute.Key("service.namespace")

	ServiceVersionKey = attribute.Key("service.version")
)

func ServiceInstanceID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServiceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServiceNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServiceVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	SessionIDKey = attribute.Key("session.id")

	SessionPreviousIDKey = attribute.Key("session.previous_id")
)

func SessionID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SessionPreviousID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	SignalRConnectionStatusKey = attribute.Key("signalr.connection.status")

	SignalRTransportKey = attribute.Key("signalr.transport")
)

var (
	SignalRConnectionStatusNormalClosure = SignalRConnectionStatusKey.String("normal_closure")

	SignalRConnectionStatusTimeout = SignalRConnectionStatusKey.String("timeout")

	SignalRConnectionStatusAppShutdown = SignalRConnectionStatusKey.String("app_shutdown")
)

var (
	SignalRTransportServerSentEvents = SignalRTransportKey.String("server_sent_events")

	SignalRTransportLongPolling = SignalRTransportKey.String("long_polling")

	SignalRTransportWebSockets = SignalRTransportKey.String("web_sockets")
)

const (
	SourceAddressKey = attribute.Key("source.address")

	SourcePortKey = attribute.Key("source.port")
)

func SourceAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SourcePort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	SystemCPULogicalNumberKey = attribute.Key("system.cpu.logical_number")

	SystemDeviceKey = attribute.Key("system.device")

	SystemFilesystemModeKey = attribute.Key("system.filesystem.mode")

	SystemFilesystemMountpointKey = attribute.Key("system.filesystem.mountpoint")

	SystemFilesystemStateKey = attribute.Key("system.filesystem.state")

	SystemFilesystemTypeKey = attribute.Key("system.filesystem.type")

	SystemMemoryStateKey = attribute.Key("system.memory.state")

	SystemPagingDirectionKey = attribute.Key("system.paging.direction")

	SystemPagingStateKey = attribute.Key("system.paging.state")

	SystemPagingTypeKey = attribute.Key("system.paging.type")

	SystemProcessStatusKey = attribute.Key("system.process.status")
)

func SystemCPULogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SystemDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SystemFilesystemMode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SystemFilesystemMountpoint(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	SystemFilesystemStateUsed = SystemFilesystemStateKey.String("used")

	SystemFilesystemStateFree = SystemFilesystemStateKey.String("free")

	SystemFilesystemStateReserved = SystemFilesystemStateKey.String("reserved")
)

var (
	SystemFilesystemTypeFat32 = SystemFilesystemTypeKey.String("fat32")

	SystemFilesystemTypeExfat = SystemFilesystemTypeKey.String("exfat")

	SystemFilesystemTypeNtfs = SystemFilesystemTypeKey.String("ntfs")

	SystemFilesystemTypeRefs = SystemFilesystemTypeKey.String("refs")

	SystemFilesystemTypeHfsplus = SystemFilesystemTypeKey.String("hfsplus")

	SystemFilesystemTypeExt4 = SystemFilesystemTypeKey.String("ext4")
)

var (
	SystemMemoryStateUsed = SystemMemoryStateKey.String("used")

	SystemMemoryStateFree = SystemMemoryStateKey.String("free")

	SystemMemoryStateBuffers = SystemMemoryStateKey.String("buffers")

	SystemMemoryStateCached = SystemMemoryStateKey.String("cached")
)

var (
	SystemPagingDirectionIn = SystemPagingDirectionKey.String("in")

	SystemPagingDirectionOut = SystemPagingDirectionKey.String("out")
)

var (
	SystemPagingStateUsed = SystemPagingStateKey.String("used")

	SystemPagingStateFree = SystemPagingStateKey.String("free")
)

var (
	SystemPagingTypeMajor = SystemPagingTypeKey.String("major")

	SystemPagingTypeMinor = SystemPagingTypeKey.String("minor")
)

var (
	SystemProcessStatusRunning = SystemProcessStatusKey.String("running")

	SystemProcessStatusSleeping = SystemProcessStatusKey.String("sleeping")

	SystemProcessStatusStopped = SystemProcessStatusKey.String("stopped")

	SystemProcessStatusDefunct = SystemProcessStatusKey.String("defunct")
)

const (
	TelemetryDistroNameKey = attribute.Key("telemetry.distro.name")

	TelemetryDistroVersionKey = attribute.Key("telemetry.distro.version")

	TelemetrySDKLanguageKey = attribute.Key("telemetry.sdk.language")

	TelemetrySDKNameKey = attribute.Key("telemetry.sdk.name")

	TelemetrySDKVersionKey = attribute.Key("telemetry.sdk.version")
)

func TelemetryDistroName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TelemetryDistroVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TelemetrySDKName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TelemetrySDKVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	TelemetrySDKLanguageCPP = TelemetrySDKLanguageKey.String("cpp")

	TelemetrySDKLanguageDotnet = TelemetrySDKLanguageKey.String("dotnet")

	TelemetrySDKLanguageErlang = TelemetrySDKLanguageKey.String("erlang")

	TelemetrySDKLanguageGo = TelemetrySDKLanguageKey.String("go")

	TelemetrySDKLanguageJava = TelemetrySDKLanguageKey.String("java")

	TelemetrySDKLanguageNodejs = TelemetrySDKLanguageKey.String("nodejs")

	TelemetrySDKLanguagePHP = TelemetrySDKLanguageKey.String("php")

	TelemetrySDKLanguagePython = TelemetrySDKLanguageKey.String("python")

	TelemetrySDKLanguageRuby = TelemetrySDKLanguageKey.String("ruby")

	TelemetrySDKLanguageRust = TelemetrySDKLanguageKey.String("rust")

	TelemetrySDKLanguageSwift = TelemetrySDKLanguageKey.String("swift")

	TelemetrySDKLanguageWebJS = TelemetrySDKLanguageKey.String("webjs")
)

const (
	TestCaseNameKey = attribute.Key("test.case.name")

	TestCaseResultStatusKey = attribute.Key("test.case.result.status")

	TestSuiteNameKey = attribute.Key("test.suite.name")

	TestSuiteRunStatusKey = attribute.Key("test.suite.run.status")
)

func TestCaseName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TestSuiteName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	TestCaseResultStatusPass = TestCaseResultStatusKey.String("pass")

	TestCaseResultStatusFail = TestCaseResultStatusKey.String("fail")
)

var (
	TestSuiteRunStatusSuccess = TestSuiteRunStatusKey.String("success")

	TestSuiteRunStatusFailure = TestSuiteRunStatusKey.String("failure")

	TestSuiteRunStatusSkipped = TestSuiteRunStatusKey.String("skipped")

	TestSuiteRunStatusAborted = TestSuiteRunStatusKey.String("aborted")

	TestSuiteRunStatusTimedOut = TestSuiteRunStatusKey.String("timed_out")

	TestSuiteRunStatusInProgress = TestSuiteRunStatusKey.String("in_progress")
)

const (
	ThreadIDKey = attribute.Key("thread.id")

	ThreadNameKey = attribute.Key("thread.name")
)

func ThreadID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ThreadName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	TLSCipherKey = attribute.Key("tls.cipher")

	TLSClientCertificateKey = attribute.Key("tls.client.certificate")

	TLSClientCertificateChainKey = attribute.Key("tls.client.certificate_chain")

	TLSClientHashMd5Key = attribute.Key("tls.client.hash.md5")

	TLSClientHashSha1Key = attribute.Key("tls.client.hash.sha1")

	TLSClientHashSha256Key = attribute.Key("tls.client.hash.sha256")

	TLSClientIssuerKey = attribute.Key("tls.client.issuer")

	TLSClientJa3Key = attribute.Key("tls.client.ja3")

	TLSClientNotAfterKey = attribute.Key("tls.client.not_after")

	TLSClientNotBeforeKey = attribute.Key("tls.client.not_before")

	TLSClientSubjectKey = attribute.Key("tls.client.subject")

	TLSClientSupportedCiphersKey = attribute.Key("tls.client.supported_ciphers")

	TLSCurveKey = attribute.Key("tls.curve")

	TLSEstablishedKey = attribute.Key("tls.established")

	TLSNextProtocolKey = attribute.Key("tls.next_protocol")

	TLSProtocolNameKey = attribute.Key("tls.protocol.name")

	TLSProtocolVersionKey = attribute.Key("tls.protocol.version")

	TLSResumedKey = attribute.Key("tls.resumed")

	TLSServerCertificateKey = attribute.Key("tls.server.certificate")

	TLSServerCertificateChainKey = attribute.Key("tls.server.certificate_chain")

	TLSServerHashMd5Key = attribute.Key("tls.server.hash.md5")

	TLSServerHashSha1Key = attribute.Key("tls.server.hash.sha1")

	TLSServerHashSha256Key = attribute.Key("tls.server.hash.sha256")

	TLSServerIssuerKey = attribute.Key("tls.server.issuer")

	TLSServerJa3sKey = attribute.Key("tls.server.ja3s")

	TLSServerNotAfterKey = attribute.Key("tls.server.not_after")

	TLSServerNotBeforeKey = attribute.Key("tls.server.not_before")

	TLSServerSubjectKey = attribute.Key("tls.server.subject")
)

func TLSCipher(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientCertificate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientCertificateChain(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientHashMd5(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientHashSha1(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientHashSha256(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientIssuer(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientJa3(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientNotAfter(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientNotBefore(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientSubject(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSClientSupportedCiphers(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSCurve(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSEstablished(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSNextProtocol(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSResumed(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerCertificate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerCertificateChain(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerHashMd5(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerHashSha1(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerHashSha256(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerIssuer(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerJa3s(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerNotAfter(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerNotBefore(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func TLSServerSubject(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	TLSProtocolNameSsl = TLSProtocolNameKey.String("ssl")

	TLSProtocolNameTLS = TLSProtocolNameKey.String("tls")
)

const (
	URLDomainKey = attribute.Key("url.domain")

	URLExtensionKey = attribute.Key("url.extension")

	URLFragmentKey = attribute.Key("url.fragment")

	URLFullKey = attribute.Key("url.full")

	URLOriginalKey = attribute.Key("url.original")

	URLPathKey = attribute.Key("url.path")

	URLPortKey = attribute.Key("url.port")

	URLQueryKey = attribute.Key("url.query")

	URLRegisteredDomainKey = attribute.Key("url.registered_domain")

	URLSchemeKey = attribute.Key("url.scheme")

	URLSubdomainKey = attribute.Key("url.subdomain")

	URLTemplateKey = attribute.Key("url.template")

	URLTopLevelDomainKey = attribute.Key("url.top_level_domain")
)

func URLDomain(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLExtension(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLFragment(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLOriginal(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLPath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLQuery(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLRegisteredDomain(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLSubdomain(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLTopLevelDomain(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	UserEmailKey = attribute.Key("user.email")

	UserFullNameKey = attribute.Key("user.full_name")

	UserHashKey = attribute.Key("user.hash")

	UserIDKey = attribute.Key("user.id")

	UserNameKey = attribute.Key("user.name")

	UserRolesKey = attribute.Key("user.roles")
)

func UserEmail(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func UserFullName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func UserHash(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func UserID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func UserName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func UserRoles(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	UserAgentNameKey = attribute.Key("user_agent.name")

	UserAgentOriginalKey = attribute.Key("user_agent.original")

	UserAgentOSNameKey = attribute.Key("user_agent.os.name")

	UserAgentOSVersionKey = attribute.Key("user_agent.os.version")

	UserAgentSyntheticTypeKey = attribute.Key("user_agent.synthetic.type")

	UserAgentVersionKey = attribute.Key("user_agent.version")
)

func UserAgentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func UserAgentOriginal(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func UserAgentOSName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func UserAgentOSVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func UserAgentVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	UserAgentSyntheticTypeBot = UserAgentSyntheticTypeKey.String("bot")

	UserAgentSyntheticTypeTest = UserAgentSyntheticTypeKey.String("test")
)

const (
	VCSChangeIDKey = attribute.Key("vcs.change.id")

	VCSChangeStateKey = attribute.Key("vcs.change.state")

	VCSChangeTitleKey = attribute.Key("vcs.change.title")

	VCSLineChangeTypeKey = attribute.Key("vcs.line_change.type")

	VCSOwnerNameKey = attribute.Key("vcs.owner.name")

	VCSProviderNameKey = attribute.Key("vcs.provider.name")

	VCSRefBaseNameKey = attribute.Key("vcs.ref.base.name")

	VCSRefBaseRevisionKey = attribute.Key("vcs.ref.base.revision")

	VCSRefBaseTypeKey = attribute.Key("vcs.ref.base.type")

	VCSRefHeadNameKey = attribute.Key("vcs.ref.head.name")

	VCSRefHeadRevisionKey = attribute.Key("vcs.ref.head.revision")

	VCSRefHeadTypeKey = attribute.Key("vcs.ref.head.type")

	VCSRefTypeKey = attribute.Key("vcs.ref.type")

	VCSRepositoryNameKey = attribute.Key("vcs.repository.name")

	VCSRepositoryURLFullKey = attribute.Key("vcs.repository.url.full")

	VCSRevisionDeltaDirectionKey = attribute.Key("vcs.revision_delta.direction")
)

func VCSChangeID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func VCSChangeTitle(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func VCSOwnerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func VCSRefBaseName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func VCSRefBaseRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func VCSRefHeadName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func VCSRefHeadRevision(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func VCSRepositoryName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func VCSRepositoryURLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

var (
	VCSChangeStateOpen = VCSChangeStateKey.String("open")

	VCSChangeStateWip = VCSChangeStateKey.String("wip")

	VCSChangeStateClosed = VCSChangeStateKey.String("closed")

	VCSChangeStateMerged = VCSChangeStateKey.String("merged")
)

var (
	VCSLineChangeTypeAdded = VCSLineChangeTypeKey.String("added")

	VCSLineChangeTypeRemoved = VCSLineChangeTypeKey.String("removed")
)

var (
	VCSProviderNameGithub = VCSProviderNameKey.String("github")

	VCSProviderNameGitlab = VCSProviderNameKey.String("gitlab")

	VCSProviderNameGitea = VCSProviderNameKey.String("gitea")

	VCSProviderNameBitbucket = VCSProviderNameKey.String("bitbucket")
)

var (
	VCSRefBaseTypeBranch = VCSRefBaseTypeKey.String("branch")

	VCSRefBaseTypeTag = VCSRefBaseTypeKey.String("tag")
)

var (
	VCSRefHeadTypeBranch = VCSRefHeadTypeKey.String("branch")

	VCSRefHeadTypeTag = VCSRefHeadTypeKey.String("tag")
)

var (
	VCSRefTypeBranch = VCSRefTypeKey.String("branch")

	VCSRefTypeTag = VCSRefTypeKey.String("tag")
)

var (
	VCSRevisionDeltaDirectionBehind = VCSRevisionDeltaDirectionKey.String("behind")

	VCSRevisionDeltaDirectionAhead = VCSRevisionDeltaDirectionKey.String("ahead")
)

const (
	WebEngineDescriptionKey = attribute.Key("webengine.description")

	WebEngineNameKey = attribute.Key("webengine.name")

	WebEngineVersionKey = attribute.Key("webengine.version")
)

func WebEngineDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func WebEngineName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func WebEngineVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ZOSSmfIDKey = attribute.Key("zos.smf.id")

	ZOSSysplexNameKey = attribute.Key("zos.sysplex.name")
)

func ZOSSmfID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ZOSSysplexName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
