package main

import "fmt"

func main() {
	text := "abracadabra"

	charFrequency := Frequência(text)

	fmt.Println("Frequência dos caracteres:")
	for char, count := range charFrequency {
		fmt.Printf("%c: %d\n", char, count)
	}

}
func Frequência(text string) map[rune]int {
	frequency := make(map[rune]int)

	for _, char := range text {
		frequency[char]++
	}

	return frequency
}
