package main

import (
	"fmt"
	"sort"
	"strings"
)

func getAnagramGroups(words []string) map[string][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		key := sortString(word)
		groups[key] = append(groups[key], word)
	}

	return groups
}

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func main() {
	words := []string{"listen", "silent", "elbow", "below", "state", "taste"}

	anagramGroups := getAnagramGroups(words)

	fmt.Println("Grupos de palavras anagramas:")
	for _, group := range anagramGroups {
		fmt.Println(group)
	}
}
