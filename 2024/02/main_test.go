package main

import (
	"testing"
)

// func Test_compute(t *testing.T) {
// 	loadFile("test.txt")
// 	type args struct {
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := compute(tt.args.line); got != tt.want {
// 				t.Errorf("compute() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_part1(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		reports []level
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				reports: reports,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.reports); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		reports []level
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				reports: reports,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.reports); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
