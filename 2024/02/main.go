package main

import (
	"fmt"
	"math"

	"github.com/defts/advent-of-code/utils"
)

func levelIsSafe(l level) bool {
	isSafe := true
	reportIncrease := false
	if l[0] < l[1] {
		reportIncrease = true
	}
	for i := 0; i < len(l)-1; i++ {
		if math.Abs(float64(l[i]-l[i+1])) > 3 || l[i] == l[i+1] {
			isSafe = false
			break
		}
		if reportIncrease {
			if l[i] > l[i+1] {
				isSafe = false
				break
			}
		} else {
			if l[i] < l[i+1] {
				isSafe = false
				break
			}
		}
	}
	return isSafe
}

func part1(report []level) int {
	defer utils.Timer("part1")()
	res := 0

	for _, level := range report {
		if levelIsSafe(level) {
			res++
		}
	}
	return res
}

func part2(report []level) int {
	defer utils.Timer("part2")()
	res := 0
	for _, level := range report {
		levelSafe := levelIsSafe(level)

		for i := range level {
			levelCopy := make([]int, len(level))
			copy(levelCopy, level)
			levelCopy = append(levelCopy[:i], levelCopy[i+1:]...)

			if levelIsSafe(levelCopy) {
				levelSafe = true
				break
			}
		}
		if levelSafe {
			res++
		}
	}
	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(reports))
	fmt.Println("part2: ", part2(reports))
}
