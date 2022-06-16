package day10

import "fmt"

type ParseStatus int

const (
	Unknown ParseStatus = iota
	Incomplete
	Corrupt
)

type ParseErr struct {
	status     ParseStatus
	unexpected rune
	expected   rune
	index      int
}

type ParseError interface {
	error
	getStatus() ParseStatus
	getUnexpectedRune() rune
	getExpectedRune() rune
	getIndex() int
}

func (e ParseErr) getUnexpectedRune() rune {
	return e.unexpected
}
func (e ParseErr) getIndex() int {
	return e.index
}

func (e ParseErr) getExpectedRune() rune {
	return e.expected
}

func (e ParseErr) getStatus() ParseStatus {
	return e.status
}
func (e ParseErr) Error() string {
	var errorType string
	switch e.status {
	case Corrupt:
		errorType = "corrupt"

	case Incomplete:
		errorType = "incomplete"

	default:
		errorType = "unknown"

	}

	return fmt.Sprintf("Parsing error %v, unknown rune %v", errorType, e.unexpected)
}
