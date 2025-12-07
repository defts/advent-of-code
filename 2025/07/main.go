package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func part1(input arg) int {
	defer utils.Timer("part1")()
	res := 0
	start := utils.Position{}

	for p, v := range input.grid.Nodes {
		if v == 'S' {
			start = p
			break
		}
	}
	visited := map[utils.Position]bool{}

	var beamFunc func(utils.Position)
	beamFunc = func(p utils.Position) {
		if visited[p] {
			return
		}
		visited[p] = true

		p = p.Move(utils.Down, 1)
		if !input.grid.PositionInBounds(p) {
			return
		}
		if input.grid.Nodes[p] == '^' {
			res++
			beamFunc(p.Move(utils.Left, 1))
			beamFunc(p.Move(utils.Right, 1))
			return
		}
		beamFunc(p)
	}

	beamFunc(start)

	return res
}

func part2(input arg) int {
	defer utils.Timer("part2")()
	res := 0

	start := utils.Position{}

	for p, v := range input.grid.Nodes {
		if v == 'S' {
			start = p
			break
		}
	}
	cache := map[utils.Position]int{}

	var beamFunc func(utils.Position) int
	beamFunc = func(p utils.Position) (n int) {
		if n, ok := cache[p]; ok {
			return n
		}
		defer func() { cache[p] = n }()

		p = p.Move(utils.Down, 1)
		if !input.grid.PositionInBounds(p) {
			return 1
		}
		if input.grid.Nodes[p] == '^' {
			return beamFunc(p.Move(utils.Left, 1)) + beamFunc(p.Move(utils.Right, 1))
		}
		return beamFunc(p)
	}

	res = beamFunc(start)

	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}
