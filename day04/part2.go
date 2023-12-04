package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	total := 0
	copies := map[int]int{}
	for cardNum, line := range lines {
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
		totalCopies := 1 + copies[cardNum]
		for i := 0; i < matches; i++ {
			copies[cardNum+i+1] += totalCopies
		}
		total += totalCopies
	}

	fmt.Println(total)
}
