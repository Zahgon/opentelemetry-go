//go:generate stringer -type=Temporality

package metricdata

type Temporality uint8

const (

	//nolint:unused
	undefinedTemporality Temporality = iota

	CumulativeTemporality

	DeltaTemporality
)

func (t Temporality) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
