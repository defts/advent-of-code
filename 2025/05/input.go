package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

type arg struct {
	ranges      [][2]int
	ingredients []int
}

var input arg

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, l := range lines {
		if l == "" {
			continue
		} else if strings.Contains(l, "-") {
			a, b := 0, 0
			fmt.Sscanf(strings.TrimSpace(l), "%d-%d", &a, &b)
			input.ranges = append(input.ranges, [2]int{a, b})
		} else {
			input.ingredients = append(input.ingredients, utils.MustParseInt(l))
		}
	}
}
