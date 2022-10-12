package recursion

import (
	"testing"
)

func TestRecursionSum(t *testing.T) {

	result := sum([]int{1, 2, 3, 4, 5})

	if result != 15 {
		t.Errorf("sum function want '%v', got '%v';", 15, result)
	}

}
