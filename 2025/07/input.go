package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

type arg struct {
	grid utils.Grid[rune]
}

var input arg

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	var g [][]rune
	for _, line := range lines {
		g = append(g, []rune(line))
	}
	input.grid = utils.NewGrid(g)
}
