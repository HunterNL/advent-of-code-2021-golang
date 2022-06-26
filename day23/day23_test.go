package day23

import (
	"os"
	"testing"
)

func TestParseLarge(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		panic(err)
	}

	parsedBoard := parseBoard(string(bytes), PARSE_LARGE)

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
func TestParseSmall(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		panic(err)
	}

	parsedBoard := parseBoard(string(bytes), PARSE_SMALL)

	expectedBoard := board{
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod(0),
		amphipod('B'),
		amphipod('A'),
		amphipod(0),
		amphipod(0),
		amphipod('C'),
		amphipod('D'),
		amphipod(0),
		amphipod(0),
		amphipod('B'),
		amphipod('C'),
		amphipod(0),
		amphipod(0),
		amphipod('D'),
		amphipod('A'),
		amphipod(0),
		amphipod(0),
	}

	if expectedBoard != parsedBoard {
		t.Error("Parsing error")
	}

}

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

	t.Logf("\n%#v\n", board)
	t.Fail()
}

func TestScoreSmall(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		panic(err)
	}

	board := parseBoard(string(bytes), PARSE_SMALL)

	expectedScore := 12521

	score := findQuickestMoves(board, 2, desiredSmallBoard)

	// g.playMoves(&score, &config, 0)

	if score != expectedScore {
		t.Errorf("Expected a score of %v, not %v\n", expectedScore, score)
	}
}

func TestScoreLarge(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		panic(err)
	}

	board := parseBoard(string(bytes), PARSE_LARGE)

	expectedScore := 44169

	score := findQuickestMoves(board, 4, desiredLargeBoard)

	// g.playMoves(&score, &config, 0)

	if score != expectedScore {
		t.Errorf("Expected a score of %v, not %v\n", expectedScore, score)
	}
}

var perfectMoves = []move{
	{19, 6},
	{20, 0},
	{15, 5},
	{16, 4},
	{17, 1},
	{11, 3},
	{3, 17},
	{12, 3},
	{3, 16},
	{13, 3},
	{14, 2},
	{3, 14},
	{4, 13},
	{21, 4},
	{4, 15},
	{5, 12},
	{22, 5},
	{2, 22},
	{7, 2},
	{2, 11},
	{8, 4},
	{9, 3},
	{4, 21},
	{3, 20},
	{5, 9},
	{6, 19},
	{1, 8},
	{0, 7},
}

func TestPerfectMovement(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		panic(err)
	}

	board := parseBoard(string(bytes), PARSE_LARGE)

	expectedScore := 44169

	config := gameConfig{
		winState: desiredLargeBoard,
		roomSize: 4,
	}

	g := game{state: board}

	for i, v := range perfectMoves {
		b := g.state
		possibleMoves := availableDestinations(&b, v.from, &config)

		found := false
		for _, to := range possibleMoves {
			if to == v.to {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected move [%v -> %v] to be possible on step %v", v.from, v.to, i)
			t.FailNow()
		}

		g = g.applyMove(v.from, v.to, &config)

	}

	if g.score != expectedScore {
		t.Errorf("Expected a score of %v, not %v\n", expectedScore, g.score)
	}
}
