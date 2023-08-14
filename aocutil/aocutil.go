package aocutil

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(strings.TrimSpace(string(file)), "\n"), nil
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
