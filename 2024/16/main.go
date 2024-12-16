package main

import (
	"container/heap"
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

type State struct {
	pos       utils.Position
	direction utils.Direction
	score     int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].score < pq[j].score }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(State)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func part1(input arg) int {
	defer utils.Timer("part1")()

	maze := utils.NewGrid(input.maze)
	maze.Nodes[input.start] = '.'
	maze.Nodes[input.end] = '.'

	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, State{input.start, utils.Right, 0})

	visited := make(map[struct {
		pos       utils.Position
		direction utils.Direction
	}]bool)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)

		// Check if we reached the end
		if current.pos == input.end {
			return current.score
		}

		// Avoid revisiting same position + direction
		visitedKey := struct {
			pos       utils.Position
			direction utils.Direction
		}{current.pos, current.direction}

		if visited[visitedKey] {
			continue
		}
		visited[visitedKey] = true

		// Move forward
		nextPos := current.pos.Move(current.direction, 1)
		if maze.PositionInBounds(nextPos) && maze.Nodes[nextPos] != '#' {
			heap.Push(pq, State{nextPos, current.direction, current.score + 1})
		}

		// Rotate left and right
		heap.Push(pq, State{current.pos, current.direction.Turn(utils.Right), current.score + 1000})
		heap.Push(pq, State{current.pos, current.direction.Turn(utils.Left), current.score + 1000})
	}

	return -1
}

func part2(input arg) int {
	defer utils.Timer("part2")()
	res := 0

	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}
