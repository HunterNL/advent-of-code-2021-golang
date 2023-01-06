package day10

import (
	"aoc2021/file"
	"log"
	"sort"
)

var charMap = map[rune]rune{
	'{': '}',
	'[': ']',
	'<': '>',
	'(': ')',
}

var scoreMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}
var completionScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func parseChunk(scanner RuneScanner) {
	startChar := scanner.scan()
	expectedPair := charMap[startChar]

	if expectedPair == 0 {
		panic(ParseErr{status: Corrupt, unexpected: startChar})
	}

	for scanner.peekFor(expectedPair) != expectedPair {
		parseChunk(scanner)
	}

	scanner.consume(expectedPair)
}

func parseLine(str string) (err ParseError) {
	defer func() {
		if rerr := recover(); rerr != nil {
			perr, ok := rerr.(ParseError)
			// log.Printf("Recover error: %v\nAsserted Error:%v\nAssertionSucces:%v\nreturn error:%v", err, perr, ok, err)

			if ok {
				err = perr
			} else {
				panic(rerr)
			}
		}

	}()

	for scanner := newScanner(str); !scanner.atEnd(); {
		parseChunk(scanner)
	}

	return
}

func completeLine(str string) (fullstring []rune, completion []rune) {
	completion = make([]rune, 0)

	for {
		err := parseLine(str)
		if err == nil {
			break
		}
		if err.getStatus() != Incomplete {
			log.Printf("Encounted non-incomplete error")
			panic(err)
		}
		rune := err.getExpectedRune()
		if rune == 0 {
			panic("Did not get rune from error")
		}
		str = str + string(rune)
		completion = append(completion, rune)
	}

	return []rune(str), completion
}

func scoreCompletion(completion []rune) int {
	sum := 0
	for _, r := range completion {
		sum = sum * 5
		sum += completionScore[r]
	}

	return sum
}

func Solve() (int, int, error) {
	lines := file.ReadFile("./day10/input.txt")
	sum := 0

	for _, e := range lines {
		err := parseLine(e)
		if err.getStatus() == Corrupt {
			sum += scoreMap[err.getUnexpectedRune()]
		}
	}

	// Part 1
	// log.Printf("Syntax score: %v", sum)

	scores := make([]int, 0, len(lines))
	for _, line := range lines {
		err := parseLine(line)
		if err.getStatus() == Incomplete {
			_, completion := completeLine(line)
			scores = append(scores, scoreCompletion(completion))
		}
	}

	sort.Ints(scores)

	middle := len(scores) / 2
	// log.Printf("completion score: %v, ", scores[middle])

	return sum, scores[middle], nil

}
