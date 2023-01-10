package day24

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type alu int

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
	*a = 0
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

// Replicates what the ALU is doing, any step-to-step changes in the original input are input here area sectionConfig
func sectionStep(a *alu, s sectionConfig, input int) {
	x := 0
	y := 0
	z := int(*a)
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

	*a = alu(z)
}

func findValidStartZStates(step sectionConfig, finalZState int, upperZLimit int) []int {
	zStates := make(map[int]bool)

	zStates[finalZState] = true

	for input := 1; input <= 9; input++ {
		for i := 1; i <= step.zDivStep; i++ {
			z := finalZState
			y := input + step.yAdd
			zets := [3]int{z, z - y, (z - y) / 26}

			for _, z1 := range zets {
				zStates[z1*step.zDivStep+i] = true
			}

		}
	}

	a := alu(0)
	out := []int{}
	for k := range zStates {
		if k < 0 {
			continue
		}
		if k > upperZLimit {
			continue
		}
		for i := 1; i <= 9; i++ {
			a = alu(k)
			a.executeStep(step, i)
			if int(a) == finalZState {
				out = append(out, k)
			}
			a.reset()
		}
	}

	return out
}

const digitCount = 14

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
const highestBacktraceSection = numSize - backTraceCount
const leftSize = numSize - backTraceCount
const rightSize = backTraceCount

// const

type zStateLookup map[int]map[int]bool

func highestPossibleZStateAfterStep(initialZState int, step sectionConfig) int {
	zState := initialZState / step.zDivStep
	zState = zState * 26
	return zState + 9 + step.yAdd
}

// Works backwards from the final step (where Z has to be 0) and bruteforces zValues in the previous step to see what zValues leads to succesfull zValues in the next step
func findValidZStates(steps []sectionConfig) zStateLookup {
	acceptedZExitState := make(zStateLookup)
	for i := 0; i < 13; i++ {
		acceptedZExitState[i] = map[int]bool{}
	}

	zCap := [14]int{}
	lastCap := 0
	for i := 0; i < len(zCap); i++ {
		zCap[i] = highestPossibleZStateAfterStep(lastCap, steps[i])
		lastCap = zCap[i]
	}

	// The state we're looking for, 0 after step 13
	acceptedZExitState[13] = map[int]bool{0: true}

	for sectionId := 13; sectionId >= highestBacktraceSection; sectionId-- {
		step := steps[sectionId]
		validExitStates := acceptedZExitState[sectionId]

		for k := range validExitStates {
			for _, zState := range findValidStartZStates(step, k, zCap[sectionId]) {
				acceptedZExitState[sectionId-1][zState] = true
			}
		}

		log.Printf("Found %v valid exit states for section %vn", len(acceptedZExitState[sectionId-1]), sectionId-1)

	}

	log.Print(len(acceptedZExitState[highestBacktraceSection-1]))

	return acceptedZExitState
}

func bruteForce(steps []sectionConfig, validzSates zStateLookup, advanceFunc func([]int, int) bool, startDigit int) int {

	leftALU := alu(0)
	leftNum := createIntSlice(startDigit, leftSize)

	log.Println("Left slice:", leftNum[:(leftSize-1)])

	for {

		// if leftNum[7] == 9 && leftNum[6] == 9 && leftNum[5] == 9 && leftNum[4] == 9 && leftNum[3] == 9 && leftNum[2] == 9 {
		// 	log.Printf("Current leftNum: %v\n", leftNum)
		// }

		leftALU.reset()
		leftALU.executeSteps(steps[:leftSize], leftNum[:leftSize])

		_, isValidZ := validzSates[highestBacktraceSection-1][int(leftALU)]

		if isValidZ {
			log.Println("Found valid starting digits!", leftNum, int(leftALU))
			rightNum := createIntSlice(startDigit, backTraceCount)

			for {
				rightALU := leftALU //Reset to a1 state every attempt
				var sectionId int
				for sectionId = leftSize; sectionId < len(steps); sectionId++ {

					rightALU.executeStep(steps[sectionId], rightNum[sectionId-leftSize])

					if sectionId == 13 && int(rightALU) == 0 {
						log.Printf("Found result: %v %v", leftNum, rightNum)
						return digitSliceToInt(append(leftNum[:], rightNum[:]...))
					}

				}
				if advanceFunc(rightNum[:], rightSize-1) {
					break
				}
			}
			panic("Somehow failed!")
		}

		if advanceFunc(leftNum[:], leftSize-1) {
			break
		}

	}

	return -1
}

func digitSliceToInt(digits []int) int {
	out := 0
	for i := len(digits) - 1; i >= 0; i-- {
		out = out + digits[i]*int(math.Pow(10, float64(len(digits)-i-1)))
	}
	return out
}

func Solve() (int, int, error) {
	sectionFile, err := os.ReadFile("./day24/input.txt")
	if err != nil {
		return -1, -1, err
	}
	_, _, steps, _ := programs(sectionFile)

	// Calculate z targets, valid for both attempts
	zStates := findValidZStates(steps)

	upper := bruteForce(steps, zStates, decrementNum, 9)
	lower := bruteForce(steps, zStates, incrementNum, 1)

	return upper, lower, nil
}
