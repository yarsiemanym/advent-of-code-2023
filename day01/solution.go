package day01

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2023/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n", parseLine)

	var lines []string
	for _, result := range results {
		lines = append(lines, result.(string))
	}

	answer := common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(lines),
		Part2: solvePart2(lines),
	}

	return answer
}

func solvePart1(lines []string) string {
	log.Debug("Solving part 1.")

	calibrationValue := 0

	for _, line := range lines {
		firstDigit := findFirstDigit(line)
		lastDigit := findLastDigit(line)
		twoDigitNumber := combine(firstDigit, lastDigit)
		log.Debugf("twoDigitNumber = %d", twoDigitNumber)
		calibrationValue += twoDigitNumber
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(calibrationValue)
}

func solvePart2(lines []string) string {
	log.Debug("Solving part 2.")

	calibrationValue := 0

	for _, line := range lines {
		line = clean(line)
		firstDigit := findFirstDigit(line)
		lastDigit := findLastDigit(line)
		twoDigitNumber := combine(firstDigit, lastDigit)
		log.Debugf("twoDigitNumber = %d", twoDigitNumber)
		calibrationValue += twoDigitNumber
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(calibrationValue)
}

func parseLine(text string) interface{} {
	if text == "" {
		return nil
	}

	line := strings.TrimSpace(text)

	log.Debugf("Parsing line '%s'", line)

	return line
}

func findFirstDigit(line string) *rune {
	characters := []rune(line)

	for i := 0; i < len(characters); i++ {
		character := characters[i]

		if unicode.IsDigit(character) {
			return &character
		}
	}

	return nil
}

func findLastDigit(line string) *rune {
	characters := []rune(line)

	for i := len(characters) - 1; i >= 0; i-- {
		character := characters[i]

		if unicode.IsDigit(character) {
			return &character
		}
	}

	return nil
}

func combine(firstDigit *rune, secondDigit *rune) int {
	if firstDigit == nil || secondDigit == nil {
		return 0
	}

	twoDigitNumber, err := strconv.Atoi(fmt.Sprintf("%c%c", *firstDigit, *secondDigit))
	common.Check(err)
	return twoDigitNumber
}

func clean(line string) string {

	newLine := []rune{}

	for i := 0; i < len(line); i++ {
		character := rune(line[i])

		if unicode.IsDigit(character) {
			newLine = append(newLine, character)
		} else {
			for key, value := range digits {
				if i+len(key) <= len(line) && line[i:i+len(key)] == key {
					newLine = append(newLine, value)
					break
				}
			}
		}
	}

	log.Debugf("Replacing words in line '%s' results in line '%s'", line, string(newLine))

	return string(newLine)
}

var digits = map[string]rune{
	"zero":  '0',
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}
