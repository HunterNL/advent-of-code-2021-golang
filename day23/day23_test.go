package day23

import (
	"math"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		panic(err)
	}

	parsedBoard := parseBoard(string(bytes))

	expectedBoard := board{
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod('B'),
		amphipod('D'),
		amphipod('D'),
		amphipod('A'),
		amphipod('C'),
		amphipod('C'),
		amphipod('B'),
		amphipod('D'),
		amphipod('B'),
		amphipod('B'),
		amphipod('A'),
		amphipod('C'),
		amphipod('D'),
		amphipod('A'),
		amphipod('C'),
		amphipod('A'),
	}

	if expectedBoard != parsedBoard {
		t.Error("Parsing error")
	}

}

// 	for position, expectedPod := range expectedBoard {
// 		if board[position] != expectedPod {
// 			t.Errorf("Expected pod %s in %v, not %s\n", string(expectedPod), position, string(board[position]))
// 		}
// 	}

// 	if len(board) != len(expectedBoard) {
// 		t.Error("Length mismatch")
// 	}

// }

func TestPrint(t *testing.T) {
	board := board{
		amphipod('A'),
		amphipod('B'),
		amphipod('C'),
		amphipod('D'),
		amphipod('E'),
		amphipod('F'),
		amphipod('G'),
		amphipod('H'),
		amphipod('I'),
		amphipod('J'),
		amphipod('K'),
		amphipod('L'),
		amphipod('M'),
		amphipod('N'),
		amphipod('O'),
		amphipod('P'),
		amphipod('Q'),
		amphipod('R'),
		amphipod('S'),
		amphipod('T'),
		amphipod('U'),
		amphipod('V'),
		amphipod('W'),
	}

	t.Logf("\n%v\n", board)
	t.Fail()
}

func TestScore(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		panic(err)
	}

	board := parseBoard(string(bytes))

	expectedScore := 44169

	g := game{state: board}

	score := math.MaxInt

	g.playMoves(&score)

	if score != expectedScore {
		t.Errorf("Expected a score of %v, not %v\n", expectedScore, score)
	}
}
