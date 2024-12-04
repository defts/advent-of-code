package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

const (
	XMAS = "XMAS"
)

func findXMAS(grid utils.Grid[rune], pos utils.Position, direction utils.Direction) bool {
	for _, r := range XMAS {
		if grid.Nodes[pos] != r {
			return false
		}
		pos = pos.Move(direction, 1)
	}
	return true
}

func part1(grid utils.Grid[rune]) int {
	defer utils.Timer("part1")()
	res := 0
	for pos := range grid.Nodes {
		for _, dir := range utils.GetAllDirections() {
			if findXMAS(grid, pos, dir) {
				res += 1
			}
		}

	}
	return res
}

func findCross(grid utils.Grid[rune], pos utils.Position) bool {
	if grid.Nodes[pos] != 'A' { // the center
		return false
	}
	ul := grid.Nodes[pos.Move(utils.UpLeft, 1)]
	dr := grid.Nodes[pos.Move(utils.DownRight, 1)]
	ur := grid.Nodes[pos.Move(utils.UpRight, 1)]
	dl := grid.Nodes[pos.Move(utils.DownLeft, 1)]
	return ((ul == 'M' && dr == 'S') || (ul == 'S' && dr == 'M')) && ((ur == 'M' && dl == 'S') || (ur == 'S' && dl == 'M'))
}

func part2(grid utils.Grid[rune]) int {
	defer utils.Timer("part2")()
	res := 0
	for pos := range grid.Nodes {
		if findCross(grid, pos) {
			res += 1
		}
	}

	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(grid))
	fmt.Println("part2: ", part2(grid))
}
