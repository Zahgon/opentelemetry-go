//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris && !windows && !zos

package resource

func platformOSDescription() (string, error) { _ = "STUB: not implemented"; return "", nil }
