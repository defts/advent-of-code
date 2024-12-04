package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var grid utils.Grid[rune]

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	var g [][]rune
	for _, line := range lines {
		g = append(g, []rune(line))
	}
	grid = utils.NewGrid(g)
}
