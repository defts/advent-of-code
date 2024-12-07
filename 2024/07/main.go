package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func part1(equations map[int][]int) int64 {
	defer utils.Timer("part1")()
	res := int64(0)

	var calculate func(index, current int, result int, num []int) bool

	calculate = func(index, current int, result int, nums []int) bool {
		// Si on atteint la fin et que le résultat correspond, c'est valide
		if index == len(nums) {
			return current == result
		}

		// On essaie soit d'additionner soit de multiplier
		return calculate(index+1, current+nums[index], result, nums) || calculate(index+1, current*nums[index], result, nums)
	}

	for k, v := range equations {
		if calculate(1, v[0], k, v) {
			res += int64(k)
		}
	}
	return res
}

func part2(equations map[int][]int) int64 {
	defer utils.Timer("part1")()
	res := int64(0)
	var calculate func(index, current int, result int, num []int) bool

	calculate = func(index, current int, result int, nums []int) bool {
		// Si on atteint la fin et que le résultat correspond, c'est valide
		if index == len(nums) {
			return current == result
		}

		// On essaie soit d'additionner soit de multiplier // soit de concaténer
		return calculate(index+1, current+nums[index], result, nums) ||
			calculate(index+1, current*nums[index], result, nums) ||
			calculate(index+1, utils.MustParseInt(fmt.Sprintf("%d%d", current, nums[index])), result, nums)
	}

	for k, v := range equations {
		if calculate(1, v[0], k, v) {
			res += int64(k)
		}
	}
	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(equations))
	fmt.Println("part2: ", part2(equations))
}
