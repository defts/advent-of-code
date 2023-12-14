package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var platform [][]rune

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	platform = [][]rune{}
	for _, l := range lines {
		platform = append(platform, []rune(l))
	}
}
