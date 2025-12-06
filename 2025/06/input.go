package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

type ope struct {
	nums      []int
	numsPart2 []int
	op        string
}

type arg struct {
	ope []ope
}

var input arg

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	input.ope = make([]ope, len(utils.StringToIntArray(lines[0])))
	for _, l := range lines[:len(lines)-1] {
		numbers := utils.StringToIntArray(l)
		for i, v := range numbers {
			input.ope[i].nums = append(input.ope[i].nums, v)
		}
	}

	transposed := ""
	lines2 := lines[:len(lines)-1]
	for x := range lines2[0] {
		s := ""
		for y := range lines2 {
			s += string(lines2[y][x])
		}
		transposed += s + " "
		if strings.TrimSpace(s) == "" {
			transposed += "\n"
		}
	}
	for i, line := range strings.Split(transposed, "\n") {
		input.ope[i].numsPart2 = utils.StringToIntArray(line)
	}

	op := strings.Fields(lines[len(lines)-1])
	for i, v := range op {
		input.ope[i].op = v
	}
}
