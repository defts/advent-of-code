package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/defts/advent-of-code/utils"
)

type robot struct {
	pos       utils.Position
	velocityX int
	velocityY int
}

var arg []robot

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)

	for _, line := range lines {
		r := robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.pos.Row, &r.pos.Col, &r.velocityX, &r.velocityY)
		arg = append(arg, r)
	}
}
