package main

import (
	"fmt"
)

func generateFibonacci(n int) map[int]int {
	fibonacci := make(map[int]int)

	a, b := 0, 1

	for i := 0; i <= n; i++ {
		fibonacci[i] = a
		a, b = b, a+b
	}

	return fibonacci
}

func main() {
	n := 15

	fibonacciSequence := generateFibonacci(n)

	fmt.Println("Sequência de Fibonacci até", n)
	for index, number := range fibonacciSequence {
		fmt.Printf("%d: %d\n", index, number)
	}
}
