package main

import "fmt"

func gameIsPossible(g game, maxBlue, maxRed, maxGreen int) bool {
	for _, s := range g.sets {
		if s.blue > maxBlue || s.red > maxRed || s.green > maxGreen {
			return false
		}
	}
	return true
}

func getPower(g game) int {
	minBlue := 0
	minRed := 0
	minGreen := 0
	for _, s := range g.sets {
		if s.blue > minBlue {
			minBlue = s.blue
		}
		if s.red > minRed {
			minRed = s.red
		}
		if s.green > minGreen {
			minGreen = s.green
		}
	}
	return minBlue * minRed * minGreen
}

func main() {
	loadFile("input.txt")
	res := 0
	for r, g := range games {
		if gameIsPossible(g, 14, 12, 13) {
			res += r + 1
		}
	}
	fmt.Println("part1: ", res)

	res = 0
	for _, g := range games {
		res += getPower(g)
	}
	fmt.Println("part2: ", res)

}
