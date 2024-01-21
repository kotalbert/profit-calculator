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

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)
	ratio := ebt / profit

	fmt.Printf("Calculated:\n"+
		"EBT: %g\n"+
		"Profit: %g\n"+
		"Ratio: %g\n",
		ebt, profit, ratio)

}

func captureInput(msg string, p any) {
	fmt.Print(msg, ": ")
	_, err := fmt.Scan(p)
	if err != nil {
		log.Fatal(err)
	}
}
