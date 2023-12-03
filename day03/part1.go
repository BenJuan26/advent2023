package main

import (
	"fmt"
	"regexp"
	"strconv"

	advent "github.com/BenJuan26/advent2023"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var digitOrDot = regexp.MustCompile(`[0-9\.]`)
var digitRegex = regexp.MustCompile(`[0-9]`)

func IsPartNumber(grid []string, y, startX, endX int) bool {
	if !digitOrDot.Match([]byte{grid[y][max(0, startX-1)]}) || !digitOrDot.Match([]byte{grid[y][min(endX+1, len(grid[0])-1)]}) {
		return true
	}
	for i := max(0, startX-1); i <= min(endX+1, len(grid[0])-1); i++ {
		if y > 0 && !digitOrDot.Match([]byte{grid[y-1][i]}) || y < len(grid)-1 && !digitOrDot.Match([]byte{grid[y+1][i]}) {
			return true
		}
	}
	return false
}

func Part1() {
	grid, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		x := 0
		for x < len(grid[0]) {
			if !digitRegex.Match([]byte{grid[y][x]}) {
				x += 1
				continue
			}
			// start of a number
			startX := x
			endX := x
			for {
				if x < len(grid[0])-1 && digitRegex.Match([]byte{grid[y][x+1]}) {
					x += 1
				} else {
					endX = x
					break
				}
			}
			if IsPartNumber(grid, y, startX, endX) {
				num, err := strconv.Atoi(grid[y][startX : endX+1])
				if err != nil {
					panic(err)
				}
				total += num
			}
			x = endX + 1
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
