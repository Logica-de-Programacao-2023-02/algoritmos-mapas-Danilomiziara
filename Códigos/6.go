package main

import (
	"fmt"
)

func sumWordCounts(wordCountsList []map[string]int) map[string]int {
	result := make(map[string]int)

	for _, wordCounts := range wordCountsList {
		for word, count := range wordCounts {
			result[word] += count
		}
	}

	return result
}

func main() {
	wordCountsList := []map[string]int{
		{"apple": 1, "banana": 2},
		{"apple": 1, "cherry": 4},
		{"banana": 1, "date": 2},
	}

	totalCounts := sumWordCounts(wordCountsList)

	fmt.Println("Contagem total das palavras:")
	for word, count := range totalCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
