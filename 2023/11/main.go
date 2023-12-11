package main

import (
	"fmt"
	"slices"
)

type pos struct {
	row int
	col int
}

func expansionOfTheUniverse(universe [][]rune, scale int) (rowsWithoutGalaxy, colsWithoutGalaxy []int) {
	// newUniverse := [][]rune{}
	for row := 0; row < len(universe); row++ {
		found := false
		for col := 0; col < len(universe[row]); col++ {
			if universe[row][col] == '#' {
				found = true
				break
			}
		}
		if !found {
			rowsWithoutGalaxy = append(rowsWithoutGalaxy, row)
		}
	}

	for col := 0; col < len(universe[0]); col++ {
		found := false
		for row := 0; row < len(universe); row++ {
			if universe[row][col] == '#' {
				found = true
				break
			}
		}
		if !found {
			colsWithoutGalaxy = append(colsWithoutGalaxy, col)
		}
	}

	return rowsWithoutGalaxy, colsWithoutGalaxy

	// // exapnd universe
	// for row := 0; row < len(universe); row++ {
	// 	if slices.Contains(rowsWithoutGalaxy, row) {
	// 		for i := 0; i < scale; i++ {
	// 			newUniverse = append(newUniverse, universe[row])
	// 		}
	// 	} else {
	// 		newUniverse = append(newUniverse, universe[row])
	// 	}
	// }
	// newUniverseExpanded := make([][]rune, len(newUniverse))
	// nbCols := len(newUniverse[0]) + (len(colsWithoutGalaxy) * scale)
	// for row := range newUniverse {
	// 	newUniverseExpanded[row] = make([]rune, nbCols)
	// 	newCol := 0
	// 	if row%1000 == 0 {
	// 		fmt.Println("row : ", row)
	// 	}
	// 	for col := range newUniverse[row] {
	// 		if slices.Contains(colsWithoutGalaxy, col) {
	// 			for i := 0; i < scale; i++ {
	// 				newUniverseExpanded[row][newCol] = newUniverse[row][col]
	// 				newCol++
	// 			}
	// 		} else {
	// 			newUniverseExpanded[row][newCol] = newUniverse[row][col]
	// 			newCol++
	// 		}
	// 	}
	// }
	// return newUniverseExpanded
}

func computeDistanceAllGalaxies(universe [][]rune, scale int) int {
	rowsWithoutGalaxy, colsWithoutGalaxy := expansionOfTheUniverse(universe, scale)
	galaxies := []pos{}
	for r := range universe {
		for c := range universe[r] {
			if universe[r][c] == '#' {
				galaxies = append(galaxies, pos{row: r, col: c})
			}
		}
	}

	computeDistance := func(a, b pos, rowsWithoutGalaxy, colsWithoutGalaxy []int, scale int) int {
		r, c := 0, 0
		if a.row > b.row {
			a, b = b, a
		}
		for i := a.row; i < b.row; i++ {
			if slices.Contains(rowsWithoutGalaxy, i) {
				r += scale
			} else {
				r++
			}
		}

		if a.col > b.col {
			a, b = b, a
		}
		for i := a.col; i < b.col; i++ {
			if slices.Contains(colsWithoutGalaxy, i) {
				c += scale
			} else {
				c++
			}
		}
		return r + c
	}

	res := 0
	for _g1 := 0; _g1 < len(galaxies)-1; _g1++ {
		g1 := galaxies[_g1]
		for _g2 := _g1 + 1; _g2 < len(galaxies); _g2++ {
			g2 := galaxies[_g2]
			if g1 != g2 {
				res += computeDistance(g1, g2, rowsWithoutGalaxy, colsWithoutGalaxy, scale)
			}
		}
	}
	return res

}

func main() {
	loadFile("input.txt")

	fmt.Println("part1 : ", computeDistanceAllGalaxies(universe, 2))
	fmt.Println("part2 : ", computeDistanceAllGalaxies(universe, 1_000_000))
}
