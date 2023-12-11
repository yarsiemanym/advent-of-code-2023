package day01

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2023/common"
)

func Test_Solve_Test1(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2023,
		Day:       1,
		InputFile: "test1.txt",
	}

	expectedCalibrationValue := "142"

	answer := Solve(puzzle)

	if answer.Part1 != expectedCalibrationValue {
		t.Errorf("Expected: %v, Actual:%v", expectedCalibrationValue, answer.Part1)
	}
}

func Test_Solve_Test2(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2023,
		Day:       1,
		InputFile: "test2.txt",
	}

	expectedCalibrationValue := "281"

	answer := Solve(puzzle)

	if answer.Part2 != expectedCalibrationValue {
		t.Errorf("Expected: %v, Actual:%v", expectedCalibrationValue, answer.Part2)
	}
}
