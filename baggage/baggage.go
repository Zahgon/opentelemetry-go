package baggage

import (
	"errors"
	"unicode/utf8"

	"go.opentelemetry.io/otel/internal/baggage"
)

const (
	maxParseErrors = 5

	maxMembers               = 64
	maxBytesPerBaggageString = 8192

	listDelimiter     = ","
	keyValueDelimiter = "="
	propertyDelimiter = ";"
)

var (
	errInvalidKey      = errors.New("invalid key")
	errInvalidValue    = errors.New("invalid value")
	errInvalidProperty = errors.New("invalid baggage list-member property")
	errInvalidMember   = errors.New("invalid baggage list-member")
	errMemberNumber    = errors.New("too many list-members in baggage-string")
	errBaggageBytes    = errors.New("baggage-string too large")
)

type Property struct {
	key, value string

	hasValue bool
}

func NewKeyProperty(key string) (Property, error) {
	_ = "STUB: not implemented"
	return *new(Property), nil
}

func NewKeyValueProperty(key, value string) (Property, error) {
	_ = "STUB: not implemented"
	return *new(Property), nil
}

func NewKeyValuePropertyRaw(key, value string) (Property, error) {
	_ = "STUB: not implemented"
	return *new(Property), nil
}

func newInvalidProperty() Property { _ = "STUB: not implemented"; return *new(Property) }

func parseProperty(property string) (Property, error) {
	_ = "STUB: not implemented"
	return *new(Property), nil
}

func (p Property) validate() error { _ = "STUB: not implemented"; return nil }

func (p Property) Key() string { _ = "STUB: not implemented"; return "" }

func (p Property) Value() (string, bool) { _ = "STUB: not implemented"; return "", false }

func (p Property) String() string { _ = "STUB: not implemented"; return "" }

type properties []Property

func fromInternalProperties(iProps []baggage.Property) properties {
	_ = "STUB: not implemented"
	return *new(properties)
}

func (p properties) asInternal() []baggage.Property { _ = "STUB: not implemented"; return nil }

func (p properties) Copy() properties { _ = "STUB: not implemented"; return *new(properties) }

func (p properties) validate() error { _ = "STUB: not implemented"; return nil }

func (p properties) String() string { _ = "STUB: not implemented"; return "" }

type Member struct {
	key, value string
	properties properties

	hasData bool
}

func NewMember(key, value string, props ...Property) (Member, error) {
	_ = "STUB: not implemented"
	return *new(Member), nil
}

func NewMemberRaw(key, value string, props ...Property) (Member, error) {
	_ = "STUB: not implemented"
	return *new(Member), nil
}

func newInvalidMember() Member { _ = "STUB: not implemented"; return *new(Member) }

func parseMember(member string) (Member, error) {
	_ = "STUB: not implemented"
	return *new(Member), nil
}

func replaceInvalidUTF8Sequences(c int, unescapeVal string) string {
	_ = "STUB: not implemented"
	return ""
}

func (m Member) validate() error { _ = "STUB: not implemented"; return nil }

func (m Member) Key() string { _ = "STUB: not implemented"; return "" }

func (m Member) Value() string { _ = "STUB: not implemented"; return "" }

func (m Member) Properties() []Property { _ = "STUB: not implemented"; return nil }

func (m Member) String() string { _ = "STUB: not implemented"; return "" }

type Baggage struct { //nolint:golint
	list baggage.List
}

func New(members ...Member) (Baggage, error) { _ = "STUB: not implemented"; return *new(Baggage), nil }

func Parse(bStr string) (Baggage, error) { _ = "STUB: not implemented"; return *new(Baggage), nil }

func (b Baggage) Member(key string) Member { _ = "STUB: not implemented"; return *new(Member) }

func (b Baggage) Members() []Member { _ = "STUB: not implemented"; return nil }

func (b Baggage) SetMember(member Member) (Baggage, error) {
	_ = "STUB: not implemented"
	return *new(Baggage), nil
}

func (b Baggage) DeleteMember(key string) Baggage { _ = "STUB: not implemented"; return *new(Baggage) }

func (b Baggage) Len() int { _ = "STUB: not implemented"; return 0 }

func (b Baggage) String() string { _ = "STUB: not implemented"; return "" }

func parsePropertyInternal(s string) (p Property, ok bool) {
	_ = "STUB: not implemented"
	return *new(Property), false
}

func skipSpace(s string, offset int) int { _ = "STUB: not implemented"; return 0 }

var safeKeyCharset = [utf8.RuneSelf]bool{

	'#':  true,
	'$':  true,
	'%':  true,
	'&':  true,
	'\'': true,

	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,

	'A': true,
	'B': true,
	'C': true,
	'D': true,
	'E': true,
	'F': true,
	'G': true,
	'H': true,
	'I': true,
	'J': true,
	'K': true,
	'L': true,
	'M': true,
	'N': true,
	'O': true,
	'P': true,
	'Q': true,
	'R': true,
	'S': true,
	'T': true,
	'U': true,
	'V': true,
	'W': true,
	'X': true,
	'Y': true,
	'Z': true,

	'^': true,
	'_': true,
	'`': true,
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'g': true,
	'h': true,
	'i': true,
	'j': true,
	'k': true,
	'l': true,
	'm': true,
	'n': true,
	'o': true,
	'p': true,
	'q': true,
	'r': true,
	's': true,
	't': true,
	'u': true,
	'v': true,
	'w': true,
	'x': true,
	'y': true,
	'z': true,

	'!': true,
	'*': true,
	'+': true,
	'-': true,
	'.': true,
	'|': true,
	'~': true,
}

func validateBaggageName(s string) bool { _ = "STUB: not implemented"; return false }

func validateBaggageValue(s string) bool { _ = "STUB: not implemented"; return false }

func validateKey(s string) bool { _ = "STUB: not implemented"; return false }

func validateKeyChar(c int32) bool { _ = "STUB: not implemented"; return false }

func validateValue(s string) bool { _ = "STUB: not implemented"; return false }

var safeValueCharset = [utf8.RuneSelf]bool{
	'!': true,

	'#':  true,
	'$':  true,
	'%':  true,
	'&':  true,
	'\'': true,
	'(':  true,
	')':  true,
	'*':  true,
	'+':  true,

	'-': true,
	'.': true,
	'/': true,
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
	':': true,

	'<': true,
	'=': true,
	'>': true,
	'?': true,
	'@': true,
	'A': true,
	'B': true,
	'C': true,
	'D': true,
	'E': true,
	'F': true,
	'G': true,
	'H': true,
	'I': true,
	'J': true,
	'K': true,
	'L': true,
	'M': true,
	'N': true,
	'O': true,
	'P': true,
	'Q': true,
	'R': true,
	'S': true,
	'T': true,
	'U': true,
	'V': true,
	'W': true,
	'X': true,
	'Y': true,
	'Z': true,
	'[': true,

	']': true,
	'^': true,
	'_': true,
	'`': true,
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'g': true,
	'h': true,
	'i': true,
	'j': true,
	'k': true,
	'l': true,
	'm': true,
	'n': true,
	'o': true,
	'p': true,
	'q': true,
	'r': true,
	's': true,
	't': true,
	'u': true,
	'v': true,
	'w': true,
	'x': true,
	'y': true,
	'z': true,
	'{': true,
	'|': true,
	'}': true,
	'~': true,
}

func validateValueChar(c int32) bool { _ = "STUB: not implemented"; return false }

func valueEscape(s string) string { _ = "STUB: not implemented"; return "" }

func shouldEscape(c byte) bool { _ = "STUB: not implemented"; return false }
