package rpcconv

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	addOptPool = &sync.Pool{New: func() any { return &[]metric.AddOption{} }}
	recOptPool = &sync.Pool{New: func() any { return &[]metric.RecordOption{} }}
)

type ClientDuration struct {
	metric.Float64Histogram
}

func NewClientDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ClientDuration, error) {
	_ = "STUB: not implemented"
	return *new(ClientDuration), nil
}

func (m ClientDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ClientDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientDuration) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ClientRequestSize struct {
	metric.Int64Histogram
}

func NewClientRequestSize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientRequestSize, error) {
	_ = "STUB: not implemented"
	return *new(ClientRequestSize), nil
}

func (m ClientRequestSize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientRequestSize) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientRequestSize) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ClientRequestsPerRPC struct {
	metric.Int64Histogram
}

func NewClientRequestsPerRPC(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientRequestsPerRPC, error) {
	_ = "STUB: not implemented"
	return *new(ClientRequestsPerRPC), nil
}

func (m ClientRequestsPerRPC) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientRequestsPerRPC) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestsPerRPC) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientRequestsPerRPC) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientRequestsPerRPC) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ClientResponseSize struct {
	metric.Int64Histogram
}

func NewClientResponseSize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientResponseSize, error) {
	_ = "STUB: not implemented"
	return *new(ClientResponseSize), nil
}

func (m ClientResponseSize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientResponseSize) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientResponseSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientResponseSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientResponseSize) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ClientResponsesPerRPC struct {
	metric.Int64Histogram
}

func NewClientResponsesPerRPC(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ClientResponsesPerRPC, error) {
	_ = "STUB: not implemented"
	return *new(ClientResponsesPerRPC), nil
}

func (m ClientResponsesPerRPC) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ClientResponsesPerRPC) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientResponsesPerRPC) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientResponsesPerRPC) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientResponsesPerRPC) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ServerDuration struct {
	metric.Float64Histogram
}

func NewServerDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ServerDuration, error) {
	_ = "STUB: not implemented"
	return *new(ServerDuration), nil
}

func (m ServerDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ServerDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerDuration) Record(ctx context.Context, val float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ServerRequestSize struct {
	metric.Int64Histogram
}

func NewServerRequestSize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ServerRequestSize, error) {
	_ = "STUB: not implemented"
	return *new(ServerRequestSize), nil
}

func (m ServerRequestSize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ServerRequestSize) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerRequestSize) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ServerRequestsPerRPC struct {
	metric.Int64Histogram
}

func NewServerRequestsPerRPC(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ServerRequestsPerRPC, error) {
	_ = "STUB: not implemented"
	return *new(ServerRequestsPerRPC), nil
}

func (m ServerRequestsPerRPC) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ServerRequestsPerRPC) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestsPerRPC) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerRequestsPerRPC) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerRequestsPerRPC) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ServerResponseSize struct {
	metric.Int64Histogram
}

func NewServerResponseSize(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ServerResponseSize, error) {
	_ = "STUB: not implemented"
	return *new(ServerResponseSize), nil
}

func (m ServerResponseSize) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ServerResponseSize) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerResponseSize) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerResponseSize) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerResponseSize) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type ServerResponsesPerRPC struct {
	metric.Int64Histogram
}

func NewServerResponsesPerRPC(
	m metric.Meter,
	opt ...metric.Int64HistogramOption,
) (ServerResponsesPerRPC, error) {
	_ = "STUB: not implemented"
	return *new(ServerResponsesPerRPC), nil
}

func (m ServerResponsesPerRPC) Inst() metric.Int64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Int64Histogram)
}

func (ServerResponsesPerRPC) Name() string { _ = "STUB: not implemented"; return "" }

func (ServerResponsesPerRPC) Unit() string { _ = "STUB: not implemented"; return "" }

func (ServerResponsesPerRPC) Description() string { _ = "STUB: not implemented"; return "" }

func (m ServerResponsesPerRPC) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}
