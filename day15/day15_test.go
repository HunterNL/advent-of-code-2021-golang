package day15

import (
	"aoc2021/file"
	"aoc2021/grid"
	"fmt"
	"reflect"
	"testing"
)

func TestFindPath(t *testing.T) {
	file := file.ReadFile("./test-input.txt")
	g, rowSize, rowCount := parseGrid(file)

	risk := 0
	path := findPath(&g, rowSize, rowCount, 0, 99)

	for _, i := range path {
		x, y := grid.ToXY(i, rowSize)
		fmt.Printf("X: %v, Y:%v\n", x, y)
		risk += int(g[i])
	}

	if risk != 40 {
		t.Errorf("Expected a risk of 40, got %v instead", risk)
	}
}

func TestMod(t *testing.T) {
	// if(mod(0) != 1) {
	// 	t.Error("Expected 2")
	// }
	if inc(1, 1) != 2 {
		t.Errorf("Expected 2 got %v\n", inc(1, 1))
	}
	if inc(8, 1) != 9 {
		t.Errorf("Expected 9 got %v\n", inc(1, 1))
	}
	if inc(9, 1) != 1 {
		t.Errorf("Expected 1 got %v\n", inc(1, 1))
	}
}

func TestRepeat(t *testing.T) {
	file := file.ReadFile("./test-input.txt")
	g, rowSize, _ := parseGrid(file)
	r := repeatRight(&g, rowSize)

	part := r[:100]
	if !reflect.DeepEqual(part, []uint8{1, 1, 6, 3, 7, 5, 1, 7, 4, 2, 2, 2, 7, 4, 8, 6, 2, 8, 5, 3, 3, 3, 8, 5, 9, 7, 3, 9, 6, 4, 4, 4, 9, 6, 1, 8, 4, 1, 7, 5, 5, 5, 1, 7, 2, 9, 5, 2, 8, 6, 1, 3, 8, 1, 3, 7, 3, 6, 7, 2, 2, 4, 9, 2, 4, 8, 4, 7, 8, 3, 3, 5, 1, 3, 5, 9, 5, 8, 9, 4, 4, 6, 2, 4, 6, 1, 6, 9, 1, 5, 5, 7, 3, 5, 7, 2, 7, 1, 2, 6}) {
		t.Error("Match failed")
	}
}
