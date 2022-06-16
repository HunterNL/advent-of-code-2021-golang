package aoc

import (
	"encoding/json"
	"testing"
)

func TestParse(t *testing.T) {
	t.SkipNow()
	jsonBytes := []byte(`{"1":{"part1": 5,"part2": 15}}`)

	dayMap := make(map[int]solution)

	decodingError := json.Unmarshal(jsonBytes, &dayMap)
	if decodingError != nil {
		t.Log(decodingError)
		t.FailNow()
	}

	if len(dayMap) != 1 {
		t.Error("Expected dayMap to have len() of 1")
	}

	day, found := dayMap[1]
	if !found {
		t.Error("Expected dayMap to have a '1' key")
	}

	if day.Part1 != 5 {
		t.Error("Expected day.part1 to equal 5")
	}

	if day.Part2 != 15 {
		t.Error("Expected day.part2 to equal 15")
	}
}
