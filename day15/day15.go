package day15

import (
	"aoc2021/file"
	"aoc2021/grid"
	"aoc2021/pool"
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
	x         int
	y         int
}

type vec2 struct {
	x int
	y int
}

type nodeheap []node

// Len is the number of elements in the collection.
func (nh nodeheap) Len() int {
	return len(nh) // TODO: Implement
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

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func validNeighbours(x, y, rowSize, rowCount int) [4]vec2 {
	neighbours := [4]vec2{}

	top := y > 0
	bottom := y < rowCount-1

	left := x > 0
	right := x < rowSize-1

	if top {
		neighbours[0] = vec2{x, y - 1}
	}

	if left {
		neighbours[1] = vec2{x - 1, y}
	}

	if right {
		neighbours[2] = vec2{x + 1, y}
	}

	if bottom {
		neighbours[3] = vec2{x, y + 1}
	}

	return neighbours
}

func findPath(g []uint8, rowSize, rowCount, start, end int) []int {
	// openSet := make(map[int]node, len(*g))
	openSet := make(nodeheap, 0, len(g))
	closedSet := make(map[int]node, len(g))
	nodePool := pool.MakePool[node](rowSize * rowCount)

	heap.Init(&openSet)

	targetX, targetY := grid.ToXY(end, rowSize)
	currentX, currentY := grid.ToXY(start, rowSize)

	current := *nodePool.Pop()

	current.index = start
	current.g = 0
	current.h = grid.Manhatten(currentX, currentY, targetX, targetY)
	current.parent = -1
	current.x = 0
	current.y = 0
	// currentBase := 0

	// openSet[current.index] = current

	heap.Push(&openSet, current)
	done := false

	for !done {
		//Grab a new node
		pop := heap.Pop(&openSet)
		current = pop.(node)
		closedSet[current.index] = current

		x := current.x
		y := current.y

		// Get all possible paths forward
		neighbours := validNeighbours(x, y, rowSize, rowCount)
		for _, i := range neighbours {
			neighbourGridIndex := i.x + i.y*rowSize
			neighbourX := i.x
			neighbourY := i.y

			var cost int = int(g[neighbourGridIndex]) * 1000

			// If it's in the closed set, update it's costs and continue
			closedNode, found := closedSet[neighbourGridIndex]
			if found {
				if closedNode.g > current.g+cost {
					closedNode.g = current.g + cost
					closedNode.parent = current.index
				}
				continue
			}

			var dist = grid.Manhatten(neighbourX, neighbourY, targetX, targetY)

			// If we've reached the end: wrap up
			if neighbourGridIndex == end {
				done = true
				n := *nodePool.Pop()
				n.g = current.g + cost
				n.h = dist
				n.index = neighbourGridIndex
				n.parent = current.index
				closedSet[neighbourGridIndex] = n
				break
			}

			// Why the heck is this here?
			// heapIndex := -1
			// for _, n := range openSet {
			// 	if n.index == neighbourGridIndex {
			// 		heapIndex = n.heapIndex
			// 		break
			// 	}
			// }

			// Insert new path into the open set
			n := *nodePool.Pop()
			n.g = current.g + cost
			n.h = dist
			n.x = neighbourX
			n.y = neighbourY
			n.index = neighbourGridIndex
			n.parent = current.index
			closedSet[neighbourGridIndex] = n
			heap.Push(&openSet, n)

		}
	}

	// Convert closed set to ordered list
	out := make([]int, 0, len(closedSet))

	for i := end; i != start; {
		n := closedSet[i]
		out = append(out, n.index)
		i = n.parent
	}

	// for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
	// 	out[i], out[j] = out[j], out[i]
	// }

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
	path := findPath(g, rowSize, rowCount, 0, rowSize*rowCount-1)

	for _, i := range path {
		x, y := grid.ToXY(i, rowSize)
		log.Printf("X: %v, Y:%v\n", x, y)
		risk += int(g[i])
	}

	return risk
}

func Solve() (int, int, error) {
	file := file.ReadFile("./day15/input.txt")
	g, rowSize, rowCount := parseGrid(file)
	g2 := repeatRight(&g, rowSize)
	g3 := repeatDown(&g2)

	risk := 0
	path := findPath(g3, rowSize*5, rowCount*5, 0, 500*500-1)

	for _, i := range path {
		x, y := grid.ToXY(i, rowSize)
		log.Printf("X: %v, Y:%v\n", x, y)
		risk += int(g3[i])
	}

	log.Printf("Total risk: %v\n", risk)

	return part1(file), risk, nil
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
