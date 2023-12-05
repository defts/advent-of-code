package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var schematic [][]rune

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, line := range lines {
		schematic = append(schematic, []rune(line))
	}
}
