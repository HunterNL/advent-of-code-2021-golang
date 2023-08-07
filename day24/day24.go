package day24

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type alu struct {
	state int
}

func (a *alu) executeSteps(steps []sectionConfig, inputs []int) {
	for i, section := range steps {
		sectionStep(a, section, inputs[i])
	}
}
func (a *alu) executeStep(step sectionConfig, in int) {
	sectionStep(a, step, in)
}

func programs(file []byte) (_, _, steps []sectionConfig, err error) {
	strSections := strings.Split(strings.TrimSpace(string(file)), "inp w\n")[1:]
	steps = []sectionConfig{}

	for _, section := range strSections {
		lines := strings.Split(strings.TrimSpace(section), "\n")
		lines = append([]string{"inp w"}, lines...)
		steps = append(steps, parseSectionConfig(lines))
	}

	return nil, nil, steps, nil
}

func (a *alu) reset() {
	a.state = 0
}

func sliceInt(in int) []int {
	out := []int{}
	for _, n := range fmt.Sprint(in) {
		digit, err := strconv.Atoi(string(n))
		if err != nil {
			panic(err)
		}
		out = append(out, digit)
	}

	return out
}

func createIntSlice(content, size int) []int {
	slice := make([]int, size)

	for i := range slice {
		slice[i] = content
	}

	return slice
}

func parseSectionConfig(lines []string) sectionConfig {
	xAddLine := lines[5]
	xAddStr := strings.Split(xAddLine, " ")[2]
	xAdd, err := strconv.Atoi(xAddStr)
	if err != nil {
		panic(err)
	}

	yAddLine := lines[15]
	yAddStr := strings.Split(yAddLine, " ")[2]
	yAdd, err := strconv.Atoi(yAddStr)
	if err != nil {
		panic(err)
	}

	zDivLine := lines[4]
	zDivStr := strings.Split(zDivLine, " ")[2]
	zDiv, err := strconv.Atoi(zDivStr)
	if err != nil {
		panic(err)
	}

	return sectionConfig{
		xAdd:     xAdd,
		yAdd:     yAdd,
		zDivStep: zDiv,
	}

}

type zState struct {
	state        int
	forwardInput int
}

type sectionConfig struct {
	zDivStep int
	xAdd     int
	yAdd     int
}

// Replicates what the ALU is doing, any step-to-step changes in the original input are input here area sectionConfig
func sectionStep(a *alu, s sectionConfig, input int) {
	x := 0
	y := 0
	z := a.state
	w := input

	x = (z % 26) + s.xAdd // [0-25]+xAdd

	if x == w { //   [1-9]
		x = 0
		y = 0
	} else {
		x = 1
		y = 25
	}

	//x : 1|0
	y++                //1|26
	z = z / s.zDivStep // z | z*26+[1-25]
	z = z * y          //z|z/26|z*26+[1-25]

	y = w + s.yAdd //[1-9]+yAdd
	y = y * x      //[1-9]+yAdd|0
	z = z + y      //((z*(1|26))/(1|26))-(0|[1-9]+yAdd)									//z|(z-[1-9]-yAdd)/26|(z-[1-9]-yAdd)*26+[1-25]				//z|z+[1-9]+yAdd

	a.state = z
}

func reverseStep(step sectionConfig, finalZ int) []zState {
	out := []zState{}

	// Tackling the x == w branch first
	// Thus x = 0, y = 1
	zTrue := finalZ * step.zDivStep
	for div := 0; div < step.zDivStep; div++ {
		zTrueDiv := zTrue + div

		zTrueMod := (zTrueDiv % 26) + step.xAdd
		if zTrueMod >= 1 && zTrueMod <= 9 {
			out = append(out, zState{state: zTrueDiv, forwardInput: zTrueMod})
		}
	}

	// And now the x != w branch
	// Here x = 1 and y=25
	for input := 1; input <= 9; input++ {
		y := input + step.yAdd
		zFalse := finalZ - y

		// Must cleanly divide by 26
		if zFalse%26 != 0 {
			continue
		}

		zFalse = zFalse / 26

		zFalse = zFalse * step.zDivStep

		for div := 0; div < step.zDivStep; div++ {
			zFalseDiv := zFalse + div //  eg 6/4 and 7/4 would result in the same, gotta attempt multiple origin values

			zFalseDiv = zFalseDiv * step.zDivStep

			zFalseMod := (zFalseDiv % 26) + step.xAdd
			if zFalseMod != input {
				out = append(out, zState{state: zFalseDiv, forwardInput: input})
			}
		}

	}

	return out
}

func findValidStartZStates(step sectionConfig, finalZState int, upperZLimit int) []zState {
	zStates := []zState{}

	inputZs := reverseStep(step, finalZState)
	for _, in := range inputZs {
		if in.state > upperZLimit {
			continue
		}

		zStates = append(zStates, in)

	}

	return zStates
}

func decrementNum(nums []int, startIndex int) bool {
	incIndex := startIndex
	for incIndex >= 0 {
		nums[incIndex]--
		if nums[incIndex] == 0 {
			nums[incIndex] = 9
			incIndex--
			continue
		}
		return false
	}
	return true
}
func incrementNum(nums []int, startIndex int) bool {
	incIndex := startIndex
	for incIndex >= 0 {
		nums[incIndex]++
		if nums[incIndex] == 10 {
			nums[incIndex] = 1
			incIndex--
			continue
		}
		return false
	}
	return true
}

const numSize = 14
const backTraceCount = 8
const bruteForceCount = numSize - backTraceCount

const leftSize = bruteForceCount
const rightSize = backTraceCount

// Utilities for using an int as a quick array of uint4's
func intPush(base, append int) int {
	base = base << 4
	base = base | append
	return base
}

func intToSlice(n int) []int {
	out := []int{}
	for n > 0 {
		out = append(out, n&15)
		n = n >> 4
	}
	return out
}

func intHasMoreNumbers(n int) bool {
	return n > 15
}

type zStateLookup map[int]map[int]int

func highestPossibleZStateAfterStep(initialZState int, step sectionConfig) int {
	zState := initialZState / step.zDivStep
	zState = zState * 26
	return zState + 9 + step.yAdd
}

// Works backwards from the final step (where Z has to be 0) and bruteforces zValues in the previous step to see what zValues leads to succesfull zValues in the next step
func findValidZStates(steps []sectionConfig) zStateLookup {
	highestBacktraceSection := leftSize
	acceptedZExitState := make(zStateLookup)
	for i := 0; i < 13; i++ {
		acceptedZExitState[i] = map[int]int{}
	}

	zCap := [14]int{}
	lastCap := 0
	for i := 0; i < len(zCap); i++ {
		zCap[i] = highestPossibleZStateAfterStep(lastCap, steps[i])
		lastCap = zCap[i]
	}

	// The state we're looking for, 0 after step 13
	acceptedZExitState[13] = map[int]int{0: 0}

	for sectionId := 13; sectionId >= highestBacktraceSection; sectionId-- {
		step := steps[sectionId]
		validExitStates := acceptedZExitState[sectionId]

		for k := range validExitStates {
			for _, zState := range findValidStartZStates(step, k, zCap[sectionId]) {
				currentState := acceptedZExitState[sectionId-1][zState.state]
				newState := intPush(currentState, zState.forwardInput)
				acceptedZExitState[sectionId-1][zState.state] = newState
			}
		}
	}

	return acceptedZExitState
}

func bruteForce(steps []sectionConfig, validzSates zStateLookup, advanceFunc func([]int, int) bool, sliceSelector func(s []int) int, startDigit int) int {

	leftALU := alu{state: 0}
	leftNum := createIntSlice(startDigit, leftSize)
	rightSteps := steps[leftSize:]
	leftSteps := steps[:leftSize]

	for {
		leftALU.reset()
		leftALU.executeSteps(leftSteps, leftNum)

		input, isValidZ := validzSates[leftSize-1][leftALU.state]

		if isValidZ {
			log.Println("Found valid starting digits!", leftNum, leftALU.state)
			for i := 0; i < rightSize; i++ {
				if intHasMoreNumbers(input) {
					input = sliceSelector(intToSlice(input))
				}

				leftNum = append(leftNum, input)
				leftALU.executeStep(rightSteps[i], input)
				input = validzSates[leftSize+i][leftALU.state]
			}
			if leftALU.state != 0 {
				log.Fatal("expected ALU to reach 0")
			}
			return digitSliceToInt(leftNum)
		}

		if advanceFunc(leftNum[:], leftSize-1) {
			break
		}

	}

	panic("No solution found")
}

func digitSliceToInt(digits []int) int {
	out := 0
	for i := len(digits) - 1; i >= 0; i-- {
		out = out + digits[i]*int(math.Pow(10, float64(len(digits)-i-1)))
	}
	return out
}

func Solve() (int, int, error) {
	log.SetFlags(0)
	sectionFile, err := os.ReadFile("./day24/input.txt")
	if err != nil {
		return -1, -1, err
	}
	_, _, steps, _ := programs(sectionFile)

	zStates := findValidZStates(steps)

	upper := bruteForce(steps, zStates, decrementNum, sliceMax, 9)
	lower := bruteForce(steps, zStates, incrementNum, sliceMin, 1)

	return upper, lower, nil
}

func sliceMin(s []int) int {
	min := math.MaxInt
	for _, i := range s {
		if i < min {
			min = i
		}
	}
	return min
}

func sliceMax(s []int) int {
	max := math.MinInt
	for _, i := range s {
		if i > max {
			max = i
		}
	}
	return max
}
