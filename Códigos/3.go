package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	total := somaValores(m)

	fmt.Printf("Soma dos valores: %d\n", total)

}
func somaValores(m map[string]int) int {
	soma := 0

	for _, value := range m {
		soma += value
	}

	return soma
}
