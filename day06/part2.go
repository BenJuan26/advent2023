package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

func Part2() {
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
	times := re.FindAllString(lines[0], -1)
	timeString := strings.Join(times, "")
	time := advent.MustAtoi(timeString)

	distances := re.FindAllString(lines[1], -1)
	distancesString := strings.Join(distances, "")
	distance := advent.MustAtoi(distancesString)

	discriminant := math.Sqrt(float64(time*time - 4*distance))
	root1 := (-float64(time) + discriminant) / -2
	root2 := (-float64(time) - discriminant) / -2

	min := math.Ceil(root1)
	if NearlyEqual(min, root1) {
		min += 1
	}

	max := math.Floor(root2)
	if NearlyEqual(max, root2) {
		max -= 1
	}

	numWays := int(max-min) + 1

	fmt.Println(numWays)
}
