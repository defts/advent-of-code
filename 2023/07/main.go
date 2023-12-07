package main

import (
	"fmt"
	"slices"
	"sort"
)

const (
	highCard     = 1
	onePair      = 2
	twoPairs     = 3
	threeOfAKind = 4
	fullHouse    = 5
	fourOfAKind  = 6
	fiveOfAKind  = 7
)

func getValueOfCard(c rune, part2 bool) int {
	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if part2 {
			return 0
		}
		return 11
	case 'T':
		return 10
	default:
		return int(c - '0')
	}
}

func computeHandValue(h hand, isPart2 bool, countJoker int) int {
	if countJoker == 5 {
		return fiveOfAKind
	}
	c := map[rune]int{}
	for _, card := range h.cards {
		if isPart2 && card == rune('J') {
			continue
		}
		c[card]++
	}

	// sort by value
	// Create slice of key-value pairs
	pairs := make([][2]interface{}, 0, len(c))
	for k, v := range c {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	// Sort slice based on values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1].(int) > pairs[j][1].(int)
	})

	// Extract sorted keys
	keys := make([]rune, len(pairs))
	for i, p := range pairs {
		keys[i] = p[0].(rune)
	}

	for _, i := range keys {
		v := c[i]

		if v+countJoker == 5 {
			return fiveOfAKind
		}
		if v+countJoker == 4 {
			return fourOfAKind
		}
		if v+countJoker == 3 {
			for ii, v2 := range c {
				if v2 == 2 && i != ii {
					return fullHouse
				}
			}
			return threeOfAKind
		}
		if v+countJoker == 2 {
			for ii, v2 := range c {
				if v2 == 2 && i != ii {
					return twoPairs
				}
			}
			for ii, v2 := range c {
				if v2 == 3 && i != ii {
					return fullHouse
				}
			}
			return onePair
		}
	}
	return highCard
}
func computeHands(hands []hand, isPart2 bool) int {
	for i, h := range hands {
		if !isPart2 {
			hands[i].value = computeHandValue(h, isPart2, 0)
		}
		if isPart2 {
			// no J, compute as normal
			if !slices.Contains(h.cards, rune('J')) {
				hands[i].value = computeHandValue(h, isPart2, 0)
			} else {
				// count number of J
				countJoker := 0
				for _, card := range h.cards {
					if card == rune('J') {
						countJoker++
					}
				}
				hands[i].value = computeHandValue(h, isPart2, countJoker)
			}
		}

	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].value == hands[j].value {
			for ii := range hands[i].cards {
				if getValueOfCard(hands[i].cards[ii], isPart2) != getValueOfCard(hands[j].cards[ii], isPart2) {
					return getValueOfCard(hands[i].cards[ii], isPart2) > getValueOfCard(hands[j].cards[ii], isPart2)
				}
			}
		}
		return hands[i].value > hands[j].value
	})

	slices.Reverse(hands)

	res := 0
	for i, h := range hands {
		res += h.bid * (i + 1)
	}
	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", computeHands(hands, false))
	fmt.Println("part2: ", computeHands(hands, true))

}
