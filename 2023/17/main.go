package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
	priorityqueue "github.com/emirpasic/gods/queues/priorityqueue"
)

type node struct {
	pos       utils.Position
	direction utils.Direction
	straight  int
}

type nodeWithCost struct {
	node
	history []utils.Position
	cost    int
}

func run(maze [][]int, mindist, maxdist int) int {
	// la fin, c'est la derniere case en bas à droite
	end := utils.Position{len(maze) - 1, len(maze[0]) - 1}

	queue := priorityqueue.NewWith(func(a, b any) int {
		p1 := a.(nodeWithCost).cost
		p2 := b.(nodeWithCost).cost
		return p1 - p2
	})

	// comme on commence au 0,0, on ajoute 0,1 vers la droite et 1,0 vers le bas
	queue.Enqueue(nodeWithCost{node{utils.Position{0, 1}, utils.Right, 1}, []utils.Position{utils.Position{0, 0}}, 0})
	queue.Enqueue(nodeWithCost{node{utils.Position{1, 0}, utils.Down, 1}, []utils.Position{utils.Position{0, 0}}, 0})
	cache := make(map[node]int)

	for !queue.Empty() {
		_n, _ := queue.Dequeue()
		n := _n.(nodeWithCost)

		if !n.pos.InBounds(maze) {
			continue
		}

		heat := maze[n.pos.Row][n.pos.Col] + n.cost

		// on ne veut sortir que si la derniere ligne droite est dans supperieur à la borne mini imposée
		if n.pos == end && n.straight >= mindist {
			return heat
		}

		// si on est deja passé par cette case, on ne veut pas y repasser si le cout est plus élevé
		if v, exists := cache[n.node]; exists {
			if v <= heat {
				continue
			}
		}
		cache[n.node] = heat

		// il faut attendre de continuer tout droit par le minimum imposé avant de tourner
		if n.straight >= mindist {
			// on peut tourner à droite ou à gauche, le demi tour n'est pas autorisé et le tout droit est géré apres.
			left := n.direction.Turn(utils.Left)
			queue.Enqueue(nodeWithCost{node{n.pos.Move(left, 1), left, 1}, append(n.history, n.pos), heat})

			right := n.direction.Turn(utils.Right)
			queue.Enqueue(nodeWithCost{node{n.pos.Move(right, 1), right, 1}, append(n.history, n.pos), heat})
		}

		// on ne veut pas depasser la borne max de ligne droite
		if n.straight < maxdist {
			queue.Enqueue(nodeWithCost{node{n.pos.Move(n.direction, 1), n.direction, n.straight + 1}, append(n.history, n.pos), heat})
		}
	}
	return -1 // Aucun chemin trouvé
}

func main() {
	loadFile("input.txt")
	fmt.Println("Part1 : ", run(maze, 1, 3))
	fmt.Println("Part2 : ", run(maze, 4, 10))

}
