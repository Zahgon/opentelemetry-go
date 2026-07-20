package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	ClientAddressKey = attribute.Key("client.address")

	ClientPortKey = attribute.Key("client.port")

	ClientSocketAddressKey = attribute.Key("client.socket.address")

	ClientSocketPortKey = attribute.Key("client.socket.port")
)

func ClientAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ClientPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ClientSocketAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ClientSocketPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HTTPMethodKey = attribute.Key("http.method")

	HTTPStatusCodeKey = attribute.Key("http.status_code")

	HTTPSchemeKey = attribute.Key("http.scheme")

	HTTPURLKey = attribute.Key("http.url")

	HTTPTargetKey = attribute.Key("http.target")

	HTTPRequestContentLengthKey = attribute.Key("http.request_content_length")

	HTTPResponseContentLengthKey = attribute.Key("http.response_content_length")
)

func HTTPMethod(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPURL(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPTarget(val string) attribute.KeyValue {
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
	NetSockPeerNameKey = attribute.Key("net.sock.peer.name")

	NetSockPeerAddrKey = attribute.Key("net.sock.peer.addr")

	NetSockPeerPortKey = attribute.Key("net.sock.peer.port")

	NetPeerNameKey = attribute.Key("net.peer.name")

	NetPeerPortKey = attribute.Key("net.peer.port")

	NetHostNameKey = attribute.Key("net.host.name")

	NetHostPortKey = attribute.Key("net.host.port")

	NetSockHostAddrKey = attribute.Key("net.sock.host.addr")

	NetSockHostPortKey = attribute.Key("net.sock.host.port")

	NetTransportKey = attribute.Key("net.transport")

	NetProtocolNameKey = attribute.Key("net.protocol.name")

	NetProtocolVersionKey = attribute.Key("net.protocol.version")

	NetSockFamilyKey = attribute.Key("net.sock.family")
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

func NetProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	DestinationDomainKey = attribute.Key("destination.domain")

	DestinationAddressKey = attribute.Key("destination.address")

	DestinationPortKey = attribute.Key("destination.port")
)

func DestinationDomain(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DestinationAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func DestinationPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HTTPRequestMethodKey = attribute.Key("http.request.method")

	HTTPResponseStatusCodeKey = attribute.Key("http.response.status_code")
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

func HTTPResponseStatusCode(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HTTPRouteKey = attribute.Key("http.route")
)

func HTTPRoute(val string) attribute.KeyValue {
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

	LogFilePathKey = attribute.Key("log.file.path")

	LogFileNameResolvedKey = attribute.Key("log.file.name_resolved")

	LogFilePathResolvedKey = attribute.Key("log.file.path_resolved")
)

func LogFileName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func LogFilePath(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func LogFileNameResolved(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func LogFilePathResolved(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	TypeKey = attribute.Key("type")

	PoolKey = attribute.Key("pool")
)

var (
	TypeHeap = TypeKey.String("heap")

	TypeNonHeap = TypeKey.String("non_heap")
)

func Pool(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	ServerAddressKey = attribute.Key("server.address")

	ServerPortKey = attribute.Key("server.port")

	ServerSocketDomainKey = attribute.Key("server.socket.domain")

	ServerSocketAddressKey = attribute.Key("server.socket.address")

	ServerSocketPortKey = attribute.Key("server.socket.port")
)

func ServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServerSocketDomain(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServerSocketAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func ServerSocketPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	SourceDomainKey = attribute.Key("source.domain")

	SourceAddressKey = attribute.Key("source.address")

	SourcePortKey = attribute.Key("source.port")
)

func SourceDomain(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SourceAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func SourcePort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	NetworkTransportKey = attribute.Key("network.transport")

	NetworkTypeKey = attribute.Key("network.type")

	NetworkProtocolNameKey = attribute.Key("network.protocol.name")

	NetworkProtocolVersionKey = attribute.Key("network.protocol.version")
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

func NetworkProtocolName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func NetworkProtocolVersion(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	NetworkConnectionTypeKey = attribute.Key("network.connection.type")

	NetworkConnectionSubtypeKey = attribute.Key("network.connection.subtype")

	NetworkCarrierNameKey = attribute.Key("network.carrier.name")

	NetworkCarrierMccKey = attribute.Key("network.carrier.mcc")

	NetworkCarrierMncKey = attribute.Key("network.carrier.mnc")

	NetworkCarrierIccKey = attribute.Key("network.carrier.icc")
)

var (
	NetworkConnectionTypeWifi = NetworkConnectionTypeKey.String("wifi")

	NetworkConnectionTypeWired = NetworkConnectionTypeKey.String("wired")

	NetworkConnectionTypeCell = NetworkConnectionTypeKey.String("cell")

	NetworkConnectionTypeUnavailable = NetworkConnectionTypeKey.String("unavailable")

	NetworkConnectionTypeUnknown = NetworkConnectionTypeKey.String("unknown")
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

func NetworkCarrierName(val string) attribute.KeyValue {
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

func NetworkCarrierIcc(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	HTTPRequestMethodOriginalKey = attribute.Key("http.request.method_original")

	HTTPRequestBodySizeKey = attribute.Key("http.request.body.size")

	HTTPResponseBodySizeKey = attribute.Key("http.response.body.size")
)

func HTTPRequestMethodOriginal(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPRequestBodySize(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPResponseBodySize(val int) attribute.KeyValue {
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

	MessagingDestinationTemplateKey = attribute.Key("messaging.destination.template")

	MessagingDestinationTemporaryKey = attribute.Key("messaging.destination.temporary")

	MessagingDestinationAnonymousKey = attribute.Key("messaging.destination.anonymous")
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
	MessagingRabbitmqDestinationRoutingKeyKey = attribute.Key("messaging.rabbitmq.destination.routing_key")
)

func MessagingRabbitmqDestinationRoutingKey(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

const (
	MessagingKafkaMessageKeyKey = attribute.Key("messaging.kafka.message.key")

	MessagingKafkaConsumerGroupKey = attribute.Key("messaging.kafka.consumer.group")

	MessagingKafkaDestinationPartitionKey = attribute.Key("messaging.kafka.destination.partition")

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

func MessagingKafkaDestinationPartition(val int) attribute.KeyValue {
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
	URLSchemeKey = attribute.Key("url.scheme")

	URLFullKey = attribute.Key("url.full")

	URLPathKey = attribute.Key("url.path")

	URLQueryKey = attribute.Key("url.query")

	URLFragmentKey = attribute.Key("url.fragment")
)

func URLScheme(val string) attribute.KeyValue {
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

func URLFragment(val string) attribute.KeyValue {
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
