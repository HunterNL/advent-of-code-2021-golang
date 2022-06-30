package day22

import (
	"fmt"
	"log"
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

func (container line) envelopes(containee line) bool {
	return container.min < containee.min && container.max > containee.max
}

// Returns 1 or more lines that don't overlap `cut`
func split(primary, cut line) []line {
	if primary.envelopes(cut) {
		return []line{{primary.min, cut.min - 1}, cut, {cut.max + 1, primary.max}}
	}

	cutFirst := cut.min < primary.min

	ints := []int{primary.min, primary.max, cut.min, cut.max}
	sort.Ints(ints)

	if cutFirst {
		return []line{{ints[1], ints[2]}, {ints[2] + 1, ints[3]}}
	} else {
		return []line{{ints[0], ints[1] - 1}, {ints[1], ints[2]}}
	}

}

// Creates a slice of cuboids out of every combination of gives lines
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

// Returns a list of cuboids that make up the given `primary` cuboid, with any
func carveCuboid(primary, cut cuboid) []cuboid {
	xLines := split(primary.x, cut.x)
	yLines := split(primary.y, cut.y)
	zLines := split(primary.z, cut.z)

	// paranoidLineOverlapCheck([][]line{xLines, yLines, zLines}, cut)

	cuboids := cuboidsFromLines(xLines, yLines, zLines)

	// paranoidCuboidOverlapCheck(cuboids)

	out := make([]cuboid, 0, len(xLines)*len(yLines)*len(zLines))
	sum := 0

	// Only include cuboids that make up the base cuboid
	for _, c := range cuboids {
		if c.overlaps(primary) {
			out = append(out, c)
			sum += c.volume()
		}
	}

	if sum > primary.volume() {
		panic("Volume grew?!") //Sanity check
	}

	return out
}

func (c1 cuboid) overlaps(c2 cuboid) bool {
	return c1.x.overlaps(c2.x) && c1.y.overlaps(c2.y) && c1.z.overlaps(c2.z)
}

func (c cuboid) volume() int {
	return c.x.size() * c.y.size() * c.z.size()
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

//Removes any areas overlapping the given area from the reactor
func removeOverlap(reactor cuboidReactor, clearArea cuboid) cuboidReactor {

	for i := 0; i < len(reactor); i++ {
		reactorCube := reactor[i]

		if reactorCube.overlaps(clearArea) {

			// Remove existing cube from reactor
			reactor[i] = reactor[len(reactor)-1]
			reactor = reactor[:len(reactor)-1]

			// Carve up the removed cube into smaller 'cubelets' making up the same area
			newCuboids := carveCuboid(reactorCube, clearArea)
			for _, cubelet := range newCuboids {

				// Add every cubelet not overlapping the area to clear
				if !cubelet.overlaps(clearArea) {
					reactor = append(reactor, cubelet)
				}
				// Cubelets not used added are discarded
			}

			//New area may overlap with multiple existing cubes, and we've messed with the array we're looping over, just restart
			i = -1 //Note the increment!
		}
	}

	return reactor
}

func Solve() (int, int) {
	bytes, err := os.ReadFile("./day22/input.txt")

	if err != nil {
		panic(err)
	}

	instructions := parseInstructions(string(bytes))

	p1r := vector3DReactor{}
	for _, step := range instructions {
		p1r.applyStep(step)
	}

	p1 := p1r.countOn()

	r := cuboidReactor{}
	for i, step := range instructions {
		r.applyStep(step)
		log.Printf("Step %v, reactor size: %v\n", i, len(r))
	}

	sum := r.countOn()

	log.Println("Count on:", sum)

	return p1, sum
}
