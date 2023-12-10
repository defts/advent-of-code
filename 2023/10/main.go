package main

import (
	"fmt"
	"slices"
)

type pos struct {
	row int
	col int
}

func validInBounds(sketch [][]rune, p pos) bool {
	return p.row >= 0 && p.row < len(sketch) && p.col >= 0 && p.col < len(sketch[0])
}

func neighbors(sketch [][]rune, p pos) []pos {
	switch sketch[p.row][p.col] {
	case '|':
		if validInBounds(sketch, pos{p.row - 1, p.col}) && validInBounds(sketch, pos{p.row + 1, p.col}) {
			return []pos{pos{p.row - 1, p.col}, pos{p.row + 1, p.col}}
		}
	case '-':
		if validInBounds(sketch, pos{p.row, p.col - 1}) && validInBounds(sketch, pos{p.row, p.col + 1}) {
			return []pos{pos{p.row, p.col - 1}, pos{p.row, p.col + 1}}
		}
	case 'L':
		if validInBounds(sketch, pos{p.row - 1, p.col}) && validInBounds(sketch, pos{p.row, p.col + 1}) {
			return []pos{pos{p.row - 1, p.col}, pos{p.row, p.col + 1}}
		}
	case 'J':
		if validInBounds(sketch, pos{p.row - 1, p.col}) && validInBounds(sketch, pos{p.row, p.col - 1}) {
			return []pos{pos{p.row - 1, p.col}, pos{p.row, p.col - 1}}
		}
	case '7':
		if validInBounds(sketch, pos{p.row + 1, p.col}) && validInBounds(sketch, pos{p.row, p.col - 1}) {
			return []pos{pos{p.row + 1, p.col}, pos{p.row, p.col - 1}}
		}
	case 'F':
		if validInBounds(sketch, pos{p.row + 1, p.col}) && validInBounds(sketch, pos{p.row, p.col + 1}) {
			return []pos{pos{p.row + 1, p.col}, pos{p.row, p.col + 1}}
		}
	case 'S':
		// start can have 4 neighbors
		n := make([]pos, 0)
		if validInBounds(sketch, pos{p.row + 1, p.col}) {
			n = append(n, pos{p.row + 1, p.col})
		}
		if validInBounds(sketch, pos{p.row - 1, p.col}) {
			n = append(n, pos{p.row - 1, p.col})
		}
		if validInBounds(sketch, pos{p.row, p.col + 1}) {
			n = append(n, pos{p.row, p.col + 1})
		}
		if validInBounds(sketch, pos{p.row, p.col - 1}) {
			n = append(n, pos{p.row, p.col - 1})
		}
		return n
	}
	return nil

}

func findStart(sketch [][]rune) pos {
	for r := 0; r < len(sketch); r++ {
		for c := 0; c < len(sketch[0]); c++ {
			if sketch[r][c] == 'S' {
				return pos{r, c}
			}
		}
	}
	return pos{-1, -1}
}

// start at start and follow neighbors by starting at top bottom left right and following neighbors until you reach start again
func findMainLoop(sketch [][]rune, start pos) []pos {
	visited := make(map[pos]bool)
	visited[start] = true
	path := []pos{start}
	for _, n := range neighbors(sketch, start) { // start neighbors
		path := []pos{start}
		previous := start
		current := n
		for {
			if current == start {
				path = append(path, current)
				return path
			}
			nexts := make([]pos, 0)
			for _, n := range neighbors(sketch, current) {
				if n != previous {
					nexts = append(nexts, n)
				}
			}

			if len(nexts) != 1 {
				break
			}
			next := nexts[0]

			if !visited[current] {
				visited[current] = true
				path = append(path, current)
			} else {
				if n == start {
					return path
				}
				break
			}
			previous = current
			current = next
		}
	}

	return path
}

func printMazePath(ketch [][]rune, loop []pos, enclosed []pos) {
	for r := 0; r < len(ketch); r++ {
		for c := 0; c < len(ketch[r]); c++ {
			if slices.Contains(loop, pos{r, c}) {
				fmt.Print(string(sketch[r][c]))
			} else {
				if slices.Contains(enclosed, pos{r, c}) {
					fmt.Print("I")
				} else {
					fmt.Print("#")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	loadFile("input.txt")

	start := findStart(sketch)
	fmt.Println(start)
	loop := findMainLoop(sketch, start)
	fmt.Println("part1: ", len(loop)/2)

	enclosed := make([]pos, 0)
	for r := 0; r < len(sketch); r++ {
		inside := false
		for c := 0; c < len(sketch[r]); c++ {
			if slices.Contains(loop, pos{r, c}) {
				if sketch[r][c] == '|' || sketch[r][c] == 'L' || sketch[r][c] == 'J' {
					inside = !inside
				}
				continue
			}
			if inside {
				enclosed = append(enclosed, pos{r, c})
			}
		}
	}
	fmt.Println("part2: ", len(enclosed))
	printMazePath(sketch, loop, enclosed)

}
