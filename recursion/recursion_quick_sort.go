package recursion

import (
	"math/rand"
	"time"
)

func quickSort(arr []int) (a []int) {
	if len(arr) < 2 {
		return arr
	} else {
		// 随机地选择用作基准值的元素。快速排序的平均运行时间为O(nlogn)。
		// 假设使用第一个元素作为基准值元素，对于已经原本就排好序的数组做快排平均时间为O(n*n)。

		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(len(arr))

		pivot := arr[randomNum]

		var less []int
		var greater []int

		for i := 0; i < len(arr); i++ {
			// 跳过自身
			if i == randomNum {
				continue
			}
			//基线条件:为空或只包含一个元素的数组是“有序”的
			if pivot <= arr[i] {
				greater = append(greater, arr[i])
			} else {
				less = append(less, arr[i])
			}
		}
		// quickSort(less) 对小于pivot的排序
		// quickSort(greater) 对大于pivot的排序
		// 合并
		return append(append(quickSort(less), pivot), quickSort(greater)...)
	}
}
