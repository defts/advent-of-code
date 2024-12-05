package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		order map[int][]int
		print [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				order: order,
				print: print,
			},
			want: 143,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.order, tt.args.print); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		order map[int][]int
		print [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				order: order,
				print: print,
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.order, tt.args.print); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
