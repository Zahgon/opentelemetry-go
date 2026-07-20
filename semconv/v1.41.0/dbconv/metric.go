package dbconv

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type ClientConnectionStateAttr string

var (
	ClientConnectionStateIdle ClientConnectionStateAttr = "idle"

	ClientConnectionStateUsed ClientConnectionStateAttr = "used"
)

type SystemNameAttr string

var (
	SystemNameOtherSQL SystemNameAttr = "other_sql"

	SystemNameSoftwareagAdabas SystemNameAttr = "softwareag.adabas"

	SystemNameActianIngres SystemNameAttr = "actian.ingres"

	SystemNameAWSDynamoDB SystemNameAttr = "aws.dynamodb"

	SystemNameAWSRedshift SystemNameAttr = "aws.redshift"

	SystemNameAzureCosmosDB SystemNameAttr = "azure.cosmosdb"

	SystemNameIntersystemsCache SystemNameAttr = "intersystems.cache"

	SystemNameCassandra SystemNameAttr = "cassandra"

	SystemNameClickHouse SystemNameAttr = "clickhouse"

	SystemNameCockroachDB SystemNameAttr = "cockroachdb"

	SystemNameCouchbase SystemNameAttr = "couchbase"

	SystemNameCouchDB SystemNameAttr = "couchdb"

	SystemNameDerby SystemNameAttr = "derby"

	SystemNameElasticsearch SystemNameAttr = "elasticsearch"

	SystemNameFirebirdSQL SystemNameAttr = "firebirdsql"

	SystemNameGCPSpanner SystemNameAttr = "gcp.spanner"

	SystemNameGeode SystemNameAttr = "geode"

	SystemNameH2database SystemNameAttr = "h2database"

	SystemNameHBase SystemNameAttr = "hbase"

	SystemNameHive SystemNameAttr = "hive"

	SystemNameHSQLDB SystemNameAttr = "hsqldb"

	SystemNameIBMDB2 SystemNameAttr = "ibm.db2"

	SystemNameIBMInformix SystemNameAttr = "ibm.informix"

	SystemNameIBMNetezza SystemNameAttr = "ibm.netezza"

	SystemNameInfluxDB SystemNameAttr = "influxdb"

	SystemNameInstantDB SystemNameAttr = "instantdb"

	SystemNameMariaDB SystemNameAttr = "mariadb"

	SystemNameMemcached SystemNameAttr = "memcached"

	SystemNameMongoDB SystemNameAttr = "mongodb"

	SystemNameMicrosoftSQLServer SystemNameAttr = "microsoft.sql_server"

	SystemNameMySQL SystemNameAttr = "mysql"

	SystemNameNeo4j SystemNameAttr = "neo4j"

	SystemNameOpenSearch SystemNameAttr = "opensearch"

	SystemNameOracleDB SystemNameAttr = "oracle.db"

	SystemNamePostgreSQL SystemNameAttr = "postgresql"

	SystemNameRedis SystemNameAttr = "redis"

	SystemNameSAPHANA SystemNameAttr = "sap.hana"

	SystemNameSAPMaxDB SystemNameAttr = "sap.maxdb"

	SystemNameSQLite SystemNameAttr = "sqlite"

	SystemNameTeradata SystemNameAttr = "teradata"

	SystemNameTrino SystemNameAttr = "trino"
)

type ErrorTypeAttr string

var (
	ErrorTypeOther ErrorTypeAttr = "_OTHER"
)

type ClientConnectionCount struct {
	metric.Int64UpDownCounter
}

var newClientConnectionCountOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of connections that are currently in state described by the `state` attribute."),
	metric.WithUnit("{connection}"),
}

func NewClientConnectionCount(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClientConnectionCount, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionCount), nil
}

func (m ClientConnectionCount) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClientConnectionCount) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionCount) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionCount) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionCount) Add(
	ctx context.Context,
	incr int64,
	clientConnectionPoolName string,
	clientConnectionState ClientConnectionStateAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConnectionCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientConnectionCountObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClientConnectionCountObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of connections that are currently in state described by the `state` attribute."),
	metric.WithUnit("{connection}"),
}

func NewClientConnectionCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClientConnectionCountObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionCountObservable), nil
}

func (m ClientConnectionCountObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClientConnectionCountObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionCountObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionCountObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionCountObservable) AttrClientConnectionPoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConnectionCountObservable) AttrClientConnectionState(val ClientConnectionStateAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientConnectionCreateTime struct {
	metric.Float64Histogram
}

var newClientConnectionCreateTimeOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The time it took to create a new connection."),
	metric.WithUnit("s"),
}

func NewClientConnectionCreateTime(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientConnectionCreateTime, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionCreateTime), nil
}

func (m ClientConnectionCreateTime) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientConnectionCreateTime) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionCreateTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionCreateTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionCreateTime) Record(
	ctx context.Context,
	val float64,
	clientConnectionPoolName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConnectionCreateTime) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientConnectionIdleMax struct {
	metric.Int64UpDownCounter
}

var newClientConnectionIdleMaxOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The maximum number of idle open connections allowed."),
	metric.WithUnit("{connection}"),
}

func NewClientConnectionIdleMax(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClientConnectionIdleMax, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionIdleMax), nil
}

func (m ClientConnectionIdleMax) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClientConnectionIdleMax) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMax) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMax) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionIdleMax) Add(
	ctx context.Context,
	incr int64,
	clientConnectionPoolName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConnectionIdleMax) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientConnectionIdleMaxObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClientConnectionIdleMaxObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The maximum number of idle open connections allowed."),
	metric.WithUnit("{connection}"),
}

func NewClientConnectionIdleMaxObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClientConnectionIdleMaxObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionIdleMaxObservable), nil
}

func (m ClientConnectionIdleMaxObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClientConnectionIdleMaxObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMaxObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMaxObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMaxObservable) AttrClientConnectionPoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientConnectionIdleMin struct {
	metric.Int64UpDownCounter
}

var newClientConnectionIdleMinOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The minimum number of idle open connections allowed."),
	metric.WithUnit("{connection}"),
}

func NewClientConnectionIdleMin(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClientConnectionIdleMin, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionIdleMin), nil
}

func (m ClientConnectionIdleMin) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClientConnectionIdleMin) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMin) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMin) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionIdleMin) Add(
	ctx context.Context,
	incr int64,
	clientConnectionPoolName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConnectionIdleMin) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientConnectionIdleMinObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClientConnectionIdleMinObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The minimum number of idle open connections allowed."),
	metric.WithUnit("{connection}"),
}

func NewClientConnectionIdleMinObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClientConnectionIdleMinObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionIdleMinObservable), nil
}

func (m ClientConnectionIdleMinObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClientConnectionIdleMinObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMinObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMinObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionIdleMinObservable) AttrClientConnectionPoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientConnectionMax struct {
	metric.Int64UpDownCounter
}

var newClientConnectionMaxOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The maximum number of open connections allowed."),
	metric.WithUnit("{connection}"),
}

func NewClientConnectionMax(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClientConnectionMax, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionMax), nil
}

func (m ClientConnectionMax) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClientConnectionMax) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionMax) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionMax) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionMax) Add(
	ctx context.Context,
	incr int64,
	clientConnectionPoolName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConnectionMax) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientConnectionMaxObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClientConnectionMaxObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The maximum number of open connections allowed."),
	metric.WithUnit("{connection}"),
}

func NewClientConnectionMaxObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClientConnectionMaxObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionMaxObservable), nil
}

func (m ClientConnectionMaxObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClientConnectionMaxObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionMaxObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionMaxObservable) Description() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionMaxObservable) AttrClientConnectionPoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientConnectionPendingRequests struct {
	metric.Int64UpDownCounter
}

var newClientConnectionPendingRequestsOpts = []metric.Int64UpDownCounterOption{
	metric.WithDescription("The number of current pending requests for an open connection."),
	metric.WithUnit("{request}"),
}

func NewClientConnectionPendingRequests(
	m metric.Meter,
	opt ...metric.Int64UpDownCounterOption,
) (ClientConnectionPendingRequests, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionPendingRequests), nil
}

func (m ClientConnectionPendingRequests) Inst() metric.Int64UpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64UpDownCounter)
}

func (ClientConnectionPendingRequests) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionPendingRequests) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionPendingRequests) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionPendingRequests) Add(
	ctx context.Context,
	incr int64,
	clientConnectionPoolName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConnectionPendingRequests) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientConnectionPendingRequestsObservable struct {
	metric.Int64ObservableUpDownCounter
}

var newClientConnectionPendingRequestsObservableOpts = []metric.Int64ObservableUpDownCounterOption{
	metric.WithDescription("The number of current pending requests for an open connection."),
	metric.WithUnit("{request}"),
}

func NewClientConnectionPendingRequestsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableUpDownCounterOption,
) (ClientConnectionPendingRequestsObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionPendingRequestsObservable), nil
}

func (m ClientConnectionPendingRequestsObservable) Inst() metric.Int64ObservableUpDownCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableUpDownCounter)
}

func (ClientConnectionPendingRequestsObservable) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClientConnectionPendingRequestsObservable) Unit() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClientConnectionPendingRequestsObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClientConnectionPendingRequestsObservable) AttrClientConnectionPoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientConnectionTimeouts struct {
	metric.Int64Counter
}

var newClientConnectionTimeoutsOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of connection timeouts that have occurred trying to obtain a connection from the pool."),
	metric.WithUnit("{timeout}"),
}

func NewClientConnectionTimeouts(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientConnectionTimeouts, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionTimeouts), nil
}

func (m ClientConnectionTimeouts) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientConnectionTimeouts) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionTimeouts) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionTimeouts) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionTimeouts) Add(
	ctx context.Context,
	incr int64,
	clientConnectionPoolName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConnectionTimeouts) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientConnectionTimeoutsObservable struct {
	metric.Int64ObservableCounter
}

var newClientConnectionTimeoutsObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of connection timeouts that have occurred trying to obtain a connection from the pool."),
	metric.WithUnit("{timeout}"),
}

func NewClientConnectionTimeoutsObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ClientConnectionTimeoutsObservable, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionTimeoutsObservable), nil
}

func (m ClientConnectionTimeoutsObservable) Inst() metric.Int64ObservableCounter {
	_ = "STUB: not implemented"
	return *new(metric.Int64ObservableCounter)
}

func (ClientConnectionTimeoutsObservable) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionTimeoutsObservable) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionTimeoutsObservable) Description() string {
	_ = "STUB: not implemented"
	return ""
}

func (ClientConnectionTimeoutsObservable) AttrClientConnectionPoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientConnectionUseTime struct {
	metric.Float64Histogram
}

var newClientConnectionUseTimeOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The time between borrowing a connection and returning it to the pool."),
	metric.WithUnit("s"),
}

func NewClientConnectionUseTime(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientConnectionUseTime, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionUseTime), nil
}

func (m ClientConnectionUseTime) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientConnectionUseTime) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionUseTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionUseTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionUseTime) Record(
	ctx context.Context,
	val float64,
	clientConnectionPoolName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConnectionUseTime) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientConnectionWaitTime struct {
	metric.Float64Histogram
}

var newClientConnectionWaitTimeOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The time it took to obtain an open connection from the pool."),
	metric.WithUnit("s"),
}

func NewClientConnectionWaitTime(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientConnectionWaitTime, error) {
	_ = "STUB: not implemented"
	return *new(ClientConnectionWaitTime), nil
}

func (m ClientConnectionWaitTime) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientConnectionWaitTime) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionWaitTime) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConnectionWaitTime) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConnectionWaitTime) Record(
	ctx context.Context,
	val float64,
	clientConnectionPoolName string,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConnectionWaitTime) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

type ClientOperationDuration struct {
	metric.Float64Histogram
}

var newClientOperationDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Duration of database client operations."),
	metric.WithUnit("s"),
}

func NewClientOperationDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientOperationDuration, error) {
	_ = "STUB: not implemented"
	return *new(ClientOperationDuration), nil
}

func (m ClientOperationDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientOperationDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientOperationDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientOperationDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientOperationDuration) Record(
	ctx context.Context,
	val float64,
	systemName SystemNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientOperationDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientOperationDuration) AttrCollectionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrOperationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrQuerySummary(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrStoredProcedureName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrNetworkPeerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrNetworkPeerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrQueryText(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientResponseReturnedRows struct {
	metric.Int64Histogram
}

var newClientResponseReturnedRowsOpts = []metric.Int64HistogramOption{
	metric.WithDescription("The actual number of records returned by the database operation."),
	metric.WithUnit("{row}"),
}

func NewClientResponseReturnedRows(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientResponseReturnedRows, error) {
	_ = "STUB: not implemented"
	return *new(ClientResponseReturnedRows), nil
}

func (m ClientResponseReturnedRows) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientResponseReturnedRows) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientResponseReturnedRows) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientResponseReturnedRows) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientResponseReturnedRows) Record(
	ctx context.Context,
	val int64,
	systemName SystemNameAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientResponseReturnedRows) RecordSet(ctx context.Context, val int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientResponseReturnedRows) AttrCollectionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrOperationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrResponseStatusCode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrQuerySummary(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrNetworkPeerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrNetworkPeerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientResponseReturnedRows) AttrQueryText(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
