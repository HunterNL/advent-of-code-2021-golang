package day14

import (
	"aoc2021/file"
	"testing"
)

func TestApplyRule(t *testing.T) {
	seed, rules := parseLines(file.ReadFile("./test-input.txt"))

	if seed != "NNCB" {
		t.Errorf("Parsing error, expected NNCB but got %v\n", seed)
	}

	seed = step(seed, &rules)
	if seed != "NCNBCHB" {
		t.Errorf("Step error, exepected NCNBCHB but got %v\n", seed)
	}

	seed = step(seed, &rules)
	if seed != "NBCCNBBBCBHCB" {
		t.Errorf("Step error, exepected NBCCNBBBCBHCB but got %v\n", seed)
	}

	seed = step(seed, &rules)
	if seed != "NBBBCNCCNBBNBNBBCHBHHBCHB" {
		t.Errorf("Step error, exepected NBBBCNCCNBBNBNBBCHBHHBCHB but got %v\n", seed)
	}
	seed = step(seed, &rules)
	if seed != "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB" {
		t.Errorf("Step error, exepected NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB but got %v\n", seed)
	}
}

func mapsMatch(small map[byte]int, large map[byte]int) bool {
	for k, v := range small {
		if large[k] != v {
			return false
		}
	}

	return true
}

func TestComplexCount(t *testing.T) {
	seed, rules := parseLines(file.ReadFile("./test-input.txt"))

	c := getCounts(seed)
	pairs := buildPairMap(&rules)

	for i := 0; i < 40; i++ {
		c = smartStep(&c, &rules, &pairs)
	}

	mostCommon2 := mostCommon2(&c)
	leastCommon2 := leastCommon2(&c)

	if mostCommon2-leastCommon2 != 2188189693529 {
		t.Errorf("Expected difference to be 2188189693529 but it was %v\nMost common:%v\nLeast Common:%v\n", mostCommon2-leastCommon2, mostCommon2, leastCommon2)
	}

	// c = smartStep(&c, &rules, &pairs)
	// validCounts := getCounts("NCNBCHB")
	// if !mapsMatch(c.elements, validCounts.elements) {
	// 	t.Errorf("Expected c and newc to deep equal\nExpected: %v\nReceived: %v\n", validCounts, c.elements)
	// }

	// c = smartStep(&c, &rules, &pairs)
	// validCounts = getCounts("NBCCNBBBCBHCB")
	// if !mapsMatch(c.elements, validCounts.elements) {
	// 	t.Errorf("Expected c and newc to deep equal\nExpected: %v\nReceived: %v\n", validCounts, c.elements)
	// }
	// c = smartStep(&c, &rules, &pairs)
	// validCounts = getCounts("NBBBCNCCNBBNBNBBCHBHHBCHB")
	// if !mapsMatch(c.elements, validCounts.elements) {
	// 	t.Errorf("Expected c and newc to deep equal\nExpected: %v\nReceived: %v\n", validCounts, c.elements)
	// }
	// c = smartStep(&c, &rules, &pairs)
	// validCounts = getCounts("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB")
	// if !mapsMatch(c.elements, validCounts.elements) {
	// 	t.Errorf("Expected c and newc to deep equal\nExpected: %v\nReceived: %v\n", validCounts, c.elements)
	// }
}
