package day18

import (
	"aoc2021/file"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func splitN(n int) string {
	return fmt.Sprintf("[%v,%v]", n/2, int(math.Ceil(float64(n)/2)))
}

func findRightBound(str string, index int) (int, error) {
	runes := []rune(str)
	for i := index; i < len(runes); i++ {
		if runes[i] == ']' {
			return i, nil
		}
	}
	return -1, errors.New("closing bracket not found")
}

func addToLeftNumber(str string, n int) string {
	intstr := ""

	runes := []rune(str)
	i := 0
	for ; i < len(runes); i++ {
		r := runes[i]
		if unicode.IsDigit(r) {
			intstr = intstr + string(r)
		} else {
			if intstr != "" {
				break // Once we leave a number, stop scanning
			}
		}
	}

	if intstr == "" {
		return str
	}

	stri, err := strconv.Atoi(intstr)
	if err != nil {
		panic(err)
	}

	leftSide := runes[0 : i-len(intstr)]
	rightSide := runes[i:]

	return string(leftSide) + fmt.Sprint(stri+n) + string(rightSide)
}
func addToRightNumber(str string, n int) string {
	intstr := ""

	runes := []rune(str)
	i := len(runes) - 1
	for ; i >= 0; i-- {
		r := runes[i]
		if unicode.IsDigit(r) {
			intstr = string(r) + intstr
		} else {
			if intstr != "" {
				break // Once we leave a number, stop scanning
			}
		}
	}

	if intstr == "" {
		return str
	}

	stri, err := strconv.Atoi(intstr)
	if err != nil {
		panic(err)
	}

	leftSide := runes[0 : i+1]
	rightSide := runes[i+len(intstr)+1:]

	return string(leftSide) + fmt.Sprint(stri+n) + string(rightSide)
}

func parsePair(str string) (int, int) {
	str = strings.Trim(str, "[]")
	as, bs := file.SplitOnce(str, ",")

	a, err := strconv.Atoi(as)
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(bs)
	if err != nil {
		panic(err)
	}

	return a, b
}

func explode(str string, leftBound int) string {
	rightBound, err := findRightBound(str, leftBound)
	if err != nil {
		panic(err)
	}

	pairLeft, pairRight := parsePair(str[leftBound:rightBound])
	leftSplit := str[:leftBound]
	rightSplit := str[rightBound+1:]

	leftSide := addToRightNumber(leftSplit, pairLeft)
	rightSide := addToLeftNumber(rightSplit, pairRight)

	return leftSide + "0" + rightSide

	// leftNumber := make([]rune, 0)
	// for i := 0; i < leftBound; i++ {
	// 	r := runes[i]
	// 	if unicode.IsDigit(r) {
	// 		leftNumber = append(leftNumber, r)
	// 	} else {
	// 		leftNumber = nil
	// 	}
	// }

}

func splitAt(str string, i int) string {
	runes := []rune(str)
	intstr := ""

	for ; i < len(runes); i++ {
		r := runes[i]
		if unicode.IsDigit(r) {
			intstr = intstr + string(r)
		} else {
			break
		}
	}

	n, err := strconv.Atoi(intstr)
	if err != nil {
		panic(err)
	}

	leftSide := runes[:i-len(intstr)]
	rightSide := runes[i:]

	return string(leftSide) + splitN(n) + string(rightSide)
}

// func reduce(a string) string {
// 	nesting := 0
// 	dirty := true
// 	for dirty {
// 		dirty = false
// 		for i, r := range a {
// 			if r == '[' {
// 				nesting++
// 				continue
// 			}
// 			if r == ']' {
// 				nesting--
// 				continue
// 			}

// 			if nesting == 5 { // Start of pair to explode
// 				dirty = true
// 				a = explode(a, i)
// 				break
// 			}
// 		}
// 	}
// 	return a
// }

func reduceOnce(a string) (string, bool) {
	nesting := 0
	strint := ""

	// Always try explode first...
	for i, r := range a {
		if r == '[' {
			nesting++
		}
		if r == ']' {
			nesting--
		}

		if nesting == 5 { // Start of pair to explode
			return explode(a, i), false
		}
	}

	/// and split second
	for i, r := range a {
		if unicode.IsDigit(r) {
			strint = strint + string(r)
		} else {
			if strint != "" {
				n, err := strconv.Atoi(strint)
				if err != nil {
					print("Parse err?", strint)
					continue
				}
				if n >= 10 {
					return splitAt(a, i-len(strint)), false
				}
			}
			strint = ""
		}
	}

	return a, true
}

func mag(s io.RuneScanner) int {
	s.ReadRune() // Initial [
	leftR, _, _ := s.ReadRune()
	var leftN int
	var rightN int

	if leftR == '[' {
		s.UnreadRune()
		leftN = mag(s)
	}

	if unicode.IsDigit(leftR) {
		var err error
		leftN, err = strconv.Atoi(string(leftR))
		if err != nil {
			panic(err)
		}
	}

	s.ReadRune() // ,

	rightR, _, _ := s.ReadRune()
	if rightR == '[' {
		s.UnreadRune()
		rightN = mag(s)
	}

	if unicode.IsDigit(rightR) {
		var err error
		rightN, err = strconv.Atoi(string(rightR))
		if err != nil {
			panic(err)
		}
	}

	s.ReadRune() // Closing ]

	return leftN*3 + rightN*2
}

func magnitude(str string) int {
	r := strings.NewReader(str)
	return mag(r)
}

func reduce(str string) string {
	done := false
	for {
		str, done = reduceOnce(str)
		if done {
			break
		}

	}
	return str
}

func sum(lines []string) string {
	acc := lines[0]
	for _, s := range lines[1:] {
		acc = add(acc, s)
		acc = reduce(acc)
	}
	return reduce(acc)
}

func add(a, b string) string {
	return "[" + a + "," + b + "]"
}

func Solve() (int, int) {
	lines := file.ReadFile("./day18/input.txt")
	total := sum(lines)
	mag := magnitude(total)

	fmt.Printf("Total magnitude: %v\n", mag)

	//Part 2
	max := 0
	for i1, s1 := range lines {
		for i2, s2 := range lines {
			if i1 == i2 {
				continue
			}
			mag := magnitude(sum([]string{s1, s2}))
			if mag > max {
				max = mag
			}
		}
	}

	fmt.Printf("Max magnitude: %v\n", max)

	return mag, max
}
