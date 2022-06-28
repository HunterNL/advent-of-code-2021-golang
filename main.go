package main

import (
	"aoc2021/day1"
	"aoc2021/day10"
	"aoc2021/day11"
	"aoc2021/day12"
	"aoc2021/day13"
	"aoc2021/day14"
	"aoc2021/day15"
	"aoc2021/day16"
	"aoc2021/day17"
	"aoc2021/day18"
	"aoc2021/day19"
	"aoc2021/day2"
	"aoc2021/day20"
	"aoc2021/day21"
	"aoc2021/day22"
	"aoc2021/day23"
	"aoc2021/day24"
	"aoc2021/day25"
	"aoc2021/day3"
	"aoc2021/day4"
	"aoc2021/day5"
	"aoc2021/day6"
	"aoc2021/day7"
	"aoc2021/day8"
	"aoc2021/day9"
	"log"
)

type dayFunc = func() (int, int)

type dayResult struct {
	part1 int
	part2 int
}

func getDays() []dayFunc {
	return []dayFunc{
		day1.Solve,
		day2.Solve,
		day3.Solve,
		day4.Solve,
		day5.Solve,
		day6.Solve,
		day7.Solve,
		day8.Solve,
		day9.Solve,
		day10.Solve,
		day11.Solve,
		day12.Solve,
		day13.Solve,
		day14.Solve,
		day15.Solve,
		day16.Solve,
		day17.Solve,
		day18.Solve,
		day19.Solve,
		day20.Solve,
		day21.Solve,
		day22.Solve,
		day23.Solve,
		day24.Solve,
		day25.Solve,
	}
}

func main() {
	// solutionsFile, err := os.ReadFile("./solutions.json")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	//
	// solutions, err := aoc.ParseSolutions(solutionsFile)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	days := getDays()

	output := []dayResult{}

	for _, dayFunc := range days {
		part1, part2 := dayFunc()

		output = append(output, dayResult{part1, part2})
	}

	for i, day := range output {
		log.Printf("Day %v:\n\tPart 1:\t%v\n\tPart 2:\t%v\n", i+1, day.part1, day.part2)
	}

	// day25.Solve()
}
