package main

import "testing"

func Test_computeDistanceAllGalaxies(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		universe [][]rune
		scale    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test - scale 2",
			args: args{
				universe: universe,
				scale:    2,
			},
			want: 374,
		},
		{
			name: "test - scale 10",
			args: args{
				universe: universe,
				scale:    10,
			},
			want: 1030,
		},
		{
			name: "test - scale 100",
			args: args{
				universe: universe,
				scale:    100,
			},
			want: 8410,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeDistanceAllGalaxies(tt.args.universe, tt.args.scale); got != tt.want {
				t.Errorf("computeDistanceAllGalaxies() = %v, want %v", got, tt.want)
			}
		})
	}
}
