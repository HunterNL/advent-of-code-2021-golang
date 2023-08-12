package day20

import (
	"log"
	"os"
	"strings"
	"testing"
)

// var test_alg = "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"

func TestParse(t *testing.T) {
	file, err := os.ReadFile("./input_test.txt")

	if err != nil {
		log.Print(string(file))
		panic(err)
	}

	alg, img := parseInput(file, 5, 0)

	if img.iterationToString(0) != `#..#.
#....
##..#
..#..
..###
` {
		t.Errorf("Parsing image failed, %v\n\n", img)
	}

	if len(alg) != 512 {
		t.Errorf("Algorithm seed size not 512 but %v\n", len(alg))
	}
}

func TestStringToNum(t *testing.T) {
	expected := 34
	result := stringToNum("...#...#.")

	if result != expected {
		t.Errorf("Expected string to result in %v instead of %v\n", expected, result)
	}
}
func TestBoolsToNum(t *testing.T) {
	// expected := 34
	// result := boolsToNum([9]bool{false, false, false, true, false, false, false, true, false})

	tests := []struct {
		input    [9]bool
		expected int
	}{
		{
			input:    [9]bool{false, false, false, false, false, false, false, false, false},
			expected: 0,
		},
		{
			input:    [9]bool{true, true, true, true, true, true, true, true, true},
			expected: 511,
		},
		{
			input:    [9]bool{false, false, false, true, false, false, false, true, false},
			expected: 34,
		},
		{
			input:    [9]bool{false, false, false, false, false, false, false, false, true},
			expected: 1,
		},
		{
			input:    [9]bool{true, false, false, false, false, false, false, false, false},
			expected: 256,
		},
	}

	for _, test := range tests {
		if test.expected != boolsToNum(test.input) {
			t.Errorf("Expected %v instead of %v\n", test.expected, boolsToNum(test.input))
		}
	}

	// if result != expected {
	// 	t.Errorf("Expected string to result in %v instead of %v\n", expected, result)
	// }
}

// func TestImageBounds(t *testing.T) {
// 	file, err := os.ReadFile("./input_test.txt")

// 	if err != nil {
// 		log.Printf(string(file))
// 		panic(err)
// 	}

// 	_, img := parseInput(file)

// 	sizeX, sizeY := img.imageBounds().size()

// 	if sizeX != 5 {
// 		t.Error("Expected sizeof 5, not", sizeX)
// 	}
// 	if sizeY != 5 {
// 		t.Error("Expected sizeof 5, not", sizeY)
// 	}

// }

func TestEnhance(t *testing.T) {
	file, err := os.ReadFile("./input_test.txt")

	if err != nil {
		log.Print(string(file))
		panic(err)
	}

	expectedImage := newImage(7)

	expectedImage.parseInto(sanitizeString(
		""+
			".##.##."+
			"#..#.#."+
			"##.#..#"+
			"####..#"+
			".#..##."+
			"..##..#"+
			"...#.#.",
	), 7, 0)

	alg, img := parseInput(file, 5, 1)
	enhancer := newEnhancer(img, alg, 1)

	enhancer.enhanceImage()

	var numBits [9]bool
	for boolIndex, neigbourIndex := range pixelNeighbours(3+3*7, img.size) {
		neighbourState := img.pixels[neigbourIndex]
		numBits[boolIndex] = ((neighbourState & 1) > 0)
	}
	num := boolsToNum(numBits)
	if num != 34 {
		t.Errorf("Expected center num to resolve to 34, not %v\n", num)
	}

	enhancedImageStr := img.iterationToString(1)
	expectedImageStr := expectedImage.iterationToString(0)

	if enhancedImageStr != expectedImageStr {
		t.Errorf("Expected \n%v\n\nbut got\n\n%v\n\n", expectedImageStr, enhancedImageStr)
	}
}

func TestOOB(t *testing.T) {
	expectedImage := newImage(5)

	var oob_alg = strings.Repeat("#.", 256) //First character is #, last character is "."

	expectedImage.parseInto(strings.Repeat("#", 5*5), 5, 0)

	img := newImage(5)
	img.parseInto(strings.Repeat(".", 3*3), 3, 1)

	enhancer := newEnhancer(img, oob_alg, 1)
	enhancer.enhanceImage()

	enhancedImageStr := img.iterationToString(1)
	expectedImageStr := expectedImage.iterationToString(0)

	if enhancedImageStr != expectedImageStr {
		t.Errorf("Expected \n%v\n\nbut got\n\n%v\n\n", expectedImageStr, enhancedImageStr)
	}
}

func TestOOBInverse(t *testing.T) {
	expectedImage := newImage(5)

	var oob_alg = strings.Repeat("#.", 256) //First character is #, last character is "."

	expectedImage.parseInto(strings.Repeat(".", 5*5), 5, 0)

	img := newImage(5)
	img.parseInto(strings.Repeat("#", 3*3), 3, 1)

	enhancer := newEnhancer(img, oob_alg, 1)
	enhancer.oobState = true
	enhancer.enhanceImage()

	enhancedImageStr := img.iterationToString(1)
	expectedImageStr := expectedImage.iterationToString(0)

	if enhancedImageStr != expectedImageStr {
		t.Errorf("Expected \n%v\n\nbut got\n\n%v\n\n", expectedImageStr, enhancedImageStr)
	}
}

func TestEnhanceCount(t *testing.T) {
	file, err := os.ReadFile("./input_test.txt")

	if err != nil {
		t.Log(string(file))
		panic(err)
	}

	alg, img := parseInput(file, 5, 2)

	enhancer := newEnhancer(img, alg, 2)

	for i := 0; i < 2; i++ {
		enhancer.enhanceImage()
		t.Log(img.countActivePixels(uint8(i)))
	}

	log.Println(img)

	// t.Fail()

	if img.countActivePixels(2) != 35 {
		t.Errorf("Expected 35 pixels lit instead of %v\n", img.countActivePixels(2))
	}

}
func TestBulkEnhanceCount(t *testing.T) {
	file, err := os.ReadFile("./input_test.txt")

	if err != nil {
		log.Print(string(file))
		panic(err)
	}

	alg, img := parseInput(file, 5, 50)
	enhancer := newEnhancer(img, alg, 50)

	enhanceCount := 50

	for i := 0; i < enhanceCount; i++ {
		enhancer.enhanceImage()
	}

	log.Println(img)

	if img.countActivePixels(uint8(enhancer.currentIteration)) != 3351 {
		t.Errorf("Expected 3351 pixels lit instead of %v\n", img.countActivePixels(uint8(enhancer.currentIteration)))
	}

}
