package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func calculateRegion(grid [][]rune, start utils.Position, visited map[utils.Position]bool) (int, int) {
	stack := []utils.Position{start}
	area, perimeter := 0, 0
	regionType := grid[start.Row][start.Col]

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
			if neighbor.Row < 0 || neighbor.Col < 0 || neighbor.Row >= len(grid) || neighbor.Col >= len(grid[0]) {
				perimeter++ // Edge of the grid
			} else if grid[neighbor.Row][neighbor.Col] != regionType {
				perimeter++ // Neighbor of a different type
			} else if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
	return area, perimeter
}

func part1(arg [][]rune) int {
	defer utils.Timer("part1")()

	res := 0
	visited := make(map[utils.Position]bool)

	for row := 0; row < len(arg); row++ {
		for col := 0; col < len(arg[0]); col++ {
			pos := utils.Position{Row: row, Col: col}
			if !visited[pos] {
				area, perimeter := calculateRegion(arg, pos, visited)
				res += area * perimeter

			}
		}
	}
	return res
}

func part2(arg [][]rune) int {
	defer utils.Timer("part2")()
	res := 0

	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(arg))
	fmt.Println("part2: ", part2(arg))
}
