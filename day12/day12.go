package day12

import (
	"aoc2021/file"
	"log"
	"strings"
)

// Pointers are used to uniquely identify a cave, it's value used for cave size
// "Cave" pointers are never used to change the pointed value
type cave *bool
type route []cave

func isUpper(str string) bool {
	return strings.ToUpper(str) == str
}

func visited(needle cave, haystack route) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func travel(g Graph, start cave, behind route) []route {
	path := append(behind, start)

	if start == g.getStringAsBoolPointer("end") {
		return []route{path}
	}

	routes := []route{}

	neighbours := g.getNeighbours(start)

	for _, n := range neighbours {
		if *n {
			routes = append(routes, travel(g, n, path)...)
		} else {
			if !visited(n, path) {
				routes = append(routes, travel(g, n, path)...)
			}
		}

	}

	return routes
}

// Returns if we travel to `to`
// Note that unlike "cave" bool pointers, this plain pointer does update `visitedAnySmallCaveTwice`
func canComplexTravel(to cave, behind route, startPtr cave, visitedAnySmallCaveTwice *bool) bool {
	// Probit reentering starting cave
	if to == startPtr {
		return false
	}

	if *to {
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

func travelComplex(g Graph, start cave, behind route, visitedAnySmallCaveTwice bool) []route {
	path := append(behind, start)

	if start == g.getStringAsBoolPointer("end") {
		return []route{path}
	}

	routes := []route{}

	neighbours := g.getNeighbours(start)

	for _, n := range neighbours {
		localDidVisitTwice := visitedAnySmallCaveTwice
		if canComplexTravel(n, path, g.getStringAsBoolPointer("start"), &localDidVisitTwice) {
			routes = append(routes, travelComplex(g, n, path, localDidVisitTwice)...)
		}
	}

	return routes
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
	routes := travel(g, start, []cave{})
	return len(routes)
}

func countComplexRoutes(g Graph) int {
	start := g.getStringAsBoolPointer("start")
	routes := travelComplex(g, start, []cave{}, false)
	return len(routes)
}

func Solve() (int, int) {
	lines := file.ReadFile("./day12/input.txt")
	g := parseGraph(lines)

	count := countRoutes(&g)
	countComplex := countComplexRoutes(&g)

	log.Printf("Simples routes: %v Complex routes: %v\n", count, countComplex)

	return count, countComplex
}
