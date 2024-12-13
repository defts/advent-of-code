package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func calculate(prize prize) (float64, float64) {
	divisor := (prize.buttonA.x*prize.buttonB.y - prize.buttonA.y*prize.buttonB.x)
	A := float64(prize.x*prize.buttonB.y-prize.y*prize.buttonB.x) / float64(divisor)
	B := float64(prize.buttonA.x*prize.y-prize.buttonA.y*prize.x) / float64(divisor)
	return A, B
}

func isValid(A, B float64) bool {
	return A >= 0 && B >= 0 && A == float64(int(A)) && B == float64(int(B))
}

func part1(prizes []prize) int {
	defer utils.Timer("part1")()
	res := 0

	for _, p := range prizes {
		A, B := calculate(p)
		if isValid(A, B) && A <= 100 && B <= 100 {
			res += int(A*3 + B)
		}
	}

	return res
}

func part2(prizes []prize) int {
	defer utils.Timer("part2")()
	res := 0

	// Offset prizes' coordinates
	for i := range prizes {
		prizes[i].x += 10000000000000
		prizes[i].y += 10000000000000
	}

	for _, p := range prizes {
		A, B := calculate(p)
		if isValid(A, B) {
			res += int(A*3 + B)
		}
	}

	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println(arg[0])

	fmt.Println("part1: ", part1(arg)) // 28262
	fmt.Println("part2: ", part2(arg)) // 101406661266314
}
