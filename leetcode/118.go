package leetcode

// dp(i,j) = 1 // j ==0 or j == row - 1
// dp(i,j) = dp(i-1, j-1) + dp(i -1, j)
// 只能自底向上，不能自顶向下
// 严格来说不是动态规划，没有求最值。
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			ans[i] = []int{1}
			continue
		}
		if i == 1 {
			ans[i] = []int{1, 1}
			continue
		}

		ans[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ans[i][j] = 1
				continue
			}
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans
}
