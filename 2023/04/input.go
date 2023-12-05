package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

type Card struct {
	Winning []int
	Having  []int
}

var cards []Card

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, l := range lines {
		c := Card{}
		ll := strings.Split(l, ":")[1]
		c.Winning = utils.StringToIntArray(strings.Split(ll, "|")[0])
		c.Having = utils.StringToIntArray(strings.Split(ll, "|")[1])
		cards = append(cards, c)
	}
}
