package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func part1(grid [][]rune) int {
	defer utils.Timer("part1")()

	positions := map[rune][]utils.Position{}
	antinodes := map[utils.Position]bool{}

	g := utils.NewGrid(grid)

	for p, pos := range g.Nodes {
		if pos != '.' {
			positions[pos] = append(positions[pos], p)
		}
	}

	for _, pos := range positions {
		for i := range pos {
			for j := range pos {
				if i != j {
					rowDiff := pos[i].RowDiff(pos[j])
					colDiff := pos[i].ColDiff(pos[j])
					antinode_1 := utils.Position{Row: pos[i].Row + rowDiff, Col: pos[i].Col + colDiff}
					antinode_2 := utils.Position{Row: pos[j].Row - rowDiff, Col: pos[j].Col - colDiff}

					if g.PositionInBounds(antinode_1) {
						antinodes[antinode_1] = true
					}
					if g.PositionInBounds(antinode_2) {
						antinodes[antinode_2] = true
					}
				}
			}
		}
	}

	return len(antinodes)
}

func part2(grid [][]rune) int {
	defer utils.Timer("part1")()

	positions := map[rune][]utils.Position{}
	antinodes := map[utils.Position]bool{}

	g := utils.NewGrid(grid)

	for p, pos := range g.Nodes {
		if pos != '.' {
			positions[pos] = append(positions[pos], p)

		}
	}

	generateAntinodes := func(start utils.Position, rowStep, colStep int) {
		current := start
		for {
			current = utils.Position{Row: current.Row + rowStep, Col: current.Col + colStep}
			if g.PositionInBounds(current) {
				antinodes[current] = true
			} else {
				break
			}
		}
	}

	// Process each group of nodes
	for _, pos := range positions {
		for i := range pos {
			for j := range pos {
				if i != j {
					// Add direct nodes as antinodes
					antinodes[pos[i]] = true
					antinodes[pos[j]] = true

					// Calculate differences
					rowDiff := pos[i].RowDiff(pos[j])
					colDiff := pos[i].ColDiff(pos[j])

					// Generate antinodes in both directions
					generateAntinodes(pos[i], rowDiff, colDiff)
					generateAntinodes(pos[j], -rowDiff, -colDiff)
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(grid))
	fmt.Println("part2: ", part2(grid))
}
