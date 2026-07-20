package messagingconv

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

type ErrorTypeAttr string

var ErrorTypeOther ErrorTypeAttr = "_OTHER"

type OperationTypeAttr string

var (
	OperationTypeCreate OperationTypeAttr = "create"

	OperationTypeSend OperationTypeAttr = "send"

	OperationTypeReceive OperationTypeAttr = "receive"

	OperationTypeProcess OperationTypeAttr = "process"

	OperationTypeSettle OperationTypeAttr = "settle"
)

type SystemAttr string

var (
	SystemActiveMQ SystemAttr = "activemq"

	SystemAWSSNS SystemAttr = "aws.sns"

	SystemAWSSQS SystemAttr = "aws_sqs"

	SystemEventGrid SystemAttr = "eventgrid"

	SystemEventHubs SystemAttr = "eventhubs"

	SystemServiceBus SystemAttr = "servicebus"

	SystemGCPPubSub SystemAttr = "gcp_pubsub"

	SystemJMS SystemAttr = "jms"

	SystemKafka SystemAttr = "kafka"

	SystemRabbitMQ SystemAttr = "rabbitmq"

	SystemRocketMQ SystemAttr = "rocketmq"

	SystemPulsar SystemAttr = "pulsar"
)

type ClientConsumedMessages struct {
	metric.Int64Counter
}

var newClientConsumedMessagesOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of messages that were delivered to the application."),
	metric.WithUnit("{message}"),
}

func NewClientConsumedMessages(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientConsumedMessages, error) {
	_ = "STUB: not implemented"
	return *new(ClientConsumedMessages), nil
}

func (m ClientConsumedMessages) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientConsumedMessages) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientConsumedMessages) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientConsumedMessages) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientConsumedMessages) Add(
	ctx context.Context,
	incr int64,
	operationName string,
	system SystemAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientConsumedMessages) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientConsumedMessages) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConsumedMessages) AttrConsumerGroupName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConsumedMessages) AttrDestinationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConsumedMessages) AttrDestinationSubscriptionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConsumedMessages) AttrDestinationTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConsumedMessages) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConsumedMessages) AttrDestinationPartitionID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientConsumedMessages) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientOperationDuration struct {
	metric.Float64Histogram
}

var newClientOperationDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Duration of messaging operation initiated by a producer or consumer client."),
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
	operationName string,
	system SystemAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientOperationDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientOperationDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrConsumerGroupName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrDestinationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrDestinationSubscriptionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrDestinationTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrOperationType(val OperationTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrDestinationPartitionID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientOperationDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ClientSentMessages struct {
	metric.Int64Counter
}

var newClientSentMessagesOpts = []metric.Int64CounterOption{
	metric.WithDescription("Number of messages producer attempted to send to the broker."),
	metric.WithUnit("{message}"),
}

func NewClientSentMessages(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ClientSentMessages, error) {
	_ = "STUB: not implemented"
	return *new(ClientSentMessages), nil
}

func (m ClientSentMessages) Inst() metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

func (ClientSentMessages) Name() string { _ = "STUB: not implemented"; return "" }

func (ClientSentMessages) Unit() string { _ = "STUB: not implemented"; return "" }

func (ClientSentMessages) Description() string { _ = "STUB: not implemented"; return "" }

func (m ClientSentMessages) Add(
	ctx context.Context,
	incr int64,
	operationName string,
	system SystemAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ClientSentMessages) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ClientSentMessages) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSentMessages) AttrDestinationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSentMessages) AttrDestinationTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSentMessages) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSentMessages) AttrDestinationPartitionID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ClientSentMessages) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

type ProcessDuration struct {
	metric.Float64Histogram
}

var newProcessDurationOpts = []metric.Float64HistogramOption{
	metric.WithDescription("Duration of processing operation."),
	metric.WithUnit("s"),
}

func NewProcessDuration(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (ProcessDuration, error) {
	_ = "STUB: not implemented"
	return *new(ProcessDuration), nil
}

func (m ProcessDuration) Inst() metric.Float64Histogram {
	_ = "STUB: not implemented"
	return *new(metric.Float64Histogram)
}

func (ProcessDuration) Name() string { _ = "STUB: not implemented"; return "" }

func (ProcessDuration) Unit() string { _ = "STUB: not implemented"; return "" }

func (ProcessDuration) Description() string { _ = "STUB: not implemented"; return "" }

func (m ProcessDuration) Record(
	ctx context.Context,
	val float64,
	operationName string,
	system SystemAttr,
	attrs ...attribute.KeyValue,
) {
	_ = "STUB: not implemented"
	return
}

func (m ProcessDuration) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	_ = "STUB: not implemented"
	return
}

func (ProcessDuration) AttrErrorType(val ErrorTypeAttr) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ProcessDuration) AttrConsumerGroupName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ProcessDuration) AttrDestinationName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ProcessDuration) AttrDestinationSubscriptionName(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ProcessDuration) AttrDestinationTemplate(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ProcessDuration) AttrServerAddress(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ProcessDuration) AttrDestinationPartitionID(val string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func (ProcessDuration) AttrServerPort(val int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}
