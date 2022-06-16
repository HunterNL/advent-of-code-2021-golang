package main

import (
	"aoc2021/day1"
	"fmt"
)

type dayFunc = func() (int, int)

type dayResult struct {
	part1 int
	part2 int
}

func getDays() []dayFunc {
	return []dayFunc{
		day1.Solve,
	}
}

func main() {
	// solutionsFile, err := os.ReadFile("./solutions.json")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// solutions, err := aoc.ParseSolutions(solutionsFile)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	days := getDays()

	output := []dayResult{}

	for _, dayFunc := range days {
		part1, part2 := dayFunc()

		output = append(output, dayResult{part1, part2})
	}

	for i, day := range output {
		fmt.Printf("Day %v:\n\tPart 1:\t%v\n\tPart 2:\t%v\n", i+1, day.part1, day.part2)
	}

	// day25.Solve()
}
