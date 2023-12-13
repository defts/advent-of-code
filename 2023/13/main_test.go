package main

import (
	"testing"
)

func Test_findVerticalReflect(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		pattern [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test vertical pattern",
			args: args{
				pattern: patterns[0],
			},
			want: 5,
		},
		{
			name: "test vertical pattern",
			args: args{
				pattern: patterns[1],
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findVerticalReflect(tt.args.pattern); got != tt.want {
				t.Errorf("verticalReflect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findHorizontalReflect(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		pattern [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test horizontal pattern",
			args: args{
				pattern: patterns[0],
			},
			want: -1,
		},
		{
			name: "test horizontal pattern",
			args: args{
				pattern: patterns[1],
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findHorizontalReflect(tt.args.pattern); got != tt.want {
				t.Errorf("findHorizontalReflect() = %v, want %v", got, tt.want)
			}
		})
	}
}
