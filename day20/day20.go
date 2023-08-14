package day20

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const PIXEL_LIT = '#'
const PIXEL_DIM = '.'

type image struct {
	pixels []uint64
	size   int
}

type bounds struct {
	min, max int
}

type imageEnhancer struct {
	img              image
	currentIteration int
	maxIterations    int
	imageBounds      bounds
	algorithm        string
	oobState         bool
}

func newEnhancer(img image, alg string, currentOffset int) *imageEnhancer {
	currentImageSize := img.size - (currentOffset * 2)
	return &imageEnhancer{
		img:              img,
		currentIteration: 0,
		maxIterations:    currentOffset,
		algorithm:        alg,
		oobState:         false,
		imageBounds: bounds{
			min: currentOffset,
			max: currentOffset + currentImageSize,
		},
	}
}

func (enhancer *imageEnhancer) enhanceImage() error {
	if enhancer.currentIteration == enhancer.maxIterations {
		return errors.New("max iterations reached")
	}

	enhancer.oobState = enhanceImage(enhancer.img, enhancer.algorithm, enhancer.oobState, uint8(enhancer.currentIteration), enhancer.imageBounds)
	enhancer.imageBounds = enhancer.imageBounds.extendBounds()
	enhancer.currentIteration++

	return nil
}

func newImage(size int) image {
	return image{
		pixels: make([]uint64, size*size),
		size:   size,
	}
}

func (img image) parseInto(input string, inputSize int, maxIterations int) error {
	for i, r := range input {
		if !(r == PIXEL_DIM || r == PIXEL_LIT) {
			return fmt.Errorf("input string contains a %q, only %q or %q are allowed", r, PIXEL_DIM, PIXEL_LIT)
		}

		x := i % inputSize
		y := i / inputSize

		if r == '#' {
			img.pixels[x+maxIterations+(y+maxIterations)*img.size] = 1
		}
	}

	return nil
}

func (i image) countActivePixels(iteration uint8) int {
	var iteration_mask uint64 = 1 << iteration
	out := 0

	for _, pixel := range i.pixels {
		if (pixel & iteration_mask) > 0 {
			out++
		}
	}

	return out
}

func pixelNeighbours(pixel, lineWidth int) [9]int {
	return [9]int{
		pixel - lineWidth - 1,
		pixel - lineWidth,
		pixel - lineWidth + 1,
		pixel - 1,
		pixel,
		pixel + 1,
		pixel + lineWidth - 1,
		pixel + lineWidth,
		pixel + lineWidth + 1,
	}
}

func sanitizeString(str string) string {
	return strings.ReplaceAll(str, "\n", "")
}

func parseInput(b []byte, startSize int, maxIterations int) (string, image) {
	alg, imgStr, found := strings.Cut(string(b), "\n\n")

	if !found {
		panic("No newline in input")
	}

	var final_size = startSize + maxIterations*2
	// var offset = maxIterations

	image := newImage(final_size)

	err := image.parseInto(strings.ReplaceAll(imgStr, "\n", ""), startSize, maxIterations)
	if err != nil {
		panic(err)
	}
	// parseImage(image, imgStr, offset)

	return strings.ReplaceAll(alg, "\n", ""), image
}

// func parseImage(i image, str string, padding int) {
// 	lines := strings.Split(str, "\n")
// 	for y, line := range lines {
// 		for x, r := range line {
// 			if r == '#' {
// 				i.pixels[x+padding+(y+padding)*i.size] = 1
// 			}
// 		}
// 	}
// }

// func (i *image) String() string {
// builder := strings.Builder{}
// bounds := i.imageBounds()
// for y := bounds.yMin; y <= bounds.yMax; y++ {
// for x := bounds.xMin; x <= bounds.xMax; x++ {
// pixel, found := (*i)[x+y*FINAL_SIZE]
// if found && pixel {
// builder.WriteRune('#')
// } else {
// builder.WriteRune('.')
// }
// }
// builder.WriteRune('\n')
// }
// return builder.String()
// }

func (i image) iterationToString(interation int) string {
	var iteration_mask uint64 = 1 << interation

	str := strings.Repeat(".", i.size*i.size)
	runes := []rune(str)
	l := len(runes)

	_ = l
	// for index := range i.pixels {
	// 	runes[index] = '#'
	// }

	for k, v := range i.pixels {
		if (v & iteration_mask) > 0 {
			runes[k] = '#'
		}
	}

	sb := strings.Builder{}
	for k, v := range runes {
		sb.WriteRune(v)

		if k%i.size == i.size-1 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()

	// builder := strings.Builder{}
	// // bounds := i.imageBounds()
	// for y := bounds.yMin; y <= bounds.yMax; y++ {
	// 	for x := bounds.xMin; x <= bounds.xMax; x++ {
	// 		pixel, found := (*i)[x+y*FINAL_SIZE]
	// 		if found && pixel {
	// 			builder.WriteRune('#')
	// 		} else {
	// 			builder.WriteRune('.')
	// 		}
	// 	}
	// 	builder.WriteRune('\n')
	// }
	// return builder.String()
}

func (b bounds) extendBounds() bounds {
	return bounds{
		b.min - 1,
		b.max + 1,
	}
}

func (b bounds) inBounds(x, y int) bool {
	if x < b.min {
		return false
	}
	if x >= b.max {
		return false
	}
	if y < b.min {
		return false
	}
	if y >= b.max {
		return false
	}

	return true
}

// func (i image) imageBounds() (b bounds) {
// 	for pos := range i {
// 		if pos.x < b.xMin {
// 			b.xMin = pos.x
// 		}
// 		if pos.x > b.xMax {
// 			b.xMax = pos.x
// 		}
// 		if pos.y < b.yMin {
// 			b.yMin = pos.y
// 		}
// 		if pos.y > b.yMax {
// 			b.yMax = pos.y
// 		}
// 	}

// 	return b
// }

func boolsToNum(b [9]bool) (out uint16) {
	for i := 0; i < 9; i++ {
		var n uint16 = 0

		if b[i] {
			n = 1
		}
		out = out << 1
		out = out | n
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

func enhanceImage(in image, algorithm string, oobState bool, iteration uint8, b bounds) bool {
	var currentPassMask uint64 = 1 << iteration
	var nextPassMask uint64 = 1 << (iteration + 1)
	// litCount := 0

	// testSb := strings.Builder{}

	setBounds := b.extendBounds()
	checkBounds := b

	for y := setBounds.min; y < setBounds.max; y++ {
		for x := setBounds.min; x < setBounds.max; x++ {
			index := x + y*in.size

			newPixelState := enhanceSinglePixel(x, y, algorithm, in, oobState, currentPassMask, checkBounds)
			if newPixelState {
				// litCount++
				in.pixels[index] = (in.pixels[index] | nextPassMask)
			}
		}

		// testSb.WriteRune('\n')
	}

	// log.Println("Pixel count", litCount)
	// log.Printf("\n%v\n", testSb.String())

	return oobCycle(oobState, algorithm)
}

func enhanceSinglePixel(x, y int, algorithm string, img image, oobState bool, bitMask uint64, b bounds) bool {
	var numBits [9]bool
	var boolState bool
	for offsetY := -1; offsetY <= 1; offsetY++ {
		sampleY := y + offsetY
		for offsetX := -1; offsetX <= 1; offsetX++ {
			samplyX := x + offsetX

			if !b.inBounds(samplyX, sampleY) {
				boolState = oobState
			} else {
				neighbourState := img.pixels[samplyX+sampleY*img.size]
				boolState = ((neighbourState & bitMask) > 0)
			}

			numBits[(offsetX+1)+((offsetY+1)*3)] = boolState
		}
	}
	algIndex := boolsToNum(numBits)
	return algorithm[algIndex] == '#'
}

func Solve() (int, int, error) {
	file, err := os.ReadFile("./day20/input.txt")
	if err != nil {
		return -1, -1, err
	}

	alg, img := parseInput(file, 100, 50)
	en := newEnhancer(img, alg, 50)

	for i := 0; i < 2; i++ {
		en.enhanceImage()
		log.Println(img.countActivePixels(uint8(en.currentIteration)))
	}

	part1Len := img.countActivePixels(uint8(en.currentIteration))

	for i := 0; i < 48; i++ {
		en.enhanceImage()
	}

	part2Len := img.countActivePixels(uint8(en.currentIteration))

	return part1Len, part2Len, nil
}
