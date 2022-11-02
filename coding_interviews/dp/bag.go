package dp

import "github.com/crusj/algorithm/common"

// 背包问题
// 把前i件物品放入承重为j的背包最大值
// dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
type bagv struct {
	w int
	v int
}

// 0 1背包
func bag(m []*bagv, size int) int {
	dp := make([][]int, len(m)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, size+1)
	}

	for i := 1; i <= len(m); i++ {
		for j := 1; j <= size; j++ {
			if j >= m[i-1].w {
				dp[i][j] = common.Max(dp[i-1][j], dp[i-1][j-m[i-1].w]+m[i-1].v)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(m)][size]
}

// 完全背包
func bag2(m []*bagv, size int) int {
	dp := make([][]int, len(m)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, size+1)
	}

	for i := 1; i <= len(m); i++ {
		for j := 1; j <= size; j++ {
			if j >= m[i-1].w {
				dp[i][j] = common.Max(dp[i-1][j], dp[i][j-m[i-1].w]+m[i-1].v)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(m)][size]
}
