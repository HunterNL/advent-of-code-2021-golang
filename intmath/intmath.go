package intmath

import "math"

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Distance(a, b int) int {
	return Abs(a - b)
}

func Min(a []int) int {
	min := math.MaxInt
	for _, n := range a {
		if n < min {
			min = n
		}
	}

	return min
}
func Max(a []int) int {
	min := math.MinInt
	for _, n := range a {
		if n > min {
			min = n
		}
	}

	return min
}
