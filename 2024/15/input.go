package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

type arg struct {
	puzzle [][]rune
	robot  utils.Position
	inputs []rune
}

var input arg

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	g := [][]rune{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			g = append(g, []rune(line))
			for i, r := range line {
				if r == '@' {
					input.robot = utils.Position{Row: i, Col: len(g) - 1}
				}
			}
			continue
		}
		input.inputs = append(input.inputs, []rune(line)...)
	}
	input.puzzle = g
}
