package main

import (
	"fmt"
	"math"
	"regexp"

	advent "github.com/BenJuan26/advent2023"
)

type Race struct {
	Time     int
	Distance int
}

func NearlyEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1e-9
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	// result is always a quadratic function with formula
	// y = -(x)(x-raceTime)
	// this multiplies out into:
	// y = -x^2 + raceTime*x
	// quadratic formula with a=-1, b=raceTime, c=0
	// to find the roots of the distance record, move y=distance over:
	// 0 = -x^2 + raceTime*x - distance
	// the use the quadratic formula to solve for x roots
	re := regexp.MustCompile(`[0-9]+`)
	races := []Race{}
	times := re.FindAllString(lines[0], -1)
	for _, t := range times {
		races = append(races, Race{Time: advent.MustAtoi(t)})
	}

	distances := re.FindAllString(lines[1], -1)
	for i, d := range distances {
		races[i].Distance = advent.MustAtoi(d)
	}

	product := 1
	for _, race := range races {
		discriminant := math.Sqrt(float64(race.Time*race.Time - 4*race.Distance))
		root1 := (-float64(race.Time) + discriminant) / -2
		root2 := (-float64(race.Time) - discriminant) / -2

		min := math.Ceil(root1)
		if NearlyEqual(min, root1) {
			min += 1
		}

		max := math.Floor(root2)
		if NearlyEqual(max, root2) {
			max -= 1
		}

		numWays := int(max-min) + 1
		product *= numWays
	}

	fmt.Println(product)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
