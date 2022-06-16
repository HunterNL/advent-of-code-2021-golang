package day2

import (
	"aoc2021/aocutil"
	"strconv"
	"strings"
)

func splitCommand(str string) (string, int) {
	split := strings.Fields(str)
	num, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return split[0], num
}

func Solve() {
	horizontal := 0
	depth := 0
	aim := 0

	commands := aocutil.ReadFile("./day2/input.txt")

	// Part 1
	// for _, str := range commands {
	// 	command, count := splitCommand(str)
	// 	switch command {
	// 	case "forward":
	// 		horizontal += count
	// 	case "up":
	// 		depth -= count
	// 	case "down":
	// 		depth += count
	// 	}
	// }

	for _, str := range commands {
		command, count := splitCommand(str)
		switch command {
		case "forward":
			horizontal += count
			depth += (aim * count)
		case "up":
			aim -= count
		case "down":
			aim += count
		}
	}

	println(horizontal * depth)

}
