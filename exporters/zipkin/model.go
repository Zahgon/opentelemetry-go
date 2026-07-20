package zipkin

import (
	zkmodel "github.com/openzipkin/zipkin-go/model"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv120 "go.opentelemetry.io/otel/semconv/v1.20.0"
	semconv121 "go.opentelemetry.io/otel/semconv/v1.21.0"
	semconv125 "go.opentelemetry.io/otel/semconv/v1.25.0"
	semconv138 "go.opentelemetry.io/otel/semconv/v1.38.0"
	semconv "go.opentelemetry.io/otel/semconv/v1.42.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	keyPeerHostname attribute.Key = "peer.hostname"
	keyPeerAddress  attribute.Key = "peer.address"
)

var defaultServiceName string

func init() {

	defaultResource := resource.Default()
	if value, exists := defaultResource.Set().Value(semconv.ServiceNameKey); exists {
		defaultServiceName = value.AsString()
	}
}

func SpanModels(batch []tracesdk.ReadOnlySpan) []zkmodel.SpanModel {
	_ = "STUB: not implemented"
	return nil
}

func getServiceName(attrs []attribute.KeyValue) string { _ = "STUB: not implemented"; return "" }

func toZipkinSpanModel(data tracesdk.ReadOnlySpan) zkmodel.SpanModel {
	_ = "STUB: not implemented"
	return *new(zkmodel.SpanModel)
}

func toZipkinSpanContext(data tracesdk.ReadOnlySpan) zkmodel.SpanContext {
	_ = "STUB: not implemented"
	return *new(zkmodel.SpanContext)
}

func toZipkinTraceID(traceID trace.TraceID) zkmodel.TraceID {
	_ = "STUB: not implemented"
	return *new(zkmodel.TraceID)
}

func toZipkinID(spanID trace.SpanID) zkmodel.ID { _ = "STUB: not implemented"; return *new(zkmodel.ID) }

func toZipkinParentID(spanID trace.SpanID) *zkmodel.ID { _ = "STUB: not implemented"; return nil }

func toZipkinKind(kind trace.SpanKind) zkmodel.Kind {
	_ = "STUB: not implemented"
	return *new(zkmodel.Kind)
}

func toZipkinAnnotations(events []tracesdk.Event) []zkmodel.Annotation {
	_ = "STUB: not implemented"
	return nil
}

func attributesToJSONMapString(attributes []attribute.KeyValue) string {
	_ = "STUB: not implemented"
	return ""
}

func attributeToStringPair(kv attribute.KeyValue) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

//nolint:staticcheck // Preserve existing Zipkin tag encoding.

var extraZipkinTagsLen = len([]attribute.Key{
	semconv.OTelStatusCodeKey,
	semconv.OTelScopeNameKey,
	semconv.OTelScopeVersionKey,
})

func toZipkinTags(data tracesdk.ReadOnlySpan) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

var remoteEndpointKeyRank = map[attribute.Key]int{
	semconv138.PeerServiceKey:         1,
	semconv.ServerAddressKey:          2,
	semconv120.NetPeerNameKey:         3,
	semconv.NetworkPeerAddressKey:     4,
	semconv121.ServerSocketDomainKey:  5,
	semconv121.ServerSocketAddressKey: 6,
	semconv120.NetSockPeerNameKey:     7,
	semconv120.NetSockPeerAddrKey:     8,
	keyPeerHostname:                   9,
	keyPeerAddress:                    10,
	semconv125.DBNameKey:              11,
}

func toZipkinRemoteEndpoint(data tracesdk.ReadOnlySpan) *zkmodel.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

func remoteEndpointPeerIPWithPort(peerIP string, portKey attribute.Key, attrs []attribute.KeyValue) *zkmodel.Endpoint {
	_ = "STUB: not implemented"
	return nil
}
