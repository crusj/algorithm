package leetcode

// dp(i,j) = 1 // j ==0 or j == row - 1
// dp(i,j) = dp(i-1, j-1) + dp(i -1, j)
// 只能自底向上，不能自顶向下
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
	}

	ans[0] = []int{1}
	ans[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ans[i][j] = 1
				continue
			}
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans[0:numRows]
}
