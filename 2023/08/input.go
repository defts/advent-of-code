package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

var inputs string

var sequences map[string][2]string

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)

	// AAA = (BBB, CCC)
	inputs = strings.TrimSpace(lines[0])

	sequences = make(map[string][2]string)

	for i := 2; i < len(lines); i++ {
		l := strings.Split(lines[i], "=")
		k := strings.TrimSpace(l[0])
		// (BBB, CCC)
		// extract BBB and CCC
		v := strings.Split(strings.TrimSpace(l[1]), ",")
		sequences[k] = [2]string{strings.Trim(strings.TrimSpace(v[0]), "("), strings.Trim(strings.TrimSpace(v[1]), ")")}
	}
}
