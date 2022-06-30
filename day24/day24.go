package day24

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var zGates = map[int]int{
	7:  308915776,
	8:  11881376,
	9:  456976,
	10: 456976,
	11: 17576,
	12: 676,
	13: 26,
}

type alu [4]int

func (a *alu) executeSteps(steps []sectionConfig, inputs []int) {
	for i, section := range steps {
		sectionStep(a, section, inputs[i])
	}
}
func (a *alu) executeStep(step sectionConfig, in int) {
	sectionStep(a, step, in)
}

func (a *alu) executeInstructions(instructions []instruction, inputs []int) error {
	ib := inputBuffer{inputs: inputs}

	for _, instruction := range instructions {
		_, isInput := instruction.(input)
		if isInput && ib.done {
			return errors.New("Input buffer exceeded")
		}

		instruction.execute(a, &ib)
	}

	return nil
}

func programs(file []byte) (wholeProgram []instruction, sections [][]instruction, steps []sectionConfig, err error) {
	fileStr := strings.TrimSpace(string(file))

	wholeProgram = parseInstructions(strings.Split(fileStr, "\n"))

	strSections := strings.Split(strings.TrimSpace(string(file)), "inp w\n")[1:]
	sections = [][]instruction{}
	steps = []sectionConfig{}

	for _, section := range strSections {
		lines := strings.Split(strings.TrimSpace(section), "\n")
		lines = append([]string{"inp w"}, lines...)
		steps = append(steps, parseSectionConfig(lines))
		sections = append(sections, parseInstructions(lines))
	}

	return wholeProgram, sections, steps, nil
}

func (a *alu) reset() {
	a[0] = 0
	a[1] = 0
	a[2] = 0
	a[3] = 0
}

type inputBuffer struct {
	currentIndex int
	inputs       []int
	done         bool
}

type operand interface {
	getValue(a alu) int
	getIndex() int
}

type literal int

func (l literal) getValue(a alu) int {
	return int(l)
}
func (l literal) getIndex() int {
	return -1
}

type variable int

func (i variable) getValue(a alu) int {
	return a[i]
}

func (i variable) getIndex() int {
	return int(i)
}

func (buffer *inputBuffer) readInt() int {
	out := buffer.inputs[buffer.currentIndex]
	buffer.currentIndex++

	if buffer.currentIndex == len(buffer.inputs) {
		buffer.done = true
	}

	return out
}

type instruction interface {
	execute(a *alu, buffer *inputBuffer)
}

type input [1]operand
type add [2]operand
type mul [2]operand
type div [2]operand
type mod [2]operand
type eql [2]operand

func (i input) execute(a *alu, buffer *inputBuffer) {
	(*a)[i[0].getIndex()] = buffer.readInt()
}

func (i add) execute(a *alu, buffer *inputBuffer) {
	a[i[0].getIndex()] = i[0].getValue(*a) + i[1].getValue(*a)
}
func (i mul) execute(a *alu, buffer *inputBuffer) {
	a[i[0].getIndex()] = i[0].getValue(*a) * i[1].getValue(*a)
}
func (i div) execute(a *alu, buffer *inputBuffer) {
	a[i[0].getIndex()] = i[0].getValue(*a) / i[1].getValue(*a)
}
func (i mod) execute(a *alu, buffer *inputBuffer) {
	a[i[0].getIndex()] = i[0].getValue(*a) % i[1].getValue(*a)
}
func (i eql) execute(a *alu, buffer *inputBuffer) {
	if i[0].getValue(*a) == i[1].getValue(*a) {
		a[i[0].getIndex()] = 1
	} else {
		a[i[0].getIndex()] = 0
	}
}

var variableMap = map[string]int{
	"x": 0,
	"y": 1,
	"z": 2,
	"w": 3,
}

func strToOperand(str string) operand {
	varIndex, found := variableMap[str]
	if found {
		return variable(varIndex)
	} else {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return literal(i)
	}
}

func parseInstructions(str []string) []instruction {
	out := make([]instruction, len(str))
	for i, v := range str {
		lineSplit := strings.Split(v, " ")

		instructionName := lineSplit[0]

		var op1 = strToOperand(lineSplit[1])
		var op2 operand

		if len(lineSplit) == 3 {
			op2 = strToOperand(lineSplit[2])
		}

		switch instructionName {
		case "inp":
			out[i] = input{op1}
		case "add":
			out[i] = add{op1, op2}
		case "mul":
			out[i] = mul{op1, op2}
		case "div":
			out[i] = div{op1, op2}
		case "mod":
			out[i] = mod{op1, op2}
		case "eql":
			out[i] = eql{op1, op2}
		}
	}

	return out
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
	a[3] = input
	x := (a[2] % 26) + s.xAdd
	a[2] = a[2] / s.zDivStep
	if x == input {
		x = 0
	} else {
		x = 1
	}
	y := 25
	y = y * x
	y++
	a[2] = a[2] * y
	y = input + s.yAdd
	y = y * x
	a[2] = a[2] + y

	a[0] = x
	a[1] = y
}

// func isValidNumber(model []int, program []instruction) bool {
// 	a := alu{}
// }

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
const highestBacktraceSection = 8
const leftSize = numSize - highestBacktraceSection

// const

type zStateLookup map[int]map[int]bool

// Works backwards from the final step (where Z has to be 0) and bruteforces zValues in the previous step to see what zValues leads to succesfull zValues in the next step
func findValidZStates(steps []sectionConfig) zStateLookup {
	acceptedZExitState := make(zStateLookup)
	for i := 0; i < 13; i++ {
		acceptedZExitState[i] = map[int]bool{}

	}

	zUpperLimit := 1
	acceptedZExitState[13] = map[int]bool{0: true}

	a0 := alu{}

	for sectionId := 13; sectionId >= highestBacktraceSection; sectionId-- {
		step := steps[sectionId]
		zUpperLimit = zUpperLimit * step.zDivStep // A high devisor increases bruteforce area
		for input := 1; input <= 9; input++ {
			for zState := 0; zState < zUpperLimit; zState++ {

				a0[2] = zState
				a0.executeStep(step, input)
				_, isValid := acceptedZExitState[sectionId][a0[2]]
				if isValid {
					acceptedZExitState[sectionId-1][zState] = true
				}
				a0.reset()
			}

		}

		log.Printf("Found %v valid exit states for section %v while Zlimit is %v\n", len(acceptedZExitState[sectionId-1]), sectionId-1, zUpperLimit)

	}

	log.Print(len(acceptedZExitState[highestBacktraceSection-1]))

	return acceptedZExitState
}

func bruteForce(steps []sectionConfig, validzSates zStateLookup, advanceFunc func([]int, int) bool, startDigit int) int {

	leftALU := alu{}
	leftNum := createIntSlice(startDigit, 8)

	log.Println("Left slice:", leftNum[:7])

	for {

		if leftNum[7] == 9 && leftNum[6] == 9 && leftNum[5] == 9 && leftNum[4] == 9 && leftNum[3] == 9 && leftNum[2] == 9 {
			log.Printf("Current leftNum: %v\n", leftNum)
		}

		leftALU.reset()
		leftALU.executeSteps(steps[:8], leftNum[:8])

		_, isValidZ := validzSates[7][leftALU[2]]

		if isValidZ {
			log.Println("Found valid starting digits!", leftNum, leftALU[2])
			rightNum := createIntSlice(startDigit, 6)

			for {
				rightALU := leftALU //Reset to a1 state every attempt
				var sectionId int
				for sectionId = 8; sectionId < len(steps); sectionId++ {

					rightALU.executeStep(steps[sectionId], rightNum[sectionId-8])

					if sectionId == 13 && rightALU[2] == 0 {
						log.Printf("Found result: %v %v", leftNum, rightNum)
						return digitSliceToInt(append(leftNum[:], rightNum[:]...))
					}

				}
				if advanceFunc(rightNum[:], 5) {
					break
				}
			}
			log.Println("Somehow failed!")
		}

		if advanceFunc(leftNum[:], 7) {
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

func Solve() (int, int) {
	sectionFile, err := os.ReadFile("./day24/input.txt")
	if err != nil {
		panic(err)
	}
	_, _, steps, _ := programs(sectionFile)

	// Calculate z targets, valid for both attempts
	zStates := findValidZStates(steps)

	upper := bruteForce(steps, zStates, decrementNum, 9)
	lower := bruteForce(steps, zStates, incrementNum, 1)

	return upper, lower
}
