package day20

import (
	"log"
	"os"
	"testing"
)

var test_alg = "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"

func TestParse(t *testing.T) {
	file, err := os.ReadFile("./input_test.txt")

	if err != nil {
		log.Printf(string(file))
		panic(err)
	}

	alg, img := parseInput(file)

	if img.String() != `#..#.
#....
##..#
..#..
..###
` {
		t.Errorf("Parsing image failed, %v\n\n", img.String())
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

func TestImageBounds(t *testing.T) {
	file, err := os.ReadFile("./input_test.txt")

	if err != nil {
		log.Printf(string(file))
		panic(err)
	}

	_, img := parseInput(file)

	sizeX, sizeY := img.imageBounds().size()

	if sizeX != 5 {
		t.Error("Expected sizeof 5, not", sizeX)
	}
	if sizeY != 5 {
		t.Error("Expected sizeof 5, not", sizeY)
	}

}

func TestEnhance(t *testing.T) {
	file, err := os.ReadFile("./input_test.txt")

	if err != nil {
		log.Printf(string(file))
		panic(err)
	}

	// 	expectedImage := parseImage(sanitizeString(`
	// .##.##.
	// #..#.#.
	// ##.#..#
	// ####..#
	// .#..##.
	// ..##..#
	// ...#.#.`))

	alg, img := parseInput(file)

	log.Printf("\n%v\n", img.String())

	var bools [9]bool
	for i, vec := range pixelNeighbours(vec2{2, 2}) {
		bools[i] = img[vec]
	}
	num := boolsToNum(bools)
	if num != 34 {
		t.Errorf("Expected center num to resolve to 34, not %v\n", num)
	}

	img2, _ := enhanceImage(img, alg, false)

	t.Logf("\n%v\n", img2.String())
	t.Fail()
}

func TestEnhanceCount(t *testing.T) {
	file, err := os.ReadFile("./input_test.txt")

	if err != nil {
		log.Printf(string(file))
		panic(err)
	}

	alg, img := parseInput(file)

	for i := 0; i < 2; i++ {
		img, _ = enhanceImage(img, alg, false)
	}

	log.Printf(img.String())

	// t.Fail()

	if len(img) != 35 {
		t.Errorf("Expected 35 pixels lit instead of %v\n", len(img))
	}

}
func TestBulkEnhanceCount(t *testing.T) {
	file, err := os.ReadFile("./input_test.txt")

	if err != nil {
		log.Printf(string(file))
		panic(err)
	}

	alg, img := parseInput(file)

	enhanceCount := 50

	for i := 0; i < enhanceCount; i++ {
		img, _ = enhanceImage(img, alg, false)
	}

	log.Printf(img.String())

	// t.Fail()

	if len(img) != 3351 {
		t.Errorf("Expected 3351 pixels lit instead of %v\n", len(img))
	}

}
