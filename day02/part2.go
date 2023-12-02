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

	for _, game := range lines {
		maxInGame := map[string]int{}
		parts := strings.Split(game, ": ")
		groups := parts[1]
		sets := strings.Split(groups, "; ")
		for _, set := range sets {
			for _, colorField := range strings.Split(set, ", ") {
				fields := strings.Split(colorField, " ")
				color := fields[1]
				num, err := strconv.Atoi(fields[0])
				if err != nil {
					panic(err)
				}
				if maxInGame[color] < num {
					maxInGame[color] = num
				}
			}
		}
		power := maxInGame["red"] * maxInGame["green"] * maxInGame["blue"]
		total += power
	}

	fmt.Println(total)
}
