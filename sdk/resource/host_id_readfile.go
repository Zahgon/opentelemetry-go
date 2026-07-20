//go:build linux || dragonfly || freebsd || netbsd || openbsd || solaris

package resource

func readFile(filename string) (string, error) { _ = "STUB: not implemented"; return "", nil }
