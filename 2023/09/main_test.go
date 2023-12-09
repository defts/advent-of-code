package main

import (
	"slices"
	"testing"
)

func Test_guessNext(t *testing.T) {
	loadFile("test.txt")
	var part2 [][]int
	for _, sequence := range sequences {
		seq := slices.Clone(sequence)
		slices.Reverse(seq)
		part2 = append(part2, seq)
	}
	type args struct {
		seq []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test part 1 0", args{sequences[0]}, 18},
		{"test part 1 1", args{sequences[1]}, 28},
		{"test part 2 2", args{sequences[2]}, 68},
		{"test part 2 0", args{part2[0]}, -3},
		{"test part 2 1", args{part2[1]}, 0},
		{"test part 2 2", args{part2[2]}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := guessNext(tt.args.seq); got != tt.want {
				t.Errorf("guessNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
