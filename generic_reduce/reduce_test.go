package reducers

import (
	"slices"
	"testing"
)

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			To:     "John Cena",
			From:   "Nicholas Cage",
			Amount: 9000,
		},
		{
			To:     "Nicholas Cage",
			From:   "Guy",
			Amount: 42,
		},
	}
	AssertEqual(t, BalanceFor(transactions, "John Cena"), 9000)
	AssertEqual(t, BalanceFor(transactions, "Nicholas Cage"), -8958)
	AssertEqual(t, BalanceFor(transactions, "Guy"), -42)

}

func AssertEqual[T comparable](t *testing.T, i1, i2 T) {
	t.Helper()
	if i1 != i2 {
		t.Errorf("expected %v and %v to be equal", i1, i2)
	}
}
func TestSum(t *testing.T) {
	t.Run("collection of variable size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	assertSlicesEqual := func(t *testing.T, s1, s2 []int) {
		t.Helper()
		if !slices.Equal(s1, s2) {
			t.Errorf("got %v want %v", s1, s2)
		}
	}
	t.Run("sum two slices of variable size", func(t *testing.T) {
		s1 := []int{0, 1, 1, 2, 3, 5, 8, 13}
		s2 := []int{2, 4, 8, 16, 32}

		got := SumAllTails(s1, s2)

		want := []int{33, 60}
		assertSlicesEqual(t, got, want)
	})
	t.Run("safely sum empty slice", func(t *testing.T) {
		emptySlice := []int{}
		got := SumAllTails(emptySlice)
		want := []int{0}

		assertSlicesEqual(t, got, want)
	})
}
