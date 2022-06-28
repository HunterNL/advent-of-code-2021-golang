package day12

import (
	"aoc2021/file"
	"log"
	"strings"
)

type route = []*bool

func isUpper(str string) bool {
	return strings.ToUpper(str) == str
}

func visited(needle *bool, haystack route) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func travel(g Graph, start *bool, behind route) []route {
	path := append(behind, start)

	if start == g.getStringAsBoolPointer("end") {
		return [][]*bool{path}
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

func visitCount(cave *bool, behind route) int {
	sum := 0
	for _, v := range behind {
		if v == cave {
			sum++
		}
	}

	return sum
}

func canComplexTravel(to *bool, behind route, startPtr *bool) bool {
	if to == startPtr {
		return false
	}

	if *to {
		// log.Printf("Route %v can travel to %v\n", behind, to)
		return true
	} else {
		// log.Printf("Visited %v before: %v Visited any small cave twice: %v\n", to, visited(to, behind), visitedAnySmallCaveTwice(behind))
		return !visited(to, behind) || !visitedAnySmallCaveTwice(behind)
	}

}

// #PERFORMANCE Turn route into a struct and set value once
func visitedAnySmallCaveTwice(route []*bool) bool {
	seen := make(map[*bool]bool)

	for _, v := range route {
		if *v {
			continue // Ignore large caves
		}
		if seen[v] {
			return true // We've seen this cave small before
		}
		seen[v] = true // Note that we've seen this small cave

	}

	return false
}

func travelComplex(g Graph, start *bool, behind route) []route {
	path := append(behind, start)

	if start == g.getStringAsBoolPointer("end") {
		return [][]*bool{path}
	}

	routes := []route{}

	neighbours := g.getNeighbours(start)

	for _, n := range neighbours {
		if canComplexTravel(n, path, g.getStringAsBoolPointer("start")) {
			routes = append(routes, travelComplex(g, n, path)...)
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
	routes := travel(g, start, []*bool{})
	return len(routes)
}

func countComplexRoutes(g Graph) int {
	start := g.getStringAsBoolPointer("start")
	routes := travelComplex(g, start, []*bool{})
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
