package main

import (
	"fmt"
	"log"
)

func main() {

	var revenue, expenses, taxRate float64

	captureInput("Revenue", &revenue)
	captureInput("Expenses", &expenses)
	captureInput("Tax rate", &taxRate)

	taxRate = taxRate / 100
	ebt := calcEBT(revenue, expenses)
	profit := calcProfit(ebt, taxRate)
	ratio := calcRatio(ebt, profit)

	printResults(ebt, profit, ratio)

}

func printResults(ebt float64, profit float64, ratio float64) {
	fmt.Printf(`Calculated:
EBT: %g
Profit: %g
Ratio: %g`, ebt, profit, ratio)
}

func calcRatio(ebt float64, profit float64) float64 {
	return ebt / profit
}

func calcProfit(ebt float64, taxRate float64) float64 {
	return ebt * (1 - taxRate)
}

func calcEBT(revenue float64, expenses float64) float64 {
	return revenue - expenses
}

func captureInput(msg string, p any) {
	fmt.Print(msg, ": ")
	_, err := fmt.Scan(p)
	if err != nil {
		log.Fatal(err)
	}
}
