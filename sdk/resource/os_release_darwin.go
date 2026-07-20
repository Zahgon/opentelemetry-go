package resource

import (
	"encoding/xml"
	"io"
	"os"
)

type plist struct {
	XMLName xml.Name `xml:"plist"`
	Dict    dict     `xml:"dict"`
}

type dict struct {
	Key    []string `xml:"key"`
	String []string `xml:"string"`
}

func osRelease() string { _ = "STUB: not implemented"; return "" }

func getPlistFile() (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func parsePlistFile(file io.Reader) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildOSRelease(properties map[string]string) string { _ = "STUB: not implemented"; return "" }
