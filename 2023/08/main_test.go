package main

import (
	"testing"
)

func Test_computeSequence(t *testing.T) {
	type args struct {
		seq    map[string][2]string
		inputs string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				map[string][2]string{
					"AAA": {"BBB", "BBB"},
					"BBB": {"AAA", "ZZZ"},
					"ZZZ": {"ZZZ", "ZZZ"},
				},
				"LLR",
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeSequence(tt.args.seq, tt.args.inputs); got != tt.want {
				t.Errorf("computeSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeSequencePart2(t *testing.T) {
	type args struct {
		seq    map[string][2]string
		inputs string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				map[string][2]string{
					"11A": {"11B", "XXX"},
					"11B": {"XXX", "11Z"},
					"11Z": {"11B", "XXX"},
					"22A": {"22B", "XXX"},
					"22B": {"22C", "22C"},
					"22C": {"22Z", "22Z"},
					"22Z": {"22B", "22B"},
					"XXX": {"XXX", "XXX"},
				},
				"LR",
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeSequencePart2(tt.args.seq, tt.args.inputs); got != tt.want {
				t.Errorf("computeSequencePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
