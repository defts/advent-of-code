package main

import (
	"testing"
)

func Test_compute(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		schematic [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				schematic: schematic,
			},
			want: 4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute(tt.args.schematic); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeRatio(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		schematic [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				schematic: schematic,
			},
			want: 467835,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeRatio(tt.args.schematic); got != tt.want {
				t.Errorf("computeRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
