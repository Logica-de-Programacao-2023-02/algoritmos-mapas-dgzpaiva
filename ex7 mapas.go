package main

import (
	"fmt"
	"strings"
)

func countLetters(word string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, letter := range word {
		letterCount[letter]++
	}

	return letterCount
}

func wordLetterCount(text string) map[string]map[rune]int {
	words := strings.Fields(text)
	wordMap := make(map[string]map[rune]int)

	for _, word := range words {
		wordMap[word] = countLetters(word)
	}

	return wordMap
}

func main() {
	text := "Esta é uma frase de exemplo"
	wordLetterCounts := wordLetterCount(text)

	for word, letterCount := range wordLetterCounts {
		fmt.Printf("%s:\n", word)
		for letter, count := range letterCount {
			fmt.Printf("  %c: %d\n", letter, count)
		}
	}
}
