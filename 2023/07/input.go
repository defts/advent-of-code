package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

type hand struct {
	cards []rune
	bid   int
	value int
}

var hands []hand

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, line := range lines {
		h := hand{}
		h.cards = []rune(strings.Split(line, " ")[0])
		h.bid = utils.MustParseInt(strings.Split(line, " ")[1])
		hands = append(hands, h)
	}
}
