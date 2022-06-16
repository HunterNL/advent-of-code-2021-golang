package day5

import (
	"aoc2021/vec2"
	"os"
	"strconv"
	"strings"
)

type grid = map[vec2.Vec2]int

func parseVec2(str string) vec2.Vec2 {
	nsec := strings.Split(str, ",")

	x, err := strconv.Atoi(nsec[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(nsec[1])

	return vec2.Vec2{X: float64(x), Y: float64(y)}
}

func parseLine(str string) vec2.Line {
	vectors := strings.Split(str, " -> ")

	return vec2.Line{
		A: parseVec2(vectors[0]),
		B: parseVec2(vectors[1]),
	}
}

func parseFile(str string) []vec2.Line {
	textLines := strings.Split(str, "\n")
	textLines = textLines[:len(textLines)-1]
	lines := make([]vec2.Line, len(textLines))
	for i, v := range textLines {
		lines[i] = parseLine(v)
	}

	return lines
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return x
}

func dirTo(x1, y1, x2, y2 int) (int, int) {
	x := sign(x2 - x1)
	y := sign(y2 - y1)
	return x, y
}

func countIntersections(g *grid) int {
	intersections := 0
	for _, n := range *g {
		if n >= 2 {
			intersections++
		}
	}
	return intersections
}

func Solve() (int, int) {
	file, err := os.ReadFile("./day5/input.txt")
	if err != nil {
		panic(err)
	}

	lines := parseFile(string(file))

	// Part 1
	part1Grid := make(grid)
	part2Grid := make(grid)

	// for _, line := range lines {
	// 	if(line.Isorthogonal())
	// 		steps := line.A.StepsTo(&line.B)
	// 		for _, step := range steps {
	// 			grid[step]++
	// 		}
	// 	}
	// }

	for _, line := range lines {
		cur := line.A
		tar := line.B
		dirX, dirY := dirTo(int(cur.X), int(cur.Y), int(tar.X), int(tar.Y))
		if line.Isorthogonal() {
			part1Grid[cur]++
		}

		part2Grid[cur]++
		for cur != tar {
			cur.X += float64(dirX)
			cur.Y += float64(dirY)

			part2Grid[cur]++

			if line.Isorthogonal() {
				part1Grid[cur]++
			}
		}
	}

	part1 := countIntersections(&part1Grid)
	part2 := countIntersections(&part2Grid)

	// intersections := 0
	// for _, n := range grid {
	// 	if n >= 2 {
	// 		intersections++
	// 	}
	// }

	// fmt.Printf("Intersections: %v", intersections)

	return part1, part2
}
