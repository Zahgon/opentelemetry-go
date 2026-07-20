//go:build linux

package resource

var platformHostIDReader hostIDReader = &hostIDReaderLinux{
	readFile: readFile,
}
