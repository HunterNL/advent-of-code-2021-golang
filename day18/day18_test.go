package day18

import (
	"aoc2021/file"
	"testing"
)

// func TestAddition(t *testing.T) {

// }

/*
[[[[[9,8],1],2],3],4] becomes [[[[0,9],2],3],4] (the 9 has no regular number to its left, so it is not added to any regular number).
[7,[6,[5,[4,[3,2]]]]] becomes [7,[6,[5,[7,0]]]] (the 2 has no regular number to its right, and so it is not added to any regular number).
[[6,[5,[4,[3,2]]]],1] becomes [[6,[5,[7,0]]],3].
[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]] becomes [[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]] (the pair [3,2] is unaffected because the pair [7,3] is further to the left; [3,2] would explode on the next action).
[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]] becomes [[3,[2,[8,0]]],[9,[5,[7,0]]]].*/

func TestExplode(t *testing.T) {
	cases := map[string]string{
		"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}

	for input, expected := range cases {
		result, _ := reduceOnce(input)
		if result != expected {
			t.Errorf("Expected %v but received %v\n", expected, result)
		}
	}
}

func TestSplit(t *testing.T) {
	cases := map[string]string{
		"[[[[0,7],4],[15,[0,13]]],[1,1]]":    "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]": "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
	}

	for input, expected := range cases {
		result, _ := reduceOnce(input)
		if result != expected {
			t.Errorf("Expected %v but received %v\n", expected, result)
		}
	}
}

func TestMagnitude(t *testing.T) {
	cases := map[string]int{
		"[9,1]":                                                 29,
		"[[1,2],[[3,4],5]]":                                     143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]":                     1384,
		"[[[[1,1],[2,2]],[3,3]],[4,4]]":                         445,
		"[[[[3,0],[5,3]],[4,4]],[5,5]]":                         791,
		"[[[[5,0],[7,4]],[5,5]],[6,6]]":                         1137,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]": 3488,
	}

	for input, expected := range cases {
		result := magnitude(input)
		if result != expected {
			t.Errorf("Expected %v but received %v\n", expected, result)
		}
	}
}

func TestSum(t *testing.T) {
	lines := []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"}
	str := sum(lines)

	if str != "[[[[1,1],[2,2]],[3,3]],[4,4]]" {
		t.Error("Failed simple sum")
	}

	lines = []string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"}
	str = sum(lines)

	if str != "[[[[3,0],[5,3]],[4,4]],[5,5]]" {
		t.Error("Failed reducing sum")
	}

	lines = file.ReadFile("./test_data.txt")
	sum := sum(lines)
	if sum != "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]" {
		t.Error("Failed complex sum")
	}
}

func TestReduceSequence(t *testing.T) {
	lines := []string{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]"}
	added := add(lines[0], lines[1])

	if added != "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]" {
		t.Error("Failed add")
	}

	e1, _ := reduceOnce(added)
	if e1 != "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]" {
		t.Error("Failed explode 1")
	}
	e2, _ := reduceOnce(e1)
	if e2 != "[[[[0,7],4],[15,[0,13]]],[1,1]]" {
		t.Error("Failed explode 2")
	}
	s1, _ := reduceOnce(e2)
	if s1 != "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]" {
		t.Error("Failed split 1")
	}
	s2, _ := reduceOnce(s1)
	if s2 != "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]" {
		t.Error("Failed split 2")
	}
	e3, _ := reduceOnce(s2)
	if e3 != "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]" {
		t.Error("Failed explode 3")
	}

}

func TestAll(t *testing.T) {
	lines := file.ReadFile("./test_data2.txt")
	sum := sum(lines)
	if sum != "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]" {
		t.Error("Failed system sum")
	}

	if magnitude(sum) != 4140 {
		t.Error("Failed system magnitude")
	}

}
