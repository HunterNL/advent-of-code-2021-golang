package aoc

import (
	"encoding/json"
)

type solution struct {
	dayNumber int
	Part1     int
	Part2     int
}

func sortDays(dayMap map[int]solution) []solution {
	out := []solution{}

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

func ParseSolutions(file []byte) (map[int]solution, error) {
	dayMap := make(map[int]solution)

	decodingError := json.Unmarshal(file, &dayMap)
	if decodingError != nil {
		return nil, decodingError
	}

	return dayMap, nil
}
