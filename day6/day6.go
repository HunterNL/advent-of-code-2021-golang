package day6

import (
	"os"
	"strconv"
	"strings"
)

func tick(oldFish [9]int) [9]int {
	var newFish [9]int

	newFish[0] = oldFish[1]
	newFish[1] = oldFish[2]
	newFish[2] = oldFish[3]
	newFish[3] = oldFish[4]
	newFish[4] = oldFish[5]
	newFish[5] = oldFish[6]
	newFish[6] = oldFish[7] + oldFish[0]
	newFish[7] = oldFish[8]
	newFish[8] = oldFish[0]

	return newFish
}

func CalcLanternFish(startFish []int, interations int) int {
	var fish [9]int

	for _, n := range startFish {
		fish[n]++
	}

	for i := 0; i < interations; i++ {
		fish = tick(fish)
	}

	sum := 0
	for _, fishCount := range fish {
		sum += fishCount
	}

	return sum

}

func Solve() (int, int, error) {
	file, err := os.ReadFile("./day6/input.txt")
	if err != nil {

	}

	strs := strings.Split(string(file), ",")
	nums := make([]int, len(strs))

	for i, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}

	part1 := CalcLanternFish(nums, 80)
	part2 := CalcLanternFish(nums, 256)

	// log.Printf("Fishes part1: %v part2: %v", part1, part2)

	return part1, part2, nil
}
