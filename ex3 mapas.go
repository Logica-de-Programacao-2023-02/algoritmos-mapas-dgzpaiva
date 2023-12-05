package main

import (
	"fmt"
)

func sumMapValues(inputMap map[string]int) int {
	sum := 0

	for _, value := range inputMap {
		sum += value
	}

	return sum
}

func main() {
	myMap := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	totalSum := sumMapValues(myMap)
	fmt.Println("Soma dos valores do mapa:", totalSum)
}
