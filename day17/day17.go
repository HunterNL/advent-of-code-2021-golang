package day17

import (
	"log"
	"os"
	"strings"
)

type target struct {
	left, right, top, bottom int
}

func iterTillDown(launchY int) int {
	dy := launchY
	yPos := 0
	for dy >= 0 {
		yPos += dy
		dy = dy - 1
		log.Println(yPos)
	}

	return yPos

}

func findValidX(t target) (int, iter int) {
	launchX := -1
	iterations := -1
	// found := false
	for found := false; !found; {
		launchX++
		found, iterations = isXHit(launchX, t)
		log.Println(found, launchX, iterations)
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
		// println("f, launchy", found, launchY)
	}

	for found {
		launchY++
		found = iterTillYHit(t, launchY, maxIterations)
		// println("f, launchy", found, launchY)
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

func Solve() (int, int, error) {
	file, err := os.ReadFile("./day17/input.txt") // You saw nothing :>
	if err != nil {
		return -1, -1, err
	}
	input := strings.TrimSpace(string(file))

	target := parseTarget(input)

	maxY := iterTillDown(85)

	launchX, iterations := findValidX(target)
	launchY := findValidY(target, iterations)

	hit := iterTillHit(target, launchX, launchY, 100)

	hitCount := countHits(target)

	log.Printf("Hit %v hitCount: %v\n", hit, hitCount)

	return maxY, hitCount, nil

}
