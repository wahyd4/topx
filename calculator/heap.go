package calculator

type LargestNumbersHeap []int

func (biggestNumbers LargestNumbersHeap) Less(first, second int) bool {
	// pop the smallest number
	return biggestNumbers[first] < biggestNumbers[second]
}

func (biggestNumbers LargestNumbersHeap) Len() int {
	return len(biggestNumbers)
}

func (biggestNumbers LargestNumbersHeap) Swap(first, second int) {
	biggestNumbers[first], biggestNumbers[second] = biggestNumbers[second], biggestNumbers[first]
}

func (biggestNumbers *LargestNumbersHeap) Push(x interface{}) {
	*biggestNumbers = append(*biggestNumbers, x.(int))
}

func (biggestNumbers *LargestNumbersHeap) Pop() interface{} {
	oldNumbers := *biggestNumbers
	arrayLength := len(oldNumbers)
	number := oldNumbers[arrayLength-1]
	*biggestNumbers = oldNumbers[0 : arrayLength-1]
	return number
}
