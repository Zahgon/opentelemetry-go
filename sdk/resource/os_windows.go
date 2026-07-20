package resource

import (
	"golang.org/x/sys/windows/registry"
)

func platformOSDescription() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getStringValue(name string, k registry.Key) string { _ = "STUB: not implemented"; return "" }

func getIntegerValue(name string, k registry.Key) uint64 { _ = "STUB: not implemented"; return 0 }

func readProductName(k registry.Key) string { _ = "STUB: not implemented"; return "" }

func readDisplayVersion(k registry.Key) string { _ = "STUB: not implemented"; return "" }

func readReleaseID(k registry.Key) string { _ = "STUB: not implemented"; return "" }

func readCurrentMajorVersionNumber(k registry.Key) string { _ = "STUB: not implemented"; return "" }

func readCurrentMinorVersionNumber(k registry.Key) string { _ = "STUB: not implemented"; return "" }

func readCurrentBuildNumber(k registry.Key) string { _ = "STUB: not implemented"; return "" }

func readUBR(k registry.Key) string { _ = "STUB: not implemented"; return "" }
