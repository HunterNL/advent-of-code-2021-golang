package day23

import (
	"math"
	"os"
	"testing"
)

func TestParseP1(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		panic(err)
	}

	board := parsep1Board(string(bytes))

	expectedBoard := map[vec2]amphipod{
		{2, 1}: amphipod('B'),
		{2, 2}: amphipod('A'),
		{4, 1}: amphipod('C'),
		{4, 2}: amphipod('D'),
		{6, 1}: amphipod('B'),
		{6, 2}: amphipod('C'),
		{8, 1}: amphipod('D'),
		{8, 2}: amphipod('A'),
	}

	for position, expectedPod := range expectedBoard {
		if board[position] != expectedPod {
			t.Errorf("Expected pod %s in %v, not %s\n", string(expectedPod), position, string(board[position]))
		}
	}

	if len(board) != len(expectedBoard) {
		t.Error("Length mismatch")
	}

}

func TestScoreP1(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		panic(err)
	}

	board := parsep1Board(string(bytes))

	score := math.MaxInt

	expectedScore := 12521

	g := gamep1{state: board}

	// allGames := []gamep1{}

	g.playMovesP1(&score)

	// score := lowestScore(allGames)

	if score != expectedScore {
		t.Errorf("Expected a score of %v, not %v\n", expectedScore, score)
	}
}
