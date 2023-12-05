package main

import (
	"math"
	"slices"

	"github.com/samber/lo"
)

func computeCard(c Card) (int, int) {
	count := 0
	for _, w := range c.Having {
		if slices.Contains(c.Winning, w) {
			count++
		}
	}
	return count, int(math.Pow(2, float64(count-1)))
}

func main() {
	loadFile("input.txt")
	total := 0
	for _, c := range cards {
		_, t := computeCard(c)
		total += t
	}
	println("Part1:", total)

	// part 2
	copies := make([]int, len(cards))
	for id := range cards {
		copies[id] = 1
	}
	for id, c := range cards {
		cb := copies[id]
		for i := 0; i < cb; i++ {
			count, _ := computeCard(c)
			for d := 0; d < count; d++ {
				copies[id+d+1]++
			}
		}
	}
	sum := lo.Reduce(copies, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	println("Part2:", sum)
}
