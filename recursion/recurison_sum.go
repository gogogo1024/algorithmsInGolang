package recursion

// 找到基线条件和递归条件
// 基线条件 arr长度为0
// 递归条件
// 递归计算和
func sum(arr []int) (t int) {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sum(arr[1:])
}
