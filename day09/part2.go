package main

import (
	"fmt"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	total := 0
	for _, line := range lines {
		nums := []int{}
		numStrings := strings.Fields(line)
		for _, ns := range numStrings {
			num := advent.MustAtoi(ns)
			nums = append(nums, num)
		}
		total += ExtrapolateBackwards(nums)
	}

	fmt.Println(total)
}
