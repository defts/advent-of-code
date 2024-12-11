package main

import (
	"fmt"
	"strconv"

	"github.com/defts/advent-of-code/utils"
)

func blink(stoneCounts map[int]int) map[int]int {
	newCounts := map[int]int{}
	for stone, count := range stoneCounts {
		switch {
		case stone == 0:
			newCounts[1] += count
		case len(strconv.Itoa(stone))%2 == 0:
			str := strconv.Itoa(stone)
			mid := len(str) / 2
			left, _ := strconv.Atoi(str[:mid])
			right, _ := strconv.Atoi(str[mid:])
			newCounts[left] += count
			newCounts[right] += count
		default:
			newCounts[stone*2024] += count
		}
	}
	return newCounts
}

func part1(arg []int) int {
	defer utils.Timer("part1")()

	stoneCounts := map[int]int{}
	for _, a := range arg {
		stoneCounts[a]++
	}

	for i := 0; i < 25; i++ {
		stoneCounts = blink(stoneCounts)
	}

	res := 0
	for _, count := range stoneCounts {
		res += count
	}
	return res
}

func part2(arg []int) int {
	defer utils.Timer("part2")()

	stoneCounts := map[int]int{}
	for _, a := range arg {
		stoneCounts[a]++
	}

	for i := 0; i < 75; i++ {
		stoneCounts = blink(stoneCounts)
	}

	res := 0
	for _, count := range stoneCounts {
		res += count
	}
	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(arg))
	fmt.Println("part2: ", part2(arg))
}
