package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

type ids struct {
	idOne int
	idTwo int
}
type arg struct {
	ids []ids
}

var input arg

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	line := utils.ReaderToLines(buff)
	for _, part := range strings.Split(line[0], ",") {
		var a, b int
		fmt.Sscanf(strings.TrimSpace(part), "%d-%d", &a, &b)
		input.ids = append(input.ids, ids{a, b})
	}
}
