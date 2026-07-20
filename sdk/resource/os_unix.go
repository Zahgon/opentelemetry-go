//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos

package resource

import (
	"os"

	"golang.org/x/sys/unix"
)

type unameProvider func(buf *unix.Utsname) (err error)

var defaultUnameProvider unameProvider = unix.Uname

var currentUnameProvider = defaultUnameProvider

func setDefaultUnameProvider() { _ = "STUB: not implemented"; return }

func setUnameProvider(unameProvider unameProvider) { _ = "STUB: not implemented"; return }

func platformOSDescription() (string, error) { _ = "STUB: not implemented"; return "", nil }

func uname() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getFirstAvailableFile(candidates []string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
