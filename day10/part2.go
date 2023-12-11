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

	total := 0
	dist, _ := GetDistances(lines)
	for y, line := range lines {
		inside := false
		for x, char := range line {
			if dist[y][x] != -1 && (char == 'L' || char == 'J' || char == '|') {
				inside = !inside
			} else if dist[y][x] == -1 && inside {
				total += 1
			}
		}
	}

	fmt.Println(total)
}
