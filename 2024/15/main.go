package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func runeToDir(d rune) utils.Direction {
	switch d {
	case '^':
		return utils.Up
	case 'v':
		return utils.Down
	case '<':
		return utils.Left
	case '>':
		return utils.Right

	}
	return utils.NotSet
}

func part1(input arg) int {
	defer utils.Timer("part1")()

	puzzle := utils.NewGrid(input.puzzle)
	puzzle.Nodes[input.robot] = '.'

	res := 0

	for _, i := range input.inputs {
		dir := runeToDir(i)
		nextMove := input.robot.Move(dir, 1)
		if puzzle.Nodes[nextMove] == '#' {
			continue
		}
		if puzzle.Nodes[nextMove] == '.' {
			input.robot = nextMove
			continue
		}
		if puzzle.Nodes[nextMove] == 'O' {
			f := nextMove.Move(dir, 1)
			for {
				if puzzle.Nodes[f] == '#' {
					break
				}
				if puzzle.Nodes[f] == '.' {
					input.robot = nextMove
					puzzle.Nodes[nextMove] = '.'
					puzzle.Nodes[f] = 'O'
					break
				}
				f = f.Move(dir, 1)
			}
		}

	}

	for p := range puzzle.Nodes {
		if puzzle.Nodes[p] == 'O' {
			res += p.Col + p.Row*100
		}
	}

	return res
}

func part2(input arg) int {
	defer utils.Timer("part2")()
	res := 0

	scaledPuzzle := make([][]rune, len(input.puzzle))

	for i := range input.puzzle {
		scaledPuzzle[i] = make([]rune, len(input.puzzle[i])*2)
		for j := range input.puzzle[i] {
			switch input.puzzle[i][j] {
			case '.':
				scaledPuzzle[i][j*2] = '.'
				scaledPuzzle[i][j*2+1] = '.'
			case '#':
				scaledPuzzle[i][j*2] = '#'
				scaledPuzzle[i][j*2+1] = '#'
			case 'O':
				scaledPuzzle[i][j*2] = '['
				scaledPuzzle[i][j*2+1] = ']'
			}
			if i == input.robot.Row && j == input.robot.Col {
				scaledPuzzle[i][j*2] = '.'
				scaledPuzzle[i][j*2+1] = '.'
			}
		}
	}

	input.robot.Col *= 2

	for _, i := range input.inputs {
		dir := runeToDir(i)
		nextMove := input.robot.Move(dir, 1)
		if scaledPuzzle[nextMove.Row][nextMove.Col] == '#' {
			continue
		}
		if scaledPuzzle[nextMove.Row][nextMove.Col] == '.' {
			input.robot = nextMove
			continue
		}
		// gerer les []
	}

	puzzle := utils.NewGrid(scaledPuzzle)
	puzzle.Nodes[input.robot] = '.'

	for p := range puzzle.Nodes {
		if puzzle.Nodes[p] == '[' {
			res += p.Col + p.Row*100
		}
	}

	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}
