package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRates float64

	fmt.Print("Enter your Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your Tax Rates: ")
	fmt.Scan(&taxRates)

	ebt := revenue - expenses

	profit := ebt * (1 - taxRates/100)

	ratio := ebt/profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
