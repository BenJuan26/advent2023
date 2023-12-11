package main

import (
	"fmt"

	advent "github.com/BenJuan26/advent2023"
)

const (
	Left  = iota
	Right = iota
	Up    = iota
	Down  = iota
)

func fromLeft(grid []string, y, x int) int {
	char := grid[y][x]
	if char == 'J' {
		return Up
	}
	if char == '7' {
		return Down
	}
	if char == '-' {
		return Right
	}

	return -1
}

func fromRight(grid []string, y, x int) int {
	char := grid[y][x]
	if char == 'L' {
		return Up
	}
	if char == 'F' {
		return Down
	}
	if char == '-' {
		return Left
	}

	return -1
}

func fromTop(grid []string, y, x int) int {
	char := grid[y][x]
	if char == 'J' {
		return Left
	}
	if char == 'L' {
		return Right
	}
	if char == '|' {
		return Down
	}

	return -1
}

func fromBottom(grid []string, y, x int) int {
	char := grid[y][x]
	if char == 'F' {
		return Right
	}
	if char == '7' {
		return Left
	}
	if char == '|' {
		return Up
	}

	return -1
}

func getConnectedDirs(grid []string, y, x int) (int, int) {
	dirs := []int{}
	if y > 0 {
		dir := fromBottom(grid, y-1, x)
		if dir >= 0 {
			dirs = append(dirs, Up)
		}
	}
	if y < len(grid)-1 {
		dir := fromTop(grid, y+1, x)
		if dir >= 0 {
			dirs = append(dirs, Down)
		}
	}
	if x > 0 {
		dir := fromRight(grid, y, x-1)
		if dir >= 0 {
			dirs = append(dirs, Left)
		}
	}
	if x < len(grid[0])-1 {
		dir := fromLeft(grid, y, x+1)
		if dir >= 0 {
			dirs = append(dirs, Right)
		}
	}

	if len(dirs) != 2 {
		panic("too many/few dirs connected")
	}

	if dirs[0] == Up || dirs[1] == Up {
		if dirs[0] == Left || dirs[1] == Left {
			grid[y] = grid[y][:x] + "J" + grid[y][x+1:]
		} else if dirs[0] == Right || dirs[1] == Right {
			grid[y] = grid[y][:x] + "L" + grid[y][x+1:]
		} else { // Down
			grid[y] = grid[y][:x] + "|" + grid[y][x+1:]
		}
	} else if dirs[0] == Left || dirs[1] == Left {
		if dirs[0] == Right || dirs[1] == Right {
			grid[y] = grid[y][:x] + "-" + grid[y][x+1:]
		} else { // Down
			grid[y] = grid[y][:x] + "7" + grid[y][x+1:]
		}
	} else { // Down and right
		grid[y] = grid[y][:x] + "F" + grid[y][x+1:]
	}

	return dirs[0], dirs[1]
}

func GetDistances(grid []string) ([][]int, int) {
	dist := [][]int{}
	for _, line := range grid {
		d := []int{}
		for range line {
			d = append(d, -1)
		}
		dist = append(dist, d)
	}

	sy := -1
	sx := -1
	for y, line := range grid {
		for x, char := range line {
			if char == 'S' {
				sy = y
				sx = x
				break
			}
		}
		if sy >= 0 {
			break
		}
	}

	dir1, dir2 := getConnectedDirs(grid, sy, sx)
	y1 := sy
	x1 := sx
	y2 := sy
	x2 := sx
	maxDist := 0
	for dist[y1][x1] == -1 && dist[y2][x2] == -1 {
		dist[y1][x1] = maxDist
		dist[y2][x2] = maxDist
		maxDist += 1
		switch dir1 {
		case Up:
			y1 -= 1
			dir1 = fromBottom(grid, y1, x1)
		case Down:
			y1 += 1
			dir1 = fromTop(grid, y1, x1)
		case Left:
			x1 -= 1
			dir1 = fromRight(grid, y1, x1)
		case Right:
			x1 += 1
			dir1 = fromLeft(grid, y1, x1)
		}

		switch dir2 {
		case Up:
			y2 -= 1
			dir2 = fromBottom(grid, y2, x2)
		case Down:
			y2 += 1
			dir2 = fromTop(grid, y2, x2)
		case Left:
			x2 -= 1
			dir2 = fromRight(grid, y2, x2)
		case Right:
			x2 += 1
			dir2 = fromLeft(grid, y2, x2)
		}
	}

	return dist, maxDist - 1
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	_, maxDist := GetDistances(lines)
	fmt.Println(maxDist)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
