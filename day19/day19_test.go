package day19

import (
	"fmt"
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

// func TestOffsetGenerator(t *testing.T) {
// 	bytes, err := os.ReadFile("./test1_input.txt")
// 	_, unresolvedScanners := parseDetection(string(bytes))

// 	if err != nil {
// 		panic(err)
// 	}

// 	offsets := getOffsets(unresolvedScanners[0], unresolvedScanners[1])

// 	found := false
// 	for _, o := range offsets {
// 		if (vec3equal(o, vec3{68, -1246, -43})) {
// 			found = true
// 		}
// 	}

// 	if !found {
// 		t.Error("Expected to find vec3{68, -1246, -43} in possible offsets")
// 	}
// }

func TestRotationoffset(t *testing.T) {
	bytes, err := os.ReadFile("./test2_input.txt")
	_, scanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	scanners[0].resolved = true

	for i := 1; i <= 4; i++ {
		found, offset, _ := hasCommonPoints(scanners[0], scanners[i], 6)
		if !found {
			t.Errorf("Expected scanner %v to be found", i)
		}
		if (!vec3equal(offset, vec3{})) {
			t.Errorf("Expected scanner %v at offset %v, not at %v", i, vec3{}, offset)
		}
	}

	// resolveScanners(scanners, 5)

	// print(scanners)
}

func TestOffset(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	_, scanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	found0, offset0, _ := hasCommonPoints(scanners[0], scanners[1], commonDetectionThreshold)
	found1, offset1, _ := hasCommonPoints(scanners[1], scanners[1], commonDetectionThreshold)
	found2, offset2, _ := hasCommonPoints(scanners[1], scanners[2], commonDetectionThreshold)
	found3, offset3, _ := hasCommonPoints(scanners[1], scanners[3], commonDetectionThreshold)
	found4, offset4, _ := hasCommonPoints(scanners[1], scanners[4], commonDetectionThreshold)

	fmt.Println(found1, offset1, found2, offset2, found3, offset3, found4, offset4)

	found := found0
	offset := offset0

	// t.Log(offset)

	if !found {
		t.Fatal("Expected hasCommonDetections to succeed")
	}

	expected := vec3{68, -1246, -43}

	if !vec3equal(offset, expected) {
		t.Fatalf("Expected offset to be %v but got %v", expected, offset)
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

	count := countMatches(scanners[0], scanners[1], rotation{0, 1, 2, -1, 1, -1}, vec3{68, -1246, -43})

	// t.Fail()

	if count != 12 {
		t.Errorf("Expected count to be 12 instead of %v", count)
	}

}

var expectedPositions = []vec3{
	{0, 0, 0},
	{68, -1246, -43},
	{1105, -1205, 1229},
	{-92, -2380, -20},
	{-20, -1133, 1061},
}

func TestOverLap(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	_, scanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	found, offset, rot := hasCommonPoints(scanners[0], scanners[1], 12)
	// scanners[1].rotation = rot
	// scanners[1].position = offset
	// scanners[1].resolved = true

	resolveScanner(scanners[0], &scanners[1], rot, offset)

	// t.Fail()

	if !found {
		t.Errorf("Expected scanner 1 to be found")
		t.FailNow()
	}

	expected := expectedPositions[1]

	if !vec3equal(offset, expected) {
		t.Fatalf("Expected offset to be %v but got %v", expected, offset)
	}

	found, offset, rot = hasCommonPoints(scanners[1], scanners[3], 12)

	if !found {
		t.Log("Expected scanner 3 to be found")
		t.FailNow()
	}

	t.Log(rot)
	t.Log(offset)
	// offset = offset.applyRotation(rot)
	t.Log(offset)

	expected = expectedPositions[3].sub(expectedPositions[1])

	if !vec3equal(offset, expected) {
		t.Fatalf("Expected %v at %v but got %v", 3, expected, offset)
	}

}

func TestAbsolutePositions(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	_, scanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	scanners[0].resolved = true

	resolveScanners(scanners, 12)

	for i, ep := range expectedPositions {
		pos := scanners[i].position //.applyRotation(scanners[i].rotation)

		if !vec3equal(pos, ep) {
			t.Errorf("Expected scanner %v at %v instead of %v", i, ep, pos)
		}
	}
}
func TestRelativePositions(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	_, scanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	scanners[0].rotation = rotation{0, 1, 2, 1, 1, 1}
	scanners[0].resolved = true

	resolveScanners(scanners, 12)

	_, offset, rot := hasCommonPoints(scanners[0], scanners[1], 12)

	v1 := expectedPositions[1]

	if !vec3equal(offset, v1) {
		t.Errorf("Expected offset %v to be %v instead of %v, rot %v", 0, v1, offset, rot)
	}

	_, offset, rot = hasCommonPoints(scanners[1], scanners[3], 12)

	offset = offset.applyRotation(scanners[1].rotation)

	v2 := expectedPositions[3].sub(expectedPositions[1])
	if !vec3equal(offset, v2) {
		t.Errorf("Expected offset %v to be %v instead of %v, diff of %v, rot of %v", "1 to 3", v2, offset, v2.sub(offset), rot)
	}

}

func TestCount(t *testing.T) {
	bytes, err := os.ReadFile("./test1_input.txt")
	_, unresolvedScanners := parseDetection(string(bytes))

	if err != nil {
		panic(err)
	}

	unresolvedScanners[0].rotation = rotation{0, 1, 2, 1, 1, 1}
	unresolvedScanners[0].resolved = true

	resolveScanners(unresolvedScanners, commonDetectionThreshold)

	for _, s := range unresolvedScanners {
		fmt.Println(s.id, s.resolved, s.position)
	}

	count := countUniqueDetections(unresolvedScanners)
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

	unresolvedScanners[0].rotation = rotation{0, 1, 2, 1, 1, 1}
	unresolvedScanners[0].resolved = true

	resolveScanners(unresolvedScanners, commonDetectionThreshold)

	maxDist := findLargestDistance(unresolvedScanners)

	if maxDist != 3621 {
		t.Errorf("Expected distance to be 3621 instead of %v\n", maxDist)
	}
}

// func TestSolve(t *testing.T) {
// 	bytes, err := os.ReadFile("./test1_input.txt")
// 	_, scanners := parseDetection(string(bytes))

// 	if err != nil {
// 		panic(err)
// 	}

// 	ocean := make([]vec3, 0, 0)

// 	for _, v := range scanners[0].detections {
// 		ocean = append(ocean, v.localPos)
// 	}

// }

// func TestRotationResolve(t *testing.T) {
// 	bytes, err := os.ReadFile("./test2_input.txt")
// 	_, unresolvedScanners := parseDetection(string(bytes))
// 	if err != nil {
// 		panic(err)
// 	}

// 	found, offset := hasCommonPoints(unresolvedScanners[0], unresolvedScanners[1])
// 	if count != 6 {
// 		t.Error("Expected 6 matches")
// 	}
// }
