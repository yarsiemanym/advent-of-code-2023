package day03

import (
	"strconv"
	"unicode"

	"github.com/yarsiemanym/advent-of-code-2023/common"
)

type schematic struct {
	plane *common.BoundedPlane
}

func parseSchematic(text string) *schematic {

	lines := common.Split(text, "\n")

	plane := common.NewBoundedPlane(len(lines)-1, len(lines[0]))

	for y, line := range lines {
		if line == "" {
			continue
		}

		for x, character := range line {
			plane.SetValueAt(common.New2DPoint(x, y), character)
		}
	}

	return &schematic{
		plane: plane,
	}
}

func (thisSchematic *schematic) FindPartNumbers() []int {

	partNumbers := []int{}

	for y := thisSchematic.plane.Span().Start().Y(); y <= thisSchematic.plane.Span().End().Y(); y++ {
		number := []rune{}
		isPartNumber := false

		for x := thisSchematic.plane.Span().Start().X(); x <= thisSchematic.plane.Span().End().X(); x++ {
			thisPoint := common.New2DPoint(x, y)

			thisCharacter := thisSchematic.plane.GetValueAt(thisPoint).(rune)
			if !unicode.IsDigit(thisCharacter) {
				if len(number) > 0 && isPartNumber {
					partNumber, err := strconv.Atoi(string(number))
					common.Check(err)
					partNumbers = append(partNumbers, partNumber)
				}

				number = []rune{}
				isPartNumber = false
				continue
			} else {
				number = append(number, thisCharacter)

				if !isPartNumber {
					neighboringPoints := thisSchematic.plane.GetMooreNeighbors(thisPoint)

					for _, neighboringPoint := range neighboringPoints {
						neighboringCharacter := thisSchematic.plane.GetValueAt(neighboringPoint).(rune)

						if !unicode.IsDigit(neighboringCharacter) && neighboringCharacter != '.' {
							isPartNumber = true
							break
						}
					}
				}
			}
		}

		if len(number) > 0 && isPartNumber {
			partNumber, err := strconv.Atoi(string(number))
			common.Check(err)
			partNumbers = append(partNumbers, partNumber)
		}
	}

	return partNumbers
}
