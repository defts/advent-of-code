package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

var equations map[int][]int

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	equations = make(map[int][]int)
	for _, l := range lines {
		r := strings.Split(l, ":")
		equations[utils.MustParseInt(r[0])] = utils.StringToIntArray(r[1])

	}
}
