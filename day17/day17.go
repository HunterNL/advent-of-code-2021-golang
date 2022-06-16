package day17

import (
	"fmt"
)

type target struct {
	left, right, top, bottom int
}

func findValidX(t target) (int, iter int) {
	launchX := -1
	iterations := -1
	// found := false
	for found := false; !found; {
		launchX++
		found, iterations = isXHit(launchX, t)
		fmt.Println(found, launchX, iterations)
	}

	// for found {
	// 	launchX++
	// 	found, iterations = isXHit(launchX, t)
	// }

	return launchX, iterations
}

func findValidY(t target, maxIterations int) int {
	launchY := -1

	found := false
	for found = false; !found; {
		launchY++
		found = iterTillYHit(t, launchY, maxIterations)
		println("f, launchy", found, launchY)
	}

	for found {
		launchY++
		found = iterTillYHit(t, launchY, maxIterations)
		println("f, launchy", found, launchY)
	}

	return launchY - 1

}

func countHits(t target) int {
	leftBound, _ := findValidX(t)
	rightBound := t.right

	upperBound := -t.bottom + 1
	lowerBound := t.bottom

	hits := 0

	for x := leftBound; x <= rightBound; x++ {

		for y := lowerBound; y < +upperBound; y++ {
			if iterTillHit(t, x, y, 500) {
				hits++
			}
		}

	}

	return hits
}

func Solve() {
	input := "target area: x=209..238, y=-86..-59"
	target := parseTarget(input)
	// x := findValidX(target)
	// y := findValidY(target)

	launchX, iterations := findValidX(target)
	launchY := findValidY(target, iterations)

	hit := iterTillHit(target, launchX, launchY, 100)

	hitCount := countHits(target)

	fmt.Printf("Hit %v hitCount: %v\n", hit, hitCount)

}
