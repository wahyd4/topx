package main

import (
	"fmt"
	"os"

	"github.com/wahyd4/mryum/calculator"
)

const (
	filePath  = "./testx3.txt"
	topNumber = 4
)

func main() {
	topNumbers, err := calculator.CalculateLargestNumbers(filePath, topNumber)
	if err != nil {
		fmt.Println("Oops: " + err.Error())
		os.Exit(1)
	}
	fmt.Printf("Top %d numbers are: %v", topNumber, topNumbers)
}
