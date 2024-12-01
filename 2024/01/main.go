package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/defts/advent-of-code/utils"
)

func part1(rowOne []int, rowTwo []int) int {
	defer utils.Timer("part1")()
	slices.Sort(rowOne)
	slices.Sort(rowTwo)
	res := 0
	for i := 0; i < len(rowOne); i++ {
		res += int(math.Abs(float64(rowOne[i] - rowTwo[i])))
	}
	return res
}

func part2(rowOne []int, rowTwo []int) int {
	defer utils.Timer("part2")()
	res := 0
	for i := 0; i < len(rowOne); i++ {
		copy := 0
		for j := 0; j < len(rowTwo); j++ {
			if rowOne[i] == rowTwo[j] {
				copy++
			}
		}
		res += (rowOne[i] * copy)
	}
	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(rowOne, rowTwo))
	fmt.Println("part2: ", part2(rowOne, rowTwo))
}
