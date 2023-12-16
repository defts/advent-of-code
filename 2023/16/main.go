package main

import (
	"fmt"
)

type pos struct {
	row int
	col int
}

type direction int

// 0 = up
// 1 = right
// 2 = down
// 3 = left

type beam struct {
	pos pos
	d   direction
}

func (b beam) nextMove(contraption [][]rune) []beam {
	currentPos := contraption[b.pos.row][b.pos.col]

	switch currentPos {
	case '.':
		switch b.d {
		case 0:
			return []beam{beam{pos{b.pos.row - 1, b.pos.col}, b.d}}
		case 1:
			return []beam{beam{pos{b.pos.row, b.pos.col + 1}, b.d}}
		case 2:
			return []beam{beam{pos{b.pos.row + 1, b.pos.col}, b.d}}
		case 3:
			return []beam{beam{pos{b.pos.row, b.pos.col - 1}, b.d}}
		}
	case '|':
		switch b.d {
		case 0:
			return []beam{beam{pos{b.pos.row - 1, b.pos.col}, b.d}}
		case 1, 3:
			return []beam{beam{pos{b.pos.row - 1, b.pos.col}, 0}, beam{pos{b.pos.row + 1, b.pos.col}, 2}}
		case 2:
			return []beam{beam{pos{b.pos.row + 1, b.pos.col}, b.d}}
		}
	case '-':
		switch b.d {
		case 0, 2:
			return []beam{beam{pos{b.pos.row, b.pos.col - 1}, 3}, beam{pos{b.pos.row, b.pos.col + 1}, 1}}

		case 1:
			return []beam{beam{pos{b.pos.row, b.pos.col + 1}, b.d}}
		case 3:
			return []beam{beam{pos{b.pos.row, b.pos.col - 1}, b.d}}
		}
	case '/':
		switch b.d {
		case 0:
			return []beam{beam{pos{b.pos.row, b.pos.col + 1}, 1}}
		case 1:
			return []beam{beam{pos{b.pos.row - 1, b.pos.col}, 0}}
		case 2:
			return []beam{beam{pos{b.pos.row, b.pos.col - 1}, 3}}
		case 3:
			return []beam{beam{pos{b.pos.row + 1, b.pos.col}, 2}}
		}
	case '\\':
		switch b.d {
		case 0:
			return []beam{beam{pos{b.pos.row, b.pos.col - 1}, 3}}
		case 1:
			return []beam{beam{pos{b.pos.row + 1, b.pos.col}, 2}}
		case 2:
			return []beam{beam{pos{b.pos.row, b.pos.col + 1}, 1}}
		case 3:
			return []beam{beam{pos{b.pos.row - 1, b.pos.col}, 0}}
		}
	}

	return []beam{}
}

func computeContraption(contraption [][]rune, start beam) map[beam]bool {
	energized := make(map[beam]bool)
	beams := []beam{start}

	for len(beams) != 0 {
		curr := beams[0]
		beams = beams[1:]
		if curr.pos.row >= len(contraption) || curr.pos.row < 0 || curr.pos.col >= len(contraption[0]) || curr.pos.col < 0 {
			continue
		}

		if energized[curr] {
			continue
		}

		energized[curr] = true

		beams = append(beams, curr.nextMove(contraption)...)
	}
	return energized
}

func getEnergizedTiles(energized map[beam]bool) map[pos]bool {
	seen := make(map[pos]bool)
	for k := range energized {
		seen[k.pos] = true
	}
	return seen
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", len(getEnergizedTiles(computeContraption(contraption, beam{pos{0, 0}, 1}))))

	// part2
	max := 0
	for row := 0; row < len(contraption); row++ {
		s := len(getEnergizedTiles(computeContraption(contraption, beam{pos{row, 0}, 1})))
		if s > max {
			max = s
		}

		s2 := len(getEnergizedTiles(computeContraption(contraption, beam{pos{row, len(contraption) - 1}, 3})))
		if s2 > max {
			max = s2
		}
	}
	for col := 0; col < len(contraption[0]); col++ {
		s := len(getEnergizedTiles(computeContraption(contraption, beam{pos{0, col}, 2})))
		if s > max {
			max = s
		}

		s2 := len(getEnergizedTiles(computeContraption(contraption, beam{pos{len(contraption) - 1, col}, 0})))
		if s2 > max {
			max = s2
		}
	}

	fmt.Println("part2: ", max)

}
