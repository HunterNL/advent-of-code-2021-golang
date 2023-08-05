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
	"io"
	"log"
	"os"
	"time"
)

type dayFunc = func() (int, int, error)

type dayResult struct {
	part1 int
	part2 int
	err   error
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
	days := getDays()

	output := []dayResult{}
	durations := []time.Duration{}

	// day24.Solve()
	// return

	log.SetOutput(io.Discard)

	for _, dayFunc := range days {
		start := time.Now()
		part1, part2, err := dayFunc()

		durations = append(durations, time.Since(start))
		output = append(output, dayResult{part1, part2, err})
	}

	log.SetOutput(os.Stdout)

	var totalDuration int64 = 0
	for i, day := range output {
		totalDuration += int64(durations[i])
		if day.err != nil {
			log.Printf("Day %2v:        Error:%v\n", i, day.err)
		} else {
			log.Printf("Day %2v: %4v ms\n", i+1, int64(durations[i]/time.Millisecond))
		}

	}

	log.Printf("Total duration: %vms\n", totalDuration/int64(time.Millisecond))

}
