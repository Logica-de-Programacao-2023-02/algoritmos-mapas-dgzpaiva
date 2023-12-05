package main

import (
	"fmt"
)

func mergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for key, value := range map1 {
		mergedMap[key] = value
	}

	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}

func main() {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"b": 20, "c": 30, "d": 40}

	merged := mergeMaps(map1, map2)

	for key, value := range merged {
		fmt.Printf("%s: %d\n", key, value)
	}
}
