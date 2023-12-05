package main

import (
	"fmt"
	"sort"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	seeds := []int{}
	seedsString := strings.Split(lines[0], ": ")[1]
	seedsStrings := strings.Fields(seedsString)
	for _, s := range seedsStrings {
		seeds = append(seeds, advent.MustAtoi(s))
	}

	maps := PopulateMaps(lines)

	locations := []int{}
	for _, seed := range seeds {
		num := seed
		for _, m := range maps {
			num = m.Map(num)
		}
		locations = append(locations, num)
	}

	sort.Ints(locations)
	fmt.Println(locations[0])
}
