package day02

import (
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2023/common"
)

type game struct {
	Id    int
	Turns []*turn
}

type turn struct {
	Id           int
	ReveledCubes map[Color]int
}

type Color string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

func parseGame(text string) interface{} {

	if text == "" {
		return nil
	}

	log.Debugf("Parsing game '%s'", text)

	gameRegex := regexp.MustCompile(`Game (\d+):(.*)`)
	matches := gameRegex.FindStringSubmatch(text)

	id, err := strconv.Atoi(matches[1])
	common.Check(err)
	turnsText := matches[2]
	game := &game{
		Id:    id,
		Turns: []*turn{},
	}

	turns := common.Split(turnsText, ";")

	for index, turnMatch := range turns {
		game.Turns = append(game.Turns, parseTurn(index+1, turnMatch))
	}

	return game
}

func parseTurn(index int, text string) *turn {
	redRegex := regexp.MustCompile(`(\d+) red`)
	redMatches := redRegex.FindStringSubmatch(text)
	redCubes := 0
	if redMatches != nil {
		value, err := strconv.Atoi(redMatches[1])
		common.Check(err)
		redCubes = value
	}

	greenRegex := regexp.MustCompile(`(\d+) green`)
	greenMatches := greenRegex.FindStringSubmatch(text)
	greenCubes := 0
	if greenMatches != nil {
		value, err := strconv.Atoi(greenMatches[1])
		common.Check(err)
		greenCubes = value
	}

	blueRegex := regexp.MustCompile(`(\d+) blue`)
	blueMatches := blueRegex.FindStringSubmatch(text)
	blueCubes := 0
	if blueMatches != nil {
		value, err := strconv.Atoi(blueMatches[1])
		common.Check(err)
		blueCubes = value
	}

	return &turn{
		Id: index,
		ReveledCubes: map[Color]int{
			Red:   redCubes,
			Green: greenCubes,
			Blue:  blueCubes,
		},
	}
}

func (thisGame *game) IsValid(bagContents map[Color]int) bool {
	for _, turn := range thisGame.Turns {
		if turn.ReveledCubes[Red] > bagContents[Red] ||
			turn.ReveledCubes[Green] > bagContents[Green] ||
			turn.ReveledCubes[Blue] > bagContents[Blue] {
			return false
		}
	}

	return true
}

func (thisGame *game) MinimumBag() map[Color]int {
	minimumBagContents := map[Color]int{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for _, turn := range thisGame.Turns {
		minimumBagContents[Red] = common.MaxInt(minimumBagContents[Red], turn.ReveledCubes[Red])
		minimumBagContents[Green] = common.MaxInt(minimumBagContents[Green], turn.ReveledCubes[Green])
		minimumBagContents[Blue] = common.MaxInt(minimumBagContents[Blue], turn.ReveledCubes[Blue])
	}

	return minimumBagContents
}
