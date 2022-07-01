package day12

// type cave bool

type Graph interface {
	addEdge(a cave, b cave)
	getNeighbours(a cave) []cave
	getStringAsBoolPointer(string) cave
}

type graphdata struct {
	vertices   []cave
	edges      [][2]cave
	lookup     map[cave][]cave
	pointermap map[string]cave
}

func newGraph() graphdata {
	return graphdata{
		pointermap: make(map[string]cave),
		vertices:   make([]cave, 0),
		edges:      make([][2]cave, 0),
		lookup:     map[cave][]cave{},
	}
}

func (g *graphdata) getStringAsBoolPointer(str string) cave {
	ptr, found := g.pointermap[str]
	if found {
		return ptr
	}

	boolean := isUpper(str)

	g.pointermap[str] = &boolean

	return &boolean
}

func (g *graphdata) addEdge(a cave, b cave) {
	g.edges = append(g.edges, [2]cave{a, b})

	g.lookup[a] = append(g.lookup[a], b)
	g.lookup[b] = append(g.lookup[b], a)
}

func (g *graphdata) getNeighbours(a cave) []cave {
	return g.lookup[a]
}
