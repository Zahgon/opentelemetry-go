package semconv

import (
	"errors"
	"fmt"
	"reflect"

	"go.opentelemetry.io/otel/attribute"
)

func ErrorType(err error) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

func errorType(err error) string { _ = "STUB: not implemented"; return "" }

var fmtWrapErrorType = reflect.TypeOf(fmt.Errorf("wrapped: %w", errors.New("err")))

func unwrapFmtWrapped(err error) error { _ = "STUB: not implemented"; return nil }
