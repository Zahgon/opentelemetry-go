//go:build darwin || dragonfly || freebsd || netbsd || openbsd || solaris

package resource

func execCommand(name string, arg ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
