package day14

import (
	"aoc2021/file"
	"fmt"
	"math"
	"strings"
)

type counts struct {
	elements map[byte]int
	pairs    map[string]int
}

func getCounts(str string) counts {
	c := counts{
		make(map[byte]int),
		make(map[string]int),
	}

	for i := 0; i < len(str)-1; i++ {
		c.elements[str[i]]++
		c.pairs[str[i:i+2]]++
	}

	c.elements[str[len(str)-1]]++

	return c
}

func countPairs(str string) map[string]int {
	pairs := make(map[string]int)

	for i := 0; i < len(str)-2; i++ {
		pairs[str[i:i+2]]++
	}

	return pairs
}

func countElements(str string) map[rune]int {
	counts := make(map[rune]int)

	for _, b := range str {
		counts[b]++
	}

	return counts
}

func mostCommon(counts *map[rune]int) int {
	max := 0
	for _, n := range *counts {
		if n > max {
			max = n
		}
	}

	return max
}
func leastCommon(counts *map[rune]int) int {
	min := math.MaxInt
	for _, n := range *counts {
		if n < min {
			min = n
		}
	}

	return min
}

func mostCommon2(c *counts) int {
	max := 0
	for _, n := range c.elements {
		if n > max {
			max = n
		}
	}

	return max
}

func leastCommon2(c *counts) int {
	min := math.MaxInt
	for _, n := range c.elements {
		if n < min {
			min = n
		}
	}

	return min
}

func parseLines(lines []string) (seed string, rules map[string]string) {
	rules = make(map[string]string)
	seed = lines[0]

	ruleLines := lines[2:]
	for _, l := range ruleLines {
		pair, insert := file.SplitOnce(l, " -> ")
		rules[pair] = insert
	}

	return seed, rules
}

func step(seed string, rules *map[string]string) string {
	sb := strings.Builder{}

	for i := 0; i < len(seed)-1; i++ {
		rule := seed[i : i+2]
		extraChar, found := (*rules)[rule]

		sb.WriteByte(rule[0])

		if found {
			sb.WriteByte(extraChar[0])
		}
	}

	sb.WriteByte(seed[len(seed)-1])

	return sb.String()
}

func buildPairMap(rules *map[string]string) map[string][2]string {
	pairMap := make(map[string][2]string)

	for pair, extra := range *rules {
		pairMap[pair] = [2]string{
			string(pair[0]) + extra,
			extra + string(pair[1]),
		}
	}

	return pairMap
}

func smartStep(c *counts, rules *map[string]string, pairMap *map[string][2]string) counts {
	newc := counts{
		elements: c.elements,
		pairs:    make(map[string]int),
	}

	for pair, extra := range *rules {
		activePairs := c.pairs[pair]

		newPairs := (*pairMap)[pair]

		newc.elements[extra[0]] += activePairs
		newc.pairs[newPairs[0]] += activePairs
		newc.pairs[newPairs[1]] += activePairs
	}

	return newc
}

func Solve() (int, int) {
	seed, rules := parseLines(file.ReadFile("./day14/input.txt"))

	// for i := 0; i < 10; i++ {
	// 	seed = step(seed, &rules)
	// }

	// pairs := countPairs(seed)

	// fmt.Printf("%v\n", pairs)
	// fmt.Printf("%v\n", len(pairs))

	// counts := countElements(seed)
	// mostCommon := mostCommon(&counts)
	// leastCommon := leastCommon(&counts)

	// fmt.Printf("Most common: %v least common: %v solution: %v\n", mostCommon, leastCommon, mostCommon-leastCommon)

	// Part 2
	c := getCounts(seed)
	pairMap := buildPairMap(&rules)

	for i := 0; i < 10; i++ {
		c = smartStep(&c, &rules, &pairMap)
		// println(i)
	}

	part1 := mostCommon2(&c) - leastCommon2(&c)

	// fmt.Printf("Elements: %v\n", c.elements)

	for i := 0; i < 30; i++ {
		c = smartStep(&c, &rules, &pairMap)
		// println(i)
	}

	mostCommon2 := mostCommon2(&c)
	leastCommon2 := leastCommon2(&c)

	fmt.Printf("Most common: %v least common: %v solution: %v\n", mostCommon2, leastCommon2, mostCommon2-leastCommon2)

	return part1, mostCommon2 - leastCommon2
}
