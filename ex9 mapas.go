package main

import (
	"fmt"
)

func fibonacciSequence(n int) map[int]int {
	fibMap := make(map[int]int)
	a, b := 0, 1
	index := 0

	for b <= n {
		fibMap[index] = b
		a, b = b, a+b
		index++
	}

	return fibMap
}

func main() {
	n := 50
	fibSequence := fibonacciSequence(n)

	for index, value := range fibSequence {
		fmt.Printf("Índice: %d, Valor: %d\n", index, value)
	}
}
