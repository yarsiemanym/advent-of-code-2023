package day03

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2023/common"
)

func Test_Solve_Test1(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2023,
		Day:       3,
		InputFile: "test1.txt",
	}

	expectedPart1Value := "Not Implemented"
	expectedPart2Value := "Not Implemented"

	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1Value {
		t.Errorf("Expected: %v, Actual:%v", expectedPart1Value, answer.Part1)
	}

	if answer.Part2 != expectedPart2Value {
		t.Errorf("Expected: %v, Actual:%v", expectedPart2Value, answer.Part2)
	}
}
