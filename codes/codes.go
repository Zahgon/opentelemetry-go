package codes

const (
	Unset Code = 0

	Error Code = 1

	Ok Code = 2

	maxCode = 3
)

type Code uint32

var codeToStr = map[Code]string{
	Unset: "Unset",
	Error: "Error",
	Ok:    "Ok",
}

var strToCode = map[string]Code{
	`"Unset"`: Unset,
	`"Error"`: Error,
	`"Ok"`:    Ok,
}

func (c Code) String() string { _ = "STUB: not implemented"; return "" }

func (c *Code) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (c *Code) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
