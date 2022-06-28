package day8

import (
	"aoc2021/file"
	"sort"
	"strconv"
	"strings"
)

type solutionMap map[string]rune

// const (
// 	TOP segment = iota
// 	TOPLEFT
// 	TOPRIGHT
// 	MIDDLE
// 	BOTTOMLEFT
// 	BOTTOMRIGHT
// 	BOTTOM
// )

// 1 : 2
// 7 : 3
// 8 : 7
// 4 : 4

// 2 : 5
// 3 : 5
// 5 : 5

// 0 : 6
// 6 : 6
// 9 : 6

func toSortedSlice(pattern string) []rune {
	runeSlice := []rune(pattern)
	sort.Slice(runeSlice, func(a int, b int) bool { return runeSlice[a] < runeSlice[b] })
	return runeSlice
}

func parseLine(line string) ([][]rune, [][]rune) {
	line = strings.Replace(line, " | ", " ", 1)
	allpatterns := strings.Fields(line)

	sortedPatterns := make([][]rune, len(allpatterns))
	for i, p := range allpatterns {
		sortedPatterns[i] = toSortedSlice(p)
	}

	return sortedPatterns[0:10], sortedPatterns[10:14]
}

func isEeasyDigit(digit string) bool {
	switch len(digit) {
	case 2, 4, 3, 7:
		{
			return true
		}
	}
	return false
}

func CountUniqueDigits(lines []string) int {
	digits := 0

	for _, line := range lines {
		_, output := parseLine(line)

		for _, digit := range output {
			switch len(digit) {
			case 2, 4, 3, 7:
				{
					digits++
				}
			}
		}

	}

	return digits
}

func SumLines(lines []string) int {
	sum := 0

	for _, line := range lines {
		// patternMap := make(digitMap)
		solutionMap := make(solutionMap)
		pattern, output := parseLine(line)

		//Solve easy digits
		pattern1 := findOneByLength(pattern, 2)
		pattern4 := findOneByLength(pattern, 4)
		pattern7 := findOneByLength(pattern, 3)
		pattern8 := findOneByLength(pattern, 7)

		// Group by length
		len5s := findByLength(pattern, 5)
		len6s := findByLength(pattern, 6)

		pattern3 := findOneContainingAll(len5s, pattern1...) // 3 is the only 5-length to contain everything in 1

		TOPLEFT := exclude(pattern4, pattern3...)[0] // Exlude 3 from 4 to find topleft

		topLeftAndMiddle := exclude(pattern4, pattern1...)
		MIDDLE := exclude(topLeftAndMiddle, TOPLEFT)

		topLeftAndBottomleft := exclude(pattern8, pattern3...)
		BOTTOMLEFT := exclude(topLeftAndBottomleft, TOPLEFT)

		pattern9 := exclude(pattern8, BOTTOMLEFT...)
		pattern0 := exclude(pattern8, MIDDLE...)
		pattern6 := excludeSlices(len6s, pattern0, pattern9)[0]

		// 2 5
		pattern2and5 := excludeSlices(len5s, pattern3)
		pattern5 := findSliceContainingRune(pattern2and5, TOPLEFT)
		pattern2 := excludeSlices(len5s, pattern3, pattern5)[0]

		solutionMap[string(pattern0)] = '0'
		solutionMap[string(pattern1)] = '1'
		solutionMap[string(pattern2)] = '2'
		solutionMap[string(pattern3)] = '3'
		solutionMap[string(pattern4)] = '4'
		solutionMap[string(pattern5)] = '5'
		solutionMap[string(pattern6)] = '6'
		solutionMap[string(pattern7)] = '7'
		solutionMap[string(pattern8)] = '8'
		solutionMap[string(pattern9)] = '9'
		solutionMap[string(pattern0)] = '0'

		outstr := make([]rune, 4, 4)
		for i, digitSlice := range output {
			outstr[i] = solutionMap[string(digitSlice)]
		}

		n, err := strconv.Atoi(string(outstr))
		if err != nil {
			panic(err)
		}
		sum += n
	}

	return sum

}

func Solve() (int, int) {
	lines := file.ReadFile("./day8/input.txt")

	digits := CountUniqueDigits(lines)

	sum := SumLines(lines)

	// log.Printf("Unique Digits: %v Output sum: %v", digits, sum)

	return digits, sum
}

func findOneByLength(strs [][]rune, searchLen int) []rune {
	for _, pattern := range strs {
		if len(pattern) == searchLen {
			return pattern
		}
	}

	panic("Did not find string of length")
}

func findByLength(strs [][]rune, searchLen int) [][]rune {
	ret := make([][]rune, 0, len(strs))
	for _, pattern := range strs {
		if len(pattern) == searchLen {
			ret = append(ret, pattern)
		}
	}

	return ret
}

func findOneContainingAll(haystack [][]rune, needles ...rune) []rune {
	for _, pattern := range haystack {
		if containsAll(pattern, needles...) {
			return pattern
		}
	}

	panic("Could not find string containing needle")
}

func findSingleMissingRune(long string, short string) rune {
	for _, r := range long {
		if !strings.ContainsRune(short, r) {
			return r
		}
	}

	panic("Did not find missing rune")
}

func findMissingRunes(long string, short string) []rune {
	out := make([]rune, 0, len(long))
	for _, r := range long {
		if !strings.ContainsRune(short, r) {
			out = append(out, r)
		}
	}

	return out
}

func containsAll(haystack []rune, needles ...rune) bool {
	for _, r := range needles {
		if !contains(haystack, r) {
			return false
		}
	}

	return true
}

func exclude(slice []rune, filter ...rune) []rune {
	out := make([]rune, 0, len(slice))
	for _, r := range slice {
		if contains(filter, r) {
			continue
		}
		out = append(out, r)
	}

	return out
}

func contains(haystack []rune, needle rune) bool {
	for _, r := range haystack {
		if r == needle {
			return true
		}
	}

	return false
}

func findSliceContainingRune(haystack [][]rune, needle rune) []rune {
	for _, slice := range haystack {
		if contains(slice, needle) {
			return slice
		}
	}
	panic("Did not find slice")
}

func containsSlice(haystack [][]rune, needle []rune) bool {
	for _, r := range haystack {
		if string(r) == string(needle) {
			return true
		}
	}

	return false
}

func excludeSlices(slices [][]rune, filter ...[]rune) [][]rune {
	out := make([][]rune, 0, len(slices))
	for _, slice := range slices {
		if containsSlice(filter, slice) {
			continue
		}
		out = append(out, slice)
	}

	return out
}
