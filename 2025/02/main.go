package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

func part1(input arg) int {
	defer utils.Timer("part1")()
	res := 0

	for _, ids := range input.ids {
		for i := ids.idOne; i <= ids.idTwo; i++ {
			curr := strconv.Itoa(i)
			if curr[:len(curr)/2] == curr[len(curr)/2:] {
				res += i
			}
		}
	}

	return res
}

func part2(input arg) int {
	defer utils.Timer("part2")()
	res := 0
	for _, ids := range input.ids {
		for i := ids.idOne; i <= ids.idTwo; i++ {
			curr := strconv.Itoa(i)
			if strings.Contains((curr + curr)[1:len(curr+curr)-1], curr) {
				res += i
			}
		}
	}

	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}
