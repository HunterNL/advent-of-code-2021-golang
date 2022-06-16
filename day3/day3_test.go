package day3

import (
	"aoc2021/aocutil"
	"os"
	"strconv"
	"testing"
)

func TestPart2(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	t.Log(dir)
	lines := aocutil.ReadFile("./test_input.txt")

	o2str := part2(lines, false)
	cooStr := part2(lines, true)

	o2Rating, err := strconv.ParseInt(o2str, 2, 64)
	if err != nil {
		panic(err)
	}

	if o2str != "10111" {
		t.Error("str not 10111")
	}

	if o2Rating != 23 {
		t.Error("O2 rating not 23")
	}

	if cooStr != "01010" {
		t.Error("coo rating not 01010")
	}
}
