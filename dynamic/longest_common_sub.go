package dynamic

// init an matrix to put the elements
func createMatrix(rows, cols int) [][]int {
	cell := make([][]int, rows)
	for i := range cell {
		cell[i] = make([]int, cols)
	}

	return cell
}

// 子字符串
// 比较大小是从a[i-1]和b[j-1], i是从1 到 len(a),j是1到len(b)
// 赋值是给cell[i][j],

 // 1. a[i-1]等于b[j-1]相等时，单元格的值等于左上角单元格的值加上1
 // (为什么是左上角：假设此时比较hi,fi,左上角对应是h和f,正好当前单元格就是hi和fi)
 // (为什么要加1：假设此时比较hi,fi,左上角对应是h和f,正好当前单元格就是hi和fi,hi的i和fi的i)

// hish ,fish的最大公共字串是ish，长度为3
func substring(a, b string) string {
	lcs := 0
	lastSubIndex := 0
	// 创建 len(a)+1  * len(b)+1 的矩阵
	cell := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1
				if cell[i][j] > lcs {
					lcs = cell[i][j]
					lastSubIndex = i
				}
			} else {
				cell[i][j] = 0
			}
		}
	}
	return a[lastSubIndex-lcs : lastSubIndex]
}

// fosh vs fish 最大公共子序列是fsh长度为3
// fosh vs fort 最大公共子序列是fo长度为2
// 填充cell[i][j]
// 当a[i-1]==b[j-1]时，和计算最长公共字串一样，等于左上方+1
// 当a[i-1]!=b[j-1]时, 选取左方cell[i][j-1]和上方cell[i-1][j]最大值

// 最大公共子序列
func subsequence(a, b string) int {
	// 创建 len(a)+1  * len(b)+1 的矩阵
	cell := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1
			} else {
				cell[i][j] = cell[i-1][j]
				if cell[i-1][j] < cell[i][j-1] {
					cell[i][j] = cell[i][j-1]
				}
			}
		}
	}
	return cell[len(a)][len(b)]
}