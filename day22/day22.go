package day22

import (
	"fmt"
	"os"
	"sort"
)

type vec3 struct{ x, y, z int }

type cuboidReactor []cuboid

type cuboid struct {
	x, y, z line
	// min, max vec3
}

type line struct {
	min, max int
}

func (l line) size() int {
	return (l.max - l.min) + 1 // [0,0] still contains the 0 cell itself
}

func (l1 line) overlaps(l2 line) bool {
	if l2.min > l1.max {
		return false
	}
	if l2.max < l1.min {
		return false
	}

	return true
}

func (l1 line) equals(l2 line) bool {
	return l1.min == l2.min && l1.max == l2.max
}

func sort2Int(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

// Safely create a line with min/max in the righ order
func newLine(a, b int) line {
	if a < b {
		return line{a, b}
	} else {
		return line{b, a}
	}
}

func (container line) envelopes(containee line) bool {
	return container.min < containee.min && container.max > containee.max
}

// Returns 1 or more lines with no overlap
// Results cannot overlap, and must contain cut.min and cut.max
func split(primary, cut line) []line {
	if !primary.overlaps(cut) {
		return []line{primary, cut}
	}

	if primary.equals(cut) {
		return []line{primary}
	}

	if cut.envelopes(primary) {
		return []line{primary}
		// return []line{{cut.min, primary.min - 1}, primary, {primary.max + 1, cut.max}}
	}

	if cut.min <= primary.min && cut.max >= primary.max {
		return []line{primary}
	}

	if primary.envelopes(cut) {
		return []line{{primary.min, cut.min - 1}, cut, {cut.max + 1, primary.max}}
	}

	if primary.min == cut.min {
		if cut.max > primary.max {
			return []line{primary}
		} else {
			return []line{cut, {cut.max + 1, primary.max}}
		}
	}

	if primary.max == cut.max {
		if cut.min < primary.min {
			return []line{primary}
		} else {
			return []line{{primary.min, cut.min - 1}, cut}
		}
	}

	if primary.max == cut.min {
		return []line{{primary.min, primary.max - 1}, {primary.max, primary.max}}
	}

	if cut.max == primary.min {
		return []line{{primary.min, primary.min}, {primary.min + 1, primary.max}}
	}

	// Last case, overlap with no common points
	// Just take all points and sort em

	cutFirst := cut.min < primary.min

	ints := []int{primary.min, primary.max, cut.min, cut.max}
	sort.Ints(ints)

	if cutFirst {
		return []line{{ints[1], ints[2]}, {ints[2] + 1, ints[3]}}
	} else {
		return []line{{ints[0], ints[1] - 1}, {ints[1], ints[2]}}
	}

	// // return []line{
	// 	{ints[0], ints[1] - 1},
	// 	{ints[1], ints[2]},
	// 	{ints[2] + 1, ints[3]},
	// }

	// if primary.min < cut.min {
	// 	return []line{{primary.min, cut.min - 1}, cut}
	// }

	// if(primary.min> cut.min) {
	// 	return []line{}
	// }

	// if(cut.max > primary.max) {
	// 	return []line{}
	// }

	// if cut.max > primary.max {
	// 	return []line{cut, {cut.max + 1, primary.max}}
	// }

	// panic("Unknown state")

}

func cuboidsFromLines(xLines, yLines, zLines []line) (out []cuboid) {
	for _, zLine := range zLines {

		for _, yLine := range yLines {

			for _, xLine := range xLines {
				out = append(out, cuboid{
					x: xLine, y: yLine, z: zLine,
				})
			}
		}
	}

	return out
}

// func (c1 cuboid) overlapVolume(c2 cuboid) int {
// 	return
// }

// func (l1 line) overlapSize(l2 line) {

// }
// func (l line) length() int {
// 	return l.max - l.min
// }

type step struct {
	area  cuboid
	state bool
}

type instructions = []step

type reactor interface {
	countOn() int
	applyStep(step)
}

func (c cuboid) includes(v vec3) bool {
	if v.x < c.x.min {
		return false
	}
	if v.x > c.x.max {
		return false
	}
	if v.y < c.y.min {
		return false
	}
	if v.y > c.y.max {
		return false
	}
	if v.z < c.z.min {
		return false
	}
	if v.z > c.z.max {
		return false
	}

	return true
}

func paranoidCuboidOverlapCheck(r []cuboid) {
	for i1, c1 := range r {
		for i2, c2 := range r {
			if i1 == i2 {
				continue
			}

			if c1.overlaps(c2) {
				panic("Overlap!")
			}
		}
	}
}

func paranoidLineOverlapCheck(axis [][]line, cut cuboid) {
	for i, cutLine := range []line{cut.x, cut.y, cut.z} {
		didOverlap := false
		for _, line := range axis[i] {
			if line.overlaps(cutLine) {
				if didOverlap {
					axisName := "xyz"[i]
					panic("Multiple lines overlap!" + string(axisName))
				} else {
					didOverlap = true
				}
			}
		}
	}
}

func carveCuboid(primary, cut cuboid) []cuboid {
	xLines := split(primary.x, cut.x)
	yLines := split(primary.y, cut.y)
	zLines := split(primary.z, cut.z)

	// paranoidLineOverlapCheck([][]line{xLines, yLines, zLines}, cut)

	cuboids := cuboidsFromLines(xLines, yLines, zLines)

	// paranoidCuboidOverlapCheck(cuboids)

	out := []cuboid{}
	sum := 0

	for _, c := range cuboids {

		if !c.overlaps(cut) && c.overlaps(primary) {
			out = append(out, c)
			sum += c.volume()
		}
	}

	if sum > primary.volume() {
		panic("Volume grew?!")
	}

	// fmt.Printf("Input|Cut|Final %v|%v|%v\n", primary.volume(), cut.volume(), sum)

	return out
}

func (c1 cuboid) overlaps(c2 cuboid) bool {
	return c1.x.overlaps(c2.x) && c1.y.overlaps(c2.y) && c1.z.overlaps(c2.z)
}

func (c cuboid) volume() int {
	return c.x.size() * c.y.size() * c.z.size()
	// return (c.max.x - c.min.x + 1) * (c.max.y - c.min.y + 1) * (c.max.y - c.min.y + 1)
}

func (b *cuboidReactor) countOn() int {
	sum := 0
	for _, c := range *b {
		sum += c.volume()
	}
	return sum
}

func (b *cuboidReactor) applyStep(s step) {

	*b = removeOverlap(*b, s.area)

	if s.state {
		*b = append(*b, s.area)
	}
}

func (c cuboid) String() string {
	return fmt.Sprintf("[%v | %v]", c.x.min, c.x.max)
}

func removeOverlap(reactor cuboidReactor, newArea cuboid) cuboidReactor {
	for i := 0; i < len(reactor); i++ {
		reactorCube := reactor[i]
		if reactorCube.overlaps(newArea) {
			// fmt.Printf("Existing cube %v [%v] overlaps with %v, splitting\n", reactorCube, i, newArea)

			// Remove cube from reactor
			reactor[i] = reactor[len(reactor)-1]
			reactor = reactor[:len(reactor)-1]

			newCuboids := carveCuboid(reactorCube, newArea)
			for _, cubelet := range newCuboids {

				// Add every cubelet not overlapping the area to clear
				if !cubelet.overlaps(newArea) {
					// fmt.Printf("Adding %v to reactor\n", cubelet)
					reactor = append(reactor, cubelet)
				} else {
					// fmt.Printf("Discarding %v\n", cubelet)
				}
			}

			//New area may overlap with multiple existing cubes, and we've messed with the array we're looping over, just restart
			i = -1 //Note the increment!
		}
	}

	return reactor
}

func Solve() {
	bytes, err := os.ReadFile("./day22/input.txt")

	if err != nil {
		panic(err)
	}

	instructions := parseInstructions(string(bytes))

	r := cuboidReactor{}
	for i, step := range instructions {
		r.applyStep(step)
		fmt.Printf("Step %v, reactor size: %v\n", i, len(r))
	}

	sum := r.countOn()

	fmt.Println("Count on:", sum)
}
