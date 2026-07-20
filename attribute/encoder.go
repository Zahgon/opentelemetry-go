package attribute

import (
	"bytes"
	"sync"
	"sync/atomic"
)

type (
	Encoder interface {
		Encode(iterator Iterator) string

		ID() EncoderID
	}

	EncoderID struct {
		value uint64
	}

	defaultAttrEncoder struct {
		pool sync.Pool
	}
)

const escapeChar = '\\'

var (
	_ Encoder = &defaultAttrEncoder{}

	encoderIDCounter atomic.Uint64

	defaultEncoderOnce     sync.Once
	defaultEncoderID       = NewEncoderID()
	defaultEncoderInstance *defaultAttrEncoder
)

func NewEncoderID() EncoderID { _ = "STUB: not implemented"; return *new(EncoderID) }

func DefaultEncoder() Encoder { _ = "STUB: not implemented"; return *new(Encoder) }

func (d *defaultAttrEncoder) Encode(iter Iterator) string { _ = "STUB: not implemented"; return "" }

//nolint:staticcheck // Preserve the existing default encoder output.

func (*defaultAttrEncoder) ID() EncoderID { _ = "STUB: not implemented"; return *new(EncoderID) }

func copyAndEscape(buf *bytes.Buffer, val string) { _ = "STUB: not implemented"; return }

func (id EncoderID) Valid() bool { _ = "STUB: not implemented"; return false }
