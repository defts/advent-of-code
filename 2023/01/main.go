package main

import (
	"fmt"
	"strconv"
)

func replaceSpelledNumbers(line string) string {
	replacedLine := ""
	for i := 0; i < len(line); i++ {
		if len(line[i:]) >= len("one") && line[i:len("one")+i] == "one" {
			replacedLine += replacedLine + "1"
			continue
		}
		if len(line[i:]) >= len("two") && line[i:len("two")+i] == "two" {
			replacedLine += replacedLine + "2"
			continue
		}
		if len(line[i:]) >= len("three") && line[i:len("three")+i] == "three" {
			replacedLine += replacedLine + "3"
			continue
		}
		if len(line[i:]) >= len("four") && line[i:len("four")+i] == "four" {
			replacedLine += replacedLine + "4"
			continue
		}
		if len(line[i:]) >= len("five") && line[i:len("five")+i] == "five" {
			replacedLine += replacedLine + "5"
			continue
		}
		if len(line[i:]) >= len("six") && line[i:len("six")+i] == "six" {
			replacedLine += replacedLine + "6"
			continue
		}
		if len(line[i:]) >= len("seven") && line[i:len("seven")+i] == "seven" {
			replacedLine += replacedLine + "7"
			continue
		}
		if len(line[i:]) >= len("eight") && line[i:len("eight")+i] == "eight" {
			replacedLine += replacedLine + "8"
			continue
		}
		if len(line[i:]) >= len("nine") && line[i:len("nine")+i] == "nine" {
			replacedLine += replacedLine + "9"
			continue
		}
		replacedLine += string(line[i])
	}

	return replacedLine
}

func compute(line string) int {
	start := -1
	end := -1
	for _, c := range line {
		if c >= '0' && c <= '9' && start == -1 {
			n, _ := strconv.Atoi(string(c))
			start = n
		}
		if c >= '0' && c <= '9' {
			n, _ := strconv.Atoi(string(c))
			end = n
		}
	}
	return start*10 + end

}

func main() {
	loadFile("input.txt")
	res := 0
	for _, line := range lines {
		res += compute(line)
	}
	fmt.Println("part1: ", res)

	res = 0
	for _, line := range lines {
		res += compute(replaceSpelledNumbers(line))
	}
	fmt.Println("part2: ", res)
}
