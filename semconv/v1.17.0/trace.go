package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	ExceptionTypeKey = attribute.Key("exception.type")

	ExceptionMessageKey = attribute.Key("exception.message")

	ExceptionStacktraceKey = attribute.Key("exception.stacktrace")
)

func ExceptionType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ExceptionMessage(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ExceptionStacktrace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	EventNameKey = attribute.Key("event.name")

	EventDomainKey = attribute.Key("event.domain")
)

var (
	EventDomainBrowser = EventDomainKey.String("browser")

	EventDomainDevice = EventDomainKey.String("device")

	EventDomainK8S = EventDomainKey.String("k8s")
)

func EventName(val string) attribute.KeyValue {
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

	CloudeventsEventTypeKey = attribute.Key("cloudevents.event_type")

	CloudeventsEventSubjectKey = attribute.Key("cloudevents.event_subject")
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

func CloudeventsEventType(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CloudeventsEventSubject(val string) attribute.KeyValue {
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

	DBSystemOpensearch = DBSystemKey.String("opensearch")

	DBSystemClickhouse = DBSystemKey.String("clickhouse")
)

func DBConnectionString(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBUser(val string) attribute.KeyValue {
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

func DBStatement(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBOperation(val string) attribute.KeyValue {
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

func DBCassandraPageSize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraTable(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraIdempotence(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraSpeculativeExecutionCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraCoordinatorID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DBCassandraCoordinatorDC(val string) attribute.KeyValue {
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
	DBSQLTableKey = attribute.Key("db.sql.table")
)

func DBSQLTable(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	OtelStatusCodeKey = attribute.Key("otel.status_code")

	OtelStatusDescriptionKey = attribute.Key("otel.status_description")
)

var (
	OtelStatusCodeOk = OtelStatusCodeKey.String("OK")

	OtelStatusCodeError = OtelStatusCodeKey.String("ERROR")
)

func OtelStatusDescription(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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

func FaaSExecution(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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

func FaaSDocumentCollection(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSDocumentTime(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSDocumentName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	FaaSTimeKey = attribute.Key("faas.time")

	FaaSCronKey = attribute.Key("faas.cron")
)

func FaaSTime(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSCron(val string) attribute.KeyValue {
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

func FaaSInvokedName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func FaaSInvokedRegion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	NetTransportKey = attribute.Key("net.transport")

	NetAppProtocolNameKey = attribute.Key("net.app.protocol.name")

	NetAppProtocolVersionKey = attribute.Key("net.app.protocol.version")

	NetSockPeerNameKey = attribute.Key("net.sock.peer.name")

	NetSockPeerAddrKey = attribute.Key("net.sock.peer.addr")

	NetSockPeerPortKey = attribute.Key("net.sock.peer.port")

	NetSockFamilyKey = attribute.Key("net.sock.family")

	NetPeerNameKey = attribute.Key("net.peer.name")

	NetPeerPortKey = attribute.Key("net.peer.port")

	NetHostNameKey = attribute.Key("net.host.name")

	NetHostPortKey = attribute.Key("net.host.port")

	NetSockHostAddrKey = attribute.Key("net.sock.host.addr")

	NetSockHostPortKey = attribute.Key("net.sock.host.port")

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

	NetTransportPipe = NetTransportKey.String("pipe")

	NetTransportInProc = NetTransportKey.String("inproc")

	NetTransportOther = NetTransportKey.String("other")
)

var (
	NetSockFamilyInet = NetSockFamilyKey.String("inet")

	NetSockFamilyInet6 = NetSockFamilyKey.String("inet6")

	NetSockFamilyUnix = NetSockFamilyKey.String("unix")
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

func NetAppProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetAppProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetSockPeerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetSockPeerAddr(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetSockPeerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetPeerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetPeerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetHostName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetHostPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetSockHostAddr(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetSockHostPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetHostCarrierName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetHostCarrierMcc(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetHostCarrierMnc(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetHostCarrierIcc(val string) attribute.KeyValue {
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
	CodeFunctionKey = attribute.Key("code.function")

	CodeNamespaceKey = attribute.Key("code.namespace")

	CodeFilepathKey = attribute.Key("code.filepath")

	CodeLineNumberKey = attribute.Key("code.lineno")

	CodeColumnKey = attribute.Key("code.column")
)

func CodeFunction(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeFilepath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeLineNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func CodeColumn(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HTTPMethodKey = attribute.Key("http.method")

	HTTPStatusCodeKey = attribute.Key("http.status_code")

	HTTPFlavorKey = attribute.Key("http.flavor")

	HTTPUserAgentKey = attribute.Key("http.user_agent")

	HTTPRequestContentLengthKey = attribute.Key("http.request_content_length")

	HTTPResponseContentLengthKey = attribute.Key("http.response_content_length")
)

var (
	HTTPFlavorHTTP10 = HTTPFlavorKey.String("1.0")

	HTTPFlavorHTTP11 = HTTPFlavorKey.String("1.1")

	HTTPFlavorHTTP20 = HTTPFlavorKey.String("2.0")

	HTTPFlavorHTTP30 = HTTPFlavorKey.String("3.0")

	HTTPFlavorSPDY = HTTPFlavorKey.String("SPDY")

	HTTPFlavorQUIC = HTTPFlavorKey.String("QUIC")
)

func HTTPMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPUserAgent(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPRequestContentLength(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPResponseContentLength(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HTTPURLKey = attribute.Key("http.url")

	HTTPResendCountKey = attribute.Key("http.resend_count")
)

func HTTPURL(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPResendCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HTTPSchemeKey = attribute.Key("http.scheme")

	HTTPTargetKey = attribute.Key("http.target")

	HTTPRouteKey = attribute.Key("http.route")

	HTTPClientIPKey = attribute.Key("http.client_ip")
)

func HTTPScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPTarget(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPRoute(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPClientIP(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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

func AWSDynamoDBTableNames(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBConsumedCapacity(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBItemCollectionMetrics(val string) attribute.KeyValue {
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

func AWSDynamoDBConsistentRead(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBProjection(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBLimit(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBAttributesToGet(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBIndexName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBSelect(val string) attribute.KeyValue {
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
	AWSDynamoDBSegmentKey = attribute.Key("aws.dynamodb.segment")

	AWSDynamoDBTotalSegmentsKey = attribute.Key("aws.dynamodb.total_segments")

	AWSDynamoDBCountKey = attribute.Key("aws.dynamodb.count")

	AWSDynamoDBScannedCountKey = attribute.Key("aws.dynamodb.scanned_count")
)

func AWSDynamoDBSegment(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBTotalSegments(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func AWSDynamoDBScannedCount(val int) attribute.KeyValue {
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
	GraphqlOperationNameKey = attribute.Key("graphql.operation.name")

	GraphqlOperationTypeKey = attribute.Key("graphql.operation.type")

	GraphqlDocumentKey = attribute.Key("graphql.document")
)

var (
	GraphqlOperationTypeQuery = GraphqlOperationTypeKey.String("query")

	GraphqlOperationTypeMutation = GraphqlOperationTypeKey.String("mutation")

	GraphqlOperationTypeSubscription = GraphqlOperationTypeKey.String("subscription")
)

func GraphqlOperationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func GraphqlDocument(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingMessageIDKey = attribute.Key("messaging.message.id")

	MessagingMessageConversationIDKey = attribute.Key("messaging.message.conversation_id")

	MessagingMessagePayloadSizeBytesKey = attribute.Key("messaging.message.payload_size_bytes")

	MessagingMessagePayloadCompressedSizeBytesKey = attribute.Key("messaging.message.payload_compressed_size_bytes")
)

func MessagingMessageID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingMessageConversationID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingMessagePayloadSizeBytes(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingMessagePayloadCompressedSizeBytes(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingDestinationNameKey = attribute.Key("messaging.destination.name")

	MessagingDestinationKindKey = attribute.Key("messaging.destination.kind")

	MessagingDestinationTemplateKey = attribute.Key("messaging.destination.template")

	MessagingDestinationTemporaryKey = attribute.Key("messaging.destination.temporary")

	MessagingDestinationAnonymousKey = attribute.Key("messaging.destination.anonymous")
)

var (
	MessagingDestinationKindQueue = MessagingDestinationKindKey.String("queue")

	MessagingDestinationKindTopic = MessagingDestinationKindKey.String("topic")
)

func MessagingDestinationName(val string) attribute.KeyValue {
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

func MessagingDestinationAnonymous(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingSourceNameKey = attribute.Key("messaging.source.name")

	MessagingSourceKindKey = attribute.Key("messaging.source.kind")

	MessagingSourceTemplateKey = attribute.Key("messaging.source.template")

	MessagingSourceTemporaryKey = attribute.Key("messaging.source.temporary")

	MessagingSourceAnonymousKey = attribute.Key("messaging.source.anonymous")
)

var (
	MessagingSourceKindQueue = MessagingSourceKindKey.String("queue")

	MessagingSourceKindTopic = MessagingSourceKindKey.String("topic")
)

func MessagingSourceName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingSourceTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingSourceTemporary(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingSourceAnonymous(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingSystemKey = attribute.Key("messaging.system")

	MessagingOperationKey = attribute.Key("messaging.operation")

	MessagingBatchMessageCountKey = attribute.Key("messaging.batch.message_count")
)

var (
	MessagingOperationPublish = MessagingOperationKey.String("publish")

	MessagingOperationReceive = MessagingOperationKey.String("receive")

	MessagingOperationProcess = MessagingOperationKey.String("process")
)

func MessagingSystem(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingBatchMessageCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingConsumerIDKey = attribute.Key("messaging.consumer.id")
)

func MessagingConsumerID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingRabbitmqDestinationRoutingKeyKey = attribute.Key("messaging.rabbitmq.destination.routing_key")
)

func MessagingRabbitmqDestinationRoutingKey(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingKafkaMessageKeyKey = attribute.Key("messaging.kafka.message.key")

	MessagingKafkaConsumerGroupKey = attribute.Key("messaging.kafka.consumer.group")

	MessagingKafkaClientIDKey = attribute.Key("messaging.kafka.client_id")

	MessagingKafkaDestinationPartitionKey = attribute.Key("messaging.kafka.destination.partition")

	MessagingKafkaSourcePartitionKey = attribute.Key("messaging.kafka.source.partition")

	MessagingKafkaMessageOffsetKey = attribute.Key("messaging.kafka.message.offset")

	MessagingKafkaMessageTombstoneKey = attribute.Key("messaging.kafka.message.tombstone")
)

func MessagingKafkaMessageKey(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaConsumerGroup(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaClientID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaDestinationPartition(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaSourcePartition(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaMessageOffset(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaMessageTombstone(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingRocketmqNamespaceKey = attribute.Key("messaging.rocketmq.namespace")

	MessagingRocketmqClientGroupKey = attribute.Key("messaging.rocketmq.client_group")

	MessagingRocketmqClientIDKey = attribute.Key("messaging.rocketmq.client_id")

	MessagingRocketmqMessageDeliveryTimestampKey = attribute.Key("messaging.rocketmq.message.delivery_timestamp")

	MessagingRocketmqMessageDelayTimeLevelKey = attribute.Key("messaging.rocketmq.message.delay_time_level")

	MessagingRocketmqMessageGroupKey = attribute.Key("messaging.rocketmq.message.group")

	MessagingRocketmqMessageTypeKey = attribute.Key("messaging.rocketmq.message.type")

	MessagingRocketmqMessageTagKey = attribute.Key("messaging.rocketmq.message.tag")

	MessagingRocketmqMessageKeysKey = attribute.Key("messaging.rocketmq.message.keys")

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

func MessagingRocketmqNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqClientGroup(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqClientID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageDeliveryTimestamp(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageDelayTimeLevel(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageGroup(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageTag(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageKeys(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	RPCSystemKey = attribute.Key("rpc.system")

	RPCServiceKey = attribute.Key("rpc.service")

	RPCMethodKey = attribute.Key("rpc.method")
)

var (
	RPCSystemGRPC = RPCSystemKey.String("grpc")

	RPCSystemJavaRmi = RPCSystemKey.String("java_rmi")

	RPCSystemDotnetWcf = RPCSystemKey.String("dotnet_wcf")

	RPCSystemApacheDubbo = RPCSystemKey.String("apache_dubbo")
)

func RPCService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCMethod(val string) attribute.KeyValue {
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
	RPCJsonrpcVersionKey = attribute.Key("rpc.jsonrpc.version")

	RPCJsonrpcRequestIDKey = attribute.Key("rpc.jsonrpc.request_id")

	RPCJsonrpcErrorCodeKey = attribute.Key("rpc.jsonrpc.error_code")

	RPCJsonrpcErrorMessageKey = attribute.Key("rpc.jsonrpc.error_message")
)

func RPCJsonrpcVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJsonrpcRequestID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJsonrpcErrorCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCJsonrpcErrorMessage(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
