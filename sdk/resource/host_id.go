package resource

import (
	"context"
)

type hostIDProvider func() (string, error)

var defaultHostIDProvider hostIDProvider = platformHostIDReader.read

var hostID = defaultHostIDProvider

type hostIDReader interface {
	read() (string, error)
}

type fileReader func(string) (string, error)

type commandExecutor func(string, ...string) (string, error)

type hostIDReaderBSD struct {
	execCommand commandExecutor
	readFile    fileReader
}

func (r *hostIDReaderBSD) read() (string, error) { _ = "STUB: not implemented"; return "", nil }

type hostIDReaderDarwin struct {
	execCommand commandExecutor
}

func (r *hostIDReaderDarwin) read() (string, error) { _ = "STUB: not implemented"; return "", nil }

type hostIDReaderLinux struct {
	readFile fileReader
}

func (r *hostIDReaderLinux) read() (string, error) { _ = "STUB: not implemented"; return "", nil }

type hostIDDetector struct{}

func (hostIDDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
