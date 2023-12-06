package main

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

func computeRace(time int) (result map[int]int) {
	result = make(map[int]int)
	for i := 0; i < time; i++ {
		speed := i
		timeLeft := time - i
		result[i] = speed * timeLeft
	}
	return result
}

func countWinningRace(time, record int) (winning map[int]int) {
	res := computeRace(time)
	winning = make(map[int]int)
	for i, r := range res {
		if r > record {
			winning[i] = r
		}
	}
	return winning
}

func main() {
	loadFile("input.txt")

	res := make([]int, 0)
	for i := range times {
		res = append(res, len(countWinningRace(times[i], records[i])))
	}

	fmt.Println("Part1: ", lo.Reduce(res, func(agg, i, _ int) int { return agg * i }, 1))

	// part 2
	part2TimesTxt := ""
	for _, v := range times {
		part2TimesTxt = fmt.Sprintf("%v%d", part2TimesTxt, v)
	}
	part2Time, _ := strconv.Atoi(part2TimesTxt)

	part2Recordstxt := ""
	for _, v := range records {
		part2Recordstxt = fmt.Sprintf("%v%d", part2Recordstxt, v)
	}
	part2Record, _ := strconv.Atoi(part2Recordstxt)

	fmt.Println("Part2: ", len(countWinningRace(part2Time, part2Record)))
}
