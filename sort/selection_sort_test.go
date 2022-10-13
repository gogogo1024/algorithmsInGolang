package sort

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSelectSort(t *testing.T) {

	ar1 := []int{35, 61, 23, 14, 175}
	sorted := selectionSort(ar1)
	fmt.Printf("%d", sorted)
	wanted := []int{14, 23, 35, 61, 175}
	if !slices.Equal(sorted, wanted) {
		t.Errorf("selectionSort on '%v'; want '%v', got '%v';", ar1, wanted, sorted)
	}
}
