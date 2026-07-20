//go:build !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris && !windows

package resource

type hostIDReaderUnsupported struct{}

func (*hostIDReaderUnsupported) read() (string, error) { _ = "STUB: not implemented"; return "", nil }

var platformHostIDReader hostIDReader = &hostIDReaderUnsupported{}
