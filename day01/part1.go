package main

import (
	"fmt"
	"strconv"

	advent "github.com/BenJuan26/advent2022"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	total := 0
	for _, line := range lines {
		firstDigit := -1
		lastDigit := -1
		for _, char := range line {
			digit, err := strconv.Atoi(string(char))
			if err == nil {
				if firstDigit < 0 {
					firstDigit = digit
				}
				lastDigit = digit
			}
		}
		value := firstDigit*10 + lastDigit
		total += value
	}

	fmt.Println(total)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
