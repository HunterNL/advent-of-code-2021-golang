package day12

// type cave bool

const LARGE_CAVE_BIT cave = 1 << 31

type Graph interface {
	addEdge(a cave, b cave)
	getNeighbours(a cave) []cave
	getStringAsBoolPointer(string) cave
}

type graphdata struct {
	edges  [][2]cave
	lookup map[cave][]cave
	idMap  map[string]cave
}

func newGraph() graphdata {
	return graphdata{
		idMap:  make(map[string]cave),
		edges:  make([][2]cave, 0),
		lookup: map[cave][]cave{},
	}
}

func (g *graphdata) getStringAsBoolPointer(str string) cave {
	ptr, found := g.idMap[str]
	if found {
		return ptr
	}

	var new_cave cave = 1 << len(g.idMap)

	isLarge := isUpper(str)

	if isLarge {
		new_cave = new_cave | LARGE_CAVE_BIT
	}

	g.idMap[str] = new_cave

	return new_cave
}

func (g *graphdata) addEdge(a cave, b cave) {
	g.edges = append(g.edges, [2]cave{a, b})

	g.lookup[a] = append(g.lookup[a], b)
	g.lookup[b] = append(g.lookup[b], a)
}

func (g *graphdata) getNeighbours(a cave) []cave {
	return g.lookup[a]
}
