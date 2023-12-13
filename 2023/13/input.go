package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

type pattern [][]rune

var patterns []pattern

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)

	for _, g := range utils.GroupLines(lines) {
		pattern := make(pattern, len(g))
		for i, l := range g {
			pattern[i] = []rune(l)
		}
		patterns = append(patterns, pattern)
	}
}
