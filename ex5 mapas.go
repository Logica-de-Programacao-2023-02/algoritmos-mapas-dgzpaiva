package main

import "fmt"

func countCharacters(input string) map[rune]int {
	charFrequency := make(map[rune]int)

	for _, char := range input {
		charFrequency[char]++
	}

	return charFrequency
}

func main() {
	text := "hello world"
	charFrequencyMap := countCharacters(text)

	for char, freq := range charFrequencyMap {
		fmt.Printf("Caractere: %c | Frequência: %d\n", char, freq)
	}
}
