package semconv

import "go.opentelemetry.io/otel/attribute"

const (
	HTTPMethodKey = attribute.Key("http.method")

	HTTPStatusCodeKey = attribute.Key("http.status_code")

	HTTPFlavorKey = attribute.Key("http.flavor")
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

const (
	HTTPSchemeKey = attribute.Key("http.scheme")

	HTTPRouteKey = attribute.Key("http.route")
)

func HTTPScheme(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

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
	HTTPRequestContentLengthKey = attribute.Key("http.request_content_length")

	HTTPResponseContentLengthKey = attribute.Key("http.response_content_length")
)

func HTTPRequestContentLength(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func HTTPResponseContentLength(val int) attribute.KeyValue {
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
	UserAgentOriginalKey = attribute.Key("user_agent.original")
)

func UserAgentOriginal(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
