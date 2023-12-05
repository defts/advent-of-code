package main

import (
	"testing"
)

func Test_gameIsPossible(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		g        game
		maxBlue  int
		maxRed   int
		maxGreen int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				g:        games[0],
				maxBlue:  14,
				maxRed:   12,
				maxGreen: 13,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				g:        games[1],
				maxBlue:  14,
				maxRed:   12,
				maxGreen: 13,
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				g:        games[2],
				maxBlue:  14,
				maxRed:   12,
				maxGreen: 13,
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				g:        games[3],
				maxBlue:  14,
				maxRed:   12,
				maxGreen: 13,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gameIsPossible(tt.args.g, tt.args.maxBlue, tt.args.maxRed, tt.args.maxGreen); got != tt.want {
				t.Errorf("gameIsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPower(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		g game
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				g: games[0],
			},
			want: 48,
		},
		{
			name: "test2",
			args: args{
				g: games[1],
			},
			want: 12,
		},
		{
			name: "test3",
			args: args{
				g: games[2],
			},
			want: 1560,
		},
		{
			name: "test4",
			args: args{
				g: games[3],
			},
			want: 630,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPower(tt.args.g); got != tt.want {
				t.Errorf("getPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
