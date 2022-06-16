package main

import (
	"aoc2021/aoc"
	"fmt"
	"os"
	"testing"
)

func TestSolutions(t *testing.T) {
	days := getDays()

	solutionFile, err := os.ReadFile("./solutions.json")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	solutions, err := aoc.ParseSolutions(solutionFile)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	for dayIndex, day := range days {
		t.Run(fmt.Sprint("Day", dayIndex+1), func(t *testing.T) {
			solution, foundSolution := solutions[dayIndex+1]

			if !foundSolution {
				t.Logf("Did not find solution for day %v", dayIndex+1)
				t.SkipNow()
			}

			part1, part2 := day()

			if part1 != solution.Part1 {
				t.Errorf("Expected part1's solution to be %v instead of %v", solution.Part1, part1)
			}

			if part2 != solution.Part2 {
				t.Errorf("Expected part2's solution to be %v instead of %v", solution.Part2, part2)
			}
		})
	}
}
