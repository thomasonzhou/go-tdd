package main

type reducer func(n1, n2 int) int

func Reduce(start int, iter []int, reducerFunc reducer) int {
	total := start
	for _, i := range iter {
		total = reducerFunc(total, i)
	}
	return total
}

func Sum(numbers []int) (sum int) {
	var addInt = func(n1, n2 int) int {
		return n1 + n2
	}
	return Reduce(0, numbers, addInt)
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
