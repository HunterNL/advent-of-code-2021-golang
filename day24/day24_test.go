package day24

import (
	"log"
	"os"
	"reflect"
	"sort"
	"testing"
)

func TestIntSlice(t *testing.T) {
	expected := []int{1, 2, 5, 8}
	actual := sliceInt(1258)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expexted %v not %v\n", expected, actual)
	}
}

func TestMonad(t *testing.T) {
	t.SkipNow()
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	_, _, steps, _ := programs(file)

	a1 := alu{state: 0}
	a1.executeSteps(steps, sliceInt(13579246899999))

	if a1.state != 0 {
		t.Error()
	}
}

func intSlicesEqual(a, b []int) bool {
	as := append([]int{}, a...)
	bs := append([]int{}, b...)
	sort.Ints(as)
	sort.Ints(bs)

	return reflect.DeepEqual(as, bs)
}

func TestBackTrace(t *testing.T) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	_, _, steps, _ := programs(file)

	targetZ := 0
	zCap := 22000

	//Bruteforce valid solutions
	a0 := alu{state: 0}

	for stepId := 13; stepId >= 0; stepId-- {
		step := steps[stepId]

		bruteForcedStates := []int{}
		for input := 1; input <= 9; input++ {
			for zState := 0; zState < zCap; zState++ {

				a0.state = zState
				a0.executeStep(step, input)
				if a0.state == targetZ {
					bruteForcedStates = append(bruteForcedStates, zState)
					t.Logf("input %v brings %v to %v\n", input, zState, targetZ)
				}
				a0.reset()
			}

		}

		calcedStates := findValidStartZStates(step, targetZ, zCap)

		calculatedIntsAsSlice := []int{}

		for _, z := range calcedStates {
			calculatedIntsAsSlice = append(calculatedIntsAsSlice, z.state)
		}

		t.Log(calcedStates)

		if !intSlicesEqual(bruteForcedStates, calculatedIntsAsSlice) {
			sort.Ints(bruteForcedStates)
			sort.Ints(calculatedIntsAsSlice)

			t.Logf("Failed on step %v\n", stepId)
			t.FailNow()
		}
	}

}

func TestSections(t *testing.T) {
	t.SkipNow() // Only used for quickly running this func in development
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	_, _, steps, _ := programs(file)

	a := alu{0}

	for i := 1; i <= 9; i++ {
		for i2 := 1; i2 <= 9; i2++ {
			for i3 := 1; i3 <= 9; i3++ {
				for i4 := 1; i4 <= 9; i4++ {
					a.executeSteps(steps, []int{i, i2, i3, i4})
					log.Printf("ALU state: %v,\n", a)
					a.reset()
				}

			}

		}

	}

	t.Fail()

}

func TestFinalSection(t *testing.T) {
	t.SkipNow() // Only used for quickly running this func in development
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	_, _, steps, _ := programs(file)

	a := alu{0}
	for zState := 0; zState < 27; zState++ {
		for i := 1; i <= 9; i++ {
			a.state = zState + 26
			a.executeSteps(steps, []int{i})
			log.Printf("ALU State for (I:%v Z:%v) %v\n", i, zState, a.state)
		}
	}

	t.Fail()
}
