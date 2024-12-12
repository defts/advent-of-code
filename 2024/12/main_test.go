package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	loadFile("test.txt")
	res := 1930
	if got := part1(arg); got != res {
		t.Errorf("part1() = %v, want %v", got, res)
	}
}

func Test_part2(t *testing.T) {
	loadFile("test.txt")
	res := 1206
	if got := part2(arg); got != res {
		t.Errorf("part2() = %v, want %v", got, res)
	}
}
