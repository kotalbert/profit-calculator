package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	revenue, errRev := captureFloat("Revenue")
	if errRev != nil {
		log.Fatal(errRev)
	}

	expenses, errExp := captureFloat("Expenses")
	if errExp != nil {
		log.Fatal(errExp)
	}

	taxRate, errTax := captureFloat("Tax rate")
	if errTax != nil {
		log.Fatal(errTax)
	}

	taxRate = taxRate / 100
	ebt := calcEBT(revenue, expenses)
	profit := calcProfit(ebt, taxRate)
	ratio := calcRatio(ebt, profit)

	report := getResultsReport(ebt, profit, ratio)
	fmt.Println(report)
	writeReportToFile(report)
}

func writeReportToFile(report string) {
	err := os.WriteFile("profit_report.txt", []byte(report), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getResultsReport(ebt float64, profit float64, ratio float64) string {
	return fmt.Sprintf(`Calculated:
EBT: %.2f
Profit: %.2f
Ratio: %.2f`, ebt, profit, ratio)
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

func captureFloat(msg string) (float64, error) {
	fmt.Print(msg, ": ")
	var userInput float64
	_, err := fmt.Scan(&userInput)
	if err != nil {
		log.Fatal(err)
	}
	// handle negative input
	if userInput <= 0 {
		return -1, errors.New("expecting positive value")
	}
	return userInput, nil
}
