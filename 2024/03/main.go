package main

import (
	"fmt"
	"regexp"

	"github.com/defts/advent-of-code/utils"
)

func part1(lines []string) int {
	defer utils.Timer("part1")()
	res := 0

	for _, l := range lines {
		r := regexp.MustCompile(`mul\(\d*,\d*\)`)
		matches := r.FindAllString(l, -1)
		for _, m := range matches {
			r := regexp.MustCompile(`\d+`)
			n := r.FindAllString(m, -1)
			res += utils.MustParseInt(n[0]) * utils.MustParseInt(n[1])
		}

	}
	return res
}

func part2(lines []string) int {
	defer utils.Timer("part2")()
	res := 0

	matches := []string{}
	r := regexp.MustCompile(`mul\(\d*,\d*\)|do\(\)|don't\(\)`)

	for _, l := range lines {
		matches = append(matches, r.FindAllString(l, -1)...)
	}

	do := true

	for _, m := range matches {
		if m == "do()" {
			do = true
			continue
		}
		if m == "don't()" {
			do = false
			continue
		}
		if do {
			num_r := regexp.MustCompile(`\d+`)
			n := num_r.FindAllString(m, -1)
			res += utils.MustParseInt(n[0]) * utils.MustParseInt(n[1])
		}
	}

	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1: ", part1(lines))
	fmt.Println("part2: ", part2(lines))
}
