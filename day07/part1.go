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
		if countsByNumber[1] == 4 || countsByNumber[1] == 1 {
			return FiveOfAKind
		}
		return FourOfAKind
	}
	if countsByTotal[3] == 1 {
		if countsByNumber[1] == 3 {
			if countsByTotal[2] == 1 {
				return FiveOfAKind
			}
			return FourOfAKind
		}
		if countsByTotal[2] == 1 {
			if countsByNumber[1] == 2 {
				return FiveOfAKind
			}
			return FullHouse
		}
		if countsByNumber[1] == 1 {
			return FourOfAKind
		}
		return ThreeOfAKind
	}
	if countsByTotal[2] == 2 {
		if countsByNumber[1] == 2 {
			return FourOfAKind
		}
		if countsByNumber[1] == 1 {
			return FullHouse
		}
		return TwoPair
	}
	if countsByTotal[2] == 1 {
		if countsByNumber[1] == 1 || countsByNumber[1] == 2 {
			return ThreeOfAKind
		}
		return OnePair
	}
	if countsByTotal[1] == 5 {
		if countsByNumber[1] == 1 {
			return OnePair
		}
		return HighCard
	}
	panic("shouldn't reach here")
}

func PopulateHands(lines []string, jokersWild bool) []Hand {
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
				if jokersWild {
					cards = append(cards, 1)
				} else {
					cards = append(cards, 11)
				}
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

	return hands
}

func SortHands(hands []Hand) {
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
}

func Part1() {
	lines, err := advent.ReadInput()
	if err != nil {
		panic(err)
	}

	hands := PopulateHands(lines, false)
	SortHands(hands)

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
