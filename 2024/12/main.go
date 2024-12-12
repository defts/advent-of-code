package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func part1(arg [][]rune) int {
	defer utils.Timer("part1")()

	res := 0
	visited := make(map[utils.Position]bool)

	board := utils.NewGrid(arg)

	for pos, v := range board.Nodes {
		if visited[pos] {
			continue
		}

		stack := []utils.Position{pos}
		area, perimeter := 0, 0
		regionType := v

		for len(stack) > 0 {
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if visited[cell] {
				continue
			}
			visited[cell] = true
			area++

			for _, dir := range utils.GetAllCartesianDirection() {
				neighbor := cell.Move(dir, 1)
				if !board.PositionInBounds(neighbor) {
					perimeter++ // Edge of the grid
				} else if arg[neighbor.Row][neighbor.Col] != regionType {
					perimeter++ // Neighbor of a different type
				} else if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}

		res += area * perimeter
	}
	return res
}

func part2(input [][]rune) int {
	board := utils.NewGrid(input)

	seen := make(map[utils.Position]bool)
	res := 0

	for pos := range board.Nodes {
		if seen[pos] {
			continue
		}
		seen[pos] = true

		area := 1
		sides := 0

		// Utilisation d'une pile pour le parcours
		stack := []utils.Position{pos}

		for len(stack) > 0 {
			// Pop depuis la fin de la pile
			n := len(stack) - 1
			loc := stack[n]
			stack = stack[:n]

			for _, dir := range utils.GetAllCartesianDirection() {
				newLoc := loc.Move(dir, 1)
				if board.Nodes[newLoc] != board.Nodes[loc] {
					// Vérification des coins
					check1 := loc.Move(dir.Turn(utils.Left), 1)
					check2 := check1.Move(dir, 1)

					if board.Nodes[check1] != board.Nodes[loc] || board.Nodes[check2] == board.Nodes[loc] {
						sides++
					}
				} else if !seen[newLoc] {
					seen[newLoc] = true
					stack = append(stack, newLoc)
					area++
				}
			}
		}
		res += area * sides
	}
	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(arg))
	fmt.Println("part2: ", part2(arg))
}
