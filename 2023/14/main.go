package main

import (
	"fmt"
)

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

func rollEast(platform [][]rune) [][]rune {
	for row := len(platform) - 1; row >= 0; row-- {
		index := -1
		for col := len(platform[0]) - 1; col >= 0; col-- {
			if platform[row][col] == '.' && index == -1 {
				index = col
			}

			if platform[row][col] == '#' {
				index = -1
			}

			if index != -1 && platform[row][col] == 'O' {
				platform[row][index] = 'O'
				platform[row][col] = '.'
				index--
			}
		}
	}
	return platform
}

func rollSouth(platform [][]rune) [][]rune {
	for col := len(platform[0]) - 1; col >= 0; col-- {
		index := -1
		for row := len(platform) - 1; row >= 0; row-- {
			if platform[row][col] == '.' && index == -1 {
				index = row
			}

			if platform[row][col] == '#' {
				index = -1
			}

			if index != -1 && platform[row][col] == 'O' {
				platform[index][col] = 'O'
				platform[row][col] = '.'
				index--
			}
		}
	}
	return platform
}

func rollWest(platform [][]rune) [][]rune {
	for row := 0; row < len(platform); row++ {
		index := -1
		for col := 0; col < len(platform[0]); col++ {
			if platform[row][col] == '.' && index == -1 {
				index = col
			}

			if platform[row][col] == '#' {
				index = -1
			}

			if index != -1 && platform[row][col] == 'O' {
				platform[row][index] = 'O'
				platform[row][col] = '.'
				index++
			}
		}
	}
	return platform
}

func makeACycle(platform [][]rune) [][]rune {
	platform = rollNorth(platform)
	platform = rollWest(platform)
	platform = rollSouth(platform)
	platform = rollEast(platform)
	return platform
}

func cacheKey(platform [][]rune) string {
	res := ""
	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform[i]); j++ {
			res += string(platform[i][j])
		}
	}
	return res
}

func makeCycles(platform [][]rune, times int) [][]rune {
	var cache = make(map[string]int)
	breakP := times
	for cycle := 0; cycle < times; cycle++ {
		platform = makeACycle(platform)
		if i, ok := cache[cacheKey(platform)]; ok {
			breakP = cycle + (times-i)%(cycle-i) - 1
		} else {
			cache[cacheKey(platform)] = cycle
		}

		if cycle == breakP {
			break
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
	fmt.Println("Part 2:", compute(makeCycles(platform, 1_000_000_000)))

}
