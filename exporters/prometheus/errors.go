package prometheus

import "errors"

var (
	errInvalidMetricType  = errors.New("invalid metric type")
	errInvalidMetric      = errors.New("invalid metric")
	errEHScaleBelowMin    = errors.New("exponential histogram scale below minimum supported")
	errBridgeNotSupported = errors.New(
		"metrics from the prometheus bridge are not supported in the prometheus exporter, and will be dropped",
	)
)
