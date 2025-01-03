package utils

import "math"

type Position struct {
	Row int
	Col int
}

func (p Position) InBounds(t [][]int) bool {
	return p.Row >= 0 && p.Col >= 0 && p.Row < len(t) && p.Col < len(t)
}

func (p Position) Delta(row, col int) Position {
	return Position{
		Row: p.Row + row,
		Col: p.Col + col,
	}
}

func (p Position) Move(direction Direction, moves int) Position {
	switch direction {
	case Up:
		return p.Delta(-moves, 0)
	case UpRight:
		return p.Delta(-moves, moves)
	case Right:
		return p.Delta(0, moves)
	case DownRight:
		return p.Delta(moves, moves)
	case Down:
		return p.Delta(moves, 0)
	case DownLeft:
		return p.Delta(moves, -moves)
	case Left:
		return p.Delta(0, -moves)
	case UpLeft:
		return p.Delta(-moves, -moves)
	}

	return p
}

func (p Position) RowDiff(p2 Position) int {
	return p.Row - p2.Row
}

func (p Position) ColDiff(p2 Position) int {
	return p.Col - p2.Col
}

func (p Position) Manhattan(p2 Position) int {
	return int(math.Abs(float64(p.Row-p2.Row)) + math.Abs(float64(p.Col-p2.Col)))
}
