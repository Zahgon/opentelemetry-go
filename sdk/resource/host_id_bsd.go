//go:build dragonfly || freebsd || netbsd || openbsd || solaris

package resource

var platformHostIDReader hostIDReader = &hostIDReaderBSD{
	execCommand: execCommand,
	readFile:    readFile,
}
