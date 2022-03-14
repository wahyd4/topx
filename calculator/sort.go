package calculator

import "sort"

func sortDesc(numbers []int) []int {
	sort.SliceStable(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	return numbers
}
