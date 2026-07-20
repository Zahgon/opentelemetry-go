package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	AWSLambdaInvokedARNKey = attribute.Key("aws.lambda.invoked_arn")
)

const (
	DBSystemKey = attribute.Key("db.system")

	DBConnectionStringKey = attribute.Key("db.connection_string")

	DBUserKey = attribute.Key("db.user")

	DBJDBCDriverClassnameKey = attribute.Key("db.jdbc.driver_classname")

	DBNameKey = attribute.Key("db.name")

	DBStatementKey = attribute.Key("db.statement")

	DBOperationKey = attribute.Key("db.operation")
)

var (
	DBSystemOtherSQL = DBSystemKey.String("other_sql")

	DBSystemMSSQL = DBSystemKey.String("mssql")

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
)

const (
	DBMSSQLInstanceNameKey = attribute.Key("db.mssql.instance_name")
)

const (
	DBCassandraPageSizeKey = attribute.Key("db.cassandra.page_size")

	DBCassandraConsistencyLevelKey = attribute.Key("db.cassandra.consistency_level")

	DBCassandraTableKey = attribute.Key("db.cassandra.table")

	DBCassandraIdempotenceKey = attribute.Key("db.cassandra.idempotence")

	DBCassandraSpeculativeExecutionCountKey = attribute.Key("db.cassandra.speculative_execution_count")

	DBCassandraCoordinatorIDKey = attribute.Key("db.cassandra.coordinator.id")

	DBCassandraCoordinatorDCKey = attribute.Key("db.cassandra.coordinator.dc")
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

const (
	DBRedisDBIndexKey = attribute.Key("db.redis.database_index")
)

const (
	DBMongoDBCollectionKey = attribute.Key("db.mongodb.collection")
)

const (
	DBSQLTableKey = attribute.Key("db.sql.table")
)

const (
	ExceptionTypeKey = attribute.Key("exception.type")

	ExceptionMessageKey = attribute.Key("exception.message")

	ExceptionStacktraceKey = attribute.Key("exception.stacktrace")

	ExceptionEscapedKey = attribute.Key("exception.escaped")
)

const (
	FaaSTriggerKey = attribute.Key("faas.trigger")

	FaaSExecutionKey = attribute.Key("faas.execution")
)

var (
	FaaSTriggerDatasource = FaaSTriggerKey.String("datasource")

	FaaSTriggerHTTP = FaaSTriggerKey.String("http")

	FaaSTriggerPubsub = FaaSTriggerKey.String("pubsub")

	FaaSTriggerTimer = FaaSTriggerKey.String("timer")

	FaaSTriggerOther = FaaSTriggerKey.String("other")
)

const (
	FaaSDocumentCollectionKey = attribute.Key("faas.document.collection")

	FaaSDocumentOperationKey = attribute.Key("faas.document.operation")

	FaaSDocumentTimeKey = attribute.Key("faas.document.time")

	FaaSDocumentNameKey = attribute.Key("faas.document.name")
)

var (
	FaaSDocumentOperationInsert = FaaSDocumentOperationKey.String("insert")

	FaaSDocumentOperationEdit = FaaSDocumentOperationKey.String("edit")

	FaaSDocumentOperationDelete = FaaSDocumentOperationKey.String("delete")
)

const (
	FaaSTimeKey = attribute.Key("faas.time")

	FaaSCronKey = attribute.Key("faas.cron")
)

const (
	FaaSColdstartKey = attribute.Key("faas.coldstart")
)

const (
	FaaSInvokedNameKey = attribute.Key("faas.invoked_name")

	FaaSInvokedProviderKey = attribute.Key("faas.invoked_provider")

	FaaSInvokedRegionKey = attribute.Key("faas.invoked_region")
)

var (
	FaaSInvokedProviderAlibabaCloud = FaaSInvokedProviderKey.String("alibaba_cloud")

	FaaSInvokedProviderAWS = FaaSInvokedProviderKey.String("aws")

	FaaSInvokedProviderAzure = FaaSInvokedProviderKey.String("azure")

	FaaSInvokedProviderGCP = FaaSInvokedProviderKey.String("gcp")

	FaaSInvokedProviderTencentCloud = FaaSInvokedProviderKey.String("tencent_cloud")
)

const (
	NetTransportKey = attribute.Key("net.transport")

	NetPeerIPKey = attribute.Key("net.peer.ip")

	NetPeerPortKey = attribute.Key("net.peer.port")

	NetPeerNameKey = attribute.Key("net.peer.name")

	NetHostIPKey = attribute.Key("net.host.ip")

	NetHostPortKey = attribute.Key("net.host.port")

	NetHostNameKey = attribute.Key("net.host.name")

	NetHostConnectionTypeKey = attribute.Key("net.host.connection.type")

	NetHostConnectionSubtypeKey = attribute.Key("net.host.connection.subtype")

	NetHostCarrierNameKey = attribute.Key("net.host.carrier.name")

	NetHostCarrierMccKey = attribute.Key("net.host.carrier.mcc")

	NetHostCarrierMncKey = attribute.Key("net.host.carrier.mnc")

	NetHostCarrierIccKey = attribute.Key("net.host.carrier.icc")
)

var (
	NetTransportTCP = NetTransportKey.String("ip_tcp")

	NetTransportUDP = NetTransportKey.String("ip_udp")

	NetTransportIP = NetTransportKey.String("ip")

	NetTransportUnix = NetTransportKey.String("unix")

	NetTransportPipe = NetTransportKey.String("pipe")

	NetTransportInProc = NetTransportKey.String("inproc")

	NetTransportOther = NetTransportKey.String("other")
)

var (
	NetHostConnectionTypeWifi = NetHostConnectionTypeKey.String("wifi")

	NetHostConnectionTypeWired = NetHostConnectionTypeKey.String("wired")

	NetHostConnectionTypeCell = NetHostConnectionTypeKey.String("cell")

	NetHostConnectionTypeUnavailable = NetHostConnectionTypeKey.String("unavailable")

	NetHostConnectionTypeUnknown = NetHostConnectionTypeKey.String("unknown")
)

var (
	NetHostConnectionSubtypeGprs = NetHostConnectionSubtypeKey.String("gprs")

	NetHostConnectionSubtypeEdge = NetHostConnectionSubtypeKey.String("edge")

	NetHostConnectionSubtypeUmts = NetHostConnectionSubtypeKey.String("umts")

	NetHostConnectionSubtypeCdma = NetHostConnectionSubtypeKey.String("cdma")

	NetHostConnectionSubtypeEvdo0 = NetHostConnectionSubtypeKey.String("evdo_0")

	NetHostConnectionSubtypeEvdoA = NetHostConnectionSubtypeKey.String("evdo_a")

	NetHostConnectionSubtypeCdma20001xrtt = NetHostConnectionSubtypeKey.String("cdma2000_1xrtt")

	NetHostConnectionSubtypeHsdpa = NetHostConnectionSubtypeKey.String("hsdpa")

	NetHostConnectionSubtypeHsupa = NetHostConnectionSubtypeKey.String("hsupa")

	NetHostConnectionSubtypeHspa = NetHostConnectionSubtypeKey.String("hspa")

	NetHostConnectionSubtypeIden = NetHostConnectionSubtypeKey.String("iden")

	NetHostConnectionSubtypeEvdoB = NetHostConnectionSubtypeKey.String("evdo_b")

	NetHostConnectionSubtypeLte = NetHostConnectionSubtypeKey.String("lte")

	NetHostConnectionSubtypeEhrpd = NetHostConnectionSubtypeKey.String("ehrpd")

	NetHostConnectionSubtypeHspap = NetHostConnectionSubtypeKey.String("hspap")

	NetHostConnectionSubtypeGsm = NetHostConnectionSubtypeKey.String("gsm")

	NetHostConnectionSubtypeTdScdma = NetHostConnectionSubtypeKey.String("td_scdma")

	NetHostConnectionSubtypeIwlan = NetHostConnectionSubtypeKey.String("iwlan")

	NetHostConnectionSubtypeNr = NetHostConnectionSubtypeKey.String("nr")

	NetHostConnectionSubtypeNrnsa = NetHostConnectionSubtypeKey.String("nrnsa")

	NetHostConnectionSubtypeLteCa = NetHostConnectionSubtypeKey.String("lte_ca")
)

const (
	PeerServiceKey = attribute.Key("peer.service")
)

const (
	EnduserIDKey = attribute.Key("enduser.id")

	EnduserRoleKey = attribute.Key("enduser.role")

	EnduserScopeKey = attribute.Key("enduser.scope")
)

const (
	ThreadIDKey = attribute.Key("thread.id")

	ThreadNameKey = attribute.Key("thread.name")
)

const (
	CodeFunctionKey = attribute.Key("code.function")

	CodeNamespaceKey = attribute.Key("code.namespace")

	CodeFilepathKey = attribute.Key("code.filepath")

	CodeLineNumberKey = attribute.Key("code.lineno")
)

const (
	HTTPMethodKey = attribute.Key("http.method")

	HTTPURLKey = attribute.Key("http.url")

	HTTPTargetKey = attribute.Key("http.target")

	HTTPHostKey = attribute.Key("http.host")

	HTTPSchemeKey = attribute.Key("http.scheme")

	HTTPStatusCodeKey = attribute.Key("http.status_code")

	HTTPFlavorKey = attribute.Key("http.flavor")

	HTTPUserAgentKey = attribute.Key("http.user_agent")

	HTTPRequestContentLengthKey = attribute.Key("http.request_content_length")

	HTTPRequestContentLengthUncompressedKey = attribute.Key("http.request_content_length_uncompressed")

	HTTPResponseContentLengthKey = attribute.Key("http.response_content_length")

	HTTPResponseContentLengthUncompressedKey = attribute.Key("http.response_content_length_uncompressed")
)

var (
	HTTPFlavorHTTP10 = HTTPFlavorKey.String("1.0")

	HTTPFlavorHTTP11 = HTTPFlavorKey.String("1.1")

	HTTPFlavorHTTP20 = HTTPFlavorKey.String("2.0")

	HTTPFlavorSPDY = HTTPFlavorKey.String("SPDY")

	HTTPFlavorQUIC = HTTPFlavorKey.String("QUIC")
)

const (
	HTTPServerNameKey = attribute.Key("http.server_name")

	HTTPRouteKey = attribute.Key("http.route")

	HTTPClientIPKey = attribute.Key("http.client_ip")
)

const (
	AWSDynamoDBTableNamesKey = attribute.Key("aws.dynamodb.table_names")

	AWSDynamoDBConsumedCapacityKey = attribute.Key("aws.dynamodb.consumed_capacity")

	AWSDynamoDBItemCollectionMetricsKey = attribute.Key("aws.dynamodb.item_collection_metrics")

	AWSDynamoDBProvisionedReadCapacityKey = attribute.Key("aws.dynamodb.provisioned_read_capacity")

	AWSDynamoDBProvisionedWriteCapacityKey = attribute.Key("aws.dynamodb.provisioned_write_capacity")

	AWSDynamoDBConsistentReadKey = attribute.Key("aws.dynamodb.consistent_read")

	AWSDynamoDBProjectionKey = attribute.Key("aws.dynamodb.projection")

	AWSDynamoDBLimitKey = attribute.Key("aws.dynamodb.limit")

	AWSDynamoDBAttributesToGetKey = attribute.Key("aws.dynamodb.attributes_to_get")

	AWSDynamoDBIndexNameKey = attribute.Key("aws.dynamodb.index_name")

	AWSDynamoDBSelectKey = attribute.Key("aws.dynamodb.select")
)

const (
	AWSDynamoDBGlobalSecondaryIndexesKey = attribute.Key("aws.dynamodb.global_secondary_indexes")

	AWSDynamoDBLocalSecondaryIndexesKey = attribute.Key("aws.dynamodb.local_secondary_indexes")
)

const (
	AWSDynamoDBExclusiveStartTableKey = attribute.Key("aws.dynamodb.exclusive_start_table")

	AWSDynamoDBTableCountKey = attribute.Key("aws.dynamodb.table_count")
)

const (
	AWSDynamoDBScanForwardKey = attribute.Key("aws.dynamodb.scan_forward")
)

const (
	AWSDynamoDBSegmentKey = attribute.Key("aws.dynamodb.segment")

	AWSDynamoDBTotalSegmentsKey = attribute.Key("aws.dynamodb.total_segments")

	AWSDynamoDBCountKey = attribute.Key("aws.dynamodb.count")

	AWSDynamoDBScannedCountKey = attribute.Key("aws.dynamodb.scanned_count")
)

const (
	AWSDynamoDBAttributeDefinitionsKey = attribute.Key("aws.dynamodb.attribute_definitions")

	AWSDynamoDBGlobalSecondaryIndexUpdatesKey = attribute.Key("aws.dynamodb.global_secondary_index_updates")
)

const (
	MessagingSystemKey = attribute.Key("messaging.system")

	MessagingDestinationKey = attribute.Key("messaging.destination")

	MessagingDestinationKindKey = attribute.Key("messaging.destination_kind")

	MessagingTempDestinationKey = attribute.Key("messaging.temp_destination")

	MessagingProtocolKey = attribute.Key("messaging.protocol")

	MessagingProtocolVersionKey = attribute.Key("messaging.protocol_version")

	MessagingURLKey = attribute.Key("messaging.url")

	MessagingMessageIDKey = attribute.Key("messaging.message_id")

	MessagingConversationIDKey = attribute.Key("messaging.conversation_id")

	MessagingMessagePayloadSizeBytesKey = attribute.Key("messaging.message_payload_size_bytes")

	MessagingMessagePayloadCompressedSizeBytesKey = attribute.Key("messaging.message_payload_compressed_size_bytes")
)

var (
	MessagingDestinationKindQueue = MessagingDestinationKindKey.String("queue")

	MessagingDestinationKindTopic = MessagingDestinationKindKey.String("topic")
)

const (
	MessagingOperationKey = attribute.Key("messaging.operation")

	MessagingConsumerIDKey = attribute.Key("messaging.consumer_id")
)

var (
	MessagingOperationReceive = MessagingOperationKey.String("receive")

	MessagingOperationProcess = MessagingOperationKey.String("process")
)

const (
	MessagingRabbitmqRoutingKeyKey = attribute.Key("messaging.rabbitmq.routing_key")
)

const (
	MessagingKafkaMessageKeyKey = attribute.Key("messaging.kafka.message_key")

	MessagingKafkaConsumerGroupKey = attribute.Key("messaging.kafka.consumer_group")

	MessagingKafkaClientIDKey = attribute.Key("messaging.kafka.client_id")

	MessagingKafkaPartitionKey = attribute.Key("messaging.kafka.partition")

	MessagingKafkaTombstoneKey = attribute.Key("messaging.kafka.tombstone")
)

const (
	MessagingRocketmqNamespaceKey = attribute.Key("messaging.rocketmq.namespace")

	MessagingRocketmqClientGroupKey = attribute.Key("messaging.rocketmq.client_group")

	MessagingRocketmqClientIDKey = attribute.Key("messaging.rocketmq.client_id")

	MessagingRocketmqMessageTypeKey = attribute.Key("messaging.rocketmq.message_type")

	MessagingRocketmqMessageTagKey = attribute.Key("messaging.rocketmq.message_tag")

	MessagingRocketmqMessageKeysKey = attribute.Key("messaging.rocketmq.message_keys")

	MessagingRocketmqConsumptionModelKey = attribute.Key("messaging.rocketmq.consumption_model")
)

var (
	MessagingRocketmqMessageTypeNormal = MessagingRocketmqMessageTypeKey.String("normal")

	MessagingRocketmqMessageTypeFifo = MessagingRocketmqMessageTypeKey.String("fifo")

	MessagingRocketmqMessageTypeDelay = MessagingRocketmqMessageTypeKey.String("delay")

	MessagingRocketmqMessageTypeTransaction = MessagingRocketmqMessageTypeKey.String("transaction")
)

var (
	MessagingRocketmqConsumptionModelClustering = MessagingRocketmqConsumptionModelKey.String("clustering")

	MessagingRocketmqConsumptionModelBroadcasting = MessagingRocketmqConsumptionModelKey.String("broadcasting")
)

const (
	RPCSystemKey = attribute.Key("rpc.system")

	RPCServiceKey = attribute.Key("rpc.service")

	RPCMethodKey = attribute.Key("rpc.method")
)

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
	RPCJsonrpcVersionKey = attribute.Key("rpc.jsonrpc.version")

	RPCJsonrpcRequestIDKey = attribute.Key("rpc.jsonrpc.request_id")

	RPCJsonrpcErrorCodeKey = attribute.Key("rpc.jsonrpc.error_code")

	RPCJsonrpcErrorMessageKey = attribute.Key("rpc.jsonrpc.error_message")
)

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
