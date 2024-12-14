package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	loadFile("test.txt")
	res := 12
	if got := part1(arg, 100, 11, 7); got != res {
		t.Errorf("part1() = %v, want %v", got, res)
	}
}
