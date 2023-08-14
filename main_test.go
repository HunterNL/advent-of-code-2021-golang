package main

import (
	"aoc2021/aoc"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"testing"
)

func TestSolutions(t *testing.T) {
	days := getDays()

	solutionFile, err := os.ReadFile("./solutions.json")

	if errors.Is(err, fs.ErrNotExist) {
		t.Log("No solutions file found, rename solutions.json.example to solutions.json or run the main binary with --write-solutions")
	}

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	solutions, err := aoc.DecodeSolutions(solutionFile)

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

			log.SetOutput(io.Discard)
			part1, part2, err := day()
			log.SetOutput(os.Stdout)

			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					t.Logf("Did not find input file for day %v\n", dayIndex+1)
					t.SkipNow()
				} else {
					t.Log(err)
					t.FailNow()
				}

			}

			if part1 != solution.Part1 {
				t.Errorf("Expected part1's solution to be %v instead of %v", solution.Part1, part1)
			}

			if part2 != solution.Part2 {
				t.Errorf("Expected part2's solution to be %v instead of %v", solution.Part2, part2)
			}
		})
	}
}

func BenchmarkSolutions(b *testing.B) {
	log.SetOutput(io.Discard)
	for dayIndex, day := range getDays() {
		b.Run(fmt.Sprint("Day", dayIndex+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				day()
			}
		})

	}

}
