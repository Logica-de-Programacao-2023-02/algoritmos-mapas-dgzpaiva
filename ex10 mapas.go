package main

import (
	"fmt"
)

func pairFrequencies(numbers []int) map[[2]int]int {
	pairFreq := make(map[[2]int]int)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			pairFreq[pair]++
		}
	}

	return pairFreq
}

func main() {
	intSlice := []int{1, 2, 2, 3, 4, 1, 3, 2, 5, 4}

	pairFreqMap := pairFrequencies(intSlice)

	for pair, freq := range pairFreqMap {
		fmt.Printf("Par: %v, Frequência: %d\n", pair, freq)
	}
}
