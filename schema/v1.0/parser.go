package schema

import (
	"io"

	"go.opentelemetry.io/otel/schema/v1.0/ast"
)

const supportedFormatMajor = 1

const supportedFormatMinor = 0

func ParseFile(schemaFilePath string) (*ast.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Parse(schemaFileContent io.Reader) (*ast.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
