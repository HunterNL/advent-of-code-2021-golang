package day3

import (
	"aoc2021/aocutil"
	"fmt"
	"strconv"
)

func countBits(bits []string) []int {
	bitCount := make([]int, len(bits[0]))

	for _, line := range bits {
		for i, r := range line {
			if r == '1' {
				bitCount[i]++
			}
		}
	}

	return bitCount
}

func part2(lines []string, invert bool) string {
	byteSize := len(lines[0])
	var keepRune rune
	var halfPoint float32

	for bitIndex := 0; bitIndex < byteSize; bitIndex++ {
		newLines := make([]string, 0, len(lines))
		bitcount := countBits(lines)
		halfPoint = 0.5 * float32(len(lines))

		fmt.Printf("HP: %T %v", halfPoint, halfPoint)

		if float32(bitcount[bitIndex]) == halfPoint {
			keepRune = '1'
			if invert {
				keepRune = '0'
			}
		} else {
			if (float32(bitcount[bitIndex]) > halfPoint) != invert {
				keepRune = '1'
			} else {
				keepRune = '0'
			}
		}

		for _, line := range lines {
			if line[bitIndex] == byte(keepRune) {
				newLines = append(newLines, line)
			}
		}

		if len(newLines) == 1 {
			return newLines[0]
		}

		lines = newLines
	}

	panic("aaaaa")
}

func Solve() {
	lines := aocutil.ReadFile("./day3/input.txt")

	byteMiddle := len(lines) / 2

	bits := countBits(lines)

	// Part 1
	gamma := 0
	epsilon := 0

	for _, b := range bits {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if b > byteMiddle {
			gamma++
		} else {
			epsilon++
		}

	}

	fmt.Printf("Gamma: %v Epsilon: %v Multiplied: %v\n", gamma, epsilon, gamma*epsilon)

	// Part 2
	oxyRating, err := strconv.ParseInt(part2(lines, false), 2, 64)
	if err != nil {
		panic(err)
	}

	cooRating, err := strconv.ParseInt(part2(lines, true), 2, 64)
	if err != nil {
		panic(err)
	}

	// Oxygen generator rating
	fmt.Printf("Oxy rating: %v cooRating: %v multiplied: %v", oxyRating, cooRating, oxyRating*cooRating)

}
