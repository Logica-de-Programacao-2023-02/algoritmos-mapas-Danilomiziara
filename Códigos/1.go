package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "a filha da xuxa se chama sasha"
	wordCount := contarPalavras(text)

	fmt.Println("Contagem de Palavras:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

}

func contarPalavras(text string) map[string]int {
	words := strings.Fields(text)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}
