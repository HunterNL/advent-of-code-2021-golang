package day21

import (
	"fmt"
	"testing"
)

func TestDice(t *testing.T) {
	g := startGame(0, 0)

	d1 := rollDeterministicDice(g)

	if d1 != 1 {
		t.Errorf("Expected roll 1 to be 1, not %v\n", d1)
	}

	d2 := rollDeterministicDice(g)

	if d2 != 2 {
		t.Errorf("Expected roll 2 to be 2, not %v\n", d2)
	}

	g.dice = 98

	d3 := rollDeterministicDice(g)

	if d3 != 99 {
		t.Errorf("Expected roll 3 to be 99, not %v\n", d3)
	}

	d4 := rollDeterministicDice(g)

	if d4 != 100 {
		t.Errorf("Expected roll 4 to be 100, not %v\n", d4)
	}

	d5 := rollDeterministicDice(g)

	if d5 != 1 {
		t.Errorf("Expected roll 5 to be 1, not %v\n", d5)
	}

}

func TestPlay(t *testing.T) {
	game := startGame(4, 8)

	playGame(game)

	if game.player1Score != 1000 {
		t.Errorf("Expected player 1 to score 1000, not %v\n", game.player1Score)
	}

	if game.player2Score != 745 {
		t.Errorf("Expected player 2 to score 745, not %v\n", game.player1Score)
	}
}

func TestScore(t *testing.T) {
	game := startGame(4, 8)

	rolls := playGame(game)

	score := rolls * *game.losingScore
	expected := 739785

	if score != expected {
		t.Errorf("Expected a score of %v not %v\n", expected, score)
	}
}

func TestOmniverseCount(t *testing.T) {
	// u :=

	o := omniverse{
		universe{
			player1Position: 4,
			player2Position: 8,
			player1Turn:     true,
		}: 1,
	}

	o = rollOmniverse(o)

	p1Wins, p2Wins := countWinners(o)

	p1Expected, p2Expected := 444356092776315, 341960390180808

	fmt.Printf("%v %T \n", p1Expected, p1Expected)

	if p1Wins != p1Expected {
		t.Errorf("Expected player 1 to win %.2e times, not %.2e\n", float64(p1Expected), float64(p1Wins))
	}
	if p2Wins != p2Expected {
		t.Errorf("Expected player 2 to win %.2e times, not %.2e\n", float64(p2Expected), float64(p2Wins))
	}

}
