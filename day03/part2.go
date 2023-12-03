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

	symbols := map[string]bool{}
	for _, line := range lines {
		for _, r := range line {
			s := string(r)
			if s < "0" || s > "9" || s != "." {
				symbols[s] = true
			}
		}
	}

	fmt.Println(symbols)
}
