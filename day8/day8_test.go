package day8

import (
	"aoc2021/file"
	"testing"
)

func TestCountUniqueDigits(t *testing.T) {
	lines := file.ReadFile("./test_input.txt")

	digits := CountUniqueDigits(lines)

	if digits != 26 {
		t.Errorf("Expected 26, got %v", digits)
	}
}

func TestLinesSum(t *testing.T) {
	lines := file.ReadFile("./test_input.txt")

	digits := SumLines(lines)

	if digits != 61229 {
		t.Errorf("Expected 61229, got %v", digits)
	}
}
