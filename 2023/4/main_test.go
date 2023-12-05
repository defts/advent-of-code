package main

import "testing"

func Test_computeCard(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		c Card
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "test1",
			args:  args{c: cards[0]},
			want:  4,
			want1: 8,
		},
		{
			name:  "test2",
			args:  args{c: cards[1]},
			want:  2,
			want1: 2,
		},
		{
			name:  "test3",
			args:  args{c: cards[2]},
			want:  2,
			want1: 2,
		},
		{
			name:  "test4",
			args:  args{c: cards[3]},
			want:  1,
			want1: 1,
		},
		{
			name:  "test5",
			args:  args{c: cards[4]},
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := computeCard(tt.args.c)
			if got != tt.want {
				t.Errorf("computeCard() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("computeCard() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
