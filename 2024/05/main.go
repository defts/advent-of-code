package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
	"golang.org/x/exp/slices"
)

func compute(order map[int][]int, print [][]int) (correct [][]int, incorrect [][]int) {
	for _, p := range print {
		keep := true
		for i := 0; i < len(p); i++ {
			if _, ok := order[p[i]]; !ok {
				continue
			}
			for j := i + 1; j < len(p); j++ {
				for _, n := range order[p[i]] {
					if p[j] == n {
						keep = false
					}
				}
			}
		}
		if keep {
			correct = append(correct, p)
		} else {
			incorrect = append(incorrect, p)
		}
	}
	return correct, incorrect
}

func part1(order map[int][]int, print [][]int) int {
	defer utils.Timer("part1")()
	res := 0
	correct, _ := compute(order, print)
	for _, c := range correct {
		res += c[(len(c)-1)/2]
	}
	return res
}

func part2(order map[int][]int, print [][]int) int {
	defer utils.Timer("part2")()
	res := 0
	_, incorrect := compute(order, print)
	for _, r := range incorrect {
		slices.SortFunc(r, func(first, second int) bool {
			if numbersAfter, exists := order[second]; exists {
				return slices.Contains(numbersAfter, first)
			}
			return false
		})
		res += r[(len(r)-1)/2]
	}
	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(order, print))
	fmt.Println("part2: ", part2(order, print))
}
