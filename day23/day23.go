package day23

import (
	"aoc2021/intmath"
	"container/heap"
	"os"
	"strings"
)

const BOARD_SIZE = 23

// Array holding all amphipod positions
type board [BOARD_SIZE]amphipod
type amphipod byte
type position int

func findQuickestMoves(startPosition board, roomSize int, winState board) int {

	openSet := gameSet{}
	seenStates := make(map[board]int, 0)

	config := gameConfig{
		winState: winState,
		roomSize: roomSize,
	}

	heap.Push(&openSet, game{state: startPosition})

	for {
		lowestItem := heap.Pop(&openSet)
		lowestGame := lowestItem.(game)

		if isWon(lowestGame.state, &config) {
			return lowestGame.score
		}

		moves := []move{}

		for i, pod := range lowestGame.state {
			if pod == amphipod(0) {
				continue
			}
			from := position(i)
			for _, to := range availableDestinations(&lowestGame.state, from, roomSize) {
				moves = append(moves, move{from: from, to: to})
			}
		}

		for _, move := range moves {
			newGame := lowestGame.applyMove(move.from, move.to, roomSize)

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

func availableDestinations(b *board, currentPosition position, roomSize int) []position {
	pod := b[currentPosition]
	out := make([]position, 0, 16)

	if isHome(pod, currentPosition) && !shouldLeaveHome(b, currentPosition, roomSize) {
		return out
	}

	if isInRoom(currentPosition) {
		for targetPosition := position(6); targetPosition >= 0; targetPosition-- {
			if canPathTo(b, currentPosition, targetPosition) {
				out = append(out, targetPosition)
			}
		}
		return out
	}

	if isInHallway(currentPosition) && canEnterHome(b, pod, roomSize) {
		roomPos := findHomePosition(b, pod, roomSize)
		if canPathTo(b, currentPosition, roomPos) {
			out = append(out, roomPos)
		}
		return out
	}

	return []position{}
}

func (g game) applyMove(from, to position, roomSize int) game {
	g.state[to] = g.state[from]
	g.state[from] = amphipod(0)

	g.score += distance(from, to) * costMap[g.state[to]]

	g.fScore = g.score + hueristic(g, roomSize)

	return g
}

func hueristic(g game, roomSize int) int {
	out := 0
	for i, pod := range g.state {
		pos := position(i)
		if pod == amphipod(0) {
			continue
		}

		if isHome(pod, pos) {
			if shouldLeaveHome(&g.state, pos, roomSize) { // Blocking something, gotta move
				homeEntry, _ := podHomePositions(pod, roomSize)
				out += (int(pos-homeEntry) + 2) * costMap[pod] // 2 is the minimal distance required to leave a room
			}
			continue
		}
		// Pod is in hallway or wrong room
		if isInHallway(pos) {
			target, _ := podHomePositions(pod, roomSize)

			out += (distance(pos, target) * costMap[pod])
			continue
		}

		//Pod is in wrong room
		out += minDistanceHome(pos, pod, roomSize) * costMap[pod]
	}

	return out
}

// Start movement logic

// An amphipod should leave home if it's blocking a different type amphipod from leaving
func shouldLeaveHome(b *board, pos position, roomSize int) bool {
	pod := b[pos]
	start, end := podHomePositions(pod, roomSize)
	for i := start + 1; i <= end; i++ {
		if b[i] != pod && isOccupied(b, i) {
			return true // If any amphipods behind the given are different, allow a move out
		}
	}

	return false
}

func findHomePosition(b *board, pod amphipod, roomSize int) position {
	start, end := podHomePositions(pod, roomSize)
	for position := end; position >= start; position-- {
		if !isOccupied(b, position) {
			return position
		}
	}

	panic("Unkown state, room full")
}

// Amphipod can only enter if no other amphipod types are present
func canEnterHome(b *board, pod amphipod, roomSize int) bool {
	start, end := podHomePositions(pod, roomSize)

	for i := end; i >= start; i-- {
		if b[i] != pod && b[i] != amphipod(0) {
			return false // Cannot enter if there's other kinds of amphipods present
		}
	}
	return true
}

// Returns the position of the room
// this also marks the biggest difference between small (part1) and large (part2) board
// the remaining room slots still exist in the array but aren't considered
func podHomePositions(pod amphipod, roomSize int) (position, position) {
	return targetRooms[pod], targetRooms[pod] + position(roomSize) - 1
}

func canPathTo(b *board, from, to position) bool {
	if isOccupied(b, to) {
		return false
	}

	path := createPath(from, to)
	for _, pos := range path {
		if isOccupied(b, pos) {
			return false
		}
	}

	return true
}

//End movement logic

// Board with own score, expected score and heap position
type game struct {
	state     board
	score     int
	fScore    int
	heapIndex int
}

var costMap = map[amphipod]int{
	amphipod('A'): 1,
	amphipod('B'): 10,
	amphipod('C'): 100,
	amphipod('D'): 1000,
}

// Global memoization of paths between two points, hashed as x*100+y
var pathMemo = map[int]path{}

type vec2 struct {
	x, y int
}

// We can reuse this for day 1, roomsize is given and taken into account
var positionMap = [BOARD_SIZE]vec2{
	0:  {0, 0},
	1:  {1, 0},
	2:  {3, 0},
	3:  {5, 0},
	4:  {7, 0},
	5:  {9, 0},
	6:  {10, 0},
	7:  {2, 1}, //Room 1
	8:  {2, 2},
	9:  {2, 3},
	10: {2, 4},
	11: {4, 1}, //Room 2
	12: {4, 2},
	13: {4, 3},
	14: {4, 4},
	15: {6, 1}, //Room 3
	16: {6, 2},
	17: {6, 3},
	18: {6, 4},
	19: {8, 1}, //Room 4
	20: {8, 2},
	21: {8, 3},
	22: {8, 4},
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

var desiredLargeBoard = board{
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

type path []position

func createPath(from, to position) []position {
	if from == to {
		panic("Positions are equal")
	}

	memoPath, memoFound := pathMemo[int(from*100+to)]
	positionMap := positionMap

	if memoFound {
		return memoPath
	}

	current := positionMap[from]
	target := positionMap[to]

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

	pathMemo[int(from*100+to)] = out

	return out
}

// For mapping a vector back to a board index
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

func isOccupied(b *board, pos position) bool {
	return b[pos] != amphipod(0)
}

func distance(from, to position) int {
	return manhatten(positionMap[from], positionMap[to])
}

func isWon(b board, cache *gameConfig) bool {
	return b == cache.winState
}

func render(b board, pm [BOARD_SIZE]vec2) string {
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

func parseBoard(s string, mode parsemode) board {

	justChars := strings.ReplaceAll(strings.Trim(s, " \n"), "#", "")
	rowOffset := 1
	if mode == PARSE_LARGE {
		rowOffset = 3
	}

	lines := strings.Split(justChars, "\n")
	topRoomContent := strings.TrimSpace(lines[2])
	bottomRoomContent := strings.TrimSpace(lines[3])

	out := board{}
	for i := 0; i < 4; i++ {
		out[HallwaySize+i*4] = amphipod(topRoomContent[i])
	}

	for i := 0; i < 4; i++ {
		out[HallwaySize+rowOffset+i*4] = amphipod(bottomRoomContent[i])
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
			out[i] = pod
		}

	}

	return out
}

func minDistanceHome(pos position, pod amphipod, roomSize int) int {
	homeStart, _ := podHomePositions(pod, roomSize)
	from := positionMap[pos]
	to := positionMap[homeStart]

	return from.y + to.y + (to.x - from.x)
}

type move struct {
	from position
	to   position
}

type gameConfig struct {
	roomSize int
	winState board
}

func Solve() (int, int, error) {
	bytes, err := os.ReadFile("./day23/input.txt")
	if err != nil {
		panic(err)
	}

	boardP1 := parseBoard(string(bytes), PARSE_SMALL)
	boardP2 := parseBoard(string(bytes), PARSE_LARGE)

	scoreP1 := findQuickestMoves(boardP1, 2, desiredSmallBoard)
	scoreP2 := findQuickestMoves(boardP2, 4, desiredLargeBoard)

	return scoreP1, scoreP2, nil
}

type gameSet []game

// Implementation of heap.Interface
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
