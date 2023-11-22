package arrays

import (
	"reflect"
	"testing"
)

// testing with array
// func TestSum(t *testing.T) {

// 	numbers := [5]int{1,2,3,4,5}

// 	got := Sum(numbers)
// 	want := 15

// 	if got != want {
// 		t.Errorf("got %d want %d given, %v",got,want,numbers)
// 	}
// }

// testing with slice
func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSum(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSum(t, got, want)
	})
}

func TestCompareSlices(t *testing.T) {
	var slice1, slice2 []int
	t.Run("Compare two nil slcies", func(t *testing.T) {
		got := compareSlices(slice1, slice2)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	slice1 = []int{1, 2, 3}
	t.Run("Compare nil slice and slice", func(t *testing.T) {
		got := compareSlices(slice1, slice2)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	slice2 = []int{1, 2, 3}
	t.Run("Compare two slices", func(t *testing.T) {
		got := compareSlices(slice1, slice2)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
