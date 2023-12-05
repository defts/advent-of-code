package main

import (
	"fmt"
	"sync"
)

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
	// minLocation = 0
	// for i := 0; i < len(seeds); i += 2 {
	// 	for j := 0; j < seeds[i+1]; j++ {
	// 		location := computeSeed(seeds[i] + j)
	// 		if location < minLocation || minLocation == 0 {
	// 			minLocation = location
	// 		}
	// 	}
	// }

	// with goroutines part2
	var wg sync.WaitGroup
	var mu sync.Mutex
	minLocation = 0
	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go func(start, count int) {
			defer wg.Done()

			localMinLocation := minLocation
			for j := 0; j < count; j++ {
				location := computeSeed(start + j)
				if location < localMinLocation || localMinLocation == 0 {
					localMinLocation = location
				}
			}

			mu.Lock()
			defer mu.Unlock()
			if localMinLocation < minLocation || minLocation == 0 {
				minLocation = localMinLocation
			}
		}(seeds[i], seeds[i+1])
	}
	wg.Wait()
	fmt.Printf("minLocation part2: %d\n", minLocation)
}
