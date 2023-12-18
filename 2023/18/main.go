package main

import (
	"fmt"
	"math"

	"github.com/defts/advent-of-code/utils"
)

func digger(digPlan []dig) int {
	digMap := make([]utils.Position, 0)
	currentPos := utils.Position{Row: 0, Col: 0}
	digMap = append(digMap, currentPos)
	perimeter := 0
	for _, d := range digPlan {
		for i := 0; i < d.distance; i++ {
			currentPos = currentPos.Move(d.direction, 1)
			digMap = append(digMap, currentPos)
			perimeter++
		}
	}

	nbRow := 1
	nbCol := 1
	deltaRow := 0
	deltaCol := 0
	for _, p := range digMap {
		if p.Row < 0 && int(math.Abs(float64(p.Row))) > deltaRow {
			deltaRow = int(math.Abs(float64(p.Row)))
		}
		if p.Col < 0 && int(math.Abs(float64(p.Col))) > deltaCol {
			deltaCol = int(math.Abs(float64(p.Col)))
		}
		if p.Row+1 > nbRow {
			nbRow = p.Row + 1
		}
		if p.Col+1 > nbCol {
			nbCol = p.Col + 1
		}
	}

	digMapRepositioned := make([]utils.Position, len(digMap))
	for _, d := range digMap {
		digMapRepositioned = append(digMapRepositioned, utils.Position{Row: d.Row + deltaRow, Col: d.Col + deltaCol})
	}

	n := len(digMapRepositioned)
	digMapRepositioned = append(digMapRepositioned, utils.Position{Row: digMapRepositioned[0].Row, Col: digMapRepositioned[0].Col})
	digMapRepositioned = append(digMapRepositioned, utils.Position{Row: digMapRepositioned[1].Row, Col: digMapRepositioned[1].Col})
	area := 0
	for i := 1; i <= n; i++ {
		area += digMapRepositioned[i].Col * (digMapRepositioned[i+1].Row - digMapRepositioned[i-1].Row)
	}

	return area/2 + perimeter/2 + 1
}

func main() {
	loadFile("input.txt")
	fmt.Println("Part 1 : ", digger(digPlan))
	fmt.Println("Part 2 : ", digger(digPlanFromColor))

}
