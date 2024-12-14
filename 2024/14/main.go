package main

import (
	"fmt"

	"github.com/defts/advent-of-code/utils"
)

func part1(arg []robot, nbSeconds int, maxX, maxY int) int {
	defer utils.Timer("part1")()

	robots := make([]robot, len(arg))
	copy(robots, arg)

	for i := 0; i < nbSeconds; i++ {
		for j := 0; j < len(robots); j++ {
			robots[j].pos.Row += robots[j].velocityX
			robots[j].pos.Col += robots[j].velocityY

			if robots[j].pos.Row < 0 {
				robots[j].pos.Row += maxX
			} else if robots[j].pos.Row >= maxX {
				robots[j].pos.Row -= maxX
			}

			if robots[j].pos.Col < 0 {
				robots[j].pos.Col += maxY
			} else if robots[j].pos.Col >= maxY {
				robots[j].pos.Col -= maxY
			}
		}
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		if r.pos.Row < maxX/2 && r.pos.Col < maxY/2 {
			q1++
		}
		if r.pos.Row > maxX/2 && r.pos.Col < maxY/2 {
			q2++
		}
		if r.pos.Row < maxX/2 && r.pos.Col > maxY/2 {
			q3++
		}
		if r.pos.Row > maxX/2 && r.pos.Col > maxY/2 {
			q4++
		}
	}

	// Output the result
	res := q1 * q2 * q3 * q4

	return res
}

func part2(arg []robot, maxX, maxY int) int {
	defer utils.Timer("part2")()

	robots := make([]robot, len(arg))
	copy(robots, arg)

	maxIterations := 1_000_000_000
	res := 0
	isDistinct := false
	for i := 0; i < maxIterations && !isDistinct; i++ {
		// find when all robots are in distinct position
		distinct := make(map[utils.Position]bool)
		for j := 0; j < len(robots); j++ {
			robots[j].pos.Row += robots[j].velocityX
			robots[j].pos.Col += robots[j].velocityY

			if robots[j].pos.Row < 0 {
				robots[j].pos.Row += maxX
			} else if robots[j].pos.Row >= maxX {
				robots[j].pos.Row -= maxX
			}

			if robots[j].pos.Col < 0 {
				robots[j].pos.Col += maxY
			} else if robots[j].pos.Col >= maxY {
				robots[j].pos.Col -= maxY
			}

			distinct[robots[j].pos] = true
		}
		if len(distinct) == len(robots) {
			res = i + 1
			isDistinct = true
			break
		}
	}

	return res
}

func main() {
	loadFile("input.txt")

	fmt.Println("part1: ", part1(arg, 100, 101, 103))
	// on va finalement chercher quand tous les robots sont à un ednroit distinct
	fmt.Println("part2: ", part2(arg, 101, 103))
}
