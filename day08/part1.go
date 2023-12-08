package main

import (
	"fmt"

	advent "github.com/BenJuan26/advent2023"
)

type Node struct {
	ID     string
	Left   string
	Right  string
	Period int
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	instructions := lines[0]
	nodes := map[string]Node{}
	for i := 2; i < len(lines); i += 1 {
		line := lines[i]
		nodeID := line[:3]
		l := line[7:10]
		r := line[12:15]
		nodes[nodeID] = Node{ID: nodeID, Left: l, Right: r}
	}

	node := nodes["AAA"]
	i := 0
	steps := 0
	for node.ID != "ZZZ" {
		if i >= len(instructions) {
			i = 0
		}
		dir := instructions[i]
		if dir == 'L' {
			node = nodes[node.Left]
		} else {
			node = nodes[node.Right]
		}
		i += 1
		steps += 1
	}

	fmt.Println(steps)
}

func main() {
	if advent.ShouldRunPart1() {
		Part1()
	}

	if advent.ShouldRunPart2() {
		Part2()
	}
}
