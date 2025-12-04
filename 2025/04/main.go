package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func part1(input arg) int {
	defer utils.Timer("part1")()
	res := 0

	for p, v := range input.grid.Nodes {
		if v == '@' {
			count := 0
			for _, dir := range utils.GetAllDirections() {
				if input.grid.Nodes[p.Move(dir, 1)] == '@' {
					count++
				}
			}
			if count < 4 {
				res++
			}
		}
	}

	return res
}

func part2(input arg) int {
	defer utils.Timer("part2")()
	res := 0
	rollsremoved := 1
	for rollsremoved != 0 {
		rollsToRemove := []utils.Position{}
		for p, v := range input.grid.Nodes {
			if v == '@' {
				count := 0
				for _, dir := range utils.GetAllDirections() {
					if input.grid.Nodes[p.Move(dir, 1)] == '@' {
						count++
					}
				}
				if count < 4 {
					rollsToRemove = append(rollsToRemove, p)
				}
			}
		}
		for _, p := range rollsToRemove {
			input.grid.Nodes[p] = '.'
		}
		rollsremoved = len(rollsToRemove)
		res += rollsremoved
	}

	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}
