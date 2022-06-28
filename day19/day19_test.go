package day19

import (
	"log"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	if err != nil {
		panic(err)
	}
	_, scanners := parseDetection(string(bytes))

	if len(scanners) != 5 {
		t.Errorf("Expected %d scanners, got %d", 5, len(scanners))
	}
}

func TestRotationoffset(t *testing.T) {
	bytes, err := os.ReadFile("./test2_input.txt")
	_, scanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	o := oceanFromStartingScanner(scanners[0])

	for i := 1; i <= 4; i++ {
		found, offset, _ := hasCommonPoints(scanners[1], o, 6)
		if !found {
			t.Errorf("Expected scanner %v to be found", i)
		}
		if (!vec3equal(offset, vec3{})) {
			t.Errorf("Expected scanner %v at offset %v, not at %v", i, vec3{}, offset)
		}
	}
}

func TestRotationCount(t *testing.T) {
	t.SkipNow()
	if len(rotations) != 24 {
		for _, r := range rotations {
			t.Log(r)
		}
		t.Fail()
	}

}

func TestOverLapCount(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	_, scanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	o := oceanFromStartingScanner(scanners[0])

	count := countMatches(&o, scanners[1], rotation{0, 1, 2, -1, 1, -1}, vec3{68, -1246, -43})

	// t.Fail()

	if count != 12 {
		t.Errorf("Expected count to be 12 instead of %v", count)
	}

}

var expectedPositions = map[int]vec3{
	0: {0, 0, 0},
	1: {68, -1246, -43},
	2: {1105, -1205, 1229},
	3: {-92, -2380, -20},
	4: {-20, -1133, 1061},
}

func TestAbsolutePositions(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	_, scanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}
	o := oceanFromStartingScanner(scanners[0])

	o.resolveScanners(scanners[1:], 12)

	for _, s := range o.scanners {
		expectedPos := expectedPositions[s.id]
		if !vec3equal(expectedPos, s.position) {
			t.Errorf("Expected scanner %v at %v instead of %v", s.id, expectedPos, s.position)
		}
	}

	if len(expectedPositions) != len(o.scanners) {
		t.Errorf("Expected %v scanners but only found %v", len(expectedPositions), len(o.scanners))
	}
}

func TestCount(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	_, unresolvedScanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	o := oceanFromStartingScanner(unresolvedScanners[0])

	o.resolveScanners(unresolvedScanners[1:], commonDetectionThreshold)

	for _, s := range unresolvedScanners {
		log.Println(s.id, s.resolved, s.position)
	}

	count := len(o.beacons)
	if count != 79 {
		t.Errorf("Expected 79 beacons but got %v\n", count)
	}

}

func TestManhatten(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	_, unresolvedScanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	o := oceanFromStartingScanner(unresolvedScanners[0])

	o.resolveScanners(unresolvedScanners[1:], commonDetectionThreshold)

	maxDist := findLargestDistance(o.scanners)

	if maxDist != 3621 {
		t.Errorf("Expected distance to be 3621 instead of %v\n", maxDist)
	}
}
