package main

import (
	"fmt"
)

func calculateExpenses(debts map[string]float64) map[string]float64 {
	total := 0.0

	for _, amount := range debts {
		total += amount
	}

	equalExpense := total / float64(len(debts))

	result := make(map[string]float64)

	for person, amount := range debts {
		result[person] = equalExpense - amount
	}

	return result
}

func main() {
	debts := map[string]float64{
		"Cândido":  20.0,
		"Pedro":    2120.0,
		"EU mesmo": 1.0,
		"fiche":    320.0,
	}

	expenses := calculateExpenses(debts)

	fmt.Println("Valor a receber/pagar:")
	for person, amount := range expenses {
		fmt.Printf("%s: %.2f\n", person, amount)
	}
}
