package otel

import (
	"go.opentelemetry.io/otel/internal/global"
)

var _ ErrorHandler = (*global.ErrDelegator)(nil)

func GetErrorHandler() ErrorHandler { _ = "STUB: not implemented"; return *new(ErrorHandler) }

func SetErrorHandler(h ErrorHandler) { _ = "STUB: not implemented"; return }

func Handle(err error) { _ = "STUB: not implemented"; return }
