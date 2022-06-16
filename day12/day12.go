package day12

import (
	"aoc2021/file"
	"fmt"
	"strings"
)

type route = []string

func isUpper(str string) bool {
	return strings.ToUpper(str) == str
}

func visited(needle string, haystack route) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func travel(g Graph, start string, behind route) []route {
	path := append(behind, start)

	if start == "end" {
		return [][]string{path}
	}

	routes := []route{}

	neighbours := g.getNeighbours(start)

	for _, n := range neighbours {
		if isUpper(n) {
			routes = append(routes, travel(g, n, path)...)
		} else {
			if !visited(n, path) {
				routes = append(routes, travel(g, n, path)...)
			}
		}

	}

	return routes
}

func visitCount(cave string, behind route) int {
	sum := 0
	for _, v := range behind {
		if v == cave {
			sum++
		}
	}

	return sum
}

func canComplexTravel(to string, behind route) bool {
	if to == "start" {
		return false
	}

	if isUpper(to) {
		// fmt.Printf("Route %v can travel to %v\n", behind, to)
		return true
	} else {
		// fmt.Printf("Visited %v before: %v Visited any small cave twice: %v\n", to, visited(to, behind), visitedAnySmallCaveTwice(behind))
		return !visited(to, behind) || !visitedAnySmallCaveTwice(behind)
	}

}

// #PERFORMANCE Turn route into a struct and set value once
func visitedAnySmallCaveTwice(route []string) bool {
	seen := make(map[string]bool)

	for _, v := range route {
		if isUpper(v) {
			continue
		}
		if seen[v] {
			return true
		}
		seen[v] = true

	}

	return false
}

func travelComplex(g Graph, start string, behind route) []route {
	path := append(behind, start)

	if start == "end" {
		return [][]string{path}
	}

	routes := []route{}

	neighbours := g.getNeighbours(start)

	for _, n := range neighbours {
		if canComplexTravel(n, path) {
			routes = append(routes, travelComplex(g, n, path)...)
		}
	}

	return routes
}

func parseGraph(lines []string) graphdata {
	g := newGraph()

	for _, line := range lines {
		a, b := file.SplitOnce(line, "-")
		g.addEdge(a, b)
	}

	return g
}

func countRoutes(g Graph) int {
	routes := travel(g, "start", []string{})
	return len(routes)
}

func countComplexRoutes(g Graph) int {
	routes := travelComplex(g, "start", []string{})
	return len(routes)
}

func Solve() (int, int) {
	lines := file.ReadFile("./day12/input.txt")
	g := parseGraph(lines)

	count := countRoutes(&g)
	countComplex := countComplexRoutes(&g)

	fmt.Printf("Simples routes: %v Complex routes: %v\n", count, countComplex)

	return count, countComplex
}
