package day10

type RuneScan struct {
	slice []rune
	index int
}

type RuneScanner interface {
	scan() rune
	peek() rune
	consume(rune)
	atEnd() bool
	peekFor(rune) rune
}

func (r *RuneScan) panicAtEnd() {
	if r.atEnd() {
		panic(ParseErr{status: Incomplete, index: r.index})
	}
}

func (r *RuneScan) scan() rune {
	r.panicAtEnd()
	r.index++
	return r.slice[r.index-1]
}

func (r *RuneScan) peek() rune {
	r.panicAtEnd()
	return r.slice[r.index]
}

func (r *RuneScan) peekFor(ru rune) rune {
	if r.atEnd() {
		panic(ParseErr{status: Incomplete, expected: ru})
	}

	return r.peek()
}

func (r *RuneScan) atEnd() bool {
	return r.index == len(r.slice)
}
func (r *RuneScan) consume(expected rune) {
	if r.atEnd() {
		panic(ParseErr{status: Incomplete, expected: expected, index: r.index})
	}

	if r.peek() == expected {
		r.scan()
	} else {
		panic(ParseErr{status: Corrupt, expected: expected, unexpected: r.peek(), index: r.index})
	}
}

func newScanner(str string) *RuneScan {
	r := RuneScan{
		slice: []rune(str),
		index: 0,
	}
	return &r
}
