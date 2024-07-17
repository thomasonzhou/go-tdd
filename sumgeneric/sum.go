package main

func Reduce[T any](start T, iter []T, reducerFunc func(T, T) T) T {
	var total = start
	for _, i := range iter {
		total = reducerFunc(total, i)
	}
	return total
}

func Sum(numbers []int) (sum int) {
	addInt := func(n1, n2 int) int {
		return n1 + n2
	}
	return Reduce(0, numbers, addInt)
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	sumTail := func(sums, currArr []int) []int {
		if len(currArr) == 0 {
			return append(sums, 0)
		} else {
			tail := currArr[1:]
			return append(sums, Sum(tail))
		}
	}
	return Reduce(sums, numbersToSum, sumTail)
}
