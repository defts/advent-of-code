package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func part1(input arg) int {
	defer utils.Timer("part1")()
	res := 0

	for _, bank := range input.banks {
		indexMax := 0
		max := ""
		for b := 2; b > 0; b-- {
			for i := indexMax; i <= len(bank)-b; i++ {
				if bank[i] > bank[indexMax] {
					indexMax = i
				}
			}
			max += string(bank[indexMax])
			indexMax++
		}
		res += utils.MustParseInt(max)
	}

	return res
}

func part2(input arg) int {
	defer utils.Timer("part2")()
	res := 0

	for _, bank := range input.banks {
		indexMax := 0
		max := ""
		for b := 12; b > 0; b-- {
			for i := indexMax; i <= len(bank)-b; i++ {
				if bank[i] > bank[indexMax] {
					indexMax = i
				}
			}
			max += string(bank[indexMax])
			indexMax++
		}
		res += utils.MustParseInt(max)
	}

	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}
