package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func part1(grid [][]rune) int {
	defer utils.Timer("part1")()
	g := utils.NewGrid(grid)
	var guard utils.Position
	guardDir := utils.Up
	for p, v := range g.Nodes {
		if v == '^' {
			guard = p
			break
		}
	}
	visited := make(map[utils.Position]bool)
	visited[guard] = true
	for {
		nextMove := guard.Move(guardDir, 1)
		if _, ok := g.Nodes[nextMove]; !ok {
			break
		}
		if g.Nodes[nextMove] == '#' {
			guardDir = guardDir.Turn(utils.Right)
			nextMove = guard.Move(guardDir, 1)
		}
		guard = nextMove
		visited[guard] = true
	}
	return len(visited)
}

func part2(grid [][]rune) int {
	defer utils.Timer("part2")()
	g := utils.NewGrid(grid)
	var guard utils.Position
	guardDir := utils.Up
	for p, v := range g.Nodes {
		if v == '^' {
			guard = p
			break
		}
	}
	startPosition := guard
	visited := make(map[utils.Position]bool)
	visited[guard] = true
	// on refait la partie 1
	for {
		nextMove := guard.Move(guardDir, 1)
		if _, ok := g.Nodes[nextMove]; !ok {
			break
		}
		if g.Nodes[nextMove] == '#' {
			guardDir = guardDir.Turn(utils.Right)
			nextMove = guard.Move(guardDir, 1)
		}
		guard = nextMove
		visited[guard] = true
	}

	found := 0
	// on remplace la position du garde
	g.Nodes[startPosition] = '.'

	// petite struct pour garder les loc, à voir pour mettre ca dans utils
	// ca permet de savoir si on boucle ou pas entre une postion et une direction
	type location struct {
		pos utils.Position
		dir utils.Direction
	}

	// pour chaque endroit ou passe le garde, on va le remplacer par un mur et voir si on boucle
	for wall := range visited {
		old_char := g.Nodes[wall] // on garde l'ancien caractère pour le remettre après
		g.Nodes[wall] = '#'

		// on part de la position de départ
		loc := location{
			pos: startPosition,
			dir: utils.Up,
		}

		v := map[utils.Position]bool{}
		v[loc.pos] = true
		loop := map[location]bool{}

		for {
			new_pos := loc.pos.Move(loc.dir, 1)
			if _, ok := g.Nodes[new_pos]; !ok {
				break
			}

			// on va se promener
			for {
				if g.Nodes[new_pos] == '#' {
					loc.dir = loc.dir.Turn(utils.Right)
					new_pos = loc.pos.Move(loc.dir, 1)
				}
				// If the new position is open space, move there
				if g.Nodes[new_pos] == '.' {
					loc.pos = new_pos
					break
				}
			}

			// on est dans une boucle ?
			if _, ok := loop[loc]; ok {
				found++
				break
			} else {
				loop[loc] = true
			}
			// on recommence
			v[loc.pos] = true
		}

		// on remet le caractère d'origine
		g.Nodes[wall] = old_char
	}

	return found
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(grid))
	fmt.Println("part2: ", part2(grid))
}
