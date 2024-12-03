package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				lines: []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"},
			},
			want: 161,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				lines: []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
