package day20

import (
	"fmt"
	"os"
	"strings"
)

type vec2 struct {
	x, y int
}

type image map[vec2]bool

type bounds struct {
	xMin, xMax, yMin, yMax int
}

func pixelNeighbours(pixel vec2) [9]vec2 {
	return [9]vec2{
		{pixel.x - 1, pixel.y - 1},
		{pixel.x, pixel.y - 1},
		{pixel.x + 1, pixel.y - 1},
		{pixel.x - 1, pixel.y},
		{pixel.x, pixel.y},
		{pixel.x + 1, pixel.y},
		{pixel.x - 1, pixel.y + 1},
		{pixel.x, pixel.y + 1},
		{pixel.x + 1, pixel.y + 1},
	}
}

func sanitizeString(str string) string {
	return strings.ReplaceAll(str, "\n", "")
}

func parseInput(b []byte) (string, image) {
	alg, imgStr, found := strings.Cut(string(b), "\n\n")

	if !found {
		panic("No newline in input")
	}

	return strings.ReplaceAll(alg, "\n", ""), parseImage(imgStr)
}

func parseImage(str string) image {
	i := make(image)
	lines := strings.Split(str, "\n")
	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				i[vec2{x, y}] = true
			}
		}
	}
	return i
}

func (i *image) String() string {
	builder := strings.Builder{}
	bounds := i.imageBounds()
	for y := bounds.yMin; y <= bounds.yMax; y++ {
		for x := bounds.xMin; x <= bounds.xMax; x++ {
			pixel, found := (*i)[vec2{x, y}]
			if found && pixel {
				builder.WriteRune('#')
			} else {
				builder.WriteRune('.')
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func (b bounds) extendBounds() bounds {
	return bounds{
		xMin: b.xMin - 1,
		xMax: b.xMax + 1,
		yMin: b.yMin - 1,
		yMax: b.yMax + 1,
	}
}

func (b bounds) allPixels() []vec2 {
	size := (b.xMax - b.yMin) * (b.yMax - b.yMin)
	vectors := make([]vec2, 0, size)

	for y := b.yMin; y <= b.yMax; y++ {
		for x := b.xMin; x <= b.xMax; x++ {
			vectors = append(vectors, vec2{x, y})
			// fmt.Printf("Appending vector %v\n", vec2{x, y})
		}
	}

	return vectors
}

func (b bounds) size() (x, y int) {
	return b.xMax - b.xMin, b.yMax - b.yMin
}

func (b bounds) inBounds(v vec2) bool {
	if v.x < b.xMin {
		return false
	}
	if v.x > b.xMax {
		return false
	}
	if v.y < b.yMin {
		return false
	}
	if v.y > b.yMax {
		return false
	}

	return true
}

func (i image) imageBounds() (b bounds) {
	for pos := range i {
		if pos.x < b.xMin {
			b.xMin = pos.x
		}
		if pos.x > b.xMax {
			b.xMax = pos.x
		}
		if pos.y < b.yMin {
			b.yMin = pos.y
		}
		if pos.y > b.yMax {
			b.yMax = pos.y
		}
	}

	return b
}

func boolsToNum(b [9]bool) int {
	out := 0
	for i := 0; i < 9; i++ {
		n := 0

		if b[i] {
			n = 1
		}
		out = out << 1
		out = out + n
	}
	return out
}

func stringToNum(str string) int {
	out := 0
	for i := 0; i < len(str); i++ {
		// for i := len(str) - 1; i >= 0; i-- {
		n := 0

		if str[i] == '#' {
			n = 1
		}
		out = out << 1
		out = out + n
	}
	return out
}

func oobCycle(state bool, alg string) bool {
	if state {
		return alg[511] == '#'
	} else {
		return alg[0] == '#'
	}
}

func enhanceImage(in image, algorithm string, oobState bool) (image, bool) {
	realImageBounds := in.imageBounds()
	allPixels := in.imageBounds().extendBounds().allPixels()
	out := make(image, len(allPixels))

	for _, pixel := range allPixels {
		var bools [9]bool
		for i, vec := range pixelNeighbours(pixel) {
			if realImageBounds.inBounds(vec) {
				bools[i] = in[vec]
			} else {
				bools[i] = oobState
			}

		}
		num := boolsToNum(bools)
		if algorithm[num] == '#' {
			out[pixel] = true
		}
	}

	fmt.Printf("Image enhanced, now sized %v\n", len(out))

	return out, oobCycle(oobState, algorithm)
}

func Solve() {
	file, err := os.ReadFile("./day20/input.txt")
	alg, img := parseInput(file)

	if err != nil {
		panic(err)
	}

	fmt.Println("Alg size:", len(alg))

	oobState := false

	// for i := 0; i < 2; i++ {
	// 	img, oobState = enhanceImage(img, alg, oobState)
	// }

	for i := 0; i < 50; i++ {
		img, oobState = enhanceImage(img, alg, oobState)
	}

	fmt.Printf("Image size: %v\n", len(img))

	// fmt.Printf("Image:\n\n%v\n%v\n\n", img.String(), alg)
}
