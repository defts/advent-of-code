package main

import (
	"fmt"
	"strconv"
)

func isDecimal(r rune) bool {
	return r >= '0' && r <= '9'
}

func isGear(r rune) bool {
	return r == '*'
}

func runToInt(rs ...rune) int {
	n := ""
	for _, r := range rs {
		n += fmt.Sprintf("%d", (int(r - '0')))
	}
	ret, _ := strconv.Atoi(n)
	return ret
}

func hasSymbol(schematic [][]rune, col, start, end int) bool {
	check := func(r rune) bool {
		return r != '.' && !isDecimal(r)
	}
	symbFound := false

	// Left or right
	if start != 0 {
		symbFound = symbFound || check(schematic[col][start-1])
	}
	if end != len(schematic[col])-1 {
		symbFound = symbFound || check(schematic[col][end+1])
	}

	// Top or down
	for line := col - 1; line <= col+1; line += 2 {
		if line < 0 || line >= len(schematic) {
			continue
		}
		for c := start - 1; c <= end+1; c++ {
			if c < 0 || c >= len(schematic[line]) {
				continue
			}
			symbFound = symbFound || check(schematic[line][c])
		}
	}

	return symbFound

}

func compute(schematic [][]rune) int {
	res := 0

	for line := 0; line < len(schematic); line++ {
		for col := 0; col < len(schematic[line]); col++ {
			if !isDecimal(schematic[line][col]) {
				continue
			}
			start := col
			end := -1
			col++
			for ; col < len(schematic[line]); col++ {
				if !isDecimal(schematic[line][col]) {
					end = col - 1
					break
				}
			}
			if end == -1 {
				end = len(schematic[line]) - 1
			}
			col = end
			number := runToInt(schematic[line][start : end+1]...)

			if hasSymbol(schematic, line, start, end) {
				res += number
			}
		}
	}
	return res
}

type pos struct {
	line int
	col  int
}

func findGears(schematic [][]rune) []pos {
	gears := make([]pos, 0)

	for line := 0; line < len(schematic); line++ {
		for col := 0; col < len(schematic[line]); col++ {
			if !isGear(schematic[line][col]) {
				continue
			}
			gears = append(gears, pos{line, col})
		}
	}
	return gears
}

func computeRatio(schematic [][]rune) int {
	res := 0

	gears := findGears(schematic)
	multiplicator := make([][]int, len(gears))

	for line := 0; line < len(schematic); line++ {
		for col := 0; col < len(schematic[line]); col++ {
			if !isDecimal(schematic[line][col]) {
				continue
			}
			start := col
			end := -1
			col++
			for ; col < len(schematic[line]); col++ {
				if !isDecimal(schematic[line][col]) {
					end = col - 1
					break
				}
			}
			if end == -1 {
				end = len(schematic[line]) - 1
			}
			col = end
			number := runToInt(schematic[line][start : end+1]...)

			for g := range gears {
				// a gear at the left or the right of the number?
				if gears[g].line == line {
					if gears[g].col == start-1 {
						multiplicator[g] = append(multiplicator[g], number)
					} else if gears[g].col == end+1 {
						multiplicator[g] = append(multiplicator[g], number)
					}
				}
				// gear at the top
				if gears[g].line == line-1 {
					if gears[g].col >= start-1 && gears[g].col <= end+1 {
						multiplicator[g] = append(multiplicator[g], number)
					}
				}
				// gear at the bottom
				if gears[g].line == line+1 {
					if gears[g].col >= start-1 && gears[g].col <= end+1 {
						multiplicator[g] = append(multiplicator[g], number)
					}
				}
			}
		}
	}

	for _, m := range multiplicator {
		if len(m) == 2 {
			res += m[0] * m[1]
		}
	}
	return res
}

func main() {
	loadFile("input.txt")
	res := compute(schematic)
	fmt.Println("part 1: ", res)

	res = computeRatio(schematic)
	fmt.Println("part 2: ", res)
}
