package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/defts/advent-of-code/utils"
)

type dig struct {
	direction utils.Direction
	distance  int
}

var digPlan = []dig{}
var digPlanFromColor = []dig{}

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)

	for _, l := range lines {
		t := strings.Split(l, " ")
		dir := utils.Up
		switch t[0] {
		case "R":
			dir = utils.Right
		case "L":
			dir = utils.Left
		case "U":
			dir = utils.Up
		case "D":
			dir = utils.Down
		}
		digPlan = append(digPlan, dig{
			direction: dir,
			distance:  utils.MustParseInt(t[1]),
		})

		stringColor := t[2]
		stringColor = strings.TrimSpace(stringColor)
		stringColor = strings.Trim(stringColor, "(#")
		stringColor = strings.Trim(stringColor, ")")

		distance := stringColor[:5]
		count, _ := strconv.ParseInt(distance, 16, 64)

		switch stringColor[5] {
		case '0':
			dir = utils.Right
		case '1':
			dir = utils.Down
		case '2':
			dir = utils.Left
		case '3':
			dir = utils.Up
		}

		digPlanFromColor = append(digPlanFromColor, dig{
			direction: dir,
			distance:  int(count),
		})
	}
}
