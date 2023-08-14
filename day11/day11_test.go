package day11

import (
	"aoc2021/file"
	"testing"
)

func TestFlashCount(t *testing.T) {
	lines, err := file.ReadFile("./test-input.txt")
	if err != nil {
		t.Error(err)
	}
	grid := parseGrid(lines)

	flashCount10 := countFlashes(&grid, 10)

	grid = parseGrid(lines)
	flashCount100 := countFlashes(&grid, 100)

	if flashCount10 != 204 {
		t.Errorf("Exepected 204 flashes, got %v", flashCount10)
	}

	if flashCount100 != 1656 {
		t.Errorf("Exepected 1656 flashes, got %v", flashCount100)
	}
}

func TestFindSyncStep(t *testing.T) {
	lines, err := file.ReadFile("./test-input.txt")
	if err != nil {
		t.Error(err)
	}
	grid := parseGrid(lines)

	syncStep := findSyncStep(&grid)

	if syncStep != 195 {
		t.Errorf("Expected sync on step 195, but got step %v", syncStep)
	}
}
