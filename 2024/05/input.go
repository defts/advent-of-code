package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

var order map[int][]int
var print [][]int

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)

	order = make(map[int][]int)

	for _, line := range lines {
		if strings.Contains(line, "|") {
			n := strings.Split(line, "|")
			order[utils.MustParseInt(n[1])] = append(order[utils.MustParseInt(n[1])], utils.MustParseInt(n[0]))
		} else if strings.Contains(line, ",") {
			n := utils.StringToIntArray(strings.Replace(line, ",", " ", -1))
			print = append(print, n)
		}
	}
}
