package day23

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func isHomeP1(pod amphipod, vec2 vec2) bool {
	return desiredp1Board[vec2] == pod
}

func isInRoomP1(p vec2) bool {
	return p.y > 0
}
func isInHallwayP1(p vec2) bool {
	return p.y == 0
}

func isEqual(p1, p2 vec2) bool {
	return p1.x == p2.x && p1.y == p2.y
}

var roomA1 = vec2{2, 1}
var roomA2 = vec2{2, 2}
var roomB1 = vec2{4, 1}
var roomB2 = vec2{4, 2}
var roomC1 = vec2{6, 1}
var roomC2 = vec2{6, 2}
var roomD1 = vec2{8, 1}
var roomD2 = vec2{8, 2}

var hallwayPositions = []vec2{
	{0, 0},
	{1, 0},
	// {2, 0},
	{3, 0},
	// {4, 0},
	{5, 0},
	// {6, 0},
	{7, 0},
	// {8, 0},
	{9, 0},
	{10, 0},
}

var desiredp1Board = p1Board{
	roomA1: amphipod('A'),
	roomA2: amphipod('A'),
	roomB1: amphipod('B'),
	roomB2: amphipod('B'),
	roomC1: amphipod('C'),
	roomC2: amphipod('C'),
	roomD1: amphipod('D'),
	roomD2: amphipod('D'),
}

var targetRoomsP1 = map[amphipod][2]vec2{
	'A': {roomA1, roomA2},
	'B': {roomB1, roomB2},
	'C': {roomC1, roomC2},
	'D': {roomD1, roomD2},
}

// func sign(a int) int {
// 	if a > 0 {
// 		return 1
// 	} else {
// 		return -1
// 	}
// }

func (b p1Board) isBlocker(pos vec2) bool {
	pod := b[pos]
	roomPositions := targetRoomsP1[pod]
	return isEqual(roomPositions[0], pos) && b[roomPositions[1]] != pod

}

type p1Board map[vec2]amphipod

func (b p1Board) availableDestinations(pod amphipod, currentPosition vec2) []vec2 {
	out := make([]vec2, 0, 10)

	if isHomeP1(pod, currentPosition) && !b.isBlocker(currentPosition) {
		return out
	}

	if isInRoomP1(currentPosition) {
		for _, targetPosition := range hallwayPositions {
			if b.canPathTo(currentPosition, targetPosition) {
				out = append(out, targetPosition)
			}
		}
		return out
	}

	if isInHallwayP1(currentPosition) && canEnterHomeP1(b, pod) {
		roomPos := roomTargetForPodP1(b, pod)
		if b.canPathTo(currentPosition, roomPos) {
			out = append(out, roomPos)
		}
		return out
	}

	return []vec2{}

	// panic("Unknown state")
}

// func isInTargetRoom(pod amphipod, currentPosition vec2) bool {
// 	targetPositions := targetRoomsP1[pod]
// 	return currentPosition == targetPositions[0] || currentPosition == targetPositions[1]
// }

func canEnterHomeP1(b p1Board, pod amphipod) bool {
	top, bottom := podHomePositionsP1(pod)
	if !b.isOccupied(top) && !b.isOccupied(bottom) {
		return true
	}

	if b.isOccupied(top) {
		return false
	}

	if b.isOccupied(bottom) && b[bottom] == pod {
		return true
	}

	return false

	// panic("Unknown state")
}

func podHomePositionsP1(pod amphipod) (vec2, vec2) {
	return targetRoomsP1[pod][0], targetRoomsP1[pod][1]
}

func roomTargetForPodP1(b p1Board, pod amphipod) vec2 {
	top, bottom := podHomePositionsP1(pod)

	if b.isOccupied(bottom) {
		return top
	} else {
		return bottom
	}
}

func (b p1Board) canPathTo(from, to vec2) bool {
	if b.isOccupied(to) {
		return false
	}

	path := createPathP1(from, to)
	for _, pos := range path {
		if b.isOccupied(pos) {
			return false
		}
	}

	return true
}

func createPathP1(from, to vec2) []vec2 {
	if isEqual(from, to) {
		panic("Positions are equal")
	}

	current := from

	out := make([]vec2, 0, 10)
	xDir := sign(to.x - from.x)
	yDir := sign(to.y - from.y)

	var addXPath = func() {
		for current.x != to.x {
			current.x += xDir
			out = append(out, current)
		}
	}
	var addYPath = func() {
		for current.y != to.y {
			current.y += yDir
			out = append(out, current)
		}
	}

	if isInRoomP1(from) {
		addYPath()
		addXPath()
	} else {
		addXPath()
		addYPath()
	}

	return out
}

func (b p1Board) isOccupied(pos vec2) bool {
	_, found := b[pos]
	return found
}

type gamep1 struct {
	state p1Board
	score int
}

func (b p1Board) isWon() bool {
	for pos, pod := range desiredp1Board {
		if b[pos] != pod {
			return false
		}
	}

	return true
}

func costP1(from, to vec2, pod amphipod) int {
	distance := manhatten(from, to)
	return map[amphipod]int{
		amphipod('A'): 1,
		amphipod('B'): 10,
		amphipod('C'): 100,
		amphipod('D'): 1000,
	}[pod] * distance
}

func (g gamep1) applyMoveP1(from, to vec2) gamep1 {
	pod := g.state[from]
	cost := costP1(from, to, pod)
	out := gamep1{score: g.score + cost, state: p1Board{}}

	for k, v := range g.state {
		out.state[k] = v
	}

	delete(out.state, from)

	out.state[to] = pod

	return out
}

func (g gamep1) playMovesP1(lowScore *int) {
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

	for pos, pod := range g.state {
		if !isHomeP1(pod, pos) || g.state.isBlocker(pos) {
			movements := g.state.availableDestinations(pod, pos)
			for _, to := range movements {
				newGame := g.applyMoveP1(pos, to)

				if printDebug {
					fmt.Print(newGame.state)
				}

				newGame.playMovesP1(lowScore)
			}
		}
	}
}

func (b p1Board) String() string {
	lines := make([][]byte, 3)

	for i := 0; i < 3; i++ {
		lines[i] = []byte("............")
	}

	for vec2, pod := range b {
		lines[vec2.y][vec2.x] = byte(pod)
	}

	sb := strings.Builder{}

	for _, v := range lines {
		sb.Write(v)
		sb.WriteRune('\n')
	}

	sb.WriteRune('\n')

	return sb.String()
}

func parsep1Board(s string) p1Board {
	parseLine := func(lineString string, y int) p1Board {
		justChars := strings.ReplaceAll(strings.Trim(lineString, " "), "#", "")
		return p1Board{
			vec2{2, y}: amphipod(justChars[0]),
			vec2{4, y}: amphipod(justChars[1]),
			vec2{6, y}: amphipod(justChars[2]),
			vec2{8, y}: amphipod(justChars[3]),
		}
	}

	lines := strings.Split(s, "\n")
	topRoomContent := parseLine(lines[2], 1)
	bottomRoomContent := parseLine(lines[3], 2)

	for k, v := range bottomRoomContent {
		topRoomContent[k] = v
	}

	return topRoomContent
}

func lowestScorep1(games []gamep1) int {
	out := games[0].score
	for i := 1; i < len(games); i++ {
		if games[i].score < out {
			out = games[i].score
		}
	}
	return out
}

func solveP1() int {
	bytes, err := os.ReadFile("./day23/input.txt")

	if err != nil {
		panic(err)
	}

	p1Board := parsep1Board(string(bytes))

	g := gamep1{state: p1Board}

	// allGames := []game{}

	fmt.Printf("Starting p1Board:\n\n%v\n\n", g.state)

	var score int = math.MaxInt

	g.playMovesP1(&score)

	// score = lowestScorep1(allGames)

	fmt.Printf("Lowest score:%v\n", score)

	return score
}
