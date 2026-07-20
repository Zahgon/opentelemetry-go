//go:build windows

package resource

type hostIDReaderWindows struct{}

func (*hostIDReaderWindows) read() (string, error) { _ = "STUB: not implemented"; return "", nil }

var platformHostIDReader hostIDReader = &hostIDReaderWindows{}
