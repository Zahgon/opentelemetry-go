package resource

import (
	"context"
	"io"
	"os"
	"regexp"
)

type containerIDProvider func() (string, error)

var (
	containerID         containerIDProvider = getContainerIDFromCGroup
	cgroupContainerIDRe                     = regexp.MustCompile(`^.*/(?:.*[-:])?([0-9a-f]+)(?:\.|\s*$)`)
)

type cgroupContainerIDDetector struct{}

const cgroupPath = "/proc/self/cgroup"

func (cgroupContainerIDDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	defaultOSStat = os.Stat
	osStat        = defaultOSStat

	defaultOSOpen = func(name string) (io.ReadCloser, error) {
		return os.Open(name)
	}
	osOpen = defaultOSOpen
)

func getContainerIDFromCGroup() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getContainerIDFromReader(reader io.Reader) string { _ = "STUB: not implemented"; return "" }

func getContainerIDFromLine(line string) string { _ = "STUB: not implemented"; return "" }
