package day1

import (
	"os"
	"strconv"
	"strings"
)

func Solve() (int, int, error) {
	file, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		return -1, -1, err
	}

	str := string(file)

	rawstrings := strings.Split(strings.TrimSpace(str), "\n")
	numlist := make([]int, len(rawstrings))
	for i, s := range rawstrings {
		numlist[i], err = strconv.Atoi(s)
		if err != nil {
			return -1, -1, err
		}
	}

	reduceCount := 0
	lastSum := numlist[0] + numlist[1] + numlist[2]

	// Part 1
	var lastNum int

	for i, n := range numlist {
		if i == 0 {
			lastNum = n
			continue
		}
		if n > lastNum {
			reduceCount++
		}
		lastNum = n
	}

	// Part 2
	reduceCount2 := 0

	for i := 1; i < len(numlist)-2; i++ {
		n := numlist[i] + numlist[i+1] + numlist[i+2]
		if n > lastSum {
			reduceCount2++
		}
		lastSum = n
	}

	// log.Printf("%d out of %d", reduceCount, len(numlist))

	return reduceCount, reduceCount2, nil

}
