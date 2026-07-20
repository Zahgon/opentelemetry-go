package semconv

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/semconv/internal"
	"go.opentelemetry.io/otel/trace"
)

var (
	HTTPSchemeHTTP  = HTTPSchemeKey.String("http")
	HTTPSchemeHTTPS = HTTPSchemeKey.String("https")
)

var sc = &internal.SemanticConventions{
	EnduserIDKey:                EnduserIDKey,
	HTTPClientIPKey:             HTTPClientIPKey,
	HTTPFlavorKey:               HTTPFlavorKey,
	HTTPHostKey:                 HTTPHostKey,
	HTTPMethodKey:               HTTPMethodKey,
	HTTPRequestContentLengthKey: HTTPRequestContentLengthKey,
	HTTPRouteKey:                HTTPRouteKey,
	HTTPSchemeHTTP:              HTTPSchemeHTTP,
	HTTPSchemeHTTPS:             HTTPSchemeHTTPS,
	HTTPServerNameKey:           HTTPServerNameKey,
	HTTPStatusCodeKey:           HTTPStatusCodeKey,
	HTTPTargetKey:               HTTPTargetKey,
	HTTPURLKey:                  HTTPURLKey,
	HTTPUserAgentKey:            HTTPUserAgentKey,
	NetHostIPKey:                NetHostIPKey,
	NetHostNameKey:              NetHostNameKey,
	NetHostPortKey:              NetHostPortKey,
	NetPeerIPKey:                NetPeerIPKey,
	NetPeerNameKey:              NetPeerNameKey,
	NetPeerPortKey:              NetPeerPortKey,
	NetTransportIP:              NetTransportIP,
	NetTransportOther:           NetTransportOther,
	NetTransportTCP:             NetTransportTCP,
	NetTransportUDP:             NetTransportUDP,
	NetTransportUnix:            NetTransportUnix,
}

func NetAttributesFromHTTPRequest(network string, request *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func EndUserAttributesFromHTTPRequest(request *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func HTTPClientAttributesFromHTTPRequest(request *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func HTTPServerMetricAttributesFromHTTPRequest(serverName string, request *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func HTTPServerAttributesFromHTTPRequest(serverName, route string, request *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func HTTPAttributesFromHTTPStatusCode(code int) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func SpanStatusFromHTTPStatusCode(code int) (codes.Code, string) {
	_ = "STUB: not implemented"
	return *new(codes.Code), ""
}

func SpanStatusFromHTTPStatusCodeAndSpanKind(code int, spanKind trace.SpanKind) (codes.Code, string) {
	_ = "STUB: not implemented"
	return *new(codes.Code), ""
}
