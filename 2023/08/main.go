package main

import (
	"fmt"
	"strings"
)

func computeSequence(seq map[string][2]string, inputs string) int {
	i := 0
	cnt := 0
	current := "AAA"
	for {
		if current == "ZZZ" {
			break
		}
		if i == len(inputs) {
			i = 0
		}
		input := inputs[i]
		if input == 'L' {
			current = seq[current][0]
		} else {
			current = seq[current][1]
		}
		i++
		cnt++
	}
	return cnt
}

func computeSequencePart2(seq map[string][2]string, inputs string) int {
	var currents []string
	for k := range seq {
		if strings.HasSuffix(k, "A") {
			currents = append(currents, k)
		}
	}

	steps := make([]int, len(currents))

	for i, cur := range currents {
		for !strings.HasSuffix(cur, "Z") {
			for _, input := range inputs {
				if input == 'L' {
					cur = seq[cur][0]
				} else {
					cur = seq[cur][1]
				}

				steps[i]++

				if strings.HasSuffix(cur, "Z") {
					break
				}
			}
		}
	}

	primeFactors := func(n int) []int {
		var factors []int
		for i := 2; i <= n; i++ {
			for n%i == 0 {
				factors = append(factors, i)
				n /= i
			}
		}
		return factors
	}

	res := 1
	for _, s := range steps {
		for _, f := range primeFactors(s) {
			if res%f != 0 {
				res *= f
			}
		}
	}
	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", computeSequence(sequences, inputs))
	fmt.Println("part2: ", computeSequencePart2(sequences, inputs))

}
