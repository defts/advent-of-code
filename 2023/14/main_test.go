package main

import "testing"

func Test_compute(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		platform [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{rollNorth(platform)}, 136},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute(tt.args.platform); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
