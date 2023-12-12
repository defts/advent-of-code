package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

type springState struct {
	springs string
	records []int
}

var springStates []springState

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	springStates = make([]springState, 0)
	for _, l := range lines {
		s := springState{}
		s.springs = strings.TrimSpace(strings.Split(l, " ")[0])
		for _, n := range strings.Split(strings.Split(l, " ")[1], ",") {
			s.records = append(s.records, utils.MustParseInt(n))
		}
		springStates = append(springStates, s)
	}
}
