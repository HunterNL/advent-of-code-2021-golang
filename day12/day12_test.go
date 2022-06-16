package day12

import (
	"aoc2021/file"
	"testing"
)

func TestCountRoutes(t *testing.T) {
	tc := map[string]int{
		"./test-input1.txt": 10,
		"./test-input2.txt": 19,
		"./test-input3.txt": 226,
	}

	for path, expected := range tc {

		lines := file.ReadFile(path)
		g := parseGraph(lines)

		count := countRoutes(&g)

		if count != expected {
			t.Errorf("Expected %v routes, got %v", expected, count)
		}
	}

}
func TestCountComplexRoutes(t *testing.T) {
	tc := map[string]int{
		"./test-input1.txt": 36,
		"./test-input2.txt": 103,
		"./test-input3.txt": 3509,
	}

	for path, expected := range tc {

		lines := file.ReadFile(path)
		g := parseGraph(lines)

		count := countComplexRoutes(&g)

		if count != expected {
			t.Errorf("Expected %v routes, got %v", expected, count)
		}
	}

}
