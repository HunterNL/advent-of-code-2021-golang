package day24

import (
	"aoc2021/file"
	"errors"
	"fmt"
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

func BruteForceDown() int {
	sectionFile, err := os.ReadFile("./day24/input.txt")
	if err != nil {
		panic(err)
	}

	// var incIndex int

	_, sections, steps, _ := programs(sectionFile)
	zUpperLimit := 1
	acceptedZExitState := make(map[int]map[int]bool)
	for i := 0; i < 13; i++ {
		acceptedZExitState[i] = map[int]bool{}

	}

	acceptedZExitState[13] = map[int]bool{0: true}

	a0 := alu{}

	for sectionId := 13; sectionId >= highestBacktraceSection; sectionId-- {
		step := steps[sectionId]
		zUpperLimit = zUpperLimit * step.zDivStep
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

		fmt.Printf("Found %v valid exit states for section %v while Zlimit is %v\n", len(acceptedZExitState[sectionId-1]), sectionId-1, zUpperLimit)

	}

	print(len(acceptedZExitState[highestBacktraceSection-1]))

	// return

	leftALU := alu{}
	// num := [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}

	leftNum := [8]int{9, 9, 9, 9, 9, 9, 9, 9}

	fmt.Println("Left slice:", leftNum[:7])

	for leftNum != [8]int{1, 1, 1, 1, 1, 1, 1, 1} {
		// if num[13] == 9 && num[12] == 9 && num[11] == 9 && num[10] == 9 && num[9] == 9 && num[8] == 9 && num[7] == 9 {
		// 	fmt.Printf("N: %v\n", num)
		// }

		if leftNum[7] == 9 && leftNum[6] == 9 && leftNum[5] == 9 && leftNum[4] == 9 && leftNum[3] == 9 && leftNum[2] == 9 {
			fmt.Printf("Current leftNum: %v\n", leftNum)
		}

		leftALU.reset()
		leftALU.executeSteps(steps[:8], leftNum[:8])

		_, isValidZ := acceptedZExitState[7][leftALU[2]]

		if isValidZ {
			fmt.Println("Found valid starting digits!", leftNum, leftALU[2])
			rightNum := [6]int{9, 9, 9, 9, 9, 9}

			for rightNum != [6]int{1, 1, 1, 1, 1, 1} {
				rightALU := leftALU //Reset to a1 state every attempt
				var sectionId int
				for sectionId = 8; sectionId < len(sections); sectionId++ {

					rightALU.executeStep(steps[sectionId], rightNum[sectionId-8])

					// _, stillValidZ := acceptedZState[sectionId][a1[2]]

					if sectionId == 13 && rightALU[2] == 0 {
						fmt.Printf("Found result: %v %v", leftNum, rightNum)
						return digitSliceToInt(append(leftNum[:], rightNum[:]...))
					}

					// if !stillValidZ {
					// 	break
					// }
				}
				decrementNum(rightNum[:], 5)
			}
			fmt.Println("Somehow failed!")
		}

		decrementNum(leftNum[:], 7)

	}

	return -1
}

func BruteForceUp() int {
	sectionFile, err := os.ReadFile("./day24/input.txt")
	if err != nil {
		panic(err)
	}

	// var incIndex int

	_, sections, steps, _ := programs(sectionFile)
	zUpperLimit := 1
	acceptedZExitState := make(map[int]map[int]bool)
	for i := 0; i < 13; i++ {
		acceptedZExitState[i] = map[int]bool{}

	}

	acceptedZExitState[13] = map[int]bool{0: true}

	a0 := alu{}

	for sectionId := 13; sectionId >= highestBacktraceSection; sectionId-- {
		step := steps[sectionId]
		zUpperLimit = zUpperLimit * step.zDivStep
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

		fmt.Printf("Found %v valid exit states for section %v while Zlimit is %v\n", len(acceptedZExitState[sectionId-1]), sectionId-1, zUpperLimit)

	}

	print(len(acceptedZExitState[highestBacktraceSection-1]))

	// return

	leftALU := alu{}
	// num := [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}

	leftNum := [8]int{1, 1, 1, 1, 1, 1, 1, 1}

	fmt.Println("Left slice:", leftNum[:7])

	for leftNum != [8]int{9, 9, 9, 9, 9, 9, 9, 9} {
		// if num[13] == 9 && num[12] == 9 && num[11] == 9 && num[10] == 9 && num[9] == 9 && num[8] == 9 && num[7] == 9 {
		// 	fmt.Printf("N: %v\n", num)
		// }

		if leftNum[7] == 9 && leftNum[6] == 9 && leftNum[5] == 9 && leftNum[4] == 9 && leftNum[3] == 9 && leftNum[2] == 9 {
			fmt.Printf("Current leftNum: %v\n", leftNum)
		}

		leftALU.reset()
		leftALU.executeSteps(steps[:8], leftNum[:8])

		_, isValidZ := acceptedZExitState[7][leftALU[2]]

		if isValidZ {
			fmt.Println("Found valid starting digits!", leftNum, leftALU[2])
			rightNum := [6]int{1, 1, 1, 1, 1, 1}

			for rightNum != [6]int{9, 9, 9, 9, 9, 9} {
				rightALU := leftALU //Reset to a1 state every attempt
				var sectionId int
				for sectionId = 8; sectionId < len(sections); sectionId++ {

					rightALU.executeStep(steps[sectionId], rightNum[sectionId-8])

					// _, stillValidZ := acceptedZState[sectionId][a1[2]]

					if sectionId == 13 && rightALU[2] == 0 {

						fmt.Printf("Found result: %v %v", leftNum, rightNum)
						return digitSliceToInt(append(leftNum[:], rightNum[:]...))
					}

					// if !stillValidZ {
					// 	break
					// }
				}
				incrementNum(rightNum[:], 5)
			}
			fmt.Println("Somehow failed!")
		}

		incrementNum(leftNum[:], 7)

	}

	return -1
}

func BruteForce() {
	program := parseInstructions(file.ReadFile("./day24/input.txt"))
	num := [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	a1 := alu{}
	var incIndex int

	for {
		if num[13] == 9 && num[12] == 9 && num[11] == 9 && num[10] == 9 && num[9] == 9 {
			fmt.Printf("N: %v\n", num)
		}

		a1.executeInstructions(program, num[:])
		// if a1[2] == 0 {
		// 	fmt.Printf("Found %v\n", num)
		// 	break
		// }

		fmt.Println(a1)
		break
		a1.reset()

		incIndex = 13
		for incIndex >= 0 {
			num[incIndex]--
			if num[incIndex] == 0 {
				num[incIndex] = 9
				incIndex--
				continue
			}
			break
		}

	}

	// instructions := parseInstructions(file.ReadFile("./day24/input.txt"))
}

func digitSliceToInt(digits []int) int {
	out := 0
	for i := len(digits) - 1; i >= 0; i-- {
		out = out + digits[i]*int(math.Pow(10, float64(len(digits)-i-1)))
	}
	return out
}

func Solve() (int, int) {
	upper := BruteForceDown()
	lower := BruteForceUp()
	return upper, lower

	sectionFile, err := os.ReadFile("./day24/input.txt")
	if err != nil {
		panic(err)
	}

	_, sections, _, _ := programs(sectionFile)

	// defaultProgram := parseInstructions(file.ReadFile("./day24/input.txt"))

	// sections := strings.Split(strings.TrimSpace(string(sectionFile)), "inp w\n")[1:]
	// programs := [][]instruction{}

	// for _, section := range sections {
	// 	lines := strings.Split(strings.TrimSpace(section), "\n")
	// 	lines = append([]string{"inp w"}, lines...)
	// 	programs = append(programs, parseInstructions(lines))
	// }

	type aluState struct{ input, zState int }

	workingInputs := make(map[int][]int)
	workingStates := make(map[int][]int)
	workingCombo := make(map[int][]aluState)
	stateMap := make(map[int]map[int]bool)

	acceptedZState := []int{0}

	inputForZState := make(map[int]map[int]int, 0)

	for sectionId := 13; sectionId >= 0; sectionId-- {
		inputForZState[sectionId] = map[int]int{}

		workingInputs[sectionId] = make([]int, 0, 10)
		workingStates[sectionId] = make([]int, 0, 10)

		stateMap[sectionId] = make(map[int]bool)

		workingCombo[sectionId] = make([]aluState, 0, 10)

		fmt.Printf("Starting section %v\n", sectionId)

		for inputDigit := 1; inputDigit <= 9; inputDigit++ {
			// fmt.Printf("Starting input %v\n", inputDigit)
			for zState := 0; zState <= 26; zState++ {
				// fmt.Printf("Starting state %v\n", zState)
				a1 := alu{0, 0, zState, 0}
				a1.executeInstructions(sections[sectionId], []int{inputDigit})

				aluZState := a1[2]

				for _, aState := range acceptedZState {
					if aluZState == aState {
						// inputForZState[sectionId][zState] = inputDigit
						// fmt.Println("Wait wat 2")
						workingInputs[sectionId] = append(workingInputs[sectionId], inputDigit)
						workingStates[sectionId] = append(workingStates[sectionId], zState)

						workingCombo[sectionId] = append(workingCombo[sectionId], aluState{inputDigit, aState})

						stateMap[sectionId][zState] = true
						// fmt.Println("Wait a sec now")
					}
				}
				// fmt.Printf("Alu state for (I:%v, Z:%v): %v\n", inputDigit, zState, a1)
			}
		}

		acceptedZState = workingStates[sectionId]
		// fmt.Printf()

		if len(stateMap[sectionId]) == 26 {
			fmt.Println("Found section that allows all digits")
			break
		}

		fmt.Printf("New target Z states (%v): %v\n", len(acceptedZState), acceptedZState)
	}

	// a2 := alu{}

	// a2.executeInstructions(defaultProgram, []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})

	// fmt.Println(a2)

	// fmt.Print(workingInputs)
	// fmt.Print(workingStates)

	return -1, -1

}
