package main

import (
	"fmt"
	"regexp"
	"slices"

	"github.com/defts/advent-of-code/utils"
)

func hash(s string) int {
	h := 0
	for _, v := range s {
		h += int(v)
		h = h * 17
		h = h % 256
	}
	return h
}

type box struct {
	label string
	focal int
}

func lensBoxes(sequence []string) [256][]box {
	boxes := [256][]box{}
	reg := regexp.MustCompile(`^(.*)([-=])(.*)$`)
	for _, seq := range sequence {
		m := reg.FindStringSubmatch(seq)
		label := m[1]
		operator := m[2]
		value := 0
		if m[3] != "" {
			value = utils.MustParseInt(m[3])
		}

		// on recupere la box courrante depuis le hash
		currentBox := boxes[hash(label)]

		// si on connait deja le label, on recupere son index
		index := -1
		for _, b := range boxes {
			for i, l := range b {
				if l.label == label {
					index = i
					break
				}
			}
		}

		// si on a trouve le label, on met a jour la box courrante
		if index >= 0 {
			if operator == "-" {
				// on supprime la box courrante
				currentBox = slices.Delete(currentBox, index, index+1)
			} else {
				// on met a jour la box courrante avec le label et la valeur
				currentBox[index] = box{label, value}
			}
		} else {
			if operator == "=" {
				currentBox = append(currentBox, box{label, value})
			}
		}

		// on met a jour la box courrante
		boxes[hash(label)] = currentBox
	}

	return boxes

}

func main() {
	loadFile("input.txt")

	part1 := 0
	for _, v := range sequence {
		part1 += hash(v)
	}
	println("part1:", part1)

	part2 := 0
	for i, b := range lensBoxes(sequence) {
		for j, l := range b {
			part2 += (i + 1) * (j + 1) * l.focal
		}
	}

	fmt.Println("part2: ", part2)

}
