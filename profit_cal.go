package main

import (
	"errors"
	"fmt"
	"os"
)

func storeFinacialInfo(ebt float64, profit float64, ratio float64) {
	var data = fmt.Sprintf("EBT: %.1f\nProfif: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	os.WriteFile("profit.txt", []byte(data), 0644)
}
func main() {
	// var revenue float64
	// var expenses float64
	// var taxRates float64

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRates, err := getUserInput("Tax Rates: ")

	if err != nil {
		fmt.Println(err)
		return
	}
	ebt, profit, ratio := calFinancials(revenue, expenses, taxRates)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Printf("Ratio to 2dp: %.2f\n", ratio)
	storeFinacialInfo(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {

	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value can't be less or equal to 0")
	}
	return userInput, nil
}

func calFinancials(uRevenue, uExpenses, uTaxRates float64) (ebt, profit, ratio float64) {

	ebt = uRevenue - uExpenses

	profit = ebt * (1 - uTaxRates/100)

	ratio = ebt / profit

	return ebt, profit, ratio
}
