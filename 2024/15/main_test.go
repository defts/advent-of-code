package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	loadFile("test.txt")
	res := 10092
	if got := part1(input); got != res {
		t.Errorf("part1() = %v, want %v", got, res)
	}
}

func Test_part2(t *testing.T) {
	loadFile("test.txt")
	res := 9021
	if got := part2(input); got != res {
		t.Errorf("part2() = %v, want %v", got, res)
	}
}
