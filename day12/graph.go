package day12

type Graph interface {
	addEdge(a string, b string)
	getNeighbours(a string) []string
}

type graphdata struct {
	vertices []string
	edges    [][2]string
	lookup   map[string][]string
}

func newGraph() graphdata {
	return graphdata{
		vertices: make([]string, 0),
		edges:    make([][2]string, 0),
		lookup:   map[string][]string{},
	}
}

func (g *graphdata) addEdge(a string, b string) {
	g.edges = append(g.edges, [2]string{a, b})

	if _, found := g.lookup[a]; !found {
		g.lookup[a] = make([]string, 0)
	}

	if _, found := g.lookup[b]; !found {
		g.lookup[b] = make([]string, 0)
	}

	g.lookup[a] = append(g.lookup[a], b)
	g.lookup[b] = append(g.lookup[b], a)
}

func (g *graphdata) getNeighbours(a string) []string {
	return g.lookup[a]
}
