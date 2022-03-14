package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wahyd4/mryum/calculator"
)

var (
	filePath  = getEnvWithDefault("INPUT", "./test3.txt")
	topNumber = getEnvWithDefault("TOPX", "3")
)

func main() {
	topNumber, err := strconv.Atoi(topNumber)
	if err != nil {
		fmt.Println("TOPX must be a number")
	}
	topNumbers, err := calculator.CalculateLargestNumbers(filePath, topNumber)
	if err != nil {
		fmt.Println("Oops: " + err.Error())
		os.Exit(1)
	}
	fmt.Printf("Top %d numbers for file %s are: %v", topNumber, filePath, topNumbers)
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
