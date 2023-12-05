package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/defts/advent-of-code/utils"
	"github.com/samber/lo"
)

func groupLines(lines []string) [][]string {
	i := 0
	var res [][]string
	var row []string
	for {
		row = append(row, lines[i])
		i++
		if i >= len(lines) {
			res = append(res, row)
			break
		}
		for ; i < len(lines); i++ {
			if lines[i] == "" {
				break
			} else {
				row = append(row, lines[i])
			}
		}
		res = append(res, row)
		row = nil
		i++
		if i >= len(lines) {
			break
		}
	}
	return res
}

var seeds = []int{}
var seedToSoil = [][]int{}
var soilToFertilizer = [][]int{}
var fertilizerToWater = [][]int{}
var waterToLight = [][]int{}
var lightToTemperature = [][]int{}
var temperatureToHumidity = [][]int{}
var humidityToLocation = [][]int{}

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines := utils.ReaderToLines(buff)
	groups := groupLines(lines)

	seedTxt := strings.Split(groups[0][0], ":")[1]
	seeds = utils.StringToIntArray(seedTxt)

	seedToSoil = lo.Map(groups[1][1:], func(x string, i int) []int { return utils.StringToIntArray(x) })
	soilToFertilizer = lo.Map(groups[2][1:], func(x string, i int) []int { return utils.StringToIntArray(x) })
	fertilizerToWater = lo.Map(groups[3][1:], func(x string, i int) []int { return utils.StringToIntArray(x) })
	waterToLight = lo.Map(groups[4][1:], func(x string, i int) []int { return utils.StringToIntArray(x) })
	lightToTemperature = lo.Map(groups[5][1:], func(x string, i int) []int { return utils.StringToIntArray(x) })
	temperatureToHumidity = lo.Map(groups[6][1:], func(x string, i int) []int { return utils.StringToIntArray(x) })
	humidityToLocation = lo.Map(groups[7][1:], func(x string, i int) []int { return utils.StringToIntArray(x) })

}
