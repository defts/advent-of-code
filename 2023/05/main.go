package main

import "fmt"

func computeMap(i int, m [][]int) (o int) {
	for _, l := range m {
		if i >= l[1] && i < l[1]+l[2] {
			return l[0] + (i - l[1])
		}
	}
	return i
}

func computeSeed(seed int) (location int) {
	soil := computeMap(seed, seedToSoil)
	fertilizer := computeMap(soil, soilToFertilizer)
	water := computeMap(fertilizer, fertilizerToWater)
	light := computeMap(water, waterToLight)
	temperature := computeMap(light, lightToTemperature)
	humidity := computeMap(temperature, temperatureToHumidity)
	return computeMap(humidity, humidityToLocation)

}

func main() {
	loadFile("input.txt")

	minLocation := 0
	for _, seed := range seeds {
		location := computeSeed(seed)
		if location < minLocation || minLocation == 0 {
			minLocation = location
		}
	}
	fmt.Printf("minLocation: %d\n", minLocation)

	// part 2
	minLocation = 0
	for i := 0; i < len(seeds); i += 2 {
		for j := 0; j < seeds[i+1]; j++ {
			location := computeSeed(seeds[i] + j)
			if location < minLocation || minLocation == 0 {
				minLocation = location
			}
		}
	}
	fmt.Printf("minLocation part2: %d\n", minLocation)
}
