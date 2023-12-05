package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	str := strings.Split(s, "")
	sort.Strings(str)
	return strings.Join(str, "")
}

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	return anagrams
}

func main() {
	wordList := []string{"listen", "silent", "enlist", "hello", "how", "who", "owl"}

	anagramGroups := findAnagrams(wordList)

	for _, group := range anagramGroups {
		if len(group) > 1 {
			fmt.Println(group)
		}
	}
}
