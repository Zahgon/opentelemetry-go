//go:build aix || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos

package resource

import (
	"io"
	"os"
)

func osRelease() string { _ = "STUB: not implemented"; return "" }

func getOSReleaseFile() (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func parseOSReleaseFile(file io.Reader) map[string]string { _ = "STUB: not implemented"; return nil }

func skip(line string) bool { _ = "STUB: not implemented"; return false }

func parse(line string) (string, string, bool) { _ = "STUB: not implemented"; return "", "", false }

func unquote(s string) string { _ = "STUB: not implemented"; return "" }

func unescape(s string) string { _ = "STUB: not implemented"; return "" }

func buildOSRelease(values map[string]string) string { _ = "STUB: not implemented"; return "" }
