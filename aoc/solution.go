package aoc

import (
	"encoding/json"
)

type Solution struct {
	Part1 int
	Part2 int
}

func (s *Solution) UnmarshalJSON(b []byte) error {
	intArray := [2]int{}
	error := json.Unmarshal(b, &intArray)
	if error != nil {
		return error
	}
	s.Part1 = intArray[0]
	s.Part2 = intArray[1]

	return nil
}

func (s Solution) MarshalJSON() ([]byte, error) {
	intArray := [2]int{}

	intArray[0] = s.Part1
	intArray[1] = s.Part2

	return json.Marshal(intArray)
}

func sortDays(dayMap map[int]Solution) []Solution {
	out := []Solution{}

	// Find lowest key, append to out and remove
	for len(dayMap) > 0 {
		lowestKey := 99
		for dayN := range dayMap {
			if dayN < lowestKey {
				lowestKey = dayN
			}
		}

		out = append(out, dayMap[lowestKey])

		delete(dayMap, lowestKey)
	}

	return out
}

func DecodeSolutions(file []byte) (map[int]Solution, error) {
	dayMap := make(map[int]Solution)

	decodingError := json.Unmarshal(file, &dayMap)
	if decodingError != nil {
		return nil, decodingError
	}

	return dayMap, nil
}

func EncodeSolutionsSlice(solutions map[int]Solution) ([]byte, error) {
	return json.MarshalIndent(solutions, "", "\t")
}
