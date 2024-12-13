package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

type button struct {
	x, y int
}

type prize struct {
	buttonA button
	buttonB button
	x, y    int
}

var arg []prize

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)

	parseLine := func(line string, fieldIndex int, sep string, prefix string) int {
		parts := strings.Fields(line)
		valueStr := strings.TrimSuffix(strings.Split(parts[fieldIndex], sep)[1], ",")
		return utils.MustParseInt(strings.TrimPrefix(valueStr, prefix))
	}

	for i := 0; i < len(lines); i = i + 4 {
		arg = append(arg, prize{
			button{
				parseLine(lines[i], 2, "+", ""),
				parseLine(lines[i], 3, "+", ""),
			},
			button{
				parseLine(lines[i+1], 2, "+", ""),
				parseLine(lines[i+1], 3, "+", ""),
			},
			parseLine(lines[i+2], 1, "=", ""),
			parseLine(lines[i+2], 2, "=", ""),
		})
	}
}
