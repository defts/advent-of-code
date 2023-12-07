package main

import (
	"testing"
)

func Test_computeHandValue(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		h hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{hands[0]}, onePair},
		{"2", args{hands[1]}, threeOfAKind},
		{"3", args{hands[2]}, twoPairs},
		{"4", args{hands[3]}, twoPairs},
		{"5", args{hands[4]}, threeOfAKind},
		{"6", args{h: hand{cards: []rune{rune('A'), rune('A'), rune('A'), rune('A'), rune('A')}}}, fiveOfAKind},
		{"7", args{h: hand{cards: []rune{rune('A'), rune('A'), rune('8'), rune('A'), rune('A')}}}, fourOfAKind},
		{"8", args{h: hand{cards: []rune{rune('2'), rune('3'), rune('3'), rune('3'), rune('2')}}}, fullHouse},
		{"9", args{h: hand{cards: []rune{rune('T'), rune('T'), rune('T'), rune('9'), rune('8')}}}, threeOfAKind},
		{"10", args{h: hand{cards: []rune{rune('2'), rune('3'), rune('4'), rune('3'), rune('2')}}}, twoPairs},
		{"11", args{h: hand{cards: []rune{rune('A'), rune('2'), rune('3'), rune('A'), rune('4')}}}, onePair},
		{"12", args{h: hand{cards: []rune{rune('2'), rune('3'), rune('4'), rune('5'), rune('6')}}}, highCard},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeHandValue(tt.args.h, false, 0); got != tt.want {
				t.Errorf("computeHandValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeHands(t *testing.T) {
	loadFile("test.txt")
	type args struct {
		hands []hand
		part2 bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{hands, false}, 6440},
		{"2", args{hands, true}, 5905},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeHands(tt.args.hands, tt.args.part2); got != tt.want {
				t.Errorf("computeHands() = %v, want %v", got, tt.want)
			}
		})
	}
}
