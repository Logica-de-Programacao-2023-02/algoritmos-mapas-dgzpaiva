package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	words := strings.Fields(text)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Esta é uma frase de exemplo. Esta frase contém palavras repetidas, como exemplo."
	wordOccurrences := countWords(text)

	for word, count := range wordOccurrences {
		fmt.Printf("%s: %d\n", word, count)
	}
}
