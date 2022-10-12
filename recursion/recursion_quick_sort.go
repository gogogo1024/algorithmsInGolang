package recursion

func quickSort(arr []int) (a []int) {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		var less []int
		var greater []int
		for i := 1; i < len(arr); i++ {
			//基线条件:为空或只包含一个元素的数组是“有序”的
			if pivot <= arr[i] {
				greater = append(greater, arr[i])
			} else {
				less = append(less, arr[i])
			}
		}
		// quickSort(less) 对小于pivot的排序
		// quickSort(less) 对大于pivot的排序
		// 合并
		return append(append(quickSort(less), pivot), quickSort(greater)...)
	}
}
