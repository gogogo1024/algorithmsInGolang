package recursion

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233...

// F(1) = 0
// F(2) = 1
// F(3) = 1
// F(4) = 2
// F(5) = 3
// ...
// F(n) = F(n-1) + F(n-2) (n > 2)

// 利用闭包缓存计算过的斐波那契函数
func wrapFibonacci() func(t int) int {
	const MAX = 50
	var fibs [MAX]int
	// 必须要先命名
	var fibonacci func(int) int

	fibonacci = func(n int) int {
		if n == 1 {
			return 0
		}
		if n == 2 {
			return 1
		}
		index := n - 1
		// 不为零说明该位置计算过
		if fibs[index] != 0 {
			return fibs[index]
		}
		num := fibonacci(n-1) + fibonacci(n-2)
		fibs[index] = num
		return num
	}
	return fibonacci

}
