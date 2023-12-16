package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var contraption [][]rune

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, l := range lines {
		contraption = append(contraption, []rune(l))
	}
}
