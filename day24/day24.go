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

type sectionConfig struct {
	zDivStep int
	xAdd     int
	yAdd     int
}

type zState struct {
	state        int
	forwardInput int
}

// Replicates what the ALU is doing, any step-to-step changes in the original input are input here area sectionConfig
func sectionStep(a *alu, s sectionConfig, input int) {
	x := 0
	y := 0
	z := a.state
	w := input

	x = (z % 26) + s.xAdd // [0-25]+xAdd
	z = z / s.zDivStep    // z | z*26+[1-25]
	y = 25
	if x == w { //   [1-9]
		x = 0
		y = 0
	} else {
		x = 1
	}

	//x : 1|0
	y++            //1|26
	z = z * y      //z|z/26|z*26+[1-25]
	y = w + s.yAdd //[1-9]+yAdd
	y = y * x      //[1-9]+yAdd|0
	z = z + y      //((z*(1|26))/(1|26))-(0|[1-9]+yAdd)									//z|(z-[1-9]-yAdd)/26|(z-[1-9]-yAdd)*26+[1-25]				//z|z+[1-9]+yAdd

	a.state = z
}

func findValidStartZStates(step sectionConfig, finalZState int, upperZLimit int, s *stats) []zState {
	zStates := make(map[int]bool)

	// zStates[finalZState] = true

	for input := 1; input <= 9; input++ {
		for i := 1; i <= step.zDivStep; i++ {
			z := finalZState
			y := input + step.yAdd
			// zets := [2]int{z, z - y, (z - y) / 26}

			zPure := z            // x == w
			zMult := (z - y) / 26 // x != w

			if (zPure*step.zDivStep+i)%26+step.xAdd == input {
				zStates[zPure*step.zDivStep+i] = true
			}

			if (zMult*step.zDivStep+i)%26+step.xAdd != input {
				zStates[zMult*step.zDivStep+i] = true
			}

		}
	}

	a := alu{state: 0}
	out := []zState{}
	for k := range zStates {

		if k < 0 {
			continue
		}
		if k > upperZLimit {
			continue
		}
		for i := 1; i <= 9; i++ {
			s.found++
			a.state = k
			a.executeStep(step, i)
			if a.state == finalZState {
				out = append(out, zState{state: k, forwardInput: i})
			} else {
				s.discarded++
			}
		}
	}

	// log.Printf("Output state %6v has %6v possible input states: %v\n", finalZState, len(out), out)

	return out
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
const backTraceCount = 7
const highestBacktraceSection = numSize - backTraceCount
const leftSize = numSize - backTraceCount
const rightSize = backTraceCount

// const

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

type stats struct {
	found     int
	discarded int
}

// Works backwards from the final step (where Z has to be 0) and bruteforces zValues in the previous step to see what zValues leads to succesfull zValues in the next step
func findValidZStates(steps []sectionConfig) zStateLookup {
	acceptedZExitState := make(zStateLookup)
	for i := 0; i < 13; i++ {
		acceptedZExitState[i] = map[int]int{}
	}

	s := stats{}

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

		// log.Printf("\n[zDiv:%3v xAdd:%3v yAdd:%3v]\n", step.zDivStep, step.xAdd, step.yAdd)

		for k := range validExitStates {
			for _, zState := range findValidStartZStates(step, k, zCap[sectionId], &s) {
				currentState := acceptedZExitState[sectionId-1][zState.state]
				newState := intPush(currentState, zState.forwardInput)
				acceptedZExitState[sectionId-1][zState.state] = newState
			}
		}

		// log.Printf("Found %v valid exit states for section %vn", len(acceptedZExitState[sectionId-1]), sectionId-1)
		// log.Printf("Found a total of %10.v valid zStates (+inputs) before discarding %10v (%2.2f%%)\n", s.found, s.discarded, (float32(s.discarded)/float32(s.found))*100.0)

	}

	log.Print(len(acceptedZExitState[highestBacktraceSection-1]))

	return acceptedZExitState
}

func bruteForce(steps []sectionConfig, validzSates zStateLookup, advanceFunc func([]int, int) bool, sliceSelector func(s []int) int, startDigit int) int {

	leftALU := alu{state: 0}
	leftNum := createIntSlice(startDigit, leftSize)
	rightSteps := steps[leftSize:]
	leftSteps := steps[:leftSize]

	for {

		// if leftNum[7] == 9 && leftNum[6] == 9 && leftNum[5] == 9 && leftNum[4] == 9 && leftNum[3] == 9 && leftNum[2] == 9 {
		// 	log.Printf("Current leftNum: %v\n", leftNum)
		// }

		leftALU.reset()
		leftALU.executeSteps(leftSteps, leftNum)

		input, isValidZ := validzSates[highestBacktraceSection-1][leftALU.state]

		if isValidZ {
			log.Println("Found valid starting digits!", leftNum, leftALU.state)
			for i := 0; i < rightSize; i++ {
				if intHasMoreNumbers(input) {
					input = sliceSelector(intToSlice(input))
				}

				leftNum = append(leftNum, input)
				leftALU.executeStep(rightSteps[i], input)
				input = validzSates[highestBacktraceSection+i][leftALU.state]
			}
			if leftALU.state != 0 {
				log.Fatal("expected ALU to reach 0")
			}
			return digitSliceToInt(leftNum)
			// rightNum := createIntSlice(startDigit, backTraceCount)

			// for {
			// 	rightALU := leftALU //Reset to a1 state every attempt
			// 	var sectionId int
			// 	for sectionId = leftSize; sectionId < len(steps); sectionId++ {

			// 		rightALU.executeStep(steps[sectionId], rightNum[sectionId-leftSize])

			// 		if sectionId == 13 && rightALU.state == 0 {
			// 			log.Printf("Found result: %v %v", leftNum, rightNum)
			// 			return digitSliceToInt(append(leftNum[:], rightNum[:]...))
			// 		}

			// 	}
			// 	if advanceFunc(rightNum[:], rightSize-1) {
			// 		break
			// 	}
			// }
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

	// Calculate z targets, valid for both attempts
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
