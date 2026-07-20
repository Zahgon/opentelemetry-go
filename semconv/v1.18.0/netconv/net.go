package netconv

import (
	"net"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/semconv/internal/v2"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
)

var nc = &internal.NetConv{
	NetHostNameKey:     semconv.NetHostNameKey,
	NetHostPortKey:     semconv.NetHostPortKey,
	NetPeerNameKey:     semconv.NetPeerNameKey,
	NetPeerPortKey:     semconv.NetPeerPortKey,
	NetSockFamilyKey:   semconv.NetSockFamilyKey,
	NetSockPeerAddrKey: semconv.NetSockPeerAddrKey,
	NetSockPeerPortKey: semconv.NetSockPeerPortKey,
	NetSockHostAddrKey: semconv.NetSockHostAddrKey,
	NetSockHostPortKey: semconv.NetSockHostPortKey,
	NetTransportOther:  semconv.NetTransportOther,
	NetTransportTCP:    semconv.NetTransportTCP,
	NetTransportUDP:    semconv.NetTransportUDP,
	NetTransportInProc: semconv.NetTransportInProc,
}

func Transport(network string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func Client(address string, conn net.Conn) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func Server(address string, ln net.Listener) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}
