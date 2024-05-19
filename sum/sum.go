package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number

	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	arrayCount := len(numbersToSum)
	sums = make([]int, arrayCount)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
