package day22

import "log"

type vector3DReactor map[vec3]bool
type vector1dReactor map[int]bool
type cuboid1dReactor []cuboid

func (r *cuboid1dReactor) countOn() int {
	sum := 0

	for _, c := range *r {
		sum += c.volume()
	}

	return sum
}

func (b *cuboid1dReactor) applyStep(s step) {
	newStep := step{
		state: s.state,
		area:  cuboid{x: s.area.x, y: line{0, 0}, z: line{0, 0}},
	}

	*b = cuboid1dReactor(removeOverlap(cuboidReactor(*b), newStep.area))

	if newStep.state {
		log.Printf("Adding step %v to reactor\n", newStep.area)
		*b = append(*b, newStep.area)
	}

	paranoidCuboidOverlapCheck([]cuboid(*b))
}

func (r *vector1dReactor) countOn() int {
	sum := 0

	for _, n := range *r {
		if n {
			sum++
		}
	}

	return sum
}

func (r *vector1dReactor) applyStep(s step) {
	for i := s.area.x.min; i <= s.area.x.max; i++ {
		(*r)[i] = s.state
	}
}

func (r vector3DReactor) applyStep(s step) {
	c := s.area

	var centerArea = cuboid{
		line{-50, 50},
		line{-50, 50},
		line{-50, 50},
	}

	if !centerArea.overlaps(c) {
		return
	}

	for z := c.z.min; z <= c.z.max; z++ {
		for y := c.y.min; y <= c.y.max; y++ {
			for x := c.x.min; x <= c.x.max; x++ {
				vec := vec3{x, y, z}
				// if centerArea.includes(vec) {
				r[vec] = s.state
				// }
			}
		}
	}
}

func (r vector3DReactor) countOn() int {
	count := 0
	for _, state := range r {
		if state {
			count++
		}
	}
	return count
}
