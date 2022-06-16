package day13

import (
	"aoc2021/file"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type xy struct {
	x, y int
}

type direction int

const (
	Unknown direction = iota
	X
	Y
)

type fold struct {
	n   int
	dir direction
}

func parseFold(input string) (map[xy]bool, []fold) {
	dotsLines, foldLines := file.SplitOnce(input, "\n\n")

	dots := make(map[xy]bool)
	folds := make([]fold, 0)

	for _, l := range strings.Split(dotsLines, "\n") {
		xs, ys := file.SplitOnce(l, ",")

		x, err := strconv.Atoi(xs)
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(ys)
		if err != nil {
			panic(err)
		}

		dots[xy{x, y}] = true
	}

	for _, l := range strings.Split(foldLines, "\n") {
		data := strings.TrimPrefix(l, "fold along ")
		dirstr := []rune(data)[0]
		dirn, err := strconv.Atoi(strings.TrimLeft(data, "xy="))
		if err != nil {
			panic(err)
		}

		var dir direction
		if dirstr == 'x' {
			dir = X
		}
		if dirstr == 'y' {
			dir = Y
		}

		folds = append(folds, fold{n: dirn, dir: dir})
	}

	return dots, folds
}

func applyFold(grid *map[xy]bool, f fold) {
	foldLine, foldN := f.dir, f.n

	for gridxy := range *grid {
		x, y := gridxy.x, gridxy.y

		if foldLine == X && x > foldN {
			(*grid)[xy{x: (foldN - (x - foldN)), y: y}] = true
			delete(*grid, gridxy)
		}

		if foldLine == Y && y > foldN {
			(*grid)[xy{x: x, y: (foldN - (y - foldN))}] = true
			delete(*grid, gridxy)
		}
	}

}

func gridBounds(grid *map[xy]bool) (int, int) {
	xbound := 0
	ybound := 0

	for xy := range *grid {
		if xy.x > xbound {
			xbound = xy.x
		}
		if xy.y > ybound {
			ybound = xy.y
		}
	}

	return xbound, ybound
}

func countFolds(grid *map[xy]bool) int {
	return len(*grid)
}

func printGrid(grid *map[xy]bool) string {

	xbound, ybound := gridBounds(grid)

	size := (xbound + 1) * (ybound + 1)

	sb := strings.Builder{}
	for i := 0; i < size; i++ {
		x := i%xbound + 1
		y := i / xbound

		println(x, y)

		if ((*grid)[xy{x: x, y: y}]) {
			sb.WriteString("#")
		} else {
			sb.WriteString(".")
		}

		if x == xbound {
			sb.WriteString("\n")
		}
	}

	fmt.Printf("Size: %v X: %v Y: %v", size, xbound, ybound)

	return sb.String()
}

func Solve() (int, int) {
	bytes, err := os.ReadFile("./day13/input.txt")
	if err != nil {
		panic(err)
	}
	str := string(bytes)

	grid1, folds1 := parseFold(str)
	applyFold(&grid1, folds1[0])
	part1 := countFolds(&grid1)

	grid, folds := parseFold(str)

	for _, f := range folds {
		applyFold(&grid, f)
	}

	count := countFolds(&grid)

	fmt.Printf("Got %v dots left\n", count)

	gridstr := printGrid(&grid)

	print(gridstr)

	return part1, -1

	// os.WriteFile("output.txt", []byte(gridstr), fs.ModeAppend)
}
