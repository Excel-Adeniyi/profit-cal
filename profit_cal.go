package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRates float64

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRates := getUserInput("Tax Rates: ")

	ebt, profit, ratio := calFinancials(revenue, expenses, taxRates)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Printf("Ratio to 2dp: %.2f\n", ratio)
}

func getUserInput(infoText string) float64 {

	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calFinancials(uRevenue, uExpenses, uTaxRates float64) (ebt, profit, ratio float64) {

	ebt = uRevenue - uExpenses

	profit = ebt * (1 - uTaxRates/100)

	ratio = ebt / profit

	return ebt, profit, ratio
}
