package recursion

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestRecursionQuickSort(t *testing.T) {

	result := quickSort([]int{1, 7, 9, 4, 3})
	want := []int{1, 3, 4, 7, 9}
	if !slices.Equal(result, want) {
		t.Errorf("quickSort want '%v', got '%v';", want, result)
	}

}
