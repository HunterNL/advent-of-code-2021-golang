package day7

import "testing"

func TestCrab(t *testing.T) {
	fuel := CalcFuel([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14})

	if fuel != 37 {
		t.Errorf("Expected fuel to be 37 instead of %v", fuel)
	}
}
