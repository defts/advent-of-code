package main

import "fmt"

func rollNorth(platform [][]rune) [][]rune {
	for col := 0; col < len(platform[0]); col++ {
		index := -1
		for row := 0; row < len(platform); row++ {
			if platform[row][col] == '.' && index == -1 {
				index = row
			}

			if platform[row][col] == '#' {
				index = -1
			}

			if index != -1 && platform[row][col] == 'O' {
				platform[index][col] = 'O'
				platform[row][col] = '.'
				index++
			}
		}
	}

	return platform
}

func compute(platform [][]rune) int {
	res := 0
	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform[i]); j++ {
			if platform[i][j] == 'O' {
				res += len(platform) - i
			}
		}
	}
	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("Part 1:", compute(rollNorth(platform)))

}
