package main

import (
	"fmt"
)

func equalizeExpenses(expenses map[string]float64) map[string]float64 {
	total := 0.0
	equalizedExpenses := make(map[string]float64)

	for _, amount := range expenses {
		total += amount
	}

	average := total / float64(len(expenses))

	for person, amountSpent := range expenses {
		equalizedExpenses[person] = amountSpent - average
	}

	return equalizedExpenses
}

func main() {
	expenses := map[string]float64{
		"Alice":  120.50,
		"Bob":    90.25,
		"Charlie": 150.75,
		"David":  110.0,
	}

	equalized := equalizeExpenses(expenses)

	for person, balance := range equalized {
		fmt.Printf("%s: %.2f\n", person, balance)
	}
}
