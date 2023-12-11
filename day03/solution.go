package day03

import (
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2023/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	answer := common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(),
		Part2: solvePart2(),
	}

	return answer
}

func solvePart1() string {
	log.Debug("Solving part 1.")

	// TODO

	log.Debug("Part 1 solved.")
	return "Not Implemented"
}

func solvePart2() string {
	log.Debug("Solving part 2.")

	// TODO

	log.Debug("Part 2 solved.")
	return "Not Implemented"
}
