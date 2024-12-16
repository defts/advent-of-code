package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

type arg struct {
	maze  [][]rune
	start utils.Position
	end   utils.Position
}

var input arg

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, line := range lines {
		input.maze = append(input.maze, []rune(line))
	}

	for i, row := range input.maze {
		for j, node := range row {
			if node == 'S' {
				input.start = utils.Position{Row: i, Col: j}
			}
			if node == 'E' {
				input.end = utils.Position{Row: i, Col: j}
			}
		}
	}

}
