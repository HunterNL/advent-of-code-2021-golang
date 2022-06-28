package day15

import (
	"aoc2021/file"
	"aoc2021/grid"
	"container/heap"
	"log"
	"strconv"
	"strings"
)

type node struct {
	g         int
	h         int
	index     int
	parent    int
	heapIndex int
}

type nodeheap []node

// Len is the number of elements in the collection.
func (nh *nodeheap) Len() int {
	return len(*nh) // TODO: Implement
}

func (nh nodeheap) Less(i int, j int) bool {
	return nh[i].g+nh[i].h < nh[j].g+nh[j].h
}

// Swap swaps the elements with indexes i and j.
func (nh nodeheap) Swap(i int, j int) {
	nh[i], nh[j] = nh[j], nh[i]
	nh[i].heapIndex = i
	nh[j].heapIndex = j
}

func (nh *nodeheap) Push(x interface{}) {
	size := len(*nh)
	n := x.(node)
	n.heapIndex = size
	*nh = append(*nh, n)
}

func (nh *nodeheap) Pop() interface{} {
	old := *nh
	n := len(old)
	item := old[n-1]
	// old[n-1] = nil  // avoid memory leak
	item.heapIndex = -1 // for safety
	*nh = old[0 : n-1]
	return item
}

func findPath(g *[]uint8, rowSize, rowCount, start, end int) []int {
	// openSet := make(map[int]node, len(*g))
	openSet := make(nodeheap, 0, len(*g))
	closedSet := make(map[int]node, len(*g))

	heap.Init(&openSet)

	targetX, targetY := grid.ToXY(end, rowSize)
	currentX, currentY := grid.ToXY(start, rowSize)

	current := node{
		index:  start,
		g:      0,
		h:      grid.Manhatten(currentX, currentY, targetX, targetY),
		parent: -1,
	}
	// currentBase := 0

	// openSet[current.index] = current

	heap.Push(&openSet, current)
	done := false

	for !done {
		// lowestScore := math.MaxInt
		// var lowestIndex int
		// for i, n := range openSet {
		// 	f := n.g + n.h
		// 	if f < lowestScore {
		// 		lowestScore = f
		// 		current = n
		// 		lowestIndex = i
		// 	}
		// }
		p := heap.Pop(&openSet)
		current = p.(node)

		// log.Printf("Removed %v from open set\n", current.index)
		// delete(openSet, lowestIndex)
		closedSet[current.index] = current

		neighbours := grid.ValidNeighbours(current.index, rowSize, rowCount)
		for _, i := range neighbours {
			_, found := closedSet[i]
			if found {
				continue
			}
			x, y := grid.ToXY(i, rowSize)
			var cost int = int((*g)[i]) * 1000
			var dist = grid.Manhatten(x, y, targetX, targetY)

			if i == end {
				done = true
				closedSet[i] = node{g: current.g + cost, h: dist, index: i, parent: current.index}
				break
			}

			heapIndex := -1
			for _, n := range openSet {
				if n.index == i {
					heapIndex = n.heapIndex
					break
				}
			}

			if heapIndex > -1 {
				// log.Printf("Fixing node %v\n", openSet[heapIndex].index)
				if openSet[heapIndex].g > current.g+cost {
					n2 := openSet[i]
					n2.g = current.g + cost
					n2.parent = current.index
					openSet[i] = n2
				}
				heap.Fix(&openSet, heapIndex)
			} else {
				// log.Printf("Adding node %v\n", i)
				heap.Push(&openSet, node{g: current.g + cost, h: dist, index: i, parent: current.index})
			}

			// log.Printf("Distance to target from %v %v: %v\n", x, y, dist)

		}

		// log.Printf("Open set size: %v closed set size: %v\n", len(openSet), len(closedSet))
	}

	// Convert closed set to ordered list
	out := make([]int, 0, len(closedSet))

	for i := end; i != start; {
		n := closedSet[i]
		out = append(out, n.index)
		i = n.parent
	}

	// log.Printf("Path: %v\n", closedSet)
	// log.Printf("Out: %v\n", out)

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return out
}

func parseGrid(lines []string) ([]uint8, int, int) {
	rowSize := len(lines[0])
	rowCount := len(lines)

	grid := make([]uint8, rowCount*rowSize)

	str := strings.Join(lines, "")
	str = strings.ReplaceAll(str, "\n", "")

	for i, r := range str {
		n, err := strconv.ParseInt(string(r), 10, 8)
		if err != nil {
			panic(err)
		}

		grid[i] = uint8(n)
	}

	return grid, rowSize, rowCount
}

func part1(file []string) int {
	g, rowSize, rowCount := parseGrid(file)

	risk := 0
	path := findPath(&g, rowSize, rowCount, 0, rowSize*rowCount-1)

	for _, i := range path {
		x, y := grid.ToXY(i, rowSize)
		log.Printf("X: %v, Y:%v\n", x, y)
		risk += int(g[i])
	}

	return risk
}

func Solve() (int, int) {
	file := file.ReadFile("./day15/input.txt")
	g, rowSize, rowCount := parseGrid(file)
	g2 := repeatRight(&g, rowSize)
	g3 := repeatDown(&g2)

	risk := 0
	path := findPath(&g3, rowSize*5, rowCount*5, 0, 500*500-1)

	for _, i := range path {
		x, y := grid.ToXY(i, rowSize)
		log.Printf("X: %v, Y:%v\n", x, y)
		risk += int(g3[i])
	}

	log.Printf("Total risk: %v\n", risk)

	return part1(file), risk
}

func inc(i, a uint8) uint8 {
	return (i+a-1)%9 + 1
}

func repeatRight(old *[]uint8, rowSize int) []uint8 {
	reps := 5
	grid := make([]uint8, len(*old)*reps)

	for i, v := range *old {
		row := i / rowSize
		grid[reps*row*rowSize+(i%rowSize)+rowSize*0] = inc(v, 0) // orginal
		grid[reps*row*rowSize+(i%rowSize)+rowSize*1] = inc(v, 1) // orginal
		grid[reps*row*rowSize+(i%rowSize)+rowSize*2] = inc(v, 2) // orginal
		grid[reps*row*rowSize+(i%rowSize)+rowSize*3] = inc(v, 3) // orginal
		grid[reps*row*rowSize+(i%rowSize)+rowSize*4] = inc(v, 4) // orginal
	}

	return grid
}
func repeatDown(old *[]uint8) []uint8 {
	size := len(*old)
	reps := 5
	grid := make([]uint8, size*reps)

	for i, v := range *old {
		grid[i+size*0] = inc(v, 0) // orginal
		grid[i+size*1] = inc(v, 1) // orginal
		grid[i+size*2] = inc(v, 2) // orginal
		grid[i+size*3] = inc(v, 3) // orginal
		grid[i+size*4] = inc(v, 4) // orginal
	}

	return grid
}

func remove(s []node, i int) []node {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
