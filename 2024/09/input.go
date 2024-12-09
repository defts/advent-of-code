package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var diskMap []int

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, l := range lines {
		for _, c := range l {
			diskMap = append(diskMap, utils.MustParseInt(string(c)))
		}
	}
}
