package day17

import (
	"aoc2021/file"
	"strconv"
	"strings"
)

func parseCoords(str string) (int, int) {
	core := strings.TrimLeft(str, "xy=")
	n1str, n2str := file.SplitOnce(core, "..")

	n1, err1 := strconv.Atoi(n1str)
	n2, err2 := strconv.Atoi(n2str)

	if err1 != nil {
		panic(err1)
	}

	if err2 != nil {
		panic(err2)
	}

	return n1, n2
}

func parseTarget(str string) target {
	xstr, ystr := file.SplitOnce(strings.TrimPrefix(str, "target area: "), ", ")
	left, right := parseCoords(xstr)
	bottom, top := parseCoords(ystr)

	return target{left: left, right: right, top: top, bottom: bottom}
}
