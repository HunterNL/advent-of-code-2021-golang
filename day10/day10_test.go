package day10

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	validInputs := []string{"()", "([])", "{()()()}", "<([{}])>", "[<>({}){}[([])<>]]", "(((((((((())))))))))"}

	corrupedInputs := map[string]rune{
		"{([(<{}[<>[]}>{[]{[(<()>": '}',
		"[[<[([]))<([[{}[[()]]]":   ')',
		"[{[{({}]{}}([{[{{{}}([]":  ']',
		"[<(<(<(<{}))><([]([]()":   ')',
		"<{([([[(<>()){}]>(<<{{":   '>',
	}

	incompleteInputs := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"(((({<>}<{<{<>}{[]{[]{}",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}

	for _, s := range validInputs {
		err := parseLine(s)

		if err != nil {
			t.Errorf("Expected %v to succeed, failed instead, error: %v", s, err)
		}
	}

	for line, expectd := range corrupedInputs {
		err := parseLine(line)

		if err.getStatus() != Corrupt {
			t.Errorf("Expected %v to be corrupt. Error: %v", line, err)
		}

		if err.getUnexpectedRune() != expectd {
			t.Errorf("Expected %v to not expect %v, got %v instead", line, expectd, err.getUnexpectedRune())
		}
	}

	for _, line := range incompleteInputs {
		err := parseLine(line)

		if err.getStatus() != Incomplete {
			t.Errorf("Extected line to be incomplete, got %v instead", err.getStatus())
		}
	}

}

func TestCompleteLine(t *testing.T) {
	tests := map[string][]rune{
		"[({(<(())[]>[[{[]{<()<>>": []rune("}}]])})]"),
		"[(()[<>])]({[<{<<[]>>(":   []rune(")}>]})"),
		"(((({<>}<{<{<>}{[]{[]{}":  []rune("}}>}>))))"),
		"{<[[]]>}<{[{[{[]{()[[[]":  []rune("]]}}]}]}>"),
		"<{([{{}}[<[[[<>{}]]]>[]]": []rune("])}>"),
	}

	for line, expected := range tests {
		_, completion := completeLine(line)
		if !reflect.DeepEqual(completion, expected) {
			t.Errorf("Line %v's autocomplete did not match, expected %v but got %v", line, expected, completion)
		}
	}
}

func TestCompletioNScore(t *testing.T) {
	score := scoreCompletion([]rune("])}>"))
	if score != 294 {
		t.Errorf("Expected a score of 294, got %v", score)
	}
}
