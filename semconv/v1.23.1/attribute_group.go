package semconv

import "go.opentelemetry.io/otel/attribute"

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
	ErrorTypeKey = attribute.Key("error.type")
)

var ErrorTypeOther = ErrorTypeKey.String("_OTHER")

const (
	FaaSInvokedNameKey = attribute.Key("faas.invoked_name")

	FaaSInvokedProviderKey = attribute.Key("faas.invoked_provider")

	FaaSInvokedRegionKey = attribute.Key("faas.invoked_region")

	FaaSTriggerKey = attribute.Key("faas.trigger")
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

	FaaSTriggerPubsub = FaaSTriggerKey.String("pubsub")

	FaaSTriggerTimer = FaaSTriggerKey.String("timer")

	FaaSTriggerOther = FaaSTriggerKey.String("other")
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
	EventDomainKey = attribute.Key("event.domain")

	EventNameKey = attribute.Key("event.name")
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
	LogRecordUIDKey = attribute.Key("log.record.uid")
)

func LogRecordUID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	LogIostreamKey = attribute.Key("log.iostream")
)

var (
	LogIostreamStdout = LogIostreamKey.String("stdout")

	LogIostreamStderr = LogIostreamKey.String("stderr")
)

const (
	LogFileNameKey = attribute.Key("log.file.name")

	LogFileNameResolvedKey = attribute.Key("log.file.name_resolved")

	LogFilePathKey = attribute.Key("log.file.path")

	LogFilePathResolvedKey = attribute.Key("log.file.path_resolved")
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

const (
	PoolNameKey = attribute.Key("pool.name")

	StateKey = attribute.Key("state")
)

var (
	StateIdle = StateKey.String("idle")

	StateUsed = StateKey.String("used")
)

func PoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	JvmBufferPoolNameKey = attribute.Key("jvm.buffer.pool.name")
)

func JvmBufferPoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	JvmMemoryPoolNameKey = attribute.Key("jvm.memory.pool.name")

	JvmMemoryTypeKey = attribute.Key("jvm.memory.type")
)

var (
	JvmMemoryTypeHeap = JvmMemoryTypeKey.String("heap")

	JvmMemoryTypeNonHeap = JvmMemoryTypeKey.String("non_heap")
)

func JvmMemoryPoolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	SystemDeviceKey = attribute.Key("system.device")
)

func SystemDevice(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	SystemCPULogicalNumberKey = attribute.Key("system.cpu.logical_number")

	SystemCPUStateKey = attribute.Key("system.cpu.state")
)

var (
	SystemCPUStateUser = SystemCPUStateKey.String("user")

	SystemCPUStateSystem = SystemCPUStateKey.String("system")

	SystemCPUStateNice = SystemCPUStateKey.String("nice")

	SystemCPUStateIdle = SystemCPUStateKey.String("idle")

	SystemCPUStateIowait = SystemCPUStateKey.String("iowait")

	SystemCPUStateInterrupt = SystemCPUStateKey.String("interrupt")

	SystemCPUStateSteal = SystemCPUStateKey.String("steal")
)

func SystemCPULogicalNumber(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	SystemMemoryStateKey = attribute.Key("system.memory.state")
)

var (
	SystemMemoryStateUsed = SystemMemoryStateKey.String("used")

	SystemMemoryStateFree = SystemMemoryStateKey.String("free")

	SystemMemoryStateShared = SystemMemoryStateKey.String("shared")

	SystemMemoryStateBuffers = SystemMemoryStateKey.String("buffers")

	SystemMemoryStateCached = SystemMemoryStateKey.String("cached")
)

const (
	SystemPagingDirectionKey = attribute.Key("system.paging.direction")

	SystemPagingStateKey = attribute.Key("system.paging.state")

	SystemPagingTypeKey = attribute.Key("system.paging.type")
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

const (
	SystemDiskDirectionKey = attribute.Key("system.disk.direction")
)

var (
	SystemDiskDirectionRead = SystemDiskDirectionKey.String("read")

	SystemDiskDirectionWrite = SystemDiskDirectionKey.String("write")
)

const (
	SystemFilesystemModeKey = attribute.Key("system.filesystem.mode")

	SystemFilesystemMountpointKey = attribute.Key("system.filesystem.mountpoint")

	SystemFilesystemStateKey = attribute.Key("system.filesystem.state")

	SystemFilesystemTypeKey = attribute.Key("system.filesystem.type")
)

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

func SystemFilesystemMode(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SystemFilesystemMountpoint(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	SystemNetworkDirectionKey = attribute.Key("system.network.direction")

	SystemNetworkStateKey = attribute.Key("system.network.state")
)

var (
	SystemNetworkDirectionTransmit = SystemNetworkDirectionKey.String("transmit")

	SystemNetworkDirectionReceive = SystemNetworkDirectionKey.String("receive")
)

var (
	SystemNetworkStateClose = SystemNetworkStateKey.String("close")

	SystemNetworkStateCloseWait = SystemNetworkStateKey.String("close_wait")

	SystemNetworkStateClosing = SystemNetworkStateKey.String("closing")

	SystemNetworkStateDelete = SystemNetworkStateKey.String("delete")

	SystemNetworkStateEstablished = SystemNetworkStateKey.String("established")

	SystemNetworkStateFinWait1 = SystemNetworkStateKey.String("fin_wait_1")

	SystemNetworkStateFinWait2 = SystemNetworkStateKey.String("fin_wait_2")

	SystemNetworkStateLastAck = SystemNetworkStateKey.String("last_ack")

	SystemNetworkStateListen = SystemNetworkStateKey.String("listen")

	SystemNetworkStateSynRecv = SystemNetworkStateKey.String("syn_recv")

	SystemNetworkStateSynSent = SystemNetworkStateKey.String("syn_sent")

	SystemNetworkStateTimeWait = SystemNetworkStateKey.String("time_wait")
)

const (
	SystemProcessesStatusKey = attribute.Key("system.processes.status")
)

var (
	SystemProcessesStatusRunning = SystemProcessesStatusKey.String("running")

	SystemProcessesStatusSleeping = SystemProcessesStatusKey.String("sleeping")

	SystemProcessesStatusStopped = SystemProcessesStatusKey.String("stopped")

	SystemProcessesStatusDefunct = SystemProcessesStatusKey.String("defunct")
)

const (
	HTTPMethodKey = attribute.Key("http.method")

	HTTPRequestContentLengthKey = attribute.Key("http.request_content_length")

	HTTPResponseContentLengthKey = attribute.Key("http.response_content_length")

	HTTPSchemeKey = attribute.Key("http.scheme")

	HTTPStatusCodeKey = attribute.Key("http.status_code")

	HTTPTargetKey = attribute.Key("http.target")

	HTTPURLKey = attribute.Key("http.url")
)

func HTTPMethod(val string) attribute.KeyValue {
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

func HTTPScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPTarget(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPURL(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	NetHostNameKey = attribute.Key("net.host.name")

	NetHostPortKey = attribute.Key("net.host.port")

	NetPeerNameKey = attribute.Key("net.peer.name")

	NetPeerPortKey = attribute.Key("net.peer.port")

	NetProtocolNameKey = attribute.Key("net.protocol.name")

	NetProtocolVersionKey = attribute.Key("net.protocol.version")

	NetSockFamilyKey = attribute.Key("net.sock.family")

	NetSockHostAddrKey = attribute.Key("net.sock.host.addr")

	NetSockHostPortKey = attribute.Key("net.sock.host.port")

	NetSockPeerAddrKey = attribute.Key("net.sock.peer.addr")

	NetSockPeerNameKey = attribute.Key("net.sock.peer.name")

	NetSockPeerPortKey = attribute.Key("net.sock.peer.port")

	NetTransportKey = attribute.Key("net.transport")
)

var (
	NetSockFamilyInet = NetSockFamilyKey.String("inet")

	NetSockFamilyInet6 = NetSockFamilyKey.String("inet6")

	NetSockFamilyUnix = NetSockFamilyKey.String("unix")
)

var (
	NetTransportTCP = NetTransportKey.String("ip_tcp")

	NetTransportUDP = NetTransportKey.String("ip_udp")

	NetTransportPipe = NetTransportKey.String("pipe")

	NetTransportInProc = NetTransportKey.String("inproc")

	NetTransportOther = NetTransportKey.String("other")
)

func NetHostName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetHostPort(val int) attribute.KeyValue {
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

func NetProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetProtocolVersion(val string) attribute.KeyValue {
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

func NetSockPeerAddr(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetSockPeerName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetSockPeerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HTTPRequestBodySizeKey = attribute.Key("http.request.body.size")

	HTTPRequestMethodKey = attribute.Key("http.request.method")

	HTTPRequestMethodOriginalKey = attribute.Key("http.request.method_original")

	HTTPRequestResendCountKey = attribute.Key("http.request.resend_count")

	HTTPResponseBodySizeKey = attribute.Key("http.response.body.size")

	HTTPResponseStatusCodeKey = attribute.Key("http.response.status_code")

	HTTPRouteKey = attribute.Key("http.route")
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

func HTTPRequestBodySize(val int) attribute.KeyValue {
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

func HTTPResponseBodySize(val int) attribute.KeyValue {
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

const (
	MessagingBatchMessageCountKey = attribute.Key("messaging.batch.message_count")

	MessagingClientIDKey = attribute.Key("messaging.client_id")

	MessagingDestinationAnonymousKey = attribute.Key("messaging.destination.anonymous")

	MessagingDestinationNameKey = attribute.Key("messaging.destination.name")

	MessagingDestinationTemplateKey = attribute.Key("messaging.destination.template")

	MessagingDestinationTemporaryKey = attribute.Key("messaging.destination.temporary")

	MessagingDestinationPublishAnonymousKey = attribute.Key("messaging.destination_publish.anonymous")

	MessagingDestinationPublishNameKey = attribute.Key("messaging.destination_publish.name")

	MessagingKafkaConsumerGroupKey = attribute.Key("messaging.kafka.consumer.group")

	MessagingKafkaDestinationPartitionKey = attribute.Key("messaging.kafka.destination.partition")

	MessagingKafkaMessageKeyKey = attribute.Key("messaging.kafka.message.key")

	MessagingKafkaMessageOffsetKey = attribute.Key("messaging.kafka.message.offset")

	MessagingKafkaMessageTombstoneKey = attribute.Key("messaging.kafka.message.tombstone")

	MessagingMessageBodySizeKey = attribute.Key("messaging.message.body.size")

	MessagingMessageConversationIDKey = attribute.Key("messaging.message.conversation_id")

	MessagingMessageEnvelopeSizeKey = attribute.Key("messaging.message.envelope.size")

	MessagingMessageIDKey = attribute.Key("messaging.message.id")

	MessagingOperationKey = attribute.Key("messaging.operation")

	MessagingRabbitmqDestinationRoutingKeyKey = attribute.Key("messaging.rabbitmq.destination.routing_key")

	MessagingRocketmqClientGroupKey = attribute.Key("messaging.rocketmq.client_group")

	MessagingRocketmqConsumptionModelKey = attribute.Key("messaging.rocketmq.consumption_model")

	MessagingRocketmqMessageDelayTimeLevelKey = attribute.Key("messaging.rocketmq.message.delay_time_level")

	MessagingRocketmqMessageDeliveryTimestampKey = attribute.Key("messaging.rocketmq.message.delivery_timestamp")

	MessagingRocketmqMessageGroupKey = attribute.Key("messaging.rocketmq.message.group")

	MessagingRocketmqMessageKeysKey = attribute.Key("messaging.rocketmq.message.keys")

	MessagingRocketmqMessageTagKey = attribute.Key("messaging.rocketmq.message.tag")

	MessagingRocketmqMessageTypeKey = attribute.Key("messaging.rocketmq.message.type")

	MessagingRocketmqNamespaceKey = attribute.Key("messaging.rocketmq.namespace")

	MessagingSystemKey = attribute.Key("messaging.system")
)

var (
	MessagingOperationPublish = MessagingOperationKey.String("publish")

	MessagingOperationCreate = MessagingOperationKey.String("create")

	MessagingOperationReceive = MessagingOperationKey.String("receive")

	MessagingOperationDeliver = MessagingOperationKey.String("deliver")
)

var (
	MessagingRocketmqConsumptionModelClustering = MessagingRocketmqConsumptionModelKey.String("clustering")

	MessagingRocketmqConsumptionModelBroadcasting = MessagingRocketmqConsumptionModelKey.String("broadcasting")
)

var (
	MessagingRocketmqMessageTypeNormal = MessagingRocketmqMessageTypeKey.String("normal")

	MessagingRocketmqMessageTypeFifo = MessagingRocketmqMessageTypeKey.String("fifo")

	MessagingRocketmqMessageTypeDelay = MessagingRocketmqMessageTypeKey.String("delay")

	MessagingRocketmqMessageTypeTransaction = MessagingRocketmqMessageTypeKey.String("transaction")
)

func MessagingBatchMessageCount(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingClientID(val string) attribute.KeyValue {
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

func MessagingDestinationTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingDestinationTemporary(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingDestinationPublishAnonymous(val bool) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingDestinationPublishName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaConsumerGroup(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaDestinationPartition(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingKafkaMessageKey(val string) attribute.KeyValue {
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

func MessagingRabbitmqDestinationRoutingKey(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqClientGroup(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageDelayTimeLevel(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageDeliveryTimestamp(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageGroup(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageKeys(val ...string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqMessageTag(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingRocketmqNamespace(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func MessagingSystem(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	NetworkCarrierIccKey = attribute.Key("network.carrier.icc")

	NetworkCarrierMccKey = attribute.Key("network.carrier.mcc")

	NetworkCarrierMncKey = attribute.Key("network.carrier.mnc")

	NetworkCarrierNameKey = attribute.Key("network.carrier.name")

	NetworkConnectionSubtypeKey = attribute.Key("network.connection.subtype")

	NetworkConnectionTypeKey = attribute.Key("network.connection.type")

	NetworkLocalAddressKey = attribute.Key("network.local.address")

	NetworkLocalPortKey = attribute.Key("network.local.port")

	NetworkPeerAddressKey = attribute.Key("network.peer.address")

	NetworkPeerPortKey = attribute.Key("network.peer.port")

	NetworkProtocolNameKey = attribute.Key("network.protocol.name")

	NetworkProtocolVersionKey = attribute.Key("network.protocol.version")

	NetworkTransportKey = attribute.Key("network.transport")

	NetworkTypeKey = attribute.Key("network.type")
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
	NetworkTransportTCP = NetworkTransportKey.String("tcp")

	NetworkTransportUDP = NetworkTransportKey.String("udp")

	NetworkTransportPipe = NetworkTransportKey.String("pipe")

	NetworkTransportUnix = NetworkTransportKey.String("unix")
)

var (
	NetworkTypeIpv4 = NetworkTypeKey.String("ipv4")

	NetworkTypeIpv6 = NetworkTypeKey.String("ipv6")
)

func NetworkCarrierIcc(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkCarrierMcc(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkCarrierMnc(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkCarrierName(val string) attribute.KeyValue {
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

const (
	RPCConnectRPCErrorCodeKey = attribute.Key("rpc.connect_rpc.error_code")

	RPCGRPCStatusCodeKey = attribute.Key("rpc.grpc.status_code")

	RPCJsonrpcErrorCodeKey = attribute.Key("rpc.jsonrpc.error_code")

	RPCJsonrpcErrorMessageKey = attribute.Key("rpc.jsonrpc.error_message")

	RPCJsonrpcRequestIDKey = attribute.Key("rpc.jsonrpc.request_id")

	RPCJsonrpcVersionKey = attribute.Key("rpc.jsonrpc.version")

	RPCMethodKey = attribute.Key("rpc.method")

	RPCServiceKey = attribute.Key("rpc.service")

	RPCSystemKey = attribute.Key("rpc.system")
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
	RPCSystemGRPC = RPCSystemKey.String("grpc")

	RPCSystemJavaRmi = RPCSystemKey.String("java_rmi")

	RPCSystemDotnetWcf = RPCSystemKey.String("dotnet_wcf")

	RPCSystemApacheDubbo = RPCSystemKey.String("apache_dubbo")

	RPCSystemConnectRPC = RPCSystemKey.String("connect_rpc")
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

func RPCMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func RPCService(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	URLFragmentKey = attribute.Key("url.fragment")

	URLFullKey = attribute.Key("url.full")

	URLPathKey = attribute.Key("url.path")

	URLQueryKey = attribute.Key("url.query")

	URLSchemeKey = attribute.Key("url.scheme")
)

func URLFragment(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLFull(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLPath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLQuery(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func URLScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	UserAgentOriginalKey = attribute.Key("user_agent.original")
)

func UserAgentOriginal(val string) attribute.KeyValue {
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
