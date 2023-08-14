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
		n := numberFromString(input)
		reduceOnce(n)
		if n.String() != expected {
			t.Errorf("Expected %v but received %v\n", expected, n)
		}
	}
}

func TestSplit(t *testing.T) {
	cases := map[string]string{
		"[[[[0,7],4],[15,[0,13]]],[1,1]]":    "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]": "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
	}

	for input, expected := range cases {
		n := numberFromString(input)
		reduceOnce(n)
		if n.String() != expected {
			t.Errorf("Expected %v but received %v\n", expected, n)
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
		result := numberFromString(input).Magnitude()
		if result != expected {
			t.Errorf("Expected %v but received %v\n", expected, result)
		}
	}
}

func convertStringSlice(sl []string) []number {
	out := []number{}
	for _, s := range sl {
		out = append(out, numberFromString(s))
	}
	return out
}

func TestSimpleSum(t *testing.T) {
	lines := convertStringSlice([]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"})
	n := sum(lines)

	if n.String() != "[[[[1,1],[2,2]],[3,3]],[4,4]]" {
		t.Error("Failed simple sum")
	}
}

func TestReplaceDigit(t *testing.T) {
	n := numberFromString("[1,2]")
	replaceNumber(n.Left(), &digit{digit: 3, parentPair: nil})

	sanityCheck(n)
}

func TestReplacePair(t *testing.T) {
	n := numberFromString("[[1,2],3]")
	sanityCheck(n)
	replaceNumber(n.Left(), createPair(&digit{digit: 3}, &digit{digit: 4}, n.Left().Parent()))

	sanityCheck(n)

	if n.String() != "[[3,4],3]" {
		t.Fail()
	}
}

func TestReducingSum(t *testing.T) {
	lines := convertStringSlice([]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"})
	str := sum(lines).String()

	expected := "[[[[3,0],[5,3]],[4,4]],[5,5]]"
	if str != expected {
		t.Errorf("Expected %v to equal %v\n", str, expected)
	}
}
func TestComplexSum(t *testing.T) {
	str, err := file.ReadFile("./test_data.txt")

	if err != nil {
		t.Error(err)
	}
	lines := convertStringSlice(str)
	sum := sum(lines).String()
	if sum != "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]" {
		t.Error("Failed complex sum")
	}
}

// func TestReduceSequence(t *testing.T) {
// 	lines := convertStringSlice([]string{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]"})
// 	added := add(lines[0], lines[1]).String()

// 	if added != "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]" {
// 		t.Error("Failed add")
// 	}

// 	e1, _ := reduceOnce(add())
// 	if string(e1) != "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]" {
// 		t.Error("Failed explode 1")
// 	}
// 	e2, _ := reduceOnce(e1)
// 	if string(e2) != "[[[[0,7],4],[15,[0,13]]],[1,1]]" {
// 		t.Error("Failed explode 2")
// 	}
// 	s1, _ := reduceOnce(e2)
// 	if string(s1) != "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]" {
// 		t.Error("Failed split 1")
// 	}
// 	s2, _ := reduceOnce(s1)
// 	if string(s2) != "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]" {
// 		t.Error("Failed split 2")
// 	}
// 	e3, _ := reduceOnce(s2)
// 	if string(e3) != "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]" {
// 		t.Error("Failed explode 3")
// 	}

// }

func TestAll(t *testing.T) {
	str, err := file.ReadFile("./test_data2.txt")
	if err != nil {
		t.Error(err)
	}
	lines := convertStringSlice(str)
	sum := sum(lines)
	if sum.String() != "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]" {
		t.Error("Failed system sum")
	}

	if sum.Magnitude() != 4140 {
		t.Error("Failed system magnitude")
	}

}
