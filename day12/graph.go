package day12

type cave bool

type Graph interface {
	addEdge(a *bool, b *bool)
	getNeighbours(a *bool) []*bool
	getStringAsBoolPointer(string) *bool
}

type graphdata struct {
	vertices   []*bool
	edges      [][2]*bool
	lookup     map[*bool][]*bool
	pointermap map[string]*bool
}

func newGraph() graphdata {
	return graphdata{
		pointermap: make(map[string]*bool),
		vertices:   make([]*bool, 0),
		edges:      make([][2]*bool, 0),
		lookup:     map[*bool][]*bool{},
	}
}

func (g *graphdata) getStringAsBoolPointer(str string) *bool {
	ptr, found := g.pointermap[str]
	if found {
		return ptr
	}

	boolean := isUpper(str)

	g.pointermap[str] = &boolean

	return &boolean
}

func (g *graphdata) addEdge(a *bool, b *bool) {
	g.edges = append(g.edges, [2]*bool{a, b})

	if _, found := g.lookup[a]; !found {
		g.lookup[a] = make([]*bool, 0)
	}

	if _, found := g.lookup[b]; !found {
		g.lookup[b] = make([]*bool, 0)
	}

	g.lookup[a] = append(g.lookup[a], b)
	g.lookup[b] = append(g.lookup[b], a)
}

func (g *graphdata) getNeighbours(a *bool) []*bool {
	return g.lookup[a]
}
