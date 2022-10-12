package recursion

import (
	"testing"

	"github.com/gogogo1024/algorithmsInGolang/util"
)

func TestRecursionQuickSort(t *testing.T) {

	result := quickSort([]int{1, 7, 9, 4, 3})
	want := []int{1, 3, 4, 7, 9}
	if !util.CompareEqualSlice(result, want) {
		t.Errorf("quickSort want '%v', got '%v';", want, result)
	}

}
