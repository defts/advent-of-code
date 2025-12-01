package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func floorDiv(a, b int) int {
	if a >= 0 {
		return a / b
	}
	return (a - b + 1) / b
}

func countZeros(lo, hi int) int {
	if lo > hi {
		return 0
	}
	return floorDiv(hi, 100) - floorDiv(lo-1, 100)
}

func part1(input arg) int {
	defer utils.Timer("part1")()
	res := 50
	resToZero := 0

	for _, i := range input.input {
		switch {
		case i[0] == 'R':
			res = (res + utils.MustParseInt(i[1:])) % 100
		case i[0] == 'L':
			res = (res - utils.MustParseInt(i[1:])) % 100
			if res < 0 {
				res += 100
			}
		}
		if res == 0 {
			resToZero++
		}
	}

	return resToZero
}

func part2(input arg) int {
	defer utils.Timer("part2")()
	res := 50
	resToZero := 0

	for _, i := range input.input {
		step := utils.MustParseInt(i[1:])
		switch {
		case i[0] == 'R':
			resToZero += countZeros(res+1, res+step)

			res = (res + step) % 100
		case i[0] == 'L':
			resToZero += countZeros(res-step, res-1)

			res = (res - step) % 100
			if res < 0 {
				res += 100
			}
		}
	}

	return resToZero
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}
