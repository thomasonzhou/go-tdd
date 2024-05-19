package main

import (
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("sum two slices of variable size", func(t *testing.T) {
		s1 := []int{0, 1, 1, 2, 3, 5, 8, 13}
		s2 := []int{2, 4, 8, 16, 32}

		got := SumAll(s1, s2)

		want := []int{33, 62}
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v given, %v %v", got, want, s1, s2)
		}
	})
}
