package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

var times []int
var records []int

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	times = utils.StringToIntArray(strings.Split(lines[0], ":")[1])
	records = utils.StringToIntArray(strings.Split(lines[1], ":")[1])
}
