package main

import (
	"fmt"
	"slices"
)

func guessNext(seq []int) int {
	if len(seq) == 1 {
		return 0
	}

	stop := false
	nextVal := 0

	currentSeq := slices.Clone(seq)
	for !stop {
		stop = true
		nextVal += currentSeq[len(currentSeq)-1]

		newSeq := make([]int, len(currentSeq)-1)
		for i := 1; i < len(currentSeq); i++ {
			newSeq[i-1] = currentSeq[i] - currentSeq[i-1]
			if newSeq[i-1] != 0 {
				stop = false
			}
		}
		currentSeq = newSeq
	}
	return nextVal
}

func main() {
	loadFile("input.txt")

	total := 0
	for _, sequence := range sequences {
		total += guessNext(sequence)
	}
	fmt.Println("part1 : ", total)

	total = 0
	for _, sequence := range sequences {
		slices.Reverse(sequence)
		total += guessNext(sequence)
	}
	fmt.Println("part2 : ", total)
}
