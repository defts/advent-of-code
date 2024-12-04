package utils

type Grid[T any] struct {
	Nodes map[Position]T
}

func NewGrid[T any](nodes [][]T) Grid[T] {
	gridMap := make(map[Position]T, len(nodes))
	for x, node := range nodes {
		for y, node := range node {
			point := Position{
				Row: x,
				Col: y,
			}
			gridMap[point] = node
		}
	}

	return Grid[T]{
		Nodes: gridMap,
	}
}

func (g Grid[T]) PositionInBounds(p Position) bool {
	return p.Row >= 0 && p.Col >= 0 && p.Row < len(g.Nodes) && p.Col < len(g.Nodes)
}