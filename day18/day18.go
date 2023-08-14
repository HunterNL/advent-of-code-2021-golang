package day18

import (
	"aoc2021/file"
	"aoc2021/scanner"
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"unicode"
)

type numberPair struct {
	left   number
	right  number
	parent *numberPair
}

type digit struct {
	digit      int
	parentPair *numberPair
}

type number interface {
	IsPair() bool
	Left() number
	Right() number
	Parent() *numberPair
	String() string
	Magnitude() int
	HasParent() bool
	SetParent(n number)
	Clone(newParent *numberPair) number
}

func (d *digit) IsPair() bool {
	return false
}

func (d *digit) Left() number {
	return d
}

func (d *digit) Right() number {
	return d
}

func (d *digit) Parent() *numberPair {
	return d.parentPair
}
func (d *digit) HasParent() bool {
	return d.parentPair != nil
}
func (d *digit) SetParent(n number) {
	pair, ok := n.(*numberPair)
	if !ok {
		panic("cast failed")
	}
	d.parentPair = pair
}

func (d *digit) String() string {
	return strconv.FormatInt(int64(d.digit), 10)
}

func (d *digit) Magnitude() int {
	return d.digit
}
func (d *digit) Clone(newParent *numberPair) number {
	a := *d
	a.parentPair = newParent
	return &a
}

func (pair *numberPair) IsPair() bool {
	return true
}

func (pair *numberPair) Left() number {
	return pair.left
}

func (pair *numberPair) Right() number {
	return pair.right
}

func (pair *numberPair) Parent() *numberPair {
	return pair.parent
}
func (pair *numberPair) HasParent() bool {
	return pair.parent != nil
}
func (pair *numberPair) Clone(newParent *numberPair) number {
	a := numberPair{}
	a.left = pair.left.Clone(&a)
	a.right = pair.right.Clone(&a)
	a.parent = newParent
	return &a
}

func (pair *numberPair) SetParent(n number) {
	parent, ok := n.(*numberPair)
	if !ok {
		panic("cast failed")
	}
	pair.parent = parent
}

func (pair *numberPair) String() string {
	return fmt.Sprintf("[%v,%v]", pair.left, pair.right)
}
func (pair *numberPair) Magnitude() int {
	return pair.left.Magnitude()*3 + pair.right.Magnitude()*2
}

func numberFromString(pairString string) number {
	buf := bytes.NewBufferString(pairString)
	s := scanner.NewScanner(buf)
	pair := &numberPair{}

	s.Expect('[')
	pair.left = parseNumberPair(s, pair)
	s.Expect(',')
	pair.right = parseNumberPair(s, pair)
	s.Expect(']')
	return pair

}

func parseDigit(s scanner.Scanner, parent *numberPair) *digit {
	n := 0
	for {
		if !unicode.IsDigit(s.Peek()) {
			break
		}
		n = n * 10
		r, _, _ := s.ReadRune()
		n = n + int(r-'0')
	}

	return &digit{digit: n, parentPair: parent}
}

func parseNumberPair(s scanner.Scanner, parent *numberPair) number {
	if unicode.IsDigit(s.Peek()) {
		// r, _, _ := s.ReadRune()
		// return &digit{digit: int(r - '0'), parentPair: parent}
		return parseDigit(s, parent)
	}
	pair := &numberPair{parent: parent}

	s.Expect('[')
	pair.left = parseNumberPair(s, pair)
	s.Expect(',')
	pair.right = parseNumberPair(s, pair)
	s.Expect(']')
	return pair
}

func splitN(n int) string {
	return fmt.Sprintf("[%v,%v]", n/2, int(math.Ceil(float64(n)/2)))
}

func findRightBound(runes []rune, index int) (int, error) {
	for i := index; i < len(runes); i++ {
		if runes[i] == ']' {
			return i, nil
		}
	}
	return -1, errors.New("closing bracket not found")
}

func findRightNeighbour(n number) *digit {
	// log.Default().Printf("Finding right neighbour of %v, parent: %v\n", n, n.Parent())
	if n.HasParent() {
		p := n.Parent()
		if p.left == n {
			return findLeftMostDigit(p.right)
		} else {
			return findRightNeighbour(p)
		}
	} else {
		return nil
	}
}

func findLeftNeighbour(n number) *digit {
	if n.HasParent() {
		p := n.Parent()
		if p.right == n {
			return findRightMostDigit(p.left)
		} else {
			return findLeftNeighbour(p)
		}
	} else {
		return nil
	}
}

func findLeftMostDigit(n number) *digit {
	if !n.IsPair() {
		d, _ := n.(*digit)
		return d
	}
	return findLeftMostDigit(n.Left())
}
func findRightMostDigit(n number) *digit {
	if !n.IsPair() {
		d, _ := n.(*digit)
		return d
	}
	return findRightMostDigit(n.Right())
}

func replaceNumber(old, new number) {
	// log.Default().Printf("Replacing %v in %v with %v\n", old, old.Parent(), new)
	p := old.Parent()
	// old.SetParent(nil)
	new.SetParent(p)
	if p.left == old {
		p.left = new
		return
	}
	if p.right == old {
		p.right = new
		return
	}
	panic("Expect old to left or right of parent")
	// log.Default().Printf("Done replacing, %v\n", new.Parent())

}

func createPair(left, right, parent number) *numberPair {
	parentPair, ok := parent.(*numberPair)
	if !ok {
		panic("cast failed")
	}

	pair := &numberPair{parent: parentPair, left: left.Clone(nil), right: right.Clone(nil)}
	pair.left.SetParent(pair)
	pair.right.SetParent(pair)

	return pair

}

func sanityCheck(n number) {
	if n.HasParent() {
		p := n.Parent()
		if n.IsPair() {
			if p.Left() != n && p.Right() != n {
				panic("Expected n to be a child of p")
			}
		}
	}
	if n.IsPair() {
		if n.Left().Parent() != n {
			panic("Expected left's parent to be n")
		}
		if n.Right().Parent() != n {
			panic("Expected right's parent to be n")
		}
		sanityCheck(n.Left())
		sanityCheck(n.Right())
	}
}

func explode(n number) {
	pair, ok := n.(*numberPair)
	if !ok {
		panic("invalid cast")
	}
	left, ok := pair.left.(*digit)
	if !ok {
		panic("invalid cast")
	}
	right, ok := pair.right.(*digit)
	if !ok {
		panic("invalid cast")
	}

	leftNeighbour := findLeftNeighbour(n)
	rightNeighbour := findRightNeighbour(n)

	if leftNeighbour != nil {
		leftNeighbour.digit = leftNeighbour.digit + left.digit
	}
	if rightNeighbour != nil {
		rightNeighbour.digit = rightNeighbour.digit + right.digit
	}

	replaceNumber(n, &digit{digit: 0, parentPair: n.Parent()})
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

func findDeepNumber(n number, currentDepth int) number {
	if !n.IsPair() {
		return nil
	} else {
		if currentDepth >= 4 {
			return n
		}
		left := findDeepNumber(n.Left(), currentDepth+1)
		if left != nil {
			return left
		}
		right := findDeepNumber(n.Right(), currentDepth+1)
		if right != nil {
			return right
		}
	}
	return nil
}

func depth(n number) int {
	i := 1
	for ; ; i++ {
		if !n.HasParent() {
			break
		} else {
			n = n.Parent()
		}
	}
	return i
}

func reduceOnce(n number) bool {
	deepn := findDeepNumber(n, 0)

	if deepn != nil {
		explode(deepn)
		return false
	}

	if highN := findMinNumber(n, 10); highN != nil {
		split(highN)
		return false
	}

	return true
}

func split(n number) {

	if n.IsPair() {
		panic("Cannot split a pair")
	}
	d := n.(*digit).digit

	left := &digit{digit: d / 2, parentPair: n.Parent()}
	right := &digit{digit: d - (d / 2), parentPair: n.Parent()}

	// replaceNumber(n, &numberPair{left: left, right: right, parent: n.Parent()})
	replaceNumber(n, createPair(left, right, n.Parent()))
}

func findMinNumber(n number, min int) number {
	if n.IsPair() {
		left := findMinNumber(n.Left(), min)
		if left != nil {
			return left
		}

		right := findMinNumber(n.Right(), min)
		if right != nil {
			return right
		}
	} else {
		digit := n.(*digit).digit
		if digit >= min {
			return n
		}
	}
	return nil
}

func reduce(n number) {
	for !reduceOnce(n) {
	}
}

func sum(numbers []number) number {
	acc := numbers[0]
	for _, s := range numbers[1:] {
		acc = add(acc, s, nil)
	}
	return acc
}

func add(a, b, parent number) number {
	p, ok := parent.(*numberPair)
	if p != nil && !ok {
		panic("no cast")
	}

	n := &numberPair{left: a, right: b, parent: p}
	a.SetParent(n)
	b.SetParent(n)

	reduce(n)

	return n
}

func Solve() (int, int, error) {
	lines, err := file.ReadFile("./day18/input.txt")
	if err != nil {
		return -1, -1, err
	}
	part1Numbers := []number{}
	for _, line := range lines {
		part1Numbers = append(part1Numbers, numberFromString(line))
	}

	total := sum(part1Numbers)
	m2 := total.Magnitude()

	log.Printf("Total magnitude: %v\n", m2)

	//Part 2
	part2Numbers := []number{}
	for _, line := range lines {
		part2Numbers = append(part2Numbers, numberFromString(line))
	}

	max := 0
	for i1, s1 := range part2Numbers {
		for i2, s2 := range part2Numbers {
			if i1 == i2 {
				continue
			}
			mag := add(s1.Clone(nil), s2.Clone(nil), nil).Magnitude()
			if mag > max {
				max = mag
			}
		}
	}

	log.Printf("Max magnitude: %v\n", max)

	return m2, max, nil
}
