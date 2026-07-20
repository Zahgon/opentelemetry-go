package stdoutlog

import (
	"io"
	"os"
)

var (
	defaultWriter      io.Writer = os.Stdout
	defaultPrettyPrint           = false
	defaultTimestamps            = true
)

type config struct {
	Writer io.Writer

	PrettyPrint bool

	Timestamps bool
}

func newConfig(options []Option) config { _ = "STUB: not implemented"; return *new(config) }

type Option interface {
	apply(config) config
}

func WithWriter(w io.Writer) Option { _ = "STUB: not implemented"; return *new(Option) }

type writerOption struct {
	W io.Writer
}

func (o writerOption) apply(cfg config) config { _ = "STUB: not implemented"; return *new(config) }

func WithPrettyPrint() Option { _ = "STUB: not implemented"; return *new(Option) }

type prettyPrintOption bool

func (o prettyPrintOption) apply(cfg config) config { _ = "STUB: not implemented"; return *new(config) }

func WithoutTimestamps() Option { _ = "STUB: not implemented"; return *new(Option) }

type timestampsOption bool

func (o timestampsOption) apply(cfg config) config { _ = "STUB: not implemented"; return *new(config) }
