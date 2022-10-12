package introduction

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {

	ar1 := []int{1, 2, 3, 4, 5}
	find1, index1 := binarySearch(ar1, 1)
	if !find1 || index1 == -1 {
		t.Errorf("binarySearch on '%v'; want true, got '%v'; want index 0,got index '%v'", ar1, find1, index1)
	}
	ar2 := []int{1, 2, 3, 4, 5}
	find2, index2 := binarySearch(ar2, -1)
	if find2 || index2 != -1 {
		t.Errorf("binarySearch on '%v'; want true, got '%v'; want index 0,got index '%v'", ar2, find2, index2)
	}

}
