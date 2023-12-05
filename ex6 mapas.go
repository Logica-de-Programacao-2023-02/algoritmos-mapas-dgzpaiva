package main

import "fmt"

func mergeWordCounts(listOfMaps []map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for _, wordMap := range listOfMaps {
		for word, count := range wordMap {
			mergedMap[word] += count
		}
	}

	return mergedMap
}

func main() {
	map1 := map[string]int{"hello": 3, "world": 2, "example": 1}
	map2 := map[string]int{"hello": 2, "world": 1, "new": 5}
	map3 := map[string]int{"hello": 1, "example": 4, "new": 2}

	listOfMaps := []map[string]int{map1, map2, map3}

	merged := mergeWordCounts(listOfMaps)

	for word, count := range merged {
		fmt.Printf("%s: %d\n", word, count)
	}
}
