package sort

// 查询数组内最小的索引
func findSmallest(array []int) (a int) {
	smallest := array[0]
	smallestIndex := 0
	for i, v := range array {
		if v < smallest {
			smallestIndex = i
			smallest = v
		}
	}
	return smallestIndex
}

// 选择排序
//第一次需要检查n个元素，但随后检查的元素数量依次为:n - 1, n – 2, ..., 2和1
// 总和是:n(n+1)/2,所以平均每次检查的元素数为:1/2 × n，运行时间为:O(n × 1/2 × n)
func selectionSort(array []int) (a []int) {
	newArray := make([]int, len(array))
	for i := range array {
		smallestIndex := findSmallest(array)
		smallest := array[smallestIndex]

		// pop the smallest and keep the origin order
		// ref :https://yourbasic.org/golang/delete-element-slice/
		copy(array[smallestIndex:], array[smallestIndex+1:])
		array = array[:len(array)-1]

		newArray[i] = smallest
	}
	return newArray
}
