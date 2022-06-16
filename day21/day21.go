package day21

import "fmt"

type game struct {
	player1Position int
	player2Position int
	player1Score    int
	player2Score    int
	dice            int
	losingScore     *int
}

type universe struct {
	player1Position int
	player2Position int
	player1Score    int
	player2Score    int
	player1Turn     bool
}

func (u *universe) isWon() bool {
	return u.player1Score >= 21 || u.player2Score >= 21
}

type omniverse = map[universe]int

func universeVariations(u universe) omniverse {
	if u.isWon() {
		return map[universe]int{u: 1}
	}
	return map[universe]int{
		advanceUniverse(u, 3): 1,
		advanceUniverse(u, 4): 3,
		advanceUniverse(u, 5): 6,
		advanceUniverse(u, 6): 7,
		advanceUniverse(u, 7): 6,
		advanceUniverse(u, 8): 3,
		advanceUniverse(u, 9): 1,
	}
}

func advanceOmniverse(o omniverse) (omniverse, bool) {
	out := make(omniverse)
	doneCount := 0

	for universe, count := range o {
		additions := universeVariations(universe)

		for extraUniverse, extraCount := range additions {
			out[extraUniverse] = count*extraCount + out[extraUniverse]
		}

		if len(additions) == 1 {
			doneCount++
		}
	}

	done := doneCount == len(out)

	return out, done
}

func advanceUniverse(u universe, diceRoll int) universe {
	if u.player1Turn {
		advancePlayer(&u.player1Position, &u.player1Score, diceRoll)
	} else {
		advancePlayer(&u.player2Position, &u.player2Score, diceRoll)
	}
	u.player1Turn = !u.player1Turn

	return u
}

func advancePlayer(position, score *int, movement int) {
	(*position) = (*position+movement-1)%10 + 1
	(*score) = *score + *position
}

func startGame(p1, p2 int) *game {
	return &game{player1Position: p1, player2Position: p2}
}

func rollDeterministicDice(g *game) int {
	g.dice = ((g.dice)%100 + 1)
	return g.dice
}

func rollDeterministicTrice(g *game) int {
	return rollDeterministicDice(g) + rollDeterministicDice(g) + rollDeterministicDice(g)
}

func playGame(g *game) int {
	rolls := 0

	for {
		advancePlayer(&g.player1Position, &g.player1Score, rollDeterministicTrice(g))
		rolls += 3
		if g.player1Score >= 1000 {
			g.losingScore = &g.player2Score
			fmt.Println("Player 1 win")
			return rolls
		}

		advancePlayer(&g.player2Position, &g.player2Score, rollDeterministicTrice(g))
		rolls += 3
		if g.player2Score >= 1000 {
			g.losingScore = &g.player1Score
			fmt.Println("Player 2 win")
			return rolls
		}
	}
}

func rollOmniverse(o omniverse) omniverse {
	done := false
	for !done {
		o, done = advanceOmniverse(o)
	}
	return o
}

func countWinners(o omniverse) (int, int) {
	p1Wins := 0
	p2Wins := 0

	for universe, count := range o {
		if universe.player1Score >= 21 {
			p1Wins += count
		} else if universe.player2Score >= 21 {
			p2Wins += count
		} else {
			panic("Universe not won!")
		}

	}

	return p1Wins, p2Wins

}

func intMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Solve() (int, int) {
	p1Game := startGame(4, 2)
	rolls := playGame(p1Game)

	score := rolls * *p1Game.losingScore

	fmt.Printf("Score: %v\n", score)

	o := omniverse{
		universe{
			player1Position: 4,
			player2Position: 2,
			player1Turn:     true,
		}: 1,
	}

	o = rollOmniverse(o)

	p1Wins, p2Wins := countWinners(o)

	fmt.Printf("P1 wins %v times\nP2 wins %v times\n", p1Wins, p2Wins)

	if p1Wins > p2Wins {
		fmt.Printf("P1 wins\n")
	} else {
		fmt.Printf("P2 wins\n")
	}

	return score, intMax(p1Wins, p2Wins)
}
