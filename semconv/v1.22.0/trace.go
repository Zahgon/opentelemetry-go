package semconv

import "go.opentelemetry.io/otel/attribute"

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
	PeerServiceKey = attribute.Key("peer.service")
)

func PeerService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	EnduserIDKey = attribute.Key("enduser.id")

	EnduserRoleKey = attribute.Key("enduser.role")

	EnduserScopeKey = attribute.Key("enduser.scope")
)

func EnduserID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func EnduserRole(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func EnduserScope(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ThreadDaemonKey = attribute.Key("thread.daemon")

	ThreadIDKey = attribute.Key("thread.id")

	ThreadNameKey = attribute.Key("thread.name")
)

func ThreadDaemon(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ThreadID(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ThreadName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	CodeColumnKey = attribute.Key("code.column")

	CodeFilepathKey = attribute.Key("code.filepath")

	CodeFunctionKey = attribute.Key("code.function")

	CodeLineNumberKey = attribute.Key("code.lineno")

	CodeNamespaceKey = attribute.Key("code.namespace")
)

func CodeColumn(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeFilepath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeFunction(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeLineNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeNamespace(val string) attribute.KeyValue {
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
	CloudeventsEventIDKey = attribute.Key("cloudevents.event_id")

	CloudeventsEventSourceKey = attribute.Key("cloudevents.event_source")

	CloudeventsEventSpecVersionKey = attribute.Key("cloudevents.event_spec_version")

	CloudeventsEventSubjectKey = attribute.Key("cloudevents.event_subject")

	CloudeventsEventTypeKey = attribute.Key("cloudevents.event_type")
)

func CloudeventsEventID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudeventsEventSource(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudeventsEventSpecVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudeventsEventSubject(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudeventsEventType(val string) attribute.KeyValue {
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
	DBConnectionStringKey = attribute.Key("db.connection_string")

	DBJDBCDriverClassnameKey = attribute.Key("db.jdbc.driver_classname")

	DBNameKey = attribute.Key("db.name")

	DBOperationKey = attribute.Key("db.operation")

	DBStatementKey = attribute.Key("db.statement")

	DBSystemKey = attribute.Key("db.system")

	DBUserKey = attribute.Key("db.user")
)

var (
	DBSystemOtherSQL = DBSystemKey.String("other_sql")

	DBSystemMSSQL = DBSystemKey.String("mssql")

	DBSystemMssqlcompact = DBSystemKey.String("mssqlcompact")

	DBSystemMySQL = DBSystemKey.String("mysql")

	DBSystemOracle = DBSystemKey.String("oracle")

	DBSystemDB2 = DBSystemKey.String("db2")

	DBSystemPostgreSQL = DBSystemKey.String("postgresql")

	DBSystemRedshift = DBSystemKey.String("redshift")

	DBSystemHive = DBSystemKey.String("hive")

	DBSystemCloudscape = DBSystemKey.String("cloudscape")

	DBSystemHSQLDB = DBSystemKey.String("hsqldb")

	DBSystemProgress = DBSystemKey.String("progress")

	DBSystemMaxDB = DBSystemKey.String("maxdb")

	DBSystemHanaDB = DBSystemKey.String("hanadb")

	DBSystemIngres = DBSystemKey.String("ingres")

	DBSystemFirstSQL = DBSystemKey.String("firstsql")

	DBSystemEDB = DBSystemKey.String("edb")

	DBSystemCache = DBSystemKey.String("cache")

	DBSystemAdabas = DBSystemKey.String("adabas")

	DBSystemFirebird = DBSystemKey.String("firebird")

	DBSystemDerby = DBSystemKey.String("derby")

	DBSystemFilemaker = DBSystemKey.String("filemaker")

	DBSystemInformix = DBSystemKey.String("informix")

	DBSystemInstantDB = DBSystemKey.String("instantdb")

	DBSystemInterbase = DBSystemKey.String("interbase")

	DBSystemMariaDB = DBSystemKey.String("mariadb")

	DBSystemNetezza = DBSystemKey.String("netezza")

	DBSystemPervasive = DBSystemKey.String("pervasive")

	DBSystemPointbase = DBSystemKey.String("pointbase")

	DBSystemSqlite = DBSystemKey.String("sqlite")

	DBSystemSybase = DBSystemKey.String("sybase")

	DBSystemTeradata = DBSystemKey.String("teradata")

	DBSystemVertica = DBSystemKey.String("vertica")

	DBSystemH2 = DBSystemKey.String("h2")

	DBSystemColdfusion = DBSystemKey.String("coldfusion")

	DBSystemCassandra = DBSystemKey.String("cassandra")

	DBSystemHBase = DBSystemKey.String("hbase")

	DBSystemMongoDB = DBSystemKey.String("mongodb")

	DBSystemRedis = DBSystemKey.String("redis")

	DBSystemCouchbase = DBSystemKey.String("couchbase")

	DBSystemCouchDB = DBSystemKey.String("couchdb")

	DBSystemCosmosDB = DBSystemKey.String("cosmosdb")

	DBSystemDynamoDB = DBSystemKey.String("dynamodb")

	DBSystemNeo4j = DBSystemKey.String("neo4j")

	DBSystemGeode = DBSystemKey.String("geode")

	DBSystemElasticsearch = DBSystemKey.String("elasticsearch")

	DBSystemMemcached = DBSystemKey.String("memcached")

	DBSystemCockroachdb = DBSystemKey.String("cockroachdb")

	DBSystemOpensearch = DBSystemKey.String("opensearch")

	DBSystemClickhouse = DBSystemKey.String("clickhouse")

	DBSystemSpanner = DBSystemKey.String("spanner")

	DBSystemTrino = DBSystemKey.String("trino")
)

func DBConnectionString(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBJDBCDriverClassname(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBOperation(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBStatement(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBUser(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DBMSSQLInstanceNameKey = attribute.Key("db.mssql.instance_name")
)

func DBMSSQLInstanceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DBCassandraConsistencyLevelKey = attribute.Key("db.cassandra.consistency_level")

	DBCassandraCoordinatorDCKey = attribute.Key("db.cassandra.coordinator.dc")

	DBCassandraCoordinatorIDKey = attribute.Key("db.cassandra.coordinator.id")

	DBCassandraIdempotenceKey = attribute.Key("db.cassandra.idempotence")

	DBCassandraPageSizeKey = attribute.Key("db.cassandra.page_size")

	DBCassandraSpeculativeExecutionCountKey = attribute.Key("db.cassandra.speculative_execution_count")

	DBCassandraTableKey = attribute.Key("db.cassandra.table")
)

var (
	DBCassandraConsistencyLevelAll = DBCassandraConsistencyLevelKey.String("all")

	DBCassandraConsistencyLevelEachQuorum = DBCassandraConsistencyLevelKey.String("each_quorum")

	DBCassandraConsistencyLevelQuorum = DBCassandraConsistencyLevelKey.String("quorum")

	DBCassandraConsistencyLevelLocalQuorum = DBCassandraConsistencyLevelKey.String("local_quorum")

	DBCassandraConsistencyLevelOne = DBCassandraConsistencyLevelKey.String("one")

	DBCassandraConsistencyLevelTwo = DBCassandraConsistencyLevelKey.String("two")

	DBCassandraConsistencyLevelThree = DBCassandraConsistencyLevelKey.String("three")

	DBCassandraConsistencyLevelLocalOne = DBCassandraConsistencyLevelKey.String("local_one")

	DBCassandraConsistencyLevelAny = DBCassandraConsistencyLevelKey.String("any")

	DBCassandraConsistencyLevelSerial = DBCassandraConsistencyLevelKey.String("serial")

	DBCassandraConsistencyLevelLocalSerial = DBCassandraConsistencyLevelKey.String("local_serial")
)

func DBCassandraCoordinatorDC(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraCoordinatorID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraIdempotence(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraPageSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraSpeculativeExecutionCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraTable(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DBRedisDBIndexKey = attribute.Key("db.redis.database_index")
)

func DBRedisDBIndex(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DBMongoDBCollectionKey = attribute.Key("db.mongodb.collection")
)

func DBMongoDBCollection(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DBElasticsearchClusterNameKey = attribute.Key("db.elasticsearch.cluster.name")

	DBElasticsearchNodeNameKey = attribute.Key("db.elasticsearch.node.name")
)

func DBElasticsearchClusterName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBElasticsearchNodeName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DBSQLTableKey = attribute.Key("db.sql.table")
)

func DBSQLTable(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DBCosmosDBClientIDKey = attribute.Key("db.cosmosdb.client_id")

	DBCosmosDBConnectionModeKey = attribute.Key("db.cosmosdb.connection_mode")

	DBCosmosDBContainerKey = attribute.Key("db.cosmosdb.container")

	DBCosmosDBOperationTypeKey = attribute.Key("db.cosmosdb.operation_type")

	DBCosmosDBRequestChargeKey = attribute.Key("db.cosmosdb.request_charge")

	DBCosmosDBRequestContentLengthKey = attribute.Key("db.cosmosdb.request_content_length")

	DBCosmosDBStatusCodeKey = attribute.Key("db.cosmosdb.status_code")

	DBCosmosDBSubStatusCodeKey = attribute.Key("db.cosmosdb.sub_status_code")
)

var (
	DBCosmosDBConnectionModeGateway = DBCosmosDBConnectionModeKey.String("gateway")

	DBCosmosDBConnectionModeDirect = DBCosmosDBConnectionModeKey.String("direct")
)

var (
	DBCosmosDBOperationTypeInvalid = DBCosmosDBOperationTypeKey.String("Invalid")

	DBCosmosDBOperationTypeCreate = DBCosmosDBOperationTypeKey.String("Create")

	DBCosmosDBOperationTypePatch = DBCosmosDBOperationTypeKey.String("Patch")

	DBCosmosDBOperationTypeRead = DBCosmosDBOperationTypeKey.String("Read")

	DBCosmosDBOperationTypeReadFeed = DBCosmosDBOperationTypeKey.String("ReadFeed")

	DBCosmosDBOperationTypeDelete = DBCosmosDBOperationTypeKey.String("Delete")

	DBCosmosDBOperationTypeReplace = DBCosmosDBOperationTypeKey.String("Replace")

	DBCosmosDBOperationTypeExecute = DBCosmosDBOperationTypeKey.String("Execute")

	DBCosmosDBOperationTypeQuery = DBCosmosDBOperationTypeKey.String("Query")

	DBCosmosDBOperationTypeHead = DBCosmosDBOperationTypeKey.String("Head")

	DBCosmosDBOperationTypeHeadFeed = DBCosmosDBOperationTypeKey.String("HeadFeed")

	DBCosmosDBOperationTypeUpsert = DBCosmosDBOperationTypeKey.String("Upsert")

	DBCosmosDBOperationTypeBatch = DBCosmosDBOperationTypeKey.String("Batch")

	DBCosmosDBOperationTypeQueryPlan = DBCosmosDBOperationTypeKey.String("QueryPlan")

	DBCosmosDBOperationTypeExecuteJavascript = DBCosmosDBOperationTypeKey.String("ExecuteJavaScript")
)

func DBCosmosDBClientID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCosmosDBContainer(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCosmosDBRequestCharge(val float64) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCosmosDBRequestContentLength(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCosmosDBStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCosmosDBSubStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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
	FaaSInvocationIDKey = attribute.Key("faas.invocation_id")
)

func FaaSInvocationID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	FaaSDocumentCollectionKey = attribute.Key("faas.document.collection")

	FaaSDocumentNameKey = attribute.Key("faas.document.name")

	FaaSDocumentOperationKey = attribute.Key("faas.document.operation")

	FaaSDocumentTimeKey = attribute.Key("faas.document.time")
)

var (
	FaaSDocumentOperationInsert = FaaSDocumentOperationKey.String("insert")

	FaaSDocumentOperationEdit = FaaSDocumentOperationKey.String("edit")

	FaaSDocumentOperationDelete = FaaSDocumentOperationKey.String("delete")
)

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

const (
	FaaSCronKey = attribute.Key("faas.cron")

	FaaSTimeKey = attribute.Key("faas.time")
)

func FaaSCron(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSTime(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	FaaSColdstartKey = attribute.Key("faas.coldstart")
)

func FaaSColdstart(val bool) attribute.KeyValue {
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
	AWSDynamoDBAttributesToGetKey = attribute.Key("aws.dynamodb.attributes_to_get")

	AWSDynamoDBConsistentReadKey = attribute.Key("aws.dynamodb.consistent_read")

	AWSDynamoDBConsumedCapacityKey = attribute.Key("aws.dynamodb.consumed_capacity")

	AWSDynamoDBIndexNameKey = attribute.Key("aws.dynamodb.index_name")

	AWSDynamoDBItemCollectionMetricsKey = attribute.Key("aws.dynamodb.item_collection_metrics")

	AWSDynamoDBLimitKey = attribute.Key("aws.dynamodb.limit")

	AWSDynamoDBProjectionKey = attribute.Key("aws.dynamodb.projection")

	AWSDynamoDBProvisionedReadCapacityKey = attribute.Key("aws.dynamodb.provisioned_read_capacity")

	AWSDynamoDBProvisionedWriteCapacityKey = attribute.Key("aws.dynamodb.provisioned_write_capacity")

	AWSDynamoDBSelectKey = attribute.Key("aws.dynamodb.select")

	AWSDynamoDBTableNamesKey = attribute.Key("aws.dynamodb.table_names")
)

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

func AWSDynamoDBSelect(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBTableNames(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSDynamoDBGlobalSecondaryIndexesKey = attribute.Key("aws.dynamodb.global_secondary_indexes")

	AWSDynamoDBLocalSecondaryIndexesKey = attribute.Key("aws.dynamodb.local_secondary_indexes")
)

func AWSDynamoDBGlobalSecondaryIndexes(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBLocalSecondaryIndexes(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSDynamoDBExclusiveStartTableKey = attribute.Key("aws.dynamodb.exclusive_start_table")

	AWSDynamoDBTableCountKey = attribute.Key("aws.dynamodb.table_count")
)

func AWSDynamoDBExclusiveStartTable(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBTableCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSDynamoDBScanForwardKey = attribute.Key("aws.dynamodb.scan_forward")
)

func AWSDynamoDBScanForward(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSDynamoDBCountKey = attribute.Key("aws.dynamodb.count")

	AWSDynamoDBScannedCountKey = attribute.Key("aws.dynamodb.scanned_count")

	AWSDynamoDBSegmentKey = attribute.Key("aws.dynamodb.segment")

	AWSDynamoDBTotalSegmentsKey = attribute.Key("aws.dynamodb.total_segments")
)

func AWSDynamoDBCount(val int) attribute.KeyValue {
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

func AWSDynamoDBTotalSegments(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	AWSDynamoDBAttributeDefinitionsKey = attribute.Key("aws.dynamodb.attribute_definitions")

	AWSDynamoDBGlobalSecondaryIndexUpdatesKey = attribute.Key("aws.dynamodb.global_secondary_index_updates")
)

func AWSDynamoDBAttributeDefinitions(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBGlobalSecondaryIndexUpdates(val ...string) attribute.KeyValue {
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

const (
	MessagingBatchMessageCountKey = attribute.Key("messaging.batch.message_count")

	MessagingClientIDKey = attribute.Key("messaging.client_id")

	MessagingOperationKey = attribute.Key("messaging.operation")

	MessagingSystemKey = attribute.Key("messaging.system")
)

var (
	MessagingOperationPublish = MessagingOperationKey.String("publish")

	MessagingOperationReceive = MessagingOperationKey.String("receive")

	MessagingOperationProcess = MessagingOperationKey.String("process")
)

func MessagingBatchMessageCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingClientID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingSystem(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	RPCMethodKey = attribute.Key("rpc.method")

	RPCServiceKey = attribute.Key("rpc.service")

	RPCSystemKey = attribute.Key("rpc.system")
)

var (
	RPCSystemGRPC = RPCSystemKey.String("grpc")

	RPCSystemJavaRmi = RPCSystemKey.String("java_rmi")

	RPCSystemDotnetWcf = RPCSystemKey.String("dotnet_wcf")

	RPCSystemApacheDubbo = RPCSystemKey.String("apache_dubbo")

	RPCSystemConnectRPC = RPCSystemKey.String("connect_rpc")
)

func RPCMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	RPCGRPCStatusCodeKey = attribute.Key("rpc.grpc.status_code")
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

const (
	RPCJsonrpcErrorCodeKey = attribute.Key("rpc.jsonrpc.error_code")

	RPCJsonrpcErrorMessageKey = attribute.Key("rpc.jsonrpc.error_message")

	RPCJsonrpcRequestIDKey = attribute.Key("rpc.jsonrpc.request_id")

	RPCJsonrpcVersionKey = attribute.Key("rpc.jsonrpc.version")
)

func RPCJsonrpcErrorCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJsonrpcErrorMessage(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJsonrpcRequestID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJsonrpcVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	RPCConnectRPCErrorCodeKey = attribute.Key("rpc.connect_rpc.error_code")
)

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
