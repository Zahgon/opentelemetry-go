package metric

import (
	"time"
)

const (
	envInterval = "OTEL_METRIC_EXPORT_INTERVAL"

	envTimeout = "OTEL_METRIC_EXPORT_TIMEOUT"
)

func envDuration(key string, defaultValue time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
