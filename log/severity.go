//go:generate stringer -type=Severity -linecomment

package log

type Severity int

const (
	SeverityUndefined Severity = 0

	SeverityTrace1 Severity = 1
	SeverityTrace2 Severity = 2
	SeverityTrace3 Severity = 3
	SeverityTrace4 Severity = 4

	SeverityDebug1 Severity = 5
	SeverityDebug2 Severity = 6
	SeverityDebug3 Severity = 7
	SeverityDebug4 Severity = 8

	SeverityInfo1 Severity = 9
	SeverityInfo2 Severity = 10
	SeverityInfo3 Severity = 11
	SeverityInfo4 Severity = 12

	SeverityWarn1 Severity = 13
	SeverityWarn2 Severity = 14
	SeverityWarn3 Severity = 15
	SeverityWarn4 Severity = 16

	SeverityError1 Severity = 17
	SeverityError2 Severity = 18
	SeverityError3 Severity = 19
	SeverityError4 Severity = 20

	SeverityFatal1 Severity = 21
	SeverityFatal2 Severity = 22
	SeverityFatal3 Severity = 23
	SeverityFatal4 Severity = 24

	SeverityTrace = SeverityTrace1
	SeverityDebug = SeverityDebug1
	SeverityInfo  = SeverityInfo1
	SeverityWarn  = SeverityWarn1
	SeverityError = SeverityError1
	SeverityFatal = SeverityFatal1
)
