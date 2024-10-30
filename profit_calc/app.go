package main

import (
	"errors"
	"fmt"
	"os"
)

const resultFile = "results.txt"

func main() {

	revenue, err := getUserInput("Revenue Amount: ")

	if err != nil {
		// return
		panic(err)
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		panic(err)
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		panic(err)
	}

	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	resultBeforeTax, resultAfterTax, ratio := calcProfit(revenue, expenses, taxRate)
	// var resultFromFile, err = readFile()

	fmt.Printf("%.1f\n", resultBeforeTax)
	fmt.Printf("%.1f\n", resultAfterTax)
	fmt.Printf("%.1f\n", ratio)
	writeFile(resultBeforeTax, resultAfterTax, ratio)
}

func writeFile(ebt, profit, ratio float64) {
	outputText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(resultFile, []byte(outputText), 0644)
}

// func readFile() (float64, error) {
// 	data, err := os.ReadFile(resultFile)

// 	if err != nil {
// 		return 0, errors.New("failed to read file")
// 	}

// 	outputText := string(data)
// 	output, err := strconv.ParseFloat(outputText, 64)

// 	if err != nil {
// 		return 0, errors.New("failed to convert file")
// 	}
// 	return output, nil
// }

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("value must be above 0")
	}
	return userInput, nil
}

func calcProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	resultBeforeTax := revenue - expenses
	resultAfterTax := resultBeforeTax * (1 - taxRate/100)
	var result float64 = resultBeforeTax / resultAfterTax

	return resultBeforeTax, resultAfterTax, result
}
