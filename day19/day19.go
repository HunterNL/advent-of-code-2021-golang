package day19

import (
	"aoc2021/intmath"
	"log"
	"os"
	"sort"
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

func (v vec3) applyRotation(r rotation) (v2 vec3) {
	v2.x = v.index(r.oX) * r.sX
	v2.y = v.index(r.oY) * r.sY
	v2.z = v.index(r.oZ) * r.sZ

	return v2
}

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
	}
	return rotations
}

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

func vec3equal(a, b vec3) bool {
	return a.x == b.x && a.y == b.y && a.z == b.z
}

func manhattan(a, b vec3) int {
	return intmath.Distance(a.x, b.x) + intmath.Distance(a.y, b.y) + intmath.Distance(a.z, b.z)
}

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
	links      map[[3]int]link
}

func parseDetection(str string) ([]detection, []scanner) {
	scanstring := strings.Split(str, "\n\n")
	detections := make([]detection, 0, len(scanstring)*30)
	scanners := make([]scanner, 0, len(scanstring))

	for i, s := range scanstring {
		lines := strings.Split(s, "\n")[1:]
		scan := scanner{id: i, links: make(map[[3]int]link)}

		for _, l := range lines {
			position := parseVec3(l)
			detection := detection{localPos: position, scannerId: i}
			detections = append(detections, detection)
			scan.detections = append(scan.detections, detection)
			scan.rotation = rotation{0, 1, 2, 1, 1, 1}
		}

		for i, leftDetection := range scan.detections {
			for i2 := len(scan.detections) - 1; i2 > i; i2-- {
				rightDetection := scan.detections[i2]
				diff := rightDetection.localPos.sub(leftDetection.localPos)

				looseInts := [3]int{intmath.Abs(diff.x), intmath.Abs(diff.y), intmath.Abs(diff.z)}
				intslice := looseInts[:]

				sort.Ints(intslice)

				link := link{
					a:         intslice[0],
					b:         intslice[1],
					c:         intslice[2],
					movement:  diff,
					detection: [2]detection{leftDetection, rightDetection},
				}

				scan.links[looseInts] = link
			}
		}

		scanners = append(scanners, scan)
	}

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

func countLinks(resolved, unresolved *scanner) int {
	count := 0
	for k := range resolved.links {
		_, found := unresolved.links[k]
		if found {
			count++
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
	for _, resolvedScanner := range o.scanners {
		if !resolvedScanner.resolved {
			continue
		}

		if countLinks(&resolvedScanner, &scan) < minCount {
			continue
		}

		rot := findRotation(&resolvedScanner, &scan)

		// for _, rot := range rotations {
		for beacon := range o.beacons {
			for _, unknownBeacon := range scan.detections {
				bWorldSpace := unknownBeacon.localPos.applyRotation(rot)
				aWorldSpace := beacon
				offset := aWorldSpace.sub(bWorldSpace)

				count := countMatches(&o, scan, rot, offset)

				if count >= minCount {
					return true, offset, rot
				}
			}
		}
		// }
	}

	return false, vec3{}, rotation{}
}

func findRotation(resolved, unresolved *scanner) (out rotation) {
	for k1, v1 := range resolved.links {
		l2, found := unresolved.links[k1]
		if !found {
			continue
		}

		for _, rot := range rotations {
			if l2.movement.applyRotation(rot) == v1.movement.applyRotation(resolved.rotation) {
				return rot
			}
		}

	}

	panic("rotation not found")
}

func removeFromSlice(slice []scanner, index int) []scanner {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

type link struct {
	a, b, c   int
	movement  vec3
	detection [2]detection
}

type ocean struct {
	beacons  map[vec3]bool
	scanners []scanner
	links    map[[3]int]link
}

func oceanFromStartingScanner(s scanner) ocean {
	o := ocean{map[vec3]bool{}, []scanner{}, map[[3]int]link{}}
	o.resolveScanner(s, vec3{}, rotation{0, 1, 2, 1, 1, 1})
	return o
}

func (o *ocean) resolveScanner(s scanner, offset vec3, rot rotation) {
	for _, detection := range s.detections {
		o.beacons[detection.localPos.applyRotation(rot).add(offset)] = true
	}

	log.Printf("Found scanner %v at offset: %v rotation: %v \n", s.id, offset, rot)

	s.resolved = true
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
			scanners = removeFromSlice(scanners, i)

			i = -1 // Gets incremented to 0 next loop
		}
	}

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

func Solve() (int, int, error) {
	file, err := os.ReadFile("./day19/input.txt")
	if err != nil {
		panic(err)
	}

	_, unresolvedScanners := parseDetection(string(file))

	o := oceanFromStartingScanner(unresolvedScanners[0])

	o.resolveScanners(unresolvedScanners[1:], commonDetectionThreshold)

	count := len(o.beacons)
	dist := findLargestDistance(o.scanners)

	log.Printf("Count: %v Max Distance: %v\n", count, dist)

	return count, dist, nil
}
