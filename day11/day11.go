package day11

import (
	"aoc2021/file"
	"log"
	"strconv"
	"strings"
)

func parseGrid(lines []string) [100]uint8 {
	str := strings.Join(lines, "")
	str = strings.ReplaceAll(str, "\n", "")

	grid := [100]uint8{}

	for i, r := range str {
		n, err := strconv.ParseInt(string(r), 10, 8)
		if err != nil {
			panic(err)
		}

		grid[i] = uint8(n)
	}

	return grid
}

func validNeighbours(index int) []int {
	x, y := index%10, index/10
	neighbours := make([]int, 0, 8)

	top := y > 0
	bottom := y < 9

	left := x > 0
	right := x < 9

	if top {
		if left {
			neighbours = append(neighbours, index-11)
		}
		neighbours = append(neighbours, index-10)
		if right {
			neighbours = append(neighbours, index-9)
		}
	}

	if left {
		neighbours = append(neighbours, index-1)
	}

	if right {
		neighbours = append(neighbours, index+1)
	}

	if bottom {
		if left {
			neighbours = append(neighbours, index+9)
		}
		neighbours = append(neighbours, index+10)
		if right {
			neighbours = append(neighbours, index+11)
		}
	}

	return neighbours

}

func flash(grid *[100]uint8, seen *map[int]bool, index int) {
	(*seen)[index] = true

	neigghbours := validNeighbours(index)

	for _, v := range neigghbours {
		if v >= 0 && v < 100 {
			grid[v] = (*grid)[v] + 1
			if grid[v] > 9 && !(*seen)[v] {
				flash(grid, seen, v)
			}
		}
	}

}

func step(gridref *[100]uint8) int {
	flashed := make(map[int]bool)

	for i, n := range gridref {
		gridref[i] = n + 1
	}

	for i, n := range gridref {
		if n > 9 && !flashed[i] {
			flash(gridref, &flashed, i)
		}
	}

	for n := range flashed {
		gridref[n] = 0
	}

	return len(flashed)
}

func countFlashes(grid *[100]uint8, iterations int) int {
	// log.Println("Initial state")
	// log.Println(toString(grid))
	sum := 0
	for i := 0; i < iterations; i++ {
		sum += step(grid)
		// log.Printf("After step %v\n", i+1)
		// log.Println(toString(grid))
	}

	return sum
}

func findSyncStep(grid *[100]uint8) int {
	iterations := 0
	for flashCount := 0; flashCount < 100; iterations++ {
		flashCount = step(grid)
	}

	return iterations
}

func toString(grid *[100]uint8) string {
	runes := [100]rune{}
	str := ""

	for i, n := range grid {
		runes[i] = []rune(strconv.Itoa(int(n)))[0]
	}

	for i := 0; i < 10; i++ {
		str += string(runes[i*10 : i*10+10])
		str += "\n"
	}

	return str
}

func Solve() (int, int, error) {
	lines := file.ReadFile("./day11/input.txt")
	grid := parseGrid(lines)

	flashes := countFlashes(&grid, 100)
	grid = parseGrid(lines)
	syncStep := findSyncStep(&grid)

	log.Printf("Flash count: %v sync step: %v", flashes, syncStep)

	return flashes, syncStep, nil

}
