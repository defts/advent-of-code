package main

import (
	"reflect"
	"testing"
)

func Test_neighbors(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		sketch [][]rune
		p      pos
	}
	tests := []struct {
		name string
		args args
		want []pos
	}{
		{
			name: "valid test L",
			args: args{
				sketch: sketch,
				p:      pos{0, 1},
			},
			want: []pos{pos{0, 0}, pos{0, 2}},
		},
		{
			name: "valid test F",
			args: args{
				sketch: sketch,
				p:      pos{0, 2},
			},
			want: []pos{pos{1, 2}, pos{0, 3}},
		},
		{
			name: "valid test J",
			args: args{
				sketch: sketch,
				p:      pos{1, 1},
			},
			want: []pos{pos{2, 1}, pos{1, 2}},
		},
		{
			name: "invalid bound",
			args: args{
				sketch: sketch,
				p:      pos{0, 4},
			},
			want: []pos{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.want) == 0 {
				if got := neighbors(tt.args.sketch, tt.args.p); got != nil {
					t.Errorf("neighbors() = %v, want %v", got, tt.want)
				}
				return
			}

			if got := neighbors(tt.args.sketch, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}
