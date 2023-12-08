package main

import (
	"fmt"

	advent "github.com/BenJuan26/advent2023"
)

func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func LCMM(nums []int) int {
	if len(nums) == 2 {
		return LCM(nums[0], nums[1])
	}
	a := nums[0]
	nums = nums[1:]
	return LCM(a, LCMM(nums))
}

func Part2() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	instructions := lines[0]
	nodes := map[string]Node{}
	currentNodes := []Node{}
	for i := 2; i < len(lines); i += 1 {
		line := lines[i]
		nodeID := line[:3]
		l := line[7:10]
		r := line[12:15]
		n := Node{ID: nodeID, Left: l, Right: r, Period: -1}
		nodes[nodeID] = n
		if n.ID[2] == 'A' {
			currentNodes = append(currentNodes, n)
		}
	}

	i := 0
	steps := 0
	done := false
	for !done {
		steps += 1
		done = true
		if i >= len(instructions) {
			i = 0
		}
		for nodeIndex, node := range currentNodes {
			if node.Period >= 0 {
				continue
			}
			nextNode := Node{}
			dir := instructions[i]
			if dir == 'L' {
				nextNode = nodes[node.Left]
			} else {
				nextNode = nodes[node.Right]
			}
			if nextNode.ID[2] == 'Z' {
				nextNode.Period = steps
			} else {
				done = false
			}
			currentNodes[nodeIndex] = nextNode
		}
		i += 1
	}

	periods := []int{}
	for _, n := range currentNodes {
		periods = append(periods, n.Period)
	}

	fmt.Println(LCMM(periods))
}
