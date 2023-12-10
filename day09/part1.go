package main

import (
	"fmt"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

func Extrapolate(vals []int) int {
	diffs := []int{}
	done := true
	for i := range vals {
		if i == len(vals)-1 {
			break
		}
		diff := vals[i+1] - vals[i]
		if diff != 0 {
			done = false
		}
		diffs = append(diffs, diff)
	}
	lastVal := vals[len(vals)-1]
	if done {
		return lastVal
	}
	return lastVal + Extrapolate(diffs)
}

func ExtrapolateBackwards(vals []int) int {
	diffs := []int{}
	done := true
	for i := range vals {
		if i == len(vals)-1 {
			break
		}
		diff := vals[i+1] - vals[i]
		if diff != 0 {
			done = false
		}
		diffs = append(diffs, diff)
	}
	firstVal := vals[0]
	if done {
		return firstVal
	}
	ext := ExtrapolateBackwards(diffs)
	return firstVal - ext
}

func Part1() {
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
		total += Extrapolate(nums)
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
