package main

import (
	"fmt"
	"strings"
)

var cache = make(map[string]int)

func cacheKey(s string, r []int) string {
	return fmt.Sprintf("%s-%s", s, fmt.Sprint(r))
}

func setCache(s string, r []int, value int) int {
	cache[cacheKey(s, r)] = value
	return value
}

func compute(springs string, records []int) int {
	if len(springs) == 0 {
		if len(records) == 0 {
			return 1
		}
		return 0
	}

	// test cache
	if value, ok := cache[cacheKey(springs, records)]; ok {
		return value
	}

	if strings.HasPrefix(springs, ".") {
		// faut enlever tous les points du début et on recommence
		res := compute(strings.TrimLeft(springs, "."), records)
		return setCache(springs, records, res)
	}

	if strings.HasPrefix(springs, "?") {
		// ca peut marcher ou pas faut recommencer
		res := compute(strings.Replace(springs, "?", ".", 1), records) + compute(strings.Replace(springs, "?", "#", 1), records)
		return setCache(springs, records, res)
	}

	// alors là, on rentre dans le dur
	if strings.HasPrefix(springs, "#") {
		// si on a pas de record, pas de record
		if len(records) == 0 || len(springs) < records[0] {
			return setCache(springs, records, 0)
		}

		// si on a un seul point dans la longueur du record, on peut pas
		if strings.ContainsAny(springs[0:records[0]], ".") {
			return setCache(springs, records, 0)
		}

		// et si on a plusieurs records...
		if len(records) > 1 {
			if len(springs) < records[0]+1 || springs[records[0]] == '#' {
				return setCache(springs, records, 0)
			}
			res := compute(springs[records[0]+1:], records[1:])
			return setCache(springs, records, res)
		} else {
			res := compute(springs[records[0]:], records[1:])
			return setCache(springs, records, res)
		}
	}

	panic("nope")
}

func (s springState) unfold() springState {
	var newS springState
	newS.springs = s.springs
	for i := 0; i < 4; i++ {
		newS.springs = fmt.Sprintf("%s?%s", newS.springs, s.springs)
	}
	for i := 0; i < 5; i++ {
		newS.records = append(newS.records, s.records...)
	}

	newS.springs = strings.TrimSpace(newS.springs)

	return newS
}

func computeSpringStates(s []springState) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += compute(s[i].springs, s[i].records)
	}
	return res
}

func main() {
	loadFile("input.txt")
	fmt.Println("part1 :", computeSpringStates(springStates))
	unfolded := make([]springState, 0)
	for _, s := range springStates {
		unfolded = append(unfolded, s.unfold())
	}
	fmt.Println("part2:", computeSpringStates(unfolded))

}
