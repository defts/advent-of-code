package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var maze = [][]int{}

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for i, l := range lines {
		maze = append(maze, make([]int, len(l)))
		for j, c := range l {
			maze[i][j] = utils.MustParseInt(string(c))
		}
	}
}
