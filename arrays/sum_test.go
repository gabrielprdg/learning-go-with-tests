package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8, 3}
		got := Sum(numbers)
		want := 23
		if got != want {
			t.Errorf("got %d but want %d and given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{2, 3, 4}
		got := Sum(numbers)
		want := 9
		if got != want {
			t.Errorf("got %d but want %d and given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 4}, []int{1, 3})
	want := []int{7, 4}

	// Go does not let you use equality operators with slices
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %d but got %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Want %d, but got %d", want, got)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 5, 6}, []int{0, 2, 2})
		want := []int{11, 4}

		checkSum(t, want, got)
	})

	t.Run("make the sum of an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 1, 7})
		want := []int{0, 8}

		checkSum(t, want, got)
	})
}
