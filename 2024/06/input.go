package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var grid [][]rune

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, l := range lines {
		grid = append(grid, []rune(l))
	}
}
