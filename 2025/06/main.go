package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func calc(initial int, nums []int, ope string) int {
	res := initial
	switch ope {
	case "+":
		for _, i := range nums {
			res += i
		}
	case "*":
		for _, i := range nums {
			res *= i
		}
	}
	return res
}

func part1(input arg) int {
	defer utils.Timer("part1")()
	res := 0

	for _, ope := range input.ope {
		res += calc(ope.nums[0], ope.nums[1:len(ope.nums)], ope.op)
	}

	return res
}

func part2(input arg) int {
	defer utils.Timer("part2")()
	res := 0

	for _, ope := range input.ope {
		res += calc(ope.numsPart2[0], ope.numsPart2[1:len(ope.numsPart2)], ope.op)
	}

	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}
