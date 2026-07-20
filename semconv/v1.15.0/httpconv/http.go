package httpconv

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/semconv/internal/v2"
	semconv "go.opentelemetry.io/otel/semconv/v1.15.0"
)

var (
	nc = &internal.NetConv{
		NetHostNameKey:     semconv.NetHostNameKey,
		NetHostPortKey:     semconv.NetHostPortKey,
		NetPeerNameKey:     semconv.NetPeerNameKey,
		NetPeerPortKey:     semconv.NetPeerPortKey,
		NetSockPeerAddrKey: semconv.NetSockPeerAddrKey,
		NetSockPeerPortKey: semconv.NetSockPeerPortKey,
		NetTransportOther:  semconv.NetTransportOther,
		NetTransportTCP:    semconv.NetTransportTCP,
		NetTransportUDP:    semconv.NetTransportUDP,
		NetTransportInProc: semconv.NetTransportInProc,
	}

	hc = &internal.HTTPConv{
		NetConv: nc,

		EnduserIDKey:                 semconv.EnduserIDKey,
		HTTPClientIPKey:              semconv.HTTPClientIPKey,
		HTTPFlavorKey:                semconv.HTTPFlavorKey,
		HTTPMethodKey:                semconv.HTTPMethodKey,
		HTTPRequestContentLengthKey:  semconv.HTTPRequestContentLengthKey,
		HTTPResponseContentLengthKey: semconv.HTTPResponseContentLengthKey,
		HTTPRouteKey:                 semconv.HTTPRouteKey,
		HTTPSchemeHTTP:               semconv.HTTPSchemeHTTP,
		HTTPSchemeHTTPS:              semconv.HTTPSchemeHTTPS,
		HTTPStatusCodeKey:            semconv.HTTPStatusCodeKey,
		HTTPTargetKey:                semconv.HTTPTargetKey,
		HTTPURLKey:                   semconv.HTTPURLKey,
		HTTPUserAgentKey:             semconv.HTTPUserAgentKey,
	}
)

func ClientResponse(resp *http.Response) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func ClientRequest(req *http.Request) []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func ClientStatus(code int) (codes.Code, string) {
	_ = "STUB: not implemented"
	return *new(codes.Code), ""
}

func ServerRequest(server string, req *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func ServerStatus(code int) (codes.Code, string) {
	_ = "STUB: not implemented"
	return *new(codes.Code), ""
}

func RequestHeader(h http.Header) []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

func ResponseHeader(h http.Header) []attribute.KeyValue { _ = "STUB: not implemented"; return nil }
