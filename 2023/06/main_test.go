package main

import (
	"reflect"
	"testing"
)

func Test_computeRace(t *testing.T) {
	type args struct {
		time int
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[int]int
	}{
		{
			name:       "test1",
			args:       args{time: 7},
			wantResult: map[int]int{0: 0, 1: 6, 2: 10, 3: 12, 4: 12, 5: 10, 6: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := computeRace(tt.args.time); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("computeRace() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_countWinningRace(t *testing.T) {
	type args struct {
		time   int
		record int
	}
	tests := []struct {
		name        string
		args        args
		wantWinning map[int]int
	}{
		{
			name:        "test1",
			args:        args{time: 7, record: 9},
			wantWinning: map[int]int{2: 10, 3: 12, 4: 12, 5: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWinning := countWinningRace(tt.args.time, tt.args.record); !reflect.DeepEqual(gotWinning, tt.wantWinning) {
				t.Errorf("countWinningRace() = %v, want %v", gotWinning, tt.wantWinning)
			}
		})
	}
}
