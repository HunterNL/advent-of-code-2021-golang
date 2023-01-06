package day9

import (
	"aoc2021/file"
	"sort"
	"strconv"
)

type heightMap map[xy]int

type xy struct {
	x int
	y int
}

func fill(grid *heightMap, origin xy, seen *map[xy]bool, xbound, ybound int) {
	x, y := origin.x, origin.y

	if x < 0 || x > xbound || y < 0 || y > ybound {
		return
	}

	if (*grid)[origin] == 9 {
		return
	}

	if (*seen)[origin] {
		return
	}

	(*seen)[origin] = true

	fill(grid, xy{x + 1, y}, seen, xbound, ybound)
	fill(grid, xy{x - 1, y}, seen, xbound, ybound)
	fill(grid, xy{x, y + 1}, seen, xbound, ybound)
	fill(grid, xy{x, y - 1}, seen, xbound, ybound)
}

func flood(grid *heightMap, origin xy, xbound, ybound int) int {
	seen := make(map[xy]bool)

	fill(grid, origin, &seen, xbound, ybound)

	return len(seen)
}

func topBasinSizes(grid *heightMap, lowestPoints map[xy]int, xbound, ybound int) []int {
	sizes := make([]int, len(lowestPoints))

	for c, _ := range lowestPoints {
		size := flood(grid, c, xbound, ybound)
		sizes = append(sizes, size)
	}

	sort.Ints(sizes)

	return sizes[len(sizes)-3 : len(sizes)]
}

func parseMap(lines []string) (heightMap, int, int) {
	cavemap := make(heightMap)
	for y, l := range lines {
		for x, r := range l {
			n, err := strconv.Atoi(string(r))

			if err != nil {
				panic(err)
			}

			cavemap[xy{x, y}] = n
		}
	}

	return cavemap, len(lines[0]) - 1, len(lines) - 1
}

func lowestPoints(cavemap heightMap, xbound, ybound int) heightMap {

	ret := make(heightMap)

	for coords, height := range cavemap {
		x, y := coords.x, coords.y

		//TOP
		if y > 0 {
			if cavemap[xy{x, y - 1}] <= height {
				// log.Println("Top isn't lower")
				continue
			}
		}

		//BOTTOM
		if y < ybound {
			if cavemap[xy{x, y + 1}] <= height {
				// log.Println("bottom isn't lower")
				continue
			}
		}

		//LEFT
		if x > 0 {
			if cavemap[xy{x - 1, y}] <= height {
				// log.Println("left isn't lower")
				continue
			}
		}

		//RIGHT
		if x < xbound {
			if cavemap[xy{x + 1, y}] <= height {
				// log.Println("right isn't lower")
				continue
			}
		}

		ret[coords] = height

	}

	return ret
}

func Solve() (int, int, error) {
	lines := file.ReadFile("./day9/input.txt")
	cavemap, xbound, ybound := parseMap(lines)
	points := lowestPoints(cavemap, xbound, ybound)
	sizes := topBasinSizes(&cavemap, points, xbound, ybound)

	part2 := sizes[0] * sizes[1] * sizes[2]

	part1 := riskScore(points)

	// log.Printf("Risk score: %v", part1)

	// log.Printf("Top basins: %v", part2)
	return part1, part2, nil
}

func riskScore(points heightMap) int {
	out := 0
	for _, v := range points {
		out += v + 1
	}
	return out
}
