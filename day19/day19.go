package day19

import (
	"aoc2021/intmath"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vec3 struct {
	x, y, z int
}

func (v vec3) index(n int) int {
	if n == 0 {
		return v.x
	} else if n == 1 {
		return v.y
	} else {
		return v.z
	}
}

type rotation struct {
	oX, oY, oZ int
	sX, sY, sZ int
}

var rotations = createRotations()

const commonDetectionThreshold = 12

// func (v vec3) applyRotation(r rotation) vec3 {
// 	for r.z > 0 {
// 		r.z = r.z - 1

// 		a := v.y

// 		v.y = -v.x
// 		v.x = a
// 	}

// 	for r.y > 0 {
// 		r.y = r.y - 1

// 		a := v.x

// 		v.x = -v.z
// 		v.z = a
// 	}

// 	for r.x > 0 {
// 		r.x = r.x - 1

// 		a := v.y

// 		v.y = -v.z
// 		v.z = a
// 	}

// 	return v
// }
func (v vec3) applyRotation(r rotation) (v2 vec3) {
	v2.x = v.index(r.oX) * r.sX
	v2.y = v.index(r.oY) * r.sY
	v2.z = v.index(r.oZ) * r.sZ

	return v2
}

// TODO Restrict to 24 cube rotations instead of all 64
func createRotations() []rotation {
	x, y, z := 0, 1, 2
	pos, neg := 1, -1

	rotations := []rotation{
		{x, y, z, pos, pos, pos},
		{x, z, y, pos, neg, pos},
		{x, y, z, pos, neg, neg},
		{x, z, y, pos, pos, neg},

		{x, y, z, neg, neg, pos},
		{x, z, y, neg, neg, neg},
		{x, y, z, neg, pos, neg},
		{x, z, y, neg, pos, pos},

		{y, x, z, pos, neg, pos},
		{y, z, x, pos, neg, neg},
		{y, x, z, pos, pos, neg},
		{y, z, x, pos, pos, pos},

		{y, x, z, neg, pos, pos},
		{y, z, x, neg, neg, pos},
		{y, x, z, neg, neg, neg},
		{y, z, x, neg, pos, neg},

		{z, y, x, pos, pos, neg},
		{z, x, y, pos, neg, neg},
		{z, y, x, pos, neg, pos},
		{z, x, y, pos, pos, pos},

		{z, y, x, neg, pos, pos},
		{z, x, y, neg, neg, pos},
		{z, y, x, neg, neg, neg},
		{z, x, y, neg, pos, neg},

		// {x, y, z, pos, pos, pos},
		// {x, y, z, neg, pos, pos},
		// {x, y, z, pos, neg, pos},
		// {x, y, z, neg, neg, pos},
		// {x, y, z, pos, pos, neg},
		// {x, y, z, neg, pos, neg},
		// {x, y, z, pos, neg, neg},
		// {x, y, z, neg, neg, neg},

		// {y, z, x, pos, pos, pos},
		// {y, z, x, neg, pos, pos},
		// {y, z, x, pos, neg, pos},
		// {y, z, x, neg, neg, pos},
		// {y, z, x, pos, pos, neg},
		// {y, z, x, neg, pos, neg},
		// {y, z, x, pos, neg, neg},
		// {y, z, x, neg, neg, neg},

		// {z, x, y, pos, pos, pos},
		// {z, x, y, neg, pos, pos},
		// {z, x, y, pos, neg, pos},
		// {z, x, y, neg, neg, pos},
		// {z, x, y, pos, pos, neg},
		// {z, x, y, neg, pos, neg},
		// {z, x, y, pos, neg, neg},
		// {z, x, y, neg, neg, neg},
	}
	return rotations
}

// // TODO Restrict to 24 cube rotations instead of all 64
// func createRotations() []rotation {
// 	rotations := make([]rotation, 0, 64)
// 	for x := 0; x <= 3; x++ {
// 		for y := 0; y <= 3; y++ {
// 			for z := 0; z <= 3; z++ {
// 				rot := rotation{x: x, y: y, z: z}
// 				rotations = append(rotations, rot)
// 			}
// 		}
// 	}
// 	return rotations
// }

// func findRotation(fromVec, toVec vec3) rotation {
// 	for x := 0; x <= 4; x++ {
// 		for y := 0; y <= 4; y++ {
// 			for z := 0; z <= 4; z++ {
// 				rot := rotation{x: x, y: y, z: z}
// 				if fromVec.applyRotation(rot) == toVec {
// 					return rot
// 				}

// 			}
// 		}
// 	}

// 	panic("Did not find rotation") // this shouldn't happen
// }

func (a vec3) sub(b vec3) vec3 {
	return vec3{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
	}
}
func (a vec3) add(b vec3) vec3 {
	return vec3{
		x: a.x + b.x,
		y: a.y + b.y,
		z: a.z + b.z,
	}
}
func (a vec3) invert() vec3 {
	return vec3{
		x: -a.x,
		y: -a.y,
		z: -a.z,
	}
}

func vec3equal(a, b vec3) bool {
	return a.x == b.x && a.y == b.y && a.z == b.z
}

// func sliceToVec3(a []int) vec3 {
// 	return vec3{
// 		a[0], a[1], a[2],
// 	}
// }

// func (a vec3) abs() vec3 {
// 	return vec3{
// 		intmath.Abs(a.x),
// 		intmath.Abs(a.y),
// 		intmath.Abs(a.z),
// 	}
// }

func manhattan(a, b vec3) int {
	return intmath.Distance(a.x, b.x) + intmath.Distance(a.y, b.y) + intmath.Distance(a.z, b.z)
}

// func (a vec3) manhatten() int {
// 	b := a.abs()
// 	return b.x + b.y + b.z
// }

type detection struct {
	scannerId int
	localPos  vec3
}

type scanner struct {
	id         int
	detections []detection
	rotation   rotation
	position   vec3
	resolved   bool
}

// func (s *scanner) createLinks() {
// 	for _, d := range s.detections {
// 		s.network[&d] = make(map[vec3]*detection, 0)

// 		for _, otherDetection := range s.detections {
// 			if d == otherDetection {
// 				continue
// 			}

// 			absVec := otherDetection.localPos.sub(d.localPos).abs()
// 			intslice := make([]int, 3, 3)
// 			intslice[0] = absVec.x
// 			intslice[1] = absVec.y
// 			intslice[2] = absVec.z

// 			sort.Ints(intslice)

// 			vec := sliceToVec3(intslice)

// 			fmt.Printf("Adding vector %v\n", vec)

// 			if _, found := s.network[&d][vec]; found {
// 				panic("Duplicate vector")
// 			}

// 			s.network[&d][vec] = &otherDetection
// 		}

// 	}
// }

func parseDetection(str string) ([]detection, []scanner) {
	scanstring := strings.Split(str, "\n\n")
	detections := make([]detection, 0, len(scanstring)*30)
	scanners := make([]scanner, 0, len(scanstring))

	for i, s := range scanstring {
		lines := strings.Split(s, "\n")[1:]
		scan := scanner{id: i}
		// detections := make([]vec3, 0, len(lines))
		for _, l := range lines {
			position := parseVec3(l)
			detection := detection{localPos: position, scannerId: i}
			detections = append(detections, detection)
			scan.detections = append(scan.detections, detection)
			scan.rotation = rotation{0, 1, 2, 1, 1, 1}
		}
		scanners = append(scanners, scan)
	}

	// for _, s := range scanners {
	// 	s.createLinks()
	// 	fmt.Printf("Done making links for %v\n", s.id)
	// }

	// scanners[0].rotationResolved = true // first scanner is absolute

	return detections, scanners
}

func unpack3(lines []string) (string, string, string) {
	return lines[0], lines[1], lines[2]
}

func parseVec3(line string) vec3 {
	xs, ys, zs := unpack3(strings.Split(line, ","))

	x, err := strconv.Atoi(xs)
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(ys)
	if err != nil {
		panic(err)
	}

	z, err := strconv.Atoi(zs)
	if err != nil {
		panic(err)
	}

	return vec3{x: x, y: y, z: z}
}

// func countBeacons(scans []scanner) int {
// 	links := findLinks(scans)
// 	return 0
// }

// func findLinks(scans []scanner) map[int][]edge {
// 	m := make(map[int][]edge)

// 	for _, s := range scans {
// 		for i1, d1 := range s.detections {
// 			for i2, d2 := range s.detections {
// 				if i1 == i2 {
// 					continue
// 				}
// 				dist := manhattan(d1.localPos, d2.localPos)

// 				_, found := m[dist]
// 				if !found {
// 					m[dist] = make([]edge, 0)
// 				}

// 				links := m[dist]
// 				links = append(links, edge{d1, d2})

// 				m[dist] = links
// 			}
// 		}
// 	}
// 	return m
// }

// func findLink(baseScan, b scanner) (error, rotation, vec3) {
// 	// m := make(map[int][]edge)

// 	for _, baseDetection := range baseScan.detections {
// 		n := baseScan.network[&baseDetection]

// 		for baseVector, baseNeighbour := range n {

// 			for _, d2 := range b.detections {
// 				n2 := b.network[&d2]

// 				for v2, e2 := range n2 {
// 					if v2 == baseVector {
// 						// Got a match
// 						originVec := baseNeighbour.localPos.sub(baseDetection.localPos)
// 						relVec := d2.localPos.sub(e2.localPos)

// 						return nil, findRotation(originVec, relVec), baseScan.position.add(baseDetection.localPos).sub(e2.localPos)
// 					}
// 				}
// 			}

// 		}
// 	}

// 	return errors.New("Not found"), rotation{}, vec3{}
// }

// func findCommonLinks(s1, s2 scanner) (error, edge, edge) {
// 	m := findLinks([]scanner{s1, s2})
// 	for _, edge := range m {
// 		if len(edge) >= 2 {
// 			return nil, edge[0], edge[1]
// 		}
// 	}

// 	return errors.New("Did not find common link"), edge{}, edge{}

// }

// func findScannerInSlice(scannerId int, slice []scanner) int {
// 	for i, s := range slice {
// 		if s.id == scannerId {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func getOffsets(a, b scanner) []vec3 {
// 	offsets := make([]vec3, 0, len(a.detections)*len(b.detections)*len(rotations))

// 	for _, d1 := range a.detections {
// 		for _, d2 := range b.detections {
// 			for _, rot := range rotations {
// 				offset := d1.localPos.sub(d2.localPos.applyRotation(rot))
// 				offsets = append(offsets, offset)
// 			}
// 		}
// 	}

// 	return offsets
// }

func countMatches2(ocean []vec3, s scanner, rot rotation, offset vec3) int {
	count := 0
	for _, detection := range s.detections {
		for _, beacon := range ocean {
			if vec3equal(detection.localPos.applyRotation(rot).sub(offset), beacon) {
				count++
			}
		}
	}

	return count
}

func countMatches(o *ocean, unresolved scanner, rot rotation, offset vec3) int {
	count := 0

	for _, detection := range unresolved.detections {
		realPos := detection.localPos.applyRotation(rot).add(offset)

		if _, found := o.beacons[realPos]; found {
			count++
		}
	}

	return count
}

func hasCommonPoints(scan scanner, o ocean, minCount int) (bool, vec3, rotation) {
	for _, rot := range rotations {
		for beacon := range o.beacons {
			for _, unknownBeacon := range scan.detections {
				bWorldSpace := unknownBeacon.localPos.applyRotation(rot)
				aWorldSpace := beacon
				offset := bWorldSpace.sub(aWorldSpace).invert()
				// offset := unknownBeacon.localPos.applyRotation(rot).sub(resolvedBeacon.localPos.applyRotation(resolved.rotation).add(resolved.position)).invert()

				count := countMatches(&o, scan, rot, offset)

				if count >= minCount {
					// fmt.Println(rot)
					return true, offset, rot
					// return true, offset.applyRotation(resolved.rotation), rot
				}
			}
		}

	}

	return false, vec3{}, rotation{}
	// WORKING
	// var count int

	// for _, rot := range rotations {

	// 	for _, d1 := range resolved.detections {
	// 		for _, d2 := range unresolved.detections {
	// 			offset := d2.localPos.applyRotation(rot).sub(d1.localPos)
	// 			count = 0

	// 			for _, d11 := range resolved.detections {
	// 				for _, d22 := range unresolved.detections {
	// 					if vec3equal(d22.localPos.applyRotation(rot).sub(offset), d11.localPos) {
	// 						count++
	// 					}
	// 				}
	// 			}

	// 			if count >= commonDetectionThreshold {
	// 				fmt.Println(rot)
	// 				return true, offset.invert()
	// 			}
	// 		}
	// 	}

	// }

	// return false, vec3{}
	// END WORKD

	// for _, offset := range []vec3{{68, -1246, -43}} {
	// 	count = 0
	// 	for _, d1 := range resolved.detections {
	// 		for _, d2 := range unresolved.detections {
	// 			for _, rot := range rotations {
	// 				if(d2.localPos.applyRotation(rot).sub(offset))
	// 			}
	// 			v1 := d1.localPos
	// 			v2 := d2.localPos.add(offset)

	// 			if vec3equal(v1, v2) {
	// 				fmt.Println("Success", offset)
	// 				count++
	// 			} else {
	// 				fmt.Println(d1.localPos, "+", offset, "=", d2.localPos)
	// 			}
	// 		}

	// 	}
	// 	if count > 0 {
	// 		fmt.Printf("%v\n", count)
	// 	}

	// 	if count >= commonDetectionThreshold {
	// 		return true, offset
	// 	}
	// }

	// return false, vec3{}, count
}

// func commonCount(resolved, unresolved scanner) (bool, vec3, int) {
// 	// offsets := getOffsets(resolved, unresolved)
// 	var count int
// 	for _, offset := range []vec3{{68, -1246, -43}} {
// 		count = 0
// 		for _, d1 := range resolved.detections {
// 			for _, d2 := range unresolved.detections {
// 				for _, rot := range rotations {
// 					if(d2.localPos.applyRotation(rot).sub(offset))
// 				}
// 				v1 := d1.localPos
// 				v2 := d2.localPos.add(offset)

// 				if vec3equal(v1, v2) {
// 					fmt.Println("Success", offset)
// 					count++
// 				} else {
// 					fmt.Println(d1.localPos, "+", offset, "=", d2.localPos)
// 				}
// 			}

// 		}
// 		if count > 0 {
// 			fmt.Printf("%v\n", count)
// 		}

// 		if count >= commonDetectionThreshold {
// 			return true, offset
// 		}
// 	}

// 	return false, vec3{}, count
// }

func resolveScanner(unresolved *scanner, rot rotation, offset vec3) {
	fmt.Printf("Found scanner %v at offset: %v rotation: %v \n", unresolved.id, offset, rot)

	unresolved.resolved = true
	unresolved.rotation = rot
	unresolved.position = offset
}

func removeFromSlice(slice []scanner, index int) []scanner {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

type ocean struct {
	beacons  map[vec3]bool
	scanners []scanner
}

func (o *ocean) resolveScanner(s scanner, offset vec3, rot rotation) {
	for _, detection := range s.detections {
		o.beacons[detection.localPos.applyRotation(rot).add(offset)] = true
	}

	s.position = offset
	s.rotation = rot

	o.scanners = append(o.scanners, s)
}

func (o *ocean) resolveScanners(scanners []scanner, count int) {

	for i := 0; i < len(scanners); i++ {
		scanner := scanners[i]
		found, offset, rot := hasCommonPoints(scanners[i], *o, count)
		if found {

			o.resolveScanner(scanner, offset, rot)
			resolveScanner(&(scanners[i]), rot, offset)
			scanners = removeFromSlice(scanners, i)

			i = -1
		}
	}

}

// func resolveScanners(scanners []scanner, count int) {
// 	var resolvedScanners int = 0

// 	for resolvedScanners < len(scanners) {
// 		resolvedScanners = 0
// 	r:
// 		for i := range scanners {
// 			s := &scanners[i]
// 			if s.resolved {
// 				resolvedScanners++
// 			} else {
// 				for i2 := range scanners {
// 					ref := &scanners[i2]
// 					if !ref.resolved || ref.id == s.id {
// 						continue
// 					}
// 					found, offset, rot := hasCommonPoints(*ref, *s, count)
// 					if found {
// 						fmt.Printf("Found scanner %v at from scanner %v, offset: %v rotation: %v \n", s.id, ref.id, offset, rot)
// 						s.resolved = true
// 						s.rotation = rot
// 						s.position = ref.position.applyRotation(ref.rotation).add(offset).applyRotation(ref.rotation)
// 						// s.position = ref.position.add(offset.applyRotation(ref.rotation)).applyRotation(ref.rotation).applyRotation(ref.rotation)
// 						break r
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func resolveScanners(unresolvedScanners, resolvedScanners *[]scanner) {
// 	resolveScanner := func(s *scanner, position vec3) {
// 		s.position = position

// 		index := findScannerInSlice(s.id, *unresolvedScanners)
// 		if index == -1 {
// 			panic("didn't find scanner, shouldnt happen")
// 		}

// 		resolvedScanners = append(resolvedScanners, s)

// 		// Move last scanner into the old one's spot and shrink slice by one
// 		unresolvedScanners[index] = unresolvedScanners[len(unresolvedScanners)-1]
// 		unresolvedScanners = unresolvedScanners[:len(unresolvedScanners)-1]
// 	}

// 	resolveScanner(&unresolvedScanners[0], vec3{})

// 	for len(unresolvedScanners) > 0 {
// 	restart:
// 		for _, resolvedScanner := range resolvedScanners {
// 			for _, unresolvedScanner := range unresolvedScanners {
// 				if ok, offset := hasCommonPoints(resolvedScanner, unresolvedScanner); ok {
// 					resolveScanner(&unresolvedScanner, resolvedScanner.position.add(offset))
// 					fmt.Printf("Solved scanner %v \n", offset)
// 					break restart
// 				}

// 			}
// 		}
// 	}
// }

func countUniqueDetections(scanners []scanner) int {
	m := make(map[vec3]bool)

	for _, s := range scanners {
		for _, d := range s.detections {
			m[s.position.add(d.localPos.applyRotation(s.rotation))] = true
		}
	}

	return len(m)
}

func findLargestDistance(scanners []scanner) int {
	maxDist := 0

	for _, s1 := range scanners {
		for _, s2 := range scanners {
			if s1.id == s2.id {
				continue
			}

			dist := manhattan(s1.position, s2.position)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return maxDist
}

func Solve() (int, int) {
	file, err := os.ReadFile("./day19/input.txt")
	if err != nil {
		panic(err)
	}

	_, unresolvedScanners := parseDetection(string(file))

	unresolvedScanners[0].resolved = true

	// resolvedScanners := make([]scanner, 0, len(unresolvedScanners))

	o := ocean{map[vec3]bool{}, []scanner{}}
	o.resolveScanner(unresolvedScanners[0], vec3{}, rotation{0, 1, 2, 1, 1, 1})

	o.resolveScanners(unresolvedScanners[1:], commonDetectionThreshold)

	count := len(o.beacons)
	dist := findLargestDistance(o.scanners)

	fmt.Printf("Count: %v Max Distance: %v\n", count, dist)

	// fmt.Printf("%v\n", resolvedScanners)

	// Parse all
	// While unsolved scanners remain
	// For every un
	// for _, resolved := range scanners {
	// 	if !resolved.rotationResolved {
	// 		continue
	// 	}

	// 	for _, unresolved := range scanners {
	// 		if unresolved.rotationResolved || unresolved.id == resolved.id {
	// 			continue
	// 		}
	// 		// error, e1, e2 := findCommonLinks(resolved, unresolved)

	// 		// if error == nil {
	// 		// 	continue
	// 		// }
	// 	}

	// }

	// find 2 edges in a single scanner with overlap with a known good scanner
	// resolve its rotatation
	// apply to scanner

	return count, dist

}
