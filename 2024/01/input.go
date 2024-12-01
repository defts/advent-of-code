package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var rowOne []int
var rowTwo []int

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, line := range lines {
		val := utils.StringToIntArray(line)
		rowOne = append(rowOne, val[0])
		rowTwo = append(rowTwo, val[1])
	}
}
