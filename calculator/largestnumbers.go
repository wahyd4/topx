package calculator

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func calculateIncomingNumber(numbers *LargestNumbersHeap, newNumber int) *LargestNumbersHeap {
	smallest := heap.Pop(numbers).(int)

	if smallest < newNumber {
		heap.Push(numbers, newNumber)
	} else {
		heap.Push(numbers, smallest)
	}

	return numbers
}

func CalculateLargestNumbers(filePath string, topX int) ([]int, error) {
	if topX <= 0 {
		return []int{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return []int{}, fmt.Errorf("Could not open file %s with error %w", filePath, err)
	}
	scanner := bufio.NewScanner(file)

	largestNumbersHeap := LargestNumbersHeap{}

	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err != nil {
			return []int{}, fmt.Errorf("fail to parse text %s with error %w", text, err)
		}
		if len(largestNumbersHeap) < topX {
			largestNumbersHeap = append(largestNumbersHeap, number)
			heap.Init(&largestNumbersHeap)
		} else {
			calculateIncomingNumber(&largestNumbersHeap, number)
		}
	}
	heap.Init(&largestNumbersHeap)
	return largestNumbersHeap, nil
}
