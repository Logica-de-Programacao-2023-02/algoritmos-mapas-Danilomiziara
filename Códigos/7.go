package main

import (
	"fmt"
	"strings"
)

func countLettersInWords(sentence string) map[string]map[rune]int {
	wordCounts := make(map[string]map[rune]int)

	words := strings.Fields(sentence)

	for _, word := range words {
		letterCount := make(map[rune]int)

		for _, char := range word {
			letterCount[char]++
		}

		wordCounts[word] = letterCount
	}

	return wordCounts
}

func main() {
	sentence := "corre berg"

	letterCounts := countLettersInWords(sentence)

	fmt.Println("Contagem de letras em cada palavra:")
	for word, letterCount := range letterCounts {
		fmt.Printf("%s: %v\n", word, letterCount)
	}
}
