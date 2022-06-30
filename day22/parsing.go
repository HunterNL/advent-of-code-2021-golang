package day22

import (
	"strconv"
	"strings"
)

func parseRange(str string) (min, max int) {
	leftStr, rightStr, _ := strings.Cut(strings.TrimLeft(str, "xyz="), "..")
	left, _ := strconv.Atoi(leftStr)
	right, _ := strconv.Atoi(rightStr)

	if left < right {
		min = left
		max = right
	} else {
		max = left
		min = right
	}

	return min, max
}

func parseCuboid(str string) cuboid {
	axis := strings.Split(str, ",")

	xMin, xMax := parseRange(axis[0])
	yMin, yMax := parseRange(axis[1])
	zMin, zMax := parseRange(axis[2])

	return cuboid{
		line{xMin, xMax},
		line{yMin, yMax},
		line{zMin, zMax},
	}

}

func parseStep(str string) step {
	stateStr, vectors, _ := strings.Cut(str, " ")

	return step{
		state: stateStr == "on",
		area:  parseCuboid(vectors),
	}

}

func parseInstructions(str string) instructions {
	lines := strings.Split(str, "\n")
	ins := make(instructions, len(lines))

	for i, str := range lines {
		ins[i] = parseStep(str)
	}

	return ins
}
