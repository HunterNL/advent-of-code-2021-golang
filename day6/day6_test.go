package day6

import "testing"

func TestFish(t *testing.T) {
	fishCount := CalcLanternFish([]int{3, 4, 3, 1, 2}, 80)

	if fishCount != 5934 {
		t.Error("Expected fishcount of 5934, got", fishCount)
	}
}
