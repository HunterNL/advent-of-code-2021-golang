package day24

import (
	"aoc2021/file"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestInput(t *testing.T) {
	a := alu{}

	programStr := []string{
		"inp x",
		"mul x -1",
	}

	program := parseInstructions(programStr)

	a.executeInstructions(program, []int{5})

	if a[0] != -5 {
		t.Errorf("Expected alu's c to be -5, not %v (%v)", a[0], a)
	}
}

func TestCompare(t *testing.T) {

	program := parseInstructions([]string{
		"inp z",
		"inp x",
		"mul z 3",
		"eql z x",
	})

	a1 := alu{}
	a1.executeInstructions(program, []int{3, 9})

	if a1[2] != 1 {
		t.Error("Expected a[2] to be 1")
	}

	a2 := alu{}
	a2.executeInstructions(program, []int{3, 10})

	if a2[2] != 0 {
		t.Error("Expected a[2] to be 0")
	}

}

func TestIntSlice(t *testing.T) {
	expected := []int{1, 2, 5, 8}
	actual := sliceInt(1258)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expexted %v not %v\n", expected, actual)
	}
}

func TestMonad(t *testing.T) {
	program := parseInstructions(file.ReadFile("./input.txt"))

	a1 := alu{}
	a1.executeInstructions(program, sliceInt(13579246899999))

	if a1[2] != 0 {
		t.Error()
	}
}

func TestMonadInputs(t *testing.T) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	sections := strings.Split(strings.TrimSpace(string(file)), "inp w\n")[1:]
	programs := [][]instruction{}

	for _, section := range sections {
		lines := strings.Split(strings.TrimSpace(section), "\n")
		lines = append([]string{"inp w"}, lines...)
		programs = append(programs, parseInstructions(lines))
	}

	for inputDigit := 1; inputDigit <= 9; inputDigit++ {
		for zState := 0; zState <= 25; zState++ {
			a1 := alu{0, 0, zState, 0}
			a1.executeInstructions(programs[13], []int{inputDigit})
			log.Printf("Alu state for (I:%v, Z:%v): %v\n", inputDigit, zState, a1)
		}

	}

	// log.Print(sections)

	t.Fail()
}

func TestSections(t *testing.T) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	_, sections, _, _ := programs(file)

	a := alu{}

	for i := 1; i <= 9; i++ {
		for i2 := 1; i2 <= 9; i2++ {
			for i3 := 1; i3 <= 9; i3++ {
				for i4 := 1; i4 <= 9; i4++ {
					a.executeInstructions(sections[0], []int{i, i2, i3, i4})
					log.Printf("ALU state: %v,\n", a)
					a.reset()
				}

			}

		}

	}

	t.Fail()

}

func TestFinalSection(t *testing.T) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	_, sections, _, _ := programs(file)

	program := sections[13]

	for zState := 0; zState < 27; zState++ {
		for i := 1; i <= 9; i++ {
			a1 := alu{}
			a1[2] = zState + 26
			a1.executeInstructions(program, []int{i})
			log.Printf("ALU State for (I:%v Z:%v) %v\n", i, zState, a1)
		}
	}

	t.Fail()
}

func TestStepParity(t *testing.T) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	input := [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}

	program, _, steps, _ := programs(file)

	oldAlu := alu{}
	oldAlu.executeInstructions(program, input[:])

	newAlu := alu{}
	newAlu.executeSteps(steps, input[:])

	if oldAlu != newAlu {
		t.Error("Expected ALU result to be identical")
	}

}
