package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

var writtenDigits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func ParseDigit(chars string) int {
	lastChar := chars[len(chars)-1]
	digit, err := strconv.Atoi(string(lastChar))
	if err == nil {
		return digit
	}

	for i, writtenDigit := range writtenDigits {
		if strings.HasSuffix(chars, writtenDigit) {
			return i + 1
		}
	}

	return -1
}

func Part2() {
	lines, err := advent.ReadTestInput()
	if err != nil {
		panic(err)
	}

	total := 0
	for _, line := range lines {
		firstDigit := -1
		lastDigit := -1
		cursorStart := 0
		cursorEnd := 1
		for cursorEnd < len(line) {
			digit := ParseDigit(line[cursorStart:cursorEnd])
			cursorEnd += 1
			if digit >= 0 {
				if firstDigit < 0 {
					firstDigit = digit
				}
				lastDigit = digit
				cursorStart = cursorEnd + 1
				cursorEnd += 1
			}
		}
		value := firstDigit*10 + lastDigit
		total += value
	}

	fmt.Println(total)
}
