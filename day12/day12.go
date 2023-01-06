package day12

import (
	"aoc2021/file"
	"log"
	"strings"
)

// Caves are a single bit in a uint32, large status being indicated by LARGE_CAVE_BIT
type cave uint32

// Routes are a combination of cave bits
type route uint32

func isUpper(str string) bool {
	return strings.ToUpper(str) == str
}

func visited(needle cave, haystack route) bool {
	return needle&cave(haystack) > 0
}

func travel(g Graph, start cave, behind route, routeCount *int) {
	path := behind | route(start)

	if start == g.getStringAsBoolPointer("end") {
		*routeCount++
		return
	}

	neighbours := g.getNeighbours(start)

	for _, n := range neighbours {
		if n&LARGE_CAVE_BIT > 0 {
			travel(g, n, path, routeCount)
		} else {
			if !visited(n, path) {
				travel(g, n, path, routeCount)
			}
		}

	}
}

// Returns if we travel to `to`
// Note that unlike "cave" bool pointers, this plain pointer does update `visitedAnySmallCaveTwice`
func canComplexTravel(to cave, behind route, startPtr cave, visitedAnySmallCaveTwice *bool) bool {
	// Probit reentering starting cave
	if to == startPtr {
		return false
	}

	isLarge := to&LARGE_CAVE_BIT > 0

	if isLarge {
		return true // Large caves are always fine
	}

	if visited(to, behind) {
		if *visitedAnySmallCaveTwice {
			return false // Prohibit visiting twice
		} else {
			*visitedAnySmallCaveTwice = true
			return true
		}
	}
	return true // Unvisited cave
}

func travelComplex(g Graph, start cave, behind route, visitedAnySmallCaveTwice bool, routeCount *int) {
	path := behind | route(start)

	if start == g.getStringAsBoolPointer("end") {
		*routeCount++
		return
	}

	neighbours := g.getNeighbours(start)

	for _, n := range neighbours {
		localDidVisitTwice := visitedAnySmallCaveTwice
		if canComplexTravel(n, path, g.getStringAsBoolPointer("start"), &localDidVisitTwice) {
			travelComplex(g, n, path, localDidVisitTwice, routeCount)
		}
	}

}

func parseGraph(lines []string) graphdata {
	g := newGraph()

	for _, line := range lines {
		a, b := file.SplitOnce(line, "-")
		g.addEdge(g.getStringAsBoolPointer(a), g.getStringAsBoolPointer(b))
	}

	return g
}

func countRoutes(g Graph) int {
	start := g.getStringAsBoolPointer("start")
	count := 0
	travel(g, start, route(0), &count)
	return count
}

func countComplexRoutes(g Graph) int {
	start := g.getStringAsBoolPointer("start")
	count := 0
	travelComplex(g, start, route(0), false, &count)
	return count
}

func Solve() (int, int, error) {
	lines := file.ReadFile("./day12/input.txt")
	g := parseGraph(lines)

	log.Println("cave count", len(g.idMap))

	count := countRoutes(&g)
	countComplex := countComplexRoutes(&g)

	log.Printf("Simples routes: %v Complex routes: %v\n", count, countComplex)

	return count, countComplex, nil
}
