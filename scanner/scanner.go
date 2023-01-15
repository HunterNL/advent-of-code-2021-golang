package scanner

import (
	"fmt"
	"io"
)

type Scanner interface {
	io.RuneScanner
	Peek() rune
	Consume()
	Expect(rune)
}

type scan struct {
	io.RuneScanner
}

func (s *scan) Peek() rune {
	r, _, err := s.ReadRune()
	if err != nil {
		panic(err)
	}
	s.UnreadRune()
	return r
}

func (s *scan) Consume() {
	s.ReadRune()
}

func (s *scan) Expect(expected rune) {
	if r, _, err := s.ReadRune(); r != expected {
		if err != nil {
			panic(fmt.Sprintf("Error while expecting %v: %v", expected, err))
		}

		panic(fmt.Sprintf("Expected %v but got %v", expected, r))

	}
}

func NewScanner(io io.RuneScanner) Scanner {
	return &scan{io}
}
