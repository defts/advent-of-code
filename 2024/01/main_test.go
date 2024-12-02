package main

import "testing"

func Test_part1(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		rowOne []int
		rowTwo []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				rowOne: rowOne,
				rowTwo: rowTwo,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.rowOne, tt.args.rowTwo); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		rowOne []int
		rowTwo []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				rowOne: rowOne,
				rowTwo: rowTwo,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.rowOne, tt.args.rowTwo); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
