package aoc

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	return strings.Split(strings.TrimSpace(string(file)), "\n")
}

func CommaStringToInts(str string) []int {
	substrings := strings.Split(str, ",")
	ret := make([]int, len(substrings))

	for i, s := range substrings {
		int, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ret[i] = int
	}

	return ret
}
