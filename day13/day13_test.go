package day13

import (
	"os"
	"testing"
)

func TestCountDots(t *testing.T) {
	bytes, err := os.ReadFile("./test-input.txt")
	if err != nil {
		panic(err)
	}
	str := string(bytes)

	grid, folds := parseFold(str)

	applyFold(&grid, folds[0])

	count := countFolds(&grid)

	if count != 17 {
		t.Errorf("Expected 17 dots, got %v", count)
	}

}
