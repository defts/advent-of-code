package main

import (
	"testing"
)

func Test_computeMap(t *testing.T) {
	// load test file
	loadFile("test.txt")

	type args struct {
		i int
		m [][]int
	}
	tests := []struct {
		name  string
		args  args
		wantO int
	}{
		{
			name: "seedToSoil",
			args: args{
				i: 79,
				m: seedToSoil,
			},
			wantO: 81,
		},
		{
			name: "seedToSoil",
			args: args{
				i: 14,
				m: seedToSoil,
			},
			wantO: 14,
		},
		{
			name: "seedToSoil",
			args: args{
				i: 55,
				m: seedToSoil,
			},
			wantO: 57,
		},
		{
			name: "seedToSoil",
			args: args{
				i: 13,
				m: seedToSoil,
			},
			wantO: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotO := computeMap(tt.args.i, tt.args.m); gotO != tt.wantO {
				t.Errorf("computeMap() = %v, want %v", gotO, tt.wantO)
			}
		})
	}
}

func Test_computeSeed(t *testing.T) {
	// load test file
	loadFile("test.txt")

	type args struct {
		seed int
	}
	tests := []struct {
		name         string
		args         args
		wantLocation int
	}{
		{
			name: "seed 79",
			args: args{
				seed: 79,
			},
			wantLocation: 82,
		},
		{
			name: "seed 14",
			args: args{
				seed: 14,
			},
			wantLocation: 43,
		},
		{
			name: "seed 55",
			args: args{
				seed: 55,
			},
			wantLocation: 86,
		},
		{
			name: "seed 13",
			args: args{
				seed: 13,
			},
			wantLocation: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLocation := computeSeed(tt.args.seed); gotLocation != tt.wantLocation {
				t.Errorf("computeSeed() = %v, want %v", gotLocation, tt.wantLocation)
			}
		})
	}
}
