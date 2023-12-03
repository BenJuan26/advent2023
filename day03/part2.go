package main

import (
	"fmt"
	"strconv"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

func GetAsteriskPos(grid []string, y, startX, endX int) (int, int) {
	if startX > 0 && string(grid[y][startX-1]) == "*" {
		return y, startX - 1
	}
	if endX < len(grid[0])-1 && string(grid[y][endX+1]) == "*" {
		return y, endX + 1
	}
	for i := max(0, startX-1); i <= min(endX+1, len(grid[0])-1); i++ {
		if y > 0 && string(grid[y-1][i]) == "*" {
			return y - 1, i
		}
		if y < len(grid)-1 && string(grid[y+1][i]) == "*" {
			return y + 1, i
		}
	}
	return -1, -1
}

func GetNumBoundsAroundAsterisk(grid []string, asteriskPosY, asteriskPosX int) (int, int, int) {
	y := -1
	x := max(0, asteriskPosX-1)
	for ; x <= min(len(grid[0])-1, asteriskPosX+1); x++ {
		if asteriskPosY > 0 && digitRegex.Match([]byte{grid[asteriskPosY-1][x]}) {
			y = asteriskPosY - 1
			break
		}
		if asteriskPosY < len(grid)-1 && digitRegex.Match([]byte{grid[asteriskPosY+1][x]}) {
			y = asteriskPosY + 1
			break
		}
	}

	// not above or below
	if y < 0 {
		// to the left
		if asteriskPosX > 0 && digitRegex.Match([]byte{grid[asteriskPosY][asteriskPosX-1]}) {
			endX := asteriskPosX - 1
			startX := endX - 1
			for startX >= 0 && digitRegex.Match([]byte{grid[asteriskPosY][startX]}) {
				startX -= 1
			}
			startX += 1
			return asteriskPosY, startX, endX
		}

		// to the right
		if asteriskPosX < len(grid[0])-1 && digitRegex.Match([]byte{grid[asteriskPosY][asteriskPosX+1]}) {
			startX := asteriskPosX + 1
			endX := startX + 1
			for endX < len(grid[0]) && digitRegex.Match([]byte{grid[asteriskPosY][endX]}) {
				endX += 1
			}
			endX -= 1
			return asteriskPosY, startX, endX
		}

		// no number
		return -1, -1, -1
	}

	// part of the number is at grid[y][x]; need to find the start and end
	startX := x - 1
	for startX >= 0 && digitRegex.Match([]byte{grid[y][startX]}) {
		startX -= 1
	}
	startX += 1

	endX := x + 1
	for endX < len(grid[0]) && digitRegex.Match([]byte{grid[y][endX]}) {
		endX += 1
	}
	endX -= 1

	return y, startX, endX
}

func GetGearRatio(grid []string, y, startX, endX int) int {
	asteriskPosY, asteriskPosX := GetAsteriskPos(grid, y, startX, endX)
	if asteriskPosX < 0 {
		return 0 // no asterisk
	}
	// get the first number before erasing
	num1, err := strconv.Atoi(grid[y][startX : endX+1])
	if err != nil {
		panic(err)
	}
	// erase the number in the grid
	grid[y] = grid[y][:startX] + strings.Repeat(".", endX-startX+1) + grid[y][endX+1:]

	numY, numStartX, numEndX := GetNumBoundsAroundAsterisk(grid, asteriskPosY, asteriskPosX)
	if numY < 0 {
		return 0 // no second number
	}

	num2, err := strconv.Atoi(grid[numY][numStartX : numEndX+1])
	if err != nil {
		panic(err)
	}

	// don't need to bother erasing the second number
	return num1 * num2
}

func Part2() {
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
			total += GetGearRatio(grid, y, startX, endX)
			x = endX + 1
		}
	}

	fmt.Println(total)
}
