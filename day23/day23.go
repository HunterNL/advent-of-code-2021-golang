package day23

import (
	"aoc2021/intmath"
	"fmt"
	"math"
	"os"
	"strings"
)

type path []position
type pathSpec struct{ from, to position }

type vec2 struct {
	x, y int
}

var positionMap = [23]vec2{
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

func findVectorIndex(v vec2) int {
	if v.y == 0 {
		if v.x == 2 || v.x == 4 || v.x == 6 || v.x == 8 {
			return -1 // No positions needed for places we can't stand
		}
	}

	for key, value := range positionMap {
		if v.x == value.x && v.y == value.y {
			return key
		}
	}

	panic("Did not find vector")
}

type amphipod byte

type position int

func manhatten(v1, v2 vec2) int {
	return intmath.Distance(v1.x, v2.x) + intmath.Distance(v1.y, v2.y)
}

func isHome(pod amphipod, position position) bool {
	return desiredBoard[position] == pod
}

func isInRoom(p position) bool {
	return p > 6
}
func isInHallway(p position) bool {
	return p < 7
}

var pathMemo = make(map[int]path)
var pathCostMemo = make(map[int]int)

// var roomA1 = {2, 1}
// var roomA2 = {2, 2}
// var roomB1 = {4, 1}
// var roomB2 = {4, 2}
// var roomC1 = {6, 1}
// var roomC2 = {6, 2}
// var roomD1 = {8, 1}
// var roomD2 = {8, 2}
var HallwaySize = 7
var RoomSize = 4

//
var roomA = position(HallwaySize + RoomSize*0)
var roomB = position(HallwaySize + RoomSize*1)
var roomC = position(HallwaySize + RoomSize*2)
var roomD = position(HallwaySize + RoomSize*3)

// var hallwayPositions = []position{
// 	{0, 0},
// 	{1, 0},
// 	// {2, 0},
// 	{3, 0},
// 	// {4, 0},
// 	{5, 0},
// 	// {6, 0},
// 	{7, 0},
// 	// {8, 0},
// 	{9, 0},
// 	{10, 0},
// }

var desiredBoard = board{
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

// func (b board) isBlocker(pos position) bool {
// 	pod := b[pos]
// 	roomPositions := targetRooms[pod]
// 	return isEqual(roomPositions[0], pos) && b[roomPositions[1]] != pod
// }

func roomIsEmpty(b board, startIndex int) bool {
	for i := startIndex; i < startIndex+RoomSize; i++ {
		if b[i] != amphipod(0) {
			return false
		}
	}

	return true
}

// func (b board) canLeaveRoom(pos position) bool {

// 	roomPositions := targetRooms[pod]
// 	return isEqual(roomPositions[0], pos) && b[roomPositions[1]] != pod

// }

func (b board) shouldLeaveHome(pod amphipod, pos position) bool {
	start, end := podHomePositions(pod)
	for i := start + 1; i <= end; i++ {
		if b[i] != b[pos] && b[i] != amphipod(0) {
			return true // If any amphipods behind the given are different, allow a move out
		}
	}

	return false
}

type board [23]amphipod

func (b board) availableDestinations(currentPosition position) []position {
	pod := b[currentPosition]
	out := make([]position, 0, 16)

	if isHome(pod, currentPosition) && !b.shouldLeaveHome(pod, currentPosition) {
		return out
	}

	if isInRoom(currentPosition) {
		for targetPosition := position(6); targetPosition >= 0; targetPosition-- {
			if b.canPathTo(currentPosition, targetPosition) {
				out = append(out, targetPosition)
			}
		}
		return out
	}

	if isInHallway(currentPosition) && canEnterHome(b, pod) {
		roomPos := findHomePosition(b, pod)
		if b.canPathTo(currentPosition, roomPos) {
			out = append(out, roomPos)
		}
		return out
	}

	return []position{}

	// panic("Unknown state")
}

// func isInTargetRoom(pod amphipod, currentPosition position) bool {
// 	targetPositions := targetRooms[pod]
// 	return currentPosition == targetPositions[0] || currentPosition == targetPositions[1]
// }

func roomIsEnpty(b board, startPos position) {

}

func findHomePosition(b board, pod amphipod) position {
	start, end := podHomePositions(pod)
	for position := end; position >= start; position-- {
		if !b.isOccupied(position) {
			return position
		}
	}

	panic("Unkown state, room full")
}

func canEnterHome(b board, pod amphipod) bool {
	start, end := podHomePositions(pod)

	for i := end; i >= start; i-- {
		if b[i] != pod && b[i] != amphipod(0) {
			return false // Cannot enter if there's other kinds of amphipods present
		}
	}

	return true

	// panic("Unknown state")
}

func podHomePositions(pod amphipod) (position, position) {
	return targetRooms[pod], targetRooms[pod] + 3
}

func roomTargetForPod(b board, pod amphipod) position {
	top, bottom := podHomePositions(pod)

	if b.isOccupied(bottom) {
		return top
	} else {
		return bottom
	}
}

func (b board) canPathTo(from, to position) bool {
	if b.isOccupied(to) {
		return false
	}

	path := createPath(from, to)
	for _, pos := range path {
		if b.isOccupied(pos) {
			return false
		}
	}

	return true
}

func createPath(from, to position) []position {
	if from == to {
		panic("Positions are equal")
	}

	memoPath, memoFound := pathMemo[int(from*100+to)]

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
			posIndex := findVectorIndex(current)
			if posIndex > -1 {
				out = append(out, position(findVectorIndex(current)))
			}
		}
	}
	var addYPath = func() {
		for current.y != target.y {
			current.y += yDir
			posIndex := findVectorIndex(current)
			if posIndex > -1 {
				out = append(out, position(findVectorIndex(current)))
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

	return out
}

func (b board) isOccupied(pos position) bool {
	return b[pos] != amphipod(0)
}

type game struct {
	state board
	score int
}

func (b board) isWon() bool {
	return b == desiredBoard
}

var costMap = map[amphipod]int{
	amphipod('A'): 1,
	amphipod('B'): 10,
	amphipod('C'): 100,
	amphipod('D'): 1000,
}

func cost(from, to position, pod amphipod) int {
	cost := costMap[pod]
	memoScore, memoFound := pathCostMemo[int(from*100+to)]

	if memoFound {
		return memoScore * cost
	}

	distance := manhatten(positionMap[from], positionMap[to])

	pathCostMemo[int(from*100+to)] = distance

	return distance * cost
}

func (g game) applyMove(from, to position) game {
	pod := g.state[from]
	g.state[from] = amphipod(0)
	g.state[to] = pod
	g.score += cost(from, to, pod)

	return g
}

// func (g game) applyMove(from, to position) game {
// 	out := game{state: g.state, score: g.score}
// 	pod := out.state[from]
// 	out.state[from] = amphipod(0)
// 	out.state[to] = pod
// 	out.score += cost(from, to, pod)

// 	return out
// }

var printDebug bool = false

func (g game) playMoves(lowScore *int) {

	if g.state.isWon() {
		if g.score < *lowScore {
			*lowScore = g.score
			fmt.Println(g.score)
		}
		// *finishedGames = append(*finishedGames, g)
		return
	}

	if g.score > *lowScore {
		return // Optimization, we're never going to beat the highscore, abort
	}

	// fmt.Print("----------------------------------------------\n")
	// fmt.Print("Current state:\n")
	// fmt.Print(g.state)

	// Start at the end, play around with the heavy hitting amphipods first
	for pos := position(22); pos >= 0; pos-- {
		if !g.state.isOccupied(pos) {
			continue
		}

		movements := g.state.availableDestinations(pos)

		if len(movements) == 0 {
			continue
		}

		// fmt.Printf("Position %v (%v) has %v movements available\n", pos, positionMap[pos], len(movements))
		// fmt.Printf("Current state:\n%v\n", g.state)

		movements = g.state.availableDestinations(pos)

		for _, to := range movements {

			// fmt.Printf("From %v to %v (%v to %v)\n", pos, to, positionMap[pos], positionMap[to])
			newGame := g.applyMove(pos, to)

			// fmt.Print(newGame.state)

			newGame.playMoves(lowScore)
		}
	}
}

func (b board) String() string {
	lines := make([][]byte, 5)

	for i := 0; i < 5; i++ {
		lines[i] = []byte("...........")
	}

	for position, pod := range b {
		pos := positionMap[position]
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

func parseBoard(s string) board {
	justChars := strings.ReplaceAll(strings.Trim(s, " \n"), "#", "")

	lines := strings.Split(justChars, "\n")
	topRoomContent := strings.TrimSpace(lines[2])
	bottomRoomContent := strings.TrimSpace(lines[3])

	return board{
		7 + 0:  amphipod(topRoomContent[0]),
		7 + 4:  amphipod(topRoomContent[1]),
		7 + 8:  amphipod(topRoomContent[2]),
		7 + 12: amphipod(topRoomContent[3]),

		7 + 1:  amphipod('D'),
		7 + 5:  amphipod('C'),
		7 + 9:  amphipod('B'),
		7 + 13: amphipod('A'),

		7 + 2:  amphipod('D'),
		7 + 6:  amphipod('B'),
		7 + 10: amphipod('A'),
		7 + 14: amphipod('C'),

		7 + 3:  amphipod(bottomRoomContent[0]),
		7 + 7:  amphipod(bottomRoomContent[1]),
		7 + 11: amphipod(bottomRoomContent[2]),
		7 + 15: amphipod(bottomRoomContent[3]),
	}
}

func lowestScore(games []game) int {
	out := games[0].score
	for i := 1; i < len(games); i++ {
		if games[i].score < out {
			out = games[i].score
		}
	}
	return out
}

func Solve() (int, int) {
	bytes, err := os.ReadFile("./day23/input.txt")

	if err != nil {
		panic(err)
	}

	board := parseBoard(string(bytes))

	g := game{state: board}

	// allGames := []game{}

	fmt.Printf("Starting board:\n\n%v\n\n", g.state)

	var score int = math.MaxInt

	g.playMoves(&score)

	// score := lowestScore(allGames)

	fmt.Printf("Lowest score:%v\n", score)

	return -1, score
}
