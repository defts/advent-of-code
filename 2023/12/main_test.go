package main

import (
	"reflect"
	"testing"
)

func Test_compute(t *testing.T) {
	loadFile("test.txt")
	springStatesUnfolded := make([]springState, 0)
	for _, s := range springStates {
		springStatesUnfolded = append(springStatesUnfolded, s.unfold())
	}
	type args struct {
		springs string
		records []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "part 1 - test0",
			args: args{
				springs: springStates[0].springs,
				records: springStates[0].records,
			},
			want: 1,
		},
		{
			name: "part 1 - test1",
			args: args{
				springs: springStates[1].springs,
				records: springStates[1].records,
			},
			want: 4,
		},
		{
			name: "part 1 - test2",
			args: args{
				springs: springStates[2].springs,
				records: springStates[2].records,
			},
			want: 1,
		},
		{
			name: "part 1 - test3",
			args: args{
				springs: springStates[3].springs,
				records: springStates[3].records,
			},
			want: 1,
		},
		{
			name: "part 1 - test4",
			args: args{
				springs: springStates[4].springs,
				records: springStates[4].records,
			},
			want: 4,
		},
		{
			name: "part 1 - test5",
			args: args{
				springs: springStates[5].springs,
				records: springStates[5].records,
			},
			want: 10,
		},
		{
			name: "part 2 - test0",
			args: args{
				springs: springStatesUnfolded[0].springs,
				records: springStatesUnfolded[0].records,
			},
			want: 1,
		},
		{
			name: "part 2 - test1",
			args: args{
				springs: springStatesUnfolded[1].springs,
				records: springStatesUnfolded[1].records,
			},
			want: 16384,
		},
		{
			name: "part 2 - test2",
			args: args{
				springs: springStatesUnfolded[2].springs,
				records: springStatesUnfolded[2].records,
			},
			want: 1,
		},
		{
			name: "part 2 - test3",
			args: args{
				springs: springStatesUnfolded[3].springs,
				records: springStatesUnfolded[3].records,
			},
			want: 16,
		},
		{
			name: "part 2 - test4",
			args: args{
				springs: springStatesUnfolded[4].springs,
				records: springStatesUnfolded[4].records,
			},
			want: 2500,
		},
		{
			name: "part 1 - test5",
			args: args{
				springs: springStatesUnfolded[5].springs,
				records: springStatesUnfolded[5].records,
			},
			want: 506250,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute(tt.args.springs, tt.args.records); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_springState_unflod(t *testing.T) {
	loadFile("test.txt")
	type fields struct {
		springs string
		records []int
	}
	tests := []struct {
		name   string
		fields fields
		want   springState
	}{
		{
			name: "test0",
			fields: fields{
				springs: springStates[0].springs,
				records: springStates[0].records,
			},
			want: springState{
				springs: "???.###????.###????.###????.###????.###",
				records: []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := springState{
				springs: tt.fields.springs,
				records: tt.fields.records,
			}
			if got := s.unfold(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("springState.unflod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeSpringState(t *testing.T) {
	loadFile("test.txt")
	springStatesUnfolded := make([]springState, 0)
	for _, s := range springStates {
		springStatesUnfolded = append(springStatesUnfolded, s.unfold())
	}
	type args struct {
		s []springState
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "part 1 - test0",
			args: args{
				s: springStates,
			},
			want: 21,
		},
		{
			name: "part 2 - test0",
			args: args{
				s: springStatesUnfolded,
			},
			want: 525152,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeSpringStates(tt.args.s); got != tt.want {
				t.Errorf("computeSpringState() = %v, want %v", got, tt.want)
			}
		})
	}
}
