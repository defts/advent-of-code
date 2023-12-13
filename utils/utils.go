package utils

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ReaderToLines(input io.Reader) []string {
	scanner := bufio.NewScanner(input)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func StringToIntArray(input string) []int {
	numStrs := strings.Fields(input)
	var intArray []int
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil
		}
		intArray = append(intArray, num)
	}
	return intArray
}

func MustParseInt(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return num
}

func GroupLines(lines []string) [][]string {
	i := 0
	var res [][]string
	var row []string
	for {
		row = append(row, lines[i])
		i++
		if i >= len(lines) {
			res = append(res, row)
			break
		}
		for ; i < len(lines); i++ {
			if lines[i] == "" {
				break
			} else {
				row = append(row, lines[i])
			}
		}
		res = append(res, row)
		row = nil
		i++
		if i >= len(lines) {
			break
		}
	}
	return res
}
