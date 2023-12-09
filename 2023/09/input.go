package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var sequences [][]int

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	sequences = make([][]int, 0)
	for _, l := range lines {
		sequences = append(sequences, utils.StringToIntArray(l))
	}
}
