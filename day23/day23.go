package day23

import (
	"aoc2021/intmath"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type path []position

type vec2 struct {
	x, y int
}

const BOARD_SIZE = 23

type largeBoard [BOARD_SIZE]amphipod

func (b *largeBoard) amphipodAt(i position) amphipod {
	return b[i]
}

type gameSet []game
type gameMap map[largeBoard]int

// Len is the number of elements in the collection.
func (gs gameSet) Len() int {
	return len(gs)
}

func (gs gameSet) Less(i int, j int) bool {
	return gs[i].fScore < gs[j].fScore
}

// Swap swaps the elements with indexes i and j.
func (gs gameSet) Swap(i int, j int) {
	gs[i], gs[j] = gs[j], gs[i]
	gs[i].heapIndex = i
	gs[j].heapIndex = j
}

func (gs *gameSet) Push(x any) {
	n := len(*gs)
	item := x.(game)
	item.heapIndex = n
	*gs = append(*gs, item)
}

func (gs *gameSet) Pop() any {
	old := *gs
	n := len(old)
	item := old[n-1]
	// old[n-1] = nil      // avoid memory leak
	item.heapIndex = -1 // for safety
	*gs = old[0 : n-1]
	return item
}

// type smallBoard [15]amphipod

type amphipod byte
type position int
type pathMemo map[int]path

var positionMapLarge = [BOARD_SIZE]vec2{
	0:  {0, 0},
	1:  {1, 0},
	2:  {3, 0},
	3:  {5, 0},
	4:  {7, 0},
	5:  {9, 0},
	6:  {10, 0},
	7:  {2, 1}, //R1
	8:  {2, 2},
	9:  {2, 3},
	10: {2, 4},
	11: {4, 1}, //R2
	12: {4, 2},
	13: {4, 3},
	14: {4, 4},
	15: {6, 1}, //R3
	16: {6, 2},
	17: {6, 3},
	18: {6, 4},
	19: {8, 1}, //R4
	20: {8, 2},
	21: {8, 3},
	22: {8, 4},
}

func findVectorIndex(v vec2, positionMap *[BOARD_SIZE]vec2) int {
	if v.y == 0 {
		if v.x == 2 || v.x == 4 || v.x == 6 || v.x == 8 {
			return -1 // No positions needed for places we can't stand
		}
	}

	for key, value := range *positionMap {
		if v.x == value.x && v.y == value.y {
			return key
		}
	}

	panic("Did not find vector")
}

func manhatten(v1, v2 vec2) int {
	return intmath.Distance(v1.x, v2.x) + intmath.Distance(v1.y, v2.y)
}

func isHome(pod amphipod, position position) bool {
	return desiredLargeBoard[position] == pod
}

func isInRoom(p position) bool {
	return p > 6
}
func isInHallway(p position) bool {
	return p < 7
}

const HallwaySize = 7
const RoomSize = 4

var roomA = position(HallwaySize + RoomSize*0)
var roomB = position(HallwaySize + RoomSize*1)
var roomC = position(HallwaySize + RoomSize*2)
var roomD = position(HallwaySize + RoomSize*3)

var desiredLargeBoard = [BOARD_SIZE]amphipod{
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod('A'),
	amphipod('A'),
	amphipod('A'),
	amphipod('A'),
	amphipod('B'),
	amphipod('B'),
	amphipod('B'),
	amphipod('B'),
	amphipod('C'),
	amphipod('C'),
	amphipod('C'),
	amphipod('C'),
	amphipod('D'),
	amphipod('D'),
	amphipod('D'),
	amphipod('D'),
}

var desiredSmallBoard = [BOARD_SIZE]amphipod{
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod(0),
	amphipod('A'),
	amphipod('A'),
	amphipod(0),
	amphipod(0),
	amphipod('B'),
	amphipod('B'),
	amphipod(0),
	amphipod(0),
	amphipod('C'),
	amphipod('C'),
	amphipod(0),
	amphipod(0),
	amphipod('D'),
	amphipod('D'),
	amphipod(0),
	amphipod(0),
}

var targetRooms = map[amphipod]position{
	'A': roomA,
	'B': roomB,
	'C': roomC,
	'D': roomD,
}

func sign(a int) int {
	if a > 0 {
		return 1
	} else {
		return -1
	}
}

func shouldLeaveHome(b *largeBoard, pos position, config *gameCache) bool {
	pod := b.amphipodAt(pos)
	start, end := podHomePositions(pod, config)
	for i := start + 1; i <= end; i++ {
		if b.amphipodAt(i) != pod && isOccupied(b, i) {
			return true // If any amphipods behind the given are different, allow a move out
		}
	}

	return false
}

func availableDestinations(b *largeBoard, currentPosition position, cache *gameCache) []position {
	pod := b.amphipodAt(currentPosition)
	out := make([]position, 0, 16)

	if isHome(pod, currentPosition) && !shouldLeaveHome(b, currentPosition, cache) {
		return out
	}

	if isInRoom(currentPosition) {
		for targetPosition := position(6); targetPosition >= 0; targetPosition-- {
			if canPathTo(b, currentPosition, targetPosition, cache) {
				out = append(out, targetPosition)
			}
		}
		return out
	}

	if isInHallway(currentPosition) && canEnterHome(b, pod, cache) {
		roomPos := findHomePosition(b, pod, cache)
		if canPathTo(b, currentPosition, roomPos, cache) {
			out = append(out, roomPos)
		}
		return out
	}

	return []position{}
}

func findHomePosition(b *largeBoard, pod amphipod, config *gameCache) position {
	start, end := podHomePositions(pod, config)
	for position := end; position >= start; position-- {
		if !isOccupied(b, position) {
			return position
		}
	}

	panic("Unkown state, room full")
}

func canEnterHome(b *largeBoard, pod amphipod, config *gameCache) bool {
	start, end := podHomePositions(pod, config)

	for i := end; i >= start; i-- {
		if b[i] != pod && b[i] != amphipod(0) {
			return false // Cannot enter if there's other kinds of amphipods present
		}
	}
	return true
}

func podHomePositions(pod amphipod, config *gameCache) (position, position) {
	return targetRooms[pod], targetRooms[pod] + position(config.roomSize) - 1
}

func canPathTo(b *largeBoard, from, to position, cache *gameCache) bool {
	if isOccupied(b, to) {
		return false
	}

	path := createPath(from, to, cache)
	for _, pos := range path {
		if isOccupied(b, pos) {
			return false
		}
	}

	return true
}

func createPath(from, to position, cache *gameCache) []position {
	if from == to {
		panic("Positions are equal")
	}

	memoPath, memoFound := cache.pathMemo[int(from*100+to)]
	positionMap := cache.positionMap

	if memoFound {
		return memoPath
	}

	current := cache.positionMap[from]
	target := cache.positionMap[to]

	out := make([]position, 0, 10)
	xDir := sign(target.x - current.x)
	yDir := sign(target.y - current.y)

	var addXPath = func() {
		for current.x != target.x {
			current.x += xDir
			posIndex := findVectorIndex(current, &positionMap)
			if posIndex > -1 {
				out = append(out, position(findVectorIndex(current, &positionMap)))
			}
		}
	}
	var addYPath = func() {
		for current.y != target.y {
			current.y += yDir
			posIndex := findVectorIndex(current, &positionMap)
			if posIndex > -1 {
				out = append(out, position(findVectorIndex(current, &positionMap)))
			}
		}
	}

	if isInRoom(from) {
		addYPath()
		addXPath()
	} else {
		addXPath()
		addYPath()
	}

	cache.pathMemo[int(from*100+to)] = out

	return out
}

func isOccupied(b *largeBoard, pos position) bool {
	return b[pos] != amphipod(0)
}

// Frequently copied game data
type game struct {
	state     largeBoard
	score     int
	fScore    int
	heapIndex int
}

// Data commonly referenced and altered but not copied around
type gameCache struct {
	pathMemo     pathMemo
	distanceMemo map[int]int
	positionMap  [BOARD_SIZE]vec2
	roomMap      map[amphipod]position
	roomSize     int
	winState     largeBoard
}

var costMap = map[amphipod]int{
	amphipod('A'): 1,
	amphipod('B'): 10,
	amphipod('C'): 100,
	amphipod('D'): 1000,
}

func distance(cache *gameCache, from, to position) int {
	return manhatten(cache.positionMap[from], cache.positionMap[to])
}

func (g game) applyMove(from, to position, cache *gameCache) game {
	g.state[to] = g.state[from]
	g.state[from] = amphipod(0)

	g.score += distance(cache, from, to) * costMap[g.state[to]]

	g.fScore = g.score + hueristic(g, cache)

	return g
}

func isWon(board largeBoard, cache *gameCache) bool {
	return board == cache.winState
}

func render(b largeBoard, pm [BOARD_SIZE]vec2) string {
	lines := make([][]byte, 5)

	for i := 0; i < 5; i++ {
		lines[i] = []byte("###########")
	}

	for pos, pod := range b {
		pos := pm[position(pos)]
		lines[pos.y][pos.x] = byte(pod)
	}

	sb := strings.Builder{}

	for _, v := range lines {
		sb.Write(v)
		sb.WriteRune('\n')
	}

	sb.WriteRune('\n')

	return sb.String()
}

type parsemode int

const (
	PARSE_SMALL parsemode = iota
	PARSE_LARGE
)

func parseBoard(s string, mode parsemode) largeBoard {

	justChars := strings.ReplaceAll(strings.Trim(s, " \n"), "#", "")
	rowOffset := 1
	if mode == PARSE_LARGE {
		rowOffset = 3
	}

	lines := strings.Split(justChars, "\n")
	topRoomContent := strings.TrimSpace(lines[2])
	bottomRoomContent := strings.TrimSpace(lines[3])

	board := largeBoard{}
	for i := 0; i < 4; i++ {
		board[HallwaySize+i*4] = amphipod(topRoomContent[i])
	}

	for i := 0; i < 4; i++ {
		board[HallwaySize+rowOffset+i*4] = amphipod(bottomRoomContent[i])
	}

	if mode == PARSE_LARGE {
		day2Content := map[int]amphipod{
			7 + 1:  amphipod('D'),
			7 + 5:  amphipod('C'),
			7 + 9:  amphipod('B'),
			7 + 13: amphipod('A'),

			7 + 2:  amphipod('D'),
			7 + 6:  amphipod('B'),
			7 + 10: amphipod('A'),
			7 + 14: amphipod('C'),
		}

		for i, pod := range day2Content {
			board[i] = pod
		}

	}

	return board
}

func hueristic(g game, config *gameCache) int {
	out := 0
	for i, pod := range g.state {
		pos := position(i)
		if pod == amphipod(0) {
			continue
		}

		if isHome(pod, pos) {
			if shouldLeaveHome(&g.state, pos, config) { // Blocking something, gotta move
				homeEntry, _ := podHomePositions(pod, config)
				out += (int(pos-homeEntry) + 2) * costMap[pod] // 2 is the minimal distance required to leave a room
			}
			continue
		}
		// Pod is in hallway or wrong room
		if isInHallway(pos) {
			target, _ := podHomePositions(pod, config)

			out += (distance(config, pos, target) * costMap[pod])
			continue
		}

		//Pod is in wrong room
		out += minDistanceHome(pos, pod, config) * costMap[pod]
	}

	return out
}

func minDistanceHome(pos position, pod amphipod, config *gameCache) int {
	homeStart, _ := podHomePositions(pod, config)
	from := config.positionMap[pos]
	to := config.positionMap[homeStart]

	return from.y + to.y + (to.x - from.x)
}

type move struct {
	from position
	to   position
}

func findQuickestMoves(startPosition game, config *gameCache) int {
	openSet := gameSet{}
	seenStates := gameMap{}

	heap.Push(&openSet, startPosition)

	for {
		lowestItem := heap.Pop(&openSet)
		lowestGame := lowestItem.(game)

		if isWon(lowestGame.state, config) {
			return lowestGame.score
		}

		moves := []move{}

		for i, pod := range lowestGame.state {
			if pod == amphipod(0) {
				continue
			}
			from := position(i)
			for _, to := range availableDestinations(&lowestGame.state, from, config) {
				moves = append(moves, move{from: from, to: to})
			}
		}

		for _, move := range moves {
			newGame := lowestGame.applyMove(move.from, move.to, config)

			if score, found := seenStates[newGame.state]; found {
				if newGame.score < score {
					seenStates[newGame.state] = score
				} else {
					// State already exists, with a better score, don't bother
					continue
				}
			} else {
				seenStates[newGame.state] = newGame.score
			}

			heap.Push(&openSet, newGame)

		}
	}
}

func Solve() (int, int) {
	bytes, err := os.ReadFile("./day23/input.txt")
	if err != nil {
		panic(err)
	}

	configp1 := gameCache{
		pathMemo:     map[int]path{},
		distanceMemo: map[int]int{},
		positionMap:  positionMapLarge,
		winState:     desiredSmallBoard,
		roomMap:      targetRooms,
		roomSize:     2,
	}

	// p1 := solveP1()
	boardP1 := parseBoard(string(bytes), PARSE_SMALL)
	p1Score := findQuickestMoves(game{state: boardP1}, &configp1)

	board := parseBoard(string(bytes), PARSE_LARGE)

	g := game{state: board}

	config := gameCache{
		pathMemo:     map[int]path{},
		distanceMemo: map[int]int{},
		positionMap:  positionMapLarge,
		winState:     desiredLargeBoard,
		roomMap:      targetRooms,
		roomSize:     4,
	}

	score := findQuickestMoves(g, &config)

	// allGames := []game{}

	fmt.Printf("Starting board:\n\n%v\n\n", g.state)

	// var score int = math.MaxInt

	// g.playMoves(&score, &cache, 0)

	// score := lowestScore(allGames)

	fmt.Printf("Lowest score:%v\n", score)

	// fmt.Println(p1, score)

	return p1Score, score
}
