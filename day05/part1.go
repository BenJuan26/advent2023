package main

import (
	"fmt"
	"sort"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

type MapRange struct {
	Source      int
	Destination int
	Length      int
}

type Map struct {
	Ranges []MapRange
}

func (m *Map) Sort() {
	sort.Slice(m.Ranges, func(i, j int) bool {
		return m.Ranges[i].Source < m.Ranges[j].Source
	})
}

func (m *Map) Map(num int) int {
	for _, r := range m.Ranges {
		if num >= r.Source && r.Source+r.Length > num {
			return r.Destination + (num - r.Source)
		}
	}
	return num
}

func PopulateMaps(lines []string) []Map {
	lineNum := 3
	currentMap := Map{}
	maps := []Map{}
	for lineNum < len(lines) {
		line := lines[lineNum]
		if line == "" {
			currentMap.Sort()
			maps = append(maps, currentMap)
			currentMap = Map{}
			lineNum += 2
			continue
		}
		fields := strings.Fields(line)
		r := MapRange{
			Destination: advent.MustAtoi(fields[0]),
			Source:      advent.MustAtoi(fields[1]),
			Length:      advent.MustAtoi(fields[2]),
		}
		currentMap.Ranges = append(currentMap.Ranges, r)
		lineNum += 1
	}
	// no blank line for the last map
	currentMap.Sort()
	maps = append(maps, currentMap)

	return maps
}

func Part1() {
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

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
