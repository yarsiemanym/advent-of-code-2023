package day03

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2023/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	schematic := parseSchematic(common.ReadFile(puzzle.InputFile))

	answer := common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(schematic),
		Part2: solvePart2(schematic),
	}

	return answer
}

func solvePart1(schematic *schematic) string {
	log.Debug("Solving part 1.")

	partNumbers := schematic.FindPartNumbers()
	log.Debugf("partNumbers = %v", partNumbers)
	sum := common.SumInt(partNumbers...)

	log.Debug("Part 1 solved.")
	return strconv.Itoa(sum)
}

func solvePart2(schematic *schematic) string {
	log.Debug("Solving part 2.")

	// TODO

	log.Debug("Part 2 solved.")
	return "Not Implemented"
}
