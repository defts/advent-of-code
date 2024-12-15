package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func runeToDir(d rune) utils.Direction {
	dirMap := map[rune]utils.Direction{
		'^': utils.Up,
		'v': utils.Down,
		'<': utils.Left,
		'>': utils.Right,
	}
	return dirMap[d]
}

func part1(input arg) int {
	defer utils.Timer("part1")()

	puzzle := utils.NewGrid(input.puzzle)
	puzzle.Nodes[input.robot] = '.'

	res := 0

	moveObstacle := func(nextMove utils.Position, dir utils.Direction) {
		for f := nextMove.Move(dir, 1); ; f = f.Move(dir, 1) {
			if puzzle.Nodes[f] == '#' {
				break
			}
			if puzzle.Nodes[f] == '.' {
				input.robot = nextMove
				puzzle.Nodes[nextMove] = '.'
				puzzle.Nodes[f] = 'O'
				break
			}
		}
	}

	for _, i := range input.inputs {
		dir := runeToDir(i)
		nextMove := input.robot.Move(dir, 1)

		switch puzzle.Nodes[nextMove] {
		case '#':
			continue
		case '.':
			input.robot = nextMove
		case 'O':
			moveObstacle(nextMove, dir)
		}
	}

	for p, val := range puzzle.Nodes {
		if val == 'O' {
			res += p.Col + p.Row*100
		}
	}

	return res
}

func orderBoxPos(pos [2]utils.Position) [2]utils.Position {
	if pos[0].Col < pos[1].Col {
		return pos
	}
	return [2]utils.Position{pos[1], pos[0]}
}

func findMovedBoxes(puzzle utils.Grid[rune], pos [2]utils.Position, dir utils.Direction) [][2]utils.Position {
	newLoc1 := pos[0].Move(dir, 1)
	newLoc2 := pos[1].Move(dir, 1)

	list := [][2]utils.Position{}
	pos = orderBoxPos(pos)
	list = append(list, pos)

	newLoc := orderBoxPos([2]utils.Position{newLoc1, newLoc2})

	switch dir {
	case utils.Right:
		if puzzle.Nodes[newLoc[1]] == '[' {
			moved := [2]utils.Position{newLoc[1], newLoc[1].Move(dir, 1)}
			list = append(list, findMovedBoxes(puzzle, moved, dir)...)
		}
	case utils.Left:
		if puzzle.Nodes[newLoc[0]] == ']' {
			moved := [2]utils.Position{newLoc[0], newLoc[0].Move(dir, 1)}
			list = append(list, findMovedBoxes(puzzle, moved, dir)...)
		}
	case utils.Up, utils.Down:
		if puzzle.Nodes[newLoc[0]] == '[' {
			moved := [2]utils.Position{newLoc[0], newLoc[0].Move(utils.Right, 1)}
			list = append(list, findMovedBoxes(puzzle, moved, dir)...)
		}
		if puzzle.Nodes[newLoc[0]] == ']' {
			moved := [2]utils.Position{newLoc[0], newLoc[0].Move(utils.Left, 1)}
			list = append(list, findMovedBoxes(puzzle, moved, dir)...)
		}
		if puzzle.Nodes[newLoc[1]] == '[' {
			moved := [2]utils.Position{newLoc[1], newLoc[1].Move(utils.Right, 1)}
			list = append(list, findMovedBoxes(puzzle, moved, dir)...)
		}
	}
	return list
}

func moveBoxes(puzzle utils.Grid[rune], boxes [][2]utils.Position, dir utils.Direction) utils.Grid[rune] {
	currentLeft := map[utils.Position]bool{}
	currentRight := map[utils.Position]bool{}

	for _, box := range boxes {
		box = orderBoxPos(box)
		currentLeft[box[0]] = true
		currentRight[box[1]] = true
	}

	finalLeft := map[utils.Position]bool{}
	finalRight := map[utils.Position]bool{}

	for loc := range currentLeft {
		newLoc := loc.Move(dir, 1)
		finalLeft[newLoc] = true
	}
	for loc := range currentRight {
		newLoc := loc.Move(dir, 1)
		finalRight[newLoc] = true
	}

	for loc := range currentLeft {
		if _, ok := finalLeft[loc]; !ok {
			puzzle.Nodes[loc] = '.'
		}
	}
	for loc := range currentRight {
		if _, ok := finalRight[loc]; !ok {
			puzzle.Nodes[loc] = '.'
		}
	}

	for loc := range finalLeft {
		if _, ok := currentLeft[loc]; !ok {
			puzzle.Nodes[loc] = '['
		}
	}
	for loc := range finalRight {
		if _, ok := currentRight[loc]; !ok {
			puzzle.Nodes[loc] = ']'
		}
	}

	return puzzle
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

	puzzle := utils.NewGrid(scaledPuzzle)
	puzzle.Nodes[input.robot] = '.'

	boxCanMove := func(pos [2]utils.Position, dir utils.Direction) bool {
		if puzzle.Nodes[pos[1].Move(dir, 1)] == '#' || puzzle.Nodes[pos[0].Move(dir, 1)] == '#' {
			return false
		}
		return true
	}

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

		var boxPos [2]utils.Position
		if puzzle.Nodes[nextMove] == '[' {
			boxPos = [2]utils.Position{nextMove, nextMove.Move(utils.Right, 1)}
		} else {
			boxPos = [2]utils.Position{nextMove, nextMove.Move(utils.Left, 1)}
		}
		boxesToMove := findMovedBoxes(puzzle, boxPos, dir)
		canMove := true
		for _, box := range boxesToMove {
			if !boxCanMove(box, dir) {
				canMove = false
				break
			}
		}
		if !canMove {
			continue
		}
		puzzle = moveBoxes(puzzle, boxesToMove, dir)
		input.robot = nextMove
	}

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
