package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

type set struct {
	blue  int
	red   int
	green int
}

type game struct {
	sets []set
}

var games []game

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	for _, line := range lines {
		g := game{}
		sets := strings.Split(line, ":")[1]
		for _, c := range strings.Split(sets, ";") {
			s := set{}
			for _, d := range strings.Split(c, ",") {
				v := strings.Split(strings.TrimSpace(d), " ")
				switch strings.TrimSpace(v[1]) {
				case "blue":
					n, _ := strconv.Atoi(v[0])
					s.blue = n
				case "red":
					n, _ := strconv.Atoi(v[0])
					s.red = n
				case "green":
					n, _ := strconv.Atoi(v[0])
					s.green = n
				}
			}
			g.sets = append(g.sets, s)
		}
		games = append(games, g)
	}
}
