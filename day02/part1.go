package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2022"
)

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	total := 0

	for i, game := range lines {
		possible := true
		parts := strings.Split(game, ": ")
		gameId := i + 1
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
				if maxCubes[color] < num {
					possible = false
					break
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			total += gameId
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
