package utils

type Grid[T any] struct {
	Nodes map[Position]T
	Rows  int
	Cols  int
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
		Rows:  len(nodes),
		Cols:  len(nodes[0]),
	}
}

func (g Grid[T]) PositionInBounds(p Position) bool {
	_, ok := g.Nodes[p]
	return ok
}

func (g Grid[T]) Print() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			if node, ok := g.Nodes[Position{Row: i, Col: j}]; ok {
				// test cast node to string
				if str, ok := any(node).(rune); ok {
					print(string(str))
				} else {
					print(node)
				}
			}
		}
		println()
	}
}
