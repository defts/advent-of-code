package main

import (
	"testing"
)

func Test_compute(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				line: lines[0],
			},
			want: 12,
		},
		{
			name: "test2",
			args: args{
				line: lines[1],
			},
			want: 38,
		},
		{
			name: "test3",
			args: args{
				line: lines[2],
			},
			want: 15,
		},
		{
			name: "test4",
			args: args{
				line: lines[3],
			},
			want: 77,
		},
		{
			name: "part2 - test5",
			args: args{
				line: replaceSpelledNumbers(lines[4]),
			},
			want: 29,
		},
		{
			name: "part2 - test6",
			args: args{
				line: replaceSpelledNumbers(lines[5]),
			},
			want: 83,
		},
		{
			name: "part2 - test7",
			args: args{
				line: replaceSpelledNumbers(lines[6]),
			},
			want: 13,
		},
		{
			name: "part2 - test8",
			args: args{
				line: replaceSpelledNumbers(lines[7]),
			},
			want: 24,
		},
		{
			name: "part2 - test9",
			args: args{
				line: replaceSpelledNumbers(lines[8]),
			},
			want: 42,
		},
		{
			name: "part2 - test10",
			args: args{
				line: replaceSpelledNumbers(lines[9]),
			},
			want: 14,
		},
		{
			name: "part2 - test11",
			args: args{
				line: replaceSpelledNumbers(lines[10]),
			},
			want: 76,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute(tt.args.line); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
