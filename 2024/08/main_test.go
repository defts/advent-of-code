package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	loadFile("test.txt")
	res := 14
	if got := part1(grid); got != res {
		t.Errorf("part1() = %v, want %v", got, res)
	}
}

func Test_part2(t *testing.T) {
	loadFile("test.txt")
	res := 34
	if got := part2(grid); got != res {
		t.Errorf("part2() = %v, want %v", got, res)
	}
}
