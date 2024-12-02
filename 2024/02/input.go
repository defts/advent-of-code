package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

type level []int

var reports []level

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, line := range lines {
		val := utils.StringToIntArray(line)
		reports = append(reports, val)
	}
}
