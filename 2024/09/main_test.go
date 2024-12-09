package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	loadFile("test.txt")
	res := 1928
	if got := part1(diskMap); got != res {
		t.Errorf("part1() = %v, want %v", got, res)
	}
}

func Test_part2(t *testing.T) {
	loadFile("test.txt")
	res := 2858
	if got := part2(diskMap); got != res {
		t.Errorf("part2() = %v, want %v", got, res)
	}
}
