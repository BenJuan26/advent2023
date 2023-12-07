package main

import (
	"fmt"
	"sort"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

const (
	HighCard     = iota
	OnePair      = iota
	TwoPair      = iota
	ThreeOfAKind = iota
	FullHouse    = iota
	FourOfAKind  = iota
	FiveOfAKind  = iota
)

type Hand struct {
	Cards []int
	Bid   int
}

func (h Hand) Type() int {
	countsByNumber := map[int]int{}
	countsByTotal := map[int]int{}
	for _, card := range h.Cards {
		countsByNumber[card] += 1
		n := countsByNumber[card]
		if n > 1 {
			countsByTotal[n-1] -= 1
		}
		countsByTotal[n] += 1
	}

	if countsByTotal[5] == 1 {
		return FiveOfAKind
	}
	if countsByTotal[4] == 1 {
		return FourOfAKind
	}
	if countsByTotal[3] == 1 {
		if countsByTotal[2] == 1 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	}
	if countsByTotal[2] == 2 {
		return TwoPair
	}
	if countsByTotal[2] == 1 {
		return OnePair
	}
	if countsByTotal[1] == 5 {
		return HighCard
	}
	panic("shouldn't reach here")
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	hands := []Hand{}
	for _, line := range lines {
		fields := strings.Split(line, " ")
		hand := Hand{}
		cards := []int{}
		for _, card := range fields[0] {
			switch card {
			case 'A':
				cards = append(cards, 14)
			case 'K':
				cards = append(cards, 13)
			case 'Q':
				cards = append(cards, 12)
			case 'J':
				cards = append(cards, 11)
			case 'T':
				cards = append(cards, 10)
			default:
				cards = append(cards, advent.MustAtoi(string(card)))
			}
		}
		hand.Cards = cards
		hand.Bid = advent.MustAtoi(fields[1])
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		type1 := hands[i].Type()
		type2 := hands[j].Type()
		if type1 != type2 {
			return type1 < type2
		}
		for index := 0; index < 5; index++ {
			if hands[i].Cards[index] == hands[j].Cards[index] {
				continue
			}
			return hands[i].Cards[index] < hands[j].Cards[index]
		}

		return false
	})

	total := 0
	for i, hand := range hands {
		rank := i + 1
		total += rank * hand.Bid
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
