package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func part1(arg [][]rune) int {
	defer utils.Timer("part1")()

	res := 0

	starts := []utils.Position{}

	grid := utils.NewGrid(arg)

	for g, v := range grid.Nodes {
		if v == '0' {
			starts = append(starts, g)
		}
	}

	explore := func(start utils.Position) int {
		queue := []utils.Position{start}
		visited := map[utils.Position]bool{start: true}
		count := 0

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			// Si on atteint un '9', on compte
			if grid.Nodes[current] == '9' {
				count++
				continue
			}

			// Vérifier les voisins
			for _, dir := range utils.GetAllCartesianDirection() {
				neighbor := current.Move(dir, 1)

				// Vérifier si le voisin est dans les limites et non visité
				if _, exists := grid.Nodes[neighbor]; !exists {
					continue
				}
				if visited[neighbor] {
					continue
				}

				// Vérifier la progression de la hauteur
				if grid.Nodes[neighbor] == grid.Nodes[current]+1 {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}

		return count
	}

	// Calculer le score pour chaque position de départ
	for _, start := range starts {
		res += explore(start)
	}

	return res
}

func part2(arg [][]rune) int {
	defer utils.Timer("part2")()

	res := 0

	// Liste des positions de départ (les '0')
	starts := []utils.Position{}

	// Grille pour navigation
	grid := utils.NewGrid(arg)

	// Trouver toutes les positions avec '0' comme valeur
	for pos, val := range grid.Nodes {
		if val == '0' {
			starts = append(starts, pos)
		}
	}

	var explore func(current utils.Position, visited map[utils.Position]bool) int
	explore = func(current utils.Position, visited map[utils.Position]bool) int {
		// Si on atteint un '9', on a trouvé un chemin
		if grid.Nodes[current] == '9' {
			return 1
		}

		paths := 0

		// Vérifier les voisins
		for _, dir := range utils.GetAllCartesianDirection() {
			neighbor := current.Move(dir, 1)

			// Vérifier si le voisin est dans les limites et non visité
			if _, exists := grid.Nodes[neighbor]; !exists {
				continue
			}
			if visited[neighbor] {
				continue
			}

			// Vérifier la progression de la hauteur
			if grid.Nodes[neighbor] == grid.Nodes[current]+1 {
				// Marquer comme visité pour ce chemin
				visited[neighbor] = true
				paths += explore(neighbor, visited)
				// Défaire la visite pour d'autres chemins
				delete(visited, neighbor)
			}
		}

		return paths
	}

	// Calculer le rating pour chaque position de départ
	for _, start := range starts {
		visited := map[utils.Position]bool{start: true}
		res += explore(start, visited)
	}

	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(arg))
	fmt.Println("part2: ", part2(arg))
}
