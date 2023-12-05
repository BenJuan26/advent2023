package advent2023

import (
	"bufio"
	"os"
	"strconv"
)

func ReadTestInput() ([]string, error) {
	return ReadFromFile("test.txt")
}

func ReadInput() ([]string, error) {
	return ReadFromFile("input.txt")
}

func ReadFromFile(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func ShouldRunPart1() bool {
	if len(os.Args) > 1 {
		if os.Args[1] == "part2" {
			return false
		} else if os.Args[1] == "part1" {
			return true
		} else {
			panic("invalid argument " + os.Args[1])
		}
	}

	return true
}

func ShouldRunPart2() bool {
	if len(os.Args) > 1 {
		if os.Args[1] == "part2" {
			return true
		} else if os.Args[1] == "part1" {
			return false
		} else {
			panic("invalid argument " + os.Args[1])
		}
	}

	return false
}
