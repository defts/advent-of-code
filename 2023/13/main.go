package main

import "fmt"

func findVerticalReflect(pattern [][]rune) int {
p:
	for col := 0; col < len(pattern[0])-1; col++ {
		for a := 0; a <= col; a++ {
			b := col + col - a + 1
			if b >= len(pattern[0]) {
				continue
			}
			for row := range pattern {
				if pattern[row][a] != pattern[row][b] {
					// on sort de la boucle interne
					continue p
				}
			}
		}
		return col + 1
	}

	return -1
}

func findHorizontalReflect(pattern [][]rune) int {
p:
	for row := 0; row < len(pattern)-1; row++ {
		for a := 0; a <= row; a++ {
			b := row + row - a + 1
			if b >= len(pattern) {
				continue
			}
			for col := range pattern[0] {
				if pattern[a][col] != pattern[b][col] {
					// on sort de la boucle interne
					continue p
				}
			}
		}
		return row + 1
	}

	return -1
}

func main() {
	loadFile("input.txt")
	res := 0
	for _, p := range patterns {
		h := findHorizontalReflect(p)
		if h != -1 {
			res += h * 100
		} else {
			v := findVerticalReflect(p)
			if v != -1 {
				res += v
			}
		}
	}
	fmt.Println("part1:", res)
}
