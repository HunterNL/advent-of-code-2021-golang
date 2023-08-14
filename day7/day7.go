package day7

import (
	"aoc2021/intmath"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//https://adventofcode.com/2021/day/7

func average(list []int) int {
	sum := 0
	for _, n := range list {
		sum += n
	}

	return int(math.Round(float64(sum) / float64(len(list))))
}

func mean(list []int) int {
	sort.Ints(list)
	return list[len(list)/2]
}

func CalcFuel(crabPos []int) int {
	av := mean(crabPos)
	fuel := 0

	for _, crab := range crabPos {
		fuel += intmath.Distance(av, crab)
	}

	return fuel
}

func fuelCost(m int) int {
	return (m * (m + 1)) / 2
}

func CalcFuelExpensive(crabPos []int) int {
	width := intmath.Max(crabPos)
	minFuel := math.MaxInt

	for i := 0; i <= width; i++ {
		fuel := 0
		for _, crab := range crabPos {
			fuel += fuelCost(intmath.Distance(i, crab))
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}

func Solve() (int, int, error) {
	file, err := os.ReadFile("./day7/input.txt")
	if err != nil {
		return -1, -1, err
	}

	strs := strings.Split(string(file), ",")
	nums := make([]int, len(strs))

	for i, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}

	part1 := CalcFuel(nums)
	part2 := CalcFuelExpensive(nums)

	// log.Printf("Part 1: Crab movement: %v\n", part1)
	// log.Printf("Part 2: Crab movement: %v\n", part2)

	return part1, part2, nil
}
