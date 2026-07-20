package schema

import (
	"io"

	"go.opentelemetry.io/otel/schema/v1.1/ast"
)

const supportedFormatMajor = 1

const supportedFormatMinor = 1

func ParseFile(schemaFilePath string) (*ast.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Parse(schemaFileContent io.Reader) (*ast.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
