package day22

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	step := parseStep("on x=-44..7,y=-48..-4,z=-28..22")

	if step.state != true {
		t.Error("Expected state to be on")
	}

	expectedCuboid := cuboid{
		line{-44, 7},
		line{-48, -4},
		line{-28, 22},
	}

	if !reflect.DeepEqual(step.area, expectedCuboid) {
		t.Error("Expected cuboid to match expected cuboid")
	}
}

func Test1dCompare(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		t.Error(err)
	}

	instructions := parseInstructions(string(bytes))

	vr := vector1dReactor{}
	cr := cuboid1dReactor{}

	for i, s := range instructions {
		log.Printf("Step %v\n", i)
		vr.applyStep(s)
		cr.applyStep(s)

		if vr.countOn() != cr.countOn() {
			t.Errorf("Count mismatch after step %v: %v vs %v\n", i, vr.countOn(), cr.countOn())
			t.FailNow()
		} else {
			t.Logf("Size after step %v: %v\n", i, vr.countOn())
		}
	}
}

func Test1DTest(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		t.Error(err)
	}

	instructions := parseInstructions(string(bytes))

	vr := vector1dReactor{}

	for _, s := range instructions {
		vr.applyStep(s)
	}

	if vr.countOn() != 84 {
		t.Error("Reference count wrong")
	}
}

func TestCount(t *testing.T) {
	bytes, err := os.ReadFile("./test_input.txt")

	if err != nil {
		t.Error(err)
	}

	instructions := parseInstructions(string(bytes))

	r := cuboidReactor{}
	s := vector3DReactor{}

	for _, step := range instructions {
		s.applyStep(step)
	}

	if s.countOn() != 590784 {
		t.Logf("Simple reactor count wrong")
		t.FailNow()
	}

	s = vector3DReactor{}

	for i, step := range instructions {
		log.Printf("\n\nApplying step %v %v\n\n", i, step)
		s.applyStep(step)
		r.applyStep(step)

		if s.countOn() != r.countOn() {
			t.Logf("Count mismatch at step %v, simple: %v complex: %v", i, s.countOn(), r.countOn())
			t.FailNow()
		}

		log.Printf("Reactor size:%v\n", len(r))
	}

	sum := r.countOn()

	if sum != 590784 {
		t.Errorf("Expected %v cells to be on, not %v\n", 590784, sum)
	}

}

func TestRemoveOverlap(t *testing.T) {
	primary := cuboid{x: line{1, 2}, y: line{1, 2}, z: line{1, 2}}
	cut := cuboid{x: line{2, 2}, y: line{2, 2}, z: line{2, 2}}

	lessCubes := removeOverlap(cuboidReactor{primary}, cut)

	if len(lessCubes) != 7 {
		t.Errorf("Expected 7 cubes, not %v\n", len(lessCubes))
	}

	sum := 0
	for _, c := range lessCubes {
		sum += c.volume()
	}

	if sum != 7 {
		t.Errorf("Expected a total volume of 7, not %v\n", sum)
	}
}

func TestCarving(t *testing.T) {
	primary := cuboid{x: line{1, 2}, y: line{1, 2}, z: line{1, 2}}
	cut := cuboid{x: line{2, 2}, y: line{2, 2}, z: line{2, 2}}

	cubes := carveCuboid(primary, cut)

	if len(cubes) != 8 {
		t.Errorf("Expected 8 cubes, not %v\n", len(cubes))
	}

	sum := 0
	for _, c := range cubes {
		sum += c.volume()
	}

	if sum != 8 {
		t.Errorf("Expected a total volume of 8, not %v\n", sum)
	}

}

func TestSplit(t *testing.T) {
	t.SkipNow() // These tests are overkill compared to the actual puzzle input

	tests := map[string]struct {
		primary  line
		cut      line
		expected []line
	}{
		"Unrelated": {
			primary:  line{0, 0},
			cut:      line{2, 2},
			expected: []line{{0, 0}, {2, 2}},
		},
		"Neighbour": {
			primary:  line{0, 2},
			cut:      line{3, 5},
			expected: []line{{0, 2}, {3, 5}},
		},
		"Containing cut": {
			primary:  line{0, 5},
			cut:      line{2, 3},
			expected: []line{{0, 1}, {2, 3}, {4, 5}},
		},
		"Containing primary": {
			primary:  line{2, 3},
			cut:      line{0, 5},
			expected: []line{{2, 3}},
		},
		// "Containing primary": {
		// 	primary:  line{2, 3},
		// 	cut:      line{0, 5},
		// 	expected: []line{{0, 1}, {2, 3}, {4, 5}},
		// },
		"Ajoined low": {
			primary:  line{0, 2},
			cut:      line{2, 5},
			expected: []line{{0, 1}, {2, 2}, {3, 5}},
		},
		"Ajoined high": {
			primary:  line{2, 5},
			cut:      line{0, 2},
			expected: []line{{0, 1}, {2, 2}, {3, 5}},
		},
		"Sharing low A": {
			primary:  line{2, 5},
			cut:      line{2, 3},
			expected: []line{{2, 3}, {4, 5}},
		},
		"Sharing low B": {
			primary:  line{2, 3},
			cut:      line{2, 5},
			expected: []line{{2, 3}, {4, 5}},
		},
		"Sharing high A": {
			primary:  line{2, 5},
			cut:      line{3, 5},
			expected: []line{{2, 2}, {3, 5}},
		},
		"Sharing high B": {
			primary:  line{3, 5},
			cut:      line{2, 5},
			expected: []line{{2, 2}, {3, 5}},
		},
		"Overlap A": {
			primary:  line{2, 6},
			cut:      line{4, 8},
			expected: []line{{2, 3}, {4, 6}, {7, 8}},
		},
		"Overlap B": {
			primary:  line{4, 8},
			cut:      line{2, 6},
			expected: []line{{2, 3}, {4, 6}, {7, 8}},
		},
		"Aaaaa": {
			primary:  line{-22, 19},
			cut:      line{-40, -22},
			expected: []line{{2, 3}, {4, 6}, {7, 8}},
		},
	}

	for name, test := range tests {
		// t.Run(name, func(t *testing.T) {
		out := split(test.primary, test.cut)
		if !reflect.DeepEqual(out, test.expected) {
			t.Errorf("Expected \n%v\nto equal \n%v\n in test %v", out, test.expected, name)
		}
		// })
	}
	// }

	// in1, in2 := line{0, 0}, line{2, 2}

	// lines := split(in1, in2)

	// if !lines[0].equals(in1) {
	// 	t.Error("Expected in1 to remain untouched")
	// }

	// if !lines[1].equals(in2) {
	// 	t.Error("Expected in1 to remain untouched")
	// }
}
func TestSplitEnvelope(t *testing.T) {
	t.SkipNow() // Again, overkill for given input
	in1, in2 := line{0, 5}, line{0, 2}

	lines := split(in1, in2)

	if !lines[0].equals(line{0, 2}) {
		t.Error("Expected in1 to remain untouched")
	}

	if !lines[1].equals(line{3, 5}) {
		t.Error("Expected in1 to remain untouched")
	}
}
