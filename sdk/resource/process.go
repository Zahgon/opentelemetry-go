package resource

import (
	"context"
	"os"
	"os/user"
	"runtime"
)

type (
	pidProvider            func() int
	executablePathProvider func() (string, error)
	commandArgsProvider    func() []string
	ownerProvider          func() (*user.User, error)
	runtimeNameProvider    func() string
	runtimeVersionProvider func() string
	runtimeOSProvider      func() string
	runtimeArchProvider    func() string
)

var (
	defaultPidProvider            pidProvider            = os.Getpid
	defaultExecutablePathProvider executablePathProvider = os.Executable
	defaultCommandArgsProvider    commandArgsProvider    = func() []string { return os.Args }
	defaultOwnerProvider          ownerProvider          = user.Current
	defaultRuntimeNameProvider    runtimeNameProvider    = func() string {
		if runtime.Compiler == "gc" {
			return "go"
		}
		return runtime.Compiler
	}
	defaultRuntimeVersionProvider runtimeVersionProvider = runtime.Version
	defaultRuntimeOSProvider      runtimeOSProvider      = func() string { return runtime.GOOS }
	defaultRuntimeArchProvider    runtimeArchProvider    = func() string { return runtime.GOARCH }
)

var (
	pid            = defaultPidProvider
	executablePath = defaultExecutablePathProvider
	commandArgs    = defaultCommandArgsProvider
	owner          = defaultOwnerProvider
	runtimeName    = defaultRuntimeNameProvider
	runtimeVersion = defaultRuntimeVersionProvider
	runtimeOS      = defaultRuntimeOSProvider
	runtimeArch    = defaultRuntimeArchProvider
)

func setDefaultOSProviders() { _ = "STUB: not implemented"; return }

func setOSProviders(
	pidProvider pidProvider,
	executablePathProvider executablePathProvider,
	commandArgsProvider commandArgsProvider,
) {
	_ = "STUB: not implemented"
	return
}

func setDefaultRuntimeProviders() { _ = "STUB: not implemented"; return }

func setRuntimeProviders(
	runtimeNameProvider runtimeNameProvider,
	runtimeVersionProvider runtimeVersionProvider,
	runtimeOSProvider runtimeOSProvider,
	runtimeArchProvider runtimeArchProvider,
) {
	_ = "STUB: not implemented"
	return
}

func setDefaultUserProviders() { _ = "STUB: not implemented"; return }

func setUserProviders(ownerProvider ownerProvider) { _ = "STUB: not implemented"; return }

type (
	processPIDDetector                struct{}
	processExecutableNameDetector     struct{}
	processExecutablePathDetector     struct{}
	processCommandArgsDetector        struct{}
	processOwnerDetector              struct{}
	processRuntimeNameDetector        struct{}
	processRuntimeVersionDetector     struct{}
	processRuntimeDescriptionDetector struct{}
)

func (processPIDDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (processExecutableNameDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (processExecutablePathDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (processCommandArgsDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (processOwnerDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (processRuntimeNameDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (processRuntimeVersionDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (processRuntimeDescriptionDetector) Detect(context.Context) (*Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
