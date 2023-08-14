package day25

import (
	"aoc2021/file"
	"log"
	"strings"
)

type cucumber rune

type grid [][]cucumber

func makeGrid(x, y int) grid {
	g := make([][]cucumber, y)
	for i := 0; i < y; i++ {
		g[i] = make([]cucumber, x)

		for i2 := 0; i2 < x; i2++ {
			g[i][i2] = cucumber('.')
		}
	}
	return g
}

func parseLine(str string) []cucumber {
	out := make([]cucumber, len(str))
	for i, c := range str {
		out[i] = cucumber(c)
	}
	return out
}

func parseGrid(strs []string) (grid, int, int) {
	x := len(strs[0])
	y := len(strs)

	grid := make(grid, y)
	for i, l := range strs {
		grid[i] = parseLine(l)

		// log.Printf("\n%v\n", grid)
	}
	return grid, x, y

}

func (g grid) String() string {
	sb := strings.Builder{}

	for Y, line := range g {
		for _, cucumber := range line {
			sb.WriteRune(rune(cucumber))
		}

		if Y < len(g)-1 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

func (g *grid) replace(x, y int, from, to cucumber) {
	for iY := 0; iY < y; iY++ {
		for iX := 0; iX < x; iX++ {
			if (*g)[iY][iX] == from {
				(*g)[iY][iX] = to
			}
		}
	}
}

func (g *grid) step(x, y int) (grid, bool) {
	moved := false
	// intermediateHorizontal := makeGrid(x, y)
	intermediate := makeGrid(x, y)
	// intermediateVertical := makeGrid(x, y)
	out := makeGrid(x, y)

	for iY := 0; iY < y; iY++ {
		for iX := 0; iX < x; iX++ {
			if (*g)[iY][iX] == cucumber('>') {
				if (*g)[iY][(iX+1)%x] == cucumber('.') {
					// if (*g)[iY][(iX+1)%x] == cucumber('.') && out[iY][(iX+1)%x] == cucumber('.') {
					intermediate[iY][(iX+1)%x] = cucumber('>')
					out[iY][(iX+1)%x] = cucumber('}')
					// out[iY][(iX+1)%x] = cucumber('>')
					moved = true
				} else {
					intermediate[iY][iX] = cucumber('>')
					out[iY][iX] = cucumber('>')
				}
			}
		}
	}

	out.replace(x, y, '}', '>')

	// log.Println(intermediate)

	for iY := 0; iY < y; iY++ {
		for iX := 0; iX < x; iX++ {
			if (*g)[iY][iX] == cucumber('v') {
				if (*g)[(iY+1)%y][iX] != cucumber('v') && intermediate[(iY+1)%y][iX] == cucumber('.') {
					out[(iY+1)%y][iX] = cucumber('|')
					moved = true
				} else {
					out[iY][iX] = cucumber('v')
				}
			}
		}
	}

	out.replace(x, y, '|', 'v')

	// for iY := 0; iY < y; iY++ {
	// 	for iX := 0; iX < x; iX++ {
	// 		if out[iY][iX] == cucumber('|') {
	// 			out[iY][iX] = cucumber('v')
	// 		}
	// 	}
	// }

	// log.Println(out)

	return out, moved
}

func Solve() (int, int, error) {
	lines, err := file.ReadFile("./day25/input.txt")
	if err != nil {
		return -1, -1, err
	}
	grid, x, y := parseGrid(lines)

	moved := true
	steps := 0
	for moved {
		steps++
		grid, moved = grid.step(x, y)
	}

	log.Printf("Steps: %v\n", steps)

	return steps, -1, nil
}
