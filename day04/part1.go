package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	total := 0

	for _, line := range lines {
		card := strings.Split(line, ": ")[1]
		cardParts := strings.Split(card, " | ")
		winningNums := map[int]bool{}
		for _, numString := range strings.Fields(cardParts[0]) {
			num, err := strconv.Atoi(numString)
			if err != nil {
				panic(err)
			}
			winningNums[num] = true
		}
		matches := 0
		for _, numString := range strings.Fields(cardParts[1]) {
			num, err := strconv.Atoi(numString)
			if err != nil {
				panic(err)
			}
			if winningNums[num] {
				matches += 1
			}
		}
		if matches > 0 {
			total += int(math.Pow(2., float64(matches-1)))
		}
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
