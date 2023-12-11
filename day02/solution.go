package day02

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2023/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseGame)
	games := []*game{}

	for _, result := range results {
		games = append(games, result.(*game))
	}

	answer := common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(games),
		Part2: solvePart2(games),
	}

	return answer
}

func solvePart1(games []*game) string {
	log.Debug("Solving part 1.")

	sumOfValidGameIds := 0

	for _, game := range games {
		if game.IsValid(bagContents) {
			sumOfValidGameIds += game.Id
		}
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(sumOfValidGameIds)
}

func solvePart2(games []*game) string {
	log.Debug("Solving part 2.")

	power := 0

	for _, game := range games {
		minimumBagContents := game.MinimumBag()

		power += minimumBagContents[Red] * minimumBagContents[Green] * minimumBagContents[Blue]
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(power)
}

var bagContents = map[Color]int{
	Red:   12,
	Green: 13,
	Blue:  14,
}
