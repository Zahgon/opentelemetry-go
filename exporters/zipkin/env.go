package zipkin

const (
	envEndpoint = "OTEL_EXPORTER_ZIPKIN_ENDPOINT"
)

func envOr(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }
