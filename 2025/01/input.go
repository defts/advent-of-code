package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

type arg struct {
	input []string
}

var input arg

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	input.input = lines
}
