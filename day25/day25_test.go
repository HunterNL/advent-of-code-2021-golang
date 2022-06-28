package day25

import (
	"log"
	"reflect"
	"testing"
)

func TestWrapRight(t *testing.T) {
	testGrid, x, y := parseGrid([]string{
		".>>",
	})

	expectedGrid, _, _ := parseGrid([]string{
		">>.",
	})

	grid, _ := testGrid.step(x, y)

	if !reflect.DeepEqual(grid, expectedGrid) {
		t.Errorf("Expected\n%v\nto equal\n%v\n", grid, expectedGrid)
	}
}

func TestWrapBottom(t *testing.T) {
	testGrid, x, y := parseGrid([]string{
		".",
		"v",
		"v",
	})

	expectedGrid, _, _ := parseGrid([]string{
		"v",
		"v",
		".",
	})

	grid, _ := testGrid.step(x, y)

	if !reflect.DeepEqual(grid, expectedGrid) {
		t.Errorf("Expected\n%v\nto equal\n%v\n", grid, expectedGrid)
	}
}

func TestWrap(t *testing.T) {
	testGrid, x, y := parseGrid([]string{
		"...>...",
		".......",
		"......>",
		"v.....>",
		"......>",
		".......",
		"..vvv..",
	})

	expectedStep1, _, _ := parseGrid([]string{
		"..vv>..",
		".......",
		">......",
		"v.....>",
		">......",
		".......",
		"....v..",
	})
	expectedStep2, _, _ := parseGrid([]string{
		"....v>.",
		"..vv...",
		".>.....",
		"......>",
		"v>.....",
		".......",
		".......",
	})
	expectedStep3, _, _ := parseGrid([]string{
		"......>",
		"..v.v..",
		"..>v...",
		">......",
		"..>....",
		"v......",
		".......",
	})

	expectedStep4, _, _ := parseGrid([]string{
		">......",
		"..v....",
		"..>.v..",
		".>.v...",
		"...>...",
		".......",
		"v......",
	})

	steps := []grid{expectedStep1, expectedStep2, expectedStep3, expectedStep4}

	for i, v := range steps {
		testGrid, _ = testGrid.step(x, y)
		if !reflect.DeepEqual(testGrid, v) {
			t.Logf("Misstep at step %v\nExpected:\n%v\nGot:\n%v", i, v, testGrid)
			t.FailNow()
		}
	}

}

func TestSteps(t *testing.T) {

	grid, x, y := parseGrid([]string{
		"v...>>.vv>",
		".vv>>.vv..",
		">>.>v>...v",
		">>v>>.>.v.",
		"v>v.vv.v..",
		">.>>..v...",
		".vv..>.>v.",
		"v.v..>>v.v",
		"....v..v.>",
	})

	expectedGrid1, _, _ := parseGrid([]string{
		"....>.>v.>",
		"v.v>.>v.v.",
		">v>>..>v..",
		">>v>v>.>.v",
		".>v.v...v.",
		"v>>.>vvv..",
		"..v...>>..",
		"vv...>>vv.",
		">.v.v..v.v",
	})

	t.Log(expectedGrid1)

	expectedGrid58, _, _ := parseGrid([]string{
		"..>>v>vv..",
		"..v.>>vv..",
		"..>>v>>vv.",
		"..>>>>>vv.",
		"v......>vv",
		"v>v....>>v",
		"vvv.....>>",
		">vv......>",
		".>v.vv.v..",
	})

	grid, _ = grid.step(x, y)

	if !reflect.DeepEqual(grid, expectedGrid1) {
		t.Logf("Expected\n%v\nto equal\n%v\n", grid, expectedGrid1)
		t.FailNow()
	}

	for i := 0; i < 57; i++ {
		grid, _ = grid.step(x, y)
	}

	if !reflect.DeepEqual(grid, expectedGrid58) {
		t.Errorf("Expected grids to equal")
	}
}

func TestGridParse(t *testing.T) {
	grid, x, y := parseGrid([]string{
		"v...>>.vv>",
		".vv>>.vv..",
		">>.>v>...v",
		">>v>>.>.v.",
		"v>v.vv.v..",
		">.>>..v...",
		".vv..>.>v.",
		"v.v..>>v.v",
		"....v..v.>",
	})

	log.Print(grid)

	t.Fail()

	if x != 10 {
		t.Errorf("Expected X of 10, not %v\n", x)
	}

	if y != 9 {
		t.Errorf("Expected Y of 9, not %v\n", y)
	}

	if !reflect.DeepEqual(grid[0], []cucumber{'v', '.', '.', '.', '>', '>', '.', 'v', 'v', '>'}) {
		t.Error("Expected first line to parse correctly")
	}
}

func TestSettle(t *testing.T) {
	grid, x, y := parseGrid([]string{
		"v...>>.vv>",
		".vv>>.vv..",
		">>.>v>...v",
		">>v>>.>.v.",
		"v>v.vv.v..",
		">.>>..v...",
		".vv..>.>v.",
		"v.v..>>v.v",
		"....v..v.>",
	})

	moved := true
	steps := 0
	for moved {
		steps++
		grid, moved = grid.step(x, y)
	}
	if steps != 58 {
		t.Errorf("Expected example to settle after 58 steps, not %v\n", steps)
	}
}
