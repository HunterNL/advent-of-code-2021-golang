package day5

import (
	"os"
	"strconv"
	"strings"
)

const GRID_SIZE = 1000

type grid map[int]int

type vec2 = struct{ x, y int }

type line struct {
	a, b vec2
}

func parseVec2(str string) vec2 {
	nsec := strings.Split(str, ",")

	x, err := strconv.Atoi(nsec[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(nsec[1])

	return vec2{x, y}
}

func parseLine(str string) line {
	vectors := strings.Split(str, " -> ")

	return line{
		a: parseVec2(vectors[0]),
		b: parseVec2(vectors[1]),
	}
}

func parseFile(str string) []line {
	textLines := strings.Split(str, "\n")
	textLines = textLines[:len(textLines)-1]
	lines := make([]line, len(textLines))
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

func Solve() (int, int, error) {
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
		cur := line.a
		tar := line.b
		dirX, dirY := dirTo(int(cur.x), int(cur.y), int(tar.x), int(tar.y))
		if line.Isorthogonal() {
			part1Grid[cur.x*GRID_SIZE+cur.y]++
		}

		part2Grid[cur.x*GRID_SIZE+cur.y]++
		for cur != tar {
			cur.x += dirX
			cur.y += dirY

			part2Grid[cur.x*GRID_SIZE+cur.y]++

			if line.Isorthogonal() {
				part1Grid[cur.x*GRID_SIZE+cur.y]++
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

	// log.Printf("Intersections: %v", intersections)

	return part1, part2, nil
}

func (l line) Isorthogonal() bool {
	if l.a.x == l.b.x {
		return true
	}

	if l.a.y == l.b.y {
		return true
	}

	return false
}
