package main

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/defts/advent-of-code/utils"
)

func part1(input arg) int {
	defer utils.Timer("part1")()
	res := 0

	for _, i := range input.ingredients {
		for _, r := range input.ranges {
			if i >= r[0] && i <= r[1] {
				res++
				break
			}
		}
	}

	return res
}

func part2(input arg) int {
	defer utils.Timer("part2")()
	res := 0

	// sort ranges
	slices.SortFunc(input.ranges, func(a, b [2]int) int { return cmp.Compare(a[0], b[0]) })

	aggRanges := [][2]int{input.ranges[0]}
	for _, r := range input.ranges[1:] {
		if last := &aggRanges[len(aggRanges)-1]; r[0] <= last[1] {
			last[1] = max(last[1], r[1])
		} else {
			aggRanges = append(aggRanges, r)
		}
	}

	for _, r := range aggRanges {
		res += r[1] - r[0] + 1
	}

	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}
