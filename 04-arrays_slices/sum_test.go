package arraysslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3,9}

	// Go does not let you use equality operators with slices.
	if !slices.Equal(got,want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

func TestFindMaxNumber(t *testing.T) {
	assert := func(t *testing.T,got, expected int, numbers []int) {
		t.Helper()
		if expected != got {
			t.Errorf("got %d expected %d, given: %v", got, expected, numbers)
		}
	}
	t.Run("get max number from slice", func(t *testing.T) {
		numbers := []int{3,6,4,1,0,8}
		expected := 8
		got := FindMaxNumber(numbers)
		assert(t, got, expected, numbers)
	})

	t.Run("get max number from empty slice", func (t *testing.T){
		numbers := []int{}
		expected := 0;
		got := FindMaxNumber(numbers)
		assert(t, got, expected, numbers)
	})
	
}

func TestSplitMiddle(t *testing.T) {
	compareSlices := func (t *testing.T, got, expected [2][]int, given []int) {
		t.Helper()
		for i := 0; i < 2; i++ {
			if !slices.Equal(got[i], expected[i]) {
				t.Errorf("got %v expected %v given %v", got, expected, given)
				break;
			}
		}
	}
	t.Run("Split Empty Slice", func(t *testing.T){
		slice := []int{}	
		got := SplitMiddle(slice)
		expected := [2][]int{}
		compareSlices(t, got, expected, slice)
	})

	t.Run("Split Even Slice", func(t *testing.T){
		slice := []int{1, 2, 3, 4}	
		got := SplitMiddle(slice)
		expected := [2][]int{
			{1, 2}, 
			{3, 4},
		}
		compareSlices(t, got, expected, slice)
	})

	t.Run("Split Odd Slice", func(t *testing.T){
		slice := []int{1, 2, 3, }	
		got := SplitMiddle(slice)
		expected := [2][]int{
			{1,}, 
			{2, 3},
		}
		compareSlices(t,got,expected, slice)
	})
}