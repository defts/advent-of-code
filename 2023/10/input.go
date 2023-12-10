package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var sketch [][]rune

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	sketch = make([][]rune, 0)
	for _, l := range lines {
		sketch = append(sketch, []rune(l))
	}
}
