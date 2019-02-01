package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum array 1-10", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := 55
		got := Sum(arr)

		assertCorrectAnswer(t, got, want)
	})

	t.Run("sum arr [20 20]", func(t *testing.T) {
		arr := []int{20, 20}
		got := Sum(arr)
		want := 40

		assertCorrectAnswer(t, got, want)
	})
}

func assertCorrectAnswer(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Expected '%d'; got '%d''", want, got)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("2-arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertCorretAnswerSlices(t, got, want)
	})

	t.Run("3-arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
		want := []int{6, 15, 24}
		assertCorretAnswerSlices(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("test-2-slices", func(t *testing.T) {
		got := SumAllTails([]int{0, 1, 2, 3}, []int{0, 1, 2, 3})
		want := []int{6, 6}
		assertCorretAnswerSlices(t, got, want)
	})

	t.Run("test-1-slice", func(t *testing.T) {
		got := SumAllTails([]int{0, 1, 2})
		want := []int{3}
		assertCorretAnswerSlices(t, got, want)
	})

	t.Run("test-3-arrays", func(t *testing.T) {
		got := SumAllTails([]int{0, 2}, []int{0, 9, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		want := []int{2, 12, 54}
		assertCorretAnswerSlices(t, got, want)
	})

	t.Run("empty-slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 4, 5})
		want := []int{0, 9}
		assertCorretAnswerSlices(t, got, want)
	})
}

func assertCorretAnswerSlices(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("go: %v; want %v", got, want)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
