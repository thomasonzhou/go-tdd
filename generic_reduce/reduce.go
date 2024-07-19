package reducers

type Transaction struct {
	To, From string
	Amount   float64
}

func BalanceFor(transactions []Transaction, name string) (finalBalance float64) {
	balanceChange := func(currentBalance float64, t Transaction) float64 {
		if t.To == name {
			currentBalance += t.Amount
		}
		if t.From == name {
			currentBalance -= t.Amount
		}
		return currentBalance
	}
	return Reduce(0, transactions, balanceChange)
}

func Reduce[A, B any](start A, iter []B, reducerFunc func(A, B) A) A {
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
