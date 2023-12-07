package main

import (
	"fmt"

	advent "github.com/BenJuan26/advent2023"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	hands := PopulateHands(lines, true)
	SortHands(hands)

	total := 0
	for i, hand := range hands {
		rank := i + 1
		total += rank * hand.Bid
	}

	fmt.Println(total)
}
