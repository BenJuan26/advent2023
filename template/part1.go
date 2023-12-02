package main

import (
	"fmt"

	advent "github.com/BenJuan26/advent2023"
)

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		line += " " // do something
	}

	fmt.Println("answer")
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
