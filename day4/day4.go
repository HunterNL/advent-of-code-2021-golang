package day4

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type board struct {
	rows      [10][5]int
	nums      [25]int
	score     int
	hitIndex  int
	hitNumber int
}

func sumIntSlice(slice []int) int {
	var sum int
	for _, n := range slice {
		sum += n
	}

	return sum
}

func contains(haystack *[5]int, needle int) bool {
	for _, n := range haystack {
		if n == needle {
			return true
		}
	}

	return false
}

// aaargh why
func contains2(haystack *[25]int, needle int) bool {
	for _, n := range haystack {
		if n == needle {
			return true
		}
	}

	return false
}

func rowHitIndex(row *[5]int, sequence []int) int {
	hits := 0
	for i, n := range sequence {
		if contains(row, n) {
			hits++
			if hits == 5 {
				return i
			}
		}
	}

	panic("Row is never completed")
}

func boardHitIndex(b *board, sequence []int) {
	hitIndex := math.MaxInt

	for _, row := range b.rows {
		hit := rowHitIndex(&row, sequence)
		if hit < hitIndex {
			hitIndex = hit
		}
	}

	b.hitIndex = hitIndex
	b.hitNumber = sequence[hitIndex]
}

func parseBoard(str string) *board {
	board := new(board)

	numstr := strings.ReplaceAll(str, "\n", " ")
	nums := strings.Fields(numstr)
	ints := new([25]int)

	for i, num := range nums {
		n2, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		ints[i] = n2
	}

	for i := 0; i < 25; i++ {
		board.rows[i/5][i%5] = ints[i]
		board.rows[5+i/5][i%5] = ints[(i*5+i/5)%25] // 0 5 10 15 20 1 6
	}

	board.nums = *ints
	return board
}

func Solve() (int, int) {
	file, err := os.ReadFile("./day4/input.txt")
	if err != nil {
		panic(err)
	}

	// String wrestling
	texts := strings.Split(string(file), "\n\n")
	sequenceString := strings.Split(texts[0], ",")

	// Parsing bingo sequence
	sequence := make([]int, len(sequenceString))
	for i, s := range sequenceString {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		sequence[i] = n
	}

	// Parsing boards
	boardCount := len(texts) - 1
	boards := make([]*board, boardCount)
	for i := 1; i <= boardCount; i++ {
		boards[i-1] = parseBoard(texts[i])
	}

	// Finding lowest board
	lowestBoard := &board{hitIndex: math.MaxInt}
	for _, b := range boards {
		boardHitIndex(b, sequence)

		if b.hitIndex < lowestBoard.hitIndex {
			lowestBoard = b
		}
	}

	//
	highestBoard := &board{hitIndex: math.MinInt}
	for _, b := range boards {
		boardHitIndex(b, sequence)

		if b.hitIndex > highestBoard.hitIndex {
			highestBoard = b
		}
	}

	unrolledLowest := make([]int, 0, len(sequence))
	unrolledHighest := make([]int, 0, len(sequence))

	for i := len(sequence) - 1; i > lowestBoard.hitIndex; i-- {
		if contains2(&lowestBoard.nums, sequence[i]) {
			unrolledLowest = append(unrolledLowest, sequence[i])
		}
	}

	for i := len(sequence) - 1; i > highestBoard.hitIndex; i-- {
		if contains2(&highestBoard.nums, sequence[i]) {
			unrolledHighest = append(unrolledHighest, sequence[i])
		}
	}

	sumLowest := sumIntSlice(unrolledLowest)
	sumHighest := sumIntSlice(unrolledHighest)

	// fmt.Printf("Winner Sum: %v Hit Number: %v Product: %v\n", sumLowest, lowestBoard.hitNumber, sumLowest*lowestBoard.hitNumber)
	// fmt.Printf("Looser Sum: %v Hit Number: %v Product: %v\n", sumHighest, highestBoard.hitNumber, sumHighest*highestBoard.hitNumber)

	part1 := sumLowest * lowestBoard.hitNumber
	part2 := sumHighest * highestBoard.hitNumber

	return part1, part2

}

// https://adventofcode.com/2021/day/4
